package models

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
	ChangHwa      string `xml:"彰化縣" json:"Changhwa"`
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
