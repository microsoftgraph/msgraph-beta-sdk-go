package graph
import (
    "strings"
    "errors"
)
type SecurityBaselinePolicySourceType int

const (
    DEVICECONFIGURATION_SECURITYBASELINEPOLICYSOURCETYPE SecurityBaselinePolicySourceType = iota
    DEVICEINTENT_SECURITYBASELINEPOLICYSOURCETYPE
)

func (i SecurityBaselinePolicySourceType) String() string {
    return []string{"DEVICECONFIGURATION", "DEVICEINTENT"}[i]
}
func ParseSecurityBaselinePolicySourceType(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "DEVICECONFIGURATION":
            return DEVICECONFIGURATION_SECURITYBASELINEPOLICYSOURCETYPE, nil
        case "DEVICEINTENT":
            return DEVICEINTENT_SECURITYBASELINEPOLICYSOURCETYPE, nil
    }
    return 0, errors.New("Unknown SecurityBaselinePolicySourceType value: " + v)
}
func SerializeSecurityBaselinePolicySourceType(values []SecurityBaselinePolicySourceType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
