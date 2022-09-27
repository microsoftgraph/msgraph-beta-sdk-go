package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Teamwork 
type Teamwork struct {
    Entity
    // A collection of deleted teams.
    deletedTeams []DeletedTeamable
    // The Teams devices provisioned for the tenant.
    devices []TeamworkDeviceable
    // Represents tenant-wide settings for all Teams apps in the tenant.
    teamsAppSettings TeamsAppSettingsable
    // The templates associated with a team.
    teamTemplates []TeamTemplateable
    // A workforce integration with shifts.
    workforceIntegrations []WorkforceIntegrationable
}
// NewTeamwork instantiates a new Teamwork and sets the default values.
func NewTeamwork()(*Teamwork) {
    m := &Teamwork{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.teamwork";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateTeamworkFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTeamworkFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTeamwork(), nil
}
// GetDeletedTeams gets the deletedTeams property value. A collection of deleted teams.
func (m *Teamwork) GetDeletedTeams()([]DeletedTeamable) {
    return m.deletedTeams
}
// GetDevices gets the devices property value. The Teams devices provisioned for the tenant.
func (m *Teamwork) GetDevices()([]TeamworkDeviceable) {
    return m.devices
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Teamwork) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["deletedTeams"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateDeletedTeamFromDiscriminatorValue , m.SetDeletedTeams)
    res["devices"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateTeamworkDeviceFromDiscriminatorValue , m.SetDevices)
    res["teamsAppSettings"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateTeamsAppSettingsFromDiscriminatorValue , m.SetTeamsAppSettings)
    res["teamTemplates"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateTeamTemplateFromDiscriminatorValue , m.SetTeamTemplates)
    res["workforceIntegrations"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateWorkforceIntegrationFromDiscriminatorValue , m.SetWorkforceIntegrations)
    return res
}
// GetTeamsAppSettings gets the teamsAppSettings property value. Represents tenant-wide settings for all Teams apps in the tenant.
func (m *Teamwork) GetTeamsAppSettings()(TeamsAppSettingsable) {
    return m.teamsAppSettings
}
// GetTeamTemplates gets the teamTemplates property value. The templates associated with a team.
func (m *Teamwork) GetTeamTemplates()([]TeamTemplateable) {
    return m.teamTemplates
}
// GetWorkforceIntegrations gets the workforceIntegrations property value. A workforce integration with shifts.
func (m *Teamwork) GetWorkforceIntegrations()([]WorkforceIntegrationable) {
    return m.workforceIntegrations
}
// Serialize serializes information the current object
func (m *Teamwork) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetDeletedTeams() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetDeletedTeams())
        err = writer.WriteCollectionOfObjectValues("deletedTeams", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDevices() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetDevices())
        err = writer.WriteCollectionOfObjectValues("devices", cast)
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
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetTeamTemplates())
        err = writer.WriteCollectionOfObjectValues("teamTemplates", cast)
        if err != nil {
            return err
        }
    }
    if m.GetWorkforceIntegrations() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetWorkforceIntegrations())
        err = writer.WriteCollectionOfObjectValues("workforceIntegrations", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDeletedTeams sets the deletedTeams property value. A collection of deleted teams.
func (m *Teamwork) SetDeletedTeams(value []DeletedTeamable)() {
    m.deletedTeams = value
}
// SetDevices sets the devices property value. The Teams devices provisioned for the tenant.
func (m *Teamwork) SetDevices(value []TeamworkDeviceable)() {
    m.devices = value
}
// SetTeamsAppSettings sets the teamsAppSettings property value. Represents tenant-wide settings for all Teams apps in the tenant.
func (m *Teamwork) SetTeamsAppSettings(value TeamsAppSettingsable)() {
    m.teamsAppSettings = value
}
// SetTeamTemplates sets the teamTemplates property value. The templates associated with a team.
func (m *Teamwork) SetTeamTemplates(value []TeamTemplateable)() {
    m.teamTemplates = value
}
// SetWorkforceIntegrations sets the workforceIntegrations property value. A workforce integration with shifts.
func (m *Teamwork) SetWorkforceIntegrations(value []WorkforceIntegrationable)() {
    m.workforceIntegrations = value
}
