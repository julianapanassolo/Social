$('#formulario-cadastro').on('submit', criarUsuario);

function criarUsuario(evento) {
    evento.preventDefault();
    console.log("Dentro da função usuário!")

    if ($('#senha').val() != $('#confimar-senha').val()) {
        Swal.fire("Ops...", "As senhas devem ser iguais", "error")        
            return
    }

    $.ajax({
        url: "/usuarios",
        method: "POST",
        data: {
            nome: $('#nome').val(),
            email: $('#email').val(),
            nick: $('#nick').val(),
            senha: $('#senha').val()
        }
    }).done(function() {
        alert("Usuário cadastrado com sucesso!");
    }).fail(function(erro) {
        console.log(erro);
        alert("Erro ao cadastrar o usuário!")
    });
}