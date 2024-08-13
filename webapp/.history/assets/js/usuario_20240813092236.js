$('#para-de-seguir').on('click', pararDeSeguir);
$('#seguir').on('click', seguir);

function pararDeSeguir() {
    const usuarioId = $(this).data('usuario-id');
    $(this).prop('disabled', true);

    $.ajax({
        url: `usuarios/${usuarioId}/parar-de-seguir`,
        method: "POST"
    }).done
}

function seguir() {

}