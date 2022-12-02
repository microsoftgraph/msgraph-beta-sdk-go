package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementImportedDeviceIdentitiesImportDeviceIdentityListPostRequestBody provides operations to call the importDeviceIdentityList method.
type DeviceManagementImportedDeviceIdentitiesImportDeviceIdentityListPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The importedDeviceIdentities property
    importedDeviceIdentities []ImportedDeviceIdentityable
    // The overwriteImportedDeviceIdentities property
    overwriteImportedDeviceIdentities *bool
}
// NewDeviceManagementImportedDeviceIdentitiesImportDeviceIdentityListPostRequestBody instantiates a new DeviceManagementImportedDeviceIdentitiesImportDeviceIdentityListPostRequestBody and sets the default values.
func NewDeviceManagementImportedDeviceIdentitiesImportDeviceIdentityListPostRequestBody()(*DeviceManagementImportedDeviceIdentitiesImportDeviceIdentityListPostRequestBody) {
    m := &DeviceManagementImportedDeviceIdentitiesImportDeviceIdentityListPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDeviceManagementImportedDeviceIdentitiesImportDeviceIdentityListPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementImportedDeviceIdentitiesImportDeviceIdentityListPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementImportedDeviceIdentitiesImportDeviceIdentityListPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementImportedDeviceIdentitiesImportDeviceIdentityListPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementImportedDeviceIdentitiesImportDeviceIdentityListPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["importedDeviceIdentities"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateImportedDeviceIdentityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ImportedDeviceIdentityable, len(val))
            for i, v := range val {
                res[i] = v.(ImportedDeviceIdentityable)
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
func (m *DeviceManagementImportedDeviceIdentitiesImportDeviceIdentityListPostRequestBody) GetImportedDeviceIdentities()([]ImportedDeviceIdentityable) {
    return m.importedDeviceIdentities
}
// GetOverwriteImportedDeviceIdentities gets the overwriteImportedDeviceIdentities property value. The overwriteImportedDeviceIdentities property
func (m *DeviceManagementImportedDeviceIdentitiesImportDeviceIdentityListPostRequestBody) GetOverwriteImportedDeviceIdentities()(*bool) {
    return m.overwriteImportedDeviceIdentities
}
// Serialize serializes information the current object
func (m *DeviceManagementImportedDeviceIdentitiesImportDeviceIdentityListPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *DeviceManagementImportedDeviceIdentitiesImportDeviceIdentityListPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetImportedDeviceIdentities sets the importedDeviceIdentities property value. The importedDeviceIdentities property
func (m *DeviceManagementImportedDeviceIdentitiesImportDeviceIdentityListPostRequestBody) SetImportedDeviceIdentities(value []ImportedDeviceIdentityable)() {
    m.importedDeviceIdentities = value
}
// SetOverwriteImportedDeviceIdentities sets the overwriteImportedDeviceIdentities property value. The overwriteImportedDeviceIdentities property
func (m *DeviceManagementImportedDeviceIdentitiesImportDeviceIdentityListPostRequestBody) SetOverwriteImportedDeviceIdentities(value *bool)() {
    m.overwriteImportedDeviceIdentities = value
}
