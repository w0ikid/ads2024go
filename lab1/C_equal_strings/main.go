package main

func processString(s string) string{
	// code
}

func compare(s1, s2 string) string{
	if processString(s1) == processString(s2){
		return "Yes"
	} else {
		return "No"
	}
}

// Answer






























































// func processString(s string) string{
// 	result := []rune{}
// 	for _, ch := range s {
// 		if ch == '#'{
// 			if len(result) > 0{
// 				result = result[:len(result)-1]
// 			}
// 		} else {
// 			result = append(result, ch)
// 		}
// 	}
// 	return string(result)
// }