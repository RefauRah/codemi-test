package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type identity struct {
	NoLoker       int
	TipeIdentitas string
	NoIdentitas   string
}

func main() {
	var identities []identity
	var lokerCount int
	i := 0
	for {
		_, textArr := inputIdentities()
		switch textArr[0] {
		case "init":
			if len(textArr) == 2 {
				intArr, _ := strconv.Atoi(textArr[1])
				lokerCount = intArr
				fmt.Println("Berhasil membuat loker dengan jumlah", lokerCount)
			} else {
				fmt.Println("perintah salah!")
			}
		case "status":
			if len(textArr) == 1 {
				if lokerCount <= 0 {
					fmt.Println("loker belum di inisialisasi")
				}

				fmt.Println("No Loker	Tipe Identitas	No Identitas")
				for _, v := range identities {
					fmt.Println(v.NoLoker, "		", v.TipeIdentitas, "		", v.NoIdentitas)
				}
			} else {
				fmt.Println("perintah salah!")
			}
		case "input":
			if len(textArr) == 3 {
				if lokerCount <= 0 {
					fmt.Println("loker belum di inisialisasi")
				}

				if len(identities) >= lokerCount {
					fmt.Println("loker penuh!")
					break
				}

				var identity identity
				i++
				identity.NoLoker = i
				identity.NoIdentitas = textArr[2]
				identity.TipeIdentitas = textArr[1]

				identities = append(identities, identity)
				fmt.Println("kartu identitas disimpan di loker nomor", identity.NoLoker)
			} else {
				fmt.Println("perintah salah!")
			}
		case "leave":
			if len(textArr) == 2 {
				if lokerCount <= 0 {
					fmt.Println("loker belum di inisialisasi")
				}

				intArr, _ := strconv.Atoi(textArr[1])
				index := intArr - 1
				identities = append(identities[:index], identities[index+1:]...)

				fmt.Println("loker nomor", intArr, "berhasil dikosongkan")
			} else {
				fmt.Println("perintah salah!")
			}
		case "find":
			if len(textArr) == 2 {
				if lokerCount <= 0 {
					fmt.Println("loker belum di inisialisasi")
				}
				ok, loc := find(identities, textArr[1])
				if ok {
					fmt.Println("kartu identitas tersebut berada di loker", loc)
				} else {
					fmt.Println("nomor identitas tidak ditemukan")
				}
			} else {
				fmt.Println("perintah salah!")
			}
		case "search":
			if len(textArr) == 2 {
				if lokerCount <= 0 {
					fmt.Println("loker belum di inisialisasi")
				}
				for _, v := range identities {
					if textArr[1] == v.TipeIdentitas {
						fmt.Print(v.NoIdentitas, ", ")
					}
				}
				fmt.Println()
			} else {
				fmt.Println("perintah salah!")
			}
		case "exit":
			os.Exit(0)
		default:
			fmt.Println("perintah tidak ditemukan!")
			continue
		}
	}
}

func inputIdentities() (string, []string) {
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	text = strings.ToLower(strings.Replace(text, "\n", "", -1))
	textArr := strings.Fields(text)

	return text, textArr
}

func find(a []identity, x string) (bool, int) {
	for i, n := range a {
		if x == n.NoIdentitas {
			return true, i + 1
		}
	}
	return false, 0
}
