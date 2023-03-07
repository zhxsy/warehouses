package utils

import (
	"bufio"
	"io"
	"os"
)

// 按行读取文件内容
func ReadFile(fileName string, contentChan chan string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		str, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		} else if err != nil {
			if err != nil {
				return err
			}
		}
		contentChan <- string(str)
	}
	close(contentChan)
	return nil
}

// 写入文件
func WriteFile(fileName, content string) {
	file, _ := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_TRUNC|os.O_SYNC, 0644)
	file.Write([]byte(content))
	file.Close()
}
