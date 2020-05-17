package main

import "github.com/streadway/amqp"

type MqClient struct {
	conn         *amqp.Connection
	channel      *amqp.Channel
	done         chan error
	block        chan amqp.Blocking
	channelClose chan *amqp.Error
	clientClose  chan *amqp.Error


}

func (m *MqClient) Push()  {

}






func NewClient() *MqClient  {
	return  &MqClient{}
}
