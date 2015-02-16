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
	reader *bufio.Reader
	writer *bufio.Writer
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
		switch cmd[0] {
		case "LOGIN":
			if len(cmd) != 3 {
				client.Write("ERR ARGS")
				client.conn.Close()
			}
			client.Write("OK WELCOME")
			// TODO find the user and try to auth him
		case "LOGOUT":
			client.Write("OK BYE")
			client.conn.Close()
		default:
			client.Write("ERR UNKWNCMD")
			client.conn.Close()
		}
	}
}

func NewClient(conn net.Conn) *Client {
	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)

	client := &Client{
		conn: conn,
		reader: reader,
		writer: writer,
	}

	client.Listen()

	return client
}

func handleConnection(conn net.Conn) {
	log.Log("Incoming connection from " + conn.RemoteAddr().String())
	NewClient(conn)
}

