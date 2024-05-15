package models
// Possible values for username source or email source.
type UserEmailSource int

const (
    // User principal name.
    USERPRINCIPALNAME_USEREMAILSOURCE UserEmailSource = iota
    // Primary SMTP address.
    PRIMARYSMTPADDRESS_USEREMAILSOURCE
)

func (i UserEmailSource) String() string {
    return []string{"userPrincipalName", "primarySmtpAddress"}[i]
}
func ParseUserEmailSource(v string) (any, error) {
    result := USERPRINCIPALNAME_USEREMAILSOURCE
    switch v {
        case "userPrincipalName":
            result = USERPRINCIPALNAME_USEREMAILSOURCE
        case "primarySmtpAddress":
            result = PRIMARYSMTPADDRESS_USEREMAILSOURCE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeUserEmailSource(values []UserEmailSource) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i UserEmailSource) isMultiValue() bool {
    return false
}
