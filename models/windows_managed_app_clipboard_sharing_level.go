// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package models
// Represents the level to which the device's clipboard may be shared between apps
type WindowsManagedAppClipboardSharingLevel int

const (
    // Org users can paste data from and cut/copy data to any account, document, location or application.
    ANYDESTINATIONANYSOURCE_WINDOWSMANAGEDAPPCLIPBOARDSHARINGLEVEL WindowsManagedAppClipboardSharingLevel = iota
    // Org users cannot cut, copy or paste data to or from external accounts, documents, locations or applications from or into the org context.
    NONE_WINDOWSMANAGEDAPPCLIPBOARDSHARINGLEVEL
)

func (i WindowsManagedAppClipboardSharingLevel) String() string {
    return []string{"anyDestinationAnySource", "none"}[i]
}
func ParseWindowsManagedAppClipboardSharingLevel(v string) (any, error) {
    result := ANYDESTINATIONANYSOURCE_WINDOWSMANAGEDAPPCLIPBOARDSHARINGLEVEL
    switch v {
        case "anyDestinationAnySource":
            result = ANYDESTINATIONANYSOURCE_WINDOWSMANAGEDAPPCLIPBOARDSHARINGLEVEL
        case "none":
            result = NONE_WINDOWSMANAGEDAPPCLIPBOARDSHARINGLEVEL
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeWindowsManagedAppClipboardSharingLevel(values []WindowsManagedAppClipboardSharingLevel) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i WindowsManagedAppClipboardSharingLevel) isMultiValue() bool {
    return false
}
