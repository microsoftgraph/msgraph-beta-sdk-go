package generatekey

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type GenerateKeyRequestBody struct {
    additionalData map[string]interface{};
    exp *int64;
    kty *string;
    nbf *int64;
    use *string;
}
func NewGenerateKeyRequestBody()(*GenerateKeyRequestBody) {
    m := &GenerateKeyRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *GenerateKeyRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *GenerateKeyRequestBody) GetExp()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.exp
    }
}
func (m *GenerateKeyRequestBody) GetKty()(*string) {
    if m == nil {
        return nil
    } else {
        return m.kty
    }
}
func (m *GenerateKeyRequestBody) GetNbf()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.nbf
    }
}
func (m *GenerateKeyRequestBody) GetUse()(*string) {
    if m == nil {
        return nil
    } else {
        return m.use
    }
}
func (m *GenerateKeyRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["exp"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetExp(val)
        return nil
    }
    res["kty"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetKty(val)
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
func (m *GenerateKeyRequestBody) IsNil()(bool) {
    return m == nil
}
func (m *GenerateKeyRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteInt64Value("exp", m.GetExp())
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
func (m *GenerateKeyRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *GenerateKeyRequestBody) SetExp(value *int64)() {
    m.exp = value
}
func (m *GenerateKeyRequestBody) SetKty(value *string)() {
    m.kty = value
}
func (m *GenerateKeyRequestBody) SetNbf(value *int64)() {
    m.nbf = value
}
func (m *GenerateKeyRequestBody) SetUse(value *string)() {
    m.use = value
}
