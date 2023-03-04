package command

import (
	"fmt"
	"github.com/guonaihong/gout"
	"github.com/schollz/progressbar/v3"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
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

func Run(c *cli.Context) error {

	driver := c.String("driver")
	switch driver {
	case "sqlite":
		InitSqlite(c.String("db"))
	case "memory":
		InitSqlite(":memory:")
	case "database":
		InitSql(
			c.String("host"),
			c.String("port"),
			c.String("username"),
			c.String("password"),
			c.String("dbname"),
			c.String("table"),
		)
	default:
		logrus.Error("无效的 driver")
	}

	syncAll()

	logrus.Info(fmt.Sprintf("数据同步完成，其中省级行政区：%d，城市：%d，区县：%d，乡镇街道：%d", count.ProvinceNum, count.CityNum, count.AreaNum, count.StreetNum))

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

func syncAll() {
	provinceCount := len(provinces)
	count.ProvinceNum = provinceCount
	bar := progressbar.Default(int64(provinceCount))
	for key, val := range provinces {
		save(Response{ID: key, Name: val, Pid: 0})

		city(key)
		_ = bar.Add(1)
	}
}

var cityResponse []Response

func city(id int) {
	result := syncById(id)
	cityResponse = append(cityResponse, result...)
	count.CityNum = count.CityNum + len(cityResponse)

	if len(result) > 0 {
		save(result...)
	}

	for _, val := range result {
		area(val.ID)
	}
}

var areaResponse []Response

func area(id int) {
	result := syncById(id)
	areaResponse = append(areaResponse, result...)
	count.AreaNum = count.AreaNum + len(areaResponse)

	if len(result) > 0 {
		save(result...)
	}

	for _, val := range result {
		street(val.ID)
	}
}

var streetResponse []Response

func street(id int) {
	result := syncById(id)
	streetResponse = append(streetResponse, result...)
	count.StreetNum = count.StreetNum + len(streetResponse)

	if len(result) > 0 {
		save(result...)
	}
}
