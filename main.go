package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/google/uuid"
)

type Vec2 struct {
	X int
	Y int
}

type GridSection struct {
	ID           string
	ConnectedIDs []string
}

type Grid [][]GridSection

func main() {
	widthPtr := flag.Int("width", 4, "Width of the x axis")
	heightPtr := flag.Int("height", 4, "Height of the y axis")
	flag.Parse()

	file, err := os.OpenFile("servers.csv", os.O_RDWR|os.O_TRUNC, os.ModePerm)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	grid := GenerateGrid(widthPtr, heightPtr)
	grid = ConnectNodes(grid)

	writer := csv.NewWriter(file)
	// defer writer.Flush()

	writer.Write([]string{"ID", "ConnectedIDs"})

	for h := 0; h < *heightPtr; h++ {
		for w := 0; w < *widthPtr; w++ {
			toWrite := []string{
				grid[h][w].ID,
				strings.Join(grid[h][w].ConnectedIDs, ","),
			}

			if err := writer.Write(toWrite); err != nil {
				fmt.Println(err.Error())
			}
		}
	}

	writer.Flush()
}

func GenerateGrid(width, height *int) Grid {
	var grid Grid

	for h := 0; h < *height; h++ {
		row := []GridSection{}

		for w := 0; w < *width; w++ {
			gridSection := GridSection{
				ID:           uuid.NewString(),
				ConnectedIDs: []string{},
			}

			row = append(row, gridSection)
		}

		grid = append(grid, row)
	}

	return grid
}

func ConnectNodes(grid Grid) Grid {
	width := len(grid[0])
	height := len(grid)

	for h := 0; h < height; h++ {
		for w := 0; w < width; w++ {
			neighors := GetNeighbors(Vec2{X: w, Y: h}, width, height)

			for _, n := range neighors {
				id := grid[n.Y][n.X].ID
				grid[h][w].ConnectedIDs = append(grid[h][w].ConnectedIDs, id)
			}
		}

	}

	return grid
}

func GetNeighbors(pos Vec2, width, height int) []Vec2 {
	neighbors := []Vec2{}
	directions := []Vec2{
		{X: -1, Y: 0}, // Left
		{X: 1, Y: 0},  // Right
		{X: 0, Y: -1}, // Up
		{X: 0, Y: 1},  // Down
	}

	for _, dir := range directions {
		col := pos.X + dir.X
		row := pos.Y + dir.Y

		if col >= 0 && col < width && row >= 0 && row < height {
			neighbors = append(neighbors, Vec2{X: col, Y: row})
		}
	}

	return neighbors
}
