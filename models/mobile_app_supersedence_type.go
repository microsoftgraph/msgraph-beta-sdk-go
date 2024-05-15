package models
// Indicates the supersedence type associated with a relationship between two mobile apps.
type MobileAppSupersedenceType int

const (
    // Indicates that the child app should be updated by the internal logic of the parent app.
    UPDATE_MOBILEAPPSUPERSEDENCETYPE MobileAppSupersedenceType = iota
    // Indicates that the child app should be uninstalled before installing the parent app.
    REPLACE_MOBILEAPPSUPERSEDENCETYPE
)

func (i MobileAppSupersedenceType) String() string {
    return []string{"update", "replace"}[i]
}
func ParseMobileAppSupersedenceType(v string) (any, error) {
    result := UPDATE_MOBILEAPPSUPERSEDENCETYPE
    switch v {
        case "update":
            result = UPDATE_MOBILEAPPSUPERSEDENCETYPE
        case "replace":
            result = REPLACE_MOBILEAPPSUPERSEDENCETYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeMobileAppSupersedenceType(values []MobileAppSupersedenceType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i MobileAppSupersedenceType) isMultiValue() bool {
    return false
}
