package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type Appinfo struct {
	Appid   int    `json:"appid"`
	Pkgname string `json:"pkg_name"`
}

type Appinfos struct {
	AppInfo []Appinfo `json:"appinfos"`
}

func (a Appinfos) String() string {
	var res string
	for _, v := range a.AppInfo {
		if v.Pkgname != "" {
			res += fmt.Sprintf("%d\n%s\n", v.Appid, v.Pkgname)
		}
	}
	fmt.Print(res)
	return res
}

func GetAppinfo(appid string, ch chan string) {
	urlStr := "http://ipush.baidu.com/index.php/normalTools/apikeyQuery/get_mongo_data/"
	urlStr += appid
	client := &http.Client{}
	resp, err := client.Get(urlStr)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	if index := strings.LastIndex(string(body), "</div>"); index != -1 {
		body = body[index+1:]
	}
	var appinfos Appinfos
	err = json.Unmarshal(body, &(appinfos.AppInfo))
	if err != nil {
		fmt.Println("error:", err)
	}
	ch <- appinfos.String()
}

func main() {
	infile, err := os.Open("./appid")
	if err != nil {
		fmt.Println("error: ", err)
	}
	outfile, err := os.OpenFile("./out1", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println("error: ", err)
	}
	defer func() {
		infile.Close()
		outfile.Close()
	}()

	reader := bufio.NewReader(infile)
	ch := make(chan string, 100)
	allNum := 0
	sucNum := 0
Loop:
	for {
		readString, err := reader.ReadString('\n')
		switch err {
		case io.EOF:
			break Loop
		case nil:
			allNum++
			appidString := readString[0 : len(readString)-1]
			go GetAppinfo(appidString, ch)
		default:
			fmt.Println("error: ", err)
		}
	}
	for i := 0; i < allNum; i++ {
		appinfo := <-ch
		if appinfo != "" {
			sucNum++
		}
		io.WriteString(outfile, appinfo)
	}
	fmt.Println("success: ", sucNum, ",all: ", allNum)
}
