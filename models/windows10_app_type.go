package models
// Windows 10 Application type.
type Windows10AppType int

const (
    // Desktop.
    DESKTOP_WINDOWS10APPTYPE Windows10AppType = iota
    // Universal.
    UNIVERSAL_WINDOWS10APPTYPE
)

func (i Windows10AppType) String() string {
    return []string{"desktop", "universal"}[i]
}
func ParseWindows10AppType(v string) (any, error) {
    result := DESKTOP_WINDOWS10APPTYPE
    switch v {
        case "desktop":
            result = DESKTOP_WINDOWS10APPTYPE
        case "universal":
            result = UNIVERSAL_WINDOWS10APPTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeWindows10AppType(values []Windows10AppType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i Windows10AppType) isMultiValue() bool {
    return false
}
