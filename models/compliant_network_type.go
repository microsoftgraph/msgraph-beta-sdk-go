package models
type CompliantNetworkType int

const (
    ALLTENANTCOMPLIANTNETWORKS_COMPLIANTNETWORKTYPE CompliantNetworkType = iota
    UNKNOWNFUTUREVALUE_COMPLIANTNETWORKTYPE
)

func (i CompliantNetworkType) String() string {
    return []string{"allTenantCompliantNetworks", "unknownFutureValue"}[i]
}
func ParseCompliantNetworkType(v string) (any, error) {
    result := ALLTENANTCOMPLIANTNETWORKS_COMPLIANTNETWORKTYPE
    switch v {
        case "allTenantCompliantNetworks":
            result = ALLTENANTCOMPLIANTNETWORKS_COMPLIANTNETWORKTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_COMPLIANTNETWORKTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeCompliantNetworkType(values []CompliantNetworkType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CompliantNetworkType) isMultiValue() bool {
    return false
}
