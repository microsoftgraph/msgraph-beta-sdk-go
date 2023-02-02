package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceHealthScriptPolicyState contains properties for policy run state of the device health script.
type DeviceHealthScriptPolicyState struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // A list of the assignment filter ids used for health script applicability evaluation
    assignmentFilterIds []string
    // Indicates the type of execution status of the device management script.
    detectionState *RunState
    // The Intune device Id
    deviceId *string
    // Display name of the device
    deviceName *string
    // The next timestamp of when the device health script is expected to execute
    expectedStateUpdateDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Key of the device health script policy state is a concatenation of the MT sideCar policy Id and Intune device Id
    id *string
    // The last timestamp of when the device health script executed
    lastStateUpdateDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The last time that Intune Managment Extension synced with Intune
    lastSyncDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The OdataType property
    odataType *string
    // Value of the OS Version in string
    osVersion *string
    // The MT sideCar policy Id
    policyId *string
    // Display name of the device health script
    policyName *string
    // Error from the detection script after remediation
    postRemediationDetectionScriptError *string
    // Detection script output after remediation
    postRemediationDetectionScriptOutput *string
    // Error from the detection script before remediation
    preRemediationDetectionScriptError *string
    // Output of the detection script before remediation
    preRemediationDetectionScriptOutput *string
    // Error output of the remediation script
    remediationScriptError *string
    // Indicates the type of execution status of the device management script.
    remediationState *RemediationState
    // Name of the user whom ran the device health script
    userName *string
}
// NewDeviceHealthScriptPolicyState instantiates a new deviceHealthScriptPolicyState and sets the default values.
func NewDeviceHealthScriptPolicyState()(*DeviceHealthScriptPolicyState) {
    m := &DeviceHealthScriptPolicyState{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateDeviceHealthScriptPolicyStateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceHealthScriptPolicyStateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceHealthScriptPolicyState(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceHealthScriptPolicyState) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAssignmentFilterIds gets the assignmentFilterIds property value. A list of the assignment filter ids used for health script applicability evaluation
func (m *DeviceHealthScriptPolicyState) GetAssignmentFilterIds()([]string) {
    return m.assignmentFilterIds
}
// GetDetectionState gets the detectionState property value. Indicates the type of execution status of the device management script.
func (m *DeviceHealthScriptPolicyState) GetDetectionState()(*RunState) {
    return m.detectionState
}
// GetDeviceId gets the deviceId property value. The Intune device Id
func (m *DeviceHealthScriptPolicyState) GetDeviceId()(*string) {
    return m.deviceId
}
// GetDeviceName gets the deviceName property value. Display name of the device
func (m *DeviceHealthScriptPolicyState) GetDeviceName()(*string) {
    return m.deviceName
}
// GetExpectedStateUpdateDateTime gets the expectedStateUpdateDateTime property value. The next timestamp of when the device health script is expected to execute
func (m *DeviceHealthScriptPolicyState) GetExpectedStateUpdateDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.expectedStateUpdateDateTime
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
    return m.id
}
// GetLastStateUpdateDateTime gets the lastStateUpdateDateTime property value. The last timestamp of when the device health script executed
func (m *DeviceHealthScriptPolicyState) GetLastStateUpdateDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastStateUpdateDateTime
}
// GetLastSyncDateTime gets the lastSyncDateTime property value. The last time that Intune Managment Extension synced with Intune
func (m *DeviceHealthScriptPolicyState) GetLastSyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastSyncDateTime
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *DeviceHealthScriptPolicyState) GetOdataType()(*string) {
    return m.odataType
}
// GetOsVersion gets the osVersion property value. Value of the OS Version in string
func (m *DeviceHealthScriptPolicyState) GetOsVersion()(*string) {
    return m.osVersion
}
// GetPolicyId gets the policyId property value. The MT sideCar policy Id
func (m *DeviceHealthScriptPolicyState) GetPolicyId()(*string) {
    return m.policyId
}
// GetPolicyName gets the policyName property value. Display name of the device health script
func (m *DeviceHealthScriptPolicyState) GetPolicyName()(*string) {
    return m.policyName
}
// GetPostRemediationDetectionScriptError gets the postRemediationDetectionScriptError property value. Error from the detection script after remediation
func (m *DeviceHealthScriptPolicyState) GetPostRemediationDetectionScriptError()(*string) {
    return m.postRemediationDetectionScriptError
}
// GetPostRemediationDetectionScriptOutput gets the postRemediationDetectionScriptOutput property value. Detection script output after remediation
func (m *DeviceHealthScriptPolicyState) GetPostRemediationDetectionScriptOutput()(*string) {
    return m.postRemediationDetectionScriptOutput
}
// GetPreRemediationDetectionScriptError gets the preRemediationDetectionScriptError property value. Error from the detection script before remediation
func (m *DeviceHealthScriptPolicyState) GetPreRemediationDetectionScriptError()(*string) {
    return m.preRemediationDetectionScriptError
}
// GetPreRemediationDetectionScriptOutput gets the preRemediationDetectionScriptOutput property value. Output of the detection script before remediation
func (m *DeviceHealthScriptPolicyState) GetPreRemediationDetectionScriptOutput()(*string) {
    return m.preRemediationDetectionScriptOutput
}
// GetRemediationScriptError gets the remediationScriptError property value. Error output of the remediation script
func (m *DeviceHealthScriptPolicyState) GetRemediationScriptError()(*string) {
    return m.remediationScriptError
}
// GetRemediationState gets the remediationState property value. Indicates the type of execution status of the device management script.
func (m *DeviceHealthScriptPolicyState) GetRemediationState()(*RemediationState) {
    return m.remediationState
}
// GetUserName gets the userName property value. Name of the user whom ran the device health script
func (m *DeviceHealthScriptPolicyState) GetUserName()(*string) {
    return m.userName
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
    m.additionalData = value
}
// SetAssignmentFilterIds sets the assignmentFilterIds property value. A list of the assignment filter ids used for health script applicability evaluation
func (m *DeviceHealthScriptPolicyState) SetAssignmentFilterIds(value []string)() {
    m.assignmentFilterIds = value
}
// SetDetectionState sets the detectionState property value. Indicates the type of execution status of the device management script.
func (m *DeviceHealthScriptPolicyState) SetDetectionState(value *RunState)() {
    m.detectionState = value
}
// SetDeviceId sets the deviceId property value. The Intune device Id
func (m *DeviceHealthScriptPolicyState) SetDeviceId(value *string)() {
    m.deviceId = value
}
// SetDeviceName sets the deviceName property value. Display name of the device
func (m *DeviceHealthScriptPolicyState) SetDeviceName(value *string)() {
    m.deviceName = value
}
// SetExpectedStateUpdateDateTime sets the expectedStateUpdateDateTime property value. The next timestamp of when the device health script is expected to execute
func (m *DeviceHealthScriptPolicyState) SetExpectedStateUpdateDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.expectedStateUpdateDateTime = value
}
// SetId sets the id property value. Key of the device health script policy state is a concatenation of the MT sideCar policy Id and Intune device Id
func (m *DeviceHealthScriptPolicyState) SetId(value *string)() {
    m.id = value
}
// SetLastStateUpdateDateTime sets the lastStateUpdateDateTime property value. The last timestamp of when the device health script executed
func (m *DeviceHealthScriptPolicyState) SetLastStateUpdateDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastStateUpdateDateTime = value
}
// SetLastSyncDateTime sets the lastSyncDateTime property value. The last time that Intune Managment Extension synced with Intune
func (m *DeviceHealthScriptPolicyState) SetLastSyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastSyncDateTime = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *DeviceHealthScriptPolicyState) SetOdataType(value *string)() {
    m.odataType = value
}
// SetOsVersion sets the osVersion property value. Value of the OS Version in string
func (m *DeviceHealthScriptPolicyState) SetOsVersion(value *string)() {
    m.osVersion = value
}
// SetPolicyId sets the policyId property value. The MT sideCar policy Id
func (m *DeviceHealthScriptPolicyState) SetPolicyId(value *string)() {
    m.policyId = value
}
// SetPolicyName sets the policyName property value. Display name of the device health script
func (m *DeviceHealthScriptPolicyState) SetPolicyName(value *string)() {
    m.policyName = value
}
// SetPostRemediationDetectionScriptError sets the postRemediationDetectionScriptError property value. Error from the detection script after remediation
func (m *DeviceHealthScriptPolicyState) SetPostRemediationDetectionScriptError(value *string)() {
    m.postRemediationDetectionScriptError = value
}
// SetPostRemediationDetectionScriptOutput sets the postRemediationDetectionScriptOutput property value. Detection script output after remediation
func (m *DeviceHealthScriptPolicyState) SetPostRemediationDetectionScriptOutput(value *string)() {
    m.postRemediationDetectionScriptOutput = value
}
// SetPreRemediationDetectionScriptError sets the preRemediationDetectionScriptError property value. Error from the detection script before remediation
func (m *DeviceHealthScriptPolicyState) SetPreRemediationDetectionScriptError(value *string)() {
    m.preRemediationDetectionScriptError = value
}
// SetPreRemediationDetectionScriptOutput sets the preRemediationDetectionScriptOutput property value. Output of the detection script before remediation
func (m *DeviceHealthScriptPolicyState) SetPreRemediationDetectionScriptOutput(value *string)() {
    m.preRemediationDetectionScriptOutput = value
}
// SetRemediationScriptError sets the remediationScriptError property value. Error output of the remediation script
func (m *DeviceHealthScriptPolicyState) SetRemediationScriptError(value *string)() {
    m.remediationScriptError = value
}
// SetRemediationState sets the remediationState property value. Indicates the type of execution status of the device management script.
func (m *DeviceHealthScriptPolicyState) SetRemediationState(value *RemediationState)() {
    m.remediationState = value
}
// SetUserName sets the userName property value. Name of the user whom ran the device health script
func (m *DeviceHealthScriptPolicyState) SetUserName(value *string)() {
    m.userName = value
}
