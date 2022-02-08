package ediscovery
import (
    "strings"
    "errors"
)
// 
type DataSourceScopes int

const (
    NONE_DATASOURCESCOPES DataSourceScopes = iota
    ALLTENANTMAILBOXES_DATASOURCESCOPES
    ALLTENANTSITES_DATASOURCESCOPES
    ALLCASECUSTODIANS_DATASOURCESCOPES
    ALLCASENONCUSTODIALDATASOURCES_DATASOURCESCOPES
    UNKNOWNFUTUREVALUE_DATASOURCESCOPES
)

func (i DataSourceScopes) String() string {
    return []string{"NONE", "ALLTENANTMAILBOXES", "ALLTENANTSITES", "ALLCASECUSTODIANS", "ALLCASENONCUSTODIALDATASOURCES", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseDataSourceScopes(v string) (interface{}, error) {
    result := NONE_DATASOURCESCOPES
    switch strings.ToUpper(v) {
        case "NONE":
            result = NONE_DATASOURCESCOPES
        case "ALLTENANTMAILBOXES":
            result = ALLTENANTMAILBOXES_DATASOURCESCOPES
        case "ALLTENANTSITES":
            result = ALLTENANTSITES_DATASOURCESCOPES
        case "ALLCASECUSTODIANS":
            result = ALLCASECUSTODIANS_DATASOURCESCOPES
        case "ALLCASENONCUSTODIALDATASOURCES":
            result = ALLCASENONCUSTODIALDATASOURCES_DATASOURCESCOPES
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_DATASOURCESCOPES
        default:
            return 0, errors.New("Unknown DataSourceScopes value: " + v)
    }
    return &result, nil
}
func SerializeDataSourceScopes(values []DataSourceScopes) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
