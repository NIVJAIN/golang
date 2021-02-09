package comms

import (
	"fmt"

	"github.com/streadway/amqp"
)

//Publish publishes a request to the amqp queue
func (c *Connection) Publish(m Message) error {
	fmt.Println("publishung.....", c.exchange)
	select { //non blocking channel - if there is no error will go to default where we do nothing
	case err := <-c.err:
		if err != nil {
			fmt.Println("Errorn reconnet", err)
			c.Reconnect()
		}
	default:
	}

	p := amqp.Publishing{
		Headers:       amqp.Table{"type": m.Body.Type},
		ContentType:   m.ContentType,
		CorrelationId: m.CorrelationID,
		Body:          m.Body.Data,
		ReplyTo:       m.ReplyTo,
	}
	// if err := c.channel.Publish(c.exchange, m.Queue, false, false, p); err != nil {
	if err := c.channel.Publish(c.exchange, "my_routing_key", false, false, p); err != nil {
		return fmt.Errorf("Error in Publishing: %s", err)
	}
	fmt.Println("published.....", m.Queue)
	return nil
}
