package main

import "fmt"

type sender interface {
	send(message string)
}
type email struct{}
type whatsapp struct{}
type messanger struct{}

func (e email) send(message string) {
	fmt.Println("Mail: ", message)
}

func (w whatsapp) send(message string) {
	fmt.Println("whatsapp message: ", message)
}

func (m messanger) send(message string) {
	fmt.Println("Messanger: ", message)
}

func createMessage(s sender) {
	s.send("Sended")
}

func main() {
	e := email{}
	e.send("how are you vivek!")
	createMessage(e)
}
