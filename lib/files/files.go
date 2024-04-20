package files

import (
	"io"
	"net/http"
	"os"
)

func DownloadFile(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
			return nil, err
	}

	return data, err
}

func TempFileNameCut(tempFileName string) string {
	return tempFileName + ".cut.mp3"
}

func DeleteFile(name string) error {
	return os.Remove(name)
}