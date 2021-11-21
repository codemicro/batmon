package main

import (
	"bytes"
	_ "embed"
	"io/ioutil"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep/wav"
)

const batteryPath = "/sys/class/power_supply/BAT0/capacity"

var (
	notificationSent bool

	notificationBinary  = "notify-send"
	notificationCommand = []string{"--urgency=critical", "Battery low", "Battery has reached 20% charge - you should plug in now to extend the battery's lifetime"}
)

//go:embed ab-ap.wav
var toneFile []byte

func run() error {

	// Prepare audio - https://github.com/faiface/beep/wiki/Hello,-Beep!
	// This need only be done once
	tr := bytes.NewReader(toneFile)
	audioStreamer, audioFormat, err := wav.Decode(tr)
	if err != nil {
		return err
	}
	speaker.Init(audioFormat.SampleRate, audioFormat.SampleRate.N(time.Second/10))

	for {
		// Read battery level
		fcont, err := ioutil.ReadFile(batteryPath)
		if err != nil {
			return err
		}

		batteryCharge, err := strconv.Atoi(strings.TrimSpace(string(fcont)))
		if err != nil {
			return err
		}

		// Check if battery is 20% or less
		if batteryCharge <= 20 {
			if !notificationSent {
				// Send notification
				err = exec.Command(notificationBinary, notificationCommand...).Run()
				if err != nil {
					panic(err)
				}
				notificationSent = true
				// Make noise
				speaker.Play(audioStreamer)
			}
		} else {
			notificationSent = false
		}

		// Pause for a bit
		time.Sleep(time.Second * 60)
	}
}

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}
