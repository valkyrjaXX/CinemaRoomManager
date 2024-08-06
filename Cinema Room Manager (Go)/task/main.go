package main

import (
	"fmt"
)

func main() {
	fmt.Println("Enter the number of rows:")
	var numOfRows int
	fmt.Scan(&numOfRows)

	fmt.Println("Enter the number of seats in each row:")
	var numOfSeats int
	fmt.Scan(&numOfSeats)

	cinemaRoom := NewCinemaRoom(numOfRows, numOfSeats)

	options := []struct {
		Description string
		Handler     func() int
	}{
		{
			Description: "Exit",
			Handler:     func() int { return 0 },
		},
		{
			Description: "Show the seats",
			Handler: func() int {
				cinemaRoom.PrintAvailability()
				return 1
			},
		},
		{
			Description: "Buy a ticket",
			Handler: func() int {
				for {
					fmt.Println("Enter a row number:")
					var rowNumber int
					fmt.Scan(&rowNumber)

					fmt.Println("Enter a seat number in that row:")
					var seatNumber int
					fmt.Scan(&seatNumber)

					ticket, err := cinemaRoom.BookTicket(rowNumber, seatNumber)
					if err != nil {
						fmt.Println(err.Error())
						continue
					}

					fmt.Println(fmt.Sprintf("Ticket price: $%.0f", ticket))
					return 1
				}
			},
		},
		{
			Description: "Statistics",
			Handler: func() int {
				boughtSeats, currentIncome, totalIncome := cinemaRoom.CalculateStatistics()
				fmt.Println(fmt.Sprintf("Number of purchased tickets: %d", boughtSeats))
				fmt.Println(fmt.Sprintf("Percentage: %.2f%%", (float32(boughtSeats*100))/float32(cinemaRoom.totalSeats)))
				fmt.Println(fmt.Sprintf("Current income: $%.0f", currentIncome))
				fmt.Println(fmt.Sprintf("Total income: $%.0f", totalIncome))
				return 1
			},
		},
	}

	for {
		printOptions(options)

		var optionNum int
		fmt.Scan(&optionNum)

		result := options[optionNum].Handler()
		if result == 0 {
			return
		}
	}
}

func printOptions(options []struct {
	Description string
	Handler     func() int
}) {
	for idx, option := range options {
		if idx == 0 {
			defer fmt.Println(fmt.Sprintf("%d.  %s", idx, option.Description))
		} else {
			fmt.Println(fmt.Sprintf("%d.  %s", idx, option.Description))
		}
	}
}

//  cinema

type CinemaRoom struct {
	numOfRows  int
	numOfSeats int
	totalSeats int
	seats      [][]*Seat
}

func (r *CinemaRoom) PrintAvailability() {
	fmt.Println("Cinema:")
	fmt.Print(" ")
	for seat := 1; seat <= r.numOfSeats; seat++ {
		fmt.Print(fmt.Sprintf(" %d", seat))
	}
	fmt.Println()

	for row := 0; row < r.numOfRows; row++ {
		fmt.Print(row + 1)
		for seat := 0; seat < r.numOfSeats; seat++ {
			if r.seats[row][seat].Booked {
				fmt.Print(fmt.Sprintf(" B"))
			} else {
				fmt.Print(fmt.Sprintf(" S"))
			}
		}
		fmt.Println()
	}
}

func (r *CinemaRoom) BookTicket(rowNum, seatNum int) (float32, error) {
	if rowNum < 0 || rowNum > r.numOfRows || seatNum < 0 || seatNum > r.numOfSeats {
		return 0, fmt.Errorf("Wrong input!")
	}

	if r.seats[rowNum-1][seatNum-1].Booked {
		return 0, fmt.Errorf("That ticket has already been purchased!")
	}

	r.seats[rowNum-1][seatNum-1].Booked = true
	return r.seats[rowNum-1][seatNum-1].Price, nil
}

func (r *CinemaRoom) CalculateStatistics() (int, float32, float32) {
	var boughtSeats int
	var currentIncome, totalIncome float32

	for row := 0; row < r.numOfRows; row++ {
		for seat := 0; seat < r.numOfSeats; seat++ {
			if r.seats[row][seat].Booked {
				boughtSeats++
				currentIncome += r.seats[row][seat].Price
			}
			totalIncome += r.seats[row][seat].Price
		}
	}

	return boughtSeats, currentIncome, totalIncome
}

func NewCinemaRoom(numOfRows, numOfSeats int) *CinemaRoom {
	totalSeats := numOfRows * numOfSeats
	if totalSeats <= 60 {
		roomSeats := make([][]*Seat, numOfRows)
		for row := 0; row < numOfRows; row++ {
			roomSeats[row] = make([]*Seat, numOfSeats)
			for seat := 0; seat < numOfSeats; seat++ {
				roomSeats[row][seat] = &Seat{Price: 10}
			}
		}

		return &CinemaRoom{
			numOfRows:  numOfRows,
			numOfSeats: numOfSeats,
			totalSeats: numOfRows * numOfSeats,
			seats:      roomSeats,
		}
	}

	halfRows := numOfRows / 2
	roomSeats := make([][]*Seat, numOfRows)
	for row := 0; row < numOfRows; row++ {
		roomSeats[row] = make([]*Seat, numOfSeats)
		for seat := 0; seat < numOfSeats; seat++ {
			if row < halfRows {
				roomSeats[row][seat] = &Seat{Price: 10}
			} else {
				roomSeats[row][seat] = &Seat{Price: 8}
			}
		}
	}

	return &CinemaRoom{
		numOfRows:  numOfRows,
		numOfSeats: numOfSeats,
		totalSeats: numOfRows * numOfSeats,
		seats:      roomSeats,
	}
}

type Seat struct {
	Booked bool
	Price  float32
}
