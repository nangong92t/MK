package main

import (
	"analysis"
	"fmt"
	"fount"
	"os"
)

func main() {
	fmt.Println(fount.Color("\n从前有座山，山上有个庙，庙里有个老和尚在给小和尚讲故事...\n", "magenta"))

	if len(os.Args) < 2 {
		fmt.Println(fount.Color("Fatal error: 请输入项目绝对路径", "red"))
		os.Exit(-1)
	}
	fmt.Println("正在查找项目中符合要求的图片...\n\n")
	imgList, err := analysis.GetImgSrc(os.Args[1])
	if nil != err {
		fmt.Println(fount.Color("Fatal error: "+err.Error(), "red"))
	}

	for i, num := 0, len(imgList); i < num; i++ {
		fmt.Println(imgList[i], "\n")
	}
	fmt.Println(len(imgList))
}
