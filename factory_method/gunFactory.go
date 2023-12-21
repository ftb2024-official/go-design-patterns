package factory_method

import "fmt"

func getGun(gunType string) (IGun, error) {
	if gunType == "ak47" {
		return newAK47(), nil
	} else if gunType == "musket" {
		return newMusket(), nil
	}

	return nil, fmt.Errorf("Wrong gun type passed...")
}
