package question13

var values = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanToInt(s string) (ans int) {
	n := len(s)
	for i := range s {
		value := values[s[i]]
		if i < n-1 && value < values[s[i+1]] {
			ans -= value
		} else {
			ans += value
		}
	}
	return
}
