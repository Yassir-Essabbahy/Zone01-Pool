package piscine

func Rot14(s string) string {
	str := ""
	for _, ch := range s {
		if ch >= 'a' && ch <= 'z' {
			str += string((ch-'a'+14)%26 + 'a')
		} else if ch >= 'A' && ch <= 'Z' {
			str += string((ch-'A'+14)%26 + 'A')
		} else {
			str += string(ch)
		}
	}
	return str
}
