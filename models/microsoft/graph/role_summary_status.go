package graph
import (
    "strings"
    "errors"
)
type RoleSummaryStatus int

const (
    OK_ROLESUMMARYSTATUS RoleSummaryStatus = iota
    BAD_ROLESUMMARYSTATUS
)

func (i RoleSummaryStatus) String() string {
    return []string{"OK", "BAD"}[i]
}
func ParseRoleSummaryStatus(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "OK":
            return OK_ROLESUMMARYSTATUS, nil
        case "BAD":
            return BAD_ROLESUMMARYSTATUS, nil
    }
    return 0, errors.New("Unknown RoleSummaryStatus value: " + v)
}
func SerializeRoleSummaryStatus(values []RoleSummaryStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
