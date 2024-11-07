package networkaccess
type NetworkTrafficOperationStatus int

const (
    SUCCESS_NETWORKTRAFFICOPERATIONSTATUS NetworkTrafficOperationStatus = iota
    FAILURE_NETWORKTRAFFICOPERATIONSTATUS
    UNKNOWNFUTUREVALUE_NETWORKTRAFFICOPERATIONSTATUS
)

func (i NetworkTrafficOperationStatus) String() string {
    return []string{"success", "failure", "unknownFutureValue"}[i]
}
func ParseNetworkTrafficOperationStatus(v string) (any, error) {
    result := SUCCESS_NETWORKTRAFFICOPERATIONSTATUS
    switch v {
        case "success":
            result = SUCCESS_NETWORKTRAFFICOPERATIONSTATUS
        case "failure":
            result = FAILURE_NETWORKTRAFFICOPERATIONSTATUS
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_NETWORKTRAFFICOPERATIONSTATUS
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeNetworkTrafficOperationStatus(values []NetworkTrafficOperationStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i NetworkTrafficOperationStatus) isMultiValue() bool {
    return false
}
