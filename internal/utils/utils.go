package utils

import "fmt"

func ParseCategories(categories []string) (string){
	parsed := ""
	if len(categories) == 0 {
		return parsed
	}

	for i := 0; i < len(categories) - 1; i++ {
		parsed += fmt.Sprintf("%s, ", categories[i])
	}

	parsed += fmt.Sprintf("%s", categories[len(categories) - 1])

	return parsed
}
