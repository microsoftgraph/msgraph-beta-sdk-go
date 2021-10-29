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
    switch strings.ToUpper(v) {
        case "NOTMANAGED":
            return NOTMANAGED_MICROSOFTMANAGEDDESKTOPTYPE, nil
        case "PREMIUMMANAGED":
            return PREMIUMMANAGED_MICROSOFTMANAGEDDESKTOPTYPE, nil
        case "STANDARDMANAGED":
            return STANDARDMANAGED_MICROSOFTMANAGEDDESKTOPTYPE, nil
        case "STARTERMANAGED":
            return STARTERMANAGED_MICROSOFTMANAGEDDESKTOPTYPE, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_MICROSOFTMANAGEDDESKTOPTYPE, nil
    }
    return 0, errors.New("Unknown MicrosoftManagedDesktopType value: " + v)
}
func SerializeMicrosoftManagedDesktopType(values []MicrosoftManagedDesktopType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
