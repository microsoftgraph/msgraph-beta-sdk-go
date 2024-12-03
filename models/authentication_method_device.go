package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type AuthenticationMethodDevice struct {
    Entity
}
// NewAuthenticationMethodDevice instantiates a new AuthenticationMethodDevice and sets the default values.
func NewAuthenticationMethodDevice()(*AuthenticationMethodDevice) {
    m := &AuthenticationMethodDevice{
        Entity: *NewEntity(),
    }
    return m
}
// CreateAuthenticationMethodDeviceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAuthenticationMethodDeviceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "#microsoft.graph.hardwareOathTokenAuthenticationMethodDevice":
                        return NewHardwareOathTokenAuthenticationMethodDevice(), nil
                }
            }
        }
    }
    return NewAuthenticationMethodDevice(), nil
}
// GetDisplayName gets the displayName property value. The displayName property
// returns a *string when successful
func (m *AuthenticationMethodDevice) GetDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("displayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AuthenticationMethodDevice) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["hardwareOathDevices"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateHardwareOathTokenAuthenticationMethodDeviceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]HardwareOathTokenAuthenticationMethodDeviceable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(HardwareOathTokenAuthenticationMethodDeviceable)
                }
            }
            m.SetHardwareOathDevices(res)
        }
        return nil
    }
    return res
}
// GetHardwareOathDevices gets the hardwareOathDevices property value. The hardwareOathDevices property
// returns a []HardwareOathTokenAuthenticationMethodDeviceable when successful
func (m *AuthenticationMethodDevice) GetHardwareOathDevices()([]HardwareOathTokenAuthenticationMethodDeviceable) {
    val, err := m.GetBackingStore().Get("hardwareOathDevices")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]HardwareOathTokenAuthenticationMethodDeviceable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *AuthenticationMethodDevice) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    if m.GetHardwareOathDevices() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetHardwareOathDevices()))
        for i, v := range m.GetHardwareOathDevices() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("hardwareOathDevices", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDisplayName sets the displayName property value. The displayName property
func (m *AuthenticationMethodDevice) SetDisplayName(value *string)() {
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// SetHardwareOathDevices sets the hardwareOathDevices property value. The hardwareOathDevices property
func (m *AuthenticationMethodDevice) SetHardwareOathDevices(value []HardwareOathTokenAuthenticationMethodDeviceable)() {
    err := m.GetBackingStore().Set("hardwareOathDevices", value)
    if err != nil {
        panic(err)
    }
}
type AuthenticationMethodDeviceable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDisplayName()(*string)
    GetHardwareOathDevices()([]HardwareOathTokenAuthenticationMethodDeviceable)
    SetDisplayName(value *string)()
    SetHardwareOathDevices(value []HardwareOathTokenAuthenticationMethodDeviceable)()
}
