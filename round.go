package round

import

// Compilation creates round robin listings for a number of players. For the
// sake of performance
"fmt"

func Compilation(playerNum int) [][]uint16 {
	var odd bool
	if playerNum%2 == 1 {
		odd = true
		playerNum++
	}
	fi := make([][]uint16, playerNum-1)
	t := make([]uint16, (playerNum-1)*2)
	for i := 1; i < playerNum; i++ {
		t[len(t)-i] = uint16(i + 1)
		t[len(t)-playerNum-i+1] = uint16(i + 1)
	}
	if odd {
		t[len(t)-playerNum+1] = 0
		t[0] = 0
	}
	half := playerNum / 2
	for i, _ := range fi {
		fi[i] = make([]uint16, playerNum)
		fi[i][0] = 1
		for j := half - 1; j > 0; j-- {
			fi[i][j*2] = t[playerNum-j+i-1]
		}
		for j := 1; j <= half; j++ {
			fi[i][j*2-1] = t[j+i-1]
		}
	}
	return fi
}

func PrintMatchups(people []string) {
	result := Compilation(len(people))
	for _, r := range result {
		for i, pi := range r {
			if i%2 == 0 {
				fmt.Print(" [")
			}
			if pi == 0 {
				fmt.Print("bye")
			} else {
				fmt.Print(people[int(pi-1)])
			}
			if i%2 == 1 {
				fmt.Print("],")
			} else {
				fmt.Print(", ")
			}
		}
		fmt.Println("")
	}
}
