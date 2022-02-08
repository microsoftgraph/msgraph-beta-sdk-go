package graph
import (
    "strings"
    "errors"
)
// 
type MicrosoftStoreForBusinessPortalSelectionOptions int

const (
    NONE_MICROSOFTSTOREFORBUSINESSPORTALSELECTIONOPTIONS MicrosoftStoreForBusinessPortalSelectionOptions = iota
    COMPANYPORTAL_MICROSOFTSTOREFORBUSINESSPORTALSELECTIONOPTIONS
    PRIVATESTORE_MICROSOFTSTOREFORBUSINESSPORTALSELECTIONOPTIONS
)

func (i MicrosoftStoreForBusinessPortalSelectionOptions) String() string {
    return []string{"NONE", "COMPANYPORTAL", "PRIVATESTORE"}[i]
}
func ParseMicrosoftStoreForBusinessPortalSelectionOptions(v string) (interface{}, error) {
    result := NONE_MICROSOFTSTOREFORBUSINESSPORTALSELECTIONOPTIONS
    switch strings.ToUpper(v) {
        case "NONE":
            result = NONE_MICROSOFTSTOREFORBUSINESSPORTALSELECTIONOPTIONS
        case "COMPANYPORTAL":
            result = COMPANYPORTAL_MICROSOFTSTOREFORBUSINESSPORTALSELECTIONOPTIONS
        case "PRIVATESTORE":
            result = PRIVATESTORE_MICROSOFTSTOREFORBUSINESSPORTALSELECTIONOPTIONS
        default:
            return 0, errors.New("Unknown MicrosoftStoreForBusinessPortalSelectionOptions value: " + v)
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
