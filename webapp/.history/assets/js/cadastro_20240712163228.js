$('#formulario-cadastro').on('submit', criarUsuario);

function criarUsuario(evento) {
    evento.preventDefault();
    console.log("Dentro da função usuário!")

    if ($('#senha').v != $('#confimar-senha')) {
        alert("As senhas devem ser iguais")
            return
    }
}