package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementResourceAccessProfilesQueryByPlatformTypePostRequestBody provides operations to call the queryByPlatformType method.
type DeviceManagementResourceAccessProfilesQueryByPlatformTypePostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Supported platform types for policies.
    platformType *PolicyPlatformType
}
// NewDeviceManagementResourceAccessProfilesQueryByPlatformTypePostRequestBody instantiates a new DeviceManagementResourceAccessProfilesQueryByPlatformTypePostRequestBody and sets the default values.
func NewDeviceManagementResourceAccessProfilesQueryByPlatformTypePostRequestBody()(*DeviceManagementResourceAccessProfilesQueryByPlatformTypePostRequestBody) {
    m := &DeviceManagementResourceAccessProfilesQueryByPlatformTypePostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDeviceManagementResourceAccessProfilesQueryByPlatformTypePostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementResourceAccessProfilesQueryByPlatformTypePostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementResourceAccessProfilesQueryByPlatformTypePostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementResourceAccessProfilesQueryByPlatformTypePostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementResourceAccessProfilesQueryByPlatformTypePostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["platformType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParsePolicyPlatformType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPlatformType(val.(*PolicyPlatformType))
        }
        return nil
    }
    return res
}
// GetPlatformType gets the platformType property value. Supported platform types for policies.
func (m *DeviceManagementResourceAccessProfilesQueryByPlatformTypePostRequestBody) GetPlatformType()(*PolicyPlatformType) {
    return m.platformType
}
// Serialize serializes information the current object
func (m *DeviceManagementResourceAccessProfilesQueryByPlatformTypePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetPlatformType() != nil {
        cast := (*m.GetPlatformType()).String()
        err := writer.WriteStringValue("platformType", &cast)
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
func (m *DeviceManagementResourceAccessProfilesQueryByPlatformTypePostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetPlatformType sets the platformType property value. Supported platform types for policies.
func (m *DeviceManagementResourceAccessProfilesQueryByPlatformTypePostRequestBody) SetPlatformType(value *PolicyPlatformType)() {
    m.platformType = value
}
