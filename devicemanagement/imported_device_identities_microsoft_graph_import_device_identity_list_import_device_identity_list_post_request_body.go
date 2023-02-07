package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// ImportedDeviceIdentitiesMicrosoftGraphImportDeviceIdentityListImportDeviceIdentityListPostRequestBody 
type ImportedDeviceIdentitiesMicrosoftGraphImportDeviceIdentityListImportDeviceIdentityListPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The importedDeviceIdentities property
    importedDeviceIdentities []ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ImportedDeviceIdentityable
    // The overwriteImportedDeviceIdentities property
    overwriteImportedDeviceIdentities *bool
}
// NewImportedDeviceIdentitiesMicrosoftGraphImportDeviceIdentityListImportDeviceIdentityListPostRequestBody instantiates a new ImportedDeviceIdentitiesMicrosoftGraphImportDeviceIdentityListImportDeviceIdentityListPostRequestBody and sets the default values.
func NewImportedDeviceIdentitiesMicrosoftGraphImportDeviceIdentityListImportDeviceIdentityListPostRequestBody()(*ImportedDeviceIdentitiesMicrosoftGraphImportDeviceIdentityListImportDeviceIdentityListPostRequestBody) {
    m := &ImportedDeviceIdentitiesMicrosoftGraphImportDeviceIdentityListImportDeviceIdentityListPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateImportedDeviceIdentitiesMicrosoftGraphImportDeviceIdentityListImportDeviceIdentityListPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateImportedDeviceIdentitiesMicrosoftGraphImportDeviceIdentityListImportDeviceIdentityListPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewImportedDeviceIdentitiesMicrosoftGraphImportDeviceIdentityListImportDeviceIdentityListPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ImportedDeviceIdentitiesMicrosoftGraphImportDeviceIdentityListImportDeviceIdentityListPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ImportedDeviceIdentitiesMicrosoftGraphImportDeviceIdentityListImportDeviceIdentityListPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["overwriteImportedDeviceIdentities"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOverwriteImportedDeviceIdentities(val)
        }
        return nil
    }
    return res
}
// GetImportedDeviceIdentities gets the importedDeviceIdentities property value. The importedDeviceIdentities property
func (m *ImportedDeviceIdentitiesMicrosoftGraphImportDeviceIdentityListImportDeviceIdentityListPostRequestBody) GetImportedDeviceIdentities()([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ImportedDeviceIdentityable) {
    return m.importedDeviceIdentities
}
// GetOverwriteImportedDeviceIdentities gets the overwriteImportedDeviceIdentities property value. The overwriteImportedDeviceIdentities property
func (m *ImportedDeviceIdentitiesMicrosoftGraphImportDeviceIdentityListImportDeviceIdentityListPostRequestBody) GetOverwriteImportedDeviceIdentities()(*bool) {
    return m.overwriteImportedDeviceIdentities
}
// Serialize serializes information the current object
func (m *ImportedDeviceIdentitiesMicrosoftGraphImportDeviceIdentityListImportDeviceIdentityListPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *ImportedDeviceIdentitiesMicrosoftGraphImportDeviceIdentityListImportDeviceIdentityListPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetImportedDeviceIdentities sets the importedDeviceIdentities property value. The importedDeviceIdentities property
func (m *ImportedDeviceIdentitiesMicrosoftGraphImportDeviceIdentityListImportDeviceIdentityListPostRequestBody) SetImportedDeviceIdentities(value []ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ImportedDeviceIdentityable)() {
    m.importedDeviceIdentities = value
}
// SetOverwriteImportedDeviceIdentities sets the overwriteImportedDeviceIdentities property value. The overwriteImportedDeviceIdentities property
func (m *ImportedDeviceIdentitiesMicrosoftGraphImportDeviceIdentityListImportDeviceIdentityListPostRequestBody) SetOverwriteImportedDeviceIdentities(value *bool)() {
    m.overwriteImportedDeviceIdentities = value
}
