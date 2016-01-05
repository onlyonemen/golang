// files project main.go
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type label struct {
	Name  string
	Count int
}

var labels = []label{{"normal", 0}, {"back", 0}, {"land", 0}, {"neptune", 0}, {"pod", 0}, {"smurf", 0}, {"teardrop", 0}, {"ipsweep", 0}, {"nmap", 0}, {"portsweep", 0}, {"satan", 0}, {"ftp_write", 0}, {"guess_passwd", 0}, {"imap", 0}, {"multihop", 0}, {"phf", 0}, {"spy", 0}, {"warezclient", 0}, {"warezmaster", 0}, {"buffer_overflow", 0}, {"loadmodule", 0}, {"perl", 0}, {"rootkit", 0}}

//var labels = []string{"normal", "back", "land", "neptune"}

func main() {
	readLine("C:/Users/h154651/Downloads/kddcup.data/kddcup.data.corrected")
//	readConfig("./setting.txt")
}
func readLine(fileName string) {
	fmt.Println("start to read file...")
	f, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	rd := bufio.NewReader(f)
	i := 0

	for {
		line, err := rd.ReadString('\n') //以'\n'为结束符读入一行

		if err != nil || io.EOF == err {
			break
		}
		i++
		line = strings.TrimSpace(line)      //去掉首尾的换行符
		line = strings.TrimRight(line, ".") //去掉末尾的.
		a := strings.Split(line, ",")       //根据逗号分组
		for j := 0; j < len(labels); j++ {
			if a[41] == labels[j].Name {
				labels[j].Count++
				
			}
		}
	}
	fmt.Println("total is:" + strconv.Itoa(i))
	for n := 0; n < len(labels); n++ {
		fmt.Println("NO",n,":", labels[n].Count)
	}

}
func readConfig(fileName string) {
	f, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	rd := bufio.NewReader(f)
	for {
		line, err := rd.ReadString('\n') //以'\n'为结束符读入一行

		if err != nil || io.EOF == err {
			break
		}
		line = strings.TrimSpace(line) //去掉首尾的换行符
		a := strings.Split(line, "=")  //根据逗号分组
		fmt.Println("setting:" + strings.TrimSpace(a[1]))

	}
}
