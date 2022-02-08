package graph
import (
    "strings"
    "errors"
)
// 
type RoleSummaryStatus int

const (
    OK_ROLESUMMARYSTATUS RoleSummaryStatus = iota
    BAD_ROLESUMMARYSTATUS
)

func (i RoleSummaryStatus) String() string {
    return []string{"OK", "BAD"}[i]
}
func ParseRoleSummaryStatus(v string) (interface{}, error) {
    result := OK_ROLESUMMARYSTATUS
    switch strings.ToUpper(v) {
        case "OK":
            result = OK_ROLESUMMARYSTATUS
        case "BAD":
            result = BAD_ROLESUMMARYSTATUS
        default:
            return 0, errors.New("Unknown RoleSummaryStatus value: " + v)
    }
    return &result, nil
}
func SerializeRoleSummaryStatus(values []RoleSummaryStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
