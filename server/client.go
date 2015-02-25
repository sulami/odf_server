package server

import (
	"bufio"
	"net"
	"strings"
)

type Client interface {
	Read()
	Write(string)
	Listen()
	parseCmd(string)
}

type GameClient struct {
	conn net.Conn
	reader *bufio.Reader
	writer *bufio.Writer
	game Game
}

func (client *GameClient) Read() {
	for {
		line, _ := client.reader.ReadString('\n')
		client.parseCmd(line)
	}
}

func (client *GameClient) Write(msg string) {
	client.writer.WriteString(msg + "\n")
	client.writer.Flush()
}

func (client *GameClient) Listen() {
	go client.Read()
}

func (client *GameClient) parseCmd(line string) {
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
	var c Client
	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)
	game := DefaultGame{}

	client := &GameClient{
		conn: conn,
		reader: reader,
		writer: writer,
		game: game,
	}

	client.Listen()
	c = client

	return &c
}

