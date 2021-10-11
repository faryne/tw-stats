package main

import "github.com/faryne/tw-stats/services"

func main() {
	// 產生主資料
	//services.GenerateGeneralData()
	// 產生性別失業率資料
	services.DoGenderUnemployment()
}
