package utils

type Driver struct {
	DriverID string
	Name     string
	Lat      float64
	Long     float64
	Rating   int
}

func GenerateDriverData() []Driver {
	return []Driver{
		{
			DriverID: "001",
			Name:     "John Doe",
			Lat:      10.12345,
			Long:     20.67890,
			Rating:   5,
		},
		{
			DriverID: "002",
			Name:     "Bob Stone",
			Lat:      15.23456,
			Long:     30.45678,
			Rating:   4,
		},
		{
			DriverID: "003",
			Name:     "Harry Potter",
			Lat:      20.34567,
			Long:     40.12345,
			Rating:   3,
		},
		{
			DriverID: "004",
			Name:     "Bob Smith",
			Lat:      25.67890,
			Long:     50.12345,
			Rating:   2,
		},
		{
			DriverID: "005",
			Name:     "Alice Johnson",
			Lat:      30.23456,
			Long:     60.45678,
			Rating:   1,
		},
		{
			DriverID: "006",
			Name:     "Michael Brown",
			Lat:      35.34567,
			Long:     70.67890,
			Rating:   0,
		},
	}
}
