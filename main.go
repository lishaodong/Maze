package main

import (
	"github.com/lishaodong/Taipei-Torrent/torrent"
	"log"
	"os"
	"os/signal"
	"strings"
	"fmt"
	"github.com/lishaodong/Maze/chat"
	"strconv"
	"time"
	"net"
)
var (
	launchers map[int]*torrent.Launcher
	chats map[int]*chat.Chat
	ss []string
	local string
	APP_VER = "0.1.1.0227"
)

func main() {
	quit := make(chan bool)
	launchers = make(map[int]*torrent.Launcher)
	chats=make(map[int]*chat.Chat)
	local=GetAddr()

	go func() {
		quitChan := listenSigInt()
		select {
		case <-quitChan:
			log.Printf("got control-C")
			for _,l:=range launchers {
				l.Quit()
			}
			quit<-true
		}
	}()
	ss = make([]string, 10)
	ss[0] = "/Users/dong/configosst.torrent"


	launchers[1] = torrent.NewLauncher(ss[0:1])
	chats[1]=chat.NewChat("37143")
	launchers[2] = torrent.NewLauncher(ss[0:1])
	chats[2]=chat.NewChat("37145")
	go clientStart(1)
	<-quit
}

func clientStart(sn int){
	launcher :=launchers[sn]
	c:=chats[sn]
	var input string
	

	go launcher.Launch()

	go func(){
		var peer string
		for{
			peer=<-launcher.AddPeerChan
			if(strings.Split(peer,":")[0]==local){
				fmt.Println("same ip, rejected")
				continue
			}
			break
			fmt.Println("get Peer:"+peer)
		}
		go c.Start()
		if(true) {
			go c.ConnectPort(changePort(peer))

		}
		i:=0
		for i<0{
			time.Sleep(1*time.Second)
			c.Write("expample form "+strconv.Itoa(sn))
			i++
		}
		go func(){

			fmt.Println("Chat Begin")
			for {
				input=ScanLine()
				c.Write(input)
			}
		}()
	}()


}




func listenSigInt() chan os.Signal {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	return c
}
//
//fmt.Println("new chat, remote address:"+peer.address)
//c:=chat.NewChat("37143",changePort(peer.address))
//go c.Start()
//go c.Connect()
//c.Write("example")


func changePort(address string) (changed string){
	index:=strings.Index(address, ":")
	s:=[]byte(address)[0:index]
	changed =string(s)+":37143"
	fmt.Println("changed port:"+changed)
	return
}

func ScanLine() string {
	var c byte
	var err error
	var b []byte
	for ; err == nil; {
		_, err = fmt.Scanf("%c", &c)

		if c != '\n' {
			b = append(b, c)
		} else {
			break;
		}
	}

	return string(b)
}

func GetAddr() string { //Get ip
	conn, err := net.Dial("udp", "baidu.com:80")
	if err != nil {
		fmt.Println(err.Error())
		return "Erorr"
	}
	defer conn.Close()
	return strings.Split(conn.LocalAddr().String(), ":")[0]
}

