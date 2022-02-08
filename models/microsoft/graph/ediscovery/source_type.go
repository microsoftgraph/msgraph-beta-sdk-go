package ediscovery
import (
    "strings"
    "errors"
)
// 
type SourceType int

const (
    MAILBOX_SOURCETYPE SourceType = iota
    SITE_SOURCETYPE
)

func (i SourceType) String() string {
    return []string{"MAILBOX", "SITE"}[i]
}
func ParseSourceType(v string) (interface{}, error) {
    result := MAILBOX_SOURCETYPE
    switch strings.ToUpper(v) {
        case "MAILBOX":
            result = MAILBOX_SOURCETYPE
        case "SITE":
            result = SITE_SOURCETYPE
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
