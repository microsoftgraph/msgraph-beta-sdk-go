package models
// Process identifier types for MacOS Privacy Preferences
type MacOSProcessIdentifierType int

const (
    // Indicates an app with a bundle ID.
    BUNDLEID_MACOSPROCESSIDENTIFIERTYPE MacOSProcessIdentifierType = iota
    // Indicates a file path for a process.
    PATH_MACOSPROCESSIDENTIFIERTYPE
)

func (i MacOSProcessIdentifierType) String() string {
    return []string{"bundleID", "path"}[i]
}
func ParseMacOSProcessIdentifierType(v string) (any, error) {
    result := BUNDLEID_MACOSPROCESSIDENTIFIERTYPE
    switch v {
        case "bundleID":
            result = BUNDLEID_MACOSPROCESSIDENTIFIERTYPE
        case "path":
            result = PATH_MACOSPROCESSIDENTIFIERTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeMacOSProcessIdentifierType(values []MacOSProcessIdentifierType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i MacOSProcessIdentifierType) isMultiValue() bool {
    return false
}
