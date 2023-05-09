package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httputil"
	_ "net/http/pprof"
)

func main() {
	//resp, err := http.Get("https://www.baidu.com")
	//if err != nil {
	//	panic(err)
	//}
	//defer resp.Body.Close()
	////	resp直接打印输出为内存地址、需要 DumpResponse 转储响应得到响应正文 如果使用日志格式输出就为二进制，需要进行格式化
	//s, err := httputil.DumpResponse(resp, true)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("%s\n", s)

	//基于web方式注册路由
	http.HandleFunc("/GetBody", HttpBody)
	http.HandleFunc("/Form", HttpForm)
	http.HandleFunc("/Url", HttpUrl)
	http.ListenAndServe(":8080", nil)

	//	模拟移动端发起请求
	//	1.创建一个http请求 如果是get请求body一般为空，如果post则为请求参数
	request, err := http.NewRequest(http.MethodGet, "https://www.imooc.com", nil)
	if err != nil {
		panic(err)
	}
	//	2.模拟移动端则需要设置移动端的ua
	request.Header.Add("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 13_2_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0.3 Mobile/15E148 Safari/604.1")
	// 如果不使用DefaultClient也可以自己声明一个
	/*
		Transport 多使用在代理服务器上
		CheckRedirect 重定向
		CookieJar cookie 模拟登录时可以使用
	*/

	client := http.Client{CheckRedirect: func(req *http.Request, via []*http.Request) error {
		//	匿名函数中 第一个为目标路径、第二个切片用于存放重定向的路径 即第二种参数路径跳转第一个参数路径
		//	如果返回nil执行重定向
		fmt.Println("Redirect", req)
		return nil
	}}
	response, err := client.Do(request)
	//	3.设置客户端http请求
	if err != nil {
		panic(response)
	}
	//	4.同样的需要进行 转储响应 如果是post这使用 ioutil.ReadAll()
	s, err := httputil.DumpResponse(response, true)
	fmt.Printf("%s\n", s)
}

func HttpBody(w http.ResponseWriter, r *http.Request) {
	//body := r.Body
	//fmt.Printf("读取请求body内容,只能读取一次,内容为%s",body)
	//如果想重复读取body的内容。则可以调用http库的getbody.底层为匿名函数。默认为空。需要手动赋值
	if r.GetBody == nil {
		fmt.Println("body is null")
	} else {
		fmt.Printf("body is %T", r.GetBody)
	}

}

func HttpUrl(w http.ResponseWriter, r *http.Request) {
	//获取request的请求参数
	//将数据json序列化
	data, err := json.Marshal(r.URL)
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, string(data))
}

func HttpForm(w http.ResponseWriter, r *http.Request) {
	//格式化请求数据为表单.不建议一会表单，一会json。最好都使用json
	// 直接格式化返回的数据为空
	fmt.Fprintf(w, "数据1=%v\n", r.Form)
	// 需要先调用ParseForm
	err := r.ParseForm()
	if err != nil {
		panic(err)
	}
	//此时在调用返回的数据就不给空
	fmt.Fprintf(w, "数据2=%v\n", r.Form)
}
