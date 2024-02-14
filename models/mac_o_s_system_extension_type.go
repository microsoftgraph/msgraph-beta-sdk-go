package models
import (
    "errors"
    "math"
    "strings"
)
// Flag enum representing the allowed macOS system extension types.
type MacOSSystemExtensionType int

const (
    // Enables driver extensions.
    DRIVEREXTENSIONSALLOWED_MACOSSYSTEMEXTENSIONTYPE = 1
    // Enables network extensions.
    NETWORKEXTENSIONSALLOWED_MACOSSYSTEMEXTENSIONTYPE = 2
    // Enables endpoint security extensions.
    ENDPOINTSECURITYEXTENSIONSALLOWED_MACOSSYSTEMEXTENSIONTYPE = 4
)

func (i MacOSSystemExtensionType) String() string {
    var values []string
    options := []string{"driverExtensionsAllowed", "networkExtensionsAllowed", "endpointSecurityExtensionsAllowed"}
    for p := 0; p < 3; p++ {
        mantis := MacOSSystemExtensionType(int(math.Pow(2, float64(p))))
        if i&mantis == mantis {
            values = append(values, options[p])
        }
    }
    return strings.Join(values, ",")
}
func ParseMacOSSystemExtensionType(v string) (any, error) {
    var result MacOSSystemExtensionType
    values := strings.Split(v, ",")
    for _, str := range values {
        switch str {
            case "driverExtensionsAllowed":
                result |= DRIVEREXTENSIONSALLOWED_MACOSSYSTEMEXTENSIONTYPE
            case "networkExtensionsAllowed":
                result |= NETWORKEXTENSIONSALLOWED_MACOSSYSTEMEXTENSIONTYPE
            case "endpointSecurityExtensionsAllowed":
                result |= ENDPOINTSECURITYEXTENSIONSALLOWED_MACOSSYSTEMEXTENSIONTYPE
            default:
                return 0, errors.New("Unknown MacOSSystemExtensionType value: " + v)
        }
    }
    return &result, nil
}
func SerializeMacOSSystemExtensionType(values []MacOSSystemExtensionType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i MacOSSystemExtensionType) isMultiValue() bool {
    return true
}
