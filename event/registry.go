package event

var counter uint16

func Register() Id {
	counter++
	return Id(counter)
}
