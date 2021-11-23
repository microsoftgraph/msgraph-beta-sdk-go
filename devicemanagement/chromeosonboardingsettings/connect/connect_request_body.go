package connect

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ConnectRequestBody 
type ConnectRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    ownerUserPrincipalName *string;
    // 
    serviceAccountCredentials *string;
}
// NewConnectRequestBody instantiates a new connectRequestBody and sets the default values.
func NewConnectRequestBody()(*ConnectRequestBody) {
    m := &ConnectRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ConnectRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
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
// GetServiceAccountCredentials gets the serviceAccountCredentials property value. 
func (m *ConnectRequestBody) GetServiceAccountCredentials()(*string) {
    if m == nil {
        return nil
    } else {
        return m.serviceAccountCredentials
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ConnectRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
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
    res["serviceAccountCredentials"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetServiceAccountCredentials(val)
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
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ConnectRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetOwnerUserPrincipalName sets the ownerUserPrincipalName property value. 
func (m *ConnectRequestBody) SetOwnerUserPrincipalName(value *string)() {
    m.ownerUserPrincipalName = value
}
// SetServiceAccountCredentials sets the serviceAccountCredentials property value. 
func (m *ConnectRequestBody) SetServiceAccountCredentials(value *string)() {
    m.serviceAccountCredentials = value
}
