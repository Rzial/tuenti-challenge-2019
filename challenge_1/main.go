package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
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
	n            int

	withOnion    int
	withoutOnion int
}

type ChallengeResult struct {
	n         int

	tortillas int
}

func readChallenge(n int, scanner *bufio.Scanner) *Challenge {
	challenge := Challenge{}

	if !scanner.Scan() { return nil }
	_, err := fmt.Sscanf(scanner.Text(), "%d %d", &challenge.withOnion, &challenge.withoutOnion)
	if err != nil { return nil }

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

func resolve(challenge Challenge) ChallengeResult {
	onionTortillas := (challenge.withOnion / 2) + (challenge.withOnion % 2)
	cursedTortillas := (challenge.withoutOnion / 2) + (challenge.withoutOnion % 2)

	return ChallengeResult {
		n: challenge.n,
		tortillas: onionTortillas + cursedTortillas,
	}
}

func writeChallenge(result ChallengeResult) {
	fmt.Printf("Case #%d: %d", result.n, result.tortillas)
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
