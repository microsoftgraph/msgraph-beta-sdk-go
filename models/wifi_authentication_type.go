package models
// Wi-Fi Authentication Type Settings.
type WifiAuthenticationType int

const (
    // None
    NONE_WIFIAUTHENTICATIONTYPE WifiAuthenticationType = iota
    // User Authentication
    USER_WIFIAUTHENTICATIONTYPE
    // Machine Authentication
    MACHINE_WIFIAUTHENTICATIONTYPE
    // Machine or User Authentication
    MACHINEORUSER_WIFIAUTHENTICATIONTYPE
    // Guest Authentication
    GUEST_WIFIAUTHENTICATIONTYPE
)

func (i WifiAuthenticationType) String() string {
    return []string{"none", "user", "machine", "machineOrUser", "guest"}[i]
}
func ParseWifiAuthenticationType(v string) (any, error) {
    result := NONE_WIFIAUTHENTICATIONTYPE
    switch v {
        case "none":
            result = NONE_WIFIAUTHENTICATIONTYPE
        case "user":
            result = USER_WIFIAUTHENTICATIONTYPE
        case "machine":
            result = MACHINE_WIFIAUTHENTICATIONTYPE
        case "machineOrUser":
            result = MACHINEORUSER_WIFIAUTHENTICATIONTYPE
        case "guest":
            result = GUEST_WIFIAUTHENTICATIONTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeWifiAuthenticationType(values []WifiAuthenticationType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i WifiAuthenticationType) isMultiValue() bool {
    return false
}
