package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type ManagementTemplateCollection struct {
    Entity
    // 
    description *string;
    // 
    displayName *string;
    // 
    managementTemplates []ManagementTemplate;
}
// Instantiates a new managementTemplateCollection and sets the default values.
func NewManagementTemplateCollection()(*ManagementTemplateCollection) {
    m := &ManagementTemplateCollection{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the description property value. 
func (m *ManagementTemplateCollection) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// Gets the displayName property value. 
func (m *ManagementTemplateCollection) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the managementTemplates property value. 
func (m *ManagementTemplateCollection) GetManagementTemplates()([]ManagementTemplate) {
    if m == nil {
        return nil
    } else {
        return m.managementTemplates
    }
}
// The deserialization information for the current model
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
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
    {
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
// Sets the description property value. 
// Parameters:
//  - value : Value to set for the description property.
func (m *ManagementTemplateCollection) SetDescription(value *string)() {
    m.description = value
}
// Sets the displayName property value. 
// Parameters:
//  - value : Value to set for the displayName property.
func (m *ManagementTemplateCollection) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the managementTemplates property value. 
// Parameters:
//  - value : Value to set for the managementTemplates property.
func (m *ManagementTemplateCollection) SetManagementTemplates(value []ManagementTemplate)() {
    m.managementTemplates = value
}
