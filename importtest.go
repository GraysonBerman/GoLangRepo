//This code is testing multiple functions: pinging, reading and recording input from terminal, printing out to terminal

//https://gobyexample.com/command-line-arguments
//https://tutorialedge.net/golang/reading-console-input-golang/

package main

import (
  "bufio"
  "strings"
  "os"
  "fmt" //main package stuff (format)
  "github.com/sparrc/go-ping" //ping packages
)


func main(){
  fmt.Println("Type in the website's name that you want to ping and then hit enter")
  reader := bufio.NewReader(os.Stdin)
  fmt.Println("--------------")

  //The for loop below is under construction
  //I'm learning how to read and store input from terminal
  for {
    fmt.Print("->")
    text, _ := reader.ReadString('\n')
    //convert CRLF to LF
    text = strings.Replace(text, "\n", "", -1)

    if strings.Compare("hi", text) == 0{
      fmt.Println("hello, Yourself")
    } else{
    fmt.Println("You said " + text)
  }
  }

  pinger, err := ping.NewPinger("my.chapman.edu") //pings this website
  if err != nil{ //catches if there is an error pinging the website
    panic(err)
  }

  pinger.Count = 3 //ping 3 times
  pinger.Run() //run pinger
  stats := pinger.Statistics()
  fmt.Printf("%v\n", stats.Addr)//stats = statistics, the struct is in github.com (has ping stuff)
}
