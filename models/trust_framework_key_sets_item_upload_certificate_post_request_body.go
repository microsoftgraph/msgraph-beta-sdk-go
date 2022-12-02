package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TrustFrameworkKeySetsItemUploadCertificatePostRequestBody provides operations to call the uploadCertificate method.
type TrustFrameworkKeySetsItemUploadCertificatePostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The key property
    key *string
}
// NewTrustFrameworkKeySetsItemUploadCertificatePostRequestBody instantiates a new TrustFrameworkKeySetsItemUploadCertificatePostRequestBody and sets the default values.
func NewTrustFrameworkKeySetsItemUploadCertificatePostRequestBody()(*TrustFrameworkKeySetsItemUploadCertificatePostRequestBody) {
    m := &TrustFrameworkKeySetsItemUploadCertificatePostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateTrustFrameworkKeySetsItemUploadCertificatePostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTrustFrameworkKeySetsItemUploadCertificatePostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTrustFrameworkKeySetsItemUploadCertificatePostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TrustFrameworkKeySetsItemUploadCertificatePostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TrustFrameworkKeySetsItemUploadCertificatePostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    return res
}
// GetKey gets the key property value. The key property
func (m *TrustFrameworkKeySetsItemUploadCertificatePostRequestBody) GetKey()(*string) {
    return m.key
}
// Serialize serializes information the current object
func (m *TrustFrameworkKeySetsItemUploadCertificatePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("key", m.GetKey())
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
func (m *TrustFrameworkKeySetsItemUploadCertificatePostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetKey sets the key property value. The key property
func (m *TrustFrameworkKeySetsItemUploadCertificatePostRequestBody) SetKey(value *string)() {
    m.key = value
}
