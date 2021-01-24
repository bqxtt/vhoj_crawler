package main

import "github.com/vhoj_crawler/pkg/crawler/hdu"

func main() {
	//client := &http.Client{}
	//req, _ := http.NewRequest("GET", "http://poj.org/problem?id=1000", nil)
	//resp, err := client.Do(req)
	//if err != nil {
	//	fmt.Println("err")
	//	return
	//}
	//body, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	fmt.Printf("read err, %v", err)
	//	return
	//}
	//fmt.Println(string(body))
	//poj := &poj.POJCrawler{}
	//poj.Crawl("1")
	hdu := &hdu.HDUCrawler{}
	hdu.Crawl("1000")
}
