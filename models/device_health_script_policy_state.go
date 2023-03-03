package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// DeviceHealthScriptPolicyState contains properties for policy run state of the device health script.
type DeviceHealthScriptPolicyState struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewDeviceHealthScriptPolicyState instantiates a new deviceHealthScriptPolicyState and sets the default values.
func NewDeviceHealthScriptPolicyState()(*DeviceHealthScriptPolicyState) {
    m := &DeviceHealthScriptPolicyState{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateDeviceHealthScriptPolicyStateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceHealthScriptPolicyStateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceHealthScriptPolicyState(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceHealthScriptPolicyState) GetAdditionalData()(map[string]any) {
    val , err :=  m.backingStore.Get("additionalData")
    if err != nil {
        panic(err)
    }
    if val == nil {
        var value = make(map[string]any);
        m.SetAdditionalData(value);
    }
    return val.(map[string]any)
}
// GetAssignmentFilterIds gets the assignmentFilterIds property value. A list of the assignment filter ids used for health script applicability evaluation
func (m *DeviceHealthScriptPolicyState) GetAssignmentFilterIds()([]string) {
    val, err := m.GetBackingStore().Get("assignmentFilterIds")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetBackingStore gets the backingStore property value. Stores model information.
func (m *DeviceHealthScriptPolicyState) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetDetectionState gets the detectionState property value. Indicates the type of execution status of the device management script.
func (m *DeviceHealthScriptPolicyState) GetDetectionState()(*RunState) {
    val, err := m.GetBackingStore().Get("detectionState")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*RunState)
    }
    return nil
}
// GetDeviceId gets the deviceId property value. The Intune device Id
func (m *DeviceHealthScriptPolicyState) GetDeviceId()(*string) {
    val, err := m.GetBackingStore().Get("deviceId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDeviceName gets the deviceName property value. Display name of the device
func (m *DeviceHealthScriptPolicyState) GetDeviceName()(*string) {
    val, err := m.GetBackingStore().Get("deviceName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetExpectedStateUpdateDateTime gets the expectedStateUpdateDateTime property value. The next timestamp of when the device health script is expected to execute
func (m *DeviceHealthScriptPolicyState) GetExpectedStateUpdateDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("expectedStateUpdateDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceHealthScriptPolicyState) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["assignmentFilterIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetAssignmentFilterIds(res)
        }
        return nil
    }
    res["detectionState"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseRunState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDetectionState(val.(*RunState))
        }
        return nil
    }
    res["deviceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceId(val)
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
    res["expectedStateUpdateDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExpectedStateUpdateDateTime(val)
        }
        return nil
    }
    res["id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
        }
        return nil
    }
    res["lastStateUpdateDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastStateUpdateDateTime(val)
        }
        return nil
    }
    res["lastSyncDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastSyncDateTime(val)
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    res["osVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOsVersion(val)
        }
        return nil
    }
    res["policyId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPolicyId(val)
        }
        return nil
    }
    res["policyName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPolicyName(val)
        }
        return nil
    }
    res["postRemediationDetectionScriptError"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPostRemediationDetectionScriptError(val)
        }
        return nil
    }
    res["postRemediationDetectionScriptOutput"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPostRemediationDetectionScriptOutput(val)
        }
        return nil
    }
    res["preRemediationDetectionScriptError"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPreRemediationDetectionScriptError(val)
        }
        return nil
    }
    res["preRemediationDetectionScriptOutput"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPreRemediationDetectionScriptOutput(val)
        }
        return nil
    }
    res["remediationScriptError"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRemediationScriptError(val)
        }
        return nil
    }
    res["remediationState"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseRemediationState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRemediationState(val.(*RemediationState))
        }
        return nil
    }
    res["userName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserName(val)
        }
        return nil
    }
    return res
}
// GetId gets the id property value. Key of the device health script policy state is a concatenation of the MT sideCar policy Id and Intune device Id
func (m *DeviceHealthScriptPolicyState) GetId()(*string) {
    val, err := m.GetBackingStore().Get("id")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetLastStateUpdateDateTime gets the lastStateUpdateDateTime property value. The last timestamp of when the device health script executed
func (m *DeviceHealthScriptPolicyState) GetLastStateUpdateDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("lastStateUpdateDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetLastSyncDateTime gets the lastSyncDateTime property value. The last time that Intune Managment Extension synced with Intune
func (m *DeviceHealthScriptPolicyState) GetLastSyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("lastSyncDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *DeviceHealthScriptPolicyState) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetOsVersion gets the osVersion property value. Value of the OS Version in string
func (m *DeviceHealthScriptPolicyState) GetOsVersion()(*string) {
    val, err := m.GetBackingStore().Get("osVersion")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPolicyId gets the policyId property value. The MT sideCar policy Id
func (m *DeviceHealthScriptPolicyState) GetPolicyId()(*string) {
    val, err := m.GetBackingStore().Get("policyId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPolicyName gets the policyName property value. Display name of the device health script
func (m *DeviceHealthScriptPolicyState) GetPolicyName()(*string) {
    val, err := m.GetBackingStore().Get("policyName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPostRemediationDetectionScriptError gets the postRemediationDetectionScriptError property value. Error from the detection script after remediation
func (m *DeviceHealthScriptPolicyState) GetPostRemediationDetectionScriptError()(*string) {
    val, err := m.GetBackingStore().Get("postRemediationDetectionScriptError")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPostRemediationDetectionScriptOutput gets the postRemediationDetectionScriptOutput property value. Detection script output after remediation
func (m *DeviceHealthScriptPolicyState) GetPostRemediationDetectionScriptOutput()(*string) {
    val, err := m.GetBackingStore().Get("postRemediationDetectionScriptOutput")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPreRemediationDetectionScriptError gets the preRemediationDetectionScriptError property value. Error from the detection script before remediation
func (m *DeviceHealthScriptPolicyState) GetPreRemediationDetectionScriptError()(*string) {
    val, err := m.GetBackingStore().Get("preRemediationDetectionScriptError")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPreRemediationDetectionScriptOutput gets the preRemediationDetectionScriptOutput property value. Output of the detection script before remediation
func (m *DeviceHealthScriptPolicyState) GetPreRemediationDetectionScriptOutput()(*string) {
    val, err := m.GetBackingStore().Get("preRemediationDetectionScriptOutput")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetRemediationScriptError gets the remediationScriptError property value. Error output of the remediation script
func (m *DeviceHealthScriptPolicyState) GetRemediationScriptError()(*string) {
    val, err := m.GetBackingStore().Get("remediationScriptError")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetRemediationState gets the remediationState property value. Indicates the type of execution status of the device management script.
func (m *DeviceHealthScriptPolicyState) GetRemediationState()(*RemediationState) {
    val, err := m.GetBackingStore().Get("remediationState")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*RemediationState)
    }
    return nil
}
// GetUserName gets the userName property value. Name of the user whom ran the device health script
func (m *DeviceHealthScriptPolicyState) GetUserName()(*string) {
    val, err := m.GetBackingStore().Get("userName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *DeviceHealthScriptPolicyState) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAssignmentFilterIds() != nil {
        err := writer.WriteCollectionOfStringValues("assignmentFilterIds", m.GetAssignmentFilterIds())
        if err != nil {
            return err
        }
    }
    if m.GetDetectionState() != nil {
        cast := (*m.GetDetectionState()).String()
        err := writer.WriteStringValue("detectionState", &cast)
        if err != nil {
            return err
        }
    }
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
        err := writer.WriteTimeValue("expectedStateUpdateDateTime", m.GetExpectedStateUpdateDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("lastStateUpdateDateTime", m.GetLastStateUpdateDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("lastSyncDateTime", m.GetLastSyncDateTime())
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
        err := writer.WriteStringValue("osVersion", m.GetOsVersion())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("policyId", m.GetPolicyId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("policyName", m.GetPolicyName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("postRemediationDetectionScriptError", m.GetPostRemediationDetectionScriptError())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("postRemediationDetectionScriptOutput", m.GetPostRemediationDetectionScriptOutput())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("preRemediationDetectionScriptError", m.GetPreRemediationDetectionScriptError())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("preRemediationDetectionScriptOutput", m.GetPreRemediationDetectionScriptOutput())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("remediationScriptError", m.GetRemediationScriptError())
        if err != nil {
            return err
        }
    }
    if m.GetRemediationState() != nil {
        cast := (*m.GetRemediationState()).String()
        err := writer.WriteStringValue("remediationState", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("userName", m.GetUserName())
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
func (m *DeviceHealthScriptPolicyState) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetAssignmentFilterIds sets the assignmentFilterIds property value. A list of the assignment filter ids used for health script applicability evaluation
func (m *DeviceHealthScriptPolicyState) SetAssignmentFilterIds(value []string)() {
    err := m.GetBackingStore().Set("assignmentFilterIds", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the backingStore property value. Stores model information.
func (m *DeviceHealthScriptPolicyState) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetDetectionState sets the detectionState property value. Indicates the type of execution status of the device management script.
func (m *DeviceHealthScriptPolicyState) SetDetectionState(value *RunState)() {
    err := m.GetBackingStore().Set("detectionState", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceId sets the deviceId property value. The Intune device Id
func (m *DeviceHealthScriptPolicyState) SetDeviceId(value *string)() {
    err := m.GetBackingStore().Set("deviceId", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceName sets the deviceName property value. Display name of the device
func (m *DeviceHealthScriptPolicyState) SetDeviceName(value *string)() {
    err := m.GetBackingStore().Set("deviceName", value)
    if err != nil {
        panic(err)
    }
}
// SetExpectedStateUpdateDateTime sets the expectedStateUpdateDateTime property value. The next timestamp of when the device health script is expected to execute
func (m *DeviceHealthScriptPolicyState) SetExpectedStateUpdateDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("expectedStateUpdateDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetId sets the id property value. Key of the device health script policy state is a concatenation of the MT sideCar policy Id and Intune device Id
func (m *DeviceHealthScriptPolicyState) SetId(value *string)() {
    err := m.GetBackingStore().Set("id", value)
    if err != nil {
        panic(err)
    }
}
// SetLastStateUpdateDateTime sets the lastStateUpdateDateTime property value. The last timestamp of when the device health script executed
func (m *DeviceHealthScriptPolicyState) SetLastStateUpdateDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("lastStateUpdateDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetLastSyncDateTime sets the lastSyncDateTime property value. The last time that Intune Managment Extension synced with Intune
func (m *DeviceHealthScriptPolicyState) SetLastSyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("lastSyncDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *DeviceHealthScriptPolicyState) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetOsVersion sets the osVersion property value. Value of the OS Version in string
func (m *DeviceHealthScriptPolicyState) SetOsVersion(value *string)() {
    err := m.GetBackingStore().Set("osVersion", value)
    if err != nil {
        panic(err)
    }
}
// SetPolicyId sets the policyId property value. The MT sideCar policy Id
func (m *DeviceHealthScriptPolicyState) SetPolicyId(value *string)() {
    err := m.GetBackingStore().Set("policyId", value)
    if err != nil {
        panic(err)
    }
}
// SetPolicyName sets the policyName property value. Display name of the device health script
func (m *DeviceHealthScriptPolicyState) SetPolicyName(value *string)() {
    err := m.GetBackingStore().Set("policyName", value)
    if err != nil {
        panic(err)
    }
}
// SetPostRemediationDetectionScriptError sets the postRemediationDetectionScriptError property value. Error from the detection script after remediation
func (m *DeviceHealthScriptPolicyState) SetPostRemediationDetectionScriptError(value *string)() {
    err := m.GetBackingStore().Set("postRemediationDetectionScriptError", value)
    if err != nil {
        panic(err)
    }
}
// SetPostRemediationDetectionScriptOutput sets the postRemediationDetectionScriptOutput property value. Detection script output after remediation
func (m *DeviceHealthScriptPolicyState) SetPostRemediationDetectionScriptOutput(value *string)() {
    err := m.GetBackingStore().Set("postRemediationDetectionScriptOutput", value)
    if err != nil {
        panic(err)
    }
}
// SetPreRemediationDetectionScriptError sets the preRemediationDetectionScriptError property value. Error from the detection script before remediation
func (m *DeviceHealthScriptPolicyState) SetPreRemediationDetectionScriptError(value *string)() {
    err := m.GetBackingStore().Set("preRemediationDetectionScriptError", value)
    if err != nil {
        panic(err)
    }
}
// SetPreRemediationDetectionScriptOutput sets the preRemediationDetectionScriptOutput property value. Output of the detection script before remediation
func (m *DeviceHealthScriptPolicyState) SetPreRemediationDetectionScriptOutput(value *string)() {
    err := m.GetBackingStore().Set("preRemediationDetectionScriptOutput", value)
    if err != nil {
        panic(err)
    }
}
// SetRemediationScriptError sets the remediationScriptError property value. Error output of the remediation script
func (m *DeviceHealthScriptPolicyState) SetRemediationScriptError(value *string)() {
    err := m.GetBackingStore().Set("remediationScriptError", value)
    if err != nil {
        panic(err)
    }
}
// SetRemediationState sets the remediationState property value. Indicates the type of execution status of the device management script.
func (m *DeviceHealthScriptPolicyState) SetRemediationState(value *RemediationState)() {
    err := m.GetBackingStore().Set("remediationState", value)
    if err != nil {
        panic(err)
    }
}
// SetUserName sets the userName property value. Name of the user whom ran the device health script
func (m *DeviceHealthScriptPolicyState) SetUserName(value *string)() {
    err := m.GetBackingStore().Set("userName", value)
    if err != nil {
        panic(err)
    }
}
// DeviceHealthScriptPolicyStateable 
type DeviceHealthScriptPolicyStateable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAssignmentFilterIds()([]string)
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetDetectionState()(*RunState)
    GetDeviceId()(*string)
    GetDeviceName()(*string)
    GetExpectedStateUpdateDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetId()(*string)
    GetLastStateUpdateDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetLastSyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetOdataType()(*string)
    GetOsVersion()(*string)
    GetPolicyId()(*string)
    GetPolicyName()(*string)
    GetPostRemediationDetectionScriptError()(*string)
    GetPostRemediationDetectionScriptOutput()(*string)
    GetPreRemediationDetectionScriptError()(*string)
    GetPreRemediationDetectionScriptOutput()(*string)
    GetRemediationScriptError()(*string)
    GetRemediationState()(*RemediationState)
    GetUserName()(*string)
    SetAssignmentFilterIds(value []string)()
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetDetectionState(value *RunState)()
    SetDeviceId(value *string)()
    SetDeviceName(value *string)()
    SetExpectedStateUpdateDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetId(value *string)()
    SetLastStateUpdateDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetLastSyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetOdataType(value *string)()
    SetOsVersion(value *string)()
    SetPolicyId(value *string)()
    SetPolicyName(value *string)()
    SetPostRemediationDetectionScriptError(value *string)()
    SetPostRemediationDetectionScriptOutput(value *string)()
    SetPreRemediationDetectionScriptError(value *string)()
    SetPreRemediationDetectionScriptOutput(value *string)()
    SetRemediationScriptError(value *string)()
    SetRemediationState(value *RemediationState)()
    SetUserName(value *string)()
}
