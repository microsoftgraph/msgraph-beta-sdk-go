package managedtenants
type ManagementCategory int

const (
    CUSTOM_MANAGEMENTCATEGORY ManagementCategory = iota
    DEVICES_MANAGEMENTCATEGORY
    IDENTITY_MANAGEMENTCATEGORY
    DATA_MANAGEMENTCATEGORY
    UNKNOWNFUTUREVALUE_MANAGEMENTCATEGORY
)

func (i ManagementCategory) String() string {
    return []string{"custom", "devices", "identity", "data", "unknownFutureValue"}[i]
}
func ParseManagementCategory(v string) (any, error) {
    result := CUSTOM_MANAGEMENTCATEGORY
    switch v {
        case "custom":
            result = CUSTOM_MANAGEMENTCATEGORY
        case "devices":
            result = DEVICES_MANAGEMENTCATEGORY
        case "identity":
            result = IDENTITY_MANAGEMENTCATEGORY
        case "data":
            result = DATA_MANAGEMENTCATEGORY
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_MANAGEMENTCATEGORY
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeManagementCategory(values []ManagementCategory) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ManagementCategory) isMultiValue() bool {
    return false
}
