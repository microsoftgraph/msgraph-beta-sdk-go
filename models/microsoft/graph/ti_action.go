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
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            return UNKNOWN_TIACTION, nil
        case "ALLOW":
            return ALLOW_TIACTION, nil
        case "BLOCK":
            return BLOCK_TIACTION, nil
        case "ALERT":
            return ALERT_TIACTION, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_TIACTION, nil
    }
    return 0, errors.New("Unknown TiAction value: " + v)
}
func SerializeTiAction(values []TiAction) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
