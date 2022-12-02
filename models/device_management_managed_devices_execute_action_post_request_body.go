package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementManagedDevicesExecuteActionPostRequestBody provides operations to call the executeAction method.
type DeviceManagementManagedDevicesExecuteActionPostRequestBody struct {
    // The actionName property
    actionName *ManagedDeviceRemoteAction
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The carrierUrl property
    carrierUrl *string
    // The deprovisionReason property
    deprovisionReason *string
    // The deviceIds property
    deviceIds []string
    // The deviceName property
    deviceName *string
    // The keepEnrollmentData property
    keepEnrollmentData *bool
    // The keepUserData property
    keepUserData *bool
    // The notificationBody property
    notificationBody *string
    // The notificationTitle property
    notificationTitle *string
    // The organizationalUnitPath property
    organizationalUnitPath *string
    // The persistEsimDataPlan property
    persistEsimDataPlan *bool
}
// NewDeviceManagementManagedDevicesExecuteActionPostRequestBody instantiates a new DeviceManagementManagedDevicesExecuteActionPostRequestBody and sets the default values.
func NewDeviceManagementManagedDevicesExecuteActionPostRequestBody()(*DeviceManagementManagedDevicesExecuteActionPostRequestBody) {
    m := &DeviceManagementManagedDevicesExecuteActionPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDeviceManagementManagedDevicesExecuteActionPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementManagedDevicesExecuteActionPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementManagedDevicesExecuteActionPostRequestBody(), nil
}
// GetActionName gets the actionName property value. The actionName property
func (m *DeviceManagementManagedDevicesExecuteActionPostRequestBody) GetActionName()(*ManagedDeviceRemoteAction) {
    return m.actionName
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementManagedDevicesExecuteActionPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetCarrierUrl gets the carrierUrl property value. The carrierUrl property
func (m *DeviceManagementManagedDevicesExecuteActionPostRequestBody) GetCarrierUrl()(*string) {
    return m.carrierUrl
}
// GetDeprovisionReason gets the deprovisionReason property value. The deprovisionReason property
func (m *DeviceManagementManagedDevicesExecuteActionPostRequestBody) GetDeprovisionReason()(*string) {
    return m.deprovisionReason
}
// GetDeviceIds gets the deviceIds property value. The deviceIds property
func (m *DeviceManagementManagedDevicesExecuteActionPostRequestBody) GetDeviceIds()([]string) {
    return m.deviceIds
}
// GetDeviceName gets the deviceName property value. The deviceName property
func (m *DeviceManagementManagedDevicesExecuteActionPostRequestBody) GetDeviceName()(*string) {
    return m.deviceName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementManagedDevicesExecuteActionPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["actionName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseManagedDeviceRemoteAction)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActionName(val.(*ManagedDeviceRemoteAction))
        }
        return nil
    }
    res["carrierUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCarrierUrl(val)
        }
        return nil
    }
    res["deprovisionReason"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeprovisionReason(val)
        }
        return nil
    }
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
    res["deviceName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceName(val)
        }
        return nil
    }
    res["keepEnrollmentData"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKeepEnrollmentData(val)
        }
        return nil
    }
    res["keepUserData"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKeepUserData(val)
        }
        return nil
    }
    res["notificationBody"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotificationBody(val)
        }
        return nil
    }
    res["notificationTitle"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotificationTitle(val)
        }
        return nil
    }
    res["organizationalUnitPath"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOrganizationalUnitPath(val)
        }
        return nil
    }
    res["persistEsimDataPlan"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPersistEsimDataPlan(val)
        }
        return nil
    }
    return res
}
// GetKeepEnrollmentData gets the keepEnrollmentData property value. The keepEnrollmentData property
func (m *DeviceManagementManagedDevicesExecuteActionPostRequestBody) GetKeepEnrollmentData()(*bool) {
    return m.keepEnrollmentData
}
// GetKeepUserData gets the keepUserData property value. The keepUserData property
func (m *DeviceManagementManagedDevicesExecuteActionPostRequestBody) GetKeepUserData()(*bool) {
    return m.keepUserData
}
// GetNotificationBody gets the notificationBody property value. The notificationBody property
func (m *DeviceManagementManagedDevicesExecuteActionPostRequestBody) GetNotificationBody()(*string) {
    return m.notificationBody
}
// GetNotificationTitle gets the notificationTitle property value. The notificationTitle property
func (m *DeviceManagementManagedDevicesExecuteActionPostRequestBody) GetNotificationTitle()(*string) {
    return m.notificationTitle
}
// GetOrganizationalUnitPath gets the organizationalUnitPath property value. The organizationalUnitPath property
func (m *DeviceManagementManagedDevicesExecuteActionPostRequestBody) GetOrganizationalUnitPath()(*string) {
    return m.organizationalUnitPath
}
// GetPersistEsimDataPlan gets the persistEsimDataPlan property value. The persistEsimDataPlan property
func (m *DeviceManagementManagedDevicesExecuteActionPostRequestBody) GetPersistEsimDataPlan()(*bool) {
    return m.persistEsimDataPlan
}
// Serialize serializes information the current object
func (m *DeviceManagementManagedDevicesExecuteActionPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetActionName() != nil {
        cast := (*m.GetActionName()).String()
        err := writer.WriteStringValue("actionName", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("carrierUrl", m.GetCarrierUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("deprovisionReason", m.GetDeprovisionReason())
        if err != nil {
            return err
        }
    }
    if m.GetDeviceIds() != nil {
        err := writer.WriteCollectionOfStringValues("deviceIds", m.GetDeviceIds())
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
        err := writer.WriteBoolValue("keepEnrollmentData", m.GetKeepEnrollmentData())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("keepUserData", m.GetKeepUserData())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("notificationBody", m.GetNotificationBody())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("notificationTitle", m.GetNotificationTitle())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("organizationalUnitPath", m.GetOrganizationalUnitPath())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("persistEsimDataPlan", m.GetPersistEsimDataPlan())
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
// SetActionName sets the actionName property value. The actionName property
func (m *DeviceManagementManagedDevicesExecuteActionPostRequestBody) SetActionName(value *ManagedDeviceRemoteAction)() {
    m.actionName = value
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementManagedDevicesExecuteActionPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetCarrierUrl sets the carrierUrl property value. The carrierUrl property
func (m *DeviceManagementManagedDevicesExecuteActionPostRequestBody) SetCarrierUrl(value *string)() {
    m.carrierUrl = value
}
// SetDeprovisionReason sets the deprovisionReason property value. The deprovisionReason property
func (m *DeviceManagementManagedDevicesExecuteActionPostRequestBody) SetDeprovisionReason(value *string)() {
    m.deprovisionReason = value
}
// SetDeviceIds sets the deviceIds property value. The deviceIds property
func (m *DeviceManagementManagedDevicesExecuteActionPostRequestBody) SetDeviceIds(value []string)() {
    m.deviceIds = value
}
// SetDeviceName sets the deviceName property value. The deviceName property
func (m *DeviceManagementManagedDevicesExecuteActionPostRequestBody) SetDeviceName(value *string)() {
    m.deviceName = value
}
// SetKeepEnrollmentData sets the keepEnrollmentData property value. The keepEnrollmentData property
func (m *DeviceManagementManagedDevicesExecuteActionPostRequestBody) SetKeepEnrollmentData(value *bool)() {
    m.keepEnrollmentData = value
}
// SetKeepUserData sets the keepUserData property value. The keepUserData property
func (m *DeviceManagementManagedDevicesExecuteActionPostRequestBody) SetKeepUserData(value *bool)() {
    m.keepUserData = value
}
// SetNotificationBody sets the notificationBody property value. The notificationBody property
func (m *DeviceManagementManagedDevicesExecuteActionPostRequestBody) SetNotificationBody(value *string)() {
    m.notificationBody = value
}
// SetNotificationTitle sets the notificationTitle property value. The notificationTitle property
func (m *DeviceManagementManagedDevicesExecuteActionPostRequestBody) SetNotificationTitle(value *string)() {
    m.notificationTitle = value
}
// SetOrganizationalUnitPath sets the organizationalUnitPath property value. The organizationalUnitPath property
func (m *DeviceManagementManagedDevicesExecuteActionPostRequestBody) SetOrganizationalUnitPath(value *string)() {
    m.organizationalUnitPath = value
}
// SetPersistEsimDataPlan sets the persistEsimDataPlan property value. The persistEsimDataPlan property
func (m *DeviceManagementManagedDevicesExecuteActionPostRequestBody) SetPersistEsimDataPlan(value *bool)() {
    m.persistEsimDataPlan = value
}
