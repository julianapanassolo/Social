INSERT INTO usuarios (nome, nick, email, senha)
VALUES
("Usuário 1", "usuario_1", "usuario1@gmail.com", "$2a$10$fSh9PfkbQB3lhvdiDLcPk.BEq9xIghgkT0KuJapzanvNLu0221cD6"), -- usuario1
("Usuário 2", "usuario_2", "usuario2@gmail.com", "$2a$10$fSh9PfkbQB3lhvdiDLcPk.BEq9xIghgkT0KuJapzanvNLu0221cD6"), -- usuario 2
("Usuário 3", "usuario_3", "usuario3@gmail.com", "$2a$10$fSh9PfkbQB3lhvdiDLcPk.BEq9xIghgkT0KuJapzanvNLu0221cD6"); -- usuario3

INSERT INTO seguidores(usuario_id, seguidor_id)
VALUES
(1, 2),
(3, 1),
(1, 3);


INSERT INTO publicacoes(titulo, conteudo, autor_id)
VALUES
("Publicação testes 1","Essa é a publicação testes 4! Que Bom!", 13),
("Publicação testes 2","Essa é a publicação testes 3! Que Bom!", 6),
("Publicação testes 4","Essa é a publicação testes 2! Que Bom!", 2);