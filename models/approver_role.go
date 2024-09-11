package models
import (
    "math"
    "strings"
)
type ApproverRole int

const (
    OWNER_APPROVERROLE = 1
    APPROVER_APPROVERROLE = 2
    UNKNOWNFUTUREVALUE_APPROVERROLE = 4
)

func (i ApproverRole) String() string {
    var values []string
    options := []string{"owner", "approver", "unknownFutureValue"}
    for p := 0; p < 3; p++ {
        mantis := ApproverRole(int(math.Pow(2, float64(p))))
        if i&mantis == mantis {
            values = append(values, options[p])
        }
    }
    return strings.Join(values, ",")
}
func ParseApproverRole(v string) (any, error) {
    var result ApproverRole
    values := strings.Split(v, ",")
    for _, str := range values {
        switch str {
            case "owner":
                result |= OWNER_APPROVERROLE
            case "approver":
                result |= APPROVER_APPROVERROLE
            case "unknownFutureValue":
                result |= UNKNOWNFUTUREVALUE_APPROVERROLE
            default:
                return nil, nil
        }
    }
    return &result, nil
}
func SerializeApproverRole(values []ApproverRole) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ApproverRole) isMultiValue() bool {
    return true
}
