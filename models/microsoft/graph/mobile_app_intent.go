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
    switch strings.ToUpper(v) {
        case "AVAILABLE":
            return AVAILABLE_MOBILEAPPINTENT, nil
        case "NOTAVAILABLE":
            return NOTAVAILABLE_MOBILEAPPINTENT, nil
        case "REQUIREDINSTALL":
            return REQUIREDINSTALL_MOBILEAPPINTENT, nil
        case "REQUIREDUNINSTALL":
            return REQUIREDUNINSTALL_MOBILEAPPINTENT, nil
        case "REQUIREDANDAVAILABLEINSTALL":
            return REQUIREDANDAVAILABLEINSTALL_MOBILEAPPINTENT, nil
        case "AVAILABLEINSTALLWITHOUTENROLLMENT":
            return AVAILABLEINSTALLWITHOUTENROLLMENT_MOBILEAPPINTENT, nil
        case "EXCLUDE":
            return EXCLUDE_MOBILEAPPINTENT, nil
    }
    return 0, errors.New("Unknown MobileAppIntent value: " + v)
}
func SerializeMobileAppIntent(values []MobileAppIntent) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
