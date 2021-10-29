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
    switch strings.ToUpper(v) {
        case "NOTCONFIGURED":
            return NOTCONFIGURED_WINDOWSPRIVACYDATAACCESSLEVEL, nil
        case "FORCEALLOW":
            return FORCEALLOW_WINDOWSPRIVACYDATAACCESSLEVEL, nil
        case "FORCEDENY":
            return FORCEDENY_WINDOWSPRIVACYDATAACCESSLEVEL, nil
        case "USERINCONTROL":
            return USERINCONTROL_WINDOWSPRIVACYDATAACCESSLEVEL, nil
    }
    return 0, errors.New("Unknown WindowsPrivacyDataAccessLevel value: " + v)
}
func SerializeWindowsPrivacyDataAccessLevel(values []WindowsPrivacyDataAccessLevel) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
