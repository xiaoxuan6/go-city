package command

import (
	"errors"
	"fmt"
	"github.com/gocolly/colly"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
)

const govUrl = "http://www.stats.gov.cn/sj/tjbz/tjyqhdmhcxhfdm/2022/"

func Exec(c *cli.Context) error {

	if err := base(c); err != nil {
		return errors.New(err.Error())
	}

	var num int
	var wg sync.WaitGroup
	var m sync.Mutex
	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		for {
			select {
			case data, ok := <-ch:
				if !ok {
					break
				}

				if len(data) < 1 {
					break
				}

				m.Lock()
				num = num + len(data)
				m.Unlock()
				save(data...)

			case <-time.After(time.Second * 3):
				goto Loop
			}
		}
	Loop:
		wg.Done()
	}(&wg)

	logrus.Info("正在扒取数据中...")
	syncTjj()
	wg.Wait()
	logrus.Info(fmt.Sprintf("数据同步完成, 总共同步 %d 条数据", num))
	return nil
}

func syncTjj() {
	var res []Response

	request(govUrl, "a[href]", func(e *colly.HTMLElement) {
		href := e.Attr("href")
		text := e.Text
		if strings.HasSuffix(href, ".html") {
			idStr := strings.Replace(href, ".html", "", -1)
			id, _ := strconv.Atoi(idStr + strings.Repeat("0", 12-len(idStr)))
			res = append(res, Response{
				ID:   id,
				Name: text,
				Pid:  0,
			})

			go provinceChildren(id, govUrl+href)
		}
	})

	ch <- res
}

// http://www.stats.gov.cn/sj/tjbz/tjyqhdmhcxhfdm/2022/11.html
func provinceChildren(key int, url string) {
	var res []Response
	request(url, ".citytr", func(e *colly.HTMLElement) {
		text := e.Text
		id, nameStr, idStr := check(text)
		res = append(res, Response{
			ID:   id,
			Name: nameStr,
			Pid:  key,
		})

		url = strings.Replace(e.Request.URL.String(), ".html", "", -1) + "/" + strings.TrimRight(idStr, "0") + ".html"
		cityChildren(id, url)
	})
	ch <- res
}

// http://www.stats.gov.cn/sj/tjbz/tjyqhdmhcxhfdm/2022/11/1101.html
func cityChildren(key int, url string) {
	var res []Response
	request(url, ".countytr", func(e *colly.HTMLElement) {
		text := e.Text
		var href string
		e.ForEach("a[href]", func(i int, element *colly.HTMLElement) {
			href = element.Attr("href")
		})

		id, nameStr, _ := check(text)
		res = append(res, Response{
			ID:   id,
			Name: nameStr,
			Pid:  key,
		})

		url = strings.Replace(e.Request.URL.String(), ".html", "", -1)
		index := strings.LastIndex(url, "/")
		countyChildren(id, url[:index]+"/"+href)
	})
	ch <- res
}

// http://www.stats.gov.cn/sj/tjbz/tjyqhdmhcxhfdm/2022/11/01/110101.html
// http://www.stats.gov.cn/sj/tjbz/tjyqhdmhcxhfdm/2022/11/01/110102.html
// http://www.stats.gov.cn/sj/tjbz/tjyqhdmhcxhfdm/2022/11/01/110105.html
func countyChildren(key int, url string) {
	var res []Response
	request(url, ".towntr", func(e *colly.HTMLElement) {
		text := e.Text
		var href string
		e.ForEach("a[href]", func(i int, element *colly.HTMLElement) {
			href = element.Attr("href")
		})

		id, nameStr, _ := check(text)
		res = append(res, Response{
			ID:   id,
			Name: nameStr,
			Pid:  key,
		})

		url = strings.Replace(e.Request.URL.String(), ".html", "", -1)
		index := strings.LastIndex(url, "/")
		townChildren(id, url[:index]+"/"+href)
	})
	ch <- res
}

// http://www.stats.gov.cn/sj/tjbz/tjyqhdmhcxhfdm/2022/11/01/02/110102001.html
// http://www.stats.gov.cn/sj/tjbz/tjyqhdmhcxhfdm/2022/50/02/29/500229002.html
func townChildren(key int, url string) {
	var res []Response
	request(url, ".villagetr", func(e *colly.HTMLElement) {
		id := e.DOM.Find("td:first-child").First().Text()
		idStr, _ := strconv.Atoi(id)
		nameStr := e.DOM.Find("td:nth-child(3)").First().Text()
		res = append(res, Response{
			ID:   idStr,
			Name: nameStr,
			Pid:  key,
		})
	})
	ch <- res
}

func request(url, selector string, fn func(e *colly.HTMLElement)) {
	c := colly.NewCollector()
	c.OnHTML(selector, fn)
	_ = c.Visit(url)
}

func check(text string) (int, string, string) {
	idStr := regexp.MustCompile("[0-9]{12}").FindString(text)
	nameStr := strings.TrimSpace(strings.Replace(text, idStr, "", -1))
	id, _ := strconv.Atoi(idStr)
	return id, nameStr, idStr
}
