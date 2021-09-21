## 行政院主計總處-縣市指標

### 產生緣由
1. 因為原資料是將每年度的所有指標全部產生於一個檔案，導致檔案內容過大，只是要讀取某一個指標就變得有點困擾。
2. Golang 初心者如我想用這玩意解析 XML 試試看。

### 資料來源
[縣市指標](https://data.gov.tw/dataset/10935)

### 使用方式
1. 請到 __docs__ 目錄內找尋想要查看的指標
2. 目錄內會有該指標各年度的資料，e.g. __2015.json__ 代表是該指標 2015 年資料
3. __index.json__ 則為全部年度彙總資料

### 欄位說明
| 欄位名稱 | 欄位意義 |
| ---- | ---- | 
| Name | 指標名稱 |
| Unit | 單位 |
| Explain | 指標定義內容 |
| ByYear | 年度，使用民國紀年 |
| Total | 總計。此欄位可能為 0 或是 「...」|
| Taiwan | 臺灣地區的值 |
| NewTaipei | 新北市的值 |
| Taipei | 臺北市的值 | 
| Taoyuan | 桃園市的值 |
| Taichung | 臺中市的值 |
| Tainan | 臺南市的值 | 
| Kaohsiung | 高雄市的值 | 
| Ilan | 宜蘭縣的值 |
| HsinchuCounty | 新竹縣的值 |
| Miaoli | 苗栗縣的值 |
| Changhwa | 彰化縣的值 |
| Nantou | 南投縣的值 |
| Yunlin | 雲林縣的值 |
| ChiaYiCounty | 嘉義縣的值 |
| Pingtung | 屏東縣的值 |
| Taitung | 臺東縣的值 | 
| Hualien | 花蓮縣的值 | 
| Penghu | 澎湖縣的值 | 
| Keelung | 基隆市的值 | 
| HsinchuCity | 新竹市的值 | 
| ChiaYiCity | 嘉義市的值 | 
| Kinmen | 金門縣的值 | 
| Matsu | 連江縣的值 | 