package main

import (
	"log"
	"os"
)

func doSomething() error {
	err := os.MkdirAll("newdir", 0755)
	if err != nil {
		return err
	}
	// (2) 次にディレクトリが削除される
	defer os.RemoveAll("newdir")

	f, err := os.Create("newdir/newfile")
	if err != nil {
		return err
	}
	// (1) 最初にファイルハンドルが閉じられる
	defer f.Close()
	return nil
}

// deferの呼び出し順によるエラー
func MyTempFile() (*os.File, error) {
	file, err := ioutil.TempFile("", "temp")
	if err != nil {
		return nil, err
	}
	defer file.Close() // Closeが遅すぎる

	if err = os.Rename(file.Name(), file.Name()+".go"); err != nil {
		return nil, err
	}
	return file, nil
}

// fにはdeferが呼び出し時の値が渡される
func doSomething2() {
	f, err := os.Open("test1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	f.Write([]byte("Hello"))

	f, err = os.Open("test2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close(f)

	f.Write([]byte("World"))
}

func main() {
	doSomething()
}
