package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceAppManagementMobileAppsItemUpdateRelationshipsPostRequestBody provides operations to call the updateRelationships method.
type DeviceAppManagementMobileAppsItemUpdateRelationshipsPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The relationships property
    relationships []MobileAppRelationshipable
}
// NewDeviceAppManagementMobileAppsItemUpdateRelationshipsPostRequestBody instantiates a new DeviceAppManagementMobileAppsItemUpdateRelationshipsPostRequestBody and sets the default values.
func NewDeviceAppManagementMobileAppsItemUpdateRelationshipsPostRequestBody()(*DeviceAppManagementMobileAppsItemUpdateRelationshipsPostRequestBody) {
    m := &DeviceAppManagementMobileAppsItemUpdateRelationshipsPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDeviceAppManagementMobileAppsItemUpdateRelationshipsPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceAppManagementMobileAppsItemUpdateRelationshipsPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceAppManagementMobileAppsItemUpdateRelationshipsPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceAppManagementMobileAppsItemUpdateRelationshipsPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceAppManagementMobileAppsItemUpdateRelationshipsPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["relationships"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMobileAppRelationshipFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MobileAppRelationshipable, len(val))
            for i, v := range val {
                res[i] = v.(MobileAppRelationshipable)
            }
            m.SetRelationships(res)
        }
        return nil
    }
    return res
}
// GetRelationships gets the relationships property value. The relationships property
func (m *DeviceAppManagementMobileAppsItemUpdateRelationshipsPostRequestBody) GetRelationships()([]MobileAppRelationshipable) {
    return m.relationships
}
// Serialize serializes information the current object
func (m *DeviceAppManagementMobileAppsItemUpdateRelationshipsPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetRelationships() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRelationships()))
        for i, v := range m.GetRelationships() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("relationships", cast)
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
func (m *DeviceAppManagementMobileAppsItemUpdateRelationshipsPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetRelationships sets the relationships property value. The relationships property
func (m *DeviceAppManagementMobileAppsItemUpdateRelationshipsPostRequestBody) SetRelationships(value []MobileAppRelationshipable)() {
    m.relationships = value
}
