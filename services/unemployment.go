package services

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/faryne/tw-stats/constants"
	"github.com/faryne/tw-stats/models"
	"io/fs"
	"net/http"
	"os"
	"reflect"
	"strings"
)

// 資料來源：https://data.gov.tw/dataset/33442
// 性別失業率指標
var genderUnemploymentByYearFiles = map[int64]string{
	2015: "https://www.dgbas.gov.tw/public/data/open/Cen/Mp04037.xml",
	2016: "https://www.dgbas.gov.tw/public/data/open/Cen/Mp04037A105.xml",
	2017: "https://www.dgbas.gov.tw/public/data/open/Cen/Mp04037A106.xml",
	2018: "https://www.dgbas.gov.tw/public/data/open/Cen/Mp04037A107.xml",
	2019: "https://www.dgbas.gov.tw/public/data/open/Cen/Mp04037A108.xml",
	2020: "https://www.dgbas.gov.tw/public/data/open/Cen/Mp04037A109.xml",
}

type genderUnemploymentRoot struct {
	Root []genderUnemployment `xml:"人力資源調查重要指標"`
}
type genderUnemployment struct {
	OldArea string  `xml:"按地區別分_District_or_region"`
	Area    string  `xml:"地區別_District_or_region"`
	Male    float64 `xml:"失業率_男_Unemployment_rate_Male"`
	Female  float64 `xml:"失業率_女_Unemployment_rate_Female"`
}

var maleByYears = make(map[string]models.Data)
var femaleByYears = make(map[string]models.Data)

func DoGenderUnemployment() {
	// 檢查目錄
	maleDir := constants.RootDirName + "/" + "年度失業率-男"
	femaleDir := constants.RootDirName + "/" + "年度失業率-女"
	checkFolderAndBuild(maleDir)
	checkFolderAndBuild(femaleDir)
	for k, url := range genderUnemploymentByYearFiles {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println("cannot get file of %d", k)
			return
		}
		defer resp.Body.Close()
		//解析 xml
		reader := xml.NewDecoder(resp.Body)
		var stat genderUnemploymentRoot
		reader.Decode(&stat)

		// 賦值
		var result1 = models.Data{
			Name:  "年度失業率-男",
			Unit:  "%",
			Def:   "該年度男性失業率",
			Year:  k,
			Total: "---",
		}
		var result2 = models.Data{
			Name:  "年度失業率-女",
			Unit:  "%",
			Def:   "該年度女性失業率",
			Year:  k,
			Total: "---",
		}
		reflectField1 := reflect.ValueOf(result1)
		reflectField2 := reflect.ValueOf(result2)
		for _, v := range stat.Root {
			var area = v.OldArea
			if area == "" {
				area = v.Area
			}
			for i := 0; i < reflectField1.Type().NumField(); i++ {
				var city = reflectField1.Type().Field(i).Tag.Get("xml")

				if strings.Contains(area, city) {
					reflect.ValueOf(&result1).Elem().Field(i).SetString(fmt.Sprintf("%.2f", v.Male))
				}
			}
			for i := 0; i < reflectField2.Type().NumField(); i++ {
				var city = reflectField2.Type().Field(i).Tag.Get("xml")
				if strings.Contains(area, city) {
					reflect.ValueOf(&result2).Elem().Field(i).SetString(fmt.Sprintf("%.2f", v.Male))
				}
			}
		}

		maleByYears[fmt.Sprintf("%d", k)] = result1
		femaleByYears[fmt.Sprintf("%d", k)] = result2
		//寫入各年度檔案
		var maleFile = maleDir + fmt.Sprintf("/%d", k) + ".json"
		var fpMale *os.File
		if fpMale, err = os.OpenFile(maleFile, os.O_CREATE|os.O_WRONLY, fs.ModePerm); err != nil {
			fmt.Println(err)
			return
		}
		outputMale, _ := json.MarshalIndent(result1, "", "    ")
		if _, err := fpMale.Write(outputMale); err != nil {
			fmt.Println(err)
		}
		defer fpMale.Close()

		var femaleFile = femaleDir + fmt.Sprintf("/%d", k) + ".json"
		var fpFemale *os.File
		if fpFemale, err = os.OpenFile(femaleFile, os.O_CREATE|os.O_WRONLY, fs.ModePerm); err != nil {
			fmt.Println(err)
			return
		}
		outputFemale, _ := json.MarshalIndent(result2, "", "    ")
		if _, err := fpMale.Write(outputFemale); err != nil {
			fmt.Println(err)
		}
		defer fpFemale.Close()
	}
	// 寫入目錄檔
	var maleIndex = maleDir + "/index.json"
	fpMaleIndex, err := os.OpenFile(maleIndex, os.O_CREATE|os.O_WRONLY, fs.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}
	outputMale, _ := json.MarshalIndent(maleByYears, "", "    ")
	if _, err := fpMaleIndex.Write(outputMale); err != nil {
		fmt.Println(err)
	}
	fpMaleIndex.Close()

	var femaleIndex = femaleDir + "index.json"
	fpFemaleIndex, err := os.OpenFile(femaleIndex, os.O_CREATE|os.O_WRONLY, fs.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}
	outputFemale, _ := json.MarshalIndent(femaleByYears, "", "    ")
	if _, err := fpFemaleIndex.Write(outputFemale); err != nil {
		fmt.Println(err)
	}
	fpFemaleIndex.Close()
}
