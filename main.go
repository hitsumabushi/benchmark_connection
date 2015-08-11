package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strconv"
	"time"

	"github.com/codegangsta/cli"
)

const (
	Version = "1.0"
)

func main() {
	log.SetFlags(0)
	app := cli.NewApp()
	app.Name = "Benchmark connection"
	app.Version = Version
	app.Author = "hitsumabushi"
	app.Email = "053023.math@gmail.com"
	app.Commands = Commands
	app.EnableBashCompletion = true
	app.Flags = []cli.Flag{
		cli.IntFlag{
			Name: "endport, ep",
		},
		cli.IntFlag{
			Name: "beginport, bp",
		},
	}
	app.Run(os.Args)
}

var Commands = []cli.Command{
	commandServer,
	commandClient,
}

var commandServer = cli.Command{
	Name:    "server",
	Aliases: []string{"s"},
	Usage:   "Launch Server",
	Action:  doServer,
	Flags:   []cli.Flag{},
}

var commandClient = cli.Command{
	Name:    "client",
	Aliases: []string{"c"},
	Usage:   "Launch Client",
	Action:  doClient,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name: "host",
		},
		cli.StringFlag{
			Name: "signature, i",
		},
	},
}

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
		fmt.Println(count)
	}
}

func listenPort(n int, count *int) {
	ln, err := net.Listen("tcp4", ":"+strconv.Itoa(n))
	if err != nil {
		log.Println(err)
		return
	}
	defer ln.Close()
	(*count)++

	for {
		// wait connection
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
		}
		go func(c net.Conn) {
			io.Copy(c, c)
			fmt.Printf("%#v", c)
			c.Close()
		}(conn)
		time.Sleep(time.Duration(10) * time.Second)
	}
}

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
		time.Sleep(time.Duration(10) * time.Second)
		fmt.Println("Client Conections: " + strconv.Itoa(count))
	}
}

func connectToHost(p int, host, signature string, count *int) {
	conn, _ := net.Dial("tcp", host+":"+strconv.Itoa(p))
	defer conn.Close()
	(*count)++
	for {
		time.Sleep(time.Duration(10) * time.Second)
		conn.Write([]byte("client[" + signature + "]:" + strconv.Itoa(p)))
	}
}
