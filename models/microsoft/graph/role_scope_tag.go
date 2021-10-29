package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type RoleScopeTag struct {
    Entity
    // The list of assignments for this Role Scope Tag.
    assignments []RoleScopeTagAutoAssignment;
    // Description of the Role Scope Tag.
    description *string;
    // The display or friendly name of the Role Scope Tag.
    displayName *string;
    // Description of the Role Scope Tag. This property is read-only.
    isBuiltIn *bool;
}
// Instantiates a new roleScopeTag and sets the default values.
func NewRoleScopeTag()(*RoleScopeTag) {
    m := &RoleScopeTag{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the assignments property value. The list of assignments for this Role Scope Tag.
func (m *RoleScopeTag) GetAssignments()([]RoleScopeTagAutoAssignment) {
    if m == nil {
        return nil
    } else {
        return m.assignments
    }
}
// Gets the description property value. Description of the Role Scope Tag.
func (m *RoleScopeTag) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// Gets the displayName property value. The display or friendly name of the Role Scope Tag.
func (m *RoleScopeTag) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the isBuiltIn property value. Description of the Role Scope Tag. This property is read-only.
func (m *RoleScopeTag) GetIsBuiltIn()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isBuiltIn
    }
}
// The deserialization information for the current model
func (m *RoleScopeTag) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["assignments"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewRoleScopeTagAutoAssignment() })
        if err != nil {
            return err
        }
        res := make([]RoleScopeTagAutoAssignment, len(val))
        for i, v := range val {
            res[i] = *(v.(*RoleScopeTagAutoAssignment))
        }
        m.SetAssignments(res)
        return nil
    }
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDescription(val)
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDisplayName(val)
        return nil
    }
    res["isBuiltIn"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsBuiltIn(val)
        return nil
    }
    return res
}
func (m *RoleScopeTag) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *RoleScopeTag) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAssignments()))
        for i, v := range m.GetAssignments() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("assignments", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isBuiltIn", m.GetIsBuiltIn())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the assignments property value. The list of assignments for this Role Scope Tag.
// Parameters:
//  - value : Value to set for the assignments property.
func (m *RoleScopeTag) SetAssignments(value []RoleScopeTagAutoAssignment)() {
    m.assignments = value
}
// Sets the description property value. Description of the Role Scope Tag.
// Parameters:
//  - value : Value to set for the description property.
func (m *RoleScopeTag) SetDescription(value *string)() {
    m.description = value
}
// Sets the displayName property value. The display or friendly name of the Role Scope Tag.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *RoleScopeTag) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the isBuiltIn property value. Description of the Role Scope Tag. This property is read-only.
// Parameters:
//  - value : Value to set for the isBuiltIn property.
func (m *RoleScopeTag) SetIsBuiltIn(value *bool)() {
    m.isBuiltIn = value
}
