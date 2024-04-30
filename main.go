package main

import {
    "log"
    "chat-server-golang/network"
}


func init() {
    log.Println("start golang-server!!")
}

func main() {
    n := network.NewServer()

    n.StartServer()
}