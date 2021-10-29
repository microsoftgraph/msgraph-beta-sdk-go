package executeaction

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// 
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
}
// Instantiates a new executeActionRequestBody and sets the default values.
func NewExecuteActionRequestBody()(*ExecuteActionRequestBody) {
    m := &ExecuteActionRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the actionName property value. 
func (m *ExecuteActionRequestBody) GetActionName()(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ManagedDeviceRemoteAction) {
    if m == nil {
        return nil
    } else {
        return m.actionName
    }
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ExecuteActionRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the carrierUrl property value. 
func (m *ExecuteActionRequestBody) GetCarrierUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.carrierUrl
    }
}
// Gets the deprovisionReason property value. 
func (m *ExecuteActionRequestBody) GetDeprovisionReason()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deprovisionReason
    }
}
// Gets the deviceIds property value. 
func (m *ExecuteActionRequestBody) GetDeviceIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.deviceIds
    }
}
// Gets the deviceName property value. 
func (m *ExecuteActionRequestBody) GetDeviceName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceName
    }
}
// Gets the keepEnrollmentData property value. 
func (m *ExecuteActionRequestBody) GetKeepEnrollmentData()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.keepEnrollmentData
    }
}
// Gets the keepUserData property value. 
func (m *ExecuteActionRequestBody) GetKeepUserData()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.keepUserData
    }
}
// Gets the notificationBody property value. 
func (m *ExecuteActionRequestBody) GetNotificationBody()(*string) {
    if m == nil {
        return nil
    } else {
        return m.notificationBody
    }
}
// Gets the notificationTitle property value. 
func (m *ExecuteActionRequestBody) GetNotificationTitle()(*string) {
    if m == nil {
        return nil
    } else {
        return m.notificationTitle
    }
}
// Gets the organizationalUnitPath property value. 
func (m *ExecuteActionRequestBody) GetOrganizationalUnitPath()(*string) {
    if m == nil {
        return nil
    } else {
        return m.organizationalUnitPath
    }
}
// The deserialization information for the current model
func (m *ExecuteActionRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["actionName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ParseManagedDeviceRemoteAction)
        if err != nil {
            return err
        }
        cast := val.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ManagedDeviceRemoteAction)
        m.SetActionName(&cast)
        return nil
    }
    res["carrierUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCarrierUrl(val)
        return nil
    }
    res["deprovisionReason"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDeprovisionReason(val)
        return nil
    }
    res["deviceIds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetDeviceIds(res)
        return nil
    }
    res["deviceName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDeviceName(val)
        return nil
    }
    res["keepEnrollmentData"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetKeepEnrollmentData(val)
        return nil
    }
    res["keepUserData"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetKeepUserData(val)
        return nil
    }
    res["notificationBody"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetNotificationBody(val)
        return nil
    }
    res["notificationTitle"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetNotificationTitle(val)
        return nil
    }
    res["organizationalUnitPath"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetOrganizationalUnitPath(val)
        return nil
    }
    return res
}
func (m *ExecuteActionRequestBody) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *ExecuteActionRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetActionName() != nil {
        cast := m.GetActionName().String()
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
    {
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the actionName property value. 
// Parameters:
//  - value : Value to set for the actionName property.
func (m *ExecuteActionRequestBody) SetActionName(value *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ManagedDeviceRemoteAction)() {
    m.actionName = value
}
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *ExecuteActionRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the carrierUrl property value. 
// Parameters:
//  - value : Value to set for the carrierUrl property.
func (m *ExecuteActionRequestBody) SetCarrierUrl(value *string)() {
    m.carrierUrl = value
}
// Sets the deprovisionReason property value. 
// Parameters:
//  - value : Value to set for the deprovisionReason property.
func (m *ExecuteActionRequestBody) SetDeprovisionReason(value *string)() {
    m.deprovisionReason = value
}
// Sets the deviceIds property value. 
// Parameters:
//  - value : Value to set for the deviceIds property.
func (m *ExecuteActionRequestBody) SetDeviceIds(value []string)() {
    m.deviceIds = value
}
// Sets the deviceName property value. 
// Parameters:
//  - value : Value to set for the deviceName property.
func (m *ExecuteActionRequestBody) SetDeviceName(value *string)() {
    m.deviceName = value
}
// Sets the keepEnrollmentData property value. 
// Parameters:
//  - value : Value to set for the keepEnrollmentData property.
func (m *ExecuteActionRequestBody) SetKeepEnrollmentData(value *bool)() {
    m.keepEnrollmentData = value
}
// Sets the keepUserData property value. 
// Parameters:
//  - value : Value to set for the keepUserData property.
func (m *ExecuteActionRequestBody) SetKeepUserData(value *bool)() {
    m.keepUserData = value
}
// Sets the notificationBody property value. 
// Parameters:
//  - value : Value to set for the notificationBody property.
func (m *ExecuteActionRequestBody) SetNotificationBody(value *string)() {
    m.notificationBody = value
}
// Sets the notificationTitle property value. 
// Parameters:
//  - value : Value to set for the notificationTitle property.
func (m *ExecuteActionRequestBody) SetNotificationTitle(value *string)() {
    m.notificationTitle = value
}
// Sets the organizationalUnitPath property value. 
// Parameters:
//  - value : Value to set for the organizationalUnitPath property.
func (m *ExecuteActionRequestBody) SetOrganizationalUnitPath(value *string)() {
    m.organizationalUnitPath = value
}
