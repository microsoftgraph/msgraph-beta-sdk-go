package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type AccessPackageLocalizedContent struct {
    additionalData map[string]interface{};
    defaultText *string;
    localizedTexts []AccessPackageLocalizedText;
}
func NewAccessPackageLocalizedContent()(*AccessPackageLocalizedContent) {
    m := &AccessPackageLocalizedContent{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *AccessPackageLocalizedContent) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *AccessPackageLocalizedContent) GetDefaultText()(*string) {
    if m == nil {
        return nil
    } else {
        return m.defaultText
    }
}
func (m *AccessPackageLocalizedContent) GetLocalizedTexts()([]AccessPackageLocalizedText) {
    if m == nil {
        return nil
    } else {
        return m.localizedTexts
    }
}
func (m *AccessPackageLocalizedContent) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["defaultText"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDefaultText(val)
        return nil
    }
    res["localizedTexts"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAccessPackageLocalizedText() })
        if err != nil {
            return err
        }
        res := make([]AccessPackageLocalizedText, len(val))
        for i, v := range val {
            res[i] = *(v.(*AccessPackageLocalizedText))
        }
        m.SetLocalizedTexts(res)
        return nil
    }
    return res
}
func (m *AccessPackageLocalizedContent) IsNil()(bool) {
    return m == nil
}
func (m *AccessPackageLocalizedContent) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("defaultText", m.GetDefaultText())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetLocalizedTexts()))
        for i, v := range m.GetLocalizedTexts() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("localizedTexts", cast)
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
func (m *AccessPackageLocalizedContent) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *AccessPackageLocalizedContent) SetDefaultText(value *string)() {
    m.defaultText = value
}
func (m *AccessPackageLocalizedContent) SetLocalizedTexts(value []AccessPackageLocalizedText)() {
    m.localizedTexts = value
}
