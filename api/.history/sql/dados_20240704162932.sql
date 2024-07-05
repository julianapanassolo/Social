INSERT INTO usuarios (nome, nick, email, senha)
VALUES
("Usuário 1", "usuario_1", "usuario1@gmail.com", "$2a$10$fSh9PfkbQB3lhvdiDLcPk.BEq9xIghgkT0KuJapzanvNLu0221cD6"), -- usuario1
("Usuário 2", "usuario_2", "usuario2@gmail.com", "$2a$10$fSh9PfkbQB3lhvdiDLcPk.BEq9xIghgkT0KuJapzanvNLu0221cD6"), -- usuario 2
("Usuário 3", "usuario_3", "usuario3@gmail.com", "$2a$10$fSh9PfkbQB3lhvdiDLcPk.BEq9xIghgkT0KuJapzanvNLu0221cD6"); -- usuario3

INSERT INTO seguidores(usuario_id, seguidores_id)
VALUES
(1, 2),
(3, 1),
(1, 3);

INSERT INTO publicacoes(titulo, conteudo, autor_id)
VALUES
("Publicação do usuário 1", "Essa é a publicação do usuário 1! Que Bom!", 1)
("Publicação do usuário 2", "Essa é a publicação do usuário 1! Que Bom!", 1)
("Publicação do usuário 1", "Essa é a publicação do usuário 1! Que Bom!", 1)