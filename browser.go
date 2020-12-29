package wen

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"runtime"
)

// ChromeExecutable returns a string which points to the preferred Chrome
// executable file.
var ChromeExecutable = LocateChrome

// LocateChrome returns a path to the Chrome binary, or an empty string if
// Chrome installation is not found.
func LocateChrome() string {

	// If env variable "LORCACHROME" specified and it exists
	if path, ok := os.LookupEnv("LORCACHROME"); ok {
		if _, err := os.Stat(path); err == nil {
			return path
		}
	}

	var paths []string
	switch runtime.GOOS {
	case "darwin":
		paths = []string{
			"/Applications/Google Chrome.app/Contents/MacOS/Google Chrome",
			"/Applications/Google Chrome Canary.app/Contents/MacOS/Google Chrome Canary",
			"/Applications/Chromium.app/Contents/MacOS/Chromium",
			"/usr/bin/google-chrome-stable",
			"/usr/bin/google-chrome",
			"/usr/bin/chromium",
			"/usr/bin/chromium-browser",
		}
	case "windows":
		paths = []string{
			os.Getenv("ProgramFiles(x86)") + "/Microsoft/edge/Application/msedge.exe",
			os.Getenv("LocalAppData") + "/Google/Chrome/Application/chrome.exe",
			os.Getenv("ProgramFiles") + "/Google/Chrome/Application/chrome.exe",
			os.Getenv("ProgramFiles(x86)") + "/Google/Chrome/Application/chrome.exe",
			os.Getenv("LocalAppData") + "/Chromium/Application/chrome.exe",
			os.Getenv("ProgramFiles") + "/Chromium/Application/chrome.exe",
			os.Getenv("ProgramFiles(x86)") + "/Chromium/Application/chrome.exe",
		}
	default:
		paths = []string{
			"/usr/bin/google-chrome-stable",
			"/usr/bin/google-chrome",
			"/usr/bin/chromium",
			"/usr/bin/chromium-browser",
			"/snap/bin/chromium",
		}
	}

	for _, path := range paths {
		if _, err := os.Stat(path); os.IsNotExist(err) {
			continue
		}
		return path
	}
	return ""
}

type BrowserUI struct {
	cmd    *exec.Cmd
	done   chan struct{}
	tmpDir string
}

// --start-maximized
// --start-normal
// --start-minimized
// --start-fullscreen
var defaultChromeArgs = []string{
	"--disable-background-networking",
	"--disable-background-timer-throttling",
	"--disable-backgrounding-occluded-windows",
	"--disable-breakpad",
	"--disable-client-side-phishing-detection",
	"--disable-default-apps",
	"--disable-dev-shm-usage",
	"--disable-infobars",
	"--disable-extensions",
	"--disable-features=site-per-process",
	"--disable-hang-monitor",
	"--disable-ipc-flooding-protection",
	"--disable-popup-blocking",
	"--disable-prompt-on-repost",
	"--disable-renderer-backgrounding",
	"--disable-sync",
	"--disable-translate",
	"--disable-windows10-custom-titlebar",
	"--metrics-recording-only",
	"--no-first-run",
	"--no-default-browser-check",
	"--safebrowsing-disable-auto-update",
	"--enable-automation",
	"--password-store=basic",
	"--use-mock-keychain",
}

// New returns a new HTML5 UI for the given URL, user profile directory, window
// size and other options passed to the browser engine. If URL is an empty
// string - a blank page is displayed. If user profile directory is an empty
// string - a temporary directory is created and it will be removed on
// ui.Close(). You might want to use "--headless" custom CLI argument to test
// your UI code.

func (u *BrowserUI) Done() <-chan struct{} {
	return u.done
}

func (u *BrowserUI) Close() error {
	// ignore err, as the chrome process might be already dead, when user close the window.
	log.Print("close")
	if state := u.cmd.ProcessState; state == nil || !state.Exited() {
		return u.cmd.Process.Kill()
	}
	<-u.done
	if u.tmpDir != "" {
		if err := os.RemoveAll(u.tmpDir); err != nil {
			return err
		}
	}
	return nil
}

// New returns a new HTML5 UI for the given URL, user profile directory, window
// size and other options passed to the browser engine. If URL is an empty
// string - a blank page is displayed. If user profile directory is an empty
// string - a temporary directory is created and it will be removed on
// ui.Close(). You might want to use "--headless" custom CLI argument to test
// your UI code.
func ChromeApp(url, dir string, width, height int, customArgs ...string) error {
	if url == "" {
		url = "data:text/html,<html></html>"
	}
	tmpDir := ""
	if dir == "" {
		name, err := ioutil.TempDir("", "topo")
		if err != nil {
			return err
		}
		dir, tmpDir = name, name
	}
	args := append(defaultChromeArgs, fmt.Sprintf("--app=%s", url))
	args = append(args, fmt.Sprintf("--user-data-dir=%s", dir))
	if width > 0 && height > 0 {
		args = append(args, fmt.Sprintf("--window-size=%d,%d", width, height))
	}
	args = append(args, customArgs...)
	args = append(args, "--remote-debugging-port=0")
	if ChromeExecutable() == "" {
		return nil
	}
	cmd := exec.Command(ChromeExecutable(), args...)
	if err := cmd.Start(); err != nil {
		return err
	}
	done := make(chan struct{})

	go func() {
		cmd.Wait()
		close(done)
	}()
	u := &BrowserUI{cmd: cmd, done: done, tmpDir: tmpDir}
	go func() {
		// Wait until the interrupt signal arrives or browser window is closed
		sigc := make(chan os.Signal)
		signal.Notify(sigc, os.Interrupt)
		select {
		case <-sigc:
		case <-u.Done():
		}
		log.Println("exiting...")
		u.Close()
		os.Exit(0)
	}()
	return nil
}
