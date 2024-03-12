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
        key := "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"

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
