insert into usuarios (nome, nick, email, senha)
VALUES
("Usuário 1", "usuario_1", "usuario1@gmail.com", "$2a$10$fSh9PfkbQB3lhvdiDLcPk.BEq9xIghgkT0KuJapzanvNLu0221cD6"), -- usuario1
("Usuário 2", "usuario_2", "usuario2@gmail.com", "$2a$10$4/RUfIjGmYe.RSmSjV//YeuaK.nvzD.KvAOgEXZZb8JVlKrs25w2a"), -- usuario 2
("Usuário 3", "usuario_3", "usuario3@gmail.com", "$2a$10$7KG1QKSfshoJJm7m/Fk19Oyd4IjLWlMQAdwCXU1pDyeCQO1BUGbD."), -- usuario3

insert into seguidores(usuario_id, seguidores_id)
VALUES
(1, 2);
(3, 1);
(1, 3);