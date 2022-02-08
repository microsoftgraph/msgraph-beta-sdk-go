package ediscovery
import (
    "strings"
    "errors"
)
// 
type DataSourceContainerStatus int

const (
    ACTIVE_DATASOURCECONTAINERSTATUS DataSourceContainerStatus = iota
    RELEASED_DATASOURCECONTAINERSTATUS
    UNKNOWNFUTUREVALUE_DATASOURCECONTAINERSTATUS
)

func (i DataSourceContainerStatus) String() string {
    return []string{"ACTIVE", "RELEASED", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseDataSourceContainerStatus(v string) (interface{}, error) {
    result := ACTIVE_DATASOURCECONTAINERSTATUS
    switch strings.ToUpper(v) {
        case "ACTIVE":
            result = ACTIVE_DATASOURCECONTAINERSTATUS
        case "RELEASED":
            result = RELEASED_DATASOURCECONTAINERSTATUS
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_DATASOURCECONTAINERSTATUS
        default:
            return 0, errors.New("Unknown DataSourceContainerStatus value: " + v)
    }
    return &result, nil
}
func SerializeDataSourceContainerStatus(values []DataSourceContainerStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
