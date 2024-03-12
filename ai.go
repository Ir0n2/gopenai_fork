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

	prompter("Give me a good country song")
	format()


}

func format() {

        num := printLines("testlogfile")

        for i := 0; num-2 >= i; i++ {

                str := readList("testlogfile", i)
                //str is a line from testlogfile and test is str but split up at every curly bracket
                test := strings.Split(str, "{")

                //res2 removes the remaining bracket garbage.
                res2 := strings.TrimSuffix(test[3], "} }]}")

                //this print statement prints just the word
                fmt.Printf(res2)

        }
        fmt.Println("")

	deleteFile("testlogfile")
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
