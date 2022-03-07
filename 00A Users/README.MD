# Users

Sua função é escrever um código de inserção de usuários, baseado no que existe hoje, e criar os testes
necessários para garantir um código eficiente e de fácil manutenção.

O esquema do banco de dados apresentado abaixo é MySQL e foi copiado do banco original.

Sua função é escrever um código simples, escrever os testes unitários para garantir a integridade do 
código e testes de desempenho/conceito para garantir a melhor forma de arquivar os dados.

```mysql
CREATE TABLE `Users` (
		`Id` varchar(26) NOT NULL,
		`CreateAt` bigint(20) DEFAULT NULL,
		`UpdateAt` bigint(20) DEFAULT NULL,
		`DeleteAt` bigint(20) DEFAULT NULL,
		`Username` varchar(64) DEFAULT NULL,
		`Password` varchar(128) DEFAULT NULL,
		`AuthData` varchar(128) DEFAULT NULL,
		`AuthService` varchar(32) DEFAULT NULL,
		`Email` varchar(128) DEFAULT NULL,
		`EmailVerified` tinyint(1) DEFAULT NULL,
		`Nickname` varchar(64) DEFAULT NULL,
		`FirstName` varchar(64) DEFAULT NULL,
		`LastName` varchar(64) DEFAULT NULL,
		`Position` varchar(64) DEFAULT NULL,
		`Roles` text,
		`AllowMarketing` tinyint(1) DEFAULT NULL,
		`Props` text,
		`NotifyProps` text,
		`LastPasswordUpdate` bigint(20) DEFAULT NULL,
		`LastPictureUpdate` bigint(20) DEFAULT NULL,
		`Followers` int(11) DEFAULT '0',
		`Following` int(11) DEFAULT '0',
		`FailedAttempts` int(11) DEFAULT NULL,
		`Locale` varchar(5) DEFAULT NULL,
		`MfaActive` tinyint(1) DEFAULT NULL,
		`MfaSecret` varchar(128) DEFAULT NULL,
		`Bio` text,
		`Twitter` varchar(45) DEFAULT NULL,
		`Facebook` varchar(45) DEFAULT NULL,
		`Badge` varchar(45) DEFAULT 'member',
		`UserCompany` varchar(1) DEFAULT '0',
		`Ring` varchar(45) DEFAULT 'green',
		PRIMARY KEY (`Id`),
		UNIQUE KEY `Username` (`Username`),
		UNIQUE KEY `AuthData` (`AuthData`),
		UNIQUE KEY `Email` (`Email`),
		KEY `idx_users_email` (`Email`),
		KEY `idx_users_update_at` (`UpdateAt`),
		KEY `idx_users_create_at` (`CreateAt`),
		KEY `idx_users_delete_at` (`DeleteAt`),
		FULLTEXT KEY `idx_users_all_txt` (`Username`,`FirstName`,`LastName`,`Nickname`,`Email`),
		FULLTEXT KEY `idx_users_all_no_full_name_txt` (`Username`,`Nickname`,`Email`),
		FULLTEXT KEY `idx_users_names_txt` (`Username`,`FirstName`,`LastName`,`Nickname`),
		FULLTEXT KEY `idx_users_names_no_full_name_txt` (`Username`,`Nickname`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4
```

No código original o campo Badge contém as seguintes entradas:

| Badge            |
|------------------|
| member           |
| master           |
| premium          |
| (nada)           |
| admin            |
| private          |
| collab           |
| master_vitalicio |
| tradersclub      |
| tc               |
| gold             |
| collab_silver    |
| collab_b2b       |

O campo UserCompany contém os valores:

| UserCompany |
|-------------|
| (nada)      |
| Null        |
| 0           |
| 2           |
| 4           |
| 5           |

O campo Ring contém os valores:

| Ring         |
|--------------|
| silver       |
| transparent  |
| gold         |
| green        |
| bronze       |

O campo Roles é um campo texto que pode ter os valores abaixo separados por espaço.

Monte um teste de prova de conceito e messa o tempo de resposta usando `Role LIKE ?` e `JOIN` para 
medir os tempos de resposta.

| Roles         |
|---------------|
| (nada)        |
| system_user   |
| master_member | 
| allow_news    |
| influencer    |  
| tradersclub   | 
| commercial_n2 |    
| system_admin  |  
| journalist    |
| mover         |
| manager       |
| streamer      |   
| system_collab |    
| verified      |  
| customer      |
| multisession  |
| editor_scoop  |
| commercial_n1 |
| store         |  
| tools         |  
| moderator     |  
| support       |  
| cs_financial  |  
| educational   |  
| financial     |  
| c_manager     |  
| premium       |  






