package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IosWiFiConfigurationable 
type IosWiFiConfigurationable interface {
    DeviceConfigurationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetConnectAutomatically()(*bool)
    GetConnectWhenNetworkNameIsHidden()(*bool)
    GetDisableMacAddressRandomization()(*bool)
    GetNetworkName()(*string)
    GetPreSharedKey()(*string)
    GetProxyAutomaticConfigurationUrl()(*string)
    GetProxyManualAddress()(*string)
    GetProxyManualPort()(*int32)
    GetProxySettings()(*WiFiProxySetting)
    GetSsid()(*string)
    GetType()(*string)
    GetWiFiSecurityType()(*WiFiSecurityType)
    SetConnectAutomatically(value *bool)()
    SetConnectWhenNetworkNameIsHidden(value *bool)()
    SetDisableMacAddressRandomization(value *bool)()
    SetNetworkName(value *string)()
    SetPreSharedKey(value *string)()
    SetProxyAutomaticConfigurationUrl(value *string)()
    SetProxyManualAddress(value *string)()
    SetProxyManualPort(value *int32)()
    SetProxySettings(value *WiFiProxySetting)()
    SetSsid(value *string)()
    SetType(value *string)()
    SetWiFiSecurityType(value *WiFiSecurityType)()
}
