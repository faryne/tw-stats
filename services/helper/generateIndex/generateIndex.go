package generateIndex

import (
	"encoding/json"
	"fmt"
	"github.com/faryne/tw-stats/constants"
	"io/fs"
	"os"
)

var totalIndexes = make(map[string]string)

func Add(value string) {
	keyLength := len(totalIndexes)
	totalIndexes[fmt.Sprintf("%04d", keyLength-1)] = value
}

func Generate() {
	// 寫入總目錄檔
	fpMain, err := os.OpenFile(constants.RootDirName+"/index.json", os.O_CREATE|os.O_WRONLY, fs.ModePerm)
	if err != nil {
		fmt.Println("cannot generate main index file")
	} else {
		mainIndex, _ := json.MarshalIndent(totalIndexes, "", "    ")
		fpMain.Write(mainIndex)
		fpMain.Close()
	}
}
