package main

import (
	"log"
	"net"
	"os"
	"strconv"
	"time"

	"github.com/codegangsta/cli"
)

func doClient(c *cli.Context) {
	beginport := c.GlobalInt("beginport")
	endport := c.GlobalInt("endport")
	host := c.String("host")
	signature := c.String("signature")
	if signature == "" {
		signature, _ = os.Hostname()
	}
	if beginport == 0 || endport == 0 || host == "" {
		log.Println("Invalid arguments")
		return
	}

	count := 0
	for p := beginport; p < endport; p++ {
		go connectToHost(p, host, signature, &count)
	}
	for {
		time.Sleep(time.Duration(Wait) * time.Second)
		log.Println("client[" + signature + "]:" + strconv.Itoa(count))
	}
}

func connectToHost(p int, host, signature string, count *int) {
	for {
		conn, err := net.Dial("tcp", host+":"+strconv.Itoa(p))
		if err == nil {
			(*count)++
			defer client_close(&conn, count)
			for {
				conn.Write([]byte("client[" + signature + "]: " + strconv.Itoa(p)))
				time.Sleep(time.Duration(Wait) * time.Second)
			}
		} else {
			log.Println("cannot connect :" + host + ":" + strconv.Itoa(p))
			time.Sleep(time.Duration(Wait) * time.Second)
		}
	}
}

func client_close(c *net.Conn, count *int) {
	(*c).Close()
	(*count)--
}
