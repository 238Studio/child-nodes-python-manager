package py

import "github.com/238Studio/child-nodes-net-service/ws"

// PythonManager python管理器对象
type PythonManager struct {
	Websocket *ws.ModelMessageChan //websocket服务
}

// PythonScript 脚本结构体
type PythonScript struct {
	ScriptName string //脚本名
	state      int    //脚本状态。分为未加载、已加载、锁定三种状态
}

// PythonScriptRequest 请求结构体
type PythonScriptRequest struct {
	priority int //指令优先级，0～255,越小优先级越高
	index    int //索引在队列中的下标
}
