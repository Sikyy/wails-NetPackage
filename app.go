package main

import (
	"changeme/internal/define"
	"changeme/internal/services"
	"context"
	"fmt"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

// Network list 枚举网卡
func (a *App) Networklist() interface{} {
	ifacename, err := services.Networklist()
	if err != nil {
		return map[string]interface{}{
			"code": 500,
			"msg":  err.Error(),
		}
	}

	return map[string]interface{}{
		"code": 200,
		"data": ifacename,
	}
}

// 进行数据包的捕获
func (a *App) CaptureTraffic() {

	sessionInfoCh := make(chan define.SessionInfoFront)
	go services.CaptureTraffic(sessionInfoCh)

	for tabelinfo := range sessionInfoCh {
		// 使用 EventsEmit 方法触发事件并传递 tabelinfo 数据
		runtime.EventsEmit(a.ctx, "captureTraffic", tabelinfo)
		fmt.Println("前端渲染数据:", tabelinfo)
		// time.Sleep(time.Hour)
	}

}

// StopCaptureTraffic 停止数据包的捕获
func (a *App) StopCaptureTraffic() {
	services.StopCapture()
}
