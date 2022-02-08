package graph
import (
    "strings"
    "errors"
)
// 
type MobileAppIntent int

const (
    AVAILABLE_MOBILEAPPINTENT MobileAppIntent = iota
    NOTAVAILABLE_MOBILEAPPINTENT
    REQUIREDINSTALL_MOBILEAPPINTENT
    REQUIREDUNINSTALL_MOBILEAPPINTENT
    REQUIREDANDAVAILABLEINSTALL_MOBILEAPPINTENT
    AVAILABLEINSTALLWITHOUTENROLLMENT_MOBILEAPPINTENT
    EXCLUDE_MOBILEAPPINTENT
)

func (i MobileAppIntent) String() string {
    return []string{"AVAILABLE", "NOTAVAILABLE", "REQUIREDINSTALL", "REQUIREDUNINSTALL", "REQUIREDANDAVAILABLEINSTALL", "AVAILABLEINSTALLWITHOUTENROLLMENT", "EXCLUDE"}[i]
}
func ParseMobileAppIntent(v string) (interface{}, error) {
    result := AVAILABLE_MOBILEAPPINTENT
    switch strings.ToUpper(v) {
        case "AVAILABLE":
            result = AVAILABLE_MOBILEAPPINTENT
        case "NOTAVAILABLE":
            result = NOTAVAILABLE_MOBILEAPPINTENT
        case "REQUIREDINSTALL":
            result = REQUIREDINSTALL_MOBILEAPPINTENT
        case "REQUIREDUNINSTALL":
            result = REQUIREDUNINSTALL_MOBILEAPPINTENT
        case "REQUIREDANDAVAILABLEINSTALL":
            result = REQUIREDANDAVAILABLEINSTALL_MOBILEAPPINTENT
        case "AVAILABLEINSTALLWITHOUTENROLLMENT":
            result = AVAILABLEINSTALLWITHOUTENROLLMENT_MOBILEAPPINTENT
        case "EXCLUDE":
            result = EXCLUDE_MOBILEAPPINTENT
        default:
            return 0, errors.New("Unknown MobileAppIntent value: " + v)
    }
    return &result, nil
}
func SerializeMobileAppIntent(values []MobileAppIntent) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
