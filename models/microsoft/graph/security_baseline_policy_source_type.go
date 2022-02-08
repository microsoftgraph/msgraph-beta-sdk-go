package graph
import (
    "strings"
    "errors"
)
// 
type SecurityBaselinePolicySourceType int

const (
    DEVICECONFIGURATION_SECURITYBASELINEPOLICYSOURCETYPE SecurityBaselinePolicySourceType = iota
    DEVICEINTENT_SECURITYBASELINEPOLICYSOURCETYPE
)

func (i SecurityBaselinePolicySourceType) String() string {
    return []string{"DEVICECONFIGURATION", "DEVICEINTENT"}[i]
}
func ParseSecurityBaselinePolicySourceType(v string) (interface{}, error) {
    result := DEVICECONFIGURATION_SECURITYBASELINEPOLICYSOURCETYPE
    switch strings.ToUpper(v) {
        case "DEVICECONFIGURATION":
            result = DEVICECONFIGURATION_SECURITYBASELINEPOLICYSOURCETYPE
        case "DEVICEINTENT":
            result = DEVICEINTENT_SECURITYBASELINEPOLICYSOURCETYPE
        default:
            return 0, errors.New("Unknown SecurityBaselinePolicySourceType value: " + v)
    }
    return &result, nil
}
func SerializeSecurityBaselinePolicySourceType(values []SecurityBaselinePolicySourceType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
