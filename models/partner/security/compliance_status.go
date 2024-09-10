package security
type ComplianceStatus int

const (
    COMPLIANT_COMPLIANCESTATUS ComplianceStatus = iota
    NONCOMPLAINT_COMPLIANCESTATUS
    UNKNOWNFUTUREVALUE_COMPLIANCESTATUS
)

func (i ComplianceStatus) String() string {
    return []string{"compliant", "noncomplaint", "unknownFutureValue"}[i]
}
func ParseComplianceStatus(v string) (any, error) {
    result := COMPLIANT_COMPLIANCESTATUS
    switch v {
        case "compliant":
            result = COMPLIANT_COMPLIANCESTATUS
        case "noncomplaint":
            result = NONCOMPLAINT_COMPLIANCESTATUS
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_COMPLIANCESTATUS
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeComplianceStatus(values []ComplianceStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ComplianceStatus) isMultiValue() bool {
    return false
}
