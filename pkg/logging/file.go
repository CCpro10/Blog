package logging

import (
	"Blog/pkg/file"
	"Blog/pkg/setting"
	"fmt"
	"os"
	"time"
)

//获取相对短路径目录
func getLogFilePath() string {
	return fmt.Sprintf("%s%s", setting.AppSetting.RuntimeRootPath, setting.AppSetting.LogSavePath)
}

//获取日志文件名
func getLogFileName() string {

	return fmt.Sprintf("%s%s.%s",
		setting.AppSetting.LogSaveName,
		time.Now().Format(setting.AppSetting.TimeFormat),
		setting.AppSetting.LogFileExt,
	)

}

func openLogFile(fileName, filePath string) (*os.File, error) {
	dir, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("os.Getwd err: %v", err)
	}

	src := dir + "/" + filePath
	perm := file.CheckPermission(src)
	if perm == true {
		return nil, fmt.Errorf("file.CheckPermission Permission denied src: %s", src)
	}

	err = file.IsNotExistMkDir(src)
	if err != nil {
		return nil, fmt.Errorf("file.IsNotExistMkDir src: %s, err: %v", src, err)
	}

	f, err := file.Open(src+fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, fmt.Errorf("Fail to OpenFile :%v", err)
	}

	return f, nil
}

////打开日志文件,没有就创建,返回文件句柄
//func openLogFile(filePath string) *os.File {
//	//根据文件路径寻找文件
//	_, err := os.Stat(filePath)
//	switch {
//	case os.IsNotExist(err):
//		mkDir()//如果文件不存在则会创建目录
//	case os.IsPermission(err):
//		log.Fatalf("Permission :%v", err)
//	}
//	//创建文件
//	handle, err := os.OpenFile(filePath, os.O_APPEND | os.O_CREATE | os.O_WRONLY, 0644)
//	if err != nil {
//		log.Fatalf("Fail to OpenFile :%v", err)
//	}
//
//	return handle
//}

////创建目录路径
//func mkDir() {
//	//获取根路径
//	dir, _ := os.Getwd()
//
//	//log.Println(dir) 因为在init函数中被调用,所以根路径dir为 ..\Blog
//	//创建目录
//	err := os.MkdirAll(dir + "/" + getLogFilePath(), os.ModePerm)
//	if err != nil {
//		panic(err)
//	}
//}
