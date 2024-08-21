package models
type ProvisionState int

const (
    NOTPROVISIONED_PROVISIONSTATE ProvisionState = iota
    PROVISIONINGINPROGRESS_PROVISIONSTATE
    PROVISIONINGFAILED_PROVISIONSTATE
    PROVISIONINGCOMPLETED_PROVISIONSTATE
    UNKNOWNFUTUREVALUE_PROVISIONSTATE
)

func (i ProvisionState) String() string {
    return []string{"notProvisioned", "provisioningInProgress", "provisioningFailed", "provisioningCompleted", "unknownFutureValue"}[i]
}
func ParseProvisionState(v string) (any, error) {
    result := NOTPROVISIONED_PROVISIONSTATE
    switch v {
        case "notProvisioned":
            result = NOTPROVISIONED_PROVISIONSTATE
        case "provisioningInProgress":
            result = PROVISIONINGINPROGRESS_PROVISIONSTATE
        case "provisioningFailed":
            result = PROVISIONINGFAILED_PROVISIONSTATE
        case "provisioningCompleted":
            result = PROVISIONINGCOMPLETED_PROVISIONSTATE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_PROVISIONSTATE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeProvisionState(values []ProvisionState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ProvisionState) isMultiValue() bool {
    return false
}
