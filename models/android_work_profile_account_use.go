package models
import (
    "errors"
)
// An enum representing possible values for account use in work profile.
type AndroidWorkProfileAccountUse int

const (
    // Allow additon of all accounts except Google accounts in Android Work Profile.
    ALLOWALLEXCEPTGOOGLEACCOUNTS_ANDROIDWORKPROFILEACCOUNTUSE AndroidWorkProfileAccountUse = iota
    // Block any account from being added in Android Work Profile. 
    BLOCKALL_ANDROIDWORKPROFILEACCOUNTUSE
    // Allow addition of all accounts (including Google accounts) in Android Work Profile.
    ALLOWALL_ANDROIDWORKPROFILEACCOUNTUSE
    // Unknown future value for evolvable enum patterns.
    UNKNOWNFUTUREVALUE_ANDROIDWORKPROFILEACCOUNTUSE
)

func (i AndroidWorkProfileAccountUse) String() string {
    return []string{"allowAllExceptGoogleAccounts", "blockAll", "allowAll", "unknownFutureValue"}[i]
}
func ParseAndroidWorkProfileAccountUse(v string) (any, error) {
    result := ALLOWALLEXCEPTGOOGLEACCOUNTS_ANDROIDWORKPROFILEACCOUNTUSE
    switch v {
        case "allowAllExceptGoogleAccounts":
            result = ALLOWALLEXCEPTGOOGLEACCOUNTS_ANDROIDWORKPROFILEACCOUNTUSE
        case "blockAll":
            result = BLOCKALL_ANDROIDWORKPROFILEACCOUNTUSE
        case "allowAll":
            result = ALLOWALL_ANDROIDWORKPROFILEACCOUNTUSE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_ANDROIDWORKPROFILEACCOUNTUSE
        default:
            return 0, errors.New("Unknown AndroidWorkProfileAccountUse value: " + v)
    }
    return &result, nil
}
func SerializeAndroidWorkProfileAccountUse(values []AndroidWorkProfileAccountUse) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
