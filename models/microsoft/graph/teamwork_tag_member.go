package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// TeamworkTagMember 
type TeamworkTagMember struct {
    Entity
    // The member's display name.
    displayName *string;
    // 
    tenantId *string;
    // 
    userId *string;
}
// NewTeamworkTagMember instantiates a new teamworkTagMember and sets the default values.
func NewTeamworkTagMember()(*TeamworkTagMember) {
    m := &TeamworkTagMember{
        Entity: *NewEntity(),
    }
    return m
}
// GetDisplayName gets the displayName property value. The member's display name.
func (m *TeamworkTagMember) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetTenantId gets the tenantId property value. 
func (m *TeamworkTagMember) GetTenantId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.tenantId
    }
}
// GetUserId gets the userId property value. 
func (m *TeamworkTagMember) GetUserId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userId
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TeamworkTagMember) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
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
    res["userId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserId(val)
        }
        return nil
    }
    return res
}
func (m *TeamworkTagMember) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *TeamworkTagMember) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
        err = writer.WriteStringValue("tenantId", m.GetTenantId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userId", m.GetUserId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDisplayName sets the displayName property value. The member's display name.
func (m *TeamworkTagMember) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetTenantId sets the tenantId property value. 
func (m *TeamworkTagMember) SetTenantId(value *string)() {
    if m != nil {
        m.tenantId = value
    }
}
// SetUserId sets the userId property value. 
func (m *TeamworkTagMember) SetUserId(value *string)() {
    if m != nil {
        m.userId = value
    }
}
