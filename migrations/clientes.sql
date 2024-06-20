CREATE TABLE clientes (
    id SERIAL PRIMARY KEY,
    nome VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL,
    telefone VARCHAR(20) NOT NULL,
    morada TEXT NOT NULL,
    doencas TEXT,
    outros TEXT
);

INSERT INTO clientes (nome, email, telefone, morada, doencas, outros)
VALUES
    ('João Silva', 'joao.silva@example.com', '123456789', 'Rua A, 123', 'Hipertensão', 'Sem outros'),
    ('Maria Santos', 'maria.santos@example.com', '987654321', 'Rua B, 456', 'Diabetes', 'Sem outros'),
    ('Pedro Oliveira', 'pedro.oliveira@example.com', '111111111', 'Rua C, 789', NULL, 'Alérgico a penicilina'),
    ('Ana Pereira', 'ana.pereira@example.com', '222222222', 'Rua D, 101', 'Asma', 'Sem outros'),
    ('Carlos Costa', 'carlos.costa@example.com', '333333333', 'Rua E, 121', 'Hipertensão', 'Sem outros');