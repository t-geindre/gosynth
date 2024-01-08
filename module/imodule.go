package module

type IModule interface {
	Write(port Port, value float64)
	Connect(srcPort Port, destModule IModule, destPort Port)
	Disconnect(srcPort Port, destModule IModule, destPort Port)
	Update()
	Dispose()
	// SendInput Thread-safe way to send a command to a module
	SendInput(port Port, value float64)
	// ReceiveInput Thread-safe way to read data sent to a module
	ReceiveInput(port Port) *float64
	// ReceiveOutput Thread-safe way to get data output from a module
	ReceiveOutput(port Port) *float64
}
