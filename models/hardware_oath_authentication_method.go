package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type HardwareOathAuthenticationMethod struct {
    AuthenticationMethod
}
// NewHardwareOathAuthenticationMethod instantiates a new HardwareOathAuthenticationMethod and sets the default values.
func NewHardwareOathAuthenticationMethod()(*HardwareOathAuthenticationMethod) {
    m := &HardwareOathAuthenticationMethod{
        AuthenticationMethod: *NewAuthenticationMethod(),
    }
    odataTypeValue := "#microsoft.graph.hardwareOathAuthenticationMethod"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateHardwareOathAuthenticationMethodFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateHardwareOathAuthenticationMethodFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewHardwareOathAuthenticationMethod(), nil
}
// GetDevice gets the device property value. Exposes the hardware OATH method in the directory.
// returns a HardwareOathTokenAuthenticationMethodDeviceable when successful
func (m *HardwareOathAuthenticationMethod) GetDevice()(HardwareOathTokenAuthenticationMethodDeviceable) {
    val, err := m.GetBackingStore().Get("device")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(HardwareOathTokenAuthenticationMethodDeviceable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *HardwareOathAuthenticationMethod) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AuthenticationMethod.GetFieldDeserializers()
    res["device"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateHardwareOathTokenAuthenticationMethodDeviceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDevice(val.(HardwareOathTokenAuthenticationMethodDeviceable))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *HardwareOathAuthenticationMethod) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AuthenticationMethod.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("device", m.GetDevice())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDevice sets the device property value. Exposes the hardware OATH method in the directory.
func (m *HardwareOathAuthenticationMethod) SetDevice(value HardwareOathTokenAuthenticationMethodDeviceable)() {
    err := m.GetBackingStore().Set("device", value)
    if err != nil {
        panic(err)
    }
}
type HardwareOathAuthenticationMethodable interface {
    AuthenticationMethodable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDevice()(HardwareOathTokenAuthenticationMethodDeviceable)
    SetDevice(value HardwareOathTokenAuthenticationMethodDeviceable)()
}
