package main

import (
	"fmt"
)

// 从用户端获取脚本的启动参数
func getManagerConfig() ManagerConfig {
	refreshTime := 0
	imgPath := ""

	fmt.Println("请输入模版图片文件夹，如img，不输入默认img文件夹")
	fmt.Scanln(&imgPath)

	if imgPath == "" {
		imgPath = "img"
	}

	imgPath = "./" + imgPath + "/"

	fmt.Println("请输入刷新间隔：（单位毫秒, 不输入默认500毫秒，建议使用500毫秒）")
	fmt.Scanln(&refreshTime)

	if refreshTime > 5000 || refreshTime < 0 {
		panic("refreshTime 过于离谱")
	}

	if refreshTime == 0 {
		refreshTime = 500
	}

	return ManagerConfig{RefreshTime: refreshTime, ImgPath: imgPath}
}

func main() {
	cfg := getManagerConfig()
	manager := NewManager(cfg)

	manager.work()
}
