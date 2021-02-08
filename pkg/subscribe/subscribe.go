package subscribe

import (
	"io/ioutil"
	"net/http"
	"os"
)

const rawFilePath = "/.v2sub/raw.txt"

func Update(url string) error {
	client := http.Client{}

	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	res, err := client.Do(request)
	if err != nil {
		return err
	}
	rawData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil
	}
	homedir, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	if f, err := os.Stat(homedir + "/.v2sub"); err != nil || !f.IsDir() {
		if err := os.Mkdir(homedir+"/.v2sub", os.ModePerm); err != nil {
			return err
		}
	}

	if err = ioutil.WriteFile(homedir+rawFilePath, rawData, os.ModePerm); err != nil {
		return err
	}

	return nil
}
