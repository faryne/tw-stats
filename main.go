package main

import (
	"github.com/faryne/tw-stats/services"
	"github.com/faryne/tw-stats/services/helper/generateIndex"
)

func main() {
	// 產生主資料
	services.GenerateGeneralData()
	// 產生性別失業率資料
	services.DoGenderUnemployment()
	// 產生總目錄檔
	generateIndex.Generate()
}
