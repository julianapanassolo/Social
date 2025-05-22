import Swal from "sweetalert2";

$('#login').on('submit', fazerLogin);

function fazerLogin(evento: { preventDefault: () => void; }) {
    evento.preventDefault();

    $.ajax({
        url: `/login`,
        method: "POST",
        data: {
            email: $('#email').val(),
            senha: $('#senha').val(),
        }
    }).done(function () {
        window.location.replace("/home");
    }).fail(function () {
        Swal.fire("Ops...", "Usuário ou senha incorretos!", "error");
    });
}