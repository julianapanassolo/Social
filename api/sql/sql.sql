CREATE DATABASE IF NOT EXISTS social;
USE social;

DROP TABLE IF EXISTS publicacoes;
DROP TABLE IF EXISTS seguidores;
DROP TABLE IF EXISTS usuarios;

CREATE TABLE usuarios(
    id INT AUTO_INCREMENT PRIMARY KEY,
    nome VARCHAR(50) NOT NULL,
    nick VARCHAR(50) NOT NULL UNIQUE,
    email VARCHAR(50) NOT NULL UNIQUE,
    senha VARCHAR(255) NOT NULL, -- Aumentei o tamanho para armazenar hashes
    criadoEm TIMESTAMP DEFAULT CURRENT_TIMESTAMP()
 ) ENGINE=InnoDB;

 CREATE TABLE seguidores(
    usuario_id INT NOT NULL,
    seguidor_id INT NOT NULL,
    FOREIGN KEY fk_seguidores_usuario(usuario_id) -- Nome explícito da FK
    REFERENCES usuarios(id)
    ON DELETE CASCADE,
    FOREIGN KEY fk_seguidores_seguidor(seguidor_id) -- Nome explícito da FK
    REFERENCES usuarios(id)
    ON DELETE CASCADE,
    PRIMARY KEY(usuario_id, seguidor_id)
 ) ENGINE=InnoDB;

 CREATE TABLE publicacoes(
   id INT AUTO_INCREMENT PRIMARY KEY,
   titulo VARCHAR(50) NOT NULL,
   conteudo TEXT NOT NULL, -- Use TEXT para conteúdos maiores
   autor_id INT NOT NULL,
   FOREIGN KEY fk_publicacoes_autor(autor_id) -- Nome explícito da FK
   REFERENCES usuarios(id)
   ON DELETE CASCADE,
   curtidas INT DEFAULT 0,
   criadaEm TIMESTAMP DEFAULT CURRENT_TIMESTAMP()
 ) ENGINE=InnoDB;

-- Inserção de dados (mantendo a senha hasheada)
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
("Publicação testes","Essa é a publicação testes! Que Oba!", 1), -- Corrigi o ID para corresponder aos usuários inseridos
("Outra publicação","Conteúdo desta outra publicação.", 2),
("Mais um teste","Testando mais uma vez o sistema de publicações.", 3);