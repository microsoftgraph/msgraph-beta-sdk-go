package executeaction

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// ExecuteActionRequestBody 
type ExecuteActionRequestBody struct {
    // 
    actionName *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ManagedDeviceRemoteAction;
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    carrierUrl *string;
    // 
    deprovisionReason *string;
    // 
    deviceIds []string;
    // 
    deviceName *string;
    // 
    keepEnrollmentData *bool;
    // 
    keepUserData *bool;
    // 
    notificationBody *string;
    // 
    notificationTitle *string;
    // 
    organizationalUnitPath *string;
    // 
    persistEsimDataPlan *bool;
}
// NewExecuteActionRequestBody instantiates a new executeActionRequestBody and sets the default values.
func NewExecuteActionRequestBody()(*ExecuteActionRequestBody) {
    m := &ExecuteActionRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetActionName gets the actionName property value. 
func (m *ExecuteActionRequestBody) GetActionName()(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ManagedDeviceRemoteAction) {
    if m == nil {
        return nil
    } else {
        return m.actionName
    }
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ExecuteActionRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetCarrierUrl gets the carrierUrl property value. 
func (m *ExecuteActionRequestBody) GetCarrierUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.carrierUrl
    }
}
// GetDeprovisionReason gets the deprovisionReason property value. 
func (m *ExecuteActionRequestBody) GetDeprovisionReason()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deprovisionReason
    }
}
// GetDeviceIds gets the deviceIds property value. 
func (m *ExecuteActionRequestBody) GetDeviceIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.deviceIds
    }
}
// GetDeviceName gets the deviceName property value. 
func (m *ExecuteActionRequestBody) GetDeviceName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceName
    }
}
// GetKeepEnrollmentData gets the keepEnrollmentData property value. 
func (m *ExecuteActionRequestBody) GetKeepEnrollmentData()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.keepEnrollmentData
    }
}
// GetKeepUserData gets the keepUserData property value. 
func (m *ExecuteActionRequestBody) GetKeepUserData()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.keepUserData
    }
}
// GetNotificationBody gets the notificationBody property value. 
func (m *ExecuteActionRequestBody) GetNotificationBody()(*string) {
    if m == nil {
        return nil
    } else {
        return m.notificationBody
    }
}
// GetNotificationTitle gets the notificationTitle property value. 
func (m *ExecuteActionRequestBody) GetNotificationTitle()(*string) {
    if m == nil {
        return nil
    } else {
        return m.notificationTitle
    }
}
// GetOrganizationalUnitPath gets the organizationalUnitPath property value. 
func (m *ExecuteActionRequestBody) GetOrganizationalUnitPath()(*string) {
    if m == nil {
        return nil
    } else {
        return m.organizationalUnitPath
    }
}
// GetPersistEsimDataPlan gets the persistEsimDataPlan property value. 
func (m *ExecuteActionRequestBody) GetPersistEsimDataPlan()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.persistEsimDataPlan
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ExecuteActionRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["actionName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ParseManagedDeviceRemoteAction)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActionName(val.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ManagedDeviceRemoteAction))
        }
        return nil
    }
    res["carrierUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCarrierUrl(val)
        }
        return nil
    }
    res["deprovisionReason"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeprovisionReason(val)
        }
        return nil
    }
    res["deviceIds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
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
    res["deviceName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceName(val)
        }
        return nil
    }
    res["keepEnrollmentData"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKeepEnrollmentData(val)
        }
        return nil
    }
    res["keepUserData"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKeepUserData(val)
        }
        return nil
    }
    res["notificationBody"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotificationBody(val)
        }
        return nil
    }
    res["notificationTitle"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotificationTitle(val)
        }
        return nil
    }
    res["organizationalUnitPath"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOrganizationalUnitPath(val)
        }
        return nil
    }
    res["persistEsimDataPlan"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
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
func (m *ExecuteActionRequestBody) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ExecuteActionRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
// SetActionName sets the actionName property value. 
func (m *ExecuteActionRequestBody) SetActionName(value *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ManagedDeviceRemoteAction)() {
    if m != nil {
        m.actionName = value
    }
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ExecuteActionRequestBody) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetCarrierUrl sets the carrierUrl property value. 
func (m *ExecuteActionRequestBody) SetCarrierUrl(value *string)() {
    if m != nil {
        m.carrierUrl = value
    }
}
// SetDeprovisionReason sets the deprovisionReason property value. 
func (m *ExecuteActionRequestBody) SetDeprovisionReason(value *string)() {
    if m != nil {
        m.deprovisionReason = value
    }
}
// SetDeviceIds sets the deviceIds property value. 
func (m *ExecuteActionRequestBody) SetDeviceIds(value []string)() {
    if m != nil {
        m.deviceIds = value
    }
}
// SetDeviceName sets the deviceName property value. 
func (m *ExecuteActionRequestBody) SetDeviceName(value *string)() {
    if m != nil {
        m.deviceName = value
    }
}
// SetKeepEnrollmentData sets the keepEnrollmentData property value. 
func (m *ExecuteActionRequestBody) SetKeepEnrollmentData(value *bool)() {
    if m != nil {
        m.keepEnrollmentData = value
    }
}
// SetKeepUserData sets the keepUserData property value. 
func (m *ExecuteActionRequestBody) SetKeepUserData(value *bool)() {
    if m != nil {
        m.keepUserData = value
    }
}
// SetNotificationBody sets the notificationBody property value. 
func (m *ExecuteActionRequestBody) SetNotificationBody(value *string)() {
    if m != nil {
        m.notificationBody = value
    }
}
// SetNotificationTitle sets the notificationTitle property value. 
func (m *ExecuteActionRequestBody) SetNotificationTitle(value *string)() {
    if m != nil {
        m.notificationTitle = value
    }
}
// SetOrganizationalUnitPath sets the organizationalUnitPath property value. 
func (m *ExecuteActionRequestBody) SetOrganizationalUnitPath(value *string)() {
    if m != nil {
        m.organizationalUnitPath = value
    }
}
// SetPersistEsimDataPlan sets the persistEsimDataPlan property value. 
func (m *ExecuteActionRequestBody) SetPersistEsimDataPlan(value *bool)() {
    if m != nil {
        m.persistEsimDataPlan = value
    }
}
