package graph
import (
    "strings"
    "errors"
)
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
    switch strings.ToUpper(v) {
        case "ACTIVE":
            return ACTIVE_USAGERIGHTSTATE, nil
        case "INACTIVE":
            return INACTIVE_USAGERIGHTSTATE, nil
        case "WARNING":
            return WARNING_USAGERIGHTSTATE, nil
        case "SUSPENDED":
            return SUSPENDED_USAGERIGHTSTATE, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_USAGERIGHTSTATE, nil
    }
    return 0, errors.New("Unknown UsageRightState value: " + v)
}
func SerializeUsageRightState(values []UsageRightState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
