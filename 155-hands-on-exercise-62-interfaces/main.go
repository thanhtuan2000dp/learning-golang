package main

import (
	"fmt"
	"math"
)

type square struct {
	length float64
	width  float64
}

type circle struct {
	radius float64
}

func (s square) Area() float64 {
	return s.length * s.width
}

func (c circle) Area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

type shape interface {
	Area() float64
}

func info(s shape) float64 {
	return s.Area()
}

func main() {
	c1 := circle{
		radius: 2,
	}

	s1 := square{
		length: 2,
		width:  3,
	}

	fmt.Println("Area of Circle:", info(c1))
	fmt.Println("Area of Square:", info(s1))
}

/*
create a type SQUARE
length float64
width float64
create a type CIRCLE
radius float64
attach a method to each that calculates AREA and returns it
circle area= π r 2
math.Pi
math.Pow
square area = L * W
create a type SHAPE that defines an interface as anything that has the AREA method
create a func INFO which takes type shape and then prints the area
create a value of type square
create a value of type circle
use func info to print the area of square
use func info to print the area of circle*/
