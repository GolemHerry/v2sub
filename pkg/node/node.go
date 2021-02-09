package node

import (
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/olekukonko/tablewriter"
)

func Init() {
	homedir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	f, err := os.Open(homedir + "/.v2sub/raw.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	decodeReader := base64.NewDecoder(base64.StdEncoding, f)
	dest, err := ioutil.ReadAll(decodeReader)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(dest), "\n")

	for i := 0; i < len(lines)-1; i++ {
		if len(lines[i+1]) < 8 || lines[i+1][:5] != "vmess" {
			continue
		}
		nodes = append(nodes, parseNode(lines[i+1]))
	}
	log.Println("Init done")
}

func parseNode(line string) Node {
	if len(line) < 8 || line[:5] != "vmess" {
		return Node{}
	}

	dest := make([]byte, base64.StdEncoding.EncodedLen(len(line)-8))
	n, err := base64.StdEncoding.Decode(dest, []byte(line[8:]))
	if err != nil {
		panic(err)
	}

	var node Node
	if err := json.Unmarshal(dest[:n], &node); err != nil {
		panic(err)
	}

	return node
}

func ListNode() {
	data := make([][]string, 0)

	for _, n := range nodes {
		data = append(data, []string{n.Name, n.Host, n.Path,
			n.Add, n.Net, n.Vip, strconv.Itoa(n.Port),
			n.Id, strconv.Itoa(n.Class), "IP", "time"})
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"name", "host", "path", "addition",
		"net", "vip", "port", "id", "class", "ip", "time"})
	table.SetBorder(false)
	table.AppendBulk(data)
	table.Render()
}
