package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AndroidDeviceOwnerGlobalProxyDirect android Device Owner Global Proxy.
type AndroidDeviceOwnerGlobalProxyDirect struct {
    AndroidDeviceOwnerGlobalProxy
    // The OdataType property
    OdataType *string
}
// NewAndroidDeviceOwnerGlobalProxyDirect instantiates a new androidDeviceOwnerGlobalProxyDirect and sets the default values.
func NewAndroidDeviceOwnerGlobalProxyDirect()(*AndroidDeviceOwnerGlobalProxyDirect) {
    m := &AndroidDeviceOwnerGlobalProxyDirect{
        AndroidDeviceOwnerGlobalProxy: *NewAndroidDeviceOwnerGlobalProxy(),
    }
    odataTypeValue := "#microsoft.graph.androidDeviceOwnerGlobalProxyDirect"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateAndroidDeviceOwnerGlobalProxyDirectFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAndroidDeviceOwnerGlobalProxyDirectFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAndroidDeviceOwnerGlobalProxyDirect(), nil
}
// GetExcludedHosts gets the excludedHosts property value. The excluded hosts
func (m *AndroidDeviceOwnerGlobalProxyDirect) GetExcludedHosts()([]string) {
    val, err := m.GetBackingStore().Get("excludedHosts")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AndroidDeviceOwnerGlobalProxyDirect) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AndroidDeviceOwnerGlobalProxy.GetFieldDeserializers()
    res["excludedHosts"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetExcludedHosts(res)
        }
        return nil
    }
    res["host"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHost(val)
        }
        return nil
    }
    res["port"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPort(val)
        }
        return nil
    }
    return res
}
// GetHost gets the host property value. The host name
func (m *AndroidDeviceOwnerGlobalProxyDirect) GetHost()(*string) {
    val, err := m.GetBackingStore().Get("host")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPort gets the port property value. The port
func (m *AndroidDeviceOwnerGlobalProxyDirect) GetPort()(*int32) {
    val, err := m.GetBackingStore().Get("port")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// Serialize serializes information the current object
func (m *AndroidDeviceOwnerGlobalProxyDirect) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AndroidDeviceOwnerGlobalProxy.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetExcludedHosts() != nil {
        err = writer.WriteCollectionOfStringValues("excludedHosts", m.GetExcludedHosts())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("host", m.GetHost())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("port", m.GetPort())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetExcludedHosts sets the excludedHosts property value. The excluded hosts
func (m *AndroidDeviceOwnerGlobalProxyDirect) SetExcludedHosts(value []string)() {
    err := m.GetBackingStore().Set("excludedHosts", value)
    if err != nil {
        panic(err)
    }
}
// SetHost sets the host property value. The host name
func (m *AndroidDeviceOwnerGlobalProxyDirect) SetHost(value *string)() {
    err := m.GetBackingStore().Set("host", value)
    if err != nil {
        panic(err)
    }
}
// SetPort sets the port property value. The port
func (m *AndroidDeviceOwnerGlobalProxyDirect) SetPort(value *int32)() {
    err := m.GetBackingStore().Set("port", value)
    if err != nil {
        panic(err)
    }
}
// AndroidDeviceOwnerGlobalProxyDirectable 
type AndroidDeviceOwnerGlobalProxyDirectable interface {
    AndroidDeviceOwnerGlobalProxyable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetExcludedHosts()([]string)
    GetHost()(*string)
    GetPort()(*int32)
    SetExcludedHosts(value []string)()
    SetHost(value *string)()
    SetPort(value *int32)()
}
