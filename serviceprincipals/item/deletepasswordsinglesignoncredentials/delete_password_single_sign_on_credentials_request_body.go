package deletepasswordsinglesignoncredentials

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DeletePasswordSingleSignOnCredentialsRequestBody 
type DeletePasswordSingleSignOnCredentialsRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    id *string;
}
// NewDeletePasswordSingleSignOnCredentialsRequestBody instantiates a new deletePasswordSingleSignOnCredentialsRequestBody and sets the default values.
func NewDeletePasswordSingleSignOnCredentialsRequestBody()(*DeletePasswordSingleSignOnCredentialsRequestBody) {
    m := &DeletePasswordSingleSignOnCredentialsRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeletePasswordSingleSignOnCredentialsRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetId gets the id property value. 
func (m *DeletePasswordSingleSignOnCredentialsRequestBody) GetId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.id
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeletePasswordSingleSignOnCredentialsRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["id"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
        }
        return nil
    }
    return res
}
func (m *DeletePasswordSingleSignOnCredentialsRequestBody) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *DeletePasswordSingleSignOnCredentialsRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("id", m.GetId())
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
func (m *DeletePasswordSingleSignOnCredentialsRequestBody) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetId sets the id property value. 
func (m *DeletePasswordSingleSignOnCredentialsRequestBody) SetId(value *string)() {
    if m != nil {
        m.id = value
    }
}
