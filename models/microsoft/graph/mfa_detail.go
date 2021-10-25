package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type MfaDetail struct {
    additionalData map[string]interface{};
    authDetail *string;
    authMethod *string;
}
func NewMfaDetail()(*MfaDetail) {
    m := &MfaDetail{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *MfaDetail) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *MfaDetail) GetAuthDetail()(*string) {
    if m == nil {
        return nil
    } else {
        return m.authDetail
    }
}
func (m *MfaDetail) GetAuthMethod()(*string) {
    if m == nil {
        return nil
    } else {
        return m.authMethod
    }
}
func (m *MfaDetail) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["authDetail"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAuthDetail(val)
        return nil
    }
    res["authMethod"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAuthMethod(val)
        return nil
    }
    return res
}
func (m *MfaDetail) IsNil()(bool) {
    return m == nil
}
func (m *MfaDetail) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("authDetail", m.GetAuthDetail())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("authMethod", m.GetAuthMethod())
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
func (m *MfaDetail) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *MfaDetail) SetAuthDetail(value *string)() {
    m.authDetail = value
}
func (m *MfaDetail) SetAuthMethod(value *string)() {
    m.authMethod = value
}
