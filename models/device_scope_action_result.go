package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceScopeActionResult the result of the triggered device scope action.
type DeviceScopeActionResult struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Trigger on the service to either START or STOP computing metrics data based on a device scope configuration.
    deviceScopeAction *string
    // The unique identifier of the device scope the action was triggered on.
    deviceScopeId *string
    // The message indicates the reason the device scope action failed to trigger.
    failedMessage *string
    // The OdataType property
    odataType *string
    // Indicates the status of the attempted device scope action
    status *DeviceScopeActionStatus
}
// NewDeviceScopeActionResult instantiates a new deviceScopeActionResult and sets the default values.
func NewDeviceScopeActionResult()(*DeviceScopeActionResult) {
    m := &DeviceScopeActionResult{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    odataTypeValue := "#microsoft.graph.deviceScopeActionResult";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateDeviceScopeActionResultFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceScopeActionResultFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceScopeActionResult(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceScopeActionResult) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetDeviceScopeAction gets the deviceScopeAction property value. Trigger on the service to either START or STOP computing metrics data based on a device scope configuration.
func (m *DeviceScopeActionResult) GetDeviceScopeAction()(*string) {
    return m.deviceScopeAction
}
// GetDeviceScopeId gets the deviceScopeId property value. The unique identifier of the device scope the action was triggered on.
func (m *DeviceScopeActionResult) GetDeviceScopeId()(*string) {
    return m.deviceScopeId
}
// GetFailedMessage gets the failedMessage property value. The message indicates the reason the device scope action failed to trigger.
func (m *DeviceScopeActionResult) GetFailedMessage()(*string) {
    return m.failedMessage
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceScopeActionResult) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["deviceScopeAction"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDeviceScopeAction)
    res["deviceScopeId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDeviceScopeId)
    res["failedMessage"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetFailedMessage)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    res["status"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseDeviceScopeActionStatus , m.SetStatus)
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *DeviceScopeActionResult) GetOdataType()(*string) {
    return m.odataType
}
// GetStatus gets the status property value. Indicates the status of the attempted device scope action
func (m *DeviceScopeActionResult) GetStatus()(*DeviceScopeActionStatus) {
    return m.status
}
// Serialize serializes information the current object
func (m *DeviceScopeActionResult) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("deviceScopeAction", m.GetDeviceScopeAction())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("deviceScopeId", m.GetDeviceScopeId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("failedMessage", m.GetFailedMessage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    if m.GetStatus() != nil {
        cast := (*m.GetStatus()).String()
        err := writer.WriteStringValue("status", &cast)
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
func (m *DeviceScopeActionResult) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetDeviceScopeAction sets the deviceScopeAction property value. Trigger on the service to either START or STOP computing metrics data based on a device scope configuration.
func (m *DeviceScopeActionResult) SetDeviceScopeAction(value *string)() {
    m.deviceScopeAction = value
}
// SetDeviceScopeId sets the deviceScopeId property value. The unique identifier of the device scope the action was triggered on.
func (m *DeviceScopeActionResult) SetDeviceScopeId(value *string)() {
    m.deviceScopeId = value
}
// SetFailedMessage sets the failedMessage property value. The message indicates the reason the device scope action failed to trigger.
func (m *DeviceScopeActionResult) SetFailedMessage(value *string)() {
    m.failedMessage = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *DeviceScopeActionResult) SetOdataType(value *string)() {
    m.odataType = value
}
// SetStatus sets the status property value. Indicates the status of the attempted device scope action
func (m *DeviceScopeActionResult) SetStatus(value *DeviceScopeActionStatus)() {
    m.status = value
}
