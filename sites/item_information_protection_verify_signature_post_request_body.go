package sites

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// ItemInformationProtectionVerifySignaturePostRequestBody 
type ItemInformationProtectionVerifySignaturePostRequestBody struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewItemInformationProtectionVerifySignaturePostRequestBody instantiates a new ItemInformationProtectionVerifySignaturePostRequestBody and sets the default values.
func NewItemInformationProtectionVerifySignaturePostRequestBody()(*ItemInformationProtectionVerifySignaturePostRequestBody) {
    m := &ItemInformationProtectionVerifySignaturePostRequestBody{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemInformationProtectionVerifySignaturePostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemInformationProtectionVerifySignaturePostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemInformationProtectionVerifySignaturePostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemInformationProtectionVerifySignaturePostRequestBody) GetAdditionalData()(map[string]any) {
    val , err :=  m.backingStore.Get("additionalData")
    if err != nil {
        panic(err)
    }
    if val == nil {
        var value = make(map[string]any);
        m.SetAdditionalData(value);
    }
    return val.(map[string]any)
}
// GetBackingStore gets the BackingStore property value. Stores model information.
func (m *ItemInformationProtectionVerifySignaturePostRequestBody) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetDigest gets the digest property value. The digest property
func (m *ItemInformationProtectionVerifySignaturePostRequestBody) GetDigest()([]byte) {
    val, err := m.GetBackingStore().Get("digest")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]byte)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemInformationProtectionVerifySignaturePostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
func (m *ItemInformationProtectionVerifySignaturePostRequestBody) GetSignature()([]byte) {
    val, err := m.GetBackingStore().Get("signature")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]byte)
    }
    return nil
}
// GetSigningKeyId gets the signingKeyId property value. The signingKeyId property
func (m *ItemInformationProtectionVerifySignaturePostRequestBody) GetSigningKeyId()(*string) {
    val, err := m.GetBackingStore().Get("signingKeyId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *ItemInformationProtectionVerifySignaturePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemInformationProtectionVerifySignaturePostRequestBody) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *ItemInformationProtectionVerifySignaturePostRequestBody) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetDigest sets the digest property value. The digest property
func (m *ItemInformationProtectionVerifySignaturePostRequestBody) SetDigest(value []byte)() {
    err := m.GetBackingStore().Set("digest", value)
    if err != nil {
        panic(err)
    }
}
// SetSignature sets the signature property value. The signature property
func (m *ItemInformationProtectionVerifySignaturePostRequestBody) SetSignature(value []byte)() {
    err := m.GetBackingStore().Set("signature", value)
    if err != nil {
        panic(err)
    }
}
// SetSigningKeyId sets the signingKeyId property value. The signingKeyId property
func (m *ItemInformationProtectionVerifySignaturePostRequestBody) SetSigningKeyId(value *string)() {
    err := m.GetBackingStore().Set("signingKeyId", value)
    if err != nil {
        panic(err)
    }
}
// ItemInformationProtectionVerifySignaturePostRequestBodyable 
type ItemInformationProtectionVerifySignaturePostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetDigest()([]byte)
    GetSignature()([]byte)
    GetSigningKeyId()(*string)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetDigest(value []byte)()
    SetSignature(value []byte)()
    SetSigningKeyId(value *string)()
}
