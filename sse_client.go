package chatglm

import (
	"bufio"
	"io"
	"log"
	"strconv"
	"strings"
)

const fieldSeparator = ":"

// SSEClient represents a Server-Sent Events (SSE) client.
type SSEClient struct {
	eventSource io.Reader
	charEnc     string
}

// NewSSEClient creates a new SSEClient instance.
func NewSSEClient(eventSource io.Reader, charEnc string) *SSEClient {
	return &SSEClient{
		eventSource: eventSource,
		charEnc:     charEnc,
	}
}

// Event represents an SSE event.
type Event struct {
	ID    string
	Event string
	Data  string
	Retry int
	Meta  map[string]string
}

// ReadEvents reads SSE events from the event source and returns them as a channel.
func (c *SSEClient) ReadEvents() <-chan Event {
	eventChan := make(chan Event)
	scanner := bufio.NewScanner(c.eventSource)
	var eventData strings.Builder

	go func() {
		defer close(eventChan)

		for scanner.Scan() {
			line := scanner.Text()

			if len(line) == 0 {
				// Empty line indicates the end of an event.
				event := parseEvent(eventData.String())
				eventChan <- event
				eventData.Reset()
			} else {
				eventData.WriteString(line)
				eventData.WriteString("\n")
			}
		}

		if scanner.Err() != nil {
			log.Printf("Error reading SSE: %v", scanner.Err())
		}
	}()

	return eventChan
}

func parseEvent(data string) Event {
	event := Event{Meta: make(map[string]string)}
	scanner := bufio.NewScanner(strings.NewReader(data))
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, fieldSeparator, 2)
		if len(parts) != 2 {
			continue
		}
		field, value := strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1])

		switch field {
		case "id":
			event.ID = value
		case "event":
			event.Event = value
		case "data":
			event.Data += value + "\n"
		case "retry":
			// Parse retry field as an integer
			retry, err := strconv.Atoi(value)
			if err == nil {
				event.Retry = retry
			}
		default:
			// Store unrecognized fields in the meta map
			event.Meta[field] = value
		}
	}
	return event
}
