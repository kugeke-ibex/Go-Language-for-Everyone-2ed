package main

import (
	"io"
	"log"
	"os"
	"os/exec"
	"sync"
	"syscall"
	// "github.com/mattn/go-shellwords"
)

func main() {
	tr(os.Stdin, os.Stdout, os.Stderr)

	// OK
	// out, err := exec.Command("ls", "-l", "foo.txt").Output()

	// NG
	// out, err := exec.Command("ls -l foo.txt").Output()

	// シェル経由でコマンド実行
	// out, err := exec.Command("sh", "-c", "ls -l").Output()

	// シェルの機能を使用したコマンド実行
	// out, err := exec.Command("sh", "-c", "some_command || handle_error").Output()

	// arg, err := shellwords.Parse("ls -l foo.txt")
	// argsは["ls", "-l", "foo.txt"]となる

	// out, err := exec.Command(args[0], args[1:]...).Output()
}

func tr(src io.Reader, dst io.Writer, errDst io.Writer) error {
	cmd := exec.Command("tr", "a-z", "A-Z")
	// 実行するコマンドtr a-z A-Z
	stdin, _ := cmd.StdinPipe()
	stdout, _ := cmd.StdoutPipe()
	stderr, _ := cmd.StderrPipe()
	// コマンドの実行を開始する
	if err := cmd.Start(); err != nil {
		return err
	}

	var wg sync.WaitGroup
	wg.Add(3)
	go func () {
		// コマンドの標準入力にsrcからコピーする
		_, err := io.Copy(stdin, src)
		if e, ok := err.(*os.PathError); ok && e.Err == syscall.EPIPE {
			// ignore EPIPE
		} else if err != nil {
			log.Println("failed to write to STDIN", err)
		}

		stdin.Close()
		wg.Done()
	}()

	go func () {
		// コマンドの標準入力をdstにコピーする
		io.Copy(dst, stdout)
		stdout.Close()
		wg.Done()
	}()

	go func () {
		// コマンドの標準エラー出力をerrDstにコピーする
		io.Copy(errDst, stderr)
		stderr.Close()
		wg.Done()
	}()

	wg.Wait()
	// 標準入出力のI/Oを行う goroutineが全て終わるまで待つ
	return cmd.Wait()
	// コマンドの終了を待つ
}
