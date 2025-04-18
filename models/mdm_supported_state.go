// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package models
// Mdm Support Status of the setting.
type MdmSupportedState int

const (
    // Mdm support status of the setting is not known.
    UNKNOWN_MDMSUPPORTEDSTATE MdmSupportedState = iota
    // Setting is supported.
    SUPPORTED_MDMSUPPORTEDSTATE
    // Setting is unsupported.
    UNSUPPORTED_MDMSUPPORTEDSTATE
    // Setting is depcrecated.
    DEPRECATED_MDMSUPPORTEDSTATE
)

func (i MdmSupportedState) String() string {
    return []string{"unknown", "supported", "unsupported", "deprecated"}[i]
}
func ParseMdmSupportedState(v string) (any, error) {
    result := UNKNOWN_MDMSUPPORTEDSTATE
    switch v {
        case "unknown":
            result = UNKNOWN_MDMSUPPORTEDSTATE
        case "supported":
            result = SUPPORTED_MDMSUPPORTEDSTATE
        case "unsupported":
            result = UNSUPPORTED_MDMSUPPORTEDSTATE
        case "deprecated":
            result = DEPRECATED_MDMSUPPORTEDSTATE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeMdmSupportedState(values []MdmSupportedState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i MdmSupportedState) isMultiValue() bool {
    return false
}
