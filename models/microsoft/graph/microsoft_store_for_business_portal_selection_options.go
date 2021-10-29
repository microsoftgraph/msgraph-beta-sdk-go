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
    switch strings.ToUpper(v) {
        case "NONE":
            return NONE_MICROSOFTSTOREFORBUSINESSPORTALSELECTIONOPTIONS, nil
        case "COMPANYPORTAL":
            return COMPANYPORTAL_MICROSOFTSTOREFORBUSINESSPORTALSELECTIONOPTIONS, nil
        case "PRIVATESTORE":
            return PRIVATESTORE_MICROSOFTSTOREFORBUSINESSPORTALSELECTIONOPTIONS, nil
    }
    return 0, errors.New("Unknown MicrosoftStoreForBusinessPortalSelectionOptions value: " + v)
}
func SerializeMicrosoftStoreForBusinessPortalSelectionOptions(values []MicrosoftStoreForBusinessPortalSelectionOptions) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
