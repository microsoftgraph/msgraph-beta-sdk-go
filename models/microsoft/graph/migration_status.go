package graph
import (
    "strings"
    "errors"
)
// 
type MigrationStatus int

const (
    READY_MIGRATIONSTATUS MigrationStatus = iota
    NEEDSREVIEW_MIGRATIONSTATUS
    ADDITIONALSTEPSREQUIRED_MIGRATIONSTATUS
    UNKNOWNFUTUREVALUE_MIGRATIONSTATUS
)

func (i MigrationStatus) String() string {
    return []string{"READY", "NEEDSREVIEW", "ADDITIONALSTEPSREQUIRED", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseMigrationStatus(v string) (interface{}, error) {
    result := READY_MIGRATIONSTATUS
    switch strings.ToUpper(v) {
        case "READY":
            result = READY_MIGRATIONSTATUS
        case "NEEDSREVIEW":
            result = NEEDSREVIEW_MIGRATIONSTATUS
        case "ADDITIONALSTEPSREQUIRED":
            result = ADDITIONALSTEPSREQUIRED_MIGRATIONSTATUS
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_MIGRATIONSTATUS
        default:
            return 0, errors.New("Unknown MigrationStatus value: " + v)
    }
    return &result, nil
}
func SerializeMigrationStatus(values []MigrationStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
