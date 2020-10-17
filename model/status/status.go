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

func (s Status) status() (Status, error) {
	if s >= InTransit && s <= Discontinued {
		return s, nil
	}
	return 0, fmt.Errorf("status undefined")
}

func (s Status) String() string {
	switch s {
	case InTransit:
		return "in transit"
	case InStock:
		return "in stock"
	case Sold:
		return "sold"
	case Discontinued:
		return "discontinued"
	default:
		return fmt.Sprintf("status %d undefined", s)
	}
}

func convertToStatus(s string) (Status, error) {
	switch s {
	case "in transit":
		return InTransit, nil
	case "in stock":
		return InStock, nil
	case "sold":
		return Sold, nil
	case "discontinued":
		return Discontinued, nil
	default:
		return 0, fmt.Errorf("wrong status")
	}
}

// MarshalJSON marshals the enum as a quoted json string
func (s Status) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	status, err := s.status()
	if err != nil {
		return []byte{}, fmt.Errorf("status %s %s", s, err)
	}
	buffer.WriteString(status.String())
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
	status, err := convertToStatus(string(j))
	if err != nil {
		return fmt.Errorf("status %s  %s", j, err)
	}
	*s = status
	return nil
}
