package models
import (
    "errors"
)
type RootDomains int

const (
    NONE_ROOTDOMAINS RootDomains = iota
    ALL_ROOTDOMAINS
    ALLFEDERATED_ROOTDOMAINS
    ALLMANAGED_ROOTDOMAINS
    ENUMERATED_ROOTDOMAINS
    ALLMANAGEDANDENUMERATEDFEDERATED_ROOTDOMAINS
    UNKNOWNFUTUREVALUE_ROOTDOMAINS
)

func (i RootDomains) String() string {
    return []string{"none", "all", "allFederated", "allManaged", "enumerated", "allManagedAndEnumeratedFederated", "unknownFutureValue"}[i]
}
func ParseRootDomains(v string) (any, error) {
    result := NONE_ROOTDOMAINS
    switch v {
        case "none":
            result = NONE_ROOTDOMAINS
        case "all":
            result = ALL_ROOTDOMAINS
        case "allFederated":
            result = ALLFEDERATED_ROOTDOMAINS
        case "allManaged":
            result = ALLMANAGED_ROOTDOMAINS
        case "enumerated":
            result = ENUMERATED_ROOTDOMAINS
        case "allManagedAndEnumeratedFederated":
            result = ALLMANAGEDANDENUMERATEDFEDERATED_ROOTDOMAINS
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_ROOTDOMAINS
        default:
            return 0, errors.New("Unknown RootDomains value: " + v)
    }
    return &result, nil
}
func SerializeRootDomains(values []RootDomains) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i RootDomains) isMultiValue() bool {
    return false
}
