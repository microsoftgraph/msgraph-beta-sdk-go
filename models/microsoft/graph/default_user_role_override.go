package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DefaultUserRoleOverride provides operations to manage the policyRoot singleton.
type DefaultUserRoleOverride struct {
    Entity
    // 
    isDefault *bool;
    // 
    rolePermissions []UnifiedRolePermissionable;
}
// NewDefaultUserRoleOverride instantiates a new defaultUserRoleOverride and sets the default values.
func NewDefaultUserRoleOverride()(*DefaultUserRoleOverride) {
    m := &DefaultUserRoleOverride{
        Entity: *NewEntity(),
    }
    return m
}
// CreateDefaultUserRoleOverrideFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDefaultUserRoleOverrideFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewDefaultUserRoleOverride(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DefaultUserRoleOverride) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["isDefault"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsDefault(val)
        }
        return nil
    }
    res["rolePermissions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUnifiedRolePermissionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UnifiedRolePermissionable, len(val))
            for i, v := range val {
                res[i] = v.(UnifiedRolePermissionable)
            }
            m.SetRolePermissions(res)
        }
        return nil
    }
    return res
}
// GetIsDefault gets the isDefault property value. 
func (m *DefaultUserRoleOverride) GetIsDefault()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isDefault
    }
}
// GetRolePermissions gets the rolePermissions property value. 
func (m *DefaultUserRoleOverride) GetRolePermissions()([]UnifiedRolePermissionable) {
    if m == nil {
        return nil
    } else {
        return m.rolePermissions
    }
}
func (m *DefaultUserRoleOverride) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *DefaultUserRoleOverride) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("isDefault", m.GetIsDefault())
        if err != nil {
            return err
        }
    }
    if m.GetRolePermissions() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetRolePermissions()))
        for i, v := range m.GetRolePermissions() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("rolePermissions", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetIsDefault sets the isDefault property value. 
func (m *DefaultUserRoleOverride) SetIsDefault(value *bool)() {
    if m != nil {
        m.isDefault = value
    }
}
// SetRolePermissions sets the rolePermissions property value. 
func (m *DefaultUserRoleOverride) SetRolePermissions(value []UnifiedRolePermissionable)() {
    if m != nil {
        m.rolePermissions = value
    }
}
