package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementDepOnboardingSettingsItemImportedAppleDeviceIdentitiesImportAppleDeviceIdentityListPostRequestBody provides operations to call the importAppleDeviceIdentityList method.
type DeviceManagementDepOnboardingSettingsItemImportedAppleDeviceIdentitiesImportAppleDeviceIdentityListPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The importedAppleDeviceIdentities property
    importedAppleDeviceIdentities []ImportedAppleDeviceIdentityable
    // The overwriteImportedDeviceIdentities property
    overwriteImportedDeviceIdentities *bool
}
// NewDeviceManagementDepOnboardingSettingsItemImportedAppleDeviceIdentitiesImportAppleDeviceIdentityListPostRequestBody instantiates a new DeviceManagementDepOnboardingSettingsItemImportedAppleDeviceIdentitiesImportAppleDeviceIdentityListPostRequestBody and sets the default values.
func NewDeviceManagementDepOnboardingSettingsItemImportedAppleDeviceIdentitiesImportAppleDeviceIdentityListPostRequestBody()(*DeviceManagementDepOnboardingSettingsItemImportedAppleDeviceIdentitiesImportAppleDeviceIdentityListPostRequestBody) {
    m := &DeviceManagementDepOnboardingSettingsItemImportedAppleDeviceIdentitiesImportAppleDeviceIdentityListPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDeviceManagementDepOnboardingSettingsItemImportedAppleDeviceIdentitiesImportAppleDeviceIdentityListPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementDepOnboardingSettingsItemImportedAppleDeviceIdentitiesImportAppleDeviceIdentityListPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementDepOnboardingSettingsItemImportedAppleDeviceIdentitiesImportAppleDeviceIdentityListPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementDepOnboardingSettingsItemImportedAppleDeviceIdentitiesImportAppleDeviceIdentityListPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementDepOnboardingSettingsItemImportedAppleDeviceIdentitiesImportAppleDeviceIdentityListPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["importedAppleDeviceIdentities"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateImportedAppleDeviceIdentityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ImportedAppleDeviceIdentityable, len(val))
            for i, v := range val {
                res[i] = v.(ImportedAppleDeviceIdentityable)
            }
            m.SetImportedAppleDeviceIdentities(res)
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
// GetImportedAppleDeviceIdentities gets the importedAppleDeviceIdentities property value. The importedAppleDeviceIdentities property
func (m *DeviceManagementDepOnboardingSettingsItemImportedAppleDeviceIdentitiesImportAppleDeviceIdentityListPostRequestBody) GetImportedAppleDeviceIdentities()([]ImportedAppleDeviceIdentityable) {
    return m.importedAppleDeviceIdentities
}
// GetOverwriteImportedDeviceIdentities gets the overwriteImportedDeviceIdentities property value. The overwriteImportedDeviceIdentities property
func (m *DeviceManagementDepOnboardingSettingsItemImportedAppleDeviceIdentitiesImportAppleDeviceIdentityListPostRequestBody) GetOverwriteImportedDeviceIdentities()(*bool) {
    return m.overwriteImportedDeviceIdentities
}
// Serialize serializes information the current object
func (m *DeviceManagementDepOnboardingSettingsItemImportedAppleDeviceIdentitiesImportAppleDeviceIdentityListPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetImportedAppleDeviceIdentities() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetImportedAppleDeviceIdentities()))
        for i, v := range m.GetImportedAppleDeviceIdentities() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
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
func (m *DeviceManagementDepOnboardingSettingsItemImportedAppleDeviceIdentitiesImportAppleDeviceIdentityListPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetImportedAppleDeviceIdentities sets the importedAppleDeviceIdentities property value. The importedAppleDeviceIdentities property
func (m *DeviceManagementDepOnboardingSettingsItemImportedAppleDeviceIdentitiesImportAppleDeviceIdentityListPostRequestBody) SetImportedAppleDeviceIdentities(value []ImportedAppleDeviceIdentityable)() {
    m.importedAppleDeviceIdentities = value
}
// SetOverwriteImportedDeviceIdentities sets the overwriteImportedDeviceIdentities property value. The overwriteImportedDeviceIdentities property
func (m *DeviceManagementDepOnboardingSettingsItemImportedAppleDeviceIdentitiesImportAppleDeviceIdentityListPostRequestBody) SetOverwriteImportedDeviceIdentities(value *bool)() {
    m.overwriteImportedDeviceIdentities = value
}
