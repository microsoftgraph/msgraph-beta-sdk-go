package models
import (
    "errors"
    "strings"
)
// Portal to which admin syncs available Microsoft Store for Business apps. This is available in the Intune Admin Console.
type MicrosoftStoreForBusinessPortalSelectionOptions int

const (
    // This option is not available for the account
    NONE_MICROSOFTSTOREFORBUSINESSPORTALSELECTIONOPTIONS MicrosoftStoreForBusinessPortalSelectionOptions = iota
    // Intune Company Portal only.
    COMPANYPORTAL_MICROSOFTSTOREFORBUSINESSPORTALSELECTIONOPTIONS
    // MSFB Private store only.
    PRIVATESTORE_MICROSOFTSTOREFORBUSINESSPORTALSELECTIONOPTIONS
)

func (i MicrosoftStoreForBusinessPortalSelectionOptions) String() string {
    var values []string
    for p := MicrosoftStoreForBusinessPortalSelectionOptions(1); p <= PRIVATESTORE_MICROSOFTSTOREFORBUSINESSPORTALSELECTIONOPTIONS; p <<= 1 {
        if i&p == p {
            values = append(values, []string{"none", "companyPortal", "privateStore"}[p])
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
