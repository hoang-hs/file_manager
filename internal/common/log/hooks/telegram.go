package hooks

import (
	"file_manager/internal/common/notice"
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
	message := fmt.Sprintf("%s \n %s \n %s", entry.Level.String(), entry.Message, entry.Stack)
	template := notice.NewTemplate("file_manager "+h.env, message, notice.MarkdownV2)
	//p := notice.NewPackage(notice.GlobalChannel, template)
	notice.NewPackage(notice.GlobalChannel, template)
	//notice.GlobalJobPool.AddPackage(p)
}
