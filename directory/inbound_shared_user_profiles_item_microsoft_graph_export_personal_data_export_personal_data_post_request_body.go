package directory

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// InboundSharedUserProfilesItemMicrosoftGraphExportPersonalDataExportPersonalDataPostRequestBody 
type InboundSharedUserProfilesItemMicrosoftGraphExportPersonalDataExportPersonalDataPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The storageLocation property
    storageLocation *string
}
// NewInboundSharedUserProfilesItemMicrosoftGraphExportPersonalDataExportPersonalDataPostRequestBody instantiates a new InboundSharedUserProfilesItemMicrosoftGraphExportPersonalDataExportPersonalDataPostRequestBody and sets the default values.
func NewInboundSharedUserProfilesItemMicrosoftGraphExportPersonalDataExportPersonalDataPostRequestBody()(*InboundSharedUserProfilesItemMicrosoftGraphExportPersonalDataExportPersonalDataPostRequestBody) {
    m := &InboundSharedUserProfilesItemMicrosoftGraphExportPersonalDataExportPersonalDataPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateInboundSharedUserProfilesItemMicrosoftGraphExportPersonalDataExportPersonalDataPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateInboundSharedUserProfilesItemMicrosoftGraphExportPersonalDataExportPersonalDataPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewInboundSharedUserProfilesItemMicrosoftGraphExportPersonalDataExportPersonalDataPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *InboundSharedUserProfilesItemMicrosoftGraphExportPersonalDataExportPersonalDataPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *InboundSharedUserProfilesItemMicrosoftGraphExportPersonalDataExportPersonalDataPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["storageLocation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStorageLocation(val)
        }
        return nil
    }
    return res
}
// GetStorageLocation gets the storageLocation property value. The storageLocation property
func (m *InboundSharedUserProfilesItemMicrosoftGraphExportPersonalDataExportPersonalDataPostRequestBody) GetStorageLocation()(*string) {
    return m.storageLocation
}
// Serialize serializes information the current object
func (m *InboundSharedUserProfilesItemMicrosoftGraphExportPersonalDataExportPersonalDataPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("storageLocation", m.GetStorageLocation())
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
func (m *InboundSharedUserProfilesItemMicrosoftGraphExportPersonalDataExportPersonalDataPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetStorageLocation sets the storageLocation property value. The storageLocation property
func (m *InboundSharedUserProfilesItemMicrosoftGraphExportPersonalDataExportPersonalDataPostRequestBody) SetStorageLocation(value *string)() {
    m.storageLocation = value
}
