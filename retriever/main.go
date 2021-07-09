package main

import (
	"fmt"
	"learninggo/retriever/mock"
	"learninggo/retriever/real"
	"time"
)

type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string, form map[string]string) string
}

func post(p Poster) {
	p.Post("https://www.baidu.com", map[string]string{
		"name": "张山",
		"age":  "18",
	})
}

//接口组合
//定义一个全局变量
const url = "https://www.baidu.com"

type PostRetriever interface {
	Poster
	Retriever
}

func session(s PostRetriever) string {
	s.Post(url, map[string]string{
		"contents": "another faked baidu.com",
	})
	return s.Get(url)
}

func downLoad(r Retriever) string {
	return r.Get("https://www.baidu.com")
}

func main() {
	var r Retriever
	Retriever := mock.Retriever{"this is a fake baidu.com"}
	r = Retriever
	fmt.Printf("%T %v\n", r, r)
	inspect(r)
	r = &real.Retriever{"Mozilla/5.0", time.Minute}
	fmt.Printf("%T %v\n", r, r)
	//fmt.Println(downLoad(r))
	inspect(r)
	//	type assertion 通过.类型名称获取类型的值
	if mockRetriever, ok := r.(mock.Retriever); ok {
		fmt.Println(mockRetriever.Contents)
	} else {
		fmt.Println("not a mock retriever")
	}

	fmt.Println("try a session")
	fmt.Println(session(Retriever))

}

func inspect(r Retriever) {
	fmt.Println("Inspecting", r)
	fmt.Printf("> %T %v\n", r, r)
	fmt.Print("> Type switch")
	switch v := r.(type) {
	case mock.Retriever:
		fmt.Println("contents:", v.Contents)
	case *real.Retriever:
		fmt.Println("UserAgent:", v.UserAgent)

	}
}
