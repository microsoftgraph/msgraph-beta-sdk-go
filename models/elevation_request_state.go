package models
import (
    "errors"
)
// Indicates state of elevation request
type ElevationRequestState int

const (
    // Default Value. Indicates that elevation request status is unavailable
    NONE_ELEVATIONREQUESTSTATE ElevationRequestState = iota
    // Initial state when request is submitted but no approval/denial action taken
    PENDING_ELEVATIONREQUESTSTATE
    // Indicates elevation request has been approved by IT Admin.
    APPROVED_ELEVATIONREQUESTSTATE
    // Indicates elevation request has been denied by IT Admin.
    DENIED_ELEVATIONREQUESTSTATE
    // Set to expire when Approved for is elapsed or ExpireDate is elapsed, whichever is sooner.
    EXPIRED_ELEVATIONREQUESTSTATE
    // Evolvable enumeration sentinel value. Do not use.
    UNKNOWNFUTUREVALUE_ELEVATIONREQUESTSTATE
    // Set to expire when Approved for is elapsed or ExpireDate is elapsed, whichever is sooner.
    REVOKED_ELEVATIONREQUESTSTATE
)

func (i ElevationRequestState) String() string {
    return []string{"none", "pending", "approved", "denied", "expired", "unknownFutureValue", "revoked"}[i]
}
func ParseElevationRequestState(v string) (any, error) {
    result := NONE_ELEVATIONREQUESTSTATE
    switch v {
        case "none":
            result = NONE_ELEVATIONREQUESTSTATE
        case "pending":
            result = PENDING_ELEVATIONREQUESTSTATE
        case "approved":
            result = APPROVED_ELEVATIONREQUESTSTATE
        case "denied":
            result = DENIED_ELEVATIONREQUESTSTATE
        case "expired":
            result = EXPIRED_ELEVATIONREQUESTSTATE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_ELEVATIONREQUESTSTATE
        case "revoked":
            result = REVOKED_ELEVATIONREQUESTSTATE
        default:
            return 0, errors.New("Unknown ElevationRequestState value: " + v)
    }
    return &result, nil
}
func SerializeElevationRequestState(values []ElevationRequestState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ElevationRequestState) isMultiValue() bool {
    return false
}
