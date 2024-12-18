package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type DeviceTemplate struct {
    DirectoryObject
}
// NewDeviceTemplate instantiates a new DeviceTemplate and sets the default values.
func NewDeviceTemplate()(*DeviceTemplate) {
    m := &DeviceTemplate{
        DirectoryObject: *NewDirectoryObject(),
    }
    odataTypeValue := "#microsoft.graph.deviceTemplate"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateDeviceTemplateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDeviceTemplateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceTemplate(), nil
}
// GetDeviceAuthority gets the deviceAuthority property value. The deviceAuthority property
// returns a *string when successful
func (m *DeviceTemplate) GetDeviceAuthority()(*string) {
    val, err := m.GetBackingStore().Get("deviceAuthority")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDeviceInstances gets the deviceInstances property value. The deviceInstances property
// returns a []Deviceable when successful
func (m *DeviceTemplate) GetDeviceInstances()([]Deviceable) {
    val, err := m.GetBackingStore().Get("deviceInstances")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]Deviceable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *DeviceTemplate) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DirectoryObject.GetFieldDeserializers()
    res["deviceAuthority"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceAuthority(val)
        }
        return nil
    }
    res["deviceInstances"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Deviceable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(Deviceable)
                }
            }
            m.SetDeviceInstances(res)
        }
        return nil
    }
    res["manufacturer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManufacturer(val)
        }
        return nil
    }
    res["model"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetModel(val)
        }
        return nil
    }
    res["mutualTlsOauthConfigurationId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMutualTlsOauthConfigurationId(val)
        }
        return nil
    }
    res["mutualTlsOauthConfigurationTenantId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMutualTlsOauthConfigurationTenantId(val)
        }
        return nil
    }
    res["operatingSystem"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOperatingSystem(val)
        }
        return nil
    }
    res["owners"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDirectoryObjectFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DirectoryObjectable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(DirectoryObjectable)
                }
            }
            m.SetOwners(res)
        }
        return nil
    }
    return res
}
// GetManufacturer gets the manufacturer property value. The manufacturer property
// returns a *string when successful
func (m *DeviceTemplate) GetManufacturer()(*string) {
    val, err := m.GetBackingStore().Get("manufacturer")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetModel gets the model property value. The model property
// returns a *string when successful
func (m *DeviceTemplate) GetModel()(*string) {
    val, err := m.GetBackingStore().Get("model")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetMutualTlsOauthConfigurationId gets the mutualTlsOauthConfigurationId property value. The mutualTlsOauthConfigurationId property
// returns a *string when successful
func (m *DeviceTemplate) GetMutualTlsOauthConfigurationId()(*string) {
    val, err := m.GetBackingStore().Get("mutualTlsOauthConfigurationId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetMutualTlsOauthConfigurationTenantId gets the mutualTlsOauthConfigurationTenantId property value. The mutualTlsOauthConfigurationTenantId property
// returns a *string when successful
func (m *DeviceTemplate) GetMutualTlsOauthConfigurationTenantId()(*string) {
    val, err := m.GetBackingStore().Get("mutualTlsOauthConfigurationTenantId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetOperatingSystem gets the operatingSystem property value. The operatingSystem property
// returns a *string when successful
func (m *DeviceTemplate) GetOperatingSystem()(*string) {
    val, err := m.GetBackingStore().Get("operatingSystem")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetOwners gets the owners property value. The owners property
// returns a []DirectoryObjectable when successful
func (m *DeviceTemplate) GetOwners()([]DirectoryObjectable) {
    val, err := m.GetBackingStore().Get("owners")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]DirectoryObjectable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *DeviceTemplate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DirectoryObject.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("deviceAuthority", m.GetDeviceAuthority())
        if err != nil {
            return err
        }
    }
    if m.GetDeviceInstances() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDeviceInstances()))
        for i, v := range m.GetDeviceInstances() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("deviceInstances", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("manufacturer", m.GetManufacturer())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("model", m.GetModel())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("mutualTlsOauthConfigurationId", m.GetMutualTlsOauthConfigurationId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("mutualTlsOauthConfigurationTenantId", m.GetMutualTlsOauthConfigurationTenantId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("operatingSystem", m.GetOperatingSystem())
        if err != nil {
            return err
        }
    }
    if m.GetOwners() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetOwners()))
        for i, v := range m.GetOwners() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("owners", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDeviceAuthority sets the deviceAuthority property value. The deviceAuthority property
func (m *DeviceTemplate) SetDeviceAuthority(value *string)() {
    err := m.GetBackingStore().Set("deviceAuthority", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceInstances sets the deviceInstances property value. The deviceInstances property
func (m *DeviceTemplate) SetDeviceInstances(value []Deviceable)() {
    err := m.GetBackingStore().Set("deviceInstances", value)
    if err != nil {
        panic(err)
    }
}
// SetManufacturer sets the manufacturer property value. The manufacturer property
func (m *DeviceTemplate) SetManufacturer(value *string)() {
    err := m.GetBackingStore().Set("manufacturer", value)
    if err != nil {
        panic(err)
    }
}
// SetModel sets the model property value. The model property
func (m *DeviceTemplate) SetModel(value *string)() {
    err := m.GetBackingStore().Set("model", value)
    if err != nil {
        panic(err)
    }
}
// SetMutualTlsOauthConfigurationId sets the mutualTlsOauthConfigurationId property value. The mutualTlsOauthConfigurationId property
func (m *DeviceTemplate) SetMutualTlsOauthConfigurationId(value *string)() {
    err := m.GetBackingStore().Set("mutualTlsOauthConfigurationId", value)
    if err != nil {
        panic(err)
    }
}
// SetMutualTlsOauthConfigurationTenantId sets the mutualTlsOauthConfigurationTenantId property value. The mutualTlsOauthConfigurationTenantId property
func (m *DeviceTemplate) SetMutualTlsOauthConfigurationTenantId(value *string)() {
    err := m.GetBackingStore().Set("mutualTlsOauthConfigurationTenantId", value)
    if err != nil {
        panic(err)
    }
}
// SetOperatingSystem sets the operatingSystem property value. The operatingSystem property
func (m *DeviceTemplate) SetOperatingSystem(value *string)() {
    err := m.GetBackingStore().Set("operatingSystem", value)
    if err != nil {
        panic(err)
    }
}
// SetOwners sets the owners property value. The owners property
func (m *DeviceTemplate) SetOwners(value []DirectoryObjectable)() {
    err := m.GetBackingStore().Set("owners", value)
    if err != nil {
        panic(err)
    }
}
type DeviceTemplateable interface {
    DirectoryObjectable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDeviceAuthority()(*string)
    GetDeviceInstances()([]Deviceable)
    GetManufacturer()(*string)
    GetModel()(*string)
    GetMutualTlsOauthConfigurationId()(*string)
    GetMutualTlsOauthConfigurationTenantId()(*string)
    GetOperatingSystem()(*string)
    GetOwners()([]DirectoryObjectable)
    SetDeviceAuthority(value *string)()
    SetDeviceInstances(value []Deviceable)()
    SetManufacturer(value *string)()
    SetModel(value *string)()
    SetMutualTlsOauthConfigurationId(value *string)()
    SetMutualTlsOauthConfigurationTenantId(value *string)()
    SetOperatingSystem(value *string)()
    SetOwners(value []DirectoryObjectable)()
}
