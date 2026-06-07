  CREATE TABLE enderecos (
      ID          VARCHAR(100) NOT NULL PRIMARY KEY,
      cep         VARCHAR(10)  NOT NULL UNIQUE,
      logradouro  VARCHAR(255),
      complemento VARCHAR(255),
      unidade     VARCHAR(50),
      bairro      VARCHAR(100),
      localidade  VARCHAR(100),
      uf          CHAR(2),
      estado      VARCHAR(100),
      regiao      VARCHAR(50),
      ddd         VARCHAR(5)
  );

  CREATE TABLE usuarios (
      ID              VARCHAR(100) NOT NULL PRIMARY KEY,
      nome            VARCHAR(150) NOT NULL,
      email           VARCHAR(200) NOT NULL UNIQUE,
      senha           VARCHAR(100) NOT NULL,
      genero          CHAR(10),
      CPF             CHAR(11)     NOT NULL UNIQUE,
      data_nascimento DATE,
      endereco_id     VARCHAR(100),

      CONSTRAINT fk_usuarios_enderecos
          FOREIGN KEY (endereco_id)
          REFERENCES enderecos(ID)
          ON DELETE SET NULL
  );

  CREATE TABLE academias (
      ID   VARCHAR(100) NOT NULL PRIMARY KEY,
      nome VARCHAR(100) NOT NULL UNIQUE,
      CNPJ VARCHAR(14)  NOT NULL UNIQUE
  );

  CREATE TABLE professores (
      id_professor          VARCHAR(100) NOT NULL PRIMARY KEY,
      id_academia_professor VARCHAR(100),
      id_usuario_professor  VARCHAR(100),

      CONSTRAINT fk_professores_academias
          FOREIGN KEY (id_academia_professor)
          REFERENCES academias(ID)
          ON DELETE SET NULL,

      CONSTRAINT fk_professores_usuarios
          FOREIGN KEY (id_usuario_professor)
          REFERENCES usuarios(ID)
          ON DELETE SET NULL
  );

  CREATE TABLE convites (
      id_convite    VARCHAR(100) NOT NULL PRIMARY KEY,
      id_academia   VARCHAR(100),
      chave_convite VARCHAR(100) NOT NULL UNIQUE,

      CONSTRAINT fk_convites_academias
          FOREIGN KEY (id_academia)
          REFERENCES academias(ID)
          ON DELETE SET NULL
  );

  CREATE TABLE solicitacoes_convite (
      id_solicitacao VARCHAR(100) NOT NULL PRIMARY KEY,
      id_usuario     VARCHAR(100),
      id_academia    VARCHAR(100),

      CONSTRAINT fk_solicitacoes_usuarios
          FOREIGN KEY (id_usuario)
          REFERENCES usuarios(ID)
          ON DELETE SET NULL,

      CONSTRAINT fk_solicitacoes_academias
          FOREIGN KEY (id_academia)
          REFERENCES academias(ID)
          ON DELETE SET NULL
  );

  CREATE TABLE alunos (
      id_aluno          VARCHAR(100) NOT NULL PRIMARY KEY,
      faixa             VARCHAR(15),
      id_usuario_aluno  VARCHAR(100),
      id_academia_aluno VARCHAR(100),

      CONSTRAINT fk_alunos_usuarios
          FOREIGN KEY (id_usuario_aluno)
          REFERENCES usuarios(ID)
          ON DELETE SET NULL,

      CONSTRAINT fk_alunos_academias
          FOREIGN KEY (id_academia_aluno)
          REFERENCES academias(ID)
          ON DELETE SET NULL
  );

  CREATE TABLE instrutores (
      id_instrutor          VARCHAR(100) NOT NULL PRIMARY KEY,
      id_usuario_instrutor  VARCHAR(100),
      id_academia_instrutor VARCHAR(100),

      CONSTRAINT fk_instrutores_usuarios
          FOREIGN KEY (id_usuario_instrutor)
          REFERENCES usuarios(ID)
          ON DELETE SET NULL,

      CONSTRAINT fk_instrutores_academias
          FOREIGN KEY (id_academia_instrutor)
          REFERENCES academias(ID)
          ON DELETE SET NULL
  );

  CREATE TABLE aulas (
      id_aula      VARCHAR(100) NOT NULL PRIMARY KEY,
      data_aula    DATETIME,
      conteudo     VARCHAR(255),
      id_academia  VARCHAR(100),
      id_instrutor VARCHAR(100),

      CONSTRAINT fk_aulas_academias
          FOREIGN KEY (id_academia)
          REFERENCES academias(ID)
          ON DELETE SET NULL,

      CONSTRAINT fk_aulas_usuarios
          FOREIGN KEY (id_instrutor)
          REFERENCES usuarios(ID)
          ON DELETE SET NULL
  );

  CREATE TABLE presencas (
      id_presenca VARCHAR(100) NOT NULL PRIMARY KEY,
      id_aluno    VARCHAR(100),
      id_aula     VARCHAR(100),
      checkin_em  DATETIME,

      CONSTRAINT fk_presencas_alunos
          FOREIGN KEY (id_aluno)
          REFERENCES alunos(id_aluno)
          ON DELETE SET NULL,

      CONSTRAINT fk_presencas_aulas
          FOREIGN KEY (id_aula)
          REFERENCES aulas(id_aula)
          ON DELETE SET NULL
  );
