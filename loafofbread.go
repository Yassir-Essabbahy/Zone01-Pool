package piscine

func LoafOfBread(str string) string {
	if str == "" {
		return "\n"
	}
	if len(str) < 5 {
		return "Invalid Output\n"
	}
	s := ""
	b := 0
	for a := 0; a < len(str); a++ {
		if b < 5 && str[a] == ' ' {
			continue
		}
		if b == 5 {
			if a != len(str)-1 && str[a+1] == ' ' {
				continue
			}
			if a == len(str)-1 {
				break
			}
			s += " "
			b = 0
			continue
		}
		s += string(str[a])
		b++
	}
	s += "\n"
	return s
}
