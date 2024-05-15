package models
// Data can be transferred from/to these classes of apps
type WindowsManagedAppDataTransferLevel int

const (
    // All apps.
    ALLAPPS_WINDOWSMANAGEDAPPDATATRANSFERLEVEL WindowsManagedAppDataTransferLevel = iota
    // No apps.
    NONE_WINDOWSMANAGEDAPPDATATRANSFERLEVEL
)

func (i WindowsManagedAppDataTransferLevel) String() string {
    return []string{"allApps", "none"}[i]
}
func ParseWindowsManagedAppDataTransferLevel(v string) (any, error) {
    result := ALLAPPS_WINDOWSMANAGEDAPPDATATRANSFERLEVEL
    switch v {
        case "allApps":
            result = ALLAPPS_WINDOWSMANAGEDAPPDATATRANSFERLEVEL
        case "none":
            result = NONE_WINDOWSMANAGEDAPPDATATRANSFERLEVEL
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeWindowsManagedAppDataTransferLevel(values []WindowsManagedAppDataTransferLevel) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i WindowsManagedAppDataTransferLevel) isMultiValue() bool {
    return false
}
