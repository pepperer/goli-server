package Script

import (
	"fmt"
	"github.com/gocolly/colly"
	"io/ioutil"
	"log"
	"os/exec"
	path2 "path"
)

func StartFindFile(path string) {
	dir := path2.Dir(path)
	fmt.Println(dir)

	bytes, e := ioutil.ReadFile(path)
	if e != nil {
		panic(e)
	}

	//fmt.Printf("%s", bytes)
	fmt.Println(string(bytes))

}

func StartWatch() {
	c := colly.NewCollector()


	err := c.Post("http://example.com/login", map[string]string{"username": "admin", "password": "admin"})
	if err != nil {
		log.Fatal(err)
	}

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		// Print link
		fmt.Printf("Link found: %q -> %s\n", e.Text, link)
		// Visit link found on page
		// Only those links are visited which are in AllowedDomains
		c.Visit(e.Request.AbsoluteURL(link))
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.OnResponse(func(response *colly.Response) {
		//fmt.Println("Response", string(response.Body))
	})

	// Set error handler
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	c.Visit("http://www.baidu.com")
}

func StartCmd() {
	shellPath := "/home/xx/test.sh"
	command := exec.Command(shellPath) //初始化Cmd
	err := command.Start()//运行脚本
	if nil != err {
		fmt.Println(err)
	}
	fmt.Println("Process PID:", command.Process.Pid)
	err = command.Wait() //等待执行完成
	if nil != err {
		fmt.Println(err)
	}
	fmt.Println("ProcessState PID:", command.ProcessState.Pid())
}
