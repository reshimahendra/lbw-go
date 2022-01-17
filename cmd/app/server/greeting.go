/*
   Package server
       greeting.go contain routine to show greeting message
       when server executed
*/
package server

import (
	"fmt"
	"strings"
)

// 'Welcome' is for console 'greetings' when the server is starting
func welcome(greetings, server string, fillerChar string, lineLength int) {
    fmt.Println()
    fmt.Println(printFiller(fillerChar, lineLength))
    fmt.Println(formatGreeting(greetings, lineLength))
    fmt.Println(formatGreeting(server, lineLength))
    fmt.Println(printFiller(fillerChar, lineLength))
    fmt.Println()
}

// formatGreeting will formating our server greeting text
func formatGreeting(greeting string, lineLength int) string {
    var fillerGreeting string
    var fillerLen int

    if len(greeting)%2 == 0 {
        fillerLen = (lineLength - len(greeting)) / 2
    } else {
        fillerLen = ((lineLength - len(greeting)) / 2) - 1
    }
    for i := 0; i < fillerLen; i++ {
        fillerGreeting += " "
    }

    return fmt.Sprintf("%s%s",fillerGreeting, strings.ToUpper(greeting))
}

// printFiller will creating a horizontal line based on given string length
func printFiller(fillerChar string, lineLength int) string {
    // print filler char 
    var filler string
    for i := 0; i <= lineLength; i++ {
            filler += fillerChar
    }
    return filler
}
