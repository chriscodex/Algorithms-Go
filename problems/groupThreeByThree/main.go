package main

/* Send a string and group 3 by 3 */
func getThreeFirst(s string) *[]string {
	arr := []string{}
	str := ""
	for i := 1; i < len(s)+1; i++ {
		if i%3 == 0 {
			str += string(s[i-1])
			arr = append(arr, str)
			str = ""
		} else {
			str += string(s[i-1])
		}
	}
	if str != "" {
		arr = append(arr, str)
	}
	return &arr
}
