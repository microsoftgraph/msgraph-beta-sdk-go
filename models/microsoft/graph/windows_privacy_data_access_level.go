package graph
import (
    "strings"
    "errors"
)
// 
type WindowsPrivacyDataAccessLevel int

const (
    NOTCONFIGURED_WINDOWSPRIVACYDATAACCESSLEVEL WindowsPrivacyDataAccessLevel = iota
    FORCEALLOW_WINDOWSPRIVACYDATAACCESSLEVEL
    FORCEDENY_WINDOWSPRIVACYDATAACCESSLEVEL
    USERINCONTROL_WINDOWSPRIVACYDATAACCESSLEVEL
)

func (i WindowsPrivacyDataAccessLevel) String() string {
    return []string{"NOTCONFIGURED", "FORCEALLOW", "FORCEDENY", "USERINCONTROL"}[i]
}
func ParseWindowsPrivacyDataAccessLevel(v string) (interface{}, error) {
    result := NOTCONFIGURED_WINDOWSPRIVACYDATAACCESSLEVEL
    switch strings.ToUpper(v) {
        case "NOTCONFIGURED":
            result = NOTCONFIGURED_WINDOWSPRIVACYDATAACCESSLEVEL
        case "FORCEALLOW":
            result = FORCEALLOW_WINDOWSPRIVACYDATAACCESSLEVEL
        case "FORCEDENY":
            result = FORCEDENY_WINDOWSPRIVACYDATAACCESSLEVEL
        case "USERINCONTROL":
            result = USERINCONTROL_WINDOWSPRIVACYDATAACCESSLEVEL
        default:
            return 0, errors.New("Unknown WindowsPrivacyDataAccessLevel value: " + v)
    }
    return &result, nil
}
func SerializeWindowsPrivacyDataAccessLevel(values []WindowsPrivacyDataAccessLevel) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
