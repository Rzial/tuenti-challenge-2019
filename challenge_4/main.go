package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
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
	n    int

	list []int
}

type ChallengeResult struct {
	n           int

	numerator   int
	denominator int
}

func readChallenge(n int, scanner *bufio.Scanner) *Challenge {
	challenge := Challenge{}

	var listLen int
	if !scanner.Scan() { return nil }
	_, err := fmt.Sscanf(scanner.Text(), "%d", &listLen)
	if err != nil { return nil }

	challenge.list = make([]int, listLen)

	if !scanner.Scan() { return nil }
	stringList := strings.Split(scanner.Text(), " ")
	for i := 0; i < listLen; i++ {
		_, err := fmt.Sscanf(stringList[i], "%d", &challenge.list[i])
		if err != nil { return nil }
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

func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}

	return a
}

func lcm(a, b int, rest ...int) int {
	result := a * b / gcd(a, b)

	for i := 0; i < len(rest); i++ {
		result = lcm(result, rest[i])
	}

	return result
}

func resolve(challenge Challenge) ChallengeResult {
	sort.Ints(challenge.list)

	maxItem := challenge.list[len(challenge.list) - 1]

	// Count how many items have the list
	itemCount := make([]int, maxItem)
	for i := 0; i < len(challenge.list); i++ {
		itemCount[challenge.list[i] - 1]++
	}

	// Get the list with items
	usableList := make([]int, 0)
	for i := 0; i < len(itemCount); i++ {
		if itemCount[i] != 0 {
			usableList = append(usableList, i + 1)
		}
	}

	// Get the lcm of those items
	var listLCM int
	switch len(usableList) {
	case 1:
		listLCM = usableList[0]
	case 2:
		listLCM = lcm(usableList[0], usableList[1])
	default:
		listLCM = lcm(usableList[0], usableList[1], usableList[2:]...)
	}

	var numerator, denominator int
	for i := 0; i < len(usableList); i++ {
		// Number of total candies
		numerator += itemCount[usableList[i] - 1] * listLCM

		// Number of atendees
		denominator += itemCount[usableList[i] - 1] * listLCM / usableList[i]
	}

	fractionGCD := gcd(numerator, denominator)

	return ChallengeResult {
		n: challenge.n,

		numerator: numerator / fractionGCD,
		denominator: denominator / fractionGCD,
	}
}

func writeChallenge(result ChallengeResult) {
	fmt.Printf("Case #%d: %d/%d", result.n, result.numerator, result.denominator)
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
