package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// TemplateAction 
type TemplateAction struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    description *string;
    // 
    displayName *string;
    // 
    licenses *LicenseDetails;
    // 
    service *string;
    // 
    settings []Setting;
    // 
    templateActionId *string;
}
// NewTemplateAction instantiates a new templateAction and sets the default values.
func NewTemplateAction()(*TemplateAction) {
    m := &TemplateAction{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TemplateAction) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDescription gets the description property value. 
func (m *TemplateAction) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetDisplayName gets the displayName property value. 
func (m *TemplateAction) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetLicenses gets the licenses property value. 
func (m *TemplateAction) GetLicenses()(*LicenseDetails) {
    if m == nil {
        return nil
    } else {
        return m.licenses
    }
}
// GetService gets the service property value. 
func (m *TemplateAction) GetService()(*string) {
    if m == nil {
        return nil
    } else {
        return m.service
    }
}
// GetSettings gets the settings property value. 
func (m *TemplateAction) GetSettings()([]Setting) {
    if m == nil {
        return nil
    } else {
        return m.settings
    }
}
// GetTemplateActionId gets the templateActionId property value. 
func (m *TemplateAction) GetTemplateActionId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.templateActionId
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TemplateAction) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
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
    res["licenses"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewLicenseDetails() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLicenses(val.(*LicenseDetails))
        }
        return nil
    }
    res["service"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetService(val)
        }
        return nil
    }
    res["settings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSetting() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Setting, len(val))
            for i, v := range val {
                res[i] = *(v.(*Setting))
            }
            m.SetSettings(res)
        }
        return nil
    }
    res["templateActionId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTemplateActionId(val)
        }
        return nil
    }
    return res
}
func (m *TemplateAction) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *TemplateAction) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("licenses", m.GetLicenses())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("service", m.GetService())
        if err != nil {
            return err
        }
    }
    if m.GetSettings() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetSettings()))
        for i, v := range m.GetSettings() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("settings", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("templateActionId", m.GetTemplateActionId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TemplateAction) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetDescription sets the description property value. 
func (m *TemplateAction) SetDescription(value *string)() {
    if m != nil {
        m.description = value
    }
}
// SetDisplayName sets the displayName property value. 
func (m *TemplateAction) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetLicenses sets the licenses property value. 
func (m *TemplateAction) SetLicenses(value *LicenseDetails)() {
    if m != nil {
        m.licenses = value
    }
}
// SetService sets the service property value. 
func (m *TemplateAction) SetService(value *string)() {
    if m != nil {
        m.service = value
    }
}
// SetSettings sets the settings property value. 
func (m *TemplateAction) SetSettings(value []Setting)() {
    if m != nil {
        m.settings = value
    }
}
// SetTemplateActionId sets the templateActionId property value. 
func (m *TemplateAction) SetTemplateActionId(value *string)() {
    if m != nil {
        m.templateActionId = value
    }
}
