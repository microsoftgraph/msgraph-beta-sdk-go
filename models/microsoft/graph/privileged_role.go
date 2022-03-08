package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// PrivilegedRole provides operations to manage the collection of privilegedRoleAssignmentRequest entities.
type PrivilegedRole struct {
    Entity
    // The assignments for this role. Read-only. Nullable.
    assignments []PrivilegedRoleAssignmentable;
    // Role name.
    name *string;
    // The settings for this role. Read-only. Nullable.
    settings PrivilegedRoleSettingsable;
    // The summary information for this role. Read-only. Nullable.
    summary PrivilegedRoleSummaryable;
}
// NewPrivilegedRole instantiates a new privilegedRole and sets the default values.
func NewPrivilegedRole()(*PrivilegedRole) {
    m := &PrivilegedRole{
        Entity: *NewEntity(),
    }
    return m
}
// CreatePrivilegedRoleFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePrivilegedRoleFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewPrivilegedRole(), nil
}
// GetAssignments gets the assignments property value. The assignments for this role. Read-only. Nullable.
func (m *PrivilegedRole) GetAssignments()([]PrivilegedRoleAssignmentable) {
    if m == nil {
        return nil
    } else {
        return m.assignments
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PrivilegedRole) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["assignments"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePrivilegedRoleAssignmentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PrivilegedRoleAssignmentable, len(val))
            for i, v := range val {
                res[i] = v.(PrivilegedRoleAssignmentable)
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
        val, err := n.GetObjectValue(CreatePrivilegedRoleSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettings(val.(PrivilegedRoleSettingsable))
        }
        return nil
    }
    res["summary"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreatePrivilegedRoleSummaryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSummary(val.(PrivilegedRoleSummaryable))
        }
        return nil
    }
    return res
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
func (m *PrivilegedRole) GetSettings()(PrivilegedRoleSettingsable) {
    if m == nil {
        return nil
    } else {
        return m.settings
    }
}
// GetSummary gets the summary property value. The summary information for this role. Read-only. Nullable.
func (m *PrivilegedRole) GetSummary()(PrivilegedRoleSummaryable) {
    if m == nil {
        return nil
    } else {
        return m.summary
    }
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
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
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
func (m *PrivilegedRole) SetAssignments(value []PrivilegedRoleAssignmentable)() {
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
func (m *PrivilegedRole) SetSettings(value PrivilegedRoleSettingsable)() {
    if m != nil {
        m.settings = value
    }
}
// SetSummary sets the summary property value. The summary information for this role. Read-only. Nullable.
func (m *PrivilegedRole) SetSummary(value PrivilegedRoleSummaryable)() {
    if m != nil {
        m.summary = value
    }
}
