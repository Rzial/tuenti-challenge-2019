package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
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
	n       int

	maxWord int
	words   []string
}

type ChallengeResult struct {
	n          int

	directMap  map[rune][]rune
	inverseMap map[rune][]rune
}

func readChallenge(n int, scanner *bufio.Scanner) *Challenge {
	challenge := Challenge{}

	var words int
	if !scanner.Scan() { return nil }
	_, err := fmt.Sscanf(scanner.Text(), "%d", &words)
	if err != nil { return nil }

	challenge.maxWord = 0
	challenge.words = make([]string, words)

	for i := 0; i < words; i++  {
		if !scanner.Scan() { return nil }
		challenge.words[i] = scanner.Text()

		if len(challenge.words[i]) > challenge.maxWord {
			challenge.maxWord = len(challenge.words[i])
		}
	}

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

func normalizePath(entry rune, directMap map[rune][]rune, inverseMap map[rune][]rune, depth int) int {
	if len(directMap[entry]) == 0 { return depth }
	if len(directMap[entry]) == 1 { return normalizePath(directMap[entry][0], directMap, inverseMap, depth + 1) }

	maxSize := 0
	longest := 0
	for i, target := range directMap[entry] {
		size := normalizePath(target, directMap, inverseMap, depth + 1)

		if size > maxSize {
			maxSize = size
			longest = i
		}
	}

	for i, target := range directMap[entry] {
		if i == longest { continue }

		sourceIndex := strings.Index(string(inverseMap[target]), string(entry))
		inverseMap[target] = append(inverseMap[target][:sourceIndex], inverseMap[target][sourceIndex + 1:]...)
	}

	directMap[entry] = directMap[entry][longest:longest+1]

	return maxSize
}

func resolve(challenge Challenge) ChallengeResult {
	transitionMap := make([][]string, challenge.maxWord)
	transitionTokens := ""
	transitionStrings := make([]string, challenge.maxWord)

	for _, word := range challenge.words {
		for i, char := range word {
			if strings.Index(transitionTokens, string(char)) == -1 {
				transitionTokens = string(append([]rune(transitionTokens), rune(char)))
			}

			if transitionStrings[i] == "" {
				transitionStrings[i] = string(char)
				continue
			}

			lastChar := rune(transitionStrings[i][len(transitionStrings[i]) - 1])
			if lastChar != char {
				// Reset the lower character groups
				for j := i + 1; j < len(transitionStrings); j++ {
					if len(transitionStrings[j]) > 1 {
						transitionMap[j] = append(transitionMap[j], transitionStrings[j])
					}

					transitionStrings[j] = ""
				}

				transitionStrings[i] = string(append([]rune(transitionStrings[i]), rune(char)))
			}
		}
	}

	for i := 0; i < len(transitionStrings); i++ {
		if len(transitionStrings[i]) > 1 {
			transitionMap[i] = append(transitionMap[i], transitionStrings[i])
		}

		transitionStrings[i] = ""
	}

	directMap := make(map[rune][]rune, len(transitionTokens))
	inverseMap := make(map[rune][]rune, len(transitionTokens))

	for _, token := range transitionTokens {
		directMap[rune(token)] = make([]rune, 0)
		inverseMap[rune(token)] = make([]rune, 0)
	}

	for _, transitionGroup := range transitionMap {
		for _, transition := range transitionGroup {
			for i := 0; i < len(transition) - 1; i++ {
				source := rune(transition[i])
				target := rune(transition[i + 1])

				directMap[source] = append(directMap[source], target)
				inverseMap[target] = append(inverseMap[target], source)
			}
		}
	}

	// Get all the entry points
	for target := range inverseMap {
		// Entry point
		if len(inverseMap[target]) == 0 {
			normalizePath(target, directMap, inverseMap, 0)
		}
	}

	return ChallengeResult {
		n: challenge.n,

		directMap: directMap,
		inverseMap: inverseMap,
	}
}

func writeChallenge(result ChallengeResult) {
	nEntries := 0
	nExits := 0
	nMultiples := 0

	for source := range result.directMap {
		if len(result.directMap[source]) == 0 {nExits++}
		if len(result.directMap[source]) > 1 {nMultiples++}
	}

	var entry rune
	for target := range result.inverseMap {
		if len(result.inverseMap[target]) == 0 {
			entry = target
			nEntries++
		}
		if len(result.inverseMap[target]) > 1 {nMultiples++}
	}

	var resultString string
	if nEntries == 1 && nExits == 1 && nMultiples == 0 {
		resultString += string(entry)
		for {
			if len(result.directMap[entry]) == 0 { break }
			entry = result.directMap[entry][0]
			resultString += " " + string(entry)
		}
	} else {
		resultString = "AMBIGUOUS"
	}

	fmt.Printf("Case #%d: %s", result.n, resultString)
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
