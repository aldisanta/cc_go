package camel

import "errors"

// Tew is ... you get it
func Tew(camelSum int, positionBets ...[]int) (int, error) {
	matchSum := 0
	// validates camelSum
	if camelSum < 2 || camelSum > 100 {
		return matchSum, errors.New("Where's the Camels ?")
	}
	// validates positionBets
	for _, pos := range positionBets {
		if camelSum != len(pos) {
			return matchSum, errors.New("Where's the Camels ?")
		}
	}
	// create match slice
	var matchSlice [][]int
	for _, camels := range positionBets {
		for i := 0; i < camelSum; i++ {
			for ii := 0; ii < camelSum; ii++ {
				if i < ii {
					matchSlice = append(matchSlice, []int{camels[i], camels[ii]})
				}
			}
		}
	}
	// search internal slice
	// var matchedSlice []int
	for _, needle := range matchSlice {
		equalNum := 0
		for _, hay := range matchSlice {
			// do slice matching
			equal := true
			for i := range hay {
				if hay[i] != needle[i] {
					equal = false
					break
				}
			}
			if equal {
				equalNum++
			}
		}
		if equalNum == len(positionBets) {
			matchSum++
		}
	}
	return matchSum / len(positionBets), nil
}
