package models
// Android username source.
type AndroidUsernameSource int

const (
    // The username.
    USERNAME_ANDROIDUSERNAMESOURCE AndroidUsernameSource = iota
    // The user principal name.
    USERPRINCIPALNAME_ANDROIDUSERNAMESOURCE
    // The user sam account name.
    SAMACCOUNTNAME_ANDROIDUSERNAMESOURCE
    // Primary SMTP address.
    PRIMARYSMTPADDRESS_ANDROIDUSERNAMESOURCE
)

func (i AndroidUsernameSource) String() string {
    return []string{"username", "userPrincipalName", "samAccountName", "primarySmtpAddress"}[i]
}
func ParseAndroidUsernameSource(v string) (any, error) {
    result := USERNAME_ANDROIDUSERNAMESOURCE
    switch v {
        case "username":
            result = USERNAME_ANDROIDUSERNAMESOURCE
        case "userPrincipalName":
            result = USERPRINCIPALNAME_ANDROIDUSERNAMESOURCE
        case "samAccountName":
            result = SAMACCOUNTNAME_ANDROIDUSERNAMESOURCE
        case "primarySmtpAddress":
            result = PRIMARYSMTPADDRESS_ANDROIDUSERNAMESOURCE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeAndroidUsernameSource(values []AndroidUsernameSource) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AndroidUsernameSource) isMultiValue() bool {
    return false
}
