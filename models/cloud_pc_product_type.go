package models
type CloudPcProductType int

const (
    ENTERPRISE_CLOUDPCPRODUCTTYPE CloudPcProductType = iota
    FRONTLINE_CLOUDPCPRODUCTTYPE
    DEVBOX_CLOUDPCPRODUCTTYPE
    POWERAUTOMATE_CLOUDPCPRODUCTTYPE
    BUSINESS_CLOUDPCPRODUCTTYPE
    UNKNOWNFUTUREVALUE_CLOUDPCPRODUCTTYPE
)

func (i CloudPcProductType) String() string {
    return []string{"enterprise", "frontline", "devBox", "powerAutomate", "business", "unknownFutureValue"}[i]
}
func ParseCloudPcProductType(v string) (any, error) {
    result := ENTERPRISE_CLOUDPCPRODUCTTYPE
    switch v {
        case "enterprise":
            result = ENTERPRISE_CLOUDPCPRODUCTTYPE
        case "frontline":
            result = FRONTLINE_CLOUDPCPRODUCTTYPE
        case "devBox":
            result = DEVBOX_CLOUDPCPRODUCTTYPE
        case "powerAutomate":
            result = POWERAUTOMATE_CLOUDPCPRODUCTTYPE
        case "business":
            result = BUSINESS_CLOUDPCPRODUCTTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_CLOUDPCPRODUCTTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeCloudPcProductType(values []CloudPcProductType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CloudPcProductType) isMultiValue() bool {
    return false
}
