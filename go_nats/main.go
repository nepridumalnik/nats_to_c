package main

import (
	"C"
	"github.com/nats-io/nats.go"

	"fmt"
	"time"
)

func print_error(err error) {
	fmt.Printf("err: %s\n", err.Error())
}

// Колбэк
func cb(msg *nats.Msg) {
	fmt.Printf("Async subscription: %s\n", string(msg.Data))
}

//export test_export
func test_export() {
	fmt.Printf("Hello from GO!\n")
}

//export nats_example
func nats_example() {
	const subject string = "test_subject"
	const error_format string = "err: %s\n"
	const timeout time.Duration = time.Second * 3

	// Создание соединение с сервером
	nc, err := nats.Connect(nats.DefaultURL)

	if nil == nc {
		print_error(err)
		return
	}

	// Асинхронная nats подписка
	{
		// Создание асинхронной подписки
		asyncSub, err := nc.Subscribe(subject, cb)

		if nil != err {
			print_error(err)
			return
		}

		// Публикация данных для подписки
		err = nc.Publish(subject, []byte("Some asynchronous published data"))

		if nil != err {
			print_error(err)
			return
		}

		// Дизлайк, отписка
		asyncSub.Unsubscribe()
		// Закрытие соединения с подпиской при продолжающейся работе колбэков
		asyncSub.Drain()
	}

	{
		// Создание синхронной подписки
		syncSub, err := nc.SubscribeSync(subject)

		if nil != err {
			print_error(err)
			return
		}

		// Публикация данных для подписки
		err = nc.Publish(subject, []byte("Some synchronous published data"))

		if nil != err {
			print_error(err)
			return
		}

		// Чтение сообщения для подписки
		msg, err := syncSub.NextMsg(timeout)

		if nil != err {
			print_error(err)
			return
		}

		if nil != msg {
			fmt.Printf("Sync subscription: %s\n", msg.Data)
		}

		syncSub.Unsubscribe()
		syncSub.Drain()
	}

	// Отключает все подписки
	nc.Drain()
	// Закрытие соединения с сервером
	nc.Close()
}

func main() {
	nats_example()
}
