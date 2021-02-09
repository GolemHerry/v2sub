package subscribe

import (
	"bufio"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/olekukonko/tablewriter"
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
	bufReader := bufio.NewReader(decodeReader)
	infos := make([]userInfo, 2)

	for i := 0; i < 2; i++ {
		row, err := bufReader.ReadBytes('\n')
		if err != nil {
			panic(err)
		}

		if len(row) < 8 || string(row[:5]) != "vmess" {
			return errors.New("Invalid protocol")
		}

		dest := make([]byte, base64.StdEncoding.DecodedLen(len(row)-8))
		n, err := base64.StdEncoding.Decode(dest, row[8:])
		if err != nil {
			return err
		}

		var info userInfo
		if err := json.Unmarshal(dest[:n], &info); err != nil {
			return err
		}
		infos[i] = info
	}
	printInfo(infos)

	return nil
}

func printInfo(infos []userInfo) {
	data := make([][]string, 0)
	for _, v := range infos {
		data = append(data, []string{v.Remark, v.Addition, v.Id, v.Ps, v.Net, ""})
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Remark", "Addition", "ID", "Ps", "Net", "type"})
	table.SetFooter([]string{"", "", "f3", "f4", "", ""})
	table.SetBorder(false)

	// table.SetHeaderColor(tablewriter.Colors{tablewriter.Bold, tablewriter.BgGreenColor},
	// 	tablewriter.Colors{tablewriter.FgHiRedColor, tablewriter.Bold, tablewriter.BgBlackColor},
	// 	tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiRedColor},
	// 	tablewriter.Colors{tablewriter.BgRedColor, tablewriter.FgWhiteColor},
	// 	tablewriter.Colors{tablewriter.BgCyanColor, tablewriter.FgWhiteColor},
	// 	tablewriter.Colors{tablewriter.BgRedColor, tablewriter.FgWhiteColor},
	// )

	// table.SetColumnColor(tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiBlackColor},
	// 	tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiRedColor},
	// 	tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiRedColor},
	// 	tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiBlackColor},
	// 	tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiBlackColor},
	// 	tablewriter.Colors{tablewriter.Bold, tablewriter.FgBlackColor})

	// table.SetFooterColor(tablewriter.Colors{}, tablewriter.Colors{},
	// 	tablewriter.Colors{tablewriter.Bold},
	// 	tablewriter.Colors{tablewriter.FgHiRedColor})

	table.AppendBulk(data)
	table.Render()
	//TODO beautify output
}
