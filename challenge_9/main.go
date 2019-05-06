package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"sync"
)

const (
	NWorkers = 4
)

type Input struct {
	cases      int

	challenges []Challenge
}

type Challenge struct {
	n      int

	op1    string
	op2    string
	result string
}

type ChallengeResult struct {
	n        int

	op1      int
	operator string
	op2      int
	result   int
}

func readChallenge(n int, scanner *bufio.Scanner) *Challenge {
	challenge := Challenge{}

	var op, eq string
	if !scanner.Scan() { return nil }

	_, err := fmt.Sscanf(scanner.Text(), "%s %s %s %s %s",
		&challenge.op1, &op, &challenge.op2, &eq, &challenge.result)

	if err != nil { return nil }
	if op != "OPERATOR" { return nil }
	if eq != "=" { return nil }

	challenge.n = n

	return &challenge
}

func getInput(reader io.Reader) *Input {
	scanner := bufio.NewScanner(reader)
	input := Input {
		cases: 0,
		challenges: make([]Challenge, 0),
	}

	if !scanner.Scan() { return nil }
	_, err := fmt.Sscanf(scanner.Text(), "%d", &input.cases)
	if err != nil { return nil }

	for i := input.cases; i > 0; i-- {
		challenge := readChallenge(input.cases - i + 1, scanner)
		if challenge == nil {
			return nil
		}

		input.challenges = append(input.challenges, *challenge)
	}

	return &input
}

func removeDuplicates(elements []int) []int {
	found := map[int]bool{}
	result := make([]int, 0)

	for v := range elements {
		if found[elements[v]] != true {
			found[elements[v]] = true
			result = append(result, elements[v])
		}
	}

	return result
}

func combinePosibleRomaji(romajiMultipliers []int, romajiUnits []int, acc int, posibleRomaji *[]int) {
	if len(romajiMultipliers) == 0 {
		if len(romajiUnits) == 0 {
			*posibleRomaji = append(*posibleRomaji, acc)
		} else  {
			*posibleRomaji = append(*posibleRomaji, acc + romajiUnits[0])
		}

		return
	}

	var multipliers, units []int
	for i, mul := range romajiMultipliers {
		multipliers = make([]int, len(romajiMultipliers))
		copy(multipliers, romajiMultipliers)

		multipliers = append(multipliers[:i], multipliers[i + 1:]...)

		if mul != 10000 && len(romajiMultipliers) >= len(romajiUnits) {
			combinePosibleRomaji(multipliers, romajiUnits, acc + mul, posibleRomaji)
		}

		for j, unit := range romajiUnits {
			if mul != 10000 && unit == 1 { continue }

			units = make([]int, len(romajiUnits))
			copy(units, romajiUnits)

			units = append(units[:j], units[j + 1:]...)
			combinePosibleRomaji(multipliers, units, acc + (mul * unit), posibleRomaji)
		}
	}
}

func kanji2romaji(number string) []int {
	kanji := map[string]int {
		"一": 1, "二": 2, "三": 3, "四": 4, "五": 5, "六": 6, "七": 7, "八": 8, "九": 9,
		"十": 10,
		"百": 100,
		"千": 1000,
		"万": 10000,
	}

	var romajiUnits, romajiMultipliers []int
	for _, n := range number {
		if kanji[string(n)] > 9 {
			romajiMultipliers = append(romajiMultipliers, kanji[string(n)])
		} else {
			romajiUnits = append(romajiUnits, kanji[string(n)])
		}
	}

	sort.Slice(romajiMultipliers, func(i, j int) bool {
		return romajiMultipliers[i] > romajiMultipliers[j]
	})

	var possibleRomaji []int
	combinePosibleRomaji(romajiMultipliers, romajiUnits, 0, &possibleRomaji)

	return removeDuplicates(possibleRomaji)
}

func resolve(challenge Challenge) ChallengeResult {
	op1map := kanji2romaji(challenge.op1)
	op2map := kanji2romaji(challenge.op2)
	resultMap := kanji2romaji(challenge.result)

	// Try ADD
	for _, op1 := range op1map {
		for _, op2 := range op2map {
			for _, result := range resultMap{
				op := op1 + op2
				if op == result {
					return ChallengeResult {
						n: challenge.n,

						op1: op1,
						operator: "+",
						op2: op2,
						result: result,
					}
				}
			}
		}
	}

	// Try DIFF
	for _, op1 := range op1map {
		for _, op2 := range op2map {
			for _, result := range resultMap{
				op := op1 - op2
				if op == result {
					return ChallengeResult {
						n: challenge.n,

						op1: op1,
						operator: "-",
						op2: op2,
						result: result,
					}
				}
			}
		}
	}

	// Try MUL
	for _, op1 := range op1map {
		for _, op2 := range op2map {
			for _, result := range resultMap{
				op := op1 * op2
				if op == result {
					return ChallengeResult {
						n: challenge.n,

						op1: op1,
						operator: "*",
						op2: op2,
						result: result,
					}
				}
			}
		}
	}

	return ChallengeResult {
		n: challenge.n,
	}
}

func writeChallenge(result ChallengeResult) {
	fmt.Printf("Case #%d: %d %s %d = %d", result.n, result.op1, result.operator, result.op2, result.result)
}

func runWorkerInput(in <- chan Challenge, out chan <- ChallengeResult, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		challenge, ok := <- in
		if !ok { return }

		out <- resolve(challenge)
	}
}

func runWorkerOutput(out <- chan ChallengeResult, output []ChallengeResult, cases int, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		if cases == 0 { return }

		result, ok := <- out
		if !ok { return }

		output[result.n - 1] = result

		cases--
	}
}

func mainSequential() {
	input := getInput(os.Stdin)
	for i := 0; i < input.cases; i++ {
		if i != 0 {
			fmt.Println()
		}

		challenge := input.challenges[i]
		writeChallenge(resolve(challenge))
	}
}

func mainParallel() {
	input := getInput(os.Stdin)
	output := make([]ChallengeResult, input.cases)

	wgInput := sync.WaitGroup{}
	wgOutput := sync.WaitGroup{}
	in, out := make(chan Challenge), make(chan ChallengeResult)

	// Start Resolvers
	wgInput.Add(NWorkers)
	for i := 0; i < NWorkers; i++ {
		go runWorkerInput(in, out, &wgInput)
	}

	// Start Printers
	wgOutput.Add(1)
	go runWorkerOutput(out, output, input.cases, &wgOutput)

	for i := 0; i < input.cases; i++ {
		in <- input.challenges[i]
	}

	wgOutput.Wait()

	close(in)
	close(out)

	wgInput.Wait()

	for i := 0; i < input.cases; i++ {
		if i != 0 {
			fmt.Println()
		}

		writeChallenge(output[i])
	}
}

func main() {
	// Single Thread
	//mainSequential()

	// Threaded
	mainParallel()
}
