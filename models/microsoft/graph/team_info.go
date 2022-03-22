package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// TeamInfo 
type TeamInfo struct {
    Entity
    // 
    displayName *string;
    // 
    team Teamable;
    // 
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
func CreateTeamInfoFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewTeamInfo(), nil
}
// GetDisplayName gets the displayName property value. 
func (m *TeamInfo) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TeamInfo) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["team"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeamFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTeam(val.(Teamable))
        }
        return nil
    }
    res["tenantId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
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
// GetTeam gets the team property value. 
func (m *TeamInfo) GetTeam()(Teamable) {
    if m == nil {
        return nil
    } else {
        return m.team
    }
}
// GetTenantId gets the tenantId property value. 
func (m *TeamInfo) GetTenantId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.tenantId
    }
}
// Serialize serializes information the current object
func (m *TeamInfo) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
// SetDisplayName sets the displayName property value. 
func (m *TeamInfo) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetTeam sets the team property value. 
func (m *TeamInfo) SetTeam(value Teamable)() {
    if m != nil {
        m.team = value
    }
}
// SetTenantId sets the tenantId property value. 
func (m *TeamInfo) SetTenantId(value *string)() {
    if m != nil {
        m.tenantId = value
    }
}
