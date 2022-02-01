package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// PrivilegedRole 
type PrivilegedRole struct {
    Entity
    // The assignments for this role. Read-only. Nullable.
    assignments []PrivilegedRoleAssignment;
    // Role name.
    name *string;
    // The settings for this role. Read-only. Nullable.
    settings *PrivilegedRoleSettings;
    // The summary information for this role. Read-only. Nullable.
    summary *PrivilegedRoleSummary;
}
// NewPrivilegedRole instantiates a new privilegedRole and sets the default values.
func NewPrivilegedRole()(*PrivilegedRole) {
    m := &PrivilegedRole{
        Entity: *NewEntity(),
    }
    return m
}
// GetAssignments gets the assignments property value. The assignments for this role. Read-only. Nullable.
func (m *PrivilegedRole) GetAssignments()([]PrivilegedRoleAssignment) {
    if m == nil {
        return nil
    } else {
        return m.assignments
    }
}
// GetName gets the name property value. Role name.
func (m *PrivilegedRole) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
// GetSettings gets the settings property value. The settings for this role. Read-only. Nullable.
func (m *PrivilegedRole) GetSettings()(*PrivilegedRoleSettings) {
    if m == nil {
        return nil
    } else {
        return m.settings
    }
}
// GetSummary gets the summary property value. The summary information for this role. Read-only. Nullable.
func (m *PrivilegedRole) GetSummary()(*PrivilegedRoleSummary) {
    if m == nil {
        return nil
    } else {
        return m.summary
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PrivilegedRole) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["assignments"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPrivilegedRoleAssignment() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PrivilegedRoleAssignment, len(val))
            for i, v := range val {
                res[i] = *(v.(*PrivilegedRoleAssignment))
            }
            m.SetAssignments(res)
        }
        return nil
    }
    res["name"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["settings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPrivilegedRoleSettings() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettings(val.(*PrivilegedRoleSettings))
        }
        return nil
    }
    res["summary"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPrivilegedRoleSummary() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSummary(val.(*PrivilegedRoleSummary))
        }
        return nil
    }
    return res
}
func (m *PrivilegedRole) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *PrivilegedRole) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
        err = writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("settings", m.GetSettings())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("summary", m.GetSummary())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAssignments sets the assignments property value. The assignments for this role. Read-only. Nullable.
func (m *PrivilegedRole) SetAssignments(value []PrivilegedRoleAssignment)() {
    if m != nil {
        m.assignments = value
    }
}
// SetName sets the name property value. Role name.
func (m *PrivilegedRole) SetName(value *string)() {
    if m != nil {
        m.name = value
    }
}
// SetSettings sets the settings property value. The settings for this role. Read-only. Nullable.
func (m *PrivilegedRole) SetSettings(value *PrivilegedRoleSettings)() {
    if m != nil {
        m.settings = value
    }
}
// SetSummary sets the summary property value. The summary information for this role. Read-only. Nullable.
func (m *PrivilegedRole) SetSummary(value *PrivilegedRoleSummary)() {
    if m != nil {
        m.summary = value
    }
}
