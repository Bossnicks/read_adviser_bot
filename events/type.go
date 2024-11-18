package events

type Fetcher interface {
	Fetch(limit int) ([]Event, error)
}

type Processor interface {
	Process(e Event) error
}

type Type interface

const (
	Unknown Type = iota
	Message
)

type Event struct {
	Type Type
	Text string
}