package internal

import (
	"leaf/gate"
	"leaf/log"
	"reflect"
	"server/msg"
)

func init()  {
	// 向当前模块（game 模块）注册 Hello 消息的消息处理函数 handleHello
	handler(&msg.Hello{},handleHello)
	handler(&msg.People{},handlePeople)
}

func handler(m interface{},h interface{})  {
	skeleton.RegisterChanRPC(reflect.TypeOf(m),h)
}

func handleHello(args []interface{})  {
	// 收到的 Hello 消息
	m := args[0].(*msg.Hello)
	// 消息的发送者
	a := args[1].(gate.Agent)

	// 输出收到的消息的内容
	log.Debug("hello %v", m.Name)

	// 给发送者回应一个 Hello 消息
	a.WriteMsg(&msg.Hello{
		Name: "client",
	})
}

func handlePeople(args []interface{})  {
	// 收到的 People 消息
	m := args[0].(*msg.People)
	// 消息的发送者
	a := args[1].(gate.Agent)

	// 输出收到的消息的内容
	log.Debug("Name %v,Sex %v,Age %v", m.Name,m.Sex,m.Age)

	// 给发送者回应一个 Hello 消息
	a.WriteMsg(&msg.People{
		Name: "client",
		Sex: "F",
		Age: 0,
	})
}