package me

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// InformationProtectionMicrosoftGraphSignDigestSignDigestPostRequestBody 
type InformationProtectionMicrosoftGraphSignDigestSignDigestPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The digest property
    digest []byte
}
// NewInformationProtectionMicrosoftGraphSignDigestSignDigestPostRequestBody instantiates a new InformationProtectionMicrosoftGraphSignDigestSignDigestPostRequestBody and sets the default values.
func NewInformationProtectionMicrosoftGraphSignDigestSignDigestPostRequestBody()(*InformationProtectionMicrosoftGraphSignDigestSignDigestPostRequestBody) {
    m := &InformationProtectionMicrosoftGraphSignDigestSignDigestPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateInformationProtectionMicrosoftGraphSignDigestSignDigestPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateInformationProtectionMicrosoftGraphSignDigestSignDigestPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewInformationProtectionMicrosoftGraphSignDigestSignDigestPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *InformationProtectionMicrosoftGraphSignDigestSignDigestPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDigest gets the digest property value. The digest property
func (m *InformationProtectionMicrosoftGraphSignDigestSignDigestPostRequestBody) GetDigest()([]byte) {
    return m.digest
}
// GetFieldDeserializers the deserialization information for the current model
func (m *InformationProtectionMicrosoftGraphSignDigestSignDigestPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    return res
}
// Serialize serializes information the current object
func (m *InformationProtectionMicrosoftGraphSignDigestSignDigestPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteByteArrayValue("digest", m.GetDigest())
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
func (m *InformationProtectionMicrosoftGraphSignDigestSignDigestPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDigest sets the digest property value. The digest property
func (m *InformationProtectionMicrosoftGraphSignDigestSignDigestPostRequestBody) SetDigest(value []byte)() {
    m.digest = value
}
