package kafka

import (
	"encoding/json"
	"errors"
	"log"
	"sync"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

//PubSub publish/subscribe pattern
type PubSub struct {
	Quit chan bool
	pub  *kafka.Producer
	sub  *kafka.Consumer
}

//NewConfig set kafka connect configs
func (ps *PubSub) NewConfig(adr, username, pass, sslkey, cacert, pem, certkey string) (err error) {
	ps.sub, err = kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":        adr,
		"security.protocol":        "SASL_SSL",
		"sasl.mechanisms":          "SCRAM-SHA-256",
		"sasl.username":            username,
		"sasl.password":            pass,
		"ssl.ca.location":          cacert,
		"ssl.certificate.location": pem,
		"ssl.key.location":         certkey,
		"ssl.key.password":         sslkey,
		"go.events.channel.enable": true,
		"group.id":                 "kortimilli",
		"session.timeout.ms":       6000,
		"auto.offset.reset":        "earliest",
		"debug":                    "generic,broker,security",
	})
	if err != nil {
		return
	}

	err = ps.sub.SubscribeTopics([]string{"kortimilliBalanceReq", "C2CReq"}, nil)
	if err != nil {
		return
	}

	ps.pub, err = kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers":        adr,
		"security.protocol":        "SASL_SSL",
		"sasl.mechanisms":          "SCRAM-SHA-256",
		"sasl.username":            username,
		"sasl.password":            pass,
		"ssl.ca.location":          cacert,
		"ssl.certificate.location": pem,
		"ssl.key.location":         certkey,
		"ssl.key.password":         sslkey,
		"debug":                    "generic,broker,security",
	})
	return
}

//Start start all consumers
func (ps *PubSub) Start(wg *sync.WaitGroup) {
	log.Printf("Created balanceReq consumer %v\n", ps.sub)
	wg.Add(1)
	run := true
	for run {
		select {
		case <-ps.Quit:
			log.Println("Caught signal to stop consumer:")
			run = false
			wg.Done()
		case ev := <-ps.sub.Events():
			switch e := ev.(type) {
			case kafka.AssignedPartitions:
				log.Println("AssignedPartitions", e)
				ps.sub.Assign(e.Partitions)
			case kafka.RevokedPartitions:
				log.Println("RevokedPartitions", e)
				ps.sub.Unassign()
			case *kafka.Message:
				if *e.TopicPartition.Topic == "kortimilliBalanceReq" {
					reqBal := balance{}
					if err := json.Unmarshal(e.Value, &reqBal); err != nil {
						log.Println("kafka.Message json.Unmarshal", err, string(e.Value))
					}
					log.Println(reqBal)
					reqBal.npc()
					go ps.pubResp("kortimilliBalanceResp", reqBal)
				} else if *e.TopicPartition.Topic == "C2CReq" {
					reqC2C := c2cReq{}
					if err := json.Unmarshal(e.Value, &reqC2C); err != nil {
						log.Println("kafka.Message json.Unmarshal", err, string(e.Value))
					}
					res := reqC2C.npc()
					go ps.pubResp("C2CResp", res)
				}
				ps.sub.Commit()
			case kafka.PartitionEOF:
				log.Println("Reached: ", e)
			case kafka.Error:
				// Errors should generally be considered as informational, the client will try to automatically recover
				log.Println("Error: ", e)
			}
		}
	}

	log.Println("Closing consumer")
	ps.sub.Close()
}


//pubResp produce response message
func (ps *PubSub) pubResp(topic string, prm interface{}) {

	vByte, err := json.Marshal(prm)
	if err != nil {
		log.Println("BalanceResp json.Marshal", err)
		return
	}

	dChan := make(chan kafka.Event)
	err = ps.pub.Produce(&kafka.Message{TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny}, Value: vByte}, dChan)
	e := <-dChan
	m := e.(*kafka.Message)
	if m.TopicPartition.Error != nil {
		log.Printf("Delivery failed: %v\n", m.TopicPartition.Error)
	}

	close(dChan)
}

