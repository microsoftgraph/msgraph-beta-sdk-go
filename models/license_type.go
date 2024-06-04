package models
// Indicates whether tenant has a valid Intune Endpoint Privilege Management license. Possible value are : 0 - notPaid, 1 - paid, 2 - trial. See LicenseType enum for more details. Default notPaid .
type LicenseType int

const (
    // Indicates the tenant has neither trial or paid license.
    NOTPAID_LICENSETYPE LicenseType = iota
    // Indicates the tenant has paid Endpoint Privilege Management license.
    PAID_LICENSETYPE
    // Indicates the tenant has trial Endpoint Privilege Management license.
    TRIAL_LICENSETYPE
    // Evolvable enumeration sentinel value. Do not use.
    UNKNOWNFUTUREVALUE_LICENSETYPE
)

func (i LicenseType) String() string {
    return []string{"notPaid", "paid", "trial", "unknownFutureValue"}[i]
}
func ParseLicenseType(v string) (any, error) {
    result := NOTPAID_LICENSETYPE
    switch v {
        case "notPaid":
            result = NOTPAID_LICENSETYPE
        case "paid":
            result = PAID_LICENSETYPE
        case "trial":
            result = TRIAL_LICENSETYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_LICENSETYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeLicenseType(values []LicenseType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i LicenseType) isMultiValue() bool {
    return false
}
