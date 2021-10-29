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
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            return UNKNOWN_TLPLEVEL, nil
        case "WHITE":
            return WHITE_TLPLEVEL, nil
        case "GREEN":
            return GREEN_TLPLEVEL, nil
        case "AMBER":
            return AMBER_TLPLEVEL, nil
        case "RED":
            return RED_TLPLEVEL, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_TLPLEVEL, nil
    }
    return 0, errors.New("Unknown TlpLevel value: " + v)
}
func SerializeTlpLevel(values []TlpLevel) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
