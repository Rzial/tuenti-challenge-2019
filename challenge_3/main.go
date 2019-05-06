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

type Points []Point

type Point struct {
	x int
	y int
}

func (p Points) Len() int {
	return len(p)
}

func (p Points) Less(i, j int) bool {
	if p[i].x == p[j].x {
		return p[i].y < p[j].y
	} else {
		return p[i].x < p[j].x
	}
}

func (p Points) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

type Challenge struct {
	n       int

	width   int
	height  int

	folds   []string
	punches Points
}

type ChallengeResult struct {
	n      int

	punches Points
}

func readChallenge(n int, scanner *bufio.Scanner) *Challenge {
	challenge := Challenge{}

	var folds, punches int

	if !scanner.Scan() { return nil }
	_, err := fmt.Sscanf(scanner.Text(), "%d %d %d %d", &challenge.width, &challenge.height, &folds, &punches)
	if err != nil { return nil }

	challenge.folds = make([]string, folds)
	challenge.punches = make([]Point, punches)

	for i := 0; i < folds; i++ {
		if !scanner.Scan() { return nil }
		_, err := fmt.Sscanf(scanner.Text(), "%s", &challenge.folds[i])
		if err != nil { return nil }
	}

	for i := 0; i < punches; i++ {
		point := Point {}

		if !scanner.Scan() { return nil }
		_, err := fmt.Sscanf(scanner.Text(), "%d %d", &point.x, &point.y)
		if err != nil { return nil }

		challenge.punches[i] = point
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

func resolve(challenge Challenge) ChallengeResult {
	width := challenge.width
	height := challenge.height

	punches := make(Points, len(challenge.punches) * (2 << uint(len(challenge.folds) - 1)))
	copy(punches, challenge.punches)

	for i := 0; i < len(challenge.folds); i++  {
		for j := 0; j < len(challenge.punches) * (2 << uint(i) >> 1); j++  {
			switch challenge.folds[i] {
			case "L":
				// Projection
				punches[(len(challenge.punches) * (2 << uint(i) >> 1)) + j].x = (width - 1) - punches[j].x
				punches[(len(challenge.punches) * (2 << uint(i) >> 1)) + j].y = punches[j].y

				// Fixing
				punches[j].x += width

			case "R":
				// Projection
				punches[(len(challenge.punches) * (2 << uint(i) >> 1)) + j].x = ((width * 2) - 1) - punches[j].x
				punches[(len(challenge.punches) * (2 << uint(i) >> 1)) + j].y = punches[j].y

				// No fixing needed
			case "T":
				// Projection
				punches[(len(challenge.punches) * (2 << uint(i) >> 1)) + j].x = punches[j].x
				punches[(len(challenge.punches) * (2 << uint(i) >> 1)) + j].y = (height - 1) - punches[j].y

				// Fixing
				punches[j].y += height
			case "B":
				// Projection
				punches[(len(challenge.punches) * (2 << uint(i) >> 1)) + j].x = punches[j].x
				punches[(len(challenge.punches) * (2 << uint(i) >> 1)) + j].y = ((height * 2) - 1) - punches[j].y

				// No fixing needed
			}
		}

		// Resizing
		switch challenge.folds[i] {
		case "L":
			width = width * 2
		case "R":
			width = width * 2
		case "T":
			height = height * 2
		case "B":
			height = height * 2
		}
	}

	sort.Sort(punches)

	return ChallengeResult {
		n: challenge.n,
		punches: punches,
	}
}

func writeChallenge(result ChallengeResult) {
	fmt.Printf("Case #%d:", result.n)
	for i := 0; i < len(result.punches); i++ {
		fmt.Println()
		fmt.Printf("%d %d", result.punches[i].x, result.punches[i].y)
	}
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
