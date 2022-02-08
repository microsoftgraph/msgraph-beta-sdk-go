package graph
import (
    "strings"
    "errors"
)
// 
type LostModeState int

const (
    DISABLED_LOSTMODESTATE LostModeState = iota
    ENABLED_LOSTMODESTATE
)

func (i LostModeState) String() string {
    return []string{"DISABLED", "ENABLED"}[i]
}
func ParseLostModeState(v string) (interface{}, error) {
    result := DISABLED_LOSTMODESTATE
    switch strings.ToUpper(v) {
        case "DISABLED":
            result = DISABLED_LOSTMODESTATE
        case "ENABLED":
            result = ENABLED_LOSTMODESTATE
        default:
            return 0, errors.New("Unknown LostModeState value: " + v)
    }
    return &result, nil
}
func SerializeLostModeState(values []LostModeState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
