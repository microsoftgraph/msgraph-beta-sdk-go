package getrelatedappstateswithuserprincipalnamewithdeviceid

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// GetRelatedAppStatesWithUserPrincipalNameWithDeviceIdResponse provides operations to call the getRelatedAppStates method.
type GetRelatedAppStatesWithUserPrincipalNameWithDeviceIdResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The value property
    value []ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MobileAppRelationshipStateable
}
// NewGetRelatedAppStatesWithUserPrincipalNameWithDeviceIdResponse instantiates a new getRelatedAppStatesWithUserPrincipalNameWithDeviceIdResponse and sets the default values.
func NewGetRelatedAppStatesWithUserPrincipalNameWithDeviceIdResponse()(*GetRelatedAppStatesWithUserPrincipalNameWithDeviceIdResponse) {
    m := &GetRelatedAppStatesWithUserPrincipalNameWithDeviceIdResponse{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateGetRelatedAppStatesWithUserPrincipalNameWithDeviceIdResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGetRelatedAppStatesWithUserPrincipalNameWithDeviceIdResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGetRelatedAppStatesWithUserPrincipalNameWithDeviceIdResponse(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *GetRelatedAppStatesWithUserPrincipalNameWithDeviceIdResponse) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GetRelatedAppStatesWithUserPrincipalNameWithDeviceIdResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["value"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateMobileAppRelationshipStateFromDiscriminatorValue , m.SetValue)
    return res
}
// GetValue gets the value property value. The value property
func (m *GetRelatedAppStatesWithUserPrincipalNameWithDeviceIdResponse) GetValue()([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MobileAppRelationshipStateable) {
    return m.value
}
// Serialize serializes information the current object
func (m *GetRelatedAppStatesWithUserPrincipalNameWithDeviceIdResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetValue() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetValue())
        err := writer.WriteCollectionOfObjectValues("value", cast)
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
func (m *GetRelatedAppStatesWithUserPrincipalNameWithDeviceIdResponse) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetValue sets the value property value. The value property
func (m *GetRelatedAppStatesWithUserPrincipalNameWithDeviceIdResponse) SetValue(value []ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MobileAppRelationshipStateable)() {
    m.value = value
}
