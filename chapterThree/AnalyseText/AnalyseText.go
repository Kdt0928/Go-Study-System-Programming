package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"strings"
)

var source = `1行目
2行目
3行目`

var source2 = "123 1.234 1.0e4 test"

var source3 = "1234, 1.2344, 1.0e5, testtest"

var csvSource = `13101,"100  ","1000003","ﾄｳｷｮｳﾄ","ﾁﾖﾀﾞｸ","ﾋﾄﾂﾊﾞｼ(1ﾁｮｳﾒ)","東京都","千代田区","一ツ橋（１丁目）",1,0,1,0,0,0
13101,"101  ","1010003","ﾄｳｷｮｳﾄ","ﾁﾖﾀﾞｸ","ﾋﾄﾂﾊﾞｼ(2ﾁｮｳﾒ)","東京都","千代田区","一ツ橋（２丁目）",1,0,1,0,0,0
13101,"100  ","1000012","ﾄｳｷｮｳﾄ","ﾁﾖﾀﾞｸ","ﾋﾋﾞﾔｺｳｴﾝ","東京都","千代田区","日比谷公園",0,0,0,0,0,0
13101,"102  ","1020093","ﾄｳｷｮｳﾄ","ﾁﾖﾀﾞｸ","ﾋﾗｶﾜﾁｮｳ","東京都","千代田区","平河町",0,0,1,0,0,0
13101,"102  ","1020071","ﾄｳｷｮｳﾄ","ﾁﾖﾀﾞｸ","ﾌｼﾞﾐ","東京都","千代田区","富士見",0,0,1,0,0,0`

func main() {
	// buifo.Readerを使って改行で分割
	reader := bufio.NewReader(strings.NewReader(source))
	for {
		line, err := reader.ReadString('\n')
		fmt.Printf("%#v\n", line)
		if err == io.EOF {
			break
		}
	}

	// 改行で分割(デフォルト)
	scanner := bufio.NewScanner(strings.NewReader(source))
	for scanner.Scan() {
		fmt.Printf("%#v\n", scanner.Text())
	}

	// 単語単位で分割
	scanner2 := bufio.NewScanner(strings.NewReader(source))
	scanner2.Split(bufio.ScanWords)
	for scanner2.Scan() {
		fmt.Printf("%#v\n", scanner2.Text())
	}

	// データを指定して解析
	reader2 := strings.NewReader(source2)
	var i int
	var f, g float64
	var s string
	fmt.Fscan(reader2, &i, &f, &g, &s)
	fmt.Printf("i=%#v f=%#v g=%#v s=%#v\n", i, f, g, s)

	reader3 := strings.NewReader(source3)
	var i3 int
	var f3, g3 float64
	var s3 string
	fmt.Fscanf(reader3, "%v, %v, %v, %v", &i3, &f3, &g3, &s3)
	fmt.Printf("i=%#v f=%#v g=%#v s=%#v\n", i3, f3, g3, s3)

	// csv形式の文字列解析
	reader4 := strings.NewReader(csvSource)
	csvReader := csv.NewReader(reader4)
	for {
		line, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		fmt.Println(line[2], line[6:9])
	}

}
