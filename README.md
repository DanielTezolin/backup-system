
# Sistema de backup

Este projeto teve origem na necessidade comum de realizar backups automatizados de bancos de dados, juntamente com um estudo meu sobre a linguagem GO.

O objetivo é simples: criar um backup da base de dados, salvá-lo no disco e fazer o upload de uma cópia para o S3. Por fim, notificar se tudo ocorreu conforme o esperado ou não no Teams.

## Como usar

Optei por utilizar o formato de configuração em arquivo, atualmente utilizando JSON, como alternativa ao uso de um banco de dados.

Existem três configurações possíveis inicialmente: `notifications`, `AWSupload`, `databases`

- **notifications**:
```json
  "notifications": {
    "active": false,
    "teamsWebhook": ""
  },
```
| Parâmetro   | Tipo       | Descrição                           |
| :---------- | :--------- | :---------------------------------- |
| `active` | `boolean` | **Obrigatório**. Este parâmetro indica ao sistema se o serviço de notificação deve ser ativado. |

| Parâmetro   | Tipo       | Descrição                           |
| :---------- | :--------- | :---------------------------------- |
| `teamsWebhook` | `string` | **É necessário** inserir o webhook do Teams para receber notificações apenas se o serviço estiver ativo. |

- **AWSupload**:
```json
    "AWSupload": {
    "active": false,
    "s3Bucket": "",
    "s3Region": "",
    "s3Key": "",
    "s3Secret": ""
  },
```
| Parâmetro   | Tipo       | Descrição                           |
| :---------- | :--------- | :---------------------------------- |
| `active` | `boolean` | **Obrigatório**. Indica se o recurso de upload para o AWS S3 está ativo ou não. |

| Parâmetro   | Tipo       | Descrição                           |
| :---------- | :--------- | :---------------------------------- |
| `s3Bucket` | `string` | **Obrigatório** O nome do bucket S3 onde os arquivos serão carregados. |

| Parâmetro   | Tipo       | Descrição                           |
| :---------- | :--------- | :---------------------------------- |
| `s3Region` | `string` | **Obrigatório** A região da AWS onde o bucket S3 está localizado. |

| Parâmetro   | Tipo       | Descrição                           |
| :---------- | :--------- | :---------------------------------- |
| `s3Key` | `string` | **Obrigatório** A chave de acesso da conta AWS usada para autenticar o upload de arquivos no bucket S3. |

| Parâmetro   | Tipo       | Descrição                           |
| :---------- | :--------- | :---------------------------------- |
| `s3Secret` | `string` | **Obrigatório** A chave de acesso secreta da conta AWS usada para autenticar o upload de arquivos no bucket S3. | 

- **databases**:
```json
  "databases": [
    {
      "dbname": "",
      "username": "",
      "password": "",
      "host": "",
      "port": ""
    }
  ]
```
| Parâmetro   | Tipo       | Descrição                           |
| :---------- | :--------- | :---------------------------------- |
| `dbname` | `string` | **Obrigatório** O nome do banco de dados. |

| Parâmetro   | Tipo       | Descrição                           |
| :---------- | :--------- | :---------------------------------- |
| `username` | `string` | **Obrigatório**  O nome de usuário usado para se conectar ao banco de dados.|

| Parâmetro   | Tipo       | Descrição                           |
| :---------- | :--------- | :---------------------------------- |
| `password` | `string` | **Obrigatório**  A senha usada para se conectar ao banco de dados. |

| Parâmetro   | Tipo       | Descrição                           |
| :---------- | :--------- | :---------------------------------- |
| `host` | `string` | **Obrigatório**  O endereço do servidor ou host onde o banco de dados está localizado. |

| Parâmetro   | Tipo       | Descrição                           |
| :---------- | :--------- | :---------------------------------- |
| `port` | `string` | **Obrigatório**  O número da porta usada para se conectar ao banco de dados. |

O arquivo de configuração deve ter uma aparência semelhante a esta:
```json
{
  "notifications": {
    "active": false,
    "teamsWebhook": ""
  },
  "AWSupload": {
    "active": false,
    "s3Bucket": "",
    "s3Region": "",
    "s3Key": "",
    "s3Secret": ""
  },
  "databases": [
    {
      "dbname": "",
      "username": "",
      "password": "",
      "host": "",
      "port": ""
    }
  ]
}
```
## Deploy

Para realizar o deploy, é crucial mapear três pastas essenciais: `logs`, `dumps` e `config`. Abaixo está um exemplo de docker compose.

```yml
version: '3'
services:
  app:
    image: sua/image:latest
    volumes:
      - ./logs:/app/logs
      - ./dumps:/app/dumps
      - ./config:/app/config
```


## Roadmap
- Aprimorar a estrutura de configuração para utilizar um arquivo YML em vez de JSON.

- Expandir o suporte no serviço de notificação para incluir mais aplicativos.

- Resolver um problema na geração de arquivos de logs.

- Desenvolver uma dashboard para visualizar o status dos backups e facilitar a edição do arquivo de configuração.

- Implementar um novo sistema que testa o dump dos bancos de dados

