package main

//==============================================================================
// 									GAME LIVE
//==============================================================================

import "fmt"
import "time"
import "math/rand"

const field_size = 30

type field_t struct {
	_old_cells [field_size][field_size]cell_t
	_cells [field_size][field_size]cell_t
}

func howMuchLivingNeighbor(this *field_t, x int, y int) int {
	var res int = 0
	if x > 0 {
		res += this._old_cells[x - 1][y]._alive
	}
	if x < field_size - 1 {
		res += this._old_cells[x + 1][y]._alive
	}
	if y > 0 {
		res += this._old_cells[x][y - 1]._alive
	}
	if y < field_size - 1 {
		res += this._old_cells[x][y + 1]._alive
	}

	if x > 0 && y > 0 {
		res += this._old_cells[x - 1][y - 1]._alive
	}

	if x > 0 && y < field_size - 1 {
		res += this._old_cells[x - 1][y + 1]._alive
	}

	if x < field_size - 1 && y > 0 {
		res += this._old_cells[x + 1][y - 1]._alive
	}

	if x < field_size - 1 && y < field_size - 1 {
		res += this._old_cells[x + 1][y + 1]._alive
	}

	return res;
}

func field_step(this *field_t) {
	for x := 0; x < field_size; x++ {
		for y := 0; y < field_size; y++ {
			this._old_cells[x][y] = this._cells[x][y]
		}
	}

	for x := 0; x < field_size; x++ {
		for y := 0; y < field_size; y++ {
			this._cells[x][y]._alive = cell_step(&this._old_cells[x][y])
		}
	}
}

func field_draw (this *field_t) {
	for x := 0; x < field_size; x++ {
		for y := 0; y < field_size; y++ {
			if this._cells[x][y]._alive == 1 {
				fmt.Printf("#.")
			} else {
				fmt.Printf("..")
			}
		}

		fmt.Printf("\n")
	}
}

type cell_t struct {
	_field *field_t
	_x int
	_y int
	_alive int
}

func cell_step(this *cell_t) int {
	var living_neighbors = howMuchLivingNeighbor(this._field, this._x, this._y)
	
	if this._alive == 1 {
		if living_neighbors == 3 || living_neighbors == 2 {
			this._alive = 1
		} else {
			this._alive = 0
		}
	} else {
		if living_neighbors == 3 {
			this._alive = 1
		} else {
			this._alive = 0
		}
	}

	return this._alive
}


func main() {
	rand.Seed(time.Now().Unix())

	field := field_t{}
	for x := 0; x < field_size; x++ {
		for y := 0; y < field_size; y++ {
			var alive int = rand.Intn(10)
			if alive > 0 {
				alive = 1
			}

			field._cells[x][y] = cell_t{_field: &field, _x: x, _y: y, _alive: alive}
		}
	}

	for {
		field_step(&field)
		time.Sleep(time.Second)
		fmt.Println("===========================================================");
		field_draw(&field)
	}
}
