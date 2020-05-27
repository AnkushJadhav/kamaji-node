package store

// Driver is an interface that every persistant storage adapter in the application must implement
type Driver interface {
	Connect() error
	Disconnect() error
}
