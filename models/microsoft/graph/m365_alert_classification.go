package graph
import (
    "strings"
    "errors"
)
// 
type M365AlertClassification int

const (
    UNKNOWN_M365ALERTCLASSIFICATION M365AlertClassification = iota
    FALSEPOSITIVE_M365ALERTCLASSIFICATION
    TRUEPOSITIVE_M365ALERTCLASSIFICATION
    BENIGNPOSITIVE_M365ALERTCLASSIFICATION
    UNKNOWNFUTUREVALUE_M365ALERTCLASSIFICATION
)

func (i M365AlertClassification) String() string {
    return []string{"UNKNOWN", "FALSEPOSITIVE", "TRUEPOSITIVE", "BENIGNPOSITIVE", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseM365AlertClassification(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            return UNKNOWN_M365ALERTCLASSIFICATION, nil
        case "FALSEPOSITIVE":
            return FALSEPOSITIVE_M365ALERTCLASSIFICATION, nil
        case "TRUEPOSITIVE":
            return TRUEPOSITIVE_M365ALERTCLASSIFICATION, nil
        case "BENIGNPOSITIVE":
            return BENIGNPOSITIVE_M365ALERTCLASSIFICATION, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_M365ALERTCLASSIFICATION, nil
    }
    return 0, errors.New("Unknown M365AlertClassification value: " + v)
}
func SerializeM365AlertClassification(values []M365AlertClassification) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
