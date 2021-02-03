# CodePix

1 - Criar o Dockerfile ele será responsavel por subir o golang e o kafka
2 - Criar o docker-compose.yaml não será necessário nos preocupar em instalar o postgress etc...

- Executar o comando:

```shell
docker-compose up -d
```

- Ele subirá todos os containers dentro de do docker-compose e exibirá os containers que foram subidos:

```shell
Starting codepix_app_1                    ... done
Starting codepix_zookeeper_1 ... done
Starting codepix_db_1                     ... done
Starting codepix_kafka_1     ... done
Starting codepix_control-center_1         ... done
Starting codepix_kafka-topics-generator_1 ... done
Creating codepix_pgadmin_1                ... done
```
por fim execute o comando:

```shell
docker exec -it codepix_app_1 bash
```

- Com isso seremos capaz de criar nossa aplicação em go
- Executar o comando:

```shell
go mod init github.com/seu_git/nome_do_repositorio
```

- Será criado o arquivo chamado go.mod

- Criar a pasta `domain/model` onde estará as entidades
- Criaremos dentro dessa pasta o arquivo `bank.go`


---

## Comando do proto:
application/grpc/pb
application/grpc/protofiles
```shell
protoc --go_out=application/grpc/pb --go_opt=paths=source_relative --go-grpc_out=application/grpc/pb --go-grpc_opt=paths=source_relative --proto_path=application/grpc/protofiles application/grpc/protofiles/*.proto
```

- Executar aplicação:

```shell
go run cmd/main.go
```