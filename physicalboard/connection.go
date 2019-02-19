package physicalboard

import (
	"fmt"
	"io"
	"log"
	"strings"
	"time"

	"github.com/jacobsa/go-serial/serial"
)

var options = serial.OpenOptions{
	PortName:        "/dev/cu.usbmodem1431",
	BaudRate:        9600,
	DataBits:        8,
	StopBits:        1,
	MinimumReadSize: 4,
}

// SendCommand sends command to arduino and
// waits for arduino to reply
func SendCommand(command string) {
	port, err := serial.Open(options)
	// for now, needs to wait for arduino to reset when opening connection
	time.Sleep(time.Second * 2)
	if err != nil {
		log.Fatalf("serial.Open: %v", err)
	}
	defer port.Close()

	_, err = port.Write([]byte(command + "\n"))
	if err != nil {
		log.Fatalf("port.Write: %v", err)
	}

	waitForFinish(port)
}

func waitForFinish(port io.ReadWriteCloser) {
	for {
		s := readLine(port)
		if s == "finished" {
			return
		}
	}
}

func readLine(port io.ReadWriteCloser) string {
	word := ""
	buf := make([]byte, 1)
	for {
		_, err := port.Read(buf)
		if err != nil {
			fmt.Println("Error while waiting for command to return")
			return ""
		}
		if buf[0] == '\n' {
			// special characters sent by arduino
			word = strings.Replace(word, string(rune(13)), "", -1)
			fmt.Printf("The word is %s\n", word)
			return word
		}
		word = word + string(buf[0])
	}
}
