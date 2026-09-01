package testbed

type Notifier interface {
	Send(to, message string) error
}
