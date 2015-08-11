package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"strconv"
	"time"

	"github.com/codegangsta/cli"
)

func doServer(c *cli.Context) {
	beginport := c.GlobalInt("beginport")
	endport := c.GlobalInt("endport")
	if beginport == 0 || endport == 0 {
		log.Println("Invalid arguments")
		return
	}
	count := 0
	for p := beginport; p < endport; p++ {
		go listenPort(p, &count)
	}
	for {
		time.Sleep(time.Duration(10) * time.Second)
		fmt.Println("Server count: " + strconv.Itoa(count))
	}
}

func listenPort(n int, count *int) {
	ln, err := net.Listen("tcp4", ":"+strconv.Itoa(n))
	if err != nil {
		log.Println(err)
		return
	}
	defer ln.Close()

	for {
		// wait connection
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
		}
		go func(c net.Conn) {
			(*count)++
			io.Copy(c, c)
			server_close(&c, count)
		}(conn)
		time.Sleep(time.Duration(10) * time.Second)
	}
}

func server_close(c *net.Conn, count *int) {
	(*c).Close()
	(*count)--
}
