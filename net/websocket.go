package net

import "github.com/238Studio/child-nodes-net-service/ws"

// PythonManagerWebsocketService 管理python管理器的websocket服务
// 传入:websocket 模块对象
// 返回:无
func PythonManagerWebsocketService(websocketService *ws.ModelMessageChan) {
	go readMessageFromWebsocket(websocketService.ReadMessage)
}

// readMessageFromWebsocket 从websocket中读取消息
// 传入:websocket消息通道
// 返回:无
func readMessageFromWebsocket(read chan ws.WebsocketMessage) {
	for {
		message := <-read
		//TODO:处理消息
	}
}
