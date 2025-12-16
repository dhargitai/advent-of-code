package main

import (
	"cmp"
	"fmt"
	"os"

	"math"
	"slices"
	"strconv"
	"strings"
)

type Point struct {
	x, y, z                       int
	distanceOrigo, distanceMaxXYZ float64
}

func (p *Point) toString() string {
	return "(" + strconv.Itoa((*p).x) + " " + strconv.Itoa((*p).y) + " " + strconv.Itoa((*p).z) + ")"
}

type PointPair struct {
	pointA, pointB Point
	distance       float64
}

func main() {
	file := "08.input.txt"
	if len(os.Args) > 1 {
		file = os.Args[1]
	}

	data, err := os.ReadFile(file)
	if err != nil {
		fmt.Println("Error reading file: ", err)
		return
	}

	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	coordinates := make([]Point, len(lines))
	maxX, maxY, maxZ := 0, 0, 0

	for i, lineString := range lines {
		line := strings.Split(lineString, ",")
		x, err := strconv.Atoi(line[0])
		if err != nil {
			fmt.Println("Error converting string to int", err)
			return
		}
		if maxX < x {
			maxX = x
		}

		y, err := strconv.Atoi(line[1])
		if err != nil {
			fmt.Println("Error converting string to int", err)
			return
		}
		if maxY < y {
			maxY = y
		}

		z, err := strconv.Atoi(line[2])
		if err != nil {
			fmt.Println("Error converting string to int", err)
			return
		}
		if maxZ < z {
			maxZ = z
		}

		distanceOrigo := math.Sqrt(math.Pow(float64(x), 2) + math.Pow(float64(y), 2) + math.Pow(float64(z), 2))

		coordinates[i] = Point{x, y, z, distanceOrigo, 0}
	}

	slices.SortFunc(coordinates, func(a, b Point) int {
		return cmp.Compare(a.distanceOrigo, b.distanceOrigo)
		// return cmp.Compare(a.x+a.y+a.z, b.x+b.y+b.z)
	})

	// distances := [][]float64{}
	pointPairs := []PointPair{}

	for i, coordinate := range coordinates {
		// fmt.Println(coordinate, ": ", coordinate.x+coordinate.y+coordinate.z)
		// distanceMaxXYZ := math.Sqrt(math.Pow((float64(coordinate.x)-float64(maxX)), 2) + math.Pow((float64(coordinate.y)-float64(maxY)), 2) + math.Pow((float64(coordinate.z)-float64(maxZ)), 2))

		if i > 0 {
			// distancesI := make([]float64, i)
			for j := 0; j < i; j++ {
				// distancesI[j] = math.Sqrt(math.Pow((float64(coordinate.x)-float64(coordinates[j].x)), 2) + math.Pow((float64(coordinate.y)-float64(coordinates[j].y)), 2) + math.Pow((float64(coordinate.z)-float64(coordinates[j].z)), 2))
				pointPairs = append(pointPairs, PointPair{
					pointA:   coordinate,
					pointB:   coordinates[j],
					distance: math.Sqrt(math.Pow((float64(coordinate.x)-float64(coordinates[j].x)), 2) + math.Pow((float64(coordinate.y)-float64(coordinates[j].y)), 2) + math.Pow((float64(coordinate.z)-float64(coordinates[j].z)), 2)),
				})
			}
			// distances = append(distances, distancesI)
		}

		// fmt.Println(coordinate, "-> dOrigo: \t", coordinate.distanceOrigo)
	}
	// fmt.Println(maxX, maxY, maxZ)

	slices.SortFunc(pointPairs, func(a, b PointPair) int {
		return cmp.Compare(a.distance, b.distance)
	})

	// fmt.Println("size: ", len(pointPairs))

	circuits := []*[]Point{}
	pointsWithCircuits := map[string]*[]Point{}

	for x, pointPair := range pointPairs {
		// fmt.Printf("%v - %v --- %v\n", pointPair.pointA.toString(), pointPair.pointB.toString(), pointPair.distance)
		if x > 999 {
			break
		}

		circuitOfA, pointAisOnCircuit := pointsWithCircuits[pointPair.pointA.toString()]
		circuitOfB, pointBisOnCircuit := pointsWithCircuits[pointPair.pointB.toString()]

		if !pointAisOnCircuit && !pointBisOnCircuit {
			// fmt.Printf("\tneither of these are on circuits, so we connect them and add to the circuits list...\n")
			circuits = append(circuits, &[]Point{pointPair.pointA, pointPair.pointB})
			pointsWithCircuits[pointPair.pointA.toString()] = circuits[len(circuits)-1]
			pointsWithCircuits[pointPair.pointB.toString()] = circuits[len(circuits)-1]
			// fmt.Printf("\tpointsWithCircuits[%v] after: %v\n", pointPair.pointA.toString(), pointsWithCircuits[pointPair.pointA.toString()])
			// fmt.Printf("\tpointsWithCircuits[%v] after: %v\n\n", pointPair.pointB.toString(), pointsWithCircuits[pointPair.pointB.toString()])
		} else if pointAisOnCircuit && !pointBisOnCircuit {
			// fmt.Printf("\t%v is on a circuit, so we add %v to that (1):\n", pointPair.pointA.toString(), pointPair.pointB.toString())
			// fmt.Printf("\tcircuit before: %v\n", (*circuitOfA))
			(*circuitOfA) = append((*circuitOfA), pointPair.pointB)
			// fmt.Printf("\tcircuit after: %v\n", (*circuitOfA))
			// fmt.Printf("\tpointsWithCircuits[%v] before: %v\n", pointPair.pointB.toString(), pointsWithCircuits[pointPair.pointB.toString()])
			pointsWithCircuits[pointPair.pointB.toString()] = circuitOfA
			// fmt.Printf("\tpointsWithCircuits[%v] after: %v\n\n", pointPair.pointB.toString(), pointsWithCircuits[pointPair.pointB.toString()])
		} else if pointBisOnCircuit && !pointAisOnCircuit {
			// fmt.Printf("\t%v is on a circuit, so we add %v to that (2):\n", pointPair.pointB.toString(), pointPair.pointA.toString())
			// fmt.Printf("\tcircuit before: %v\n", (*circuitOfB))
			(*circuitOfB) = append((*circuitOfB), pointPair.pointA)
			// fmt.Printf("\tcircuit after: %v\n", (*circuitOfB))
			// fmt.Printf("\tpointsWithCircuits[%v] before: %v\n", pointPair.pointA.toString(), pointsWithCircuits[pointPair.pointA.toString()])
			pointsWithCircuits[pointPair.pointA.toString()] = circuitOfB
			// fmt.Printf("\tpointsWithCircuits[%v] after: %v\n\n", pointPair.pointA.toString(), pointsWithCircuits[pointPair.pointA.toString()])
		} else if pointAisOnCircuit && pointBisOnCircuit && pointsWithCircuits[pointPair.pointA.toString()] != pointsWithCircuits[pointPair.pointB.toString()] {
			// fmt.Printf("\tboth %v and %v are on circuits\n", pointPair.pointA.toString(), pointPair.pointB.toString())
			// fmt.Printf("\tcircuit A before: %v\n", (*circuitOfA))
			// fmt.Printf("\tcircuit B: %v\n", (*circuitOfB))
			for _, point := range *circuitOfB {
				(*circuitOfA) = append((*circuitOfA), point)
				pointsWithCircuits[point.toString()] = circuitOfA
			}
			indexOfB := slices.Index(circuits, circuitOfB)
			circuits = append(circuits[:indexOfB], circuits[indexOfB+1:]...)
			// fmt.Printf("\tcircuit A after: %v\n", (*circuitOfA))
		}
	}

	// fmt.Println("circuits: ", circuits)
	slices.SortFunc(circuits, func(a, b *[]Point) int {
		return cmp.Compare(len((*b)), len((*a)))
	})

	result := 1
	for _, circuit := range circuits[:3] {
		// for _, circuit := range circuits {
		// fmt.Println(circuit)
		result *= len(*circuit)
	}
	fmt.Println("Pointpairs: ", len(pointPairs))
	fmt.Println("Result: ", result)
}
