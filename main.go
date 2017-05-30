package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

var changelog = "Changelog"

func main() {
	a := flag.String("a", "", "Author name")
	m := flag.String("m", "", "Commit message")
	flag.Parse()
	Wlog(*a, *m)
}

func MultiMsg() error {

	logFile, err := os.OpenFile(changelog, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer logFile.Close()
	w := bufio.NewWriter(logFile)

	scanner := bufio.NewScanner(os.Stdin)
	var text string
	fmt.Println("******************************")
	fmt.Println("* You can Exit by type \":wq\" *")
	fmt.Println("******************************")
	fmt.Println("Enter your Changelog message: ")
	for text != ":wq" { // break the loop if text == "q"
		scanner.Scan()
		text = scanner.Text()
		if text != ":wq" {
			fmt.Fprintf(w, "%s\n", text)
			w.Flush()
		}
	}
	return nil
}

func Wlog(author string, msg string) error {
	logFile, err := os.OpenFile(changelog, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	w := bufio.NewWriter(logFile)
	defer logFile.Close()
	log.SetOutput(logFile)
	if author != "" {
		log.Println(author)
	} else {
		fmt.Print("Enter Author name: ")
		var input string
		fmt.Scanln(&input)
		log.Printf("%s", input)
	}

	if msg != "" {
		fmt.Fprintf(w, "- %s\n", msg)
		w.Flush()
	} else {
		MultiMsg()
	}
	//fmt.Println(msg)
	return nil
}
