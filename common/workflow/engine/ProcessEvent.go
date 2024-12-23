package engine

import (
	"Scheduler_go/common/workflow/modelx"
	"fmt"
	"reflect"
)

type method struct {
	S interface{}    //method所在的struct，这是函数执行的第一个参数
	M reflect.Method //方法
}

// 事件池，所有的事件都会在流程引擎启动的时候注册到这里
var EventPool = make(map[string]method)

// 事件出错，则可能导致流程无法运行下去,在这里添加选项，是否忽略事件出错，让流程继续
var IgnoreEventError bool

// 注册一个struct中的所有func
// 注意,此时不会验证事件方法参数是否正确,因为此时不知道事件到底是“节点事件”还是“流程事件”
func RegisterEvents(Struct any) {
	StructValue := reflect.ValueOf(Struct)
	StructType := StructValue.Type()

	for i := 0; i < StructType.NumMethod(); i++ {
		m := StructType.Method(i)
		var method = method{Struct, m}
		EventPool[m.Name] = method
	}
}

// 验证流程事件(目前只有流程撤销事件)参数是否正确
// 流程撤销事件  func签名必须是func(struct *interface{}, ProcessInstanceID int,RevokeUserID string) error
func verifyProcEventParameters(m reflect.Method) error {
	//自定义函数必须是3个参数，参数0：*struct{} 1:int 2:String
	if m.Type.NumIn() != 3 || m.Type.NumOut() != 1 {
		fmt.Errorf("warning:事件方法 %s 入参、出参数量不匹配,此函数无法运行", m.Name)
	}

	if m.Type.In(1).Kind().String() != "int" {
		fmt.Errorf("warning:事件方法 %s 参数1不是int类型,此函数无法运行", m.Name)
	}

	if m.Type.In(2).Kind().String() != "string" {
		fmt.Errorf("warning:事件方法 %s 参数2不是string类型,此函数无法运行", m.Name)
	}

	if !TypeIsError(m.Type.Out(0)) {
		fmt.Errorf("warning:事件方法 %s 返回参数不是error类型,此函数无法运行", m.Name)
	}
	return nil
}

// 验证节点事件(1、节点开始  2、节点结束 3、任务结束)参数是否正确
// 1、节点开始、结束事件     func签名必须是func(struct *interface{}, ProcessInstanceID int, CurrentNode *Node, PrevNode Node) error
// 2、任务完成事件          func签名必须是func(struct *interface{}, TaskID int, CurrentNode *Node, PrevNode Node) error
func verifyNodeEventParameters(m reflect.Method) error {
	//自定义函数必须是4个参数，参数0：*struct{} 1:int 2:Node 3:Node
	if m.Type.NumIn() != 4 || m.Type.NumOut() != 1 {
		fmt.Errorf("warning:事件方法 %s 入参、出参数量不匹配,此函数无法运行", m.Name)
	}

	if m.Type.In(1).Kind().String() != "int" {
		fmt.Errorf("warning:事件方法 %s 参数1不是int类型,此函数无法运行", m.Name)
	}

	if m.Type.In(2).ConvertibleTo(reflect.TypeOf(&modelx.Node{})) != true {
		fmt.Errorf("warning:事件方法 %s 参数2不是*Node类型,此函数无法运行", m.Name)

	}

	if m.Type.In(3).ConvertibleTo(reflect.TypeOf(modelx.Node{})) != true {
		fmt.Errorf("warning:事件方法 %s 参数3不是Node类型,此函数无法运行", m.Name)
	}

	if !TypeIsError(m.Type.Out(0)) {
		fmt.Errorf("warning:事件方法 %s 返回参数不是error类型,此函数无法运行", m.Name)
	}
	return nil
}
