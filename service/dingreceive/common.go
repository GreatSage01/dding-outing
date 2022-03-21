package dingreceive

import (
	"github.com/blinkbean/dingtalk"

	"strings"
	"zx/global"
	"zx/model"
	"zx/utils"
)

type Command struct {
	executor ExecFunc
	isAdmin  bool // 是否需要管理员权限
	arity    int  // 参数个数
}

type ExecFunc func(args []string) []byte

var cmdTable = make(map[string]*Command)

func RegisterCommand(name string, execFunc ExecFunc, arity int, isAdmin bool) {
	cmdTable[name] = &Command{
		executor: execFunc,
		arity:    arity,
		isAdmin:  isAdmin,
	}
}

func (r *DingreceiveService) ExecDingCommand(msg model.OutGoingModel) []byte {
	global.ZX_LOG.Info(utils.Struct2json(msg))
	content := msg.Text.Content
	keyAndArgs := strings.Split(strings.TrimSpace(content), " ")
	cmdName := strings.ToLower(keyAndArgs[0])
	cmd, ok := cmdTable[cmdName]
	if !ok {
		return dingtalk.NewTextMsg("ERR: unregistered command '" + cmdName + "'").Marshaler()
	}
	if !validateArity(cmd.arity, keyAndArgs[1:]) {
		return dingtalk.NewTextMsg("ERR: wrong number of arguments for '" + cmdName + "' command").Marshaler()
	}
	if cmd.isAdmin && msg.IsAdmin != cmd.isAdmin {
		return dingtalk.NewTextMsg("ERR '" + cmdName + "': you have no right to do this operation").Marshaler()
	}
	return cmd.executor(keyAndArgs[1:])
}

//注册命令

func validateArity(arity int, cmdArgs []string) bool {
	argNum := len(cmdArgs)
	if arity >= 0 {
		return argNum == arity
	}
	return argNum >= -arity
}
