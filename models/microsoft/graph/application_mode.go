package graph
import (
    "strings"
    "errors"
)
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
    switch strings.ToUpper(v) {
        case "MANUAL":
            return MANUAL_APPLICATIONMODE, nil
        case "AUTOMATIC":
            return AUTOMATIC_APPLICATIONMODE, nil
        case "RECOMMENDED":
            return RECOMMENDED_APPLICATIONMODE, nil
    }
    return 0, errors.New("Unknown ApplicationMode value: " + v)
}
func SerializeApplicationMode(values []ApplicationMode) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
