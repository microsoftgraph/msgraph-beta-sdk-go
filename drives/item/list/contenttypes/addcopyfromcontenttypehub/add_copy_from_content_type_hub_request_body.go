package addcopyfromcontenttypehub

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AddCopyFromContentTypeHubRequestBody 
type AddCopyFromContentTypeHubRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    contentTypeId *string;
}
// NewAddCopyFromContentTypeHubRequestBody instantiates a new addCopyFromContentTypeHubRequestBody and sets the default values.
func NewAddCopyFromContentTypeHubRequestBody()(*AddCopyFromContentTypeHubRequestBody) {
    m := &AddCopyFromContentTypeHubRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AddCopyFromContentTypeHubRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetContentTypeId gets the contentTypeId property value. 
func (m *AddCopyFromContentTypeHubRequestBody) GetContentTypeId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.contentTypeId
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AddCopyFromContentTypeHubRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["contentTypeId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContentTypeId(val)
        }
        return nil
    }
    return res
}
func (m *AddCopyFromContentTypeHubRequestBody) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *AddCopyFromContentTypeHubRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("contentTypeId", m.GetContentTypeId())
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
func (m *AddCopyFromContentTypeHubRequestBody) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetContentTypeId sets the contentTypeId property value. 
func (m *AddCopyFromContentTypeHubRequestBody) SetContentTypeId(value *string)() {
    if m != nil {
        m.contentTypeId = value
    }
}
