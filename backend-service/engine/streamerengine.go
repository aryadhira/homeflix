package engine

import (
	"bytes"
	"homeflix/configuration"
	"log"
	"os/exec"
	"runtime"
	"strings"
)

type StreamerEngine struct{}

func (se *StreamerEngine) ExecutePeerflix(magnet string) {
	cfg := new(configuration.Config)
	cfgdata := cfg.GetConfig()

	//temporary only support windows
	if runtime.GOOS == "windows" {
		strcmd := "node " + cfgdata.PeerflixPath + " " + magnet + " --vlc -f " + cfgdata.DownloadPath
		parts := strings.Fields(strcmd)
		cmd := exec.Command(parts[0], parts[1:]...)
		var out bytes.Buffer
		var stderr bytes.Buffer
		cmd.Stdout = &out
		cmd.Stderr = &stderr
		err := cmd.Run()
		if err != nil {
			log.Println("error:", err, stderr.String())
			return
		}

		log.Println("out", out.String())
	}
}

func (se *StreamerEngine) StartStreaming(magneturl string) {
	log.Println("Start Streaming...")
	se.ExecutePeerflix(magneturl)
}
