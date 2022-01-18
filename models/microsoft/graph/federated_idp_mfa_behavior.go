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
    switch strings.ToUpper(v) {
        case "ACCEPTIFMFADONEBYFEDERATEDIDP":
            return ACCEPTIFMFADONEBYFEDERATEDIDP_FEDERATEDIDPMFABEHAVIOR, nil
        case "ENFORCEMFABYFEDERATEDIDP":
            return ENFORCEMFABYFEDERATEDIDP_FEDERATEDIDPMFABEHAVIOR, nil
        case "REJECTMFABYFEDERATEDIDP":
            return REJECTMFABYFEDERATEDIDP_FEDERATEDIDPMFABEHAVIOR, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_FEDERATEDIDPMFABEHAVIOR, nil
    }
    return 0, errors.New("Unknown FederatedIdpMfaBehavior value: " + v)
}
func SerializeFederatedIdpMfaBehavior(values []FederatedIdpMfaBehavior) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
