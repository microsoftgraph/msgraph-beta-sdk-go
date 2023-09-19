package models
import (
    "errors"
    "strings"
)
// Flag enum representing the allowed macOS system extension types.
type MacOSSystemExtensionType int

const (
    // Enables driver extensions.
    DRIVEREXTENSIONSALLOWED_MACOSSYSTEMEXTENSIONTYPE MacOSSystemExtensionType = iota
    // Enables network extensions.
    NETWORKEXTENSIONSALLOWED_MACOSSYSTEMEXTENSIONTYPE
    // Enables endpoint security extensions.
    ENDPOINTSECURITYEXTENSIONSALLOWED_MACOSSYSTEMEXTENSIONTYPE
)

func (i MacOSSystemExtensionType) String() string {
    var values []string
    for p := MacOSSystemExtensionType(1); p <= ENDPOINTSECURITYEXTENSIONSALLOWED_MACOSSYSTEMEXTENSIONTYPE; p <<= 1 {
        if i&p == p {
            values = append(values, []string{"driverExtensionsAllowed", "networkExtensionsAllowed", "endpointSecurityExtensionsAllowed"}[p])
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
