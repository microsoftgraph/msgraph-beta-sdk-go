package graph
import (
    "strings"
    "errors"
)
type PrintColorConfiguration int

const (
    BLACKANDWHITE_PRINTCOLORCONFIGURATION PrintColorConfiguration = iota
    GRAYSCALE_PRINTCOLORCONFIGURATION
    COLOR_PRINTCOLORCONFIGURATION
    AUTO_PRINTCOLORCONFIGURATION
)

func (i PrintColorConfiguration) String() string {
    return []string{"BLACKANDWHITE", "GRAYSCALE", "COLOR", "AUTO"}[i]
}
func ParsePrintColorConfiguration(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "BLACKANDWHITE":
            return BLACKANDWHITE_PRINTCOLORCONFIGURATION, nil
        case "GRAYSCALE":
            return GRAYSCALE_PRINTCOLORCONFIGURATION, nil
        case "COLOR":
            return COLOR_PRINTCOLORCONFIGURATION, nil
        case "AUTO":
            return AUTO_PRINTCOLORCONFIGURATION, nil
    }
    return 0, errors.New("Unknown PrintColorConfiguration value: " + v)
}
func SerializePrintColorConfiguration(values []PrintColorConfiguration) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
