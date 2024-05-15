package models
import (
    "math"
    "strings"
)
type ConditionalAccessAudienceReason int

const (
    NONE_CONDITIONALACCESSAUDIENCEREASON = 1
    RESOURCELESSREQUEST_CONDITIONALACCESSAUDIENCEREASON = 2
    CONFIDENTIALCLIENTIDTOKEN_CONDITIONALACCESSAUDIENCEREASON = 4
    CONFIDENTIALCLIENTNONIDTOKEN_CONDITIONALACCESSAUDIENCEREASON = 8
    RESOURCEMAPPING_CONDITIONALACCESSAUDIENCEREASON = 16
    RESOURCEMAPPINGDEFAULT_CONDITIONALACCESSAUDIENCEREASON = 32
    SCOPEMAPPING_CONDITIONALACCESSAUDIENCEREASON = 64
    SCOPEMAPPINGDEFAULT_CONDITIONALACCESSAUDIENCEREASON = 128
    DELEGATEDSCOPE_CONDITIONALACCESSAUDIENCEREASON = 256
    FIRSTPARTYRESOURCEDEFAULT_CONDITIONALACCESSAUDIENCEREASON = 512
    THIRDPARTYRESOURCEDEFAULT_CONDITIONALACCESSAUDIENCEREASON = 1024
    UNKNOWNFUTUREVALUE_CONDITIONALACCESSAUDIENCEREASON = 2048
)

func (i ConditionalAccessAudienceReason) String() string {
    var values []string
    options := []string{"none", "resourcelessRequest", "confidentialClientIdToken", "confidentialClientNonIdToken", "resourceMapping", "resourceMappingDefault", "scopeMapping", "scopeMappingDefault", "delegatedScope", "firstPartyResourceDefault", "thirdPartyResourceDefault", "unknownFutureValue"}
    for p := 0; p < 12; p++ {
        mantis := ConditionalAccessAudienceReason(int(math.Pow(2, float64(p))))
        if i&mantis == mantis {
            values = append(values, options[p])
        }
    }
    return strings.Join(values, ",")
}
func ParseConditionalAccessAudienceReason(v string) (any, error) {
    var result ConditionalAccessAudienceReason
    values := strings.Split(v, ",")
    for _, str := range values {
        switch str {
            case "none":
                result |= NONE_CONDITIONALACCESSAUDIENCEREASON
            case "resourcelessRequest":
                result |= RESOURCELESSREQUEST_CONDITIONALACCESSAUDIENCEREASON
            case "confidentialClientIdToken":
                result |= CONFIDENTIALCLIENTIDTOKEN_CONDITIONALACCESSAUDIENCEREASON
            case "confidentialClientNonIdToken":
                result |= CONFIDENTIALCLIENTNONIDTOKEN_CONDITIONALACCESSAUDIENCEREASON
            case "resourceMapping":
                result |= RESOURCEMAPPING_CONDITIONALACCESSAUDIENCEREASON
            case "resourceMappingDefault":
                result |= RESOURCEMAPPINGDEFAULT_CONDITIONALACCESSAUDIENCEREASON
            case "scopeMapping":
                result |= SCOPEMAPPING_CONDITIONALACCESSAUDIENCEREASON
            case "scopeMappingDefault":
                result |= SCOPEMAPPINGDEFAULT_CONDITIONALACCESSAUDIENCEREASON
            case "delegatedScope":
                result |= DELEGATEDSCOPE_CONDITIONALACCESSAUDIENCEREASON
            case "firstPartyResourceDefault":
                result |= FIRSTPARTYRESOURCEDEFAULT_CONDITIONALACCESSAUDIENCEREASON
            case "thirdPartyResourceDefault":
                result |= THIRDPARTYRESOURCEDEFAULT_CONDITIONALACCESSAUDIENCEREASON
            case "unknownFutureValue":
                result |= UNKNOWNFUTUREVALUE_CONDITIONALACCESSAUDIENCEREASON
            default:
                return nil, nil
        }
    }
    return &result, nil
}
func SerializeConditionalAccessAudienceReason(values []ConditionalAccessAudienceReason) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ConditionalAccessAudienceReason) isMultiValue() bool {
    return true
}
