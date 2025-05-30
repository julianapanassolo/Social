import React, { useState } from "react";
import $ from "jquery";
import Swal from "sweetalert2";

const Cadastro = () => {
  const [nome, setNome] = useState("");
  const [email, setEmail] = useState("");
  const [nick, setNick] = useState("");
  const [senha, setSenha] = useState("");
  const [confimarSenha, setConfimarSenha] = useState("");

  const handleSubmit = (evento: { preventDefault: () => void; }) => {
    evento.preventDefault();
    if (senha !== confimarSenha) {
      Swal.fire("Ops...", "As senhas devem ser iguais", "error");
      return;
    }

    $.ajax({
      url: "/usuarios",
      method: "POST",
      data: {
        nome: nome,
        email: email,
        nick: nick,
        senha: senha,
      },
    })
      .done(function () {
        Swal.fire(
          "Sucesso!",
          "Usuário cadastrado com sucesso!",
          "success"
        ).then(function () {
          $.ajax({
            url: "/login",
            method: "POST",
            data: {
              email: email,
              senha: senha,
            },
          })
            .done(function () {
              window.location.replace("/home");
            })
            .fail(function () {
              Swal.fire("Ops...", "Erro ao autenticar o usuário", "error");
            });
        });
      })
      .fail(function () {
        Swal.fire("Ops...!", "Erro ao cadastrar o usuário!", "error");
      });
  };

  return (
    <form onSubmit={handleSubmit}>
      <label>
        Nome:
        <input
          type="text"
          value={nome}
          onChange={(evento: { target: { value: any; }; }) => setNome(evento.target.value)}
        />
      </label>
      <label>
        Email:
        <input
          type="email"
          value={email}
          onChange={(evento: { target: { value: any; }; }) => setEmail(evento.target.value)}
        />
      </label>
      <label>
        Nick:
        <input
          type="text"
          value={nick}
          onChange={(evento: { target: { value: any; }; }) => setNick(evento.target.value)}
        />
      </label>
      <label>
        Senha:
        <input
          type="password"
          value={senha}
          onChange={(evento: { target: { value: any; }; }) => setSenha(evento.target.value)}
        />
      </label>
      <label>
        Confirmar Senha:
        <input
          type="password"
          value={confimarSenha}
          onChange={(evento: { target: { value: any; }; }) => setConfimarSenha(evento.target.value)}
        />
      </label>
      <button type="submit">Cadastrar</button>
    </form>
  );
};

export default Cadastro;

