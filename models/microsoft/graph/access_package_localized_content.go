package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AccessPackageLocalizedContent 
type AccessPackageLocalizedContent struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The fallback string, which is used when a requested localization is not available. Required.
    defaultText *string;
    // Content represented in a format for a specific locale.
    localizedTexts []AccessPackageLocalizedText;
}
// NewAccessPackageLocalizedContent instantiates a new accessPackageLocalizedContent and sets the default values.
func NewAccessPackageLocalizedContent()(*AccessPackageLocalizedContent) {
    m := &AccessPackageLocalizedContent{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AccessPackageLocalizedContent) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDefaultText gets the defaultText property value. The fallback string, which is used when a requested localization is not available. Required.
func (m *AccessPackageLocalizedContent) GetDefaultText()(*string) {
    if m == nil {
        return nil
    } else {
        return m.defaultText
    }
}
// GetLocalizedTexts gets the localizedTexts property value. Content represented in a format for a specific locale.
func (m *AccessPackageLocalizedContent) GetLocalizedTexts()([]AccessPackageLocalizedText) {
    if m == nil {
        return nil
    } else {
        return m.localizedTexts
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AccessPackageLocalizedContent) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["defaultText"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultText(val)
        }
        return nil
    }
    res["localizedTexts"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAccessPackageLocalizedText() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AccessPackageLocalizedText, len(val))
            for i, v := range val {
                res[i] = *(v.(*AccessPackageLocalizedText))
            }
            m.SetLocalizedTexts(res)
        }
        return nil
    }
    return res
}
func (m *AccessPackageLocalizedContent) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *AccessPackageLocalizedContent) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("defaultText", m.GetDefaultText())
        if err != nil {
            return err
        }
    }
    if m.GetLocalizedTexts() != nil {
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AccessPackageLocalizedContent) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetDefaultText sets the defaultText property value. The fallback string, which is used when a requested localization is not available. Required.
func (m *AccessPackageLocalizedContent) SetDefaultText(value *string)() {
    if m != nil {
        m.defaultText = value
    }
}
// SetLocalizedTexts sets the localizedTexts property value. Content represented in a format for a specific locale.
func (m *AccessPackageLocalizedContent) SetLocalizedTexts(value []AccessPackageLocalizedText)() {
    if m != nil {
        m.localizedTexts = value
    }
}
