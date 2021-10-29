package uploadsecret

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type UploadSecretRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    exp *int64;
    // 
    k *string;
    // 
    nbf *int64;
    // 
    use *string;
}
// Instantiates a new uploadSecretRequestBody and sets the default values.
func NewUploadSecretRequestBody()(*UploadSecretRequestBody) {
    m := &UploadSecretRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UploadSecretRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the exp property value. 
func (m *UploadSecretRequestBody) GetExp()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.exp
    }
}
// Gets the k property value. 
func (m *UploadSecretRequestBody) GetK()(*string) {
    if m == nil {
        return nil
    } else {
        return m.k
    }
}
// Gets the nbf property value. 
func (m *UploadSecretRequestBody) GetNbf()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.nbf
    }
}
// Gets the use property value. 
func (m *UploadSecretRequestBody) GetUse()(*string) {
    if m == nil {
        return nil
    } else {
        return m.use
    }
}
// The deserialization information for the current model
func (m *UploadSecretRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["exp"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetExp(val)
        return nil
    }
    res["k"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetK(val)
        return nil
    }
    res["nbf"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetNbf(val)
        return nil
    }
    res["use"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetUse(val)
        return nil
    }
    return res
}
func (m *UploadSecretRequestBody) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *UploadSecretRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteInt64Value("exp", m.GetExp())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("k", m.GetK())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("nbf", m.GetNbf())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("use", m.GetUse())
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
func (m *UploadSecretRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the exp property value. 
// Parameters:
//  - value : Value to set for the exp property.
func (m *UploadSecretRequestBody) SetExp(value *int64)() {
    m.exp = value
}
// Sets the k property value. 
// Parameters:
//  - value : Value to set for the k property.
func (m *UploadSecretRequestBody) SetK(value *string)() {
    m.k = value
}
// Sets the nbf property value. 
// Parameters:
//  - value : Value to set for the nbf property.
func (m *UploadSecretRequestBody) SetNbf(value *int64)() {
    m.nbf = value
}
// Sets the use property value. 
// Parameters:
//  - value : Value to set for the use property.
func (m *UploadSecretRequestBody) SetUse(value *string)() {
    m.use = value
}
