package security
import (
    "strings"
    "errors"
)
// Provides operations to manage the cases property of the microsoft.graph.security entity.
type SourceType int

const (
    MAILBOX_SOURCETYPE SourceType = iota
    SITE_SOURCETYPE
    UNKNOWNFUTUREVALUE_SOURCETYPE
)

func (i SourceType) String() string {
    return []string{"MAILBOX", "SITE", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseSourceType(v string) (interface{}, error) {
    result := MAILBOX_SOURCETYPE
    switch strings.ToUpper(v) {
        case "MAILBOX":
            result = MAILBOX_SOURCETYPE
        case "SITE":
            result = SITE_SOURCETYPE
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_SOURCETYPE
        default:
            return 0, errors.New("Unknown SourceType value: " + v)
    }
    return &result, nil
}
func SerializeSourceType(values []SourceType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
