package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"sort"
	"strconv"
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

type Moon struct {
	moonDistance    float64
	initialPosition float64
	orbitalPeriod   float64

	unobtanium     int
}

type Challenge struct {
	n         int

	moons     []Moon

	capacity  int
	shipRange float64
}

type ChallengeResult struct {
	n      int

	result []int
}

func readChallenge(n int, scanner *bufio.Scanner) *Challenge {
	challenge := Challenge{}

	var moonLen int
	if !scanner.Scan() { return nil }
	_, err := fmt.Sscanf(scanner.Text(), "%d", &moonLen)
	if err != nil { return nil }

	challenge.moons = make([]Moon, moonLen)

	if !scanner.Scan() { return nil }
	distances := strings.Split(scanner.Text(), " ")

	for i, distance := range distances {
		moonDistance, _ := strconv.ParseFloat(distance, 64)
		challenge.moons[i].moonDistance = moonDistance
	}

	if !scanner.Scan() { return nil }
	initialPositions := strings.Split(scanner.Text(), " ")

	for i, position := range initialPositions {
		moonPosition, _ := strconv.ParseFloat(position, 64)
		challenge.moons[i].initialPosition = moonPosition
	}

	if !scanner.Scan() { return nil }
	orbitalPeriods := strings.Split(scanner.Text(), " ")

	for i, period := range orbitalPeriods {
		moonOrbitalPeriod, _ := strconv.ParseFloat(period, 64)
		challenge.moons[i].orbitalPeriod = moonOrbitalPeriod
	}

	if !scanner.Scan() { return nil }
	unobtainiums := strings.Split(scanner.Text(), " ")

	for i, unobtainium := range unobtainiums {
		moonUnobtainium, _ := strconv.Atoi(unobtainium)
		challenge.moons[i].unobtanium = moonUnobtainium
	}

	if !scanner.Scan() { return nil }
	_, err = fmt.Sscanf(scanner.Text(), "%d", &challenge.capacity)
	if err != nil { return nil }

	if !scanner.Scan() { return nil }
	_, err = fmt.Sscanf(scanner.Text(), "%f", &challenge.shipRange)
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

func polarDistance(p1 []float64, p2 []float64) float64 {
	r1Sqr := math.Pow(p1[0], 2)
	r2Sqr := math.Pow(p2[0], 2)
	rest := 2 * p1[0] * p2[0] * math.Cos(p1[1] - p2[1])

	return math.Sqrt(r1Sqr + r2Sqr - rest)
}

func resolve(challenge Challenge) ChallengeResult {
	dep := 1
	minCapacity := challenge.capacity
	minRange := challenge.shipRange
	bestPath := make([]int, 0)

	var simulateWarps func(origin []float64, capacity int, shipRange float64, moons []Moon, time int, path []int)
	simulateWarps = func(origin []float64, capacity int, shipRange float64, moons []Moon, time int, path []int) {
		if capacity < minCapacity || (capacity == minCapacity && shipRange > minRange) {
			bestPath = make([]int, len(path))
			copy(bestPath, path)
			minCapacity = capacity
			minRange = shipRange
		}

		for i, moon := range moons {
			restMoons := make([]Moon, len(moons))
			copy(restMoons, moons)
			restMoons = append(restMoons[:i], restMoons[i + 1:]...)

			angle := moon.initialPosition + (((1 / moon.orbitalPeriod) * 2 * math.Pi) * float64(time))
			for {
				if angle < 0 {
					angle += 2 * math.Pi
					break
				}
				angle -= 2 * math.Pi
			}
			polarCoords := []float64 {moon.moonDistance, angle}

			distance := polarDistance(origin, polarCoords)

			// If we can't back to the planet or we cant pick all the unobtanium we dont go to this planet
			if (shipRange - distance - moon.moonDistance) < 0 || capacity - moon.unobtanium < 0 {
				continue
			}

			angle = moon.initialPosition + (((1 / moon.orbitalPeriod) * 2 * math.Pi) * float64(time + 6))
			for {
				if angle < 0 {
					angle += 2 * math.Pi
					break
				}
				angle -= 2 * math.Pi
			}
			polarCoords = []float64 {moon.moonDistance, angle}

			actualPath := make([]int, len(path) + 1)
			copy(actualPath, path)
			actualPath[len(path)] = moon.unobtanium

			dep++
			simulateWarps(polarCoords, capacity - moon.unobtanium, shipRange - distance, restMoons, time + 6, actualPath)
			dep--
		}
	}

	simulateWarps([]float64 {0, 0}, challenge.capacity, float64(challenge.shipRange), challenge.moons, 0, []int{})
	sort.Ints(bestPath)

	return ChallengeResult {
		n: challenge.n,

		result: bestPath,
	}
}

func writeChallenge(result ChallengeResult) {
	resultString := ""
	if len(result.result) == 0 {
		resultString = "None"
	} else {
		resultString = strconv.Itoa(result.result[0])
		for i := 1; i < len(result.result); i++ {
			resultString += " " + strconv.Itoa(result.result[i])
		}
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
