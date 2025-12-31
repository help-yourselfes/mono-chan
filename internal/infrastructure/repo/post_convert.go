package repo

import "encoding/json"

type mediaList = []string

func jsonMedia(media mediaList) (string, error) {
	mediaJSON, err := json.Marshal(media)
	if err != nil {
		return "", err
	}
	return string(mediaJSON), nil
}

func unJsonMedia(media string) (mediaList, error) {
	var mediaList = make(mediaList, 0)
	err := json.Unmarshal([]byte(media), &mediaList)
	if err != nil {
		return nil, err
	}

	return mediaList, nil
}
