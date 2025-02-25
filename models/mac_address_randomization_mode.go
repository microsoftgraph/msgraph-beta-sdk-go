package models
// An enum representing the possible values of Android MAC Address Randomization Mode.
type MacAddressRandomizationMode int

const (
    // Indicates the Wi-Fi framework to automatically decide the MAC randomization strategy. This can either be persistent or non-persistent randomly generated MAC addresses which are used while connecting to the network. In case of Persistent randomization, android generates a persistent randomized MAC address based on the parameters of the network profile. This MAC address remains the same until factory reset. On the other hand under the non-persistent randomization type, which is used for some networks in Android 12 or higher, the Wi-Fi module re-randomizes the MAC address at the start of every connection or the framework uses the existing randomized MAC address to connect to the network. More info: https://source.android.com/docs/core/connect/wifi-mac-randomization-behavior#types
    AUTOMATIC_MACADDRESSRANDOMIZATIONMODE MacAddressRandomizationMode = iota
    // Indicates MAC randomization is disabled and the factory MAC address is used when connecting to the internet.
    HARDWARE_MACADDRESSRANDOMIZATIONMODE
    // Evolvable enumeration sentinel value. Do not use.
    UNKNOWNFUTUREVALUE_MACADDRESSRANDOMIZATIONMODE
)

func (i MacAddressRandomizationMode) String() string {
    return []string{"automatic", "hardware", "unknownFutureValue"}[i]
}
func ParseMacAddressRandomizationMode(v string) (any, error) {
    result := AUTOMATIC_MACADDRESSRANDOMIZATIONMODE
    switch v {
        case "automatic":
            result = AUTOMATIC_MACADDRESSRANDOMIZATIONMODE
        case "hardware":
            result = HARDWARE_MACADDRESSRANDOMIZATIONMODE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_MACADDRESSRANDOMIZATIONMODE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeMacAddressRandomizationMode(values []MacAddressRandomizationMode) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i MacAddressRandomizationMode) isMultiValue() bool {
    return false
}
