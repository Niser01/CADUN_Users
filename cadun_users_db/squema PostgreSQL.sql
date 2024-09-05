-- Habilitar la extensión pgcrypto para usar digest() en las contraseñas
CREATE EXTENSION IF NOT EXISTS pgcrypto;

-- No es necesario 'USE' en PostgreSQL

DROP TABLE IF EXISTS REQUEST;
DROP TABLE IF EXISTS REQUEST_TYPES;
DROP TABLE IF EXISTS USERS_PROFILE;

-- Crear tabla REQUEST_TYPES
CREATE TABLE REQUEST_TYPES (
    id SERIAL PRIMARY KEY,  -- SERIAL reemplaza AUTO_INCREMENT
    Status VARCHAR(100) NOT NULL
);

-- Crear tabla USERS_PROFILE
CREATE TABLE USERS_PROFILE (
    id SERIAL PRIMARY KEY,
    names VARCHAR(100) NOT NULL,
    lastNames VARCHAR(100) NOT NULL,
    alias VARCHAR(100) NOT NULL,
    password VARCHAR(100) NOT NULL,
    eMail VARCHAR(100) NOT NULL,
    phoneNumber VARCHAR(100) NOT NULL,
    country VARCHAR(100) NOT NULL,
    home_address VARCHAR(100) NOT NULL
);

-- Crear tabla REQUEST
CREATE TABLE REQUEST (
    id SERIAL PRIMARY KEY,
    idUser INT NOT NULL,
    request_status INT NOT NULL,
    IAM_URL VARCHAR(255),
    PDF_URL VARCHAR(255),
    QUOTE_PDF_URL VARCHAR(255),
    FOREIGN KEY (idUser) REFERENCES USERS_PROFILE(id),
    FOREIGN KEY (request_status) REFERENCES REQUEST_TYPES(id)
);

-- Insertar datos en REQUEST_TYPES
INSERT INTO REQUEST_TYPES (Status)
VALUES ('Finalizado'),
       ('Cotización'),
       ('Rechazado'),
       ('Aceptado'),
       ('Iniciado');

-- Insertar datos en USERS_PROFILE con SHA-256 para las contraseñas usando digest()
INSERT INTO USERS_PROFILE (names, lastNames, alias, password, eMail, phoneNumber, country, home_address)
VALUES ('Pepe', 'Rodriguez', 'Pepe123', encode(digest('password123', 'sha256'), 'hex'), 'pepe@unal.edu.co', '3150000000', 'Colombia', 'Cra. 1 #1-1'),
       ('Juan', 'Torres', 'Juan123', encode(digest('password456', 'sha256'), 'hex'), 'juan@unal.edu.co', '3150000001', 'Colombia', 'Cra. 1 #1-1'),
       ('Ana', 'Cortez', 'Ana123', encode(digest('password789', 'sha256'), 'hex'), 'ana@unal.edu.co', '3150000002', 'Colombia', 'Cra. 1 #1-1');
