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
    switch strings.ToUpper(v) {
        case "NOTAPPLIED":
            return NOTAPPLIED_DATASOURCEHOLDSTATUS, nil
        case "APPLIED":
            return APPLIED_DATASOURCEHOLDSTATUS, nil
        case "APPLYING":
            return APPLYING_DATASOURCEHOLDSTATUS, nil
        case "REMOVING":
            return REMOVING_DATASOURCEHOLDSTATUS, nil
        case "PARTIAL":
            return PARTIAL_DATASOURCEHOLDSTATUS, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_DATASOURCEHOLDSTATUS, nil
    }
    return 0, errors.New("Unknown DataSourceHoldStatus value: " + v)
}
func SerializeDataSourceHoldStatus(values []DataSourceHoldStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
