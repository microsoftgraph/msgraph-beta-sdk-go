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
    if m == nil {
        return nil
    } else {
        return m.deletedTeams
    }
}
// GetDevices gets the devices property value. The Teams devices provisioned for the tenant.
func (m *Teamwork) GetDevices()([]TeamworkDeviceable) {
    if m == nil {
        return nil
    } else {
        return m.devices
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Teamwork) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["deletedTeams"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeletedTeamFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeletedTeamable, len(val))
            for i, v := range val {
                res[i] = v.(DeletedTeamable)
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
                res[i] = v.(TeamworkDeviceable)
            }
            m.SetDevices(res)
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
    res["workforceIntegrations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWorkforceIntegrationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WorkforceIntegrationable, len(val))
            for i, v := range val {
                res[i] = v.(WorkforceIntegrationable)
            }
            m.SetWorkforceIntegrations(res)
        }
        return nil
    }
    return res
}
// GetTeamsAppSettings gets the teamsAppSettings property value. Represents tenant-wide settings for all Teams apps in the tenant.
func (m *Teamwork) GetTeamsAppSettings()(TeamsAppSettingsable) {
    if m == nil {
        return nil
    } else {
        return m.teamsAppSettings
    }
}
// GetWorkforceIntegrations gets the workforceIntegrations property value. A workforce integration with shifts.
func (m *Teamwork) GetWorkforceIntegrations()([]WorkforceIntegrationable) {
    if m == nil {
        return nil
    } else {
        return m.workforceIntegrations
    }
}
// Serialize serializes information the current object
func (m *Teamwork) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetDeletedTeams() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDeletedTeams()))
        for i, v := range m.GetDeletedTeams() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("deletedTeams", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDevices() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDevices()))
        for i, v := range m.GetDevices() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
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
    if m.GetWorkforceIntegrations() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetWorkforceIntegrations()))
        for i, v := range m.GetWorkforceIntegrations() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("workforceIntegrations", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDeletedTeams sets the deletedTeams property value. A collection of deleted teams.
func (m *Teamwork) SetDeletedTeams(value []DeletedTeamable)() {
    if m != nil {
        m.deletedTeams = value
    }
}
// SetDevices sets the devices property value. The Teams devices provisioned for the tenant.
func (m *Teamwork) SetDevices(value []TeamworkDeviceable)() {
    if m != nil {
        m.devices = value
    }
}
// SetTeamsAppSettings sets the teamsAppSettings property value. Represents tenant-wide settings for all Teams apps in the tenant.
func (m *Teamwork) SetTeamsAppSettings(value TeamsAppSettingsable)() {
    if m != nil {
        m.teamsAppSettings = value
    }
}
// SetWorkforceIntegrations sets the workforceIntegrations property value. A workforce integration with shifts.
func (m *Teamwork) SetWorkforceIntegrations(value []WorkforceIntegrationable)() {
    if m != nil {
        m.workforceIntegrations = value
    }
}
