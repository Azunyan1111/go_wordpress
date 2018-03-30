package main

import (
	"net/http"
	"time"
	"log"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"github.com/Azunyan1111/go_wordpress/structs"
)

const(
	WORDPRESS_API_BASE_URL  = "https://jkh.jp/wp-json/wp/v2"
	WORDPRESS_USER  = "postman"
	WORDPRESS_PASSWORD  = "KpUeX51wW1Jd^cRA3&Ej*Ycq"
)


func main() {

}

func WpPost(title string,content string,times time.Time,categories []string, reTry int) error{
	// リトライ処理
	if reTry == 0{
		return new(error)
	}

	// 送信するデータ用意
	input, err := json.Marshal(structs.Post{Title: title, Content: content, DataGmt: times.Format(time.RFC3339), Status: "publish"})//,Categories:CategoriesToInt(categories)})
	if err != nil{
		log.Println(err)
		return WpPost(title,content,times,categories,reTry - 1)
	}
	// クライアント用意
	req, err := http.NewRequest(http.MethodPost, WORDPRESS_API_BASE_URL+"/posts", bytes.NewBuffer(input))
	if err != nil{
		log.Println(err)
		return err
	}
	req.SetBasicAuth(WORDPRESS_USER, WORDPRESS_PASSWORD)
	req.Header.Add("Content-Type", "Application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
		return err
	}
	s,err := ioutil.ReadAll(resp.Body)
	if err != nil{
		log.Println(err)
		return err
	}
	log.Println(string(s))

}