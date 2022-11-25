package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

var sesi int

func main() {
	sesi = 1
	sumPlayer := ""
	sumPlayerInt := 0
	sumDice := ""
	sumDiceInt := 0

	step := 0
	var err error
input:
	for {
		switch {
		case step == 0:
			reader := bufio.NewReader(os.Stdin)
			fmt.Print("Tolong masukan jumlah pemain : ")
			sumPlayer, _ = reader.ReadString('\n')
			sumPlayer = strings.TrimSuffix(sumPlayer, "\r\n")
			sumPlayerInt, err = strconv.Atoi(sumPlayer)
			if err != nil || sumPlayerInt == 0 {
				fmt.Print("Anda memasukan data bukan angka \n")
				step = 0
			} else {
				// fmt.Println("menyimpan data jumlah pemain " + sumPlayer)
				// fmt.Println(sumPlayerInt)
				step = 1
			}
		case step == 1:
			reader := bufio.NewReader(os.Stdin)
			fmt.Print("Tolong masukan jumlah dadu : ")
			sumDice, _ = reader.ReadString('\n')
			sumDice = strings.TrimSuffix(sumDice, "\r\n")
			sumDiceInt, err = strconv.Atoi(sumDice)
			if err != nil || sumDiceInt == 0 {
				fmt.Print("Anda memasukan data bukan angka \n")
				step = 1
			} else {
				// fmt.Println("menyimpan data jumlah dadu " + sumDice)
				// fmt.Println(sumDiceInt)
				step = 2
			}
		case step == 2:
			pointPemain := []int{}
			daduPemain := make(map[int][]int)
			if mainDadu(sumPlayerInt, sumDiceInt, pointPemain, daduPemain) {
				step = 3
			}
		default:

			break input

		}

	}

}
func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}
func removeElementByIndex[T any](slice []T, index int) []T {
	return append(slice[:index], slice[index+1:]...)
}
func mainDadu(jumlahPemain int, jumlahDadu int, pointPemain []int, daduPemain map[int][]int) bool {

	for {
		fmt.Println("\n")
		s := (fmt.Sprintf("SESI KE-%d", sesi))
		fmt.Println(s)
		fmt.Println("Lempar Dadu.......")

		for i := 1; i <= jumlahPemain; i++ {

			arrays := []int{}
			if sesi == 1 {
				pointPemain = append(pointPemain, 0)
				for k := 1; k <= jumlahDadu; k++ {
					arrays = append(arrays, randInt(1, 7))
				}

			} else {
				for k := 1; k <= len(daduPemain[i-1]); k++ {
					arrays = append(arrays, randInt(1, 7))
				}
			}
			daduPemain[i-1] = arrays
			r := fmt.Sprintf("Pemain %d (%d) %d", i, pointPemain[i-1], daduPemain[i-1])
			fmt.Println(r)
		}

		fmt.Println("setelah evaluasi")
		tempArr := make(map[int][]int)
		for i := 1; i <= jumlahPemain; i++ {
			TMP := []int{}
			for _, isiDadu := range daduPemain[i-1] {
				if isiDadu == 6 {
					pointPemain[i-1] += 1
				} else if isiDadu == 1 {
					if i == jumlahPemain {
						tempArr[0] = append(tempArr[0], 1)
					} else {
						tempArr[i] = append(tempArr[i], 1)
					}
				} else {
					TMP = append(TMP, isiDadu)
				}

				// fmt.Println(isiDadu)
			}
			tempArr[i-1] = append(tempArr[i-1], TMP...)
			// r := fmt.Sprintf("Pemain %d (%d) %d", i, pointPemain[i-1], daduPemain[i-1])
			// fmt.Println(r)
		}

		for inds, dss := range tempArr {
			daduPemain[inds] = dss

		}
		cekTinggal := 0
		for e := 1; e <= jumlahPemain; e++ {
			if len(daduPemain[e-1]) > 0 {
				cekTinggal++
			}
			r := fmt.Sprintf("Pemain %d (%d) %d", e, pointPemain[e-1], daduPemain[e-1])
			fmt.Println(r)
		}
		sesi++
		if cekTinggal <= 1 {
			menang := []int{}
			s := 0
			for _, bil := range pointPemain {
				if bil >= s && bil != 0 {
					s = bil
				}
			}
			for ind, bil := range pointPemain {
				if bil == s {
					menang = append(menang, ind+1)
				}
			}
			result := "Pemenang adalah "
			for ssss, dd := range menang {
				if ssss != len(menang)-1 {
					result += fmt.Sprintf("#%d, ", dd)

				} else {
					result += fmt.Sprintf("#%d", dd)
				}
			}
			fmt.Println("\n ", result)
			return true
		}
	}

}
