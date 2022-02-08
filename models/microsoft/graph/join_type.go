package graph
import (
    "strings"
    "errors"
)
// 
type JoinType int

const (
    UNKNOWN_JOINTYPE JoinType = iota
    AZUREADJOINED_JOINTYPE
    AZUREADREGISTERED_JOINTYPE
    HYBRIDAZUREADJOINED_JOINTYPE
)

func (i JoinType) String() string {
    return []string{"UNKNOWN", "AZUREADJOINED", "AZUREADREGISTERED", "HYBRIDAZUREADJOINED"}[i]
}
func ParseJoinType(v string) (interface{}, error) {
    result := UNKNOWN_JOINTYPE
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            result = UNKNOWN_JOINTYPE
        case "AZUREADJOINED":
            result = AZUREADJOINED_JOINTYPE
        case "AZUREADREGISTERED":
            result = AZUREADREGISTERED_JOINTYPE
        case "HYBRIDAZUREADJOINED":
            result = HYBRIDAZUREADJOINED_JOINTYPE
        default:
            return 0, errors.New("Unknown JoinType value: " + v)
    }
    return &result, nil
}
func SerializeJoinType(values []JoinType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
