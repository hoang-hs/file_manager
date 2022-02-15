package hooks

import (
	"file_manager/src/common/notice"
	"fmt"
	"go.uber.org/zap/zapcore"
)

type HookProcessor struct {
	env     string
	jobPool *notice.JobPool
}

func NewHookProcessor(env string) *HookProcessor {
	return &HookProcessor{
		env:     env,
		jobPool: notice.GlobalJobPool,
	}
}

func (h *HookProcessor) ProcessEvent(entry zapcore.Entry) {
	message := fmt.Sprintf("level: %s \n message: %s \n caller: %s \n line: %v", entry.Level.String(), entry.Message, entry.Caller.File, entry.Caller.Line)
	template := notice.NewTemplate("file_manager "+h.env, message)
	p := notice.NewPackage(notice.GlobalChannel, template)
	notice.NewPackage(notice.GlobalChannel, template)
	notice.GlobalJobPool.AddPackage(p)
}
