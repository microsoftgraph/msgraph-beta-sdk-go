package models
import (
    "errors"
)
// The Enum to specify the Office365 ProductIds that represent the Office365 Suite SKUs.
type OfficeProductId int

const (
    O365PROPLUSRETAIL_OFFICEPRODUCTID OfficeProductId = iota
    O365BUSINESSRETAIL_OFFICEPRODUCTID
    VISIOPRORETAIL_OFFICEPRODUCTID
    PROJECTPRORETAIL_OFFICEPRODUCTID
)

func (i OfficeProductId) String() string {
    return []string{"o365ProPlusRetail", "o365BusinessRetail", "visioProRetail", "projectProRetail"}[i]
}
func ParseOfficeProductId(v string) (any, error) {
    result := O365PROPLUSRETAIL_OFFICEPRODUCTID
    switch v {
        case "o365ProPlusRetail":
            result = O365PROPLUSRETAIL_OFFICEPRODUCTID
        case "o365BusinessRetail":
            result = O365BUSINESSRETAIL_OFFICEPRODUCTID
        case "visioProRetail":
            result = VISIOPRORETAIL_OFFICEPRODUCTID
        case "projectProRetail":
            result = PROJECTPRORETAIL_OFFICEPRODUCTID
        default:
            return 0, errors.New("Unknown OfficeProductId value: " + v)
    }
    return &result, nil
}
func SerializeOfficeProductId(values []OfficeProductId) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i OfficeProductId) isMultiValue() bool {
    return false
}
