package graph
import (
    "strings"
    "errors"
)
type LostModeState int

const (
    DISABLED_LOSTMODESTATE LostModeState = iota
    ENABLED_LOSTMODESTATE
)

func (i LostModeState) String() string {
    return []string{"DISABLED", "ENABLED"}[i]
}
func ParseLostModeState(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "DISABLED":
            return DISABLED_LOSTMODESTATE, nil
        case "ENABLED":
            return ENABLED_LOSTMODESTATE, nil
    }
    return 0, errors.New("Unknown LostModeState value: " + v)
}
func SerializeLostModeState(values []LostModeState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
