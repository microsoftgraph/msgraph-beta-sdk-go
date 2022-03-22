package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// CloudPcWindowsSettings 
type CloudPcWindowsSettings struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The Windows language/region tag to use for language pack configuration and localization of the Cloud PC. The default value is en-US, which corresponds to English (United States).
    language *string;
}
// NewCloudPcWindowsSettings instantiates a new cloudPcWindowsSettings and sets the default values.
func NewCloudPcWindowsSettings()(*CloudPcWindowsSettings) {
    m := &CloudPcWindowsSettings{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateCloudPcWindowsSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCloudPcWindowsSettingsFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewCloudPcWindowsSettings(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CloudPcWindowsSettings) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CloudPcWindowsSettings) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["language"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLanguage(val)
        }
        return nil
    }
    return res
}
// GetLanguage gets the language property value. The Windows language/region tag to use for language pack configuration and localization of the Cloud PC. The default value is en-US, which corresponds to English (United States).
func (m *CloudPcWindowsSettings) GetLanguage()(*string) {
    if m == nil {
        return nil
    } else {
        return m.language
    }
}
// Serialize serializes information the current object
func (m *CloudPcWindowsSettings) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("language", m.GetLanguage())
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
func (m *CloudPcWindowsSettings) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetLanguage sets the language property value. The Windows language/region tag to use for language pack configuration and localization of the Cloud PC. The default value is en-US, which corresponds to English (United States).
func (m *CloudPcWindowsSettings) SetLanguage(value *string)() {
    if m != nil {
        m.language = value
    }
}
