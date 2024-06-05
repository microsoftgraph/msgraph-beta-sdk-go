package models
// The available deployment modes for a managed Tunnel server. The deployment mode is determined during the deployment depending on the Tunnel containers, namely standalone or as part of a pod, and whether the containers are running in rootful or rootless mode.
type MicrosoftTunnelDeploymentMode int

const (
    // Default. Indicates that the Tunnel containers are deployed standalone and in rootful mode.
    STANDALONEROOTFUL_MICROSOFTTUNNELDEPLOYMENTMODE MicrosoftTunnelDeploymentMode = iota
    // Indicates that the Tunnel containers are deployed standalone and in rootless mode.
    STANDALONEROOTLESS_MICROSOFTTUNNELDEPLOYMENTMODE
    // Indicates that the Tunnel containers are deployed as part of a pod and in rootful mode.
    PODROOTFUL_MICROSOFTTUNNELDEPLOYMENTMODE
    // Indicates that the Tunnel containers are deployed as part of a pod and in rootless mode.
    PODROOTLESS_MICROSOFTTUNNELDEPLOYMENTMODE
    // Evolvable enumeration sentinel value. Do not use.
    UNKNOWNFUTUREVALUE_MICROSOFTTUNNELDEPLOYMENTMODE
)

func (i MicrosoftTunnelDeploymentMode) String() string {
    return []string{"standaloneRootful", "standaloneRootless", "podRootful", "podRootless", "unknownFutureValue"}[i]
}
func ParseMicrosoftTunnelDeploymentMode(v string) (any, error) {
    result := STANDALONEROOTFUL_MICROSOFTTUNNELDEPLOYMENTMODE
    switch v {
        case "standaloneRootful":
            result = STANDALONEROOTFUL_MICROSOFTTUNNELDEPLOYMENTMODE
        case "standaloneRootless":
            result = STANDALONEROOTLESS_MICROSOFTTUNNELDEPLOYMENTMODE
        case "podRootful":
            result = PODROOTFUL_MICROSOFTTUNNELDEPLOYMENTMODE
        case "podRootless":
            result = PODROOTLESS_MICROSOFTTUNNELDEPLOYMENTMODE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_MICROSOFTTUNNELDEPLOYMENTMODE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeMicrosoftTunnelDeploymentMode(values []MicrosoftTunnelDeploymentMode) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i MicrosoftTunnelDeploymentMode) isMultiValue() bool {
    return false
}
