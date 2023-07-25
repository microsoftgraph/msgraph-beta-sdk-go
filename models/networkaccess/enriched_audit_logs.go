package networkaccess

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// EnrichedAuditLogs 
type EnrichedAuditLogs struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entity
}
// NewEnrichedAuditLogs instantiates a new enrichedAuditLogs and sets the default values.
func NewEnrichedAuditLogs()(*EnrichedAuditLogs) {
    m := &EnrichedAuditLogs{
        Entity: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewEntity(),
    }
    return m
}
// CreateEnrichedAuditLogsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEnrichedAuditLogsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEnrichedAuditLogs(), nil
}
// GetExchange gets the exchange property value. Exchange Online enriched audit logs settings.
func (m *EnrichedAuditLogs) GetExchange()(EnrichedAuditLogsSettingsable) {
    val, err := m.GetBackingStore().Get("exchange")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(EnrichedAuditLogsSettingsable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EnrichedAuditLogs) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["exchange"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEnrichedAuditLogsSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExchange(val.(EnrichedAuditLogsSettingsable))
        }
        return nil
    }
    res["sharepoint"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEnrichedAuditLogsSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSharepoint(val.(EnrichedAuditLogsSettingsable))
        }
        return nil
    }
    res["teams"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEnrichedAuditLogsSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTeams(val.(EnrichedAuditLogsSettingsable))
        }
        return nil
    }
    return res
}
// GetSharepoint gets the sharepoint property value. SharePoint Online enriched audit logs settings.
func (m *EnrichedAuditLogs) GetSharepoint()(EnrichedAuditLogsSettingsable) {
    val, err := m.GetBackingStore().Get("sharepoint")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(EnrichedAuditLogsSettingsable)
    }
    return nil
}
// GetTeams gets the teams property value. Teams enriched audit logs settings.
func (m *EnrichedAuditLogs) GetTeams()(EnrichedAuditLogsSettingsable) {
    val, err := m.GetBackingStore().Get("teams")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(EnrichedAuditLogsSettingsable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *EnrichedAuditLogs) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("exchange", m.GetExchange())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("sharepoint", m.GetSharepoint())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("teams", m.GetTeams())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetExchange sets the exchange property value. Exchange Online enriched audit logs settings.
func (m *EnrichedAuditLogs) SetExchange(value EnrichedAuditLogsSettingsable)() {
    err := m.GetBackingStore().Set("exchange", value)
    if err != nil {
        panic(err)
    }
}
// SetSharepoint sets the sharepoint property value. SharePoint Online enriched audit logs settings.
func (m *EnrichedAuditLogs) SetSharepoint(value EnrichedAuditLogsSettingsable)() {
    err := m.GetBackingStore().Set("sharepoint", value)
    if err != nil {
        panic(err)
    }
}
// SetTeams sets the teams property value. Teams enriched audit logs settings.
func (m *EnrichedAuditLogs) SetTeams(value EnrichedAuditLogsSettingsable)() {
    err := m.GetBackingStore().Set("teams", value)
    if err != nil {
        panic(err)
    }
}
// EnrichedAuditLogsable 
type EnrichedAuditLogsable interface {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetExchange()(EnrichedAuditLogsSettingsable)
    GetSharepoint()(EnrichedAuditLogsSettingsable)
    GetTeams()(EnrichedAuditLogsSettingsable)
    SetExchange(value EnrichedAuditLogsSettingsable)()
    SetSharepoint(value EnrichedAuditLogsSettingsable)()
    SetTeams(value EnrichedAuditLogsSettingsable)()
}
