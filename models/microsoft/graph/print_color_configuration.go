package graph
import (
    "strings"
    "errors"
)
// 
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
    result := BLACKANDWHITE_PRINTCOLORCONFIGURATION
    switch strings.ToUpper(v) {
        case "BLACKANDWHITE":
            result = BLACKANDWHITE_PRINTCOLORCONFIGURATION
        case "GRAYSCALE":
            result = GRAYSCALE_PRINTCOLORCONFIGURATION
        case "COLOR":
            result = COLOR_PRINTCOLORCONFIGURATION
        case "AUTO":
            result = AUTO_PRINTCOLORCONFIGURATION
        default:
            return 0, errors.New("Unknown PrintColorConfiguration value: " + v)
    }
    return &result, nil
}
func SerializePrintColorConfiguration(values []PrintColorConfiguration) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
