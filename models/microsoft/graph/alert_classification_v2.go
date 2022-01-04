package graph
import (
    "strings"
    "errors"
)
// 
type AlertClassification_v2 int

const (
    UNKNOWN_ALERTCLASSIFICATION_V2 AlertClassification_v2 = iota
    FALSEPOSITIVE_ALERTCLASSIFICATION_V2
    TRUEPOSITIVE_ALERTCLASSIFICATION_V2
    BENIGNPOSITIVE_ALERTCLASSIFICATION_V2
    UNKNOWNFUTUREVALUE_ALERTCLASSIFICATION_V2
)

func (i AlertClassification_v2) String() string {
    return []string{"UNKNOWN", "FALSEPOSITIVE", "TRUEPOSITIVE", "BENIGNPOSITIVE", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseAlertClassification_v2(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            return UNKNOWN_ALERTCLASSIFICATION_V2, nil
        case "FALSEPOSITIVE":
            return FALSEPOSITIVE_ALERTCLASSIFICATION_V2, nil
        case "TRUEPOSITIVE":
            return TRUEPOSITIVE_ALERTCLASSIFICATION_V2, nil
        case "BENIGNPOSITIVE":
            return BENIGNPOSITIVE_ALERTCLASSIFICATION_V2, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_ALERTCLASSIFICATION_V2, nil
    }
    return 0, errors.New("Unknown AlertClassification_v2 value: " + v)
}
func SerializeAlertClassification_v2(values []AlertClassification_v2) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
