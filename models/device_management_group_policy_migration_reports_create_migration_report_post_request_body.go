package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementGroupPolicyMigrationReportsCreateMigrationReportPostRequestBody provides operations to call the createMigrationReport method.
type DeviceManagementGroupPolicyMigrationReportsCreateMigrationReportPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The groupPolicyObjectFile property
    groupPolicyObjectFile GroupPolicyObjectFileable
}
// NewDeviceManagementGroupPolicyMigrationReportsCreateMigrationReportPostRequestBody instantiates a new DeviceManagementGroupPolicyMigrationReportsCreateMigrationReportPostRequestBody and sets the default values.
func NewDeviceManagementGroupPolicyMigrationReportsCreateMigrationReportPostRequestBody()(*DeviceManagementGroupPolicyMigrationReportsCreateMigrationReportPostRequestBody) {
    m := &DeviceManagementGroupPolicyMigrationReportsCreateMigrationReportPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDeviceManagementGroupPolicyMigrationReportsCreateMigrationReportPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementGroupPolicyMigrationReportsCreateMigrationReportPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementGroupPolicyMigrationReportsCreateMigrationReportPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementGroupPolicyMigrationReportsCreateMigrationReportPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementGroupPolicyMigrationReportsCreateMigrationReportPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["groupPolicyObjectFile"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateGroupPolicyObjectFileFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGroupPolicyObjectFile(val.(GroupPolicyObjectFileable))
        }
        return nil
    }
    return res
}
// GetGroupPolicyObjectFile gets the groupPolicyObjectFile property value. The groupPolicyObjectFile property
func (m *DeviceManagementGroupPolicyMigrationReportsCreateMigrationReportPostRequestBody) GetGroupPolicyObjectFile()(GroupPolicyObjectFileable) {
    return m.groupPolicyObjectFile
}
// Serialize serializes information the current object
func (m *DeviceManagementGroupPolicyMigrationReportsCreateMigrationReportPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("groupPolicyObjectFile", m.GetGroupPolicyObjectFile())
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
func (m *DeviceManagementGroupPolicyMigrationReportsCreateMigrationReportPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetGroupPolicyObjectFile sets the groupPolicyObjectFile property value. The groupPolicyObjectFile property
func (m *DeviceManagementGroupPolicyMigrationReportsCreateMigrationReportPostRequestBody) SetGroupPolicyObjectFile(value GroupPolicyObjectFileable)() {
    m.groupPolicyObjectFile = value
}
