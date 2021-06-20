// 2.9 問題
// Q2.2 CSV出力

package main

import (
	"encoding/csv"
	"os"
)

func main() {
	// 標準出力にCSVデータを出力
	csvWriter := csv.NewWriter(os.Stdout)
	recode := []string{"encoding/csv.Writer 1st example", "encoding/csv.Writer 2nd example"}
	csvWriter.Write(recode)
	csvWriter.Flush()

	// ファイルにCSVデータを出力
	csvFile, err := os.Create("csvdata.csv")
	if err != nil {
		panic(err)
	}
	textWriter := csv.NewWriter(csvFile)
	textWriter.Write(recode)
	textWriter.Flush()
}
