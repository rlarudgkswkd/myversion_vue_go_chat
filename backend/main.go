// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"log"
	"net/http"
)

var addr = flag.String("addr", ":8090", "http service address") //플래그 선언 , package flag => command-line을 parsing해주기 위함.

func main() {
	flag.Parse()    // 플래그 선언 후에는 이렇게 명령어를 나누기 위해 이렇게 call 을 해주어야함.
	hub := newHub() //hub를 선언 및 초기화
	go hub.run()    //고루틴으로 hub 실행

	//Handle and HandleFunc add handlers to defaultservemux
	//func(w http.ResponseWriter, r *http.Request) 이게 핸들러인가봄?
	//ㅇㅇ 맞음 너거 뜯어보니 go가 매개변수안에 함수가 가능했음.
	//http.handler(...)인데 func 하고 변수때리는 형태인듯.
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, w, r) //peer로부터 request 처리
	})

	//ListenAndServe starts an HTTP server with a given address and handler.
	//ListenAndServe 는 주어진 주소와 핸들러로 http server를 시작한다.
	//nil 인거 보니 DefaultServeMux
	err := http.ListenAndServe(*addr, nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
