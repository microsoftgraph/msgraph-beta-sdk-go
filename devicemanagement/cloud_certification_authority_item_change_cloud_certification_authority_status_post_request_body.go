package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

type CloudCertificationAuthorityItemChangeCloudCertificationAuthorityStatusPostRequestBody struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewCloudCertificationAuthorityItemChangeCloudCertificationAuthorityStatusPostRequestBody instantiates a new CloudCertificationAuthorityItemChangeCloudCertificationAuthorityStatusPostRequestBody and sets the default values.
func NewCloudCertificationAuthorityItemChangeCloudCertificationAuthorityStatusPostRequestBody()(*CloudCertificationAuthorityItemChangeCloudCertificationAuthorityStatusPostRequestBody) {
    m := &CloudCertificationAuthorityItemChangeCloudCertificationAuthorityStatusPostRequestBody{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCloudCertificationAuthorityItemChangeCloudCertificationAuthorityStatusPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCloudCertificationAuthorityItemChangeCloudCertificationAuthorityStatusPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCloudCertificationAuthorityItemChangeCloudCertificationAuthorityStatusPostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CloudCertificationAuthorityItemChangeCloudCertificationAuthorityStatusPostRequestBody) GetAdditionalData()(map[string]any) {
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
func (m *CloudCertificationAuthorityItemChangeCloudCertificationAuthorityStatusPostRequestBody) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetCertificationAuthorityStatus gets the certificationAuthorityStatus property value. Enum type of possible certification authority statuses. These statuses indicate whether a certification authority is currently able to issue certificates or temporarily paused or permanently revoked.
// returns a *CloudCertificationAuthorityStatus when successful
func (m *CloudCertificationAuthorityItemChangeCloudCertificationAuthorityStatusPostRequestBody) GetCertificationAuthorityStatus()(*ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudCertificationAuthorityStatus) {
    val, err := m.GetBackingStore().Get("certificationAuthorityStatus")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudCertificationAuthorityStatus)
    }
    return nil
}
// GetCertificationAuthorityVersion gets the certificationAuthorityVersion property value. The certificationAuthorityVersion property
// returns a *int32 when successful
func (m *CloudCertificationAuthorityItemChangeCloudCertificationAuthorityStatusPostRequestBody) GetCertificationAuthorityVersion()(*int32) {
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
func (m *CloudCertificationAuthorityItemChangeCloudCertificationAuthorityStatusPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["certificationAuthorityStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ParseCloudCertificationAuthorityStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCertificationAuthorityStatus(val.(*ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudCertificationAuthorityStatus))
        }
        return nil
    }
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
    return res
}
// Serialize serializes information the current object
func (m *CloudCertificationAuthorityItemChangeCloudCertificationAuthorityStatusPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetCertificationAuthorityStatus() != nil {
        cast := (*m.GetCertificationAuthorityStatus()).String()
        err := writer.WriteStringValue("certificationAuthorityStatus", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("certificationAuthorityVersion", m.GetCertificationAuthorityVersion())
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
func (m *CloudCertificationAuthorityItemChangeCloudCertificationAuthorityStatusPostRequestBody) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *CloudCertificationAuthorityItemChangeCloudCertificationAuthorityStatusPostRequestBody) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetCertificationAuthorityStatus sets the certificationAuthorityStatus property value. Enum type of possible certification authority statuses. These statuses indicate whether a certification authority is currently able to issue certificates or temporarily paused or permanently revoked.
func (m *CloudCertificationAuthorityItemChangeCloudCertificationAuthorityStatusPostRequestBody) SetCertificationAuthorityStatus(value *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudCertificationAuthorityStatus)() {
    err := m.GetBackingStore().Set("certificationAuthorityStatus", value)
    if err != nil {
        panic(err)
    }
}
// SetCertificationAuthorityVersion sets the certificationAuthorityVersion property value. The certificationAuthorityVersion property
func (m *CloudCertificationAuthorityItemChangeCloudCertificationAuthorityStatusPostRequestBody) SetCertificationAuthorityVersion(value *int32)() {
    err := m.GetBackingStore().Set("certificationAuthorityVersion", value)
    if err != nil {
        panic(err)
    }
}
type CloudCertificationAuthorityItemChangeCloudCertificationAuthorityStatusPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetCertificationAuthorityStatus()(*ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudCertificationAuthorityStatus)
    GetCertificationAuthorityVersion()(*int32)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetCertificationAuthorityStatus(value *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudCertificationAuthorityStatus)()
    SetCertificationAuthorityVersion(value *int32)()
}
