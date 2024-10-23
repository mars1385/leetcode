package stack

func deleteAtIndex(slice []rune, index int) []rune {

	return append(slice[:index], slice[index+1:]...)

}

func IsValid(s string) bool {
	list := []rune{}

	for _, char := range s {

		if char == '(' || char == '[' || char == '{' {
			list = append(list, char)
		}

		if len(list) == 0 {
			return false
		}
		temp := list[len(list)-1]

		if (char == ')' && temp != '(') || (char == ']' && temp != '[') || (char == '}' && temp != '{') {

			return false
		}

		if char == ')' || char == ']' || char == '}' {
			list = deleteAtIndex(list, len(list)-1)
		}

	}

	return len(list) <= 0
}
