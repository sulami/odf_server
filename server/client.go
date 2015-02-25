package server

import (
	"bufio"
	"net"
	"strings"
)

type Client struct {
	conn net.Conn
	reader *bufio.Reader
	writer *bufio.Writer
	game Game
}

func (client *Client) Read() {
	for {
		line, _ := client.reader.ReadString('\n')
		client.parseCmd(line)
	}
}

func (client *Client) Write(msg string) {
	client.writer.WriteString(msg + "\n")
	client.writer.Flush()
}

func (client *Client) Listen() {
	go client.Read()
}

func (client *Client) parseCmd(line string) {
	cmd := strings.Fields(line)
	if len(cmd) != 0 {
		r, f := client.game.Parse(cmd)
		client.Write(r)
		if f {
			Log("Closing connection to " +
				client.conn.RemoteAddr().String())
			client.conn.Close()
		}
	}
}

func NewClient(conn net.Conn) *Client {
	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)
	game := Game{}

	client := &Client{
		conn: conn,
		reader: reader,
		writer: writer,
		game: game,
	}

	client.Listen()

	return client
}

