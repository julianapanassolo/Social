// Script para criação de nova publicação
$('#nova-publicacao').on('submit', criarPublicacao);
// Script para curtir publicações
$('.curtir-publicacao').on('click', curtirPublicacao);


function criarPublicacao(evento) {
    evento.preventDefault();

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
        alert("Erro ao criar a publicação! Tente novamente mais tarde.");
    });
}

function curtirPublicacao(evento) {
    console.log("Curtindo publicação...");
    
}
