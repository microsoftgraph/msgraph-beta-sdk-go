package security
import (
    "strings"
    "errors"
)
// 
type AlertClassification int

const (
    UNKNOWN_ALERTCLASSIFICATION AlertClassification = iota
    FALSEPOSITIVE_ALERTCLASSIFICATION
    TRUEPOSITIVE_ALERTCLASSIFICATION
    INFORMATIONALEXPECTEDACTIVITY_ALERTCLASSIFICATION
    UNKNOWNFUTUREVALUE_ALERTCLASSIFICATION
)

func (i AlertClassification) String() string {
    return []string{"UNKNOWN", "FALSEPOSITIVE", "TRUEPOSITIVE", "INFORMATIONALEXPECTEDACTIVITY", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseAlertClassification(v string) (interface{}, error) {
    result := UNKNOWN_ALERTCLASSIFICATION
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            result = UNKNOWN_ALERTCLASSIFICATION
        case "FALSEPOSITIVE":
            result = FALSEPOSITIVE_ALERTCLASSIFICATION
        case "TRUEPOSITIVE":
            result = TRUEPOSITIVE_ALERTCLASSIFICATION
        case "INFORMATIONALEXPECTEDACTIVITY":
            result = INFORMATIONALEXPECTEDACTIVITY_ALERTCLASSIFICATION
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_ALERTCLASSIFICATION
        default:
            return 0, errors.New("Unknown AlertClassification value: " + v)
    }
    return &result, nil
}
func SerializeAlertClassification(values []AlertClassification) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
