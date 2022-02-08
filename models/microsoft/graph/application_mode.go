package graph
import (
    "strings"
    "errors"
)
// 
type ApplicationMode int

const (
    MANUAL_APPLICATIONMODE ApplicationMode = iota
    AUTOMATIC_APPLICATIONMODE
    RECOMMENDED_APPLICATIONMODE
)

func (i ApplicationMode) String() string {
    return []string{"MANUAL", "AUTOMATIC", "RECOMMENDED"}[i]
}
func ParseApplicationMode(v string) (interface{}, error) {
    result := MANUAL_APPLICATIONMODE
    switch strings.ToUpper(v) {
        case "MANUAL":
            result = MANUAL_APPLICATIONMODE
        case "AUTOMATIC":
            result = AUTOMATIC_APPLICATIONMODE
        case "RECOMMENDED":
            result = RECOMMENDED_APPLICATIONMODE
        default:
            return 0, errors.New("Unknown ApplicationMode value: " + v)
    }
    return &result, nil
}
func SerializeApplicationMode(values []ApplicationMode) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
