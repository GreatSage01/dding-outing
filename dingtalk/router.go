package dingtalk

import (
	"app/utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

var cmdTable = make(map[string]*command)

type command struct {
	executor ExecFunc
	isAdmin  bool // 是否需要管理员权限
	arity    int  // 参数个数
}

func RegisterCommand(name string, execFunc ExecFunc, arity int, isAdmin bool) {
	cmdTable[name] = &command{
		executor: execFunc,
		arity:    arity,
		isAdmin:  isAdmin,
	}
}

func validateArity(arity int, cmdArgs []string) bool {
	argNum := len(cmdArgs)
	if arity >= 0 {
		return argNum == arity
	}
	return argNum >= -arity
}

func execDingCommand(msg outGoingModel) []byte {

	content := msg.Text.Content
	keyAndArgs := strings.Split(strings.TrimSpace(content), " ")
	fmt.Printf("keyAndArgs: %s\n", keyAndArgs)
	cmdName := strings.ToLower(keyAndArgs[0])
	fmt.Printf("cmdName: %s\n", cmdName)
	cmd, ok := cmdTable[cmdName]

	// fmt.Printf("cmd: %+v", cmd)
	// fmt.Print("\n")

	fmt.Printf("keyAndArgs[1:]: %v \n", keyAndArgs[1:])

	if !ok {
		return NewTextMsg("ERR: unregistered command '" + cmdName + "'").Marshaler()
	}
	if !validateArity(cmd.arity, keyAndArgs) {
		return NewTextMsg("ERR: wrong number of arguments for '" + cmdName + "' command").Marshaler()
	}
	if cmd.isAdmin && msg.IsAdmin != cmd.isAdmin {
		return NewTextMsg("ERR '" + cmdName + "': you have no right to do this operation").Marshaler()
	}
	return cmd.executor(keyAndArgs[1:])
}

// MyHandler实现Handler接口
type OutGoingHandler struct{}

func (h *OutGoingHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	body := r.Body
	var err error
	var buf []byte
	obj := outGoingModel{}
	buf, err = ioutil.ReadAll(body)
	if err != nil {
		return
	}

	err = json.Unmarshal(buf, &obj)
	fmt.Printf("msg: %v", utils.Struct2json(obj))
	fmt.Printf("\n")

	if err != nil {
		return
	}
	msg := execDingCommand(obj)
	if msg == nil {
		return
	}
	_, _ = w.Write(msg)
}
