package analysis

import (
	"errors"
	"os"
	"bufio"
	"strings"
	"regexp"
	"io/ioutil"
	"io"
)


/*
*	@name获取目录下所有文件的image地址
*	@param	string path 目录地址
*	@return []string	图片地址slice集合
*/
func GetImgSrc(path string) ([]string, error) {
	
	if !isDir(path) {
		return nil, errors.New("无法获取目录" + path)
	}
	
	fileList, err := readPath(path)
	
	if err != nil {
		return nil, err
	}
	
	var imgList = make([]string, 0);
	
	for i, num := 0, len(fileList); i < num; i++ {
		img, err := pregMatchImg(fileList[i])
		if err != nil {
			return nil, err
		}
		
		imgList = append(imgList, img...)
	}
	
	return imgList, nil;
}

func isDir(path string) bool {
	_, err := ioutil.ReadDir(path)
	
	if nil == err {
		return true
	}
	return false
}

func readPath(path string) ([]string, error){
	f, err := os.Open(path)
	
	if nil != err {
		return nil, errors.New("打开文件" + path + "失败")
	}
	defer f.Close()
	
	list, err := f.Readdirnames(-1)
	if nil != err {
		return nil, errors.New("读取文件夹" + path + "失败")
	}
	
	fileList := make([]string, 0)
	
	for i, num := 0, len(list); i < num; i++ {
		//忽略隐藏文件或文件夹
		if '.' == list[i][0] {
			continue
		}
		newPath := path + "/" + list[i]
		if true == isDir(newPath) {
			temp, err := readPath(newPath)
			if err != nil {
				return nil, err
			}
			fileList = append(fileList, temp...)
		} else {
			fileList = append(fileList, newPath)
		}
	}
	
	return fileList, nil
}

func pregMatchImg(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, errors.New("打开文件" + path + "失败")
	}
	defer file.Close()
	
	buf := bufio.NewReader(file)
	result := make([]string, 0)
	
	for {
		line, err := buf.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break;
			} else {
				return nil, errors.New("读取文件" + path + "发生错误")
			}
		}
		//去除前后空格和转化小写
		line = strings.TrimSpace(line)
		line = strings.ToLower(line)
		reg := regexp.MustCompile(`http:\/\/img\.m\.baidu\.com\/.+?(?:png|jpg|jpeg|gif|'|"| )`)
		imgsrc := reg.FindAllString(line, -1)
		result = append(result, imgsrc...)
	}
	return result, nil
}