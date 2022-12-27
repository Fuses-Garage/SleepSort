package main

import (
	"fmt"  //入出力に使う
	"sync" //非同期実行に使う
	"time" //sleepに使う
)

func main() {
	var answer []int //ソート後のスライスがここに入る
	var numbers []int
	var n int
	var wg sync.WaitGroup
	print("ソートする要素数を入力してください（整数）：")
	_, err := fmt.Scan(&n) //整数の個数Nを受け取る
	checkErr(err)          //エラーチェック
	wg.Add(n)              //カウンタをN個増やす
	for i := 0; i < n; i++ {
		var num int
		fmt.Printf("%d個めの要素を入力してください（整数）：", i+1)
		_, err := fmt.Scan(&num)       //整数を1つ受け取る
		numbers = append(numbers, num) //配列に追加
		checkErr(err)                  //エラーチェック

	}
	for i := 0; i < n; i++ {
		go SleepWait(numbers[i], &answer, &wg) //非同期実行
	}
	wg.Wait()                  //全部終わったら
	fmt.Printf("%d\n", answer) //答えを表示
}
func checkErr(e error) { //エラーチェック用関数
	if e != nil {
		panic(e)
	}
}
func SleepWait(n int, sl *[]int, wg *sync.WaitGroup) {
	defer wg.Done()                                 //終わったらカウンタを減らす
	time.Sleep(time.Millisecond * time.Duration(n)) //数字ミリ秒待機
	*sl = append(*sl, n)                            //答えを示すスライスにくっ付ける
}
