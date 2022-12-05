package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MeManagedDevicesExecuteActionPostRequestBody provides operations to call the executeAction method.
type MeManagedDevicesExecuteActionPostRequestBody struct {
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
// NewMeManagedDevicesExecuteActionPostRequestBody instantiates a new MeManagedDevicesExecuteActionPostRequestBody and sets the default values.
func NewMeManagedDevicesExecuteActionPostRequestBody()(*MeManagedDevicesExecuteActionPostRequestBody) {
    m := &MeManagedDevicesExecuteActionPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateMeManagedDevicesExecuteActionPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMeManagedDevicesExecuteActionPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMeManagedDevicesExecuteActionPostRequestBody(), nil
}
// GetActionName gets the actionName property value. The actionName property
func (m *MeManagedDevicesExecuteActionPostRequestBody) GetActionName()(*ManagedDeviceRemoteAction) {
    return m.actionName
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MeManagedDevicesExecuteActionPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetCarrierUrl gets the carrierUrl property value. The carrierUrl property
func (m *MeManagedDevicesExecuteActionPostRequestBody) GetCarrierUrl()(*string) {
    return m.carrierUrl
}
// GetDeprovisionReason gets the deprovisionReason property value. The deprovisionReason property
func (m *MeManagedDevicesExecuteActionPostRequestBody) GetDeprovisionReason()(*string) {
    return m.deprovisionReason
}
// GetDeviceIds gets the deviceIds property value. The deviceIds property
func (m *MeManagedDevicesExecuteActionPostRequestBody) GetDeviceIds()([]string) {
    return m.deviceIds
}
// GetDeviceName gets the deviceName property value. The deviceName property
func (m *MeManagedDevicesExecuteActionPostRequestBody) GetDeviceName()(*string) {
    return m.deviceName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MeManagedDevicesExecuteActionPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
func (m *MeManagedDevicesExecuteActionPostRequestBody) GetKeepEnrollmentData()(*bool) {
    return m.keepEnrollmentData
}
// GetKeepUserData gets the keepUserData property value. The keepUserData property
func (m *MeManagedDevicesExecuteActionPostRequestBody) GetKeepUserData()(*bool) {
    return m.keepUserData
}
// GetNotificationBody gets the notificationBody property value. The notificationBody property
func (m *MeManagedDevicesExecuteActionPostRequestBody) GetNotificationBody()(*string) {
    return m.notificationBody
}
// GetNotificationTitle gets the notificationTitle property value. The notificationTitle property
func (m *MeManagedDevicesExecuteActionPostRequestBody) GetNotificationTitle()(*string) {
    return m.notificationTitle
}
// GetOrganizationalUnitPath gets the organizationalUnitPath property value. The organizationalUnitPath property
func (m *MeManagedDevicesExecuteActionPostRequestBody) GetOrganizationalUnitPath()(*string) {
    return m.organizationalUnitPath
}
// GetPersistEsimDataPlan gets the persistEsimDataPlan property value. The persistEsimDataPlan property
func (m *MeManagedDevicesExecuteActionPostRequestBody) GetPersistEsimDataPlan()(*bool) {
    return m.persistEsimDataPlan
}
// Serialize serializes information the current object
func (m *MeManagedDevicesExecuteActionPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *MeManagedDevicesExecuteActionPostRequestBody) SetActionName(value *ManagedDeviceRemoteAction)() {
    m.actionName = value
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MeManagedDevicesExecuteActionPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetCarrierUrl sets the carrierUrl property value. The carrierUrl property
func (m *MeManagedDevicesExecuteActionPostRequestBody) SetCarrierUrl(value *string)() {
    m.carrierUrl = value
}
// SetDeprovisionReason sets the deprovisionReason property value. The deprovisionReason property
func (m *MeManagedDevicesExecuteActionPostRequestBody) SetDeprovisionReason(value *string)() {
    m.deprovisionReason = value
}
// SetDeviceIds sets the deviceIds property value. The deviceIds property
func (m *MeManagedDevicesExecuteActionPostRequestBody) SetDeviceIds(value []string)() {
    m.deviceIds = value
}
// SetDeviceName sets the deviceName property value. The deviceName property
func (m *MeManagedDevicesExecuteActionPostRequestBody) SetDeviceName(value *string)() {
    m.deviceName = value
}
// SetKeepEnrollmentData sets the keepEnrollmentData property value. The keepEnrollmentData property
func (m *MeManagedDevicesExecuteActionPostRequestBody) SetKeepEnrollmentData(value *bool)() {
    m.keepEnrollmentData = value
}
// SetKeepUserData sets the keepUserData property value. The keepUserData property
func (m *MeManagedDevicesExecuteActionPostRequestBody) SetKeepUserData(value *bool)() {
    m.keepUserData = value
}
// SetNotificationBody sets the notificationBody property value. The notificationBody property
func (m *MeManagedDevicesExecuteActionPostRequestBody) SetNotificationBody(value *string)() {
    m.notificationBody = value
}
// SetNotificationTitle sets the notificationTitle property value. The notificationTitle property
func (m *MeManagedDevicesExecuteActionPostRequestBody) SetNotificationTitle(value *string)() {
    m.notificationTitle = value
}
// SetOrganizationalUnitPath sets the organizationalUnitPath property value. The organizationalUnitPath property
func (m *MeManagedDevicesExecuteActionPostRequestBody) SetOrganizationalUnitPath(value *string)() {
    m.organizationalUnitPath = value
}
// SetPersistEsimDataPlan sets the persistEsimDataPlan property value. The persistEsimDataPlan property
func (m *MeManagedDevicesExecuteActionPostRequestBody) SetPersistEsimDataPlan(value *bool)() {
    m.persistEsimDataPlan = value
}
