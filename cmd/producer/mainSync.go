package main

import (
	"fmt"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	deliveryChan := make(chan kafka.Event)

	message := "mensagem"
	topic := "teste"

	producer := NewKafkaProducer()

	Publish(message, topic, producer, nil, deliveryChan)

	// Evento que fica lendo o deliveryChan, o programa fica parado, ou seja, sincrono
	// Fica mais lento
	e := <-deliveryChan
	msg := e.(*kafka.Message)
	if msg.TopicPartition.Error != nil {
		fmt.Println("Erro aos enviar")
	} else {
		// Vai imprimir o topic[partição]offset
		fmt.Println("Mensagem enviada:", msg.TopicPartition)
	}

	// Espera 1 segundo até receber ok
	producer.Flush(1000)

}

func NewKafkaProducer() *kafka.Producer {
	configMap := &kafka.ConfigMap{
		"bootstrap.servers": "gokafka_kafka_1:9092",
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
