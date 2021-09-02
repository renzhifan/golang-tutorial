package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {

	fileName := "./storage/" + time.Now().Format("2006-01-02") + ".log"
	fmt.Println(fileName)
	logFile, err := os.Create(fileName)
	defer logFile.Close()
	if err != nil {
		log.Fatalln("create file log.log failed")
	}

	logger := log.New(logFile, "[Debug]", log.Lshortfile)
	logger.Println("debug info is ,check list ,hello")
	logger.Println("debug info is ,check list ,hello000111")
	//logger.SetPrefix("[Info]")
	logger.SetFlags(log.Ltime)
	logger.SetOutput(logFile)
	logger.Print("Info check")
	logger.SetOutput(os.Stdout)
	logger.Print("Info check stdout")
}
