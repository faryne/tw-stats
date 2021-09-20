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
	Name      string `xml:"指標名稱" json:"指標名稱"`
	Unit      string `xml:"單位" json:"單位"`
	Def       string `xml:"定義" json:"定義"`
	Year      int64  `xml:"年別" json:"年別"`
	Total     string `xml:"總計" json:"總計"`
	Taiwan    string `xml:"台灣地區" json:"台灣地區"`
	NewTaipei string `xml:"新北市" json:"新北市"`
	Taipei    string `xml:"臺北市" json:"臺北市"`
	Taoyuan   string `xml:"桃園市" json:"桃園市"`
	Taichung  string `xml:"臺中市" json:"臺中市"`
	Tainan    string `xml:"臺南市" json:"臺南市"`
	Kaohsiung string `xml:"高雄市" json:"高雄市"`
	Ilan      string `xml:"宜蘭縣" json:"宜蘭縣"`
	Hsinchu1  string `xml:"新竹縣" json:"新竹縣"`
	Maioli    string `xml:"苗栗縣" json:"苗栗縣"`
	ChungHwa  string `xml:"彰化縣" json:"彰化縣"`
	Nantou    string `xml:"南投縣" json:"南投縣"`
	Yunlin    string `xml:"雲林縣" json:"雲林縣"`
	ChaYi1    string `xml:"嘉義縣" json:"嘉義縣"`
	Pingtung  string `xml:"屏東縣" json:"屏東縣"`
	Taitung   string `xml:"臺東縣" json:"臺東縣"`
	Hualien   string `xml:"花蓮縣" json:"花蓮縣"`
	Penghu    string `xml:"澎湖縣" json:"澎湖縣"`
	Keelung   string `xml:"基隆市" json:"基隆市"`
	Hsinchu2  string `xml:"新竹市" json:"新竹市"`
	ChaYi2    string `xml:"嘉義市" json:"嘉義市"`
	Kinmen    string `xml:"金門縣" json:"金門縣"`
	LienJiang string `xml:"連江縣" json:"連江縣"`
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
		year := strconv.FormatInt(v.Year, 10)
		subDirName := dirName + "/" + v.Name
		if err := checkFolderAndBuild(subDirName); err != nil {
			fmt.Println("cannot create subfolder: " + err.Error())
			return
		}
		fName := subDirName + "/" + year + ".json"
		var fp *os.File
		fp, err = os.Open(fName)
		// 如果檔案不存在則建立
		if err != nil {
			if os.IsNotExist(err) == false {
				fmt.Println(err)
				return
			}
			fp, err = os.Create(fName)
		}
		// 將產生的內容轉為 json
		output, _ := json.Marshal(v)
		// 寫入內容
		fp.Write(output)
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
