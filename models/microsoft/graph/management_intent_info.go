package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type ManagementIntentInfo struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The display name for the management intent. Optional. Read-only.
    managementIntentDisplayName *string;
    // The identifier for the management intent. Required. Read-only.
    managementIntentId *string;
    // The collection of management template information associated with the management intent. Optional. Read-only.
    managementTemplates []ManagementTemplateDetailedInfo;
}
// Instantiates a new managementIntentInfo and sets the default values.
func NewManagementIntentInfo()(*ManagementIntentInfo) {
    m := &ManagementIntentInfo{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ManagementIntentInfo) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the managementIntentDisplayName property value. The display name for the management intent. Optional. Read-only.
func (m *ManagementIntentInfo) GetManagementIntentDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.managementIntentDisplayName
    }
}
// Gets the managementIntentId property value. The identifier for the management intent. Required. Read-only.
func (m *ManagementIntentInfo) GetManagementIntentId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.managementIntentId
    }
}
// Gets the managementTemplates property value. The collection of management template information associated with the management intent. Optional. Read-only.
func (m *ManagementIntentInfo) GetManagementTemplates()([]ManagementTemplateDetailedInfo) {
    if m == nil {
        return nil
    } else {
        return m.managementTemplates
    }
}
// The deserialization information for the current model
func (m *ManagementIntentInfo) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["managementIntentDisplayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagementIntentDisplayName(val)
        }
        return nil
    }
    res["managementIntentId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagementIntentId(val)
        }
        return nil
    }
    res["managementTemplates"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewManagementTemplateDetailedInfo() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagementTemplateDetailedInfo, len(val))
            for i, v := range val {
                res[i] = *(v.(*ManagementTemplateDetailedInfo))
            }
            m.SetManagementTemplates(res)
        }
        return nil
    }
    return res
}
func (m *ManagementIntentInfo) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *ManagementIntentInfo) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("managementIntentDisplayName", m.GetManagementIntentDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("managementIntentId", m.GetManagementIntentId())
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
        err := writer.WriteCollectionOfObjectValues("managementTemplates", cast)
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *ManagementIntentInfo) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the managementIntentDisplayName property value. The display name for the management intent. Optional. Read-only.
// Parameters:
//  - value : Value to set for the managementIntentDisplayName property.
func (m *ManagementIntentInfo) SetManagementIntentDisplayName(value *string)() {
    m.managementIntentDisplayName = value
}
// Sets the managementIntentId property value. The identifier for the management intent. Required. Read-only.
// Parameters:
//  - value : Value to set for the managementIntentId property.
func (m *ManagementIntentInfo) SetManagementIntentId(value *string)() {
    m.managementIntentId = value
}
// Sets the managementTemplates property value. The collection of management template information associated with the management intent. Optional. Read-only.
// Parameters:
//  - value : Value to set for the managementTemplates property.
func (m *ManagementIntentInfo) SetManagementTemplates(value []ManagementTemplateDetailedInfo)() {
    m.managementTemplates = value
}
