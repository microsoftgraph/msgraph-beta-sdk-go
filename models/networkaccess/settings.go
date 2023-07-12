package networkaccess

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// Settings 
type Settings struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entity
    // The OdataType property
    OdataType *string
}
// NewSettings instantiates a new settings and sets the default values.
func NewSettings()(*Settings) {
    m := &Settings{
        Entity: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewEntity(),
    }
    return m
}
// CreateSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSettings(), nil
}
// GetConditionalAccess gets the conditionalAccess property value. The conditionalAccess property
func (m *Settings) GetConditionalAccess()(ConditionalAccessSettingsable) {
    val, err := m.GetBackingStore().Get("conditionalAccess")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(ConditionalAccessSettingsable)
    }
    return nil
}
// GetCrossTenantAccess gets the crossTenantAccess property value. The crossTenantAccess property
func (m *Settings) GetCrossTenantAccess()(CrossTenantAccessSettingsable) {
    val, err := m.GetBackingStore().Get("crossTenantAccess")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(CrossTenantAccessSettingsable)
    }
    return nil
}
// GetEnrichedAuditLogs gets the enrichedAuditLogs property value. The enrichedAuditLogs property
func (m *Settings) GetEnrichedAuditLogs()(EnrichedAuditLogsable) {
    val, err := m.GetBackingStore().Get("enrichedAuditLogs")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(EnrichedAuditLogsable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Settings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["conditionalAccess"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateConditionalAccessSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConditionalAccess(val.(ConditionalAccessSettingsable))
        }
        return nil
    }
    res["crossTenantAccess"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCrossTenantAccessSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCrossTenantAccess(val.(CrossTenantAccessSettingsable))
        }
        return nil
    }
    res["enrichedAuditLogs"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEnrichedAuditLogsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnrichedAuditLogs(val.(EnrichedAuditLogsable))
        }
        return nil
    }
    res["forwardingOptions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateForwardingOptionsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetForwardingOptions(val.(ForwardingOptionsable))
        }
        return nil
    }
    return res
}
// GetForwardingOptions gets the forwardingOptions property value. The forwardingOptions property
func (m *Settings) GetForwardingOptions()(ForwardingOptionsable) {
    val, err := m.GetBackingStore().Get("forwardingOptions")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(ForwardingOptionsable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *Settings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("conditionalAccess", m.GetConditionalAccess())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("crossTenantAccess", m.GetCrossTenantAccess())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("enrichedAuditLogs", m.GetEnrichedAuditLogs())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("forwardingOptions", m.GetForwardingOptions())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetConditionalAccess sets the conditionalAccess property value. The conditionalAccess property
func (m *Settings) SetConditionalAccess(value ConditionalAccessSettingsable)() {
    err := m.GetBackingStore().Set("conditionalAccess", value)
    if err != nil {
        panic(err)
    }
}
// SetCrossTenantAccess sets the crossTenantAccess property value. The crossTenantAccess property
func (m *Settings) SetCrossTenantAccess(value CrossTenantAccessSettingsable)() {
    err := m.GetBackingStore().Set("crossTenantAccess", value)
    if err != nil {
        panic(err)
    }
}
// SetEnrichedAuditLogs sets the enrichedAuditLogs property value. The enrichedAuditLogs property
func (m *Settings) SetEnrichedAuditLogs(value EnrichedAuditLogsable)() {
    err := m.GetBackingStore().Set("enrichedAuditLogs", value)
    if err != nil {
        panic(err)
    }
}
// SetForwardingOptions sets the forwardingOptions property value. The forwardingOptions property
func (m *Settings) SetForwardingOptions(value ForwardingOptionsable)() {
    err := m.GetBackingStore().Set("forwardingOptions", value)
    if err != nil {
        panic(err)
    }
}
// Settingsable 
type Settingsable interface {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetConditionalAccess()(ConditionalAccessSettingsable)
    GetCrossTenantAccess()(CrossTenantAccessSettingsable)
    GetEnrichedAuditLogs()(EnrichedAuditLogsable)
    GetForwardingOptions()(ForwardingOptionsable)
    SetConditionalAccess(value ConditionalAccessSettingsable)()
    SetCrossTenantAccess(value CrossTenantAccessSettingsable)()
    SetEnrichedAuditLogs(value EnrichedAuditLogsable)()
    SetForwardingOptions(value ForwardingOptionsable)()
}
