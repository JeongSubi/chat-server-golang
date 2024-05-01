package main

import (
    "log"
    "chat-server-golang/network"
)

func init() {
    log.Println("Start init!!")
}

func main() {
    n := network.NewServer()
    n.StartServer()
}