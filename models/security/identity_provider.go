// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package security
type IdentityProvider int

const (
    ENTRAID_IDENTITYPROVIDER IdentityProvider = iota
    ACTIVEDIRECTORY_IDENTITYPROVIDER
    OKTA_IDENTITYPROVIDER
    UNKNOWNFUTUREVALUE_IDENTITYPROVIDER
)

func (i IdentityProvider) String() string {
    return []string{"entraID", "activeDirectory", "okta", "unknownFutureValue"}[i]
}
func ParseIdentityProvider(v string) (any, error) {
    result := ENTRAID_IDENTITYPROVIDER
    switch v {
        case "entraID":
            result = ENTRAID_IDENTITYPROVIDER
        case "activeDirectory":
            result = ACTIVEDIRECTORY_IDENTITYPROVIDER
        case "okta":
            result = OKTA_IDENTITYPROVIDER
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_IDENTITYPROVIDER
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeIdentityProvider(values []IdentityProvider) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i IdentityProvider) isMultiValue() bool {
    return false
}
