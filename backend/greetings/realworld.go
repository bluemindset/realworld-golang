package greetings

import (
	"errors"
	"fmt"
	"math"
)

func powerUp(passfunction func(float64, float64) float64, helloworld string, p1 float64, p2 float64) (float64, error) {

	fmt.Println(helloworld)

	return passfunction(p1, p2), nil
}

func Hello(name string) (string, error) {

	var hello = func(p1 float64, p2 float64) float64 {
		return math.Sqrt(p1*p1 + p2*p2)
	}
	res, err := powerUp(hello, "yes", 5, 5)

	if err != nil {
		return "", err

	}
	fmt.Printf("PowerUp Result:  %f\n", res)

	if name == "" {
		return "", errors.New("empty name")
	}

	message := fmt.Sprintf("Hi, %v. Welcome!", name)

	return message, nil

}
