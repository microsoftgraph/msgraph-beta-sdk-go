package ediscovery
import (
    "strings"
    "errors"
)
// 
type LegalHoldStatus int

const (
    PENDING_LEGALHOLDSTATUS LegalHoldStatus = iota
    ERROR_LEGALHOLDSTATUS
    SUCCESS_LEGALHOLDSTATUS
    UNKNOWNFUTUREVALUE_LEGALHOLDSTATUS
)

func (i LegalHoldStatus) String() string {
    return []string{"PENDING", "ERROR", "SUCCESS", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseLegalHoldStatus(v string) (interface{}, error) {
    result := PENDING_LEGALHOLDSTATUS
    switch strings.ToUpper(v) {
        case "PENDING":
            result = PENDING_LEGALHOLDSTATUS
        case "ERROR":
            result = ERROR_LEGALHOLDSTATUS
        case "SUCCESS":
            result = SUCCESS_LEGALHOLDSTATUS
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_LEGALHOLDSTATUS
        default:
            return 0, errors.New("Unknown LegalHoldStatus value: " + v)
    }
    return &result, nil
}
func SerializeLegalHoldStatus(values []LegalHoldStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
