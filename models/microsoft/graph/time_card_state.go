package graph
import (
    "strings"
    "errors"
)
// 
type TimeCardState int

const (
    CLOCKEDIN_TIMECARDSTATE TimeCardState = iota
    ONBREAK_TIMECARDSTATE
    CLOCKEDOUT_TIMECARDSTATE
    UNKNOWNFUTUREVALUE_TIMECARDSTATE
)

func (i TimeCardState) String() string {
    return []string{"CLOCKEDIN", "ONBREAK", "CLOCKEDOUT", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseTimeCardState(v string) (interface{}, error) {
    result := CLOCKEDIN_TIMECARDSTATE
    switch strings.ToUpper(v) {
        case "CLOCKEDIN":
            result = CLOCKEDIN_TIMECARDSTATE
        case "ONBREAK":
            result = ONBREAK_TIMECARDSTATE
        case "CLOCKEDOUT":
            result = CLOCKEDOUT_TIMECARDSTATE
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_TIMECARDSTATE
        default:
            return 0, errors.New("Unknown TimeCardState value: " + v)
    }
    return &result, nil
}
func SerializeTimeCardState(values []TimeCardState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
