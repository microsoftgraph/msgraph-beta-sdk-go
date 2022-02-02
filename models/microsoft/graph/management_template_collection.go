package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ManagementTemplateCollection 
type ManagementTemplateCollection struct {
    Entity
    // 
    description *string;
    // 
    displayName *string;
    // 
    managementTemplates []ManagementTemplate;
}
// NewManagementTemplateCollection instantiates a new managementTemplateCollection and sets the default values.
func NewManagementTemplateCollection()(*ManagementTemplateCollection) {
    m := &ManagementTemplateCollection{
        Entity: *NewEntity(),
    }
    return m
}
// GetDescription gets the description property value. 
func (m *ManagementTemplateCollection) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetDisplayName gets the displayName property value. 
func (m *ManagementTemplateCollection) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetManagementTemplates gets the managementTemplates property value. 
func (m *ManagementTemplateCollection) GetManagementTemplates()([]ManagementTemplate) {
    if m == nil {
        return nil
    } else {
        return m.managementTemplates
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ManagementTemplateCollection) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
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
    res["managementTemplates"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewManagementTemplate() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagementTemplate, len(val))
            for i, v := range val {
                res[i] = *(v.(*ManagementTemplate))
            }
            m.SetManagementTemplates(res)
        }
        return nil
    }
    return res
}
func (m *ManagementTemplateCollection) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ManagementTemplateCollection) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
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
    if m.GetManagementTemplates() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetManagementTemplates()))
        for i, v := range m.GetManagementTemplates() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("managementTemplates", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDescription sets the description property value. 
func (m *ManagementTemplateCollection) SetDescription(value *string)() {
    if m != nil {
        m.description = value
    }
}
// SetDisplayName sets the displayName property value. 
func (m *ManagementTemplateCollection) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetManagementTemplates sets the managementTemplates property value. 
func (m *ManagementTemplateCollection) SetManagementTemplates(value []ManagementTemplate)() {
    if m != nil {
        m.managementTemplates = value
    }
}
