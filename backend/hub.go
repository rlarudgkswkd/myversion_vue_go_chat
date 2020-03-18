// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

// Hub maintains the set of active clients and broadcasts messages to the
// clients. (HUB는 지금 활동중인 고객과 Broadcast를 유지시킨다.)
type Hub struct { //Hub의 구조체 선언
	// Registered clients.
	clients map[*Client]bool // keytype *Client, Valuetype bool인 clients 선언

	// Inbound messages from the clients. (client로부터 받은 메세지, 형태는 byte)
	broadcast chan []byte

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client
}

func newHub() *Hub { //hun구조체 생성자 함수 인듯?, 구조체 필드를 사용전에 초기화 하기위한
	return &Hub{
		//channel과 map의 초기화는 make로, channel은 "<-"로 데이터를 보내거나 받음
		broadcast:  make(chan []byte),
		register:   make(chan *Client), //
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}
}

func (h *Hub) run() { //메소드 정의 (h *hub) , hub의 값을 직접 다루기 위해 포인터로 구현
	for {
		select {
		case client := <-h.register: //client에서 register채널로부터 수신
			h.clients[client] = true
		case client := <-h.unregister: //unregisterd clients sending
			if _, ok := h.clients[client]; ok { //조건문 앞에 간단한 statement 입력하여 bool 판명
				delete(h.clients, client) //unregi client hub delete
				close(client.send)        //close the channel
			}
		case message := <-h.broadcast: // message에서 broadcast채널로부터 수신
			for client := range h.clients { //클라이언트 별로 다 뿌려주기
				select {
				case client.send <- message:
				default: //default가 되어있기 때문에 대기없이 진행되고 있음
					close(client.send) //outbound message가 없기에 없는 client.send채널은 종료
					delete(h.clients, client)
				}
			}
		}
	}
}
