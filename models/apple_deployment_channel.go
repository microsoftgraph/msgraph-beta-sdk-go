package models
// Indicates the channel used to deploy the configuration profile. Available choices are DeviceChannel, UserChannel
type AppleDeploymentChannel int

const (
    // Send payload down over Device Channel.
    DEVICECHANNEL_APPLEDEPLOYMENTCHANNEL AppleDeploymentChannel = iota
    // Send payload down over User Channel.
    USERCHANNEL_APPLEDEPLOYMENTCHANNEL
)

func (i AppleDeploymentChannel) String() string {
    return []string{"deviceChannel", "userChannel"}[i]
}
func ParseAppleDeploymentChannel(v string) (any, error) {
    result := DEVICECHANNEL_APPLEDEPLOYMENTCHANNEL
    switch v {
        case "deviceChannel":
            result = DEVICECHANNEL_APPLEDEPLOYMENTCHANNEL
        case "userChannel":
            result = USERCHANNEL_APPLEDEPLOYMENTCHANNEL
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeAppleDeploymentChannel(values []AppleDeploymentChannel) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AppleDeploymentChannel) isMultiValue() bool {
    return false
}
