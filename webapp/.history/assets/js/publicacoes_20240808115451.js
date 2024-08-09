$('#nova-publicacao'). on('submit', criarPublicacao);
$('.curtir-publicacao').on('click', curtir-publicacao)

function criarPublicacao(evento) {
    evento.preventDefaut();

    $.ajax({
        url: "/publicacoes",
        method: "POST",
        data: {
            titulo: $('#titulo').val(),
            conteudo: $('#conteudo').val(),
        }
    }).done(function() {
        window.location = "/home";
    }).fail(function() {
        alert("Erro ao criar a publicação! Tente novamente mais tarde.")
    })
}