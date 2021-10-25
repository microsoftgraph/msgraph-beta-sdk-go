package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type ManagementIntentInfo struct {
    additionalData map[string]interface{};
    managementIntentDisplayName *string;
    managementIntentId *string;
    managementTemplates []ManagementTemplateDetailedInfo;
}
func NewManagementIntentInfo()(*ManagementIntentInfo) {
    m := &ManagementIntentInfo{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *ManagementIntentInfo) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *ManagementIntentInfo) GetManagementIntentDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.managementIntentDisplayName
    }
}
func (m *ManagementIntentInfo) GetManagementIntentId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.managementIntentId
    }
}
func (m *ManagementIntentInfo) GetManagementTemplates()([]ManagementTemplateDetailedInfo) {
    if m == nil {
        return nil
    } else {
        return m.managementTemplates
    }
}
func (m *ManagementIntentInfo) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["managementIntentDisplayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetManagementIntentDisplayName(val)
        return nil
    }
    res["managementIntentId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetManagementIntentId(val)
        return nil
    }
    res["managementTemplates"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewManagementTemplateDetailedInfo() })
        if err != nil {
            return err
        }
        res := make([]ManagementTemplateDetailedInfo, len(val))
        for i, v := range val {
            res[i] = *(v.(*ManagementTemplateDetailedInfo))
        }
        m.SetManagementTemplates(res)
        return nil
    }
    return res
}
func (m *ManagementIntentInfo) IsNil()(bool) {
    return m == nil
}
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
func (m *ManagementIntentInfo) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *ManagementIntentInfo) SetManagementIntentDisplayName(value *string)() {
    m.managementIntentDisplayName = value
}
func (m *ManagementIntentInfo) SetManagementIntentId(value *string)() {
    m.managementIntentId = value
}
func (m *ManagementIntentInfo) SetManagementTemplates(value []ManagementTemplateDetailedInfo)() {
    m.managementTemplates = value
}
