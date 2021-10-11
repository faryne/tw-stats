package services

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/faryne/tw-stats/constants"
	"github.com/faryne/tw-stats/models"
	"github.com/faryne/tw-stats/services/helper/generateIndex"
	"io/fs"
	"net/http"
	"os"
	"strconv"
)

func GenerateGeneralData() {
	data, err := readAndParseXML()
	if err != nil {
		fmt.Println(err)
		return
	}
	// 建立目錄
	if err := checkFolderAndBuild(constants.RootDirName); err != nil {
		fmt.Println("cannot create main folder: " + err.Error())
		return
	}

	var index = make(map[string]map[int64]models.Data)
	var idx = 0
	for _, v := range data {
		// 檢查檔案是否存在
		year := strconv.FormatInt(v.Year+1911, 10)
		subDirName := constants.RootDirName + "/" + v.Name
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
		if index[v.Name] == nil {
			idx++
			index[v.Name] = make(map[int64]models.Data)
			generateIndex.Add(v.Name)
		}
		index[v.Name][(v.Year + 1911)] = v
		// 將產生的內容轉為 json
		output, _ := json.MarshalIndent(v, "", "    ")
		// 寫入內容
		if _, err := fp.Write(output); err != nil {
			fmt.Println(err)
		}
		// 關閉檔案
		fp.Close()
	}
	// 寫入總目錄檔
	generateIndex.Generate()
	// 產生目錄檔
	for k, v := range index {
		indexFileName := constants.RootDirName + "/" + k + "/index.json"
		var fpIndex *os.File
		if fpIndex, err = os.OpenFile(indexFileName, os.O_CREATE|os.O_WRONLY, fs.ModePerm); err != nil {
			fmt.Println(err)
			return
		}
		outputIndex, _ := json.MarshalIndent(v, "", "    ")
		if _, err := fpIndex.Write(outputIndex); err != nil {
			fmt.Println(err)
		}
		fpIndex.Close()
	}
}

func readAndParseXML() ([]models.Data, error) {
	url := "https://www.dgbas.gov.tw/public/data/open/LocalStat/%E7%B8%A3%E5%B8%82%E6%8C%87%E6%A8%99.xml"
	// 抓取內容
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	doc := xml.NewDecoder(resp.Body)
	var v = models.Root{}
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
