package graph
import (
    "strings"
    "errors"
)
// 
type CompanyPortalAction int

const (
    UNKNOWN_COMPANYPORTALACTION CompanyPortalAction = iota
    REMOVE_COMPANYPORTALACTION
    RESET_COMPANYPORTALACTION
)

func (i CompanyPortalAction) String() string {
    return []string{"UNKNOWN", "REMOVE", "RESET"}[i]
}
func ParseCompanyPortalAction(v string) (interface{}, error) {
    result := UNKNOWN_COMPANYPORTALACTION
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            result = UNKNOWN_COMPANYPORTALACTION
        case "REMOVE":
            result = REMOVE_COMPANYPORTALACTION
        case "RESET":
            result = RESET_COMPANYPORTALACTION
        default:
            return 0, errors.New("Unknown CompanyPortalAction value: " + v)
    }
    return &result, nil
}
func SerializeCompanyPortalAction(values []CompanyPortalAction) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
