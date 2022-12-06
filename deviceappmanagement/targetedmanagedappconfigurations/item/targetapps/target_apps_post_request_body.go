package targetapps

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// TargetAppsPostRequestBody provides operations to call the targetApps method.
type TargetAppsPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The appGroupType property
    appGroupType *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TargetedManagedAppGroupType
    // The apps property
    apps []ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedMobileAppable
}
// NewTargetAppsPostRequestBody instantiates a new targetAppsPostRequestBody and sets the default values.
func NewTargetAppsPostRequestBody()(*TargetAppsPostRequestBody) {
    m := &TargetAppsPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateTargetAppsPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTargetAppsPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTargetAppsPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TargetAppsPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetAppGroupType gets the appGroupType property value. The appGroupType property
func (m *TargetAppsPostRequestBody) GetAppGroupType()(*ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TargetedManagedAppGroupType) {
    return m.appGroupType
}
// GetApps gets the apps property value. The apps property
func (m *TargetAppsPostRequestBody) GetApps()([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedMobileAppable) {
    return m.apps
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TargetAppsPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["appGroupType"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ParseTargetedManagedAppGroupType , m.SetAppGroupType)
    res["apps"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateManagedMobileAppFromDiscriminatorValue , m.SetApps)
    return res
}
// Serialize serializes information the current object
func (m *TargetAppsPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAppGroupType() != nil {
        cast := (*m.GetAppGroupType()).String()
        err := writer.WriteStringValue("appGroupType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetApps() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetApps())
        err := writer.WriteCollectionOfObjectValues("apps", cast)
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
func (m *TargetAppsPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetAppGroupType sets the appGroupType property value. The appGroupType property
func (m *TargetAppsPostRequestBody) SetAppGroupType(value *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TargetedManagedAppGroupType)() {
    m.appGroupType = value
}
// SetApps sets the apps property value. The apps property
func (m *TargetAppsPostRequestBody) SetApps(value []ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedMobileAppable)() {
    m.apps = value
}
