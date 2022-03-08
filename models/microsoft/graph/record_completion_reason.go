package graph
import (
    "strings"
    "errors"
)
// Provides operations to call the record method.
type RecordCompletionReason int

const (
    OPERATIONCANCELED_RECORDCOMPLETIONREASON RecordCompletionReason = iota
    STOPTONEDETECTED_RECORDCOMPLETIONREASON
    MAXRECORDDURATIONREACHED_RECORDCOMPLETIONREASON
    INITIALSILENCETIMEOUT_RECORDCOMPLETIONREASON
    MAXSILENCETIMEOUT_RECORDCOMPLETIONREASON
    PLAYPROMPTFAILED_RECORDCOMPLETIONREASON
    PLAYBEEPFAILED_RECORDCOMPLETIONREASON
    MEDIARECEIVETIMEOUT_RECORDCOMPLETIONREASON
    UNSPECIFIEDERROR_RECORDCOMPLETIONREASON
)

func (i RecordCompletionReason) String() string {
    return []string{"OPERATIONCANCELED", "STOPTONEDETECTED", "MAXRECORDDURATIONREACHED", "INITIALSILENCETIMEOUT", "MAXSILENCETIMEOUT", "PLAYPROMPTFAILED", "PLAYBEEPFAILED", "MEDIARECEIVETIMEOUT", "UNSPECIFIEDERROR"}[i]
}
func ParseRecordCompletionReason(v string) (interface{}, error) {
    result := OPERATIONCANCELED_RECORDCOMPLETIONREASON
    switch strings.ToUpper(v) {
        case "OPERATIONCANCELED":
            result = OPERATIONCANCELED_RECORDCOMPLETIONREASON
        case "STOPTONEDETECTED":
            result = STOPTONEDETECTED_RECORDCOMPLETIONREASON
        case "MAXRECORDDURATIONREACHED":
            result = MAXRECORDDURATIONREACHED_RECORDCOMPLETIONREASON
        case "INITIALSILENCETIMEOUT":
            result = INITIALSILENCETIMEOUT_RECORDCOMPLETIONREASON
        case "MAXSILENCETIMEOUT":
            result = MAXSILENCETIMEOUT_RECORDCOMPLETIONREASON
        case "PLAYPROMPTFAILED":
            result = PLAYPROMPTFAILED_RECORDCOMPLETIONREASON
        case "PLAYBEEPFAILED":
            result = PLAYBEEPFAILED_RECORDCOMPLETIONREASON
        case "MEDIARECEIVETIMEOUT":
            result = MEDIARECEIVETIMEOUT_RECORDCOMPLETIONREASON
        case "UNSPECIFIEDERROR":
            result = UNSPECIFIEDERROR_RECORDCOMPLETIONREASON
        default:
            return 0, errors.New("Unknown RecordCompletionReason value: " + v)
    }
    return &result, nil
}
func SerializeRecordCompletionReason(values []RecordCompletionReason) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
