package utils

type Driver struct {
	DriverID string
	Name     string
	Lat      float64
	Long     float64
	Rating   int
	Status   string
}

func GenerateDriverData() *[]Driver {
	driverData := []Driver{
		{
			DriverID: "001",
			Name:     "Dhruv Rathi",
			Lat:      10.12345,
			Long:     20.67890,
			Rating:   5,
			Status:   "available",
		},
		{
			DriverID: "002",
			Name:     "Bob Stone",
			Lat:      15.23456,
			Long:     30.45678,
			Rating:   4,
			Status:   "busy",
		},
		{
			DriverID: "003",
			Name:     "Ashith",
			Lat:      20.34567,
			Long:     40.12345,
			Rating:   3,
			Status:   "available",
		},
		{
			DriverID: "004",
			Name:     "Bob Smith",
			Lat:      25.67890,
			Long:     50.12345,
			Rating:   2,
			Status:   "busy",
		},
		{
			DriverID: "005",
			Name:     "Adarsh S H",
			Lat:      30.23456,
			Long:     60.45678,
			Rating:   1,
			Status:   "available",
		},
		{
			DriverID: "006",
			Name:     "Sumanth",
			Lat:      35.34567,
			Long:     70.67890,
			Rating:   0,
			Status:   "available",
		},
		{
			DriverID: "007",
			Name:     "Jeethan",
			Lat:      40.12345,
			Long:     80.23456,
			Rating:   0,
			Status:   "available",
		},
		{
			DriverID: "008",
			Name:     "Lahari Priya N",
			Lat:      45.67890,
			Long:     90.34567,
			Rating:   0,
			Status:   "available",
		},
	}
	dataLoc := &driverData
	return dataLoc
}
