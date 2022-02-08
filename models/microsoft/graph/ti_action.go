package graph
import (
    "strings"
    "errors"
)
// 
type TiAction int

const (
    UNKNOWN_TIACTION TiAction = iota
    ALLOW_TIACTION
    BLOCK_TIACTION
    ALERT_TIACTION
    UNKNOWNFUTUREVALUE_TIACTION
)

func (i TiAction) String() string {
    return []string{"UNKNOWN", "ALLOW", "BLOCK", "ALERT", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseTiAction(v string) (interface{}, error) {
    result := UNKNOWN_TIACTION
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            result = UNKNOWN_TIACTION
        case "ALLOW":
            result = ALLOW_TIACTION
        case "BLOCK":
            result = BLOCK_TIACTION
        case "ALERT":
            result = ALERT_TIACTION
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_TIACTION
        default:
            return 0, errors.New("Unknown TiAction value: " + v)
    }
    return &result, nil
}
func SerializeTiAction(values []TiAction) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
