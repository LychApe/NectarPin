package tools

import "log"

func CheckConfNil(values any, msg string) {
	if values == nil || values == "" {
		log.Fatalln("[NectarPin]: " + msg + "未配置")
	}
}
