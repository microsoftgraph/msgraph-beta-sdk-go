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
    switch strings.ToUpper(v) {
        case "READY":
            return READY_MIGRATIONSTATUS, nil
        case "NEEDSREVIEW":
            return NEEDSREVIEW_MIGRATIONSTATUS, nil
        case "ADDITIONALSTEPSREQUIRED":
            return ADDITIONALSTEPSREQUIRED_MIGRATIONSTATUS, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_MIGRATIONSTATUS, nil
    }
    return 0, errors.New("Unknown MigrationStatus value: " + v)
}
func SerializeMigrationStatus(values []MigrationStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
