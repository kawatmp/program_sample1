package main

import "fmt"
import "time"
import "math"
import "os"
import "bufio"
import "strings"

func writeFile (filename string) int {
	fpw, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	w := bufio.NewWriter(fpw)
	for i:=0; i<100; i++ {
		str := "abv"
		fmt.Fprintf(w, "%d,Dummy,%s%d\n", i, str, i)
	}
	w.Flush()
	fpw.Close()
	return 1
}


//	UTF-8のCSVファイルを読み込み
//	ある条件に一致したらハッシュに入れ、サブルーチンでループを回し、
//	結果をファイル出力する

func main() {
	//	テストデータを出力
	writeFile("test999.txt")
	//	テキストファイルの読み込み
	m := map[string]int{}
	var fp *os.File
	var err error
	fp, err = os.Open("test999.txt")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
//		fmt.Println(scanner.Text())
		arr := strings.Split(scanner.Text(), ",")
		fmt.Println(arr[0])
//		m[arr[2]] = arr[0]
	}
	fp.Close()

	//	連想配列を作成し、値を入れて、出力する
//	m := map[string]int{}
//	m["apple"] = 150
//	m["banana"] = 200
//	m["lemon"] = 300
	for k,v := range m {
		fmt.Println(k,v)
	}


	t1 := time.Now()
	fmt.Println(t1)

	//	1から10万までの間の素数を表示する
	fmt.Print("2 ")
	for i := 3; i < 100000; i+=2 {		//	偶数は飛ばす
		var k = 0
		for j :=3; j < int(math.Sqrt(float64(i))); j+=2  {		//	100の約数、2,  50    4  25   5  20
			if i % j == 0 {
				k = 1
				break
			}
		}
		if k == 0 {
//			fmt.Print(i)
//			fmt.Print(" ")
		}
	}
	fmt.Println("");

	t2 := time.Now()
	fmt.Println(t2)
}
