package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TeamInfo 
type TeamInfo struct {
    Entity
    // The displayName property
    displayName *string;
    // The team property
    team Teamable;
    // The tenantId property
    tenantId *string;
}
// NewTeamInfo instantiates a new teamInfo and sets the default values.
func NewTeamInfo()(*TeamInfo) {
    m := &TeamInfo{
        Entity: *NewEntity(),
    }
    return m
}
// CreateTeamInfoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTeamInfoFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTeamInfo(), nil
}
// GetDisplayName gets the displayName property value. The displayName property
func (m *TeamInfo) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TeamInfo) GetFieldDeserializers()(map[string]func(interface{}, i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["displayName"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["team"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeamFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTeam(val.(Teamable))
        }
        return nil
    }
    res["tenantId"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTenantId(val)
        }
        return nil
    }
    return res
}
// GetTeam gets the team property value. The team property
func (m *TeamInfo) GetTeam()(Teamable) {
    if m == nil {
        return nil
    } else {
        return m.team
    }
}
// GetTenantId gets the tenantId property value. The tenantId property
func (m *TeamInfo) GetTenantId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.tenantId
    }
}
// Serialize serializes information the current object
func (m *TeamInfo) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("team", m.GetTeam())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("tenantId", m.GetTenantId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDisplayName sets the displayName property value. The displayName property
func (m *TeamInfo) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetTeam sets the team property value. The team property
func (m *TeamInfo) SetTeam(value Teamable)() {
    if m != nil {
        m.team = value
    }
}
// SetTenantId sets the tenantId property value. The tenantId property
func (m *TeamInfo) SetTenantId(value *string)() {
    if m != nil {
        m.tenantId = value
    }
}
