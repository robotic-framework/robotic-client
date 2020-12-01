package main

import (
	"encoding/hex"
	"fmt"
	"github.com/tarm/serial"
)

func main() {
	data, _ := hex.DecodeString("244d3c007070")
	fmt.Println(data)

	conn, err := serial.OpenPort(&serial.Config{
		Name:     "/dev/cu.usbmodem143101",
		Baud:     115200,
		Size:     8,
		Parity:   serial.ParityNone,
		StopBits: serial.Stop1,
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("write data: %v\n", data)
	_, err = conn.Write(data)
	if err != nil {
		panic(err)
	}

	go func() {
		for {
			fmt.Println("waiting header")
			header := make([]byte, 3)
			_, err = conn.Read(header)
			if err != nil {
				panic(err)
			}

			fmt.Printf("got header: %s\n", string(header))
			lengthAndCode := make([]byte, 2)
			_, err = conn.Read(lengthAndCode)
			if err != nil {
				panic(err)
			}
			length := lengthAndCode[0]

			fmt.Printf("got length: %d\n", length)
			data = make([]byte, length)
			_, err = conn.Read(data)
			if err != nil {
				panic(err)
			}

			fmt.Println(hex.EncodeToString(data))
		}
	}()

	select {}
}
