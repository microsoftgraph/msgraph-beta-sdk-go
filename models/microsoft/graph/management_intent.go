package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ManagementIntent provides operations to manage the tenantRelationship singleton.
type ManagementIntent struct {
    Entity
    // The display name for the management intent. Optional. Read-only.
    displayName *string;
    // A flag indicating whether the management intent is global. Required. Read-only.
    isGlobal *bool;
    // The collection of management templates associated with the management intent. Optional. Read-only.
    managementTemplates []ManagementTemplateDetailedInfoable;
}
// NewManagementIntent instantiates a new managementIntent and sets the default values.
func NewManagementIntent()(*ManagementIntent) {
    m := &ManagementIntent{
        Entity: *NewEntity(),
    }
    return m
}
// CreateManagementIntentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateManagementIntentFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewManagementIntent(), nil
}
// GetDisplayName gets the displayName property value. The display name for the management intent. Optional. Read-only.
func (m *ManagementIntent) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ManagementIntent) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
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
    res["isGlobal"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsGlobal(val)
        }
        return nil
    }
    res["managementTemplates"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateManagementTemplateDetailedInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagementTemplateDetailedInfoable, len(val))
            for i, v := range val {
                res[i] = v.(ManagementTemplateDetailedInfoable)
            }
            m.SetManagementTemplates(res)
        }
        return nil
    }
    return res
}
// GetIsGlobal gets the isGlobal property value. A flag indicating whether the management intent is global. Required. Read-only.
func (m *ManagementIntent) GetIsGlobal()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isGlobal
    }
}
// GetManagementTemplates gets the managementTemplates property value. The collection of management templates associated with the management intent. Optional. Read-only.
func (m *ManagementIntent) GetManagementTemplates()([]ManagementTemplateDetailedInfoable) {
    if m == nil {
        return nil
    } else {
        return m.managementTemplates
    }
}
func (m *ManagementIntent) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ManagementIntent) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
        err = writer.WriteBoolValue("isGlobal", m.GetIsGlobal())
        if err != nil {
            return err
        }
    }
    if m.GetManagementTemplates() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetManagementTemplates()))
        for i, v := range m.GetManagementTemplates() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("managementTemplates", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDisplayName sets the displayName property value. The display name for the management intent. Optional. Read-only.
func (m *ManagementIntent) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetIsGlobal sets the isGlobal property value. A flag indicating whether the management intent is global. Required. Read-only.
func (m *ManagementIntent) SetIsGlobal(value *bool)() {
    if m != nil {
        m.isGlobal = value
    }
}
// SetManagementTemplates sets the managementTemplates property value. The collection of management templates associated with the management intent. Optional. Read-only.
func (m *ManagementIntent) SetManagementTemplates(value []ManagementTemplateDetailedInfoable)() {
    if m != nil {
        m.managementTemplates = value
    }
}
