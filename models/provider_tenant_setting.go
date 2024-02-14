package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ProviderTenantSetting struct {
    Entity
}
// NewProviderTenantSetting instantiates a new ProviderTenantSetting and sets the default values.
func NewProviderTenantSetting()(*ProviderTenantSetting) {
    m := &ProviderTenantSetting{
        Entity: *NewEntity(),
    }
    return m
}
// CreateProviderTenantSettingFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateProviderTenantSettingFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewProviderTenantSetting(), nil
}
// GetAzureTenantId gets the azureTenantId property value. The azureTenantId property
// returns a *string when successful
func (m *ProviderTenantSetting) GetAzureTenantId()(*string) {
    val, err := m.GetBackingStore().Get("azureTenantId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetEnabled gets the enabled property value. The enabled property
// returns a *bool when successful
func (m *ProviderTenantSetting) GetEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("enabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ProviderTenantSetting) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["azureTenantId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAzureTenantId(val)
        }
        return nil
    }
    res["enabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnabled(val)
        }
        return nil
    }
    res["lastModifiedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedDateTime(val)
        }
        return nil
    }
    res["provider"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProvider(val)
        }
        return nil
    }
    res["vendor"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVendorEscaped(val)
        }
        return nil
    }
    return res
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The lastModifiedDateTime property
// returns a *Time when successful
func (m *ProviderTenantSetting) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("lastModifiedDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetProvider gets the provider property value. The provider property
// returns a *string when successful
func (m *ProviderTenantSetting) GetProvider()(*string) {
    val, err := m.GetBackingStore().Get("provider")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetVendorEscaped gets the vendor property value. The vendor property
// returns a *string when successful
func (m *ProviderTenantSetting) GetVendorEscaped()(*string) {
    val, err := m.GetBackingStore().Get("vendorEscaped")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *ProviderTenantSetting) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("azureTenantId", m.GetAzureTenantId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("enabled", m.GetEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("provider", m.GetProvider())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("vendor", m.GetVendorEscaped())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAzureTenantId sets the azureTenantId property value. The azureTenantId property
func (m *ProviderTenantSetting) SetAzureTenantId(value *string)() {
    err := m.GetBackingStore().Set("azureTenantId", value)
    if err != nil {
        panic(err)
    }
}
// SetEnabled sets the enabled property value. The enabled property
func (m *ProviderTenantSetting) SetEnabled(value *bool)() {
    err := m.GetBackingStore().Set("enabled", value)
    if err != nil {
        panic(err)
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The lastModifiedDateTime property
func (m *ProviderTenantSetting) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("lastModifiedDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetProvider sets the provider property value. The provider property
func (m *ProviderTenantSetting) SetProvider(value *string)() {
    err := m.GetBackingStore().Set("provider", value)
    if err != nil {
        panic(err)
    }
}
// SetVendorEscaped sets the vendor property value. The vendor property
func (m *ProviderTenantSetting) SetVendorEscaped(value *string)() {
    err := m.GetBackingStore().Set("vendorEscaped", value)
    if err != nil {
        panic(err)
    }
}
type ProviderTenantSettingable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAzureTenantId()(*string)
    GetEnabled()(*bool)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetProvider()(*string)
    GetVendorEscaped()(*string)
    SetAzureTenantId(value *string)()
    SetEnabled(value *bool)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetProvider(value *string)()
    SetVendorEscaped(value *string)()
}
