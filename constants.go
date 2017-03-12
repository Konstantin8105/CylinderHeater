package main

import (
	"fmt"

	"github.com/Konstantin8105/CylinderHeater/section"
)

var modulusOfElasticityE float32 = 2.05e11 //Pa

func main() {
	fmt.Println("Hello")
	s := section.Ibeam{H: 0.4, B: 0.4, Tf: 0.030, Tw: 0.02}
	fmt.Println("Section = ", s, "\t Jx = ", s.Jx())
}
