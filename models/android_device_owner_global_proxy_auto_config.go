package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AndroidDeviceOwnerGlobalProxyAutoConfig android Device Owner Global Proxy.
type AndroidDeviceOwnerGlobalProxyAutoConfig struct {
    AndroidDeviceOwnerGlobalProxy
}
// NewAndroidDeviceOwnerGlobalProxyAutoConfig instantiates a new androidDeviceOwnerGlobalProxyAutoConfig and sets the default values.
func NewAndroidDeviceOwnerGlobalProxyAutoConfig()(*AndroidDeviceOwnerGlobalProxyAutoConfig) {
    m := &AndroidDeviceOwnerGlobalProxyAutoConfig{
        AndroidDeviceOwnerGlobalProxy: *NewAndroidDeviceOwnerGlobalProxy(),
    }
    odataTypeValue := "#microsoft.graph.androidDeviceOwnerGlobalProxyAutoConfig"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateAndroidDeviceOwnerGlobalProxyAutoConfigFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAndroidDeviceOwnerGlobalProxyAutoConfigFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAndroidDeviceOwnerGlobalProxyAutoConfig(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AndroidDeviceOwnerGlobalProxyAutoConfig) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AndroidDeviceOwnerGlobalProxy.GetFieldDeserializers()
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    res["proxyAutoConfigURL"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProxyAutoConfigURL(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *AndroidDeviceOwnerGlobalProxyAutoConfig) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetProxyAutoConfigURL gets the proxyAutoConfigURL property value. The proxy auto-config URL
func (m *AndroidDeviceOwnerGlobalProxyAutoConfig) GetProxyAutoConfigURL()(*string) {
    val, err := m.GetBackingStore().Get("proxyAutoConfigURL")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *AndroidDeviceOwnerGlobalProxyAutoConfig) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AndroidDeviceOwnerGlobalProxy.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("proxyAutoConfigURL", m.GetProxyAutoConfigURL())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AndroidDeviceOwnerGlobalProxyAutoConfig) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetProxyAutoConfigURL sets the proxyAutoConfigURL property value. The proxy auto-config URL
func (m *AndroidDeviceOwnerGlobalProxyAutoConfig) SetProxyAutoConfigURL(value *string)() {
    err := m.GetBackingStore().Set("proxyAutoConfigURL", value)
    if err != nil {
        panic(err)
    }
}
// AndroidDeviceOwnerGlobalProxyAutoConfigable 
type AndroidDeviceOwnerGlobalProxyAutoConfigable interface {
    AndroidDeviceOwnerGlobalProxyable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetOdataType()(*string)
    GetProxyAutoConfigURL()(*string)
    SetOdataType(value *string)()
    SetProxyAutoConfigURL(value *string)()
}
