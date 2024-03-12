# GopenAI - golang open ai client

## Features
- supports stream/sse

## How to use

Copy & past your api key and put your prompt where it says content.

basic usage:

```go
package main

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/blackestwhite/gopenai"
)

func main() {
	key := "YOUR-OPEN-AI-KEY"

	instance := gopenai.Setup(key)

	p := gopenai.ChatCompletionRequestBody{
		Model: "gpt-3.5-turbo",
		Messages: []gopenai.Message{
			{Role: "user", Content: "hi"},
		},
		Stream: true,
	}

	resultCh, err := instance.GenerateChatCompletion(p)
	if err != nil {
		log.Fatal(err)
	}

	for chunk := range resultCh {
		log.Println(chunk)
	}
}
```

with custom http client

```go
package main

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/blackestwhite/gopenai"
	"golang.org/x/net/proxy"
)

func main() {
	key := "YOUR-OPEN-AI-KEY"

    // open ai is blocked in my country so i use socks5 proxy to consume it
	transport := &http.Transport{}
	dialer, err := proxy.SOCKS5("tcp", "localhost:8586", nil, proxy.Direct)
	if err != nil {
		log.Fatal(err)
	}
	transport.DialContext = func(ctx context.Context, network, address string) (net.Conn, error) {
		return dialer.Dial(network, address)
	}

	client := &http.Client{Transport: transport}

	instance := gopenai.SetupCustom(key, client)

	p := gopenai.ChatCompletionRequestBody{
		Model: "gpt-3.5-turbo",
		Messages: []gopenai.Message{
			{Role: "user", Content: "hi"},
		},
		Stream: true,
	}

	resultCh, err := instance.GenerateChatCompletion(p)
	if err != nil {
		log.Fatal(err)
	}

	for chunk := range resultCh {
		log.Println(chunk)
	}
}
```

## Improved output

Here's the code for the improved output, Just use the promptGPT() function to prompt chatgpt and it will return its output as a string.
so you can prompt it easy with fmt.Println(promptGPT("say 'hello world'")). Idk waht else to say you can really do what you want with it.
It's a function for prompting chatgpt, you can do whatever really.

```go

package main

import (
        //"context"
        "log"
        //"net"
        //"net/http"
	"fmt"
	"bufio"
        "strings"
        "os"
        "github.com/blackestwhite/gopenai"
)

func main() {

fmt.Println(promptGPT("What is the best marty robbins song?"))

}

//One final function so everything works within one function call
//Now it returns the GPT output as a string. party party
func promptGPT(prompt string) string {

	prompter(prompt)
	fuck := format()
	return fuck
}

func format() string {
	out := ""	
        num := printLines("testlogfile")
	for i := 0; num-2 >= i; i++ {
		if i == 0 {
			continue
		}			
                str := readList("testlogfile", i)
                //str is a line from testlogfile and test is str but split up at every curly bracket
                test := strings.Split(str, "{")

                //res2 removes the remaining bracket garbage.
                res2 := strings.TrimSuffix(test[3], "} }]}")

                //This will add whatever word is held by res2 to the out string, which is short for output. It's the chatgpt output.
		out += res2		
		//fmt.Println(out)
        }
	fmt.Println("ChatGPT:")
	//delete the log file
	deleteFile("testlogfile")
	return out
}

func check(e error) {

        if e != nil {
                panic(e)
        }

}

//increments i through the list so i should = the amount of items in the list.
func printLines(path string) int {

        i := 0
        filePath := path
        readFile, err := os.Open(filePath)
        check(err)

        fileScanner := bufio.NewScanner(readFile)
        fileScanner.Split(bufio.ScanLines)
        var fileLines []string

        for fileScanner.Scan() {
                fileLines = append(fileLines, fileScanner.Text())
        }

        readFile.Close()

        var line string
        //this print statement is here because we need to do something with line or else it wont compile.
        fmt.Println(line)
        for _, line = range fileLines {

                i++

        }
        return i
}

//Reads list from file
func readList(path string, index int) string {

        filePath := path
        readFile, err := os.Open(filePath)

        check(err)

        fileScanner := bufio.NewScanner(readFile)
        fileScanner.Split(bufio.ScanLines)
        var fileLines []string

        for fileScanner.Scan() {
                fileLines = append(fileLines, fileScanner.Text())
        }

        readFile.Close()

        return fileLines[index]

}
//deletes file
func deleteFile(path string) {

        e := os.Remove(path)
        if e != nil {
                log.Fatal(e)
        }

}


func prompter(prompt string) {
        key := "PUT YOUR KEY HERE"

        instance := gopenai.Setup(key)

        p := gopenai.ChatCompletionRequestBody{
                Model: "gpt-3.5-turbo",
                Messages: []gopenai.Message{
                        {Role: "user", Content: prompt},
                },
                Stream: true,
        }

        resultCh, err := instance.GenerateChatCompletion(p)
        if err != nil {
                log.Fatal(err)
        }

        f, err := os.OpenFile("testlogfile", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
        if err != nil {
            log.Fatalf("error opening file: %v", err)
        }
        defer f.Close()

        log.SetOutput(f)

        for chunk := range resultCh {

                log.Println(chunk)
        }
}

```

Thank you to blackestwhite for making the package.

