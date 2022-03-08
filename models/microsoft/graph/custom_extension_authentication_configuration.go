package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// CustomExtensionAuthenticationConfiguration provides operations to manage the identityGovernance singleton.
type CustomExtensionAuthenticationConfiguration struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
}
// NewCustomExtensionAuthenticationConfiguration instantiates a new customExtensionAuthenticationConfiguration and sets the default values.
func NewCustomExtensionAuthenticationConfiguration()(*CustomExtensionAuthenticationConfiguration) {
    m := &CustomExtensionAuthenticationConfiguration{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateCustomExtensionAuthenticationConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCustomExtensionAuthenticationConfigurationFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewCustomExtensionAuthenticationConfiguration(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CustomExtensionAuthenticationConfiguration) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CustomExtensionAuthenticationConfiguration) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    return res
}
func (m *CustomExtensionAuthenticationConfiguration) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *CustomExtensionAuthenticationConfiguration) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CustomExtensionAuthenticationConfiguration) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
