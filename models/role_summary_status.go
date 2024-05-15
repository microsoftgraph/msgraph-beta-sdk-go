package models
type RoleSummaryStatus int

const (
    OK_ROLESUMMARYSTATUS RoleSummaryStatus = iota
    BAD_ROLESUMMARYSTATUS
)

func (i RoleSummaryStatus) String() string {
    return []string{"ok", "bad"}[i]
}
func ParseRoleSummaryStatus(v string) (any, error) {
    result := OK_ROLESUMMARYSTATUS
    switch v {
        case "ok":
            result = OK_ROLESUMMARYSTATUS
        case "bad":
            result = BAD_ROLESUMMARYSTATUS
        default:
            return nil, nil
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
func (i RoleSummaryStatus) isMultiValue() bool {
    return false
}
