package importappledeviceidentitylist

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// ImportAppleDeviceIdentityListPostRequestBody provides operations to call the importAppleDeviceIdentityList method.
type ImportAppleDeviceIdentityListPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The importedAppleDeviceIdentities property
    importedAppleDeviceIdentities []ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ImportedAppleDeviceIdentityable
    // The overwriteImportedDeviceIdentities property
    overwriteImportedDeviceIdentities *bool
}
// NewImportAppleDeviceIdentityListPostRequestBody instantiates a new importAppleDeviceIdentityListPostRequestBody and sets the default values.
func NewImportAppleDeviceIdentityListPostRequestBody()(*ImportAppleDeviceIdentityListPostRequestBody) {
    m := &ImportAppleDeviceIdentityListPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateImportAppleDeviceIdentityListPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateImportAppleDeviceIdentityListPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewImportAppleDeviceIdentityListPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ImportAppleDeviceIdentityListPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ImportAppleDeviceIdentityListPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["importedAppleDeviceIdentities"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateImportedAppleDeviceIdentityFromDiscriminatorValue , m.SetImportedAppleDeviceIdentities)
    res["overwriteImportedDeviceIdentities"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetOverwriteImportedDeviceIdentities)
    return res
}
// GetImportedAppleDeviceIdentities gets the importedAppleDeviceIdentities property value. The importedAppleDeviceIdentities property
func (m *ImportAppleDeviceIdentityListPostRequestBody) GetImportedAppleDeviceIdentities()([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ImportedAppleDeviceIdentityable) {
    return m.importedAppleDeviceIdentities
}
// GetOverwriteImportedDeviceIdentities gets the overwriteImportedDeviceIdentities property value. The overwriteImportedDeviceIdentities property
func (m *ImportAppleDeviceIdentityListPostRequestBody) GetOverwriteImportedDeviceIdentities()(*bool) {
    return m.overwriteImportedDeviceIdentities
}
// Serialize serializes information the current object
func (m *ImportAppleDeviceIdentityListPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetImportedAppleDeviceIdentities() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetImportedAppleDeviceIdentities())
        err := writer.WriteCollectionOfObjectValues("importedAppleDeviceIdentities", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("overwriteImportedDeviceIdentities", m.GetOverwriteImportedDeviceIdentities())
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
func (m *ImportAppleDeviceIdentityListPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetImportedAppleDeviceIdentities sets the importedAppleDeviceIdentities property value. The importedAppleDeviceIdentities property
func (m *ImportAppleDeviceIdentityListPostRequestBody) SetImportedAppleDeviceIdentities(value []ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ImportedAppleDeviceIdentityable)() {
    m.importedAppleDeviceIdentities = value
}
// SetOverwriteImportedDeviceIdentities sets the overwriteImportedDeviceIdentities property value. The overwriteImportedDeviceIdentities property
func (m *ImportAppleDeviceIdentityListPostRequestBody) SetOverwriteImportedDeviceIdentities(value *bool)() {
    m.overwriteImportedDeviceIdentities = value
}
