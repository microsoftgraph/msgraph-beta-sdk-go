package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DirectorySettingTemplate 
type DirectorySettingTemplate struct {
    DirectoryObject
    // Description of the template. Read-only.
    description *string;
    // Display name of the template. Read-only.
    displayName *string;
    // Collection of settingTemplateValues that list the set of available settings, defaults and types that make up this template.  Read-only.
    values []SettingTemplateValueable;
}
// NewDirectorySettingTemplate instantiates a new directorySettingTemplate and sets the default values.
func NewDirectorySettingTemplate()(*DirectorySettingTemplate) {
    m := &DirectorySettingTemplate{
        DirectoryObject: *NewDirectoryObject(),
    }
    return m
}
// CreateDirectorySettingTemplateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDirectorySettingTemplateFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewDirectorySettingTemplate(), nil
}
// GetDescription gets the description property value. Description of the template. Read-only.
func (m *DirectorySettingTemplate) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetDisplayName gets the displayName property value. Display name of the template. Read-only.
func (m *DirectorySettingTemplate) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DirectorySettingTemplate) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.DirectoryObject.GetFieldDeserializers()
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
    res["values"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSettingTemplateValueFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SettingTemplateValueable, len(val))
            for i, v := range val {
                res[i] = v.(SettingTemplateValueable)
            }
            m.SetValues(res)
        }
        return nil
    }
    return res
}
// GetValues gets the values property value. Collection of settingTemplateValues that list the set of available settings, defaults and types that make up this template.  Read-only.
func (m *DirectorySettingTemplate) GetValues()([]SettingTemplateValueable) {
    if m == nil {
        return nil
    } else {
        return m.values
    }
}
// Serialize serializes information the current object
func (m *DirectorySettingTemplate) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.DirectoryObject.Serialize(writer)
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
    if m.GetValues() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetValues()))
        for i, v := range m.GetValues() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("values", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDescription sets the description property value. Description of the template. Read-only.
func (m *DirectorySettingTemplate) SetDescription(value *string)() {
    if m != nil {
        m.description = value
    }
}
// SetDisplayName sets the displayName property value. Display name of the template. Read-only.
func (m *DirectorySettingTemplate) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetValues sets the values property value. Collection of settingTemplateValues that list the set of available settings, defaults and types that make up this template.  Read-only.
func (m *DirectorySettingTemplate) SetValues(value []SettingTemplateValueable)() {
    if m != nil {
        m.values = value
    }
}
