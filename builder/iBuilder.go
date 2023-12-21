package builder

// interface
type IBuilder interface {
	setWindowType()
	setDoorType()
	setNumFloor()
	getHouse() House
}

func getBuilder(builderType string) IBuilder {
	if builderType == "normal" {
		return newNormalBuilder()
	} else if builderType == "igloo" {
		return newIglooBuilder()
	}

	return nil
}
