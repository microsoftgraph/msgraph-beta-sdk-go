package groups

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemSitesItemInformationProtectionMicrosoftGraphSignDigestSignDigestPostRequestBody 
type ItemSitesItemInformationProtectionMicrosoftGraphSignDigestSignDigestPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The digest property
    digest []byte
}
// NewItemSitesItemInformationProtectionMicrosoftGraphSignDigestSignDigestPostRequestBody instantiates a new ItemSitesItemInformationProtectionMicrosoftGraphSignDigestSignDigestPostRequestBody and sets the default values.
func NewItemSitesItemInformationProtectionMicrosoftGraphSignDigestSignDigestPostRequestBody()(*ItemSitesItemInformationProtectionMicrosoftGraphSignDigestSignDigestPostRequestBody) {
    m := &ItemSitesItemInformationProtectionMicrosoftGraphSignDigestSignDigestPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemSitesItemInformationProtectionMicrosoftGraphSignDigestSignDigestPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemSitesItemInformationProtectionMicrosoftGraphSignDigestSignDigestPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemSitesItemInformationProtectionMicrosoftGraphSignDigestSignDigestPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemSitesItemInformationProtectionMicrosoftGraphSignDigestSignDigestPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDigest gets the digest property value. The digest property
func (m *ItemSitesItemInformationProtectionMicrosoftGraphSignDigestSignDigestPostRequestBody) GetDigest()([]byte) {
    return m.digest
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemSitesItemInformationProtectionMicrosoftGraphSignDigestSignDigestPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
func (m *ItemSitesItemInformationProtectionMicrosoftGraphSignDigestSignDigestPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *ItemSitesItemInformationProtectionMicrosoftGraphSignDigestSignDigestPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDigest sets the digest property value. The digest property
func (m *ItemSitesItemInformationProtectionMicrosoftGraphSignDigestSignDigestPostRequestBody) SetDigest(value []byte)() {
    m.digest = value
}
