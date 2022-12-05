package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementDepOnboardingSettingsItemEnrollmentProfilesItemUpdateDeviceProfileAssignmentPostRequestBody provides operations to call the updateDeviceProfileAssignment method.
type DeviceManagementDepOnboardingSettingsItemEnrollmentProfilesItemUpdateDeviceProfileAssignmentPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The deviceIds property
    deviceIds []string
}
// NewDeviceManagementDepOnboardingSettingsItemEnrollmentProfilesItemUpdateDeviceProfileAssignmentPostRequestBody instantiates a new DeviceManagementDepOnboardingSettingsItemEnrollmentProfilesItemUpdateDeviceProfileAssignmentPostRequestBody and sets the default values.
func NewDeviceManagementDepOnboardingSettingsItemEnrollmentProfilesItemUpdateDeviceProfileAssignmentPostRequestBody()(*DeviceManagementDepOnboardingSettingsItemEnrollmentProfilesItemUpdateDeviceProfileAssignmentPostRequestBody) {
    m := &DeviceManagementDepOnboardingSettingsItemEnrollmentProfilesItemUpdateDeviceProfileAssignmentPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDeviceManagementDepOnboardingSettingsItemEnrollmentProfilesItemUpdateDeviceProfileAssignmentPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementDepOnboardingSettingsItemEnrollmentProfilesItemUpdateDeviceProfileAssignmentPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementDepOnboardingSettingsItemEnrollmentProfilesItemUpdateDeviceProfileAssignmentPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementDepOnboardingSettingsItemEnrollmentProfilesItemUpdateDeviceProfileAssignmentPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetDeviceIds gets the deviceIds property value. The deviceIds property
func (m *DeviceManagementDepOnboardingSettingsItemEnrollmentProfilesItemUpdateDeviceProfileAssignmentPostRequestBody) GetDeviceIds()([]string) {
    return m.deviceIds
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementDepOnboardingSettingsItemEnrollmentProfilesItemUpdateDeviceProfileAssignmentPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["deviceIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetDeviceIds(res)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *DeviceManagementDepOnboardingSettingsItemEnrollmentProfilesItemUpdateDeviceProfileAssignmentPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetDeviceIds() != nil {
        err := writer.WriteCollectionOfStringValues("deviceIds", m.GetDeviceIds())
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
func (m *DeviceManagementDepOnboardingSettingsItemEnrollmentProfilesItemUpdateDeviceProfileAssignmentPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetDeviceIds sets the deviceIds property value. The deviceIds property
func (m *DeviceManagementDepOnboardingSettingsItemEnrollmentProfilesItemUpdateDeviceProfileAssignmentPostRequestBody) SetDeviceIds(value []string)() {
    m.deviceIds = value
}
