package server

import (
	"bufio"
	"errors"
	"fmt"
	"net"
	"strconv"
	"strings"
	"time"
)

type Server interface {
	Listen() error
	StopListening() error
}

type GameServer struct {
	Port	int
	Online	bool
}

func (s *GameServer) Listen() (err error) {
	Log("Server starting up...")

	if s.Online {
		err = errors.New("Server is already online")
		return
	}

	ln, e := net.Listen("tcp", ":" + strconv.Itoa(s.Port))
	if e != nil {
		err = e
		return
	}

	for {
		conn, e:= ln.Accept()
		if e != nil {
			Log(e.Error())
		}
		go handleConnection(conn)
	}

	s.Online = true
	Log("Listening on port " + strconv.Itoa(s.Port))

	return
}

func (s *GameServer) StopListening() (err error) {
	Log("Stopping server...")

	if !s.Online {
		err = errors.New("Server is not running")
		return
	}

	s.Online = false
	Log("Server stopped")

	return
}

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
		switch cmd[0] {
		case "START":
			client.Write("OK WELCOME")
			// TODO start game routine
		case "EXIT":
			client.Write("OK BYE")
			Log("Closing connection to " +
				client.conn.RemoteAddr().String())
			client.conn.Close()
		default:
			client.Write("ERR UNKWNCMD")
		}
	}
}

func NewClient(conn net.Conn) *Client {
	var c Client
	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)

	client := &GameClient{
		conn: conn,
		reader: reader,
		writer: writer,
	}

	client.Listen()
	c = client

	return &c
}

func handleConnection(conn net.Conn) {
	Log("Incoming connection from " + conn.RemoteAddr().String())
	NewClient(conn)
}

func Log(msg string) {
	fmt.Println(time.Now().Format("2006-01-02 15:04"), msg)
}

