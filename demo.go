package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func respParseHandler(result string){

	if len(result) < 1{
		return
	}

	switch result[0:1]{
		case "+":{
			//状态回复
			dealStatusReply(result)
		}
		case "-":
			//错误回复
			fmt.Println(strings.Replace(result[1:],"\r\n","",-1))

		case ":":
			//整数回复
			val := strings.Replace(result[1:],"\r\n","",-1)
			i, err := strconv.Atoi(val)
			if err != nil{
				fmt.Println(err)
				return
			}
			fmt.Println(i)

		case "$":
			//批量回复
			fmt.Println(strings.Replace(result[1:],"\r\n","",-1))
			dealBulkReply(result)

	}
}

// 处理主体响应
func dealBulkReply(reply string) (interface{}, error) {

	statusByte := reply[1:]

	// 获取响应文本第一行标示的响应字符串长度
	pos := 0

	for _, v := range statusByte {
		if v == '\r' {
			break
		}
		pos++
	}

	strlen, err := strconv.Atoi(string(statusByte[:pos]))
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	if strlen == -1 {
		return "nil", nil
	}
	nextLinePost := 1
	for _, v := range statusByte {
		if v == '\n' {
			break
		}
		nextLinePost++
	}

	result := string(statusByte[nextLinePost:nextLinePost+strlen])

	fmt.Println(result)

	return result, nil
}

// 处理状态响应
func dealStatusReply(reply string) (interface{}, error) {
	statusByte := reply[1:]

	pos := 0
	for _, v := range statusByte {
		if v == '\r' {
			break
		}
		pos++
	}
	status := statusByte[:pos]
	fmt.Println(status)

	return string(status), nil
}

func main() {
	respParseHandler("+OK\r\n")
	respParseHandler(":1000\r\n")
}