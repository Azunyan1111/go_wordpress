package go_wordpress

import (
	"net/http"
	"time"
	"log"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"github.com/Azunyan1111/go_wordpress/structs"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	"net/url"
	"os"
)

var (
	WORDPRESS_API_BASE_URL  string
	WORDPRESS_USER  string
	WORDPRESS_PASSWORD  string
	WORDPRESS_DB_URL string
	WORDPRESS_DB_USER string
	WORDPRESS_DB_PASS string
	WORDPRESS_DB_NAME string
)

var Db *gorm.DB

func main() {
	var (
		ApiBaseUrl  = os.Getenv("ApiBaseUrl")
		User  = os.Getenv("User")
		Pass  = os.Getenv("Pass")
		DbUrl = os.Getenv("DbUrl")
		DbUser = os.Getenv("DbUser")
		DbPass = os.Getenv("DbPass")
		DbName = os.Getenv("DbName")
	)
	Fast(ApiBaseUrl,User,Pass,DbUrl,DbUser,DbPass,DbName)
}

func Fast(baseURL string,wpUser string,wpPass string,dbURL string,dbUser string,dbPass string,dbName string){
	WORDPRESS_API_BASE_URL  = baseURL
	WORDPRESS_USER = wpUser
	WORDPRESS_PASSWORD = wpPass
	WORDPRESS_DB_URL = dbURL
	WORDPRESS_DB_USER = dbUser
	WORDPRESS_DB_PASS = dbPass
	WORDPRESS_DB_NAME = dbName

	db,err := gorm.Open("mysql",WORDPRESS_DB_USER+":"+WORDPRESS_DB_PASS+"@"+"tcp("+WORDPRESS_DB_URL+")/"+WORDPRESS_DB_NAME)
	if err != nil{
		log.Println("Error:db conntection found")
		panic(err)
	}
	Db = db
}

func WpPost(title string,content string,times time.Time,categories []string, reTry int) error{
	// リトライ処理
	if reTry == 0{
		return fmt.Errorf("Error: reTry end",nil)
	}

	// 送信するデータ用意
	input, err := json.Marshal(structs.Post{Title: title, Content: content,Excerpt:"", DataGmt: times.Format(time.RFC3339), Status: "publish",Categories:CategoriesToInt(categories)})
	if err != nil{
		log.Println(err)
		return WpPost(title,content,times,categories,reTry - 1)
	}

	// クライアント用意
	req, err := http.NewRequest(http.MethodPost, WORDPRESS_API_BASE_URL+"/posts/", bytes.NewBuffer(input))
	if err != nil {
		log.Println(err)
		return WpPost(title,content,times,categories,reTry - 1)
	}
	req.SetBasicAuth(WORDPRESS_USER, WORDPRESS_PASSWORD)
	req.Header.Add("Content-Type", "application/json")

	// 実行
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
		return WpPost(title,content,times,categories,reTry - 1)
	}
	s,err := ioutil.ReadAll(resp.Body)
	if err != nil{
		log.Println(err)
		return fmt.Errorf("Error: ioutil Read error",nil)
	}
	log.Println(string(s))
	return nil
}


func CategoriesToInt(category []string)[]int{
	ints := []int{}
	for _,c := range category{
		cate := SearchCategory(c)
		if cate.Id == 0{
			AddCategories(c)
			cate = SearchCategory(c)
		}
		ints = append(ints, cate.Id)
	}
	return ints
}

func AddCategories(s string){
	cate := structs.CateDb{}
	cate.Name = s
	cate.Slug = url.QueryEscape(s)
	Db.Create(&cate)

	Db.Find(&cate,"name = ?",cate.Name)

	tax := structs.CateDbTaxonomy{}
	tax.Taxonomy = "category"
	tax.TermId = cate.Id
	Db.Create(&tax)
}

func SearchCategory(s string)structs.CateDb{
	cate := structs.CateDb{}
	Db.Find(&cate,"name = ?",s)
	return cate
}

func PostExist(s string) bool{
	p := structs.PostDb{}
	Db.Find(&p,"post_title = ?",s)
	if p.Id == 0{
		return false
	}else{
		return true
	}
}

func GetCount() int{
	c := structs.Count{}
	c.Id = 1
	Db.Find(&c)
	return c.Count
}

func SetCount(i int){
	c := structs.Count{}
	c.Id = 1
	Db.Find(&c)
	c.Count = i
	Db.Save(&c)
}