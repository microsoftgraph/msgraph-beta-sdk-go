package models
import (
    "errors"
    "strings"
)
// Contains properties for detection operator.
type Win32LobAppDetectionOperator int

const (
    // Not configured.
    NOTCONFIGURED_WIN32LOBAPPDETECTIONOPERATOR Win32LobAppDetectionOperator = iota
    // Equal operator.
    EQUAL_WIN32LOBAPPDETECTIONOPERATOR
    // Not equal operator.
    NOTEQUAL_WIN32LOBAPPDETECTIONOPERATOR
    // Greater than operator.
    GREATERTHAN_WIN32LOBAPPDETECTIONOPERATOR
    // Greater than or equal operator.
    GREATERTHANOREQUAL_WIN32LOBAPPDETECTIONOPERATOR
    // Less than operator.
    LESSTHAN_WIN32LOBAPPDETECTIONOPERATOR
    // Less than or equal operator.
    LESSTHANOREQUAL_WIN32LOBAPPDETECTIONOPERATOR
)

func (i Win32LobAppDetectionOperator) String() string {
    var values []string
    for p := Win32LobAppDetectionOperator(1); p <= LESSTHANOREQUAL_WIN32LOBAPPDETECTIONOPERATOR; p <<= 1 {
        if i&p == p {
            values = append(values, []string{"notConfigured", "equal", "notEqual", "greaterThan", "greaterThanOrEqual", "lessThan", "lessThanOrEqual"}[p])
        }
    }
    return strings.Join(values, ",")
}
func ParseWin32LobAppDetectionOperator(v string) (any, error) {
    var result Win32LobAppDetectionOperator
    values := strings.Split(v, ",")
    for _, str := range values {
        switch str {
            case "notConfigured":
                result |= NOTCONFIGURED_WIN32LOBAPPDETECTIONOPERATOR
            case "equal":
                result |= EQUAL_WIN32LOBAPPDETECTIONOPERATOR
            case "notEqual":
                result |= NOTEQUAL_WIN32LOBAPPDETECTIONOPERATOR
            case "greaterThan":
                result |= GREATERTHAN_WIN32LOBAPPDETECTIONOPERATOR
            case "greaterThanOrEqual":
                result |= GREATERTHANOREQUAL_WIN32LOBAPPDETECTIONOPERATOR
            case "lessThan":
                result |= LESSTHAN_WIN32LOBAPPDETECTIONOPERATOR
            case "lessThanOrEqual":
                result |= LESSTHANOREQUAL_WIN32LOBAPPDETECTIONOPERATOR
            default:
                return 0, errors.New("Unknown Win32LobAppDetectionOperator value: " + v)
        }
    }
    return &result, nil
}
func SerializeWin32LobAppDetectionOperator(values []Win32LobAppDetectionOperator) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i Win32LobAppDetectionOperator) isMultiValue() bool {
    return true
}
