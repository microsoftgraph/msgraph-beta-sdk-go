package windowsprivacyaccesscontrols

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// WindowsPrivacyAccessControlsPostRequestBody provides operations to call the windowsPrivacyAccessControls method.
type WindowsPrivacyAccessControlsPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The windowsPrivacyAccessControls property
    windowsPrivacyAccessControls []ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsPrivacyDataAccessControlItemable
}
// NewWindowsPrivacyAccessControlsPostRequestBody instantiates a new windowsPrivacyAccessControlsPostRequestBody and sets the default values.
func NewWindowsPrivacyAccessControlsPostRequestBody()(*WindowsPrivacyAccessControlsPostRequestBody) {
    m := &WindowsPrivacyAccessControlsPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateWindowsPrivacyAccessControlsPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsPrivacyAccessControlsPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsPrivacyAccessControlsPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WindowsPrivacyAccessControlsPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsPrivacyAccessControlsPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["windowsPrivacyAccessControls"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateWindowsPrivacyDataAccessControlItemFromDiscriminatorValue , m.SetWindowsPrivacyAccessControls)
    return res
}
// GetWindowsPrivacyAccessControls gets the windowsPrivacyAccessControls property value. The windowsPrivacyAccessControls property
func (m *WindowsPrivacyAccessControlsPostRequestBody) GetWindowsPrivacyAccessControls()([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsPrivacyDataAccessControlItemable) {
    return m.windowsPrivacyAccessControls
}
// Serialize serializes information the current object
func (m *WindowsPrivacyAccessControlsPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetWindowsPrivacyAccessControls() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetWindowsPrivacyAccessControls())
        err := writer.WriteCollectionOfObjectValues("windowsPrivacyAccessControls", cast)
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
func (m *WindowsPrivacyAccessControlsPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetWindowsPrivacyAccessControls sets the windowsPrivacyAccessControls property value. The windowsPrivacyAccessControls property
func (m *WindowsPrivacyAccessControlsPostRequestBody) SetWindowsPrivacyAccessControls(value []ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsPrivacyDataAccessControlItemable)() {
    m.windowsPrivacyAccessControls = value
}
