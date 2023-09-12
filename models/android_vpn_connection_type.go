package models
import (
    "errors"
)
// Android VPN connection type.
type AndroidVpnConnectionType int

const (
    // Cisco AnyConnect.
    CISCOANYCONNECT_ANDROIDVPNCONNECTIONTYPE AndroidVpnConnectionType = iota
    // Pulse Secure.
    PULSESECURE_ANDROIDVPNCONNECTIONTYPE
    // F5 Edge Client.
    F5EDGECLIENT_ANDROIDVPNCONNECTIONTYPE
    // Dell SonicWALL Mobile Connection.
    DELLSONICWALLMOBILECONNECT_ANDROIDVPNCONNECTIONTYPE
    // Check Point Capsule VPN.
    CHECKPOINTCAPSULEVPN_ANDROIDVPNCONNECTIONTYPE
    // Citrix
    CITRIX_ANDROIDVPNCONNECTIONTYPE
    // Microsoft Tunnel.
    MICROSOFTTUNNEL_ANDROIDVPNCONNECTIONTYPE
    // NetMotion Mobility.
    NETMOTIONMOBILITY_ANDROIDVPNCONNECTIONTYPE
    // Microsoft Protect.
    MICROSOFTPROTECT_ANDROIDVPNCONNECTIONTYPE
)

func (i AndroidVpnConnectionType) String() string {
    return []string{"ciscoAnyConnect", "pulseSecure", "f5EdgeClient", "dellSonicWallMobileConnect", "checkPointCapsuleVpn", "citrix", "microsoftTunnel", "netMotionMobility", "microsoftProtect"}[i]
}
func ParseAndroidVpnConnectionType(v string) (any, error) {
    result := CISCOANYCONNECT_ANDROIDVPNCONNECTIONTYPE
    switch v {
        case "ciscoAnyConnect":
            result = CISCOANYCONNECT_ANDROIDVPNCONNECTIONTYPE
        case "pulseSecure":
            result = PULSESECURE_ANDROIDVPNCONNECTIONTYPE
        case "f5EdgeClient":
            result = F5EDGECLIENT_ANDROIDVPNCONNECTIONTYPE
        case "dellSonicWallMobileConnect":
            result = DELLSONICWALLMOBILECONNECT_ANDROIDVPNCONNECTIONTYPE
        case "checkPointCapsuleVpn":
            result = CHECKPOINTCAPSULEVPN_ANDROIDVPNCONNECTIONTYPE
        case "citrix":
            result = CITRIX_ANDROIDVPNCONNECTIONTYPE
        case "microsoftTunnel":
            result = MICROSOFTTUNNEL_ANDROIDVPNCONNECTIONTYPE
        case "netMotionMobility":
            result = NETMOTIONMOBILITY_ANDROIDVPNCONNECTIONTYPE
        case "microsoftProtect":
            result = MICROSOFTPROTECT_ANDROIDVPNCONNECTIONTYPE
        default:
            return 0, errors.New("Unknown AndroidVpnConnectionType value: " + v)
    }
    return &result, nil
}
func SerializeAndroidVpnConnectionType(values []AndroidVpnConnectionType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AndroidVpnConnectionType) isMultiValue() bool {
    return false
}
