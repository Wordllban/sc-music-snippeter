package processor

import (
	"fmt"
	"sc-music-snippeter/lib/helps"
	"sc-music-snippeter/processor/soundcloud"
	"strings"
)

const (
	soundcloudURL = "soundcloud.com"
	youtubeURL = "youtube.com"
	youtubeURLShort = "youtu.be"
	spotifyURL = "spotify.com"
)


// Receives a message from a client, recognizes the URL in the message and the platform it comes from.
// Returns file name
func UrlProcessor(message string) string {
	serviceUrl := recognizeService(message)
	urlFromMessage := helps.ExtractString(message, serviceUrl)
	
	switch serviceUrl {
		case soundcloudURL:  
			return soundcloud.ProcessSoundCloudURL(urlFromMessage)
		case youtubeURL, youtubeURLShort:
			// youtube processor
			return ""
		case spotifyURL:
			// spotify processor
			return ""
		default:
			fmt.Println("This platform is not supported")
			return ""
	}
}

func recognizeService(message string) string {
	switch {
	case strings.Contains(message, soundcloudURL):
			return soundcloudURL
	case strings.Contains(message, youtubeURL) || strings.Contains(message, youtubeURLShort):
			return youtubeURL
	case strings.Contains(message, spotifyURL):
			return spotifyURL
	default:
			return ""
	}
}