package graph
import (
    "strings"
    "errors"
)
// 
type MessageStatus int

const (
    GETTINGSTATUS_MESSAGESTATUS MessageStatus = iota
    PENDING_MESSAGESTATUS
    FAILED_MESSAGESTATUS
    DELIVERED_MESSAGESTATUS
    EXPANDED_MESSAGESTATUS
    QUARANTINED_MESSAGESTATUS
    FILTEREDASSPAM_MESSAGESTATUS
    UNKNOWNFUTUREVALUE_MESSAGESTATUS
)

func (i MessageStatus) String() string {
    return []string{"GETTINGSTATUS", "PENDING", "FAILED", "DELIVERED", "EXPANDED", "QUARANTINED", "FILTEREDASSPAM", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseMessageStatus(v string) (interface{}, error) {
    result := GETTINGSTATUS_MESSAGESTATUS
    switch strings.ToUpper(v) {
        case "GETTINGSTATUS":
            result = GETTINGSTATUS_MESSAGESTATUS
        case "PENDING":
            result = PENDING_MESSAGESTATUS
        case "FAILED":
            result = FAILED_MESSAGESTATUS
        case "DELIVERED":
            result = DELIVERED_MESSAGESTATUS
        case "EXPANDED":
            result = EXPANDED_MESSAGESTATUS
        case "QUARANTINED":
            result = QUARANTINED_MESSAGESTATUS
        case "FILTEREDASSPAM":
            result = FILTEREDASSPAM_MESSAGESTATUS
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_MESSAGESTATUS
        default:
            return 0, errors.New("Unknown MessageStatus value: " + v)
    }
    return &result, nil
}
func SerializeMessageStatus(values []MessageStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
