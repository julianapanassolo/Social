import { event } from "jquery";
import Swal, { SweetAlertResult } from "sweetalert2";

$('#nova-publicacao').on('submit', criarPublicacao);

$(document).on('click', '.curtir-publicacao', curtirPublicacao);
$(document).on('click', '.descurtir-publicacao', descurtirPublicacao);

$('#atualizar-publicacao').on('click', atualizarPublicacao);
$('deletar-publicacao').on('click', deletarPublicacao);


function criarPublicacao(evento: { preventDefault: () => void; }) {
    evento.preventDefault();

    $.ajax({
        url: "/publicacoes",
        method: "POST",
        data: {
            titulo: $('#titulo').val(),
            conteudo: $('#conteudo').val(),
        }
    }).done(function () {
        window.location.replace("/home");
    }).fail(function () {
        Swal.fire("Ops...", "Erro ao criar a publicação!", "error");
    });
}

function curtirPublicacao(evento: { preventDefault: () => void; target: any; }) {
    evento.preventDefault();

    const elementoClicado = $(evento.target);
    const publicacaoId = elementoClicado.closest('div').data('publicacao-id');



    elementoClicado.prop('disabled', true);
    $.ajax({
        url: `/publicacoes/${publicacaoId}/curtir`,
        method: "POST"
    }).done(function () {
        const contadorDeCurtidas = elementoClicado.next('span');
        const quantidadeDeCurtidas = parseInt(contadorDeCurtidas.text());

        contadorDeCurtidas.text(quantidadeDeCurtidas + 1);

        elementoClicado.addClass('descurtir-publicacao');
        elementoClicado.addClass('text-danger');
        elementoClicado.removeClass('curtir-publicacao');

    }).fail(function () {
        Swal.fire("Ops...", "Erro ao curtir a publicação!", "error");
    }).always(function () {
        elementoClicado.prop('disabled', false);
    });

}

function descurtirPublicacao(evento: { preventDefault: () => void; target: any; }) {
    evento.preventDefault();

    const elementoClicado = $(evento.target);
    const publicacaoId = elementoClicado.closest('div').data('publicacao-id');

    elementoClicado.prop('disabled', true);
    $.ajax({
        url: `/publicacoes/${publicacaoId}/descurtir`,
        method: "POST"
    }).done(function () {
        const contadorDeCurtidas = elementoClicado.next('span');
        const quantidadeDeCurtidas = parseInt(contadorDeCurtidas.text());

        contadorDeCurtidas.text(quantidadeDeCurtidas - 1);

        elementoClicado.removeClass('descurtir-publicacao');
        elementoClicado.removeClass('text-danger');
        elementoClicado.addClass('curtir-publicacao');

    }).fail(function () {
        Swal.fire("Ops...", "Erro ao descurtir a publicação!", "error");
    }).always(function () {
        elementoClicado.prop('disabled', false);
    });
}

function atualizarPublicacao(this: any) {
    $(this).prop('disabled', true);

    const publicacaoId = $(this).data('publicacao-id');


    $.ajax({
        url: `/publicacoes/${publicacaoId}`,
        method: "PUT",
        data: {
            titulo: $('#titulo').val(),
            conteudo: $('#conteudo').val()
        }
    }).done(function () {
        Swal.fire({
            title: "Sucesso!",
            text: "Publicação criada com sucesso!",
            icon: "success"
        }).then(function () {
            window.location.replace("/home");
        })

    }).catch(function () {
        Swal.fire("Ops...", "Erro ao editar a publicação!", "error");
    }).always(function () {
        $('#atualizar-publicacao').prop('disabled', false);
    })
}

function deletarPublicacao(evento: { preventDefault: () => void; target: any }) {
    evento.preventDefault();

    Swal.fire({
        title: "Atenção!",
        text: "Tem certeza que deseja excluir essa publicação? Essa ação é irreversível!",
        showCancelButton: true,
        cancelButtonText: "Cancelar",
        icon: "warning"
    }).then((confirmacao: SweetAlertResult<any>) => {
        if (!confirmacao.value) return;

        const elementoClicado = $(evento.target);
        const publicacao = elementoClicado.closest('div')
        const publicacaoId = publicacao.data('publicacao-id');

        elementoClicado.prop('disabled', true);

        $.ajax({
            url: `/publicacoes/${publicacaoId}`,
            method: "DELETE"
        }).done(function () {
            publicacao.fadeOut("slow", function () {
                $(this).remove();
            });
        }).fail(function () {
            Swal.fire("Ops...", "Erro ao excluir a publicação!", "error");
        });

    })

}