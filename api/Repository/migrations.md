CREATE TABLE addresses (                                                                                                                                                                                       
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

  CREATE TABLE users (                                                                                                                                                                                           
      ID           VARCHAR(100) NOT NULL PRIMARY KEY,
      Name         VARCHAR(150) NOT NULL,                                                                                                                                                                        
      Email        VARCHAR(200) NOT NULL UNIQUE,            
      Password     VARCHAR(100) NOT NULL,
      Gender       CHAR(10),                                                                                                                                                                                     
      CPF          CHAR(11)     NOT NULL UNIQUE,
      DateBirth    DATE,                                                                                                                                                                                         
      addresses_id VARCHAR(100),                                                                                                                                                                                 
   
      CONSTRAINT fk_users_addresses                                                                                                                                                                              
          FOREIGN KEY (addresses_id)                        
          REFERENCES addresses(ID)
          ON DELETE SET NULL
  );                                                                                                                                                                                                             
   
 select * from gyms
 select * from teachers
      
  CREATE TABLE gyms (                                                                                                                                                                                            
      ID   VARCHAR(100) NOT NULL PRIMARY KEY,               
      NAME VARCHAR(100) NOT NULL UNIQUE,
      CNPJ VARCHAR(14)  NOT NULL UNIQUE
  );                                                                                                                                                                                                             
   
  CREATE TABLE teachers (                                                                                                                                                                                        
      id_teacher        VARCHAR(100) NOT NULL PRIMARY KEY,  
      id_gyms_teachers  VARCHAR(100),
      id_users_teachers VARCHAR(100),

      CONSTRAINT fk_teachers_gyms                                                                                                                                                                                
          FOREIGN KEY (id_gyms_teachers)
          REFERENCES gyms(ID)                                                                                                                                                                                    
          ON DELETE SET NULL,                               

      CONSTRAINT fk_teachers_users
          FOREIGN KEY (id_users_teachers)
          REFERENCES users(ID)
          ON DELETE SET NULL

  );            
  
 CREATE TABLE addresses (                                                                                                                                                                                       
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

  CREATE TABLE users (                                                                                                                                                                                           
      ID           VARCHAR(100) NOT NULL PRIMARY KEY,
      Name         VARCHAR(150) NOT NULL,                                                                                                                                                                        
      Email        VARCHAR(200) NOT NULL UNIQUE,            
      Password     VARCHAR(100) NOT NULL,
      Gender       CHAR(10),                                                                                                                                                                                     
      CPF          CHAR(11)     NOT NULL UNIQUE,
      DateBirth    DATE,                                                                                                                                                                                         
      addresses_id VARCHAR(100),                                                                                                                                                                                 
   
      CONSTRAINT fk_users_addresses                                                                                                                                                                              
          FOREIGN KEY (addresses_id)                        
          REFERENCES addresses(ID)
          ON DELETE SET NULL
  );                                                                                                                                                                                                             
   
 select * from gyms
 select * from teachers
 select * from users
 select * from invites
 select * from invite_requests
      
  CREATE TABLE gyms (                                                                                                                                                                                            
      ID   VARCHAR(100) NOT NULL PRIMARY KEY,               
      NAME VARCHAR(100) NOT NULL UNIQUE,
      CNPJ VARCHAR(14)  NOT NULL UNIQUE
  );                                                                                                                                                                                                             
   
  CREATE TABLE teachers (                                                                                                                                                                                        
      id_teacher        VARCHAR(100) NOT NULL PRIMARY KEY,  
      id_gyms_teachers  VARCHAR(100),
      id_users_teachers VARCHAR(100),

      CONSTRAINT fk_teachers_gyms                                                                                                                                                                                
          FOREIGN KEY (id_gyms_teachers)
          REFERENCES gyms(ID)                                                                                                                                                                                    
          ON DELETE SET NULL,                               

      CONSTRAINT fk_teachers_users
          FOREIGN KEY (id_users_teachers)
          REFERENCES users(ID)
          ON DELETE SET NULL

  );            
  
     CREATE TABLE invites (                                                                                                                                                                                         
      id_invite  VARCHAR(100) NOT NULL PRIMARY KEY,
      id_gym     VARCHAR(100),                                                                                                                                                                                   
      invite_key VARCHAR(100) NOT NULL UNIQUE,                                                                                                                                                                   
   
      CONSTRAINT fk_invites_gyms                                                                                                                                                                                 
          FOREIGN KEY (id_gym)                              
          REFERENCES gyms(ID)                                                                                                                                                                                    
          ON DELETE SET NULL
  );  

      
  CREATE TABLE invite_requests (                                                                                                                                                                                 
      id_request VARCHAR(100) NOT NULL PRIMARY KEY,
      id_user    VARCHAR(100),                                                                                                                                                                                   
      id_gym     VARCHAR(100),                              

      CONSTRAINT fk_invite_requests_users                                                                                                                                                                        
          FOREIGN KEY (id_user)
          REFERENCES users(ID)                                                                                                                                                                                   
          ON DELETE SET NULL,                               

      CONSTRAINT fk_invite_requests_gyms
          FOREIGN KEY (id_gym)
          REFERENCES gyms(ID)
          ON DELETE SET NULL
  );
      
     
      
 select * from invite_requests
     
   
      
