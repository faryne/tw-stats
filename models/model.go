package models

type Root struct {
	Data []Data `xml:"Data"`
}
type Data struct {
	Name          string `xml:"指標名稱" json:"Name"`
	Unit          string `xml:"單位" json:"Unit"`
	Def           string `xml:"定義" json:"Explain"`
	Year          int64  `xml:"年別" json:"ByYear"`
	Total         string `xml:"總計ItemValue" json:"Total"`
	Taiwan        string `xml:"臺灣地區ItemValue" json:"Taiwan"`
	NewTaipei     string `xml:"新北市ItemValue" json:"NewTaipei"`
	Taipei        string `xml:"臺北市ItemValue" json:"Taipei"`
	Taoyuan       string `xml:"桃園市ItemValue" json:"Taoyuan"`
	Taichung      string `xml:"臺中市ItemValue" json:"Taichung"`
	Tainan        string `xml:"臺南市ItemValue" json:"Tainan"`
	Kaohsiung     string `xml:"高雄市ItemValue" json:"Kaohsiung"`
	Ilan          string `xml:"宜蘭縣ItemValue" json:"Ilan"`
	HsinchuCounty string `xml:"新竹縣ItemValue" json:"HsinchuCounty"`
	Miaoli        string `xml:"苗栗縣ItemValue" json:"Miaoli"`
	ChangHwa      string `xml:"彰化縣ItemValue" json:"Changhwa"`
	Nantou        string `xml:"南投縣ItemValue" json:"Nantou"`
	Yunlin        string `xml:"雲林縣ItemValue" json:"Yunlin"`
	ChiaYiCounty  string `xml:"嘉義縣ItemValue" json:"ChiaYiCounty"`
	Pingtung      string `xml:"屏東縣ItemValue" json:"Pingtung"`
	Taitung       string `xml:"臺東縣ItemValue" json:"Taitung"`
	Hualien       string `xml:"花蓮縣ItemValue" json:"Hualien"`
	Penghu        string `xml:"澎湖縣ItemValue" json:"Penghu"`
	Keelung       string `xml:"基隆市ItemValue" json:"Keelung"`
	HsinchuCity   string `xml:"新竹市ItemValue" json:"HsinchuCity"`
	ChiaYiCity    string `xml:"嘉義市ItemValue" json:"ChiaYiCity"`
	Kinmen        string `xml:"金門縣ItemValue" json:"Kinmen"`
	Matsu         string `xml:"連江縣ItemValue" json:"Matsu"`
}
