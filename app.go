package main

import (
	"changeme/internal/services"
	"context"
	"fmt"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
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
func (a *App) CaptureTraffic(iface string) {
	services.CaptureTraffic(iface)
}

// StopCaptureTraffic 停止数据包的捕获
func (a *App) StopCaptureTraffic() {
	services.StopCapture()
}
