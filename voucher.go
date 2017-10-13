package voucher

// Generat, generate voucher //
// input configuration.Config
func Generate(input string) []string {

	// validate input, if input = 1 then return as is
	if len(input) == 1 {
		return []string{input}
	}

	runes := []rune(input)
	subPermutations := Generate(string(runes[0 : len(input)-1]))

	result := []string{}
	for _, s := range subPermutations {
		result = append(result, merge([]rune(s), runes[len(input)-1])...)
	}

	return result
}

// Insert a char in a word
func merge(ins []rune, c rune) (result []string) {
	for i := 0; i <= len(ins); i++ {
		result = append(result, string(ins[:i])+string(c)+string(ins[i:]))
	}
	return
}
