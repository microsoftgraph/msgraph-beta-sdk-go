package ediscovery
import (
    "strings"
    "errors"
)
// 
type DataSourceHoldStatus int

const (
    NOTAPPLIED_DATASOURCEHOLDSTATUS DataSourceHoldStatus = iota
    APPLIED_DATASOURCEHOLDSTATUS
    APPLYING_DATASOURCEHOLDSTATUS
    REMOVING_DATASOURCEHOLDSTATUS
    PARTIAL_DATASOURCEHOLDSTATUS
    UNKNOWNFUTUREVALUE_DATASOURCEHOLDSTATUS
)

func (i DataSourceHoldStatus) String() string {
    return []string{"NOTAPPLIED", "APPLIED", "APPLYING", "REMOVING", "PARTIAL", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseDataSourceHoldStatus(v string) (interface{}, error) {
    result := NOTAPPLIED_DATASOURCEHOLDSTATUS
    switch strings.ToUpper(v) {
        case "NOTAPPLIED":
            result = NOTAPPLIED_DATASOURCEHOLDSTATUS
        case "APPLIED":
            result = APPLIED_DATASOURCEHOLDSTATUS
        case "APPLYING":
            result = APPLYING_DATASOURCEHOLDSTATUS
        case "REMOVING":
            result = REMOVING_DATASOURCEHOLDSTATUS
        case "PARTIAL":
            result = PARTIAL_DATASOURCEHOLDSTATUS
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_DATASOURCEHOLDSTATUS
        default:
            return 0, errors.New("Unknown DataSourceHoldStatus value: " + v)
    }
    return &result, nil
}
func SerializeDataSourceHoldStatus(values []DataSourceHoldStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
