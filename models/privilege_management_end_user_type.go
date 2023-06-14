package models
import (
    "errors"
)
// The type of user account on Windows that was used to performed the elevation.
type PrivilegeManagementEndUserType int

const (
    // Default. Unable to determine the login type of the user.
    UNDETERMINED_PRIVILEGEMANAGEMENTENDUSERTYPE PrivilegeManagementEndUserType = iota
    // The user who performed the elevation logged in using an Azure Active Directory (Azure AD) account.
    AZUREAD_PRIVILEGEMANAGEMENTENDUSERTYPE
    // The user who performed the elevation logged in using Hybrid Azure AD joined account.
    HYBRID_PRIVILEGEMANAGEMENTENDUSERTYPE
    // The user who performed the elevation logged in using a Windows local account.
    LOCAL_PRIVILEGEMANAGEMENTENDUSERTYPE
    // Evolvable emuneration sentinel value. Do not use.
    UNKNOWNFUTUREVALUE_PRIVILEGEMANAGEMENTENDUSERTYPE
)

func (i PrivilegeManagementEndUserType) String() string {
    return []string{"undetermined", "azureAd", "hybrid", "local", "unknownFutureValue"}[i]
}
func ParsePrivilegeManagementEndUserType(v string) (any, error) {
    result := UNDETERMINED_PRIVILEGEMANAGEMENTENDUSERTYPE
    switch v {
        case "undetermined":
            result = UNDETERMINED_PRIVILEGEMANAGEMENTENDUSERTYPE
        case "azureAd":
            result = AZUREAD_PRIVILEGEMANAGEMENTENDUSERTYPE
        case "hybrid":
            result = HYBRID_PRIVILEGEMANAGEMENTENDUSERTYPE
        case "local":
            result = LOCAL_PRIVILEGEMANAGEMENTENDUSERTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_PRIVILEGEMANAGEMENTENDUSERTYPE
        default:
            return 0, errors.New("Unknown PrivilegeManagementEndUserType value: " + v)
    }
    return &result, nil
}
func SerializePrivilegeManagementEndUserType(values []PrivilegeManagementEndUserType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
