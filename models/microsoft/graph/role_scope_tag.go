package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// RoleScopeTag 
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
// NewRoleScopeTag instantiates a new roleScopeTag and sets the default values.
func NewRoleScopeTag()(*RoleScopeTag) {
    m := &RoleScopeTag{
        Entity: *NewEntity(),
    }
    return m
}
// GetAssignments gets the assignments property value. The list of assignments for this Role Scope Tag.
func (m *RoleScopeTag) GetAssignments()([]RoleScopeTagAutoAssignment) {
    if m == nil {
        return nil
    } else {
        return m.assignments
    }
}
// GetDescription gets the description property value. Description of the Role Scope Tag.
func (m *RoleScopeTag) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetDisplayName gets the displayName property value. The display or friendly name of the Role Scope Tag.
func (m *RoleScopeTag) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetIsBuiltIn gets the isBuiltIn property value. Description of the Role Scope Tag. This property is read-only.
func (m *RoleScopeTag) GetIsBuiltIn()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isBuiltIn
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RoleScopeTag) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["assignments"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewRoleScopeTagAutoAssignment() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]RoleScopeTagAutoAssignment, len(val))
            for i, v := range val {
                res[i] = *(v.(*RoleScopeTagAutoAssignment))
            }
            m.SetAssignments(res)
        }
        return nil
    }
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
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
    res["isBuiltIn"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsBuiltIn(val)
        }
        return nil
    }
    return res
}
func (m *RoleScopeTag) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *RoleScopeTag) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAssignments() != nil {
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
// SetAssignments sets the assignments property value. The list of assignments for this Role Scope Tag.
func (m *RoleScopeTag) SetAssignments(value []RoleScopeTagAutoAssignment)() {
    if m != nil {
        m.assignments = value
    }
}
// SetDescription sets the description property value. Description of the Role Scope Tag.
func (m *RoleScopeTag) SetDescription(value *string)() {
    if m != nil {
        m.description = value
    }
}
// SetDisplayName sets the displayName property value. The display or friendly name of the Role Scope Tag.
func (m *RoleScopeTag) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetIsBuiltIn sets the isBuiltIn property value. Description of the Role Scope Tag. This property is read-only.
func (m *RoleScopeTag) SetIsBuiltIn(value *bool)() {
    if m != nil {
        m.isBuiltIn = value
    }
}
