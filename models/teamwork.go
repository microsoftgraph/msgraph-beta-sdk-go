package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Teamwork 
type Teamwork struct {
    Entity
}
// NewTeamwork instantiates a new teamwork and sets the default values.
func NewTeamwork()(*Teamwork) {
    m := &Teamwork{
        Entity: *NewEntity(),
    }
    return m
}
// CreateTeamworkFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTeamworkFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTeamwork(), nil
}
// GetDeletedChats gets the deletedChats property value. The deletedChats property
func (m *Teamwork) GetDeletedChats()([]DeletedChatable) {
    val, err := m.GetBackingStore().Get("deletedChats")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]DeletedChatable)
    }
    return nil
}
// GetDeletedTeams gets the deletedTeams property value. A collection of deleted teams.
func (m *Teamwork) GetDeletedTeams()([]DeletedTeamable) {
    val, err := m.GetBackingStore().Get("deletedTeams")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]DeletedTeamable)
    }
    return nil
}
// GetDevices gets the devices property value. The Teams devices provisioned for the tenant.
func (m *Teamwork) GetDevices()([]TeamworkDeviceable) {
    val, err := m.GetBackingStore().Get("devices")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]TeamworkDeviceable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Teamwork) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["deletedChats"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeletedChatFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeletedChatable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(DeletedChatable)
                }
            }
            m.SetDeletedChats(res)
        }
        return nil
    }
    res["deletedTeams"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeletedTeamFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeletedTeamable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(DeletedTeamable)
                }
            }
            m.SetDeletedTeams(res)
        }
        return nil
    }
    res["devices"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateTeamworkDeviceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TeamworkDeviceable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(TeamworkDeviceable)
                }
            }
            m.SetDevices(res)
        }
        return nil
    }
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
    res["teamsAppSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeamsAppSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTeamsAppSettings(val.(TeamsAppSettingsable))
        }
        return nil
    }
    res["teamTemplates"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateTeamTemplateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TeamTemplateable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(TeamTemplateable)
                }
            }
            m.SetTeamTemplates(res)
        }
        return nil
    }
    res["workforceIntegrations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWorkforceIntegrationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WorkforceIntegrationable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(WorkforceIntegrationable)
                }
            }
            m.SetWorkforceIntegrations(res)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *Teamwork) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetTeamsAppSettings gets the teamsAppSettings property value. Represents tenant-wide settings for all Teams apps in the tenant.
func (m *Teamwork) GetTeamsAppSettings()(TeamsAppSettingsable) {
    val, err := m.GetBackingStore().Get("teamsAppSettings")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(TeamsAppSettingsable)
    }
    return nil
}
// GetTeamTemplates gets the teamTemplates property value. The templates associated with a team.
func (m *Teamwork) GetTeamTemplates()([]TeamTemplateable) {
    val, err := m.GetBackingStore().Get("teamTemplates")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]TeamTemplateable)
    }
    return nil
}
// GetWorkforceIntegrations gets the workforceIntegrations property value. A workforce integration with shifts.
func (m *Teamwork) GetWorkforceIntegrations()([]WorkforceIntegrationable) {
    val, err := m.GetBackingStore().Get("workforceIntegrations")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]WorkforceIntegrationable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *Teamwork) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetDeletedChats() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDeletedChats()))
        for i, v := range m.GetDeletedChats() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("deletedChats", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDeletedTeams() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDeletedTeams()))
        for i, v := range m.GetDeletedTeams() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("deletedTeams", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDevices() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDevices()))
        for i, v := range m.GetDevices() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("devices", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("teamsAppSettings", m.GetTeamsAppSettings())
        if err != nil {
            return err
        }
    }
    if m.GetTeamTemplates() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTeamTemplates()))
        for i, v := range m.GetTeamTemplates() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("teamTemplates", cast)
        if err != nil {
            return err
        }
    }
    if m.GetWorkforceIntegrations() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetWorkforceIntegrations()))
        for i, v := range m.GetWorkforceIntegrations() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("workforceIntegrations", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDeletedChats sets the deletedChats property value. The deletedChats property
func (m *Teamwork) SetDeletedChats(value []DeletedChatable)() {
    err := m.GetBackingStore().Set("deletedChats", value)
    if err != nil {
        panic(err)
    }
}
// SetDeletedTeams sets the deletedTeams property value. A collection of deleted teams.
func (m *Teamwork) SetDeletedTeams(value []DeletedTeamable)() {
    err := m.GetBackingStore().Set("deletedTeams", value)
    if err != nil {
        panic(err)
    }
}
// SetDevices sets the devices property value. The Teams devices provisioned for the tenant.
func (m *Teamwork) SetDevices(value []TeamworkDeviceable)() {
    err := m.GetBackingStore().Set("devices", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *Teamwork) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetTeamsAppSettings sets the teamsAppSettings property value. Represents tenant-wide settings for all Teams apps in the tenant.
func (m *Teamwork) SetTeamsAppSettings(value TeamsAppSettingsable)() {
    err := m.GetBackingStore().Set("teamsAppSettings", value)
    if err != nil {
        panic(err)
    }
}
// SetTeamTemplates sets the teamTemplates property value. The templates associated with a team.
func (m *Teamwork) SetTeamTemplates(value []TeamTemplateable)() {
    err := m.GetBackingStore().Set("teamTemplates", value)
    if err != nil {
        panic(err)
    }
}
// SetWorkforceIntegrations sets the workforceIntegrations property value. A workforce integration with shifts.
func (m *Teamwork) SetWorkforceIntegrations(value []WorkforceIntegrationable)() {
    err := m.GetBackingStore().Set("workforceIntegrations", value)
    if err != nil {
        panic(err)
    }
}
// Teamworkable 
type Teamworkable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDeletedChats()([]DeletedChatable)
    GetDeletedTeams()([]DeletedTeamable)
    GetDevices()([]TeamworkDeviceable)
    GetOdataType()(*string)
    GetTeamsAppSettings()(TeamsAppSettingsable)
    GetTeamTemplates()([]TeamTemplateable)
    GetWorkforceIntegrations()([]WorkforceIntegrationable)
    SetDeletedChats(value []DeletedChatable)()
    SetDeletedTeams(value []DeletedTeamable)()
    SetDevices(value []TeamworkDeviceable)()
    SetOdataType(value *string)()
    SetTeamsAppSettings(value TeamsAppSettingsable)()
    SetTeamTemplates(value []TeamTemplateable)()
    SetWorkforceIntegrations(value []WorkforceIntegrationable)()
}
