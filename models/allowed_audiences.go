package models
import (
    "math"
    "strings"
)
type AllowedAudiences int

const (
    ME_ALLOWEDAUDIENCES = 1
    FAMILY_ALLOWEDAUDIENCES = 2
    CONTACTS_ALLOWEDAUDIENCES = 4
    GROUPMEMBERS_ALLOWEDAUDIENCES = 8
    ORGANIZATION_ALLOWEDAUDIENCES = 16
    FEDERATEDORGANIZATIONS_ALLOWEDAUDIENCES = 32
    EVERYONE_ALLOWEDAUDIENCES = 64
    UNKNOWNFUTUREVALUE_ALLOWEDAUDIENCES = 128
)

func (i AllowedAudiences) String() string {
    var values []string
    options := []string{"me", "family", "contacts", "groupMembers", "organization", "federatedOrganizations", "everyone", "unknownFutureValue"}
    for p := 0; p < 8; p++ {
        mantis := AllowedAudiences(int(math.Pow(2, float64(p))))
        if i&mantis == mantis {
            values = append(values, options[p])
        }
    }
    return strings.Join(values, ",")
}
func ParseAllowedAudiences(v string) (any, error) {
    var result AllowedAudiences
    values := strings.Split(v, ",")
    for _, str := range values {
        switch str {
            case "me":
                result |= ME_ALLOWEDAUDIENCES
            case "family":
                result |= FAMILY_ALLOWEDAUDIENCES
            case "contacts":
                result |= CONTACTS_ALLOWEDAUDIENCES
            case "groupMembers":
                result |= GROUPMEMBERS_ALLOWEDAUDIENCES
            case "organization":
                result |= ORGANIZATION_ALLOWEDAUDIENCES
            case "federatedOrganizations":
                result |= FEDERATEDORGANIZATIONS_ALLOWEDAUDIENCES
            case "everyone":
                result |= EVERYONE_ALLOWEDAUDIENCES
            case "unknownFutureValue":
                result |= UNKNOWNFUTUREVALUE_ALLOWEDAUDIENCES
            default:
                return nil, nil
        }
    }
    return &result, nil
}
func SerializeAllowedAudiences(values []AllowedAudiences) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AllowedAudiences) isMultiValue() bool {
    return true
}
