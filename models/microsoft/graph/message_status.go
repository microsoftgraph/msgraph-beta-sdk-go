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
    switch strings.ToUpper(v) {
        case "GETTINGSTATUS":
            return GETTINGSTATUS_MESSAGESTATUS, nil
        case "PENDING":
            return PENDING_MESSAGESTATUS, nil
        case "FAILED":
            return FAILED_MESSAGESTATUS, nil
        case "DELIVERED":
            return DELIVERED_MESSAGESTATUS, nil
        case "EXPANDED":
            return EXPANDED_MESSAGESTATUS, nil
        case "QUARANTINED":
            return QUARANTINED_MESSAGESTATUS, nil
        case "FILTEREDASSPAM":
            return FILTEREDASSPAM_MESSAGESTATUS, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_MESSAGESTATUS, nil
    }
    return 0, errors.New("Unknown MessageStatus value: " + v)
}
func SerializeMessageStatus(values []MessageStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
