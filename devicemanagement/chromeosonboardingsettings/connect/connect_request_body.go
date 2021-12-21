package connect

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ConnectRequestBody 
type ConnectRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    ownerAccessToken *string;
    // 
    ownerUserPrincipalName *string;
}
// NewConnectRequestBody instantiates a new connectRequestBody and sets the default values.
func NewConnectRequestBody()(*ConnectRequestBody) {
    m := &ConnectRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ConnectRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetOwnerAccessToken gets the ownerAccessToken property value. 
func (m *ConnectRequestBody) GetOwnerAccessToken()(*string) {
    if m == nil {
        return nil
    } else {
        return m.ownerAccessToken
    }
}
// GetOwnerUserPrincipalName gets the ownerUserPrincipalName property value. 
func (m *ConnectRequestBody) GetOwnerUserPrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.ownerUserPrincipalName
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ConnectRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["ownerAccessToken"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOwnerAccessToken(val)
        }
        return nil
    }
    res["ownerUserPrincipalName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOwnerUserPrincipalName(val)
        }
        return nil
    }
    return res
}
func (m *ConnectRequestBody) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ConnectRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("ownerAccessToken", m.GetOwnerAccessToken())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("ownerUserPrincipalName", m.GetOwnerUserPrincipalName())
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
func (m *ConnectRequestBody) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetOwnerAccessToken sets the ownerAccessToken property value. 
func (m *ConnectRequestBody) SetOwnerAccessToken(value *string)() {
    if m != nil {
        m.ownerAccessToken = value
    }
}
// SetOwnerUserPrincipalName sets the ownerUserPrincipalName property value. 
func (m *ConnectRequestBody) SetOwnerUserPrincipalName(value *string)() {
    if m != nil {
        m.ownerUserPrincipalName = value
    }
}
