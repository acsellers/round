/*
Package round implements round robin pairing for teams. You can pass an
array to PrintMatchups to print out the matchings to the console,
or use Compilation to get the indexes for the subjects.
*/
package round

import "fmt"

// Compilation creates round robin listings for a number of players. Note that
// team indexes start at 1, not 0. For the
// sake of performance, it takes team indexes and returns and array of the
// matches. If the number of players is odd, the bye team will be paired with
// a team 0.
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
	for i := range fi {
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

/*
PrintMatchups takes an array of people (or players or teams) and create round robin
matchups between them. For odd numbers of people, there will be a bye
team visible in the teams that indicates the team without a pairing
for that game time.
*/
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
