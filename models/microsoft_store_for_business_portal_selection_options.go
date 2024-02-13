package models
import (
    "errors"
    "math"
    "strings"
)
// Portal to which admin syncs available Microsoft Store for Business apps. This is available in the Intune Admin Console.
type MicrosoftStoreForBusinessPortalSelectionOptions int

const (
    // This option is not available for the account
    NONE_MICROSOFTSTOREFORBUSINESSPORTALSELECTIONOPTIONS = 1
    // Intune Company Portal only.
    COMPANYPORTAL_MICROSOFTSTOREFORBUSINESSPORTALSELECTIONOPTIONS = 2
    // MSFB Private store only.
    PRIVATESTORE_MICROSOFTSTOREFORBUSINESSPORTALSELECTIONOPTIONS = 4
)

func (i MicrosoftStoreForBusinessPortalSelectionOptions) String() string {
    var values []string
    options := []string{"none", "companyPortal", "privateStore"}
    for p := 0; p < 3; p++ {
        mantis := MicrosoftStoreForBusinessPortalSelectionOptions(int(math.Pow(2, float64(p))))
        if i&mantis == mantis {
            values = append(values, options[p])
        }
    }
    return strings.Join(values, ",")
}
func ParseMicrosoftStoreForBusinessPortalSelectionOptions(v string) (any, error) {
    var result MicrosoftStoreForBusinessPortalSelectionOptions
    values := strings.Split(v, ",")
    for _, str := range values {
        switch str {
            case "none":
                result |= NONE_MICROSOFTSTOREFORBUSINESSPORTALSELECTIONOPTIONS
            case "companyPortal":
                result |= COMPANYPORTAL_MICROSOFTSTOREFORBUSINESSPORTALSELECTIONOPTIONS
            case "privateStore":
                result |= PRIVATESTORE_MICROSOFTSTOREFORBUSINESSPORTALSELECTIONOPTIONS
            default:
                return 0, errors.New("Unknown MicrosoftStoreForBusinessPortalSelectionOptions value: " + v)
        }
    }
    return &result, nil
}
func SerializeMicrosoftStoreForBusinessPortalSelectionOptions(values []MicrosoftStoreForBusinessPortalSelectionOptions) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i MicrosoftStoreForBusinessPortalSelectionOptions) isMultiValue() bool {
    return true
}
