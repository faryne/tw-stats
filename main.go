package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/fs"
	"net/http"
	"os"
	"strconv"
)

type Root struct {
	Data []Data `xml:"Data"`
}
type Data struct {
	Name          string `xml:"指標名稱" json:"Name"`
	Unit          string `xml:"單位" json:"Unit"`
	Def           string `xml:"定義" json:"Explain"`
	Year          int64  `xml:"年別" json:"ByYear"`
	Total         string `xml:"總計" json:"Total"`
	Taiwan        string `xml:"臺灣地區" json:"Taiwan"`
	NewTaipei     string `xml:"新北市" json:"NewTaipei"`
	Taipei        string `xml:"臺北市" json:"Taipei"`
	Taoyuan       string `xml:"桃園市" json:"Taoyuan"`
	Taichung      string `xml:"臺中市" json:"Taichung"`
	Tainan        string `xml:"臺南市" json:"Tainan"`
	Kaohsiung     string `xml:"高雄市" json:"Kaohsiung"`
	Ilan          string `xml:"宜蘭縣" json:"Ilan"`
	HsinchuCounty string `xml:"新竹縣" json:"HsinchuCounty"`
	Miaoli        string `xml:"苗栗縣" json:"Miaoli"`
	ChangHwa      string `xml:"彰化縣" json:"ChangHwa"`
	Nantou        string `xml:"南投縣" json:"Nantou"`
	Yunlin        string `xml:"雲林縣" json:"Yunlin"`
	ChiaYiCounty  string `xml:"嘉義縣" json:"ChiaYiCounty"`
	Pingtung      string `xml:"屏東縣" json:"Pingtung"`
	Taitung       string `xml:"臺東縣" json:"Taitung"`
	Hualien       string `xml:"花蓮縣" json:"Hualien"`
	Penghu        string `xml:"澎湖縣" json:"Penghu"`
	Keelung       string `xml:"基隆市" json:"Keelung"`
	HsinchuCity   string `xml:"新竹市" json:"HsinchuCity"`
	ChiaYiCity    string `xml:"嘉義市" json:"ChiaYiCity"`
	Kinmen        string `xml:"金門縣" json:"Kinmen"`
	Matsu         string `xml:"連江縣" json:"Matsu"`
}

func main() {
	data, err := readAndParseXML()
	if err != nil {
		fmt.Println(err)
		return
	}
	// 建立目錄
	dirName := "./docs"
	if err := checkFolderAndBuild(dirName); err != nil {
		fmt.Println("cannot create main folder: " + err.Error())
		return
	}

	for _, v := range data {
		// 檢查檔案是否存在
		year := strconv.FormatInt(v.Year+1911, 10)
		subDirName := dirName + "/" + v.Name
		if err := checkFolderAndBuild(subDirName); err != nil {
			fmt.Println("cannot create subfolder: " + err.Error())
			return
		}
		fName := subDirName + "/" + year + ".json"
		// 如果檔案不存在則建立
		var fp *os.File
		if fp, err = os.OpenFile(fName, os.O_CREATE|os.O_WRONLY, fs.ModePerm); err != nil {
			fmt.Println(err)
			return
		}
		// 將產生的內容轉為 json
		output, _ := json.Marshal(v)
		// 寫入內容
		if _, err := fp.Write(output); err != nil {
			fmt.Println(err)
		}
		// 關閉檔案
		fp.Close()
	}
}

func readAndParseXML() ([]Data, error) {
	url := "https://www.dgbas.gov.tw/public/data/open/LocalStat/%E7%B8%A3%E5%B8%82%E6%8C%87%E6%A8%99.xml"
	// 抓取內容
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	doc := xml.NewDecoder(resp.Body)
	var v = Root{}
	e := doc.Decode(&v)
	if e != nil {
		return nil, e
	}
	return v.Data, nil
}

func checkFolderAndBuild(folder string) error {
	var err error
	h, err := os.Open(folder)
	h.Close()
	if err != nil {
		if os.IsExist(err) == false {
			return os.Mkdir(folder, fs.ModePerm)
		}
	}
	return nil
}
