package connect

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ConnectPostRequestBody provides operations to call the connect method.
type ConnectPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The ownerAccessToken property
    ownerAccessToken *string
    // The ownerUserPrincipalName property
    ownerUserPrincipalName *string
}
// NewConnectPostRequestBody instantiates a new connectPostRequestBody and sets the default values.
func NewConnectPostRequestBody()(*ConnectPostRequestBody) {
    m := &ConnectPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateConnectPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateConnectPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewConnectPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ConnectPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ConnectPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["ownerAccessToken"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOwnerAccessToken(val)
        }
        return nil
    }
    res["ownerUserPrincipalName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
// GetOwnerAccessToken gets the ownerAccessToken property value. The ownerAccessToken property
func (m *ConnectPostRequestBody) GetOwnerAccessToken()(*string) {
    return m.ownerAccessToken
}
// GetOwnerUserPrincipalName gets the ownerUserPrincipalName property value. The ownerUserPrincipalName property
func (m *ConnectPostRequestBody) GetOwnerUserPrincipalName()(*string) {
    return m.ownerUserPrincipalName
}
// Serialize serializes information the current object
func (m *ConnectPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *ConnectPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetOwnerAccessToken sets the ownerAccessToken property value. The ownerAccessToken property
func (m *ConnectPostRequestBody) SetOwnerAccessToken(value *string)() {
    m.ownerAccessToken = value
}
// SetOwnerUserPrincipalName sets the ownerUserPrincipalName property value. The ownerUserPrincipalName property
func (m *ConnectPostRequestBody) SetOwnerUserPrincipalName(value *string)() {
    m.ownerUserPrincipalName = value
}
