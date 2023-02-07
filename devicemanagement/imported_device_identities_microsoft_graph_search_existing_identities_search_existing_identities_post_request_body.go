package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// ImportedDeviceIdentitiesMicrosoftGraphSearchExistingIdentitiesSearchExistingIdentitiesPostRequestBody 
type ImportedDeviceIdentitiesMicrosoftGraphSearchExistingIdentitiesSearchExistingIdentitiesPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The importedDeviceIdentities property
    importedDeviceIdentities []ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ImportedDeviceIdentityable
}
// NewImportedDeviceIdentitiesMicrosoftGraphSearchExistingIdentitiesSearchExistingIdentitiesPostRequestBody instantiates a new ImportedDeviceIdentitiesMicrosoftGraphSearchExistingIdentitiesSearchExistingIdentitiesPostRequestBody and sets the default values.
func NewImportedDeviceIdentitiesMicrosoftGraphSearchExistingIdentitiesSearchExistingIdentitiesPostRequestBody()(*ImportedDeviceIdentitiesMicrosoftGraphSearchExistingIdentitiesSearchExistingIdentitiesPostRequestBody) {
    m := &ImportedDeviceIdentitiesMicrosoftGraphSearchExistingIdentitiesSearchExistingIdentitiesPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateImportedDeviceIdentitiesMicrosoftGraphSearchExistingIdentitiesSearchExistingIdentitiesPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateImportedDeviceIdentitiesMicrosoftGraphSearchExistingIdentitiesSearchExistingIdentitiesPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewImportedDeviceIdentitiesMicrosoftGraphSearchExistingIdentitiesSearchExistingIdentitiesPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ImportedDeviceIdentitiesMicrosoftGraphSearchExistingIdentitiesSearchExistingIdentitiesPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ImportedDeviceIdentitiesMicrosoftGraphSearchExistingIdentitiesSearchExistingIdentitiesPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["importedDeviceIdentities"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateImportedDeviceIdentityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ImportedDeviceIdentityable, len(val))
            for i, v := range val {
                res[i] = v.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ImportedDeviceIdentityable)
            }
            m.SetImportedDeviceIdentities(res)
        }
        return nil
    }
    return res
}
// GetImportedDeviceIdentities gets the importedDeviceIdentities property value. The importedDeviceIdentities property
func (m *ImportedDeviceIdentitiesMicrosoftGraphSearchExistingIdentitiesSearchExistingIdentitiesPostRequestBody) GetImportedDeviceIdentities()([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ImportedDeviceIdentityable) {
    return m.importedDeviceIdentities
}
// Serialize serializes information the current object
func (m *ImportedDeviceIdentitiesMicrosoftGraphSearchExistingIdentitiesSearchExistingIdentitiesPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetImportedDeviceIdentities() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetImportedDeviceIdentities()))
        for i, v := range m.GetImportedDeviceIdentities() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("importedDeviceIdentities", cast)
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
func (m *ImportedDeviceIdentitiesMicrosoftGraphSearchExistingIdentitiesSearchExistingIdentitiesPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetImportedDeviceIdentities sets the importedDeviceIdentities property value. The importedDeviceIdentities property
func (m *ImportedDeviceIdentitiesMicrosoftGraphSearchExistingIdentitiesSearchExistingIdentitiesPostRequestBody) SetImportedDeviceIdentities(value []ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ImportedDeviceIdentityable)() {
    m.importedDeviceIdentities = value
}
