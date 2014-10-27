package chat

import (
	"fmt"
	"net"
	"time"
	"strconv"
)


type Chat struct {
	server *Server
}

func (c *Chat) Start(){
	go c.server.Start()
	for{
		var content string
		select {
		case content=<-c.server.readChan:
			fmt.Println(content)
		}
	}
}
func (c *Chat) Connect(){
	c.server.Connect()
}

func (c *Chat) ConnectPort(port string){
	c.server.ConnectPort(port)
}

func (c *Chat) Write(content string){
	c.server.Write(content)
}

type Server struct {
	port       string
	conn net.Conn
	remoteAddr string
	accepted chan bool
	connected chan bool
	writeChan chan string
	readChan chan string
}
////////////////////////////////////////////////////////
//
//错误检查
//
////////////////////////////////////////////////////////
func checkError(err error, info string) (res bool) {

	if err != nil {
		fmt.Println(info + "  " + err.Error())
		return false
	}
	return true
}


func (s *Server)Read()(){
	buf := make([]byte, 1024)
	for {
		length, err := s.conn.Read(buf)
		if checkError(err, "Connection") == false {
			s.conn.Close()
			return
		}
		if length > 0 {
			buf[length] = 0
		}
		//fmt.Println("Rec[",conn.RemoteAddr().String(),"] Say :" ,string(buf[0:lenght]))
		recieveStr := string(buf[0:length])
		//fmt.Println("receive message:"+recieveStr)
		s.readChan<-recieveStr
	}
}

func (s *Server) Write(content string){
	s.writeChan<-content
}
func (s *Server)WriteImpl(content string){
	username := s.conn.LocalAddr().String()
	_, err := s.conn.Write([]byte(username + " Say :" + content))
		//fmt.Println(lens)
	if err != nil {
	fmt.Println(err.Error())
		s.conn.Close()
		}


}

func (s *Server) Start(){
	go s.Listen()
	var content string
	<-s.connected
	go s.Read()
	for {
		select {
		case content=<-s.writeChan:
			s.WriteImpl(content)
		}
	}
}

func (s *Server) Listen() {
	service := ":" + s.port //strconv.Itoa(port);
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err, "ResolveTCPAddr")
	l, err := net.ListenTCP("tcp", tcpAddr)
	if err!=nil{
		return
	}
	checkError(err, "ListenTCP")



		fmt.Println("Listening ..."+tcpAddr.String())
		conn, err := l.Accept()
		checkError(err, "Accept")
		fmt.Println("Accepting ...")
		s.conn=conn
	s.connected<-true
}
func (s *Server) Connect(){
	s.ConnectPort(s.remoteAddr)
}

func (s *Server) ConnectPort(port string){
	tcpAddr, err := net.ResolveTCPAddr("tcp4", port)
	_,err=net.ResolveTCPAddr("tcp4","localhost:"+s.port)
	checkError(err, "ResolveTCPAddr")
	conn, err := net.DialTCP("tcp", nil, tcpAddr)

	for err!=nil{
		time.Sleep(1*time.Second)
		conn, err = net.DialTCP("tcp", nil, tcpAddr)
		fmt.Printf("connection fail from, to %v,because %v\n",tcpAddr,err)
	}
	s.conn=conn
	fmt.Printf("connected to ",tcpAddr)
	s.connected<-true
}

func NewChat(listen string) (*Chat){
	return &Chat{
		server:&Server{
			port:listen,
			accepted: make(chan bool,1),
			connected:make(chan bool,1),
			writeChan:make(chan string, 10),
			readChan :make(chan string, 10),
		},
	}

}
////////////////////////////////////////////////////////
//
//主程序
//
//参数说明：
//	启动服务器端：  Chat server [port]				eg: Chat server 9090
//	启动客户端：    Chat client [Server Ip Addr]:[Server Port]  	eg: Chat client 192.168.0.74:9090
//
////////////////////////////////////////////////////////
func main() {
	var t string
	fmt.Scanln(&t)

	quit :=make(chan string)
	var chat *Chat

	switch t{
		case "1":
		chat = NewChat("37143")
		go chat.Start()
		go chat.server.Listen()

		chat.server.Write("from 1")
	case "2":
		chat = NewChat("37145")
		go chat.Start()
		go chat.server.Connect()

		count:=0
		for count<=10{

			chat.server.Write("from 2 "+strconv.Itoa(count)+"\n")
			count++
		}

	}

	<-quit
}
