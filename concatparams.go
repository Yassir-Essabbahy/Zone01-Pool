package piscine

func ConcatParams(args []string) string {
	s := ""
	for i := 0; i < len(args); i++ {
		if i != len(args)-1 {
			s += string(args[i]) + "\n"
		} else {
			s += string(args[i])
		}
	}
	return s
}
