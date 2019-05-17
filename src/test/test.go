package main

import (
	"encoding/binary"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		panic(err)
	}

	// Hello 消息（JSON 格式）
	// 对应游戏服务器 Hello 消息结构体
	data := []byte(`{
		"Hello": {
			"Name": "leaf"
		}
	}`)
	//对应People结构体
	data2 :=[]byte(`{
		"People": {
			"Name": "Mit",
			"Sex": "M",
			"Age": 23
		}
	}`)

	// len + data
	m := make([]byte, 2+len(data))

	// 默认使用大端序
	binary.BigEndian.PutUint16(m, uint16(len(data)))

	copy(m[2:], data)

	// len + data
	m2 := make([]byte, 2+len(data2))

	// 默认使用大端序
	binary.BigEndian.PutUint16(m2, uint16(len(data2)))

	copy(m2[2:], data2)

	// 发送消息
	conn.Write(m)
	// 发送消息
	conn.Write(m2)
}
