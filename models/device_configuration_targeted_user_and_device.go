package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceConfigurationTargetedUserAndDevice conflict summary for a set of device configuration policies.
type DeviceConfigurationTargetedUserAndDevice struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The id of the device in the checkin.
    deviceId *string
    // The name of the device in the checkin.
    deviceName *string
    // Last checkin time for this user/device pair.
    lastCheckinDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The OdataType property
    odataType *string
    // The display name of the user in the checkin
    userDisplayName *string
    // The id of the user in the checkin.
    userId *string
    // The UPN of the user in the checkin.
    userPrincipalName *string
}
// NewDeviceConfigurationTargetedUserAndDevice instantiates a new deviceConfigurationTargetedUserAndDevice and sets the default values.
func NewDeviceConfigurationTargetedUserAndDevice()(*DeviceConfigurationTargetedUserAndDevice) {
    m := &DeviceConfigurationTargetedUserAndDevice{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    odataTypeValue := "#microsoft.graph.deviceConfigurationTargetedUserAndDevice";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateDeviceConfigurationTargetedUserAndDeviceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceConfigurationTargetedUserAndDeviceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceConfigurationTargetedUserAndDevice(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceConfigurationTargetedUserAndDevice) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetDeviceId gets the deviceId property value. The id of the device in the checkin.
func (m *DeviceConfigurationTargetedUserAndDevice) GetDeviceId()(*string) {
    return m.deviceId
}
// GetDeviceName gets the deviceName property value. The name of the device in the checkin.
func (m *DeviceConfigurationTargetedUserAndDevice) GetDeviceName()(*string) {
    return m.deviceName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceConfigurationTargetedUserAndDevice) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["deviceId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDeviceId)
    res["deviceName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDeviceName)
    res["lastCheckinDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetLastCheckinDateTime)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    res["userDisplayName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetUserDisplayName)
    res["userId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetUserId)
    res["userPrincipalName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetUserPrincipalName)
    return res
}
// GetLastCheckinDateTime gets the lastCheckinDateTime property value. Last checkin time for this user/device pair.
func (m *DeviceConfigurationTargetedUserAndDevice) GetLastCheckinDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastCheckinDateTime
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *DeviceConfigurationTargetedUserAndDevice) GetOdataType()(*string) {
    return m.odataType
}
// GetUserDisplayName gets the userDisplayName property value. The display name of the user in the checkin
func (m *DeviceConfigurationTargetedUserAndDevice) GetUserDisplayName()(*string) {
    return m.userDisplayName
}
// GetUserId gets the userId property value. The id of the user in the checkin.
func (m *DeviceConfigurationTargetedUserAndDevice) GetUserId()(*string) {
    return m.userId
}
// GetUserPrincipalName gets the userPrincipalName property value. The UPN of the user in the checkin.
func (m *DeviceConfigurationTargetedUserAndDevice) GetUserPrincipalName()(*string) {
    return m.userPrincipalName
}
// Serialize serializes information the current object
func (m *DeviceConfigurationTargetedUserAndDevice) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("deviceId", m.GetDeviceId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("deviceName", m.GetDeviceName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("lastCheckinDateTime", m.GetLastCheckinDateTime())
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
    {
        err := writer.WriteStringValue("userDisplayName", m.GetUserDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("userId", m.GetUserId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("userPrincipalName", m.GetUserPrincipalName())
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
func (m *DeviceConfigurationTargetedUserAndDevice) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetDeviceId sets the deviceId property value. The id of the device in the checkin.
func (m *DeviceConfigurationTargetedUserAndDevice) SetDeviceId(value *string)() {
    m.deviceId = value
}
// SetDeviceName sets the deviceName property value. The name of the device in the checkin.
func (m *DeviceConfigurationTargetedUserAndDevice) SetDeviceName(value *string)() {
    m.deviceName = value
}
// SetLastCheckinDateTime sets the lastCheckinDateTime property value. Last checkin time for this user/device pair.
func (m *DeviceConfigurationTargetedUserAndDevice) SetLastCheckinDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastCheckinDateTime = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *DeviceConfigurationTargetedUserAndDevice) SetOdataType(value *string)() {
    m.odataType = value
}
// SetUserDisplayName sets the userDisplayName property value. The display name of the user in the checkin
func (m *DeviceConfigurationTargetedUserAndDevice) SetUserDisplayName(value *string)() {
    m.userDisplayName = value
}
// SetUserId sets the userId property value. The id of the user in the checkin.
func (m *DeviceConfigurationTargetedUserAndDevice) SetUserId(value *string)() {
    m.userId = value
}
// SetUserPrincipalName sets the userPrincipalName property value. The UPN of the user in the checkin.
func (m *DeviceConfigurationTargetedUserAndDevice) SetUserPrincipalName(value *string)() {
    m.userPrincipalName = value
}
