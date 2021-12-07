package uploadpkcs12

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// UploadPkcs12RequestBody 
type UploadPkcs12RequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    key *string;
    // 
    password *string;
}
// NewUploadPkcs12RequestBody instantiates a new uploadPkcs12RequestBody and sets the default values.
func NewUploadPkcs12RequestBody()(*UploadPkcs12RequestBody) {
    m := &UploadPkcs12RequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UploadPkcs12RequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetKey gets the key property value. 
func (m *UploadPkcs12RequestBody) GetKey()(*string) {
    if m == nil {
        return nil
    } else {
        return m.key
    }
}
// GetPassword gets the password property value. 
func (m *UploadPkcs12RequestBody) GetPassword()(*string) {
    if m == nil {
        return nil
    } else {
        return m.password
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UploadPkcs12RequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["key"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKey(val)
        }
        return nil
    }
    res["password"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPassword(val)
        }
        return nil
    }
    return res
}
func (m *UploadPkcs12RequestBody) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *UploadPkcs12RequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("key", m.GetKey())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("password", m.GetPassword())
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
func (m *UploadPkcs12RequestBody) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetKey sets the key property value. 
func (m *UploadPkcs12RequestBody) SetKey(value *string)() {
    if m != nil {
        m.key = value
    }
}
// SetPassword sets the password property value. 
func (m *UploadPkcs12RequestBody) SetPassword(value *string)() {
    if m != nil {
        m.password = value
    }
}
