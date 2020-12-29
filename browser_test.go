package wen

import (
	"testing"
	"time"
)

func TestChromeApp(t *testing.T) {
	ChromeApp("https://baidu.com", "", -1, -1, "--start-maximized")
	time.Sleep(time.Second * 10)
}
