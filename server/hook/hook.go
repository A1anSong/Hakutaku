package hook

import (
	"fmt"

	log "github.com/sirupsen/logrus"
)

type HookFn func(interface{}) error

type HookType string

var (
	InitHook       HookType = "init"
	LoadConfigHook HookType = "load_config"
	LoadDBHook     HookType = "load_db"
	CloseHook      HookType = "close"

	hooks = map[HookType][]HookInfo{}
)

type HookInfo struct {
	Name string
	Fn   HookFn
}

func Register(hookType HookType, name string, fn HookFn) {
	info := HookInfo{Name: name, Fn: fn}
	v, ok := hooks[hookType]
	if ok {
		hooks[hookType] = append(v, info)
		return
	}
	hooks[hookType] = []HookInfo{info}
}

func OnEvent(evt HookType, val interface{}) (err error) {
	hooks, ok := hooks[evt]
	if !ok {
		err = fmt.Errorf("no hook register:%s", evt)
		return
	}
	for _, v := range hooks {
		err = v.Fn(val)
		if err != nil {
			log.Errorf("%s hook call %s failed: %s", evt, v.Name, err.Error())
			return
		}
	}
	return
}
