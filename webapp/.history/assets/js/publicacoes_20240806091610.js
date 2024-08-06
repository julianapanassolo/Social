$('#nova-publicacao'). on('submit', criarPublicacao);

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
        alert("Publicação criada com sucesso!");
    }).fail(function() {
        alert("Erro ao criar a publicação! Tente novamente mais tarde.")
    })
}