package security
type SecurityRequirementType int

const (
    MFAENFORCEDFORADMINS_SECURITYREQUIREMENTTYPE SecurityRequirementType = iota
    MFAENFORCEDFORADMINSOFCUSTOMERS_SECURITYREQUIREMENTTYPE
    SECURITYALERTSPROMPTLYRESOLVED_SECURITYREQUIREMENTTYPE
    SECURITYCONTACTPROVIDED_SECURITYREQUIREMENTTYPE
    SPENDINGBUDGETSETFORCUSTOMERAZURESUBSCRIPTIONS_SECURITYREQUIREMENTTYPE
    UNKNOWNFUTUREVALUE_SECURITYREQUIREMENTTYPE
)

func (i SecurityRequirementType) String() string {
    return []string{"mfaEnforcedForAdmins", "mfaEnforcedForAdminsOfCustomers", "securityAlertsPromptlyResolved", "securityContactProvided", "spendingBudgetSetForCustomerAzureSubscriptions", "unknownFutureValue"}[i]
}
func ParseSecurityRequirementType(v string) (any, error) {
    result := MFAENFORCEDFORADMINS_SECURITYREQUIREMENTTYPE
    switch v {
        case "mfaEnforcedForAdmins":
            result = MFAENFORCEDFORADMINS_SECURITYREQUIREMENTTYPE
        case "mfaEnforcedForAdminsOfCustomers":
            result = MFAENFORCEDFORADMINSOFCUSTOMERS_SECURITYREQUIREMENTTYPE
        case "securityAlertsPromptlyResolved":
            result = SECURITYALERTSPROMPTLYRESOLVED_SECURITYREQUIREMENTTYPE
        case "securityContactProvided":
            result = SECURITYCONTACTPROVIDED_SECURITYREQUIREMENTTYPE
        case "spendingBudgetSetForCustomerAzureSubscriptions":
            result = SPENDINGBUDGETSETFORCUSTOMERAZURESUBSCRIPTIONS_SECURITYREQUIREMENTTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_SECURITYREQUIREMENTTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeSecurityRequirementType(values []SecurityRequirementType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i SecurityRequirementType) isMultiValue() bool {
    return false
}
