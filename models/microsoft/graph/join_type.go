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
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            return UNKNOWN_JOINTYPE, nil
        case "AZUREADJOINED":
            return AZUREADJOINED_JOINTYPE, nil
        case "AZUREADREGISTERED":
            return AZUREADREGISTERED_JOINTYPE, nil
        case "HYBRIDAZUREADJOINED":
            return HYBRIDAZUREADJOINED_JOINTYPE, nil
    }
    return 0, errors.New("Unknown JoinType value: " + v)
}
func SerializeJoinType(values []JoinType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
