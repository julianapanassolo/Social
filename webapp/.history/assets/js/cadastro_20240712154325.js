$('#formulario-cadastro').on('submit', criarUsuario);

function criarUsuario(evento) {
    evento.prevent
    console.log("Dentro da função usuário!")
}