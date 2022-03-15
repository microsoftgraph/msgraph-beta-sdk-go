package generatekey

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// GenerateKeyRequestBody provides operations to call the generateKey method.
type GenerateKeyRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    exp *int32;
    // 
    kty *string;
    // 
    nbf *int32;
    // 
    use *string;
}
// NewGenerateKeyRequestBody instantiates a new generateKeyRequestBody and sets the default values.
func NewGenerateKeyRequestBody()(*GenerateKeyRequestBody) {
    m := &GenerateKeyRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateGenerateKeyRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGenerateKeyRequestBodyFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewGenerateKeyRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *GenerateKeyRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetExp gets the exp property value. 
func (m *GenerateKeyRequestBody) GetExp()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.exp
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GenerateKeyRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["exp"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExp(val)
        }
        return nil
    }
    res["kty"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKty(val)
        }
        return nil
    }
    res["nbf"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNbf(val)
        }
        return nil
    }
    res["use"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUse(val)
        }
        return nil
    }
    return res
}
// GetKty gets the kty property value. 
func (m *GenerateKeyRequestBody) GetKty()(*string) {
    if m == nil {
        return nil
    } else {
        return m.kty
    }
}
// GetNbf gets the nbf property value. 
func (m *GenerateKeyRequestBody) GetNbf()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.nbf
    }
}
// GetUse gets the use property value. 
func (m *GenerateKeyRequestBody) GetUse()(*string) {
    if m == nil {
        return nil
    } else {
        return m.use
    }
}
func (m *GenerateKeyRequestBody) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *GenerateKeyRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("exp", m.GetExp())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("kty", m.GetKty())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("nbf", m.GetNbf())
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *GenerateKeyRequestBody) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetExp sets the exp property value. 
func (m *GenerateKeyRequestBody) SetExp(value *int32)() {
    if m != nil {
        m.exp = value
    }
}
// SetKty sets the kty property value. 
func (m *GenerateKeyRequestBody) SetKty(value *string)() {
    if m != nil {
        m.kty = value
    }
}
// SetNbf sets the nbf property value. 
func (m *GenerateKeyRequestBody) SetNbf(value *int32)() {
    if m != nil {
        m.nbf = value
    }
}
// SetUse sets the use property value. 
func (m *GenerateKeyRequestBody) SetUse(value *string)() {
    if m != nil {
        m.use = value
    }
}
