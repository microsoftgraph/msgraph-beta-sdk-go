// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package ediscovery
type DataSourceContainerStatus int

const (
    ACTIVE_DATASOURCECONTAINERSTATUS DataSourceContainerStatus = iota
    RELEASED_DATASOURCECONTAINERSTATUS
    UNKNOWNFUTUREVALUE_DATASOURCECONTAINERSTATUS
)

func (i DataSourceContainerStatus) String() string {
    return []string{"Active", "Released", "UnknownFutureValue"}[i]
}
func ParseDataSourceContainerStatus(v string) (any, error) {
    result := ACTIVE_DATASOURCECONTAINERSTATUS
    switch v {
        case "Active":
            result = ACTIVE_DATASOURCECONTAINERSTATUS
        case "Released":
            result = RELEASED_DATASOURCECONTAINERSTATUS
        case "UnknownFutureValue":
            result = UNKNOWNFUTUREVALUE_DATASOURCECONTAINERSTATUS
        default:
            return nil, nil
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
func (i DataSourceContainerStatus) isMultiValue() bool {
    return false
}
