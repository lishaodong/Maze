package chat

import (
	"testing"
)

func TestChat(){
	go StartServer("37143")

	go StartClient("115.27.91.133:37143")


}
