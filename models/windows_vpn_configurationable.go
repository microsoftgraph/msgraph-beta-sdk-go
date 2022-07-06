package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsVpnConfigurationable 
type WindowsVpnConfigurationable interface {
    DeviceConfigurationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetConnectionName()(*string)
    GetCustomXml()([]byte)
    GetServers()([]VpnServerable)
    GetType()(*string)
    SetConnectionName(value *string)()
    SetCustomXml(value []byte)()
    SetServers(value []VpnServerable)()
    SetType(value *string)()
}
