package stretcher

import (
	"bytes"
	"io"
	"log"
	"os"
)

var LogBuffer bytes.Buffer // グローバル変数

func init() {
	// logを os.Stderr と LogBuffer両方に書き込む
	log.SetOutPut(io.MultiWriter(os.Stderr, &LogBuffer))
}
