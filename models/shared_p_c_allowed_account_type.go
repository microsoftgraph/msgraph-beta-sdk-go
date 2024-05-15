package models
import (
    "math"
    "strings"
)
// Type of accounts that are allowed to share the PC.
type SharedPCAllowedAccountType int

const (
    // Not configured. Default value.
    NOTCONFIGURED_SHAREDPCALLOWEDACCOUNTTYPE = 1
    // Only guest accounts.
    GUEST_SHAREDPCALLOWEDACCOUNTTYPE = 2
    // Only domain-joined accounts.
    DOMAIN_SHAREDPCALLOWEDACCOUNTTYPE = 4
)

func (i SharedPCAllowedAccountType) String() string {
    var values []string
    options := []string{"notConfigured", "guest", "domain"}
    for p := 0; p < 3; p++ {
        mantis := SharedPCAllowedAccountType(int(math.Pow(2, float64(p))))
        if i&mantis == mantis {
            values = append(values, options[p])
        }
    }
    return strings.Join(values, ",")
}
func ParseSharedPCAllowedAccountType(v string) (any, error) {
    var result SharedPCAllowedAccountType
    values := strings.Split(v, ",")
    for _, str := range values {
        switch str {
            case "notConfigured":
                result |= NOTCONFIGURED_SHAREDPCALLOWEDACCOUNTTYPE
            case "guest":
                result |= GUEST_SHAREDPCALLOWEDACCOUNTTYPE
            case "domain":
                result |= DOMAIN_SHAREDPCALLOWEDACCOUNTTYPE
            default:
                return nil, nil
        }
    }
    return &result, nil
}
func SerializeSharedPCAllowedAccountType(values []SharedPCAllowedAccountType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i SharedPCAllowedAccountType) isMultiValue() bool {
    return true
}
