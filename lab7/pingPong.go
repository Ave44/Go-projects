package main

import (
    "fmt"
    "strconv"
)

type Message struct {
    value string
    receiverID int
}

func player(id int, opponentID int, channel chan Message, quit chan bool, lastMessage int) {
    fmt.Println(id, "zaczynam grę")
    text := strconv.Itoa(id) + " sending to " + strconv.Itoa(opponentID)
    myMessage := Message{value: text, receiverID: opponentID}
    received := 0
    for {
        select {
        case mes := <- channel:
            if mes.receiverID == id {
                if received >= lastMessage {
                fmt.Println(id, " kończę grę")
                    quit <- true
                } else {
                    fmt.Println(id, " odebrałem, odbijam do ", opponentID)
                    received += 1
                    channel <- myMessage
                }
            } else {
                channel <- mes
            }
        }
    }
    quit <- true
}

func main() {
    c := make (chan Message)
    quit := make (chan bool)

    go player(1, 2, c, quit, 10)
    go player(2, 3, c, quit, 10)
    go player(3, 1, c, quit, 10)
    c <- Message{value: "ping", receiverID: 1}

    <- quit
}