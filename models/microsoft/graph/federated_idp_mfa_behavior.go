package graph
import (
    "strings"
    "errors"
)
// 
type FederatedIdpMfaBehavior int

const (
    ACCEPTIFMFADONEBYFEDERATEDIDP_FEDERATEDIDPMFABEHAVIOR FederatedIdpMfaBehavior = iota
    ENFORCEMFABYFEDERATEDIDP_FEDERATEDIDPMFABEHAVIOR
    REJECTMFABYFEDERATEDIDP_FEDERATEDIDPMFABEHAVIOR
    UNKNOWNFUTUREVALUE_FEDERATEDIDPMFABEHAVIOR
)

func (i FederatedIdpMfaBehavior) String() string {
    return []string{"ACCEPTIFMFADONEBYFEDERATEDIDP", "ENFORCEMFABYFEDERATEDIDP", "REJECTMFABYFEDERATEDIDP", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseFederatedIdpMfaBehavior(v string) (interface{}, error) {
    result := ACCEPTIFMFADONEBYFEDERATEDIDP_FEDERATEDIDPMFABEHAVIOR
    switch strings.ToUpper(v) {
        case "ACCEPTIFMFADONEBYFEDERATEDIDP":
            result = ACCEPTIFMFADONEBYFEDERATEDIDP_FEDERATEDIDPMFABEHAVIOR
        case "ENFORCEMFABYFEDERATEDIDP":
            result = ENFORCEMFABYFEDERATEDIDP_FEDERATEDIDPMFABEHAVIOR
        case "REJECTMFABYFEDERATEDIDP":
            result = REJECTMFABYFEDERATEDIDP_FEDERATEDIDPMFABEHAVIOR
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_FEDERATEDIDPMFABEHAVIOR
        default:
            return 0, errors.New("Unknown FederatedIdpMfaBehavior value: " + v)
    }
    return &result, nil
}
func SerializeFederatedIdpMfaBehavior(values []FederatedIdpMfaBehavior) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
