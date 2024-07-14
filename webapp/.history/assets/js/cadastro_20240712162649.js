$('#formulario-cadastro').on('submit', criarUsuario);

function criarUsuario(evento) {
    evento.preventDefault();
    console.log("Dentro da função usuário!")

    if (senha != confirmar) {
        alert("As senhas devem ser iguais")
    }
}