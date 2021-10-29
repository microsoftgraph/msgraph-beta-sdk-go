package uploaddeptoken

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type UploadDepTokenRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    appleId *string;
    // 
    depToken *string;
}
// Instantiates a new uploadDepTokenRequestBody and sets the default values.
func NewUploadDepTokenRequestBody()(*UploadDepTokenRequestBody) {
    m := &UploadDepTokenRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UploadDepTokenRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the appleId property value. 
func (m *UploadDepTokenRequestBody) GetAppleId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.appleId
    }
}
// Gets the depToken property value. 
func (m *UploadDepTokenRequestBody) GetDepToken()(*string) {
    if m == nil {
        return nil
    } else {
        return m.depToken
    }
}
// The deserialization information for the current model
func (m *UploadDepTokenRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["appleId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAppleId(val)
        return nil
    }
    res["depToken"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDepToken(val)
        return nil
    }
    return res
}
func (m *UploadDepTokenRequestBody) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *UploadDepTokenRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("appleId", m.GetAppleId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("depToken", m.GetDepToken())
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
func (m *UploadDepTokenRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the appleId property value. 
// Parameters:
//  - value : Value to set for the appleId property.
func (m *UploadDepTokenRequestBody) SetAppleId(value *string)() {
    m.appleId = value
}
// Sets the depToken property value. 
// Parameters:
//  - value : Value to set for the depToken property.
func (m *UploadDepTokenRequestBody) SetDepToken(value *string)() {
    m.depToken = value
}
