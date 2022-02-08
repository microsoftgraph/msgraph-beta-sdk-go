package graph
import (
    "strings"
    "errors"
)
// 
type MicrosoftManagedDesktopType int

const (
    NOTMANAGED_MICROSOFTMANAGEDDESKTOPTYPE MicrosoftManagedDesktopType = iota
    PREMIUMMANAGED_MICROSOFTMANAGEDDESKTOPTYPE
    STANDARDMANAGED_MICROSOFTMANAGEDDESKTOPTYPE
    STARTERMANAGED_MICROSOFTMANAGEDDESKTOPTYPE
    UNKNOWNFUTUREVALUE_MICROSOFTMANAGEDDESKTOPTYPE
)

func (i MicrosoftManagedDesktopType) String() string {
    return []string{"NOTMANAGED", "PREMIUMMANAGED", "STANDARDMANAGED", "STARTERMANAGED", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseMicrosoftManagedDesktopType(v string) (interface{}, error) {
    result := NOTMANAGED_MICROSOFTMANAGEDDESKTOPTYPE
    switch strings.ToUpper(v) {
        case "NOTMANAGED":
            result = NOTMANAGED_MICROSOFTMANAGEDDESKTOPTYPE
        case "PREMIUMMANAGED":
            result = PREMIUMMANAGED_MICROSOFTMANAGEDDESKTOPTYPE
        case "STANDARDMANAGED":
            result = STANDARDMANAGED_MICROSOFTMANAGEDDESKTOPTYPE
        case "STARTERMANAGED":
            result = STARTERMANAGED_MICROSOFTMANAGEDDESKTOPTYPE
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_MICROSOFTMANAGEDDESKTOPTYPE
        default:
            return 0, errors.New("Unknown MicrosoftManagedDesktopType value: " + v)
    }
    return &result, nil
}
func SerializeMicrosoftManagedDesktopType(values []MicrosoftManagedDesktopType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
