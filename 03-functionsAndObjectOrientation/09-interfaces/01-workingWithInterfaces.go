package _010309_interfaces_

import "fmt"

type Food interface {
	Nutrition() string
	FoodType() string
}

type Apple struct {
}

// Empty structs are perfrectly legal.

func (c Apple) Nutrition() string {
	return "Apples are carb heavy!"
}

func (c Apple) FoodType() string {
	return "Apples are fruit."
}

type Celery struct {
}

func (c Celery) Nutrition() string {
	return "Celery has zero everything!"
}

func (c Celery) FoodType() string {
	return "Celery is a vegetable."
}

func WorkingWithInterfaces() {
	fmt.Println("Running `WorkingWithInterfaces`...")
	foods := []Food{Apple{}, Celery{}}
	for _, f := range foods {
		fmt.Println(f.Nutrition(), f.FoodType())
	}
}
