package status

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type Status uint

const (
	InTransit Status = iota
	InStock
	Sold
	Discontinued
)

var byString = map[Status]string{
	InTransit:    "in transit",
	InStock:      "in stock",
	Sold:         "sold",
	Discontinued: "discontinued",
}

var toID = func(byStrings map[Status]string) map[string]*Status {
	byIDs := make(map[string]*Status, len(byString))
	for key, value := range byStrings {
		byIDs[value] = &key
	}
	return byIDs
}(byString)

func (s Status) String() string {
	return byString[s]
}

// MarshalJSON marshals the enum as a quoted json string
func (s Status) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	status, ok := byString[s]
	if !ok {
		return []byte{}, fmt.Errorf("status %s not allowed", s)
	}
	buffer.WriteString(status)
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

// UnmarshalJSON unmashals a quoted json string to the enum value
func (s *Status) UnmarshalJSON(b []byte) error {
	var j string
	err := json.Unmarshal(b, &j)
	if err != nil {
		return err
	}
	// Note that if the string cannot be found then it will be set to the zero value, 'Created' in this case.
	s, ok := toID[j]
	if !ok {
		return fmt.Errorf("status `%s` not allowed", j)
	}
	return nil
}
