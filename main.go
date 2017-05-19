// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
// Additions and modifications Copyright 2017 Lance Hartung.

package main

import (
	"flag"
	"log"
	"net"
	"net/http"
)

var udpAddr = flag.String("udp", ":4000", "UDP listening address")
var httpAddr = flag.String("http", ":8080", "HTTP service address")

func serveHome(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	if r.URL.Path != "/" {
		http.Error(w, "Not found", 404)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", 405)
		return
	}
	http.ServeFile(w, r, "home.html")
}

func handleUDPMessages(hub *Hub) {
	serverAddr, err := net.ResolveUDPAddr("udp", *udpAddr)
	if err != nil {
		log.Fatal("ResolveUDPAddr: ", err)
	}

	conn, err := net.ListenUDP("udp", serverAddr)
	if err != nil {
		log.Fatal("ListenUDP: ", err)
	}

	buffer := make([]byte, 2048)

	for {
		n, _, err := conn.ReadFromUDP(buffer)
		if err != nil {
			log.Fatal("ReadFromUDP: ", err)
		} else {
			hub.broadcast <- buffer[0:n]
		}
	}
}

func main() {
	flag.Parse()
	hub := newHub()
	go hub.run()
	go handleUDPMessages(hub)
	http.HandleFunc("/", serveHome)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, w, r)
	})
	err := http.ListenAndServe(*httpAddr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
