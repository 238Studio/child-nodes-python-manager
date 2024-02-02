package py

const (
	UnLoad = iota //脚本未加载
	Load          //脚本已加载，可并发运行状态
	Lock          //脚本锁定运行。不可并发运行状态
)
