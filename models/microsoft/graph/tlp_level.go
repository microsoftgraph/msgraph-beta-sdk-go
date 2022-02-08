package graph
import (
    "strings"
    "errors"
)
// 
type TlpLevel int

const (
    UNKNOWN_TLPLEVEL TlpLevel = iota
    WHITE_TLPLEVEL
    GREEN_TLPLEVEL
    AMBER_TLPLEVEL
    RED_TLPLEVEL
    UNKNOWNFUTUREVALUE_TLPLEVEL
)

func (i TlpLevel) String() string {
    return []string{"UNKNOWN", "WHITE", "GREEN", "AMBER", "RED", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseTlpLevel(v string) (interface{}, error) {
    result := UNKNOWN_TLPLEVEL
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            result = UNKNOWN_TLPLEVEL
        case "WHITE":
            result = WHITE_TLPLEVEL
        case "GREEN":
            result = GREEN_TLPLEVEL
        case "AMBER":
            result = AMBER_TLPLEVEL
        case "RED":
            result = RED_TLPLEVEL
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_TLPLEVEL
        default:
            return 0, errors.New("Unknown TlpLevel value: " + v)
    }
    return &result, nil
}
func SerializeTlpLevel(values []TlpLevel) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
