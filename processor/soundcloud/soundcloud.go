package soundcloud

import (
	"fmt"
	"os"
	"os/exec"
	"sc-music-snippeter/lib/files"
	"sc-music-snippeter/lib/logger"
)

func ProcessSoundCloudURL(trackUrl string) string {
	audio, err := files.DownloadFile(trackUrl)
	if err != nil {
		panic(err)
	}

	tmpFile, err := os.CreateTemp("", "*.mp3")
	if err != nil {
		logger.LogError("Telegram: Unable to create temporary file: ", err)
		return ""
	}
	defer os.Remove(tmpFile.Name())

	// Write the audio data to the temporary file
	if _, err := tmpFile.Write(audio); err != nil {
		fmt.Println("Failed to write to temporary file:", err)
		return ""
	}
	tmpFile.Close()


	var tmpCutFile = files.TempFileNameCut(tmpFile.Name())

	// Cut audio file
	cmd := exec.Command("ffmpeg", "-i", tmpFile.Name(), "-ss", "15", "-to", "30", "-f", "mp3", tmpCutFile)
	if err := cmd.Run(); err != nil {
		fmt.Println("Failed to execute command:", err)
		return ""
	}

	return tmpCutFile
}