# GO KAFKA

Exemplo de app com integração com kafka desenvolvido em GO.

Utiliza recursos como, producer sincrono e assincrono, consumers e keys.

## Steps:

    docker-compose up -d

Verificando containers:
    
    docker-compose ps

Acessando container kafka:
    
    docker exec -it gokafka_kafka_1 bash

Criando o topic:

    kafka-topics --create --bootstrap-server=localhost:9092 --topic=teste --partitions=3

criando um consumer:
    
    kafka-console-consumer --bootstrap-server=localhost:9092 --topic=teste

Acessando o container GO(app):
    
    docker exec -it gokafka bash

Executando o producer Sync(dentro do container):

    go run cmd/producer/mainSync.go