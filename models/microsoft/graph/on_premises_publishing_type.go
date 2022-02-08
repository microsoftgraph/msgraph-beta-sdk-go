package graph
import (
    "strings"
    "errors"
)
// 
type OnPremisesPublishingType int

const (
    APPLICATIONPROXY_ONPREMISESPUBLISHINGTYPE OnPremisesPublishingType = iota
    EXCHANGEONLINE_ONPREMISESPUBLISHINGTYPE
    AUTHENTICATION_ONPREMISESPUBLISHINGTYPE
    PROVISIONING_ONPREMISESPUBLISHINGTYPE
    INTUNEPFX_ONPREMISESPUBLISHINGTYPE
    OFLINEDOMAINJOIN_ONPREMISESPUBLISHINGTYPE
    UNKNOWNFUTUREVALUE_ONPREMISESPUBLISHINGTYPE
)

func (i OnPremisesPublishingType) String() string {
    return []string{"APPLICATIONPROXY", "EXCHANGEONLINE", "AUTHENTICATION", "PROVISIONING", "INTUNEPFX", "OFLINEDOMAINJOIN", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseOnPremisesPublishingType(v string) (interface{}, error) {
    result := APPLICATIONPROXY_ONPREMISESPUBLISHINGTYPE
    switch strings.ToUpper(v) {
        case "APPLICATIONPROXY":
            result = APPLICATIONPROXY_ONPREMISESPUBLISHINGTYPE
        case "EXCHANGEONLINE":
            result = EXCHANGEONLINE_ONPREMISESPUBLISHINGTYPE
        case "AUTHENTICATION":
            result = AUTHENTICATION_ONPREMISESPUBLISHINGTYPE
        case "PROVISIONING":
            result = PROVISIONING_ONPREMISESPUBLISHINGTYPE
        case "INTUNEPFX":
            result = INTUNEPFX_ONPREMISESPUBLISHINGTYPE
        case "OFLINEDOMAINJOIN":
            result = OFLINEDOMAINJOIN_ONPREMISESPUBLISHINGTYPE
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_ONPREMISESPUBLISHINGTYPE
        default:
            return 0, errors.New("Unknown OnPremisesPublishingType value: " + v)
    }
    return &result, nil
}
func SerializeOnPremisesPublishingType(values []OnPremisesPublishingType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
