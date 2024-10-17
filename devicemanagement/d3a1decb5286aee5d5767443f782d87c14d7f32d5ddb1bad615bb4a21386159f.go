package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

type CloudCertificationAuthorityItemUploadExternallySignedCertificationAuthorityCertificatePostRequestBody struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewCloudCertificationAuthorityItemUploadExternallySignedCertificationAuthorityCertificatePostRequestBody instantiates a new CloudCertificationAuthorityItemUploadExternallySignedCertificationAuthorityCertificatePostRequestBody and sets the default values.
func NewCloudCertificationAuthorityItemUploadExternallySignedCertificationAuthorityCertificatePostRequestBody()(*CloudCertificationAuthorityItemUploadExternallySignedCertificationAuthorityCertificatePostRequestBody) {
    m := &CloudCertificationAuthorityItemUploadExternallySignedCertificationAuthorityCertificatePostRequestBody{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCloudCertificationAuthorityItemUploadExternallySignedCertificationAuthorityCertificatePostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCloudCertificationAuthorityItemUploadExternallySignedCertificationAuthorityCertificatePostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCloudCertificationAuthorityItemUploadExternallySignedCertificationAuthorityCertificatePostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CloudCertificationAuthorityItemUploadExternallySignedCertificationAuthorityCertificatePostRequestBody) GetAdditionalData()(map[string]any) {
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
// returns a BackingStore when successful
func (m *CloudCertificationAuthorityItemUploadExternallySignedCertificationAuthorityCertificatePostRequestBody) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetCertificationAuthorityVersion gets the certificationAuthorityVersion property value. The certificationAuthorityVersion property
// returns a *int32 when successful
func (m *CloudCertificationAuthorityItemUploadExternallySignedCertificationAuthorityCertificatePostRequestBody) GetCertificationAuthorityVersion()(*int32) {
    val, err := m.GetBackingStore().Get("certificationAuthorityVersion")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CloudCertificationAuthorityItemUploadExternallySignedCertificationAuthorityCertificatePostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["certificationAuthorityVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCertificationAuthorityVersion(val)
        }
        return nil
    }
    res["signedCertificate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSignedCertificate(val)
        }
        return nil
    }
    res["trustChainCertificates"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateTrustChainCertificateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TrustChainCertificateable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TrustChainCertificateable)
                }
            }
            m.SetTrustChainCertificates(res)
        }
        return nil
    }
    return res
}
// GetSignedCertificate gets the signedCertificate property value. The signedCertificate property
// returns a *string when successful
func (m *CloudCertificationAuthorityItemUploadExternallySignedCertificationAuthorityCertificatePostRequestBody) GetSignedCertificate()(*string) {
    val, err := m.GetBackingStore().Get("signedCertificate")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetTrustChainCertificates gets the trustChainCertificates property value. The trustChainCertificates property
// returns a []TrustChainCertificateable when successful
func (m *CloudCertificationAuthorityItemUploadExternallySignedCertificationAuthorityCertificatePostRequestBody) GetTrustChainCertificates()([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TrustChainCertificateable) {
    val, err := m.GetBackingStore().Get("trustChainCertificates")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TrustChainCertificateable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *CloudCertificationAuthorityItemUploadExternallySignedCertificationAuthorityCertificatePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("certificationAuthorityVersion", m.GetCertificationAuthorityVersion())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("signedCertificate", m.GetSignedCertificate())
        if err != nil {
            return err
        }
    }
    if m.GetTrustChainCertificates() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTrustChainCertificates()))
        for i, v := range m.GetTrustChainCertificates() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("trustChainCertificates", cast)
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
func (m *CloudCertificationAuthorityItemUploadExternallySignedCertificationAuthorityCertificatePostRequestBody) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *CloudCertificationAuthorityItemUploadExternallySignedCertificationAuthorityCertificatePostRequestBody) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetCertificationAuthorityVersion sets the certificationAuthorityVersion property value. The certificationAuthorityVersion property
func (m *CloudCertificationAuthorityItemUploadExternallySignedCertificationAuthorityCertificatePostRequestBody) SetCertificationAuthorityVersion(value *int32)() {
    err := m.GetBackingStore().Set("certificationAuthorityVersion", value)
    if err != nil {
        panic(err)
    }
}
// SetSignedCertificate sets the signedCertificate property value. The signedCertificate property
func (m *CloudCertificationAuthorityItemUploadExternallySignedCertificationAuthorityCertificatePostRequestBody) SetSignedCertificate(value *string)() {
    err := m.GetBackingStore().Set("signedCertificate", value)
    if err != nil {
        panic(err)
    }
}
// SetTrustChainCertificates sets the trustChainCertificates property value. The trustChainCertificates property
func (m *CloudCertificationAuthorityItemUploadExternallySignedCertificationAuthorityCertificatePostRequestBody) SetTrustChainCertificates(value []ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TrustChainCertificateable)() {
    err := m.GetBackingStore().Set("trustChainCertificates", value)
    if err != nil {
        panic(err)
    }
}
type CloudCertificationAuthorityItemUploadExternallySignedCertificationAuthorityCertificatePostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetCertificationAuthorityVersion()(*int32)
    GetSignedCertificate()(*string)
    GetTrustChainCertificates()([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TrustChainCertificateable)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetCertificationAuthorityVersion(value *int32)()
    SetSignedCertificate(value *string)()
    SetTrustChainCertificates(value []ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TrustChainCertificateable)()
}
