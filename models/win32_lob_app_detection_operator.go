package models
import (
    "math"
    "strings"
)
// Contains properties for detection operator.
type Win32LobAppDetectionOperator int

const (
    // Not configured.
    NOTCONFIGURED_WIN32LOBAPPDETECTIONOPERATOR = 1
    // Equal operator.
    EQUAL_WIN32LOBAPPDETECTIONOPERATOR = 2
    // Not equal operator.
    NOTEQUAL_WIN32LOBAPPDETECTIONOPERATOR = 4
    // Greater than operator.
    GREATERTHAN_WIN32LOBAPPDETECTIONOPERATOR = 8
    // Greater than or equal operator.
    GREATERTHANOREQUAL_WIN32LOBAPPDETECTIONOPERATOR = 16
    // Less than operator.
    LESSTHAN_WIN32LOBAPPDETECTIONOPERATOR = 32
    // Less than or equal operator.
    LESSTHANOREQUAL_WIN32LOBAPPDETECTIONOPERATOR = 64
)

func (i Win32LobAppDetectionOperator) String() string {
    var values []string
    options := []string{"notConfigured", "equal", "notEqual", "greaterThan", "greaterThanOrEqual", "lessThan", "lessThanOrEqual"}
    for p := 0; p < 7; p++ {
        mantis := Win32LobAppDetectionOperator(int(math.Pow(2, float64(p))))
        if i&mantis == mantis {
            values = append(values, options[p])
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
                return nil, nil
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
