package py

import (
	"github.com/238Studio/child-nodes-assist/arch"
	"github.com/238Studio/child-nodes-net-service/ws"
)

// Init 初始化Python管理器
// 传入:无
// 返回:PythonManager对象
func Init() *PythonManager {
	//初始化python管理器
	pythonManager := new(PythonManager)

	//初始化websocket服务
	wsService := arch.Websocket.(*ws.WebsocketServiceApp)              //从arch全局变量中获取websocket服务对象
	pythonManager.Websocket = wsService.InitModelMessageChan("python") //初始化模型消息通道

	return pythonManager
}
