package main

import "net"
import "strings"
import "bufio"
import "net/textproto"
import "fmt"

func main() {
	conn, err := net.Dial("tcp", "irc.chat.twitch.tv:6667")
	if err != nil {
		panic(err)
	}

	conn.Write([]byte("PASS " + "oauth:yourkey" + "\r\n"))
	conn.Write([]byte("NICK " + "yourusername" + "\r\n"))
	conn.Write([]byte("JOIN " + "#yourchannel" + "\r\n"))
	defer conn.Close()

	// handles reading from the connection
	tp := textproto.NewReader(bufio.NewReader(conn))

	// listens/responds to chat messages
	for {
		msg, err := tp.ReadLine()
		if err != nil {
			panic(err)
		}
		fmt.Println(msg)
		// split the msg by spaces
		msgParts := strings.Split(msg, " ")

		// if the msg contains PING you're required to
		// respond with PONG else you get kicked
		if msgParts[0] == "PING" {
			conn.Write([]byte("PONG " + msgParts[1]))
			continue
		}

		// if msg contains PRIVMSG then respond
		if msgParts[1] == "PRIVMSG" {
			// echo back the same message
			conn.Write([]byte("PRIVMSG " + msgParts[2] + " " + msgParts[3] + "\r\n"))
		}
	}
}
