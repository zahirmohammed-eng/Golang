package pkg

// purpose: The pkg/ directory contains public utility code meant to be shared across projects or exported to external developers


func About() string{
	return "About the employee"
}

func Hike(hike float64) float64{
	return hike*0.5

}