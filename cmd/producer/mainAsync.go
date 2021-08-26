package main

import (
	"fmt"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	deliveryChan := make(chan kafka.Event)

	message := "mensagem asyn"
	topic := "teste"

	producer := NewKafkaProducer()

	Publish(message, topic, producer, nil, deliveryChan)

	// Forma assync, "go" joga para uma outra trehad, ai o temrinal nãO TRAVA
	go DeliveryReport(deliveryChan)

	fmt.Println("terminou")
	producer.Flush(1000)

}

func NewKafkaProducer() *kafka.Producer {

	// acks = 0 NÃO ESPERA RETORNO DE RECEBIMENTO
	// acks = 1 espera apenas o lider dizer que recebeu
	// acks = all ou -1 espera o lieder e os demais receber
	// enable.idempotence = somente true, quando o acks = all,
	// 	por que le vai garantir que tudo receber, só foi envado uma vez e na ordem certa (lento)
	configMap := &kafka.ConfigMap{
		"bootstrap.servers":   "gokafka_kafka_1:9092",
		"delivery.timeout.ms": "0",
		"acks":                "all",
		"enable.idempotence":  "true",
	}

	p, err := kafka.NewProducer(configMap)

	if err != nil {
		log.Println(err.Error())
	}

	return p
}

func Publish(msg string, topic string, producer *kafka.Producer, key []byte, deliveryChan chan kafka.Event) error {
	message := &kafka.Message{
		Value: []byte(msg),
		TopicPartition: kafka.TopicPartition{
			Topic:     &topic,
			Partition: kafka.PartitionAny,
		},
		Key: key,
	}

	err := producer.Produce(message, deliveryChan)

	if err != nil {
		return err
	}

	return nil
}

func DeliveryReport(deliveryChannel chan kafka.Event) {
	for e := range deliveryChannel {
		switch ev := e.(type) {
		case *kafka.Message:
			if ev.TopicPartition.Error != nil {
				fmt.Println("Erro aos enviar")
			} else {
				// Vai imprimir o topic[partição]offset
				fmt.Println("Mensagem enviada:", ev.TopicPartition)
			}
		}
	}
}
