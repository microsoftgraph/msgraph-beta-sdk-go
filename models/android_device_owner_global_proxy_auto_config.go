package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AndroidDeviceOwnerGlobalProxyAutoConfig android Device Owner Global Proxy Auto Config.
type AndroidDeviceOwnerGlobalProxyAutoConfig struct {
    AndroidDeviceOwnerGlobalProxy
    // The OdataType property
    OdataType *string
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
        err = writer.WriteStringValue("proxyAutoConfigURL", m.GetProxyAutoConfigURL())
        if err != nil {
            return err
        }
    }
    return nil
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
    GetProxyAutoConfigURL()(*string)
    SetProxyAutoConfigURL(value *string)()
}
