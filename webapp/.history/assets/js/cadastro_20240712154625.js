$('#').on('submit', criarUsuario);

function criarUsuario(evento) {
    evento.preventDefaut();
    console.log("Dentro da função usuário!")
}