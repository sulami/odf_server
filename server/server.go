package server

import (
	"bufio"
	"errors"
	"strconv"
	"strings"
	"net"
)

import "github.com/sulami/odf_server/log"

type Server struct {
	Port	int
	Online	bool
}

func (s *Server) Listen() (err error) {
	log.Log("Server starting up...")

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
			log.Log(e.Error())
		}
		go handleConnection(conn)
	}

	s.Online = true
	log.Log("Listening on port " + strconv.Itoa(s.Port))

	return
}

func (s *Server) StopListening() (err error) {
	log.Log("Stopping server...")

	if !s.Online {
		err = errors.New("Server is not running")
		return
	}

	s.Online = false
	log.Log("Server stopped")

	return
}

type Client struct {
	conn net.Conn
	incoming chan string
	outgoing chan string
	reader *bufio.Reader
	writer *bufio.Writer
}

func (client *Client) Read() {
	for {
		line, _ := client.reader.ReadString('\n')
		client.incoming <- line
	}
}

func (client *Client) Write() {
	for data := range client.outgoing {
		client.writer.WriteString(data)
		client.writer.Flush()
	}
}

func (client *Client) Listen() {
	go client.Read()
	go client.Write()
}

func NewClient(conn net.Conn) *Client {
	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)

	client := &Client{
		conn: conn,
		incoming: make(chan string),
		outgoing: make(chan string),
		reader: reader,
		writer: writer,
	}

	client.Listen()

	return client
}

func handleConnection(conn net.Conn) {
	log.Log("Incoming connection from " + conn.RemoteAddr().String())

	client := NewClient(conn)

	go func() {
		for  {
			data := <-client.incoming
			parseCmd(client.conn, data)
		}
	} ()

}

func parseCmd(conn net.Conn, data string) {
	cmd := strings.Fields(data)
	if len(cmd) != 0 {
		switch cmd[0] {
		case "LOGIN":
			if len(cmd) != 3 {
				conn.Write([]byte("ERR ARGS\n"))
				conn.Close()
			}
			conn.Write([]byte("OK WELCOME\n"))
			// TODO find the user and try to auth him
		case "LOGOUT":
			conn.Write([]byte("OK BYE\n"))
			conn.Close()
		default:
			conn.Write([]byte("ERR UNKWNCMD\n"))
			conn.Close()
		}
	}
}

