package models
// Indicates the dependency type associated with a relationship between two mobile apps.
type MobileAppDependencyType int

const (
    // Indicates that the child app should be detected before installing the parent app.
    DETECT_MOBILEAPPDEPENDENCYTYPE MobileAppDependencyType = iota
    // Indicates that the child app should be installed before installing the parent app.
    AUTOINSTALL_MOBILEAPPDEPENDENCYTYPE
)

func (i MobileAppDependencyType) String() string {
    return []string{"detect", "autoInstall"}[i]
}
func ParseMobileAppDependencyType(v string) (any, error) {
    result := DETECT_MOBILEAPPDEPENDENCYTYPE
    switch v {
        case "detect":
            result = DETECT_MOBILEAPPDEPENDENCYTYPE
        case "autoInstall":
            result = AUTOINSTALL_MOBILEAPPDEPENDENCYTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeMobileAppDependencyType(values []MobileAppDependencyType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i MobileAppDependencyType) isMultiValue() bool {
    return false
}
