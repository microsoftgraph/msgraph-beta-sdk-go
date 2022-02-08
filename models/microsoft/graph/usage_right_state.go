package graph
import (
    "strings"
    "errors"
)
// 
type UsageRightState int

const (
    ACTIVE_USAGERIGHTSTATE UsageRightState = iota
    INACTIVE_USAGERIGHTSTATE
    WARNING_USAGERIGHTSTATE
    SUSPENDED_USAGERIGHTSTATE
    UNKNOWNFUTUREVALUE_USAGERIGHTSTATE
)

func (i UsageRightState) String() string {
    return []string{"ACTIVE", "INACTIVE", "WARNING", "SUSPENDED", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseUsageRightState(v string) (interface{}, error) {
    result := ACTIVE_USAGERIGHTSTATE
    switch strings.ToUpper(v) {
        case "ACTIVE":
            result = ACTIVE_USAGERIGHTSTATE
        case "INACTIVE":
            result = INACTIVE_USAGERIGHTSTATE
        case "WARNING":
            result = WARNING_USAGERIGHTSTATE
        case "SUSPENDED":
            result = SUSPENDED_USAGERIGHTSTATE
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_USAGERIGHTSTATE
        default:
            return 0, errors.New("Unknown UsageRightState value: " + v)
    }
    return &result, nil
}
func SerializeUsageRightState(values []UsageRightState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
