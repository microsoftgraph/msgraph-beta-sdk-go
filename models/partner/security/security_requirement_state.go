package security
type SecurityRequirementState int

const (
    ACTIVE_SECURITYREQUIREMENTSTATE SecurityRequirementState = iota
    PREVIEW_SECURITYREQUIREMENTSTATE
    UNKNOWNFUTUREVALUE_SECURITYREQUIREMENTSTATE
)

func (i SecurityRequirementState) String() string {
    return []string{"active", "preview", "unknownFutureValue"}[i]
}
func ParseSecurityRequirementState(v string) (any, error) {
    result := ACTIVE_SECURITYREQUIREMENTSTATE
    switch v {
        case "active":
            result = ACTIVE_SECURITYREQUIREMENTSTATE
        case "preview":
            result = PREVIEW_SECURITYREQUIREMENTSTATE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_SECURITYREQUIREMENTSTATE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeSecurityRequirementState(values []SecurityRequirementState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i SecurityRequirementState) isMultiValue() bool {
    return false
}
