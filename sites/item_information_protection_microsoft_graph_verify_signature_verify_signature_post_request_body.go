package sites

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemInformationProtectionMicrosoftGraphVerifySignatureVerifySignaturePostRequestBody 
type ItemInformationProtectionMicrosoftGraphVerifySignatureVerifySignaturePostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The digest property
    digest []byte
    // The signature property
    signature []byte
    // The signingKeyId property
    signingKeyId *string
}
// NewItemInformationProtectionMicrosoftGraphVerifySignatureVerifySignaturePostRequestBody instantiates a new ItemInformationProtectionMicrosoftGraphVerifySignatureVerifySignaturePostRequestBody and sets the default values.
func NewItemInformationProtectionMicrosoftGraphVerifySignatureVerifySignaturePostRequestBody()(*ItemInformationProtectionMicrosoftGraphVerifySignatureVerifySignaturePostRequestBody) {
    m := &ItemInformationProtectionMicrosoftGraphVerifySignatureVerifySignaturePostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemInformationProtectionMicrosoftGraphVerifySignatureVerifySignaturePostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemInformationProtectionMicrosoftGraphVerifySignatureVerifySignaturePostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemInformationProtectionMicrosoftGraphVerifySignatureVerifySignaturePostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemInformationProtectionMicrosoftGraphVerifySignatureVerifySignaturePostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDigest gets the digest property value. The digest property
func (m *ItemInformationProtectionMicrosoftGraphVerifySignatureVerifySignaturePostRequestBody) GetDigest()([]byte) {
    return m.digest
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemInformationProtectionMicrosoftGraphVerifySignatureVerifySignaturePostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["digest"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDigest(val)
        }
        return nil
    }
    res["signature"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSignature(val)
        }
        return nil
    }
    res["signingKeyId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSigningKeyId(val)
        }
        return nil
    }
    return res
}
// GetSignature gets the signature property value. The signature property
func (m *ItemInformationProtectionMicrosoftGraphVerifySignatureVerifySignaturePostRequestBody) GetSignature()([]byte) {
    return m.signature
}
// GetSigningKeyId gets the signingKeyId property value. The signingKeyId property
func (m *ItemInformationProtectionMicrosoftGraphVerifySignatureVerifySignaturePostRequestBody) GetSigningKeyId()(*string) {
    return m.signingKeyId
}
// Serialize serializes information the current object
func (m *ItemInformationProtectionMicrosoftGraphVerifySignatureVerifySignaturePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteByteArrayValue("digest", m.GetDigest())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteByteArrayValue("signature", m.GetSignature())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("signingKeyId", m.GetSigningKeyId())
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
func (m *ItemInformationProtectionMicrosoftGraphVerifySignatureVerifySignaturePostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDigest sets the digest property value. The digest property
func (m *ItemInformationProtectionMicrosoftGraphVerifySignatureVerifySignaturePostRequestBody) SetDigest(value []byte)() {
    m.digest = value
}
// SetSignature sets the signature property value. The signature property
func (m *ItemInformationProtectionMicrosoftGraphVerifySignatureVerifySignaturePostRequestBody) SetSignature(value []byte)() {
    m.signature = value
}
// SetSigningKeyId sets the signingKeyId property value. The signingKeyId property
func (m *ItemInformationProtectionMicrosoftGraphVerifySignatureVerifySignaturePostRequestBody) SetSigningKeyId(value *string)() {
    m.signingKeyId = value
}
