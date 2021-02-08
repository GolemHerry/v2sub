package subscribe

import (
	"bufio"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

const rawFilePath = "/.v2sub/raw.txt"

type userInfo struct {
	Remark   string      `json:"remark"`
	Type     interface{} `json:"type"`
	Addition string      `json:"add"`
	Id       string      `json:"id"`
	Net      string      `json:"net"`
	AlterId  int         `json:"alterId"`
	Ps       string      `json:"ps"`
}

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

func Info() error {
	homedir, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	rawFile, err := os.Open(homedir + rawFilePath)
	if err != nil {
		panic(err)
	}
	defer rawFile.Close()

	decodeReader := base64.NewDecoder(base64.StdEncoding, rawFile)
	dest, err := bufio.NewReader(decodeReader).ReadBytes('\n')
	if err != nil {
		panic(err)
	}

	if len(dest) < 8 || string(dest[:5]) != "vmess" {
		return errors.New("Invalid protocol")
	}

	var info userInfo
	if err := json.Unmarshal(dest[8:], &info); err != nil {
		return err
	}
	printInfo(info)

	return nil
}

func printInfo(info userInfo) {
	fmt.Printf("User info: \n%#v\n", info)
	//TODO beautify output
}
