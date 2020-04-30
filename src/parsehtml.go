package src

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

type Response struct {
	Status int `json:"status"`
	Message string `json:"message"`
	Data interface{} `json:"data"`
}

type SubmitData struct {
	Url string `json:"url"`
}

var hzRegexp = regexp.MustCompile("^[\u4e00-\u9fa5]$")

//func StrFilterNonChinese(src *string) {
//	strn := ""
//	for _, c := range *src {
//		if hzRegexp.MatchString(string(c)) {
//			strn += string(c)
//		}
//	}
//	*src = strn
//}

//func ParseHtml(context *gin.Context){
//	//生成client 参数为默认
//	client := &http.Client{}
//	url,_ := context.GetQuery("url")
//	request, err := http.NewRequest("GET", url, nil)
//	if err != nil {
//		result := Response{
//			Status: 400,
//			Message: "解析失败啦",
//			Data: "解析失败啦,试试其他网页地址吧",
//		}
//		context.JSON(result.Status,&result)
//	}
//	//处理返回结果
//	response, _ := client.Do(request)
//	//将结果定位到标准输出 也可以直接打印出来 或者定位到其他地方进行相应的处理
//	//stdout := os.Stdout
//	//_, err = io.Copy(stdout, response.Body)
//	if response.StatusCode == 200 {
//		body, _ := ioutil.ReadAll(response.Body)
//		res := string(body)
//		StrFilterNonChinese(&res)
//		result := Response{
//			Status: response.StatusCode,
//			Message: "解析成功啦",
//			Data: res,
//		}
//		context.JSON(result.Status,&result)
//	}
//}

func trimHtml(src string) string {
	//将HTML标签全转换成小写
	re, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllStringFunc(src, strings.ToLower)
	//去除STYLE
	re, _ = regexp.Compile("\\<style[\\S\\s]+?\\</style\\>")
	src = re.ReplaceAllString(src, "")
	//去除SCRIPT
	re, _ = regexp.Compile("\\<script[\\S\\s]+?\\</script\\>")
	src = re.ReplaceAllString(src, "")
	//去除所有尖括号内的HTML代码，并换成换行符
	re, _ = regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllString(src, "\n")
	//去除连续的换行符
	re, _ = regexp.Compile("\\s{2,}")
	src = re.ReplaceAllString(src, " ")
	return strings.TrimSpace(src)
}

func ParseHtml(context *gin.Context){
	//生成client 参数为默认
	client := &http.Client{}
	result := SubmitData{}
	err := context.BindJSON(&result)
	if err != nil{
		panic(err)
	}else {
		//context.JSON(200,result)
		fmt.Println("result.url:",result.Url)
	}

	request, err := http.NewRequest("GET", result.Url, nil)
	if err != nil {
		result := Response{
			Status: 400,
			Message: "解析失败啦",
			Data: "解析失败啦,试试其他网页地址吧",
		}
		context.JSON(result.Status,&result)
	}
	////处理返回结果
	response, _ := client.Do(request)
	//将结果定位到标准输出 也可以直接打印出来 或者定位到其他地方进行相应的处理
	//stdout := os.Stdout
	//_, err = io.Copy(stdout, response.Body)
	if response.StatusCode == 200 {
		body, _ := ioutil.ReadAll(response.Body)
		res := trimHtml(string(body))
		result := Response{
			Status: response.StatusCode,
			Message: "解析成功啦",
			Data: res,
		}
		context.JSON(result.Status,&result)
	}
}
func Index(context *gin.Context){
	context.HTML(http.StatusOK, "index.html",gin.H{
		"title": "parseHtml",
	})
}