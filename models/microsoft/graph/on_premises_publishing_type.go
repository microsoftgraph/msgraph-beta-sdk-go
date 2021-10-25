package graph
import (
    "strings"
    "errors"
)
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
    switch strings.ToUpper(v) {
        case "APPLICATIONPROXY":
            return APPLICATIONPROXY_ONPREMISESPUBLISHINGTYPE, nil
        case "EXCHANGEONLINE":
            return EXCHANGEONLINE_ONPREMISESPUBLISHINGTYPE, nil
        case "AUTHENTICATION":
            return AUTHENTICATION_ONPREMISESPUBLISHINGTYPE, nil
        case "PROVISIONING":
            return PROVISIONING_ONPREMISESPUBLISHINGTYPE, nil
        case "INTUNEPFX":
            return INTUNEPFX_ONPREMISESPUBLISHINGTYPE, nil
        case "OFLINEDOMAINJOIN":
            return OFLINEDOMAINJOIN_ONPREMISESPUBLISHINGTYPE, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_ONPREMISESPUBLISHINGTYPE, nil
    }
    return 0, errors.New("Unknown OnPremisesPublishingType value: " + v)
}
func SerializeOnPremisesPublishingType(values []OnPremisesPublishingType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
