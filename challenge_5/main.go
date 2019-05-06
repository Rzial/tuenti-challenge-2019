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

	GLin = 2
	GCol = 4

	BLin = 3
	BCol = 4
)

var Charset = []string {
	"1234567890",
	"QWERTYUIOP",
	"ASDFGHJKL;",
	"ZXCVBNM,.-",
}

type Input struct {
	cases      int

	challenges []Challenge
}

type Challenge struct {
	n       int

	sender  string
	message string
}

type ChallengeResult struct {
	n                int

	decryptedMessage string
}

func readChallenge(n int, scanner *bufio.Scanner) *Challenge {
	challenge := Challenge{}

	if !scanner.Scan() { return nil }
	challenge.sender = scanner.Text()

	if !scanner.Scan() { return nil }
	challenge.message = scanner.Text()

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
	var senderLin, senderCol int
	switch challenge.sender {
	case "B":
		senderLin = BLin
		senderCol = BCol
	case "G":
		senderLin = GLin
		senderCol = GCol
	}

	var encodedLin, encodedCol int
	for i := 0; i < len(Charset); i++ {
		encodedLin = i
		encodedCol = strings.Index(Charset[i], string(challenge.message[len(challenge.message) - 1]))

		if encodedCol != -1 {
			break
		}
	}

	offsetLin := senderLin - encodedLin
	offsetCol := senderCol - encodedCol

	decryptedMessage := []uint8(challenge.message)
	for i := 0; i < len(challenge.message); i++ {
		if challenge.message[i] == ' ' {
			continue
		}

		for j := 0; j < len(Charset); j++ {
			encodedLin = j
			encodedCol = strings.Index(Charset[j], string(challenge.message[i]))

			if encodedCol != -1 {
				break
			}
		}

		decodedLin := encodedLin + offsetLin
		decodedCol := encodedCol + offsetCol

		for {
			if decodedLin > 0 { break}
			decodedLin += len(Charset)
		}

		decodedLin %= len(Charset)

		for {
			if decodedCol > 0 { break }
			decodedCol += len(Charset[decodedLin])
		}

		decodedCol %= len(Charset[decodedLin])

		decryptedMessage[i] = Charset[decodedLin][decodedCol]
	}


	return ChallengeResult {
		n: challenge.n,

		decryptedMessage: string(decryptedMessage),
	}
}

func writeChallenge(result ChallengeResult) {
	fmt.Printf("Case #%d: %s", result.n, result.decryptedMessage)
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
