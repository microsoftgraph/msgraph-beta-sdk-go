package connect

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type ConnectRequestBody struct {
    additionalData map[string]interface{};
    ownerUserPrincipalName *string;
    serviceAccountCredentials *string;
}
func NewConnectRequestBody()(*ConnectRequestBody) {
    m := &ConnectRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *ConnectRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *ConnectRequestBody) GetOwnerUserPrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.ownerUserPrincipalName
    }
}
func (m *ConnectRequestBody) GetServiceAccountCredentials()(*string) {
    if m == nil {
        return nil
    } else {
        return m.serviceAccountCredentials
    }
}
func (m *ConnectRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["ownerUserPrincipalName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetOwnerUserPrincipalName(val)
        return nil
    }
    res["serviceAccountCredentials"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetServiceAccountCredentials(val)
        return nil
    }
    return res
}
func (m *ConnectRequestBody) IsNil()(bool) {
    return m == nil
}
func (m *ConnectRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("ownerUserPrincipalName", m.GetOwnerUserPrincipalName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("serviceAccountCredentials", m.GetServiceAccountCredentials())
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
func (m *ConnectRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *ConnectRequestBody) SetOwnerUserPrincipalName(value *string)() {
    m.ownerUserPrincipalName = value
}
func (m *ConnectRequestBody) SetServiceAccountCredentials(value *string)() {
    m.serviceAccountCredentials = value
}
