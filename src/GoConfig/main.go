package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//os.Args[0] 程序执行路径， os.Args[1] 必传，项目路径，os.Args[2]可选，当前环境 os.Args[3] 可选，是配置文件路径，默认项目路径
//下的ConfigProfile.xml文件， os.Args[4] 可选，配置文件模板扩展名，默认.tpl
func main() {
	//var absPaths = GetTmpFile("F:\\GoLearn\\src\\GoConfig", ".tpl")
	fmt.Println(os.Args[0])
	var currentEnv = GetEnvironment()
	var configMaps = GetConfig("F:\\GoLearn\\src\\GoConfig\\ConfigProfile.xml")
	configMap := configMaps[currentEnv]
	generateConfigFile("F:\\GoLearn\\src\\GoConfig", ".tpl", configMap)
}

func generateConfigFile(path string, ext string, maps map[string]string) {
	var absPaths = GetTmpFile(path, ext)
	//var outPutContent []string
	for _, v := range absPaths {
		newFilePath := strings.Replace(v, ext, ".config", 1)
		//创建新文件
		createdfile, err := os.Create(newFilePath)
		if err != nil {
			panic(err)
		}
		//打开文件
		file, err := os.Open(v)
		if err != nil {
			panic(err)
		}
		//写缓存
		w := bufio.NewWriter(createdfile)
		defer file.Close()
		defer createdfile.Close()
		//行读
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			//find the key in every line
			lineContent := scanner.Text()
			begin := strings.Index(lineContent, "{")
			if begin == -1 {
				break
			} else {
				end := strings.Index(lineContent, "}")
				key := lineContent[begin+2 : end]
				lineContent = strings.Replace(lineContent, lineContent[begin:end+1], maps[key], 1)
			}
			n4, _ := w.WriteString(lineContent + "\n")
			fmt.Println(n4)
		}
		w.Flush()

	}
}
