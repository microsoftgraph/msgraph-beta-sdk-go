package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type TeamworkTagMember struct {
    Entity
    // The member's display name.
    displayName *string;
    // 
    tenantId *string;
    // 
    userId *string;
}
// Instantiates a new teamworkTagMember and sets the default values.
func NewTeamworkTagMember()(*TeamworkTagMember) {
    m := &TeamworkTagMember{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the displayName property value. The member's display name.
func (m *TeamworkTagMember) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the tenantId property value. 
func (m *TeamworkTagMember) GetTenantId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.tenantId
    }
}
// Gets the userId property value. 
func (m *TeamworkTagMember) GetUserId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userId
    }
}
// The deserialization information for the current model
func (m *TeamworkTagMember) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDisplayName(val)
        return nil
    }
    res["tenantId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetTenantId(val)
        return nil
    }
    res["userId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetUserId(val)
        return nil
    }
    return res
}
func (m *TeamworkTagMember) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the displayName property value. The member's display name.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *TeamworkTagMember) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the tenantId property value. 
// Parameters:
//  - value : Value to set for the tenantId property.
func (m *TeamworkTagMember) SetTenantId(value *string)() {
    m.tenantId = value
}
// Sets the userId property value. 
// Parameters:
//  - value : Value to set for the userId property.
func (m *TeamworkTagMember) SetUserId(value *string)() {
    m.userId = value
}
