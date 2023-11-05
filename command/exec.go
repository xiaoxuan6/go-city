package command

import (
	"errors"
	"fmt"
	"github.com/guonaihong/gout"
	"github.com/schollz/progressbar/v3"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"sync"
	"time"
)

var provinces = map[int]string{
	1:     "北京",
	2:     "上海",
	3:     "天津",
	4:     "重庆",
	5:     "河北",
	6:     "山西",
	7:     "河南",
	8:     "辽宁",
	9:     "吉林",
	10:    "黑龙江",
	11:    "内蒙古",
	12:    "江苏",
	13:    "山东",
	14:    "安徽",
	15:    "浙江",
	16:    "福建",
	17:    "湖北",
	18:    "湖南",
	19:    "广东",
	20:    "广西",
	21:    "江西",
	22:    "四川",
	23:    "海南",
	24:    "贵州",
	25:    "云南",
	26:    "西藏",
	27:    "陕西",
	28:    "甘肃",
	29:    "青海",
	30:    "宁夏",
	31:    "新疆",
	32:    "台湾",
	84:    "钓鱼岛",
	52993: "港澳",
}

var url = "https://fts.jd.com/area/get?fid=%d"

var count struct {
	ProvinceNum int `json:"province_num"`
	CityNum     int `json:"city_num"`
	AreaNum     int `json:"area_num"`
	StreetNum   int `json:"street_num"`
}

var bar *progressbar.ProgressBar
var provinceCount = len(provinces)

func Run(c *cli.Context) error {
	startTime := time.Now()
	bar = progressbar.Default(int64(provinceCount))

	if err := base(c); err != nil {
		return errors.New(err.Error())
	}

	var wg sync.WaitGroup
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

				save(data...)
			case <-time.After(time.Second * 3):
				goto Loop
			}
		}
	Loop:
		wg.Done()
	}(&wg)

	syncAll()
	wg.Wait()
	_ = bar.Add(1)

	logrus.Info(fmt.Sprintf("数据同步完成，耗时(分钟)：%f", time.Since(startTime).Minutes()))
	logrus.Info(fmt.Sprintf("其中省级行政区：%d，城市：%d，区县：%d，乡镇街道：%d", count.ProvinceNum, count.CityNum, count.AreaNum, count.StreetNum))

	return nil
}

func syncById(pid int) (res []Response) {
	err := gout.GET(fmt.Sprintf(url, pid)).
		BindJSON(&res).
		Do()

	if err != nil {
		logrus.Error(fmt.Sprintf("Pid %d 获取数据失败, 错误：%s", pid, err.Error()))
	}

	for key, _ := range res {
		res[key].Pid = pid
	}

	return
}

var ch = make(chan []Response, 100)

func syncAll() {
	count.ProvinceNum = provinceCount

	var i = 1
	var data []Response
	for key, val := range provinces {
		data = append(data, Response{ID: key, Name: val, Pid: 0})

		city(key)

		if i < provinceCount {
			_ = bar.Add(1)
		}
		i = i + 1
	}
	ch <- data
}

func city(id int) {
	result := syncById(id)
	count.CityNum = count.CityNum + len(result)

	ch <- result

	for _, val := range result {
		// 这里最适合使用协成提高速度
		go area(val.ID)
	}
}

func area(id int) {
	result := syncById(id)
	count.AreaNum = count.AreaNum + len(result)

	ch <- result

	for _, val := range result {
		street(val.ID)
	}
}

func street(id int) {
	result := syncById(id)
	count.StreetNum = count.StreetNum + len(result)

	ch <- result
}
