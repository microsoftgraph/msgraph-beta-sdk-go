package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TrustFrameworkKeySetsItemUploadPkcs12PostRequestBody provides operations to call the uploadPkcs12 method.
type TrustFrameworkKeySetsItemUploadPkcs12PostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The key property
    key *string
    // The password property
    password *string
}
// NewTrustFrameworkKeySetsItemUploadPkcs12PostRequestBody instantiates a new TrustFrameworkKeySetsItemUploadPkcs12PostRequestBody and sets the default values.
func NewTrustFrameworkKeySetsItemUploadPkcs12PostRequestBody()(*TrustFrameworkKeySetsItemUploadPkcs12PostRequestBody) {
    m := &TrustFrameworkKeySetsItemUploadPkcs12PostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateTrustFrameworkKeySetsItemUploadPkcs12PostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTrustFrameworkKeySetsItemUploadPkcs12PostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTrustFrameworkKeySetsItemUploadPkcs12PostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TrustFrameworkKeySetsItemUploadPkcs12PostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TrustFrameworkKeySetsItemUploadPkcs12PostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["key"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKey(val)
        }
        return nil
    }
    res["password"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
// GetKey gets the key property value. The key property
func (m *TrustFrameworkKeySetsItemUploadPkcs12PostRequestBody) GetKey()(*string) {
    return m.key
}
// GetPassword gets the password property value. The password property
func (m *TrustFrameworkKeySetsItemUploadPkcs12PostRequestBody) GetPassword()(*string) {
    return m.password
}
// Serialize serializes information the current object
func (m *TrustFrameworkKeySetsItemUploadPkcs12PostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *TrustFrameworkKeySetsItemUploadPkcs12PostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetKey sets the key property value. The key property
func (m *TrustFrameworkKeySetsItemUploadPkcs12PostRequestBody) SetKey(value *string)() {
    m.key = value
}
// SetPassword sets the password property value. The password property
func (m *TrustFrameworkKeySetsItemUploadPkcs12PostRequestBody) SetPassword(value *string)() {
    m.password = value
}
