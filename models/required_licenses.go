package models
type RequiredLicenses int

const (
    NOTAPPLICABLE_REQUIREDLICENSES RequiredLicenses = iota
    MICROSOFTENTRAIDFREE_REQUIREDLICENSES
    MICROSOFTENTRAIDP1_REQUIREDLICENSES
    MICROSOFTENTRAIDP2_REQUIREDLICENSES
    MICROSOFTENTRAIDGOVERNANCE_REQUIREDLICENSES
    MICROSOFTENTRAWORKLOADID_REQUIREDLICENSES
    UNKNOWNFUTUREVALUE_REQUIREDLICENSES
)

func (i RequiredLicenses) String() string {
    return []string{"notApplicable", "microsoftEntraIdFree", "microsoftEntraIdP1", "microsoftEntraIdP2", "microsoftEntraIdGovernance", "microsoftEntraWorkloadId", "unknownFutureValue"}[i]
}
func ParseRequiredLicenses(v string) (any, error) {
    result := NOTAPPLICABLE_REQUIREDLICENSES
    switch v {
        case "notApplicable":
            result = NOTAPPLICABLE_REQUIREDLICENSES
        case "microsoftEntraIdFree":
            result = MICROSOFTENTRAIDFREE_REQUIREDLICENSES
        case "microsoftEntraIdP1":
            result = MICROSOFTENTRAIDP1_REQUIREDLICENSES
        case "microsoftEntraIdP2":
            result = MICROSOFTENTRAIDP2_REQUIREDLICENSES
        case "microsoftEntraIdGovernance":
            result = MICROSOFTENTRAIDGOVERNANCE_REQUIREDLICENSES
        case "microsoftEntraWorkloadId":
            result = MICROSOFTENTRAWORKLOADID_REQUIREDLICENSES
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_REQUIREDLICENSES
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeRequiredLicenses(values []RequiredLicenses) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i RequiredLicenses) isMultiValue() bool {
    return false
}
