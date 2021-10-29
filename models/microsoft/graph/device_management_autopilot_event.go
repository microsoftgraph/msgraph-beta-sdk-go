package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type DeviceManagementAutopilotEvent struct {
    Entity
    // Time spent in user ESP.
    accountSetupDuration *string;
    // Deployment status for the enrollment status page account setup phase. Possible values are: unknown, success, inProgress, failure, successWithTimeout, notAttempted, disabled.
    accountSetupStatus *WindowsAutopilotDeploymentState;
    // Autopilot deployment duration including enrollment.
    deploymentDuration *string;
    // Deployment end time.
    deploymentEndDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Deployment start time.
    deploymentStartDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Deployment state like Success, Failure, InProgress, SuccessWithTimeout. Possible values are: unknown, success, inProgress, failure, successWithTimeout, notAttempted, disabled.
    deploymentState *WindowsAutopilotDeploymentState;
    // Total deployment duration from enrollment to Desktop screen.
    deploymentTotalDuration *string;
    // Device id associated with the object
    deviceId *string;
    // Time spent in device enrollment.
    devicePreparationDuration *string;
    // Device registration date.
    deviceRegisteredDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Device serial number.
    deviceSerialNumber *string;
    // Time spent in device ESP.
    deviceSetupDuration *string;
    // Deployment status for the enrollment status page device setup phase. Possible values are: unknown, success, inProgress, failure, successWithTimeout, notAttempted, disabled.
    deviceSetupStatus *WindowsAutopilotDeploymentState;
    // Enrollment failure details.
    enrollmentFailureDetails *string;
    // Device enrollment start date.
    enrollmentStartDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Enrollment state like Enrolled, Failed. Possible values are: unknown, enrolled, pendingReset, failed, notContacted, blocked.
    enrollmentState *EnrollmentState;
    // Enrollment type. Possible values are: unknown, azureADJoinedWithAutopilotProfile, offlineDomainJoined, azureADJoinedUsingDeviceAuthWithAutopilotProfile, azureADJoinedUsingDeviceAuthWithoutAutopilotProfile, azureADJoinedWithOfflineAutopilotProfile, azureADJoinedWithWhiteGlove, offlineDomainJoinedWithWhiteGlove, offlineDomainJoinedWithOfflineAutopilotProfile.
    enrollmentType *WindowsAutopilotEnrollmentType;
    // Time when the event occurred .
    eventDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Managed device name.
    managedDeviceName *string;
    // Device operating system version.
    osVersion *string;
    // Policy and application status details for this device.
    policyStatusDetails []DeviceManagementAutopilotPolicyStatusDetail;
    // Count of applications targeted.
    targetedAppCount *int32;
    // Count of policies targeted.
    targetedPolicyCount *int32;
    // User principal name used to enroll the device.
    userPrincipalName *string;
    // Enrollment Status Page profile name
    windows10EnrollmentCompletionPageConfigurationDisplayName *string;
    // Enrollment Status Page profile ID
    windows10EnrollmentCompletionPageConfigurationId *string;
    // Autopilot profile name.
    windowsAutopilotDeploymentProfileDisplayName *string;
}
// Instantiates a new deviceManagementAutopilotEvent and sets the default values.
func NewDeviceManagementAutopilotEvent()(*DeviceManagementAutopilotEvent) {
    m := &DeviceManagementAutopilotEvent{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the accountSetupDuration property value. Time spent in user ESP.
func (m *DeviceManagementAutopilotEvent) GetAccountSetupDuration()(*string) {
    if m == nil {
        return nil
    } else {
        return m.accountSetupDuration
    }
}
// Gets the accountSetupStatus property value. Deployment status for the enrollment status page account setup phase. Possible values are: unknown, success, inProgress, failure, successWithTimeout, notAttempted, disabled.
func (m *DeviceManagementAutopilotEvent) GetAccountSetupStatus()(*WindowsAutopilotDeploymentState) {
    if m == nil {
        return nil
    } else {
        return m.accountSetupStatus
    }
}
// Gets the deploymentDuration property value. Autopilot deployment duration including enrollment.
func (m *DeviceManagementAutopilotEvent) GetDeploymentDuration()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deploymentDuration
    }
}
// Gets the deploymentEndDateTime property value. Deployment end time.
func (m *DeviceManagementAutopilotEvent) GetDeploymentEndDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.deploymentEndDateTime
    }
}
// Gets the deploymentStartDateTime property value. Deployment start time.
func (m *DeviceManagementAutopilotEvent) GetDeploymentStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.deploymentStartDateTime
    }
}
// Gets the deploymentState property value. Deployment state like Success, Failure, InProgress, SuccessWithTimeout. Possible values are: unknown, success, inProgress, failure, successWithTimeout, notAttempted, disabled.
func (m *DeviceManagementAutopilotEvent) GetDeploymentState()(*WindowsAutopilotDeploymentState) {
    if m == nil {
        return nil
    } else {
        return m.deploymentState
    }
}
// Gets the deploymentTotalDuration property value. Total deployment duration from enrollment to Desktop screen.
func (m *DeviceManagementAutopilotEvent) GetDeploymentTotalDuration()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deploymentTotalDuration
    }
}
// Gets the deviceId property value. Device id associated with the object
func (m *DeviceManagementAutopilotEvent) GetDeviceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceId
    }
}
// Gets the devicePreparationDuration property value. Time spent in device enrollment.
func (m *DeviceManagementAutopilotEvent) GetDevicePreparationDuration()(*string) {
    if m == nil {
        return nil
    } else {
        return m.devicePreparationDuration
    }
}
// Gets the deviceRegisteredDateTime property value. Device registration date.
func (m *DeviceManagementAutopilotEvent) GetDeviceRegisteredDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.deviceRegisteredDateTime
    }
}
// Gets the deviceSerialNumber property value. Device serial number.
func (m *DeviceManagementAutopilotEvent) GetDeviceSerialNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceSerialNumber
    }
}
// Gets the deviceSetupDuration property value. Time spent in device ESP.
func (m *DeviceManagementAutopilotEvent) GetDeviceSetupDuration()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceSetupDuration
    }
}
// Gets the deviceSetupStatus property value. Deployment status for the enrollment status page device setup phase. Possible values are: unknown, success, inProgress, failure, successWithTimeout, notAttempted, disabled.
func (m *DeviceManagementAutopilotEvent) GetDeviceSetupStatus()(*WindowsAutopilotDeploymentState) {
    if m == nil {
        return nil
    } else {
        return m.deviceSetupStatus
    }
}
// Gets the enrollmentFailureDetails property value. Enrollment failure details.
func (m *DeviceManagementAutopilotEvent) GetEnrollmentFailureDetails()(*string) {
    if m == nil {
        return nil
    } else {
        return m.enrollmentFailureDetails
    }
}
// Gets the enrollmentStartDateTime property value. Device enrollment start date.
func (m *DeviceManagementAutopilotEvent) GetEnrollmentStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.enrollmentStartDateTime
    }
}
// Gets the enrollmentState property value. Enrollment state like Enrolled, Failed. Possible values are: unknown, enrolled, pendingReset, failed, notContacted, blocked.
func (m *DeviceManagementAutopilotEvent) GetEnrollmentState()(*EnrollmentState) {
    if m == nil {
        return nil
    } else {
        return m.enrollmentState
    }
}
// Gets the enrollmentType property value. Enrollment type. Possible values are: unknown, azureADJoinedWithAutopilotProfile, offlineDomainJoined, azureADJoinedUsingDeviceAuthWithAutopilotProfile, azureADJoinedUsingDeviceAuthWithoutAutopilotProfile, azureADJoinedWithOfflineAutopilotProfile, azureADJoinedWithWhiteGlove, offlineDomainJoinedWithWhiteGlove, offlineDomainJoinedWithOfflineAutopilotProfile.
func (m *DeviceManagementAutopilotEvent) GetEnrollmentType()(*WindowsAutopilotEnrollmentType) {
    if m == nil {
        return nil
    } else {
        return m.enrollmentType
    }
}
// Gets the eventDateTime property value. Time when the event occurred .
func (m *DeviceManagementAutopilotEvent) GetEventDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.eventDateTime
    }
}
// Gets the managedDeviceName property value. Managed device name.
func (m *DeviceManagementAutopilotEvent) GetManagedDeviceName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.managedDeviceName
    }
}
// Gets the osVersion property value. Device operating system version.
func (m *DeviceManagementAutopilotEvent) GetOsVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.osVersion
    }
}
// Gets the policyStatusDetails property value. Policy and application status details for this device.
func (m *DeviceManagementAutopilotEvent) GetPolicyStatusDetails()([]DeviceManagementAutopilotPolicyStatusDetail) {
    if m == nil {
        return nil
    } else {
        return m.policyStatusDetails
    }
}
// Gets the targetedAppCount property value. Count of applications targeted.
func (m *DeviceManagementAutopilotEvent) GetTargetedAppCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.targetedAppCount
    }
}
// Gets the targetedPolicyCount property value. Count of policies targeted.
func (m *DeviceManagementAutopilotEvent) GetTargetedPolicyCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.targetedPolicyCount
    }
}
// Gets the userPrincipalName property value. User principal name used to enroll the device.
func (m *DeviceManagementAutopilotEvent) GetUserPrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userPrincipalName
    }
}
// Gets the windows10EnrollmentCompletionPageConfigurationDisplayName property value. Enrollment Status Page profile name
func (m *DeviceManagementAutopilotEvent) GetWindows10EnrollmentCompletionPageConfigurationDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.windows10EnrollmentCompletionPageConfigurationDisplayName
    }
}
// Gets the windows10EnrollmentCompletionPageConfigurationId property value. Enrollment Status Page profile ID
func (m *DeviceManagementAutopilotEvent) GetWindows10EnrollmentCompletionPageConfigurationId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.windows10EnrollmentCompletionPageConfigurationId
    }
}
// Gets the windowsAutopilotDeploymentProfileDisplayName property value. Autopilot profile name.
func (m *DeviceManagementAutopilotEvent) GetWindowsAutopilotDeploymentProfileDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.windowsAutopilotDeploymentProfileDisplayName
    }
}
// The deserialization information for the current model
func (m *DeviceManagementAutopilotEvent) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["accountSetupDuration"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAccountSetupDuration(val)
        return nil
    }
    res["accountSetupStatus"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseWindowsAutopilotDeploymentState)
        if err != nil {
            return err
        }
        cast := val.(WindowsAutopilotDeploymentState)
        m.SetAccountSetupStatus(&cast)
        return nil
    }
    res["deploymentDuration"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDeploymentDuration(val)
        return nil
    }
    res["deploymentEndDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetDeploymentEndDateTime(val)
        return nil
    }
    res["deploymentStartDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetDeploymentStartDateTime(val)
        return nil
    }
    res["deploymentState"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseWindowsAutopilotDeploymentState)
        if err != nil {
            return err
        }
        cast := val.(WindowsAutopilotDeploymentState)
        m.SetDeploymentState(&cast)
        return nil
    }
    res["deploymentTotalDuration"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDeploymentTotalDuration(val)
        return nil
    }
    res["deviceId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDeviceId(val)
        return nil
    }
    res["devicePreparationDuration"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDevicePreparationDuration(val)
        return nil
    }
    res["deviceRegisteredDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetDeviceRegisteredDateTime(val)
        return nil
    }
    res["deviceSerialNumber"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDeviceSerialNumber(val)
        return nil
    }
    res["deviceSetupDuration"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDeviceSetupDuration(val)
        return nil
    }
    res["deviceSetupStatus"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseWindowsAutopilotDeploymentState)
        if err != nil {
            return err
        }
        cast := val.(WindowsAutopilotDeploymentState)
        m.SetDeviceSetupStatus(&cast)
        return nil
    }
    res["enrollmentFailureDetails"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetEnrollmentFailureDetails(val)
        return nil
    }
    res["enrollmentStartDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetEnrollmentStartDateTime(val)
        return nil
    }
    res["enrollmentState"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseEnrollmentState)
        if err != nil {
            return err
        }
        cast := val.(EnrollmentState)
        m.SetEnrollmentState(&cast)
        return nil
    }
    res["enrollmentType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseWindowsAutopilotEnrollmentType)
        if err != nil {
            return err
        }
        cast := val.(WindowsAutopilotEnrollmentType)
        m.SetEnrollmentType(&cast)
        return nil
    }
    res["eventDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetEventDateTime(val)
        return nil
    }
    res["managedDeviceName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetManagedDeviceName(val)
        return nil
    }
    res["osVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetOsVersion(val)
        return nil
    }
    res["policyStatusDetails"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceManagementAutopilotPolicyStatusDetail() })
        if err != nil {
            return err
        }
        res := make([]DeviceManagementAutopilotPolicyStatusDetail, len(val))
        for i, v := range val {
            res[i] = *(v.(*DeviceManagementAutopilotPolicyStatusDetail))
        }
        m.SetPolicyStatusDetails(res)
        return nil
    }
    res["targetedAppCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetTargetedAppCount(val)
        return nil
    }
    res["targetedPolicyCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetTargetedPolicyCount(val)
        return nil
    }
    res["userPrincipalName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetUserPrincipalName(val)
        return nil
    }
    res["windows10EnrollmentCompletionPageConfigurationDisplayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetWindows10EnrollmentCompletionPageConfigurationDisplayName(val)
        return nil
    }
    res["windows10EnrollmentCompletionPageConfigurationId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetWindows10EnrollmentCompletionPageConfigurationId(val)
        return nil
    }
    res["windowsAutopilotDeploymentProfileDisplayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetWindowsAutopilotDeploymentProfileDisplayName(val)
        return nil
    }
    return res
}
func (m *DeviceManagementAutopilotEvent) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *DeviceManagementAutopilotEvent) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("accountSetupDuration", m.GetAccountSetupDuration())
        if err != nil {
            return err
        }
    }
    if m.GetAccountSetupStatus() != nil {
        cast := m.GetAccountSetupStatus().String()
        err = writer.WriteStringValue("accountSetupStatus", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deploymentDuration", m.GetDeploymentDuration())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("deploymentEndDateTime", m.GetDeploymentEndDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("deploymentStartDateTime", m.GetDeploymentStartDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetDeploymentState() != nil {
        cast := m.GetDeploymentState().String()
        err = writer.WriteStringValue("deploymentState", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deploymentTotalDuration", m.GetDeploymentTotalDuration())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deviceId", m.GetDeviceId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("devicePreparationDuration", m.GetDevicePreparationDuration())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("deviceRegisteredDateTime", m.GetDeviceRegisteredDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deviceSerialNumber", m.GetDeviceSerialNumber())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deviceSetupDuration", m.GetDeviceSetupDuration())
        if err != nil {
            return err
        }
    }
    if m.GetDeviceSetupStatus() != nil {
        cast := m.GetDeviceSetupStatus().String()
        err = writer.WriteStringValue("deviceSetupStatus", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("enrollmentFailureDetails", m.GetEnrollmentFailureDetails())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("enrollmentStartDateTime", m.GetEnrollmentStartDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetEnrollmentState() != nil {
        cast := m.GetEnrollmentState().String()
        err = writer.WriteStringValue("enrollmentState", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetEnrollmentType() != nil {
        cast := m.GetEnrollmentType().String()
        err = writer.WriteStringValue("enrollmentType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("eventDateTime", m.GetEventDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("managedDeviceName", m.GetManagedDeviceName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("osVersion", m.GetOsVersion())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetPolicyStatusDetails()))
        for i, v := range m.GetPolicyStatusDetails() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("policyStatusDetails", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("targetedAppCount", m.GetTargetedAppCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("targetedPolicyCount", m.GetTargetedPolicyCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userPrincipalName", m.GetUserPrincipalName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("windows10EnrollmentCompletionPageConfigurationDisplayName", m.GetWindows10EnrollmentCompletionPageConfigurationDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("windows10EnrollmentCompletionPageConfigurationId", m.GetWindows10EnrollmentCompletionPageConfigurationId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("windowsAutopilotDeploymentProfileDisplayName", m.GetWindowsAutopilotDeploymentProfileDisplayName())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the accountSetupDuration property value. Time spent in user ESP.
// Parameters:
//  - value : Value to set for the accountSetupDuration property.
func (m *DeviceManagementAutopilotEvent) SetAccountSetupDuration(value *string)() {
    m.accountSetupDuration = value
}
// Sets the accountSetupStatus property value. Deployment status for the enrollment status page account setup phase. Possible values are: unknown, success, inProgress, failure, successWithTimeout, notAttempted, disabled.
// Parameters:
//  - value : Value to set for the accountSetupStatus property.
func (m *DeviceManagementAutopilotEvent) SetAccountSetupStatus(value *WindowsAutopilotDeploymentState)() {
    m.accountSetupStatus = value
}
// Sets the deploymentDuration property value. Autopilot deployment duration including enrollment.
// Parameters:
//  - value : Value to set for the deploymentDuration property.
func (m *DeviceManagementAutopilotEvent) SetDeploymentDuration(value *string)() {
    m.deploymentDuration = value
}
// Sets the deploymentEndDateTime property value. Deployment end time.
// Parameters:
//  - value : Value to set for the deploymentEndDateTime property.
func (m *DeviceManagementAutopilotEvent) SetDeploymentEndDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.deploymentEndDateTime = value
}
// Sets the deploymentStartDateTime property value. Deployment start time.
// Parameters:
//  - value : Value to set for the deploymentStartDateTime property.
func (m *DeviceManagementAutopilotEvent) SetDeploymentStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.deploymentStartDateTime = value
}
// Sets the deploymentState property value. Deployment state like Success, Failure, InProgress, SuccessWithTimeout. Possible values are: unknown, success, inProgress, failure, successWithTimeout, notAttempted, disabled.
// Parameters:
//  - value : Value to set for the deploymentState property.
func (m *DeviceManagementAutopilotEvent) SetDeploymentState(value *WindowsAutopilotDeploymentState)() {
    m.deploymentState = value
}
// Sets the deploymentTotalDuration property value. Total deployment duration from enrollment to Desktop screen.
// Parameters:
//  - value : Value to set for the deploymentTotalDuration property.
func (m *DeviceManagementAutopilotEvent) SetDeploymentTotalDuration(value *string)() {
    m.deploymentTotalDuration = value
}
// Sets the deviceId property value. Device id associated with the object
// Parameters:
//  - value : Value to set for the deviceId property.
func (m *DeviceManagementAutopilotEvent) SetDeviceId(value *string)() {
    m.deviceId = value
}
// Sets the devicePreparationDuration property value. Time spent in device enrollment.
// Parameters:
//  - value : Value to set for the devicePreparationDuration property.
func (m *DeviceManagementAutopilotEvent) SetDevicePreparationDuration(value *string)() {
    m.devicePreparationDuration = value
}
// Sets the deviceRegisteredDateTime property value. Device registration date.
// Parameters:
//  - value : Value to set for the deviceRegisteredDateTime property.
func (m *DeviceManagementAutopilotEvent) SetDeviceRegisteredDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.deviceRegisteredDateTime = value
}
// Sets the deviceSerialNumber property value. Device serial number.
// Parameters:
//  - value : Value to set for the deviceSerialNumber property.
func (m *DeviceManagementAutopilotEvent) SetDeviceSerialNumber(value *string)() {
    m.deviceSerialNumber = value
}
// Sets the deviceSetupDuration property value. Time spent in device ESP.
// Parameters:
//  - value : Value to set for the deviceSetupDuration property.
func (m *DeviceManagementAutopilotEvent) SetDeviceSetupDuration(value *string)() {
    m.deviceSetupDuration = value
}
// Sets the deviceSetupStatus property value. Deployment status for the enrollment status page device setup phase. Possible values are: unknown, success, inProgress, failure, successWithTimeout, notAttempted, disabled.
// Parameters:
//  - value : Value to set for the deviceSetupStatus property.
func (m *DeviceManagementAutopilotEvent) SetDeviceSetupStatus(value *WindowsAutopilotDeploymentState)() {
    m.deviceSetupStatus = value
}
// Sets the enrollmentFailureDetails property value. Enrollment failure details.
// Parameters:
//  - value : Value to set for the enrollmentFailureDetails property.
func (m *DeviceManagementAutopilotEvent) SetEnrollmentFailureDetails(value *string)() {
    m.enrollmentFailureDetails = value
}
// Sets the enrollmentStartDateTime property value. Device enrollment start date.
// Parameters:
//  - value : Value to set for the enrollmentStartDateTime property.
func (m *DeviceManagementAutopilotEvent) SetEnrollmentStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.enrollmentStartDateTime = value
}
// Sets the enrollmentState property value. Enrollment state like Enrolled, Failed. Possible values are: unknown, enrolled, pendingReset, failed, notContacted, blocked.
// Parameters:
//  - value : Value to set for the enrollmentState property.
func (m *DeviceManagementAutopilotEvent) SetEnrollmentState(value *EnrollmentState)() {
    m.enrollmentState = value
}
// Sets the enrollmentType property value. Enrollment type. Possible values are: unknown, azureADJoinedWithAutopilotProfile, offlineDomainJoined, azureADJoinedUsingDeviceAuthWithAutopilotProfile, azureADJoinedUsingDeviceAuthWithoutAutopilotProfile, azureADJoinedWithOfflineAutopilotProfile, azureADJoinedWithWhiteGlove, offlineDomainJoinedWithWhiteGlove, offlineDomainJoinedWithOfflineAutopilotProfile.
// Parameters:
//  - value : Value to set for the enrollmentType property.
func (m *DeviceManagementAutopilotEvent) SetEnrollmentType(value *WindowsAutopilotEnrollmentType)() {
    m.enrollmentType = value
}
// Sets the eventDateTime property value. Time when the event occurred .
// Parameters:
//  - value : Value to set for the eventDateTime property.
func (m *DeviceManagementAutopilotEvent) SetEventDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.eventDateTime = value
}
// Sets the managedDeviceName property value. Managed device name.
// Parameters:
//  - value : Value to set for the managedDeviceName property.
func (m *DeviceManagementAutopilotEvent) SetManagedDeviceName(value *string)() {
    m.managedDeviceName = value
}
// Sets the osVersion property value. Device operating system version.
// Parameters:
//  - value : Value to set for the osVersion property.
func (m *DeviceManagementAutopilotEvent) SetOsVersion(value *string)() {
    m.osVersion = value
}
// Sets the policyStatusDetails property value. Policy and application status details for this device.
// Parameters:
//  - value : Value to set for the policyStatusDetails property.
func (m *DeviceManagementAutopilotEvent) SetPolicyStatusDetails(value []DeviceManagementAutopilotPolicyStatusDetail)() {
    m.policyStatusDetails = value
}
// Sets the targetedAppCount property value. Count of applications targeted.
// Parameters:
//  - value : Value to set for the targetedAppCount property.
func (m *DeviceManagementAutopilotEvent) SetTargetedAppCount(value *int32)() {
    m.targetedAppCount = value
}
// Sets the targetedPolicyCount property value. Count of policies targeted.
// Parameters:
//  - value : Value to set for the targetedPolicyCount property.
func (m *DeviceManagementAutopilotEvent) SetTargetedPolicyCount(value *int32)() {
    m.targetedPolicyCount = value
}
// Sets the userPrincipalName property value. User principal name used to enroll the device.
// Parameters:
//  - value : Value to set for the userPrincipalName property.
func (m *DeviceManagementAutopilotEvent) SetUserPrincipalName(value *string)() {
    m.userPrincipalName = value
}
// Sets the windows10EnrollmentCompletionPageConfigurationDisplayName property value. Enrollment Status Page profile name
// Parameters:
//  - value : Value to set for the windows10EnrollmentCompletionPageConfigurationDisplayName property.
func (m *DeviceManagementAutopilotEvent) SetWindows10EnrollmentCompletionPageConfigurationDisplayName(value *string)() {
    m.windows10EnrollmentCompletionPageConfigurationDisplayName = value
}
// Sets the windows10EnrollmentCompletionPageConfigurationId property value. Enrollment Status Page profile ID
// Parameters:
//  - value : Value to set for the windows10EnrollmentCompletionPageConfigurationId property.
func (m *DeviceManagementAutopilotEvent) SetWindows10EnrollmentCompletionPageConfigurationId(value *string)() {
    m.windows10EnrollmentCompletionPageConfigurationId = value
}
// Sets the windowsAutopilotDeploymentProfileDisplayName property value. Autopilot profile name.
// Parameters:
//  - value : Value to set for the windowsAutopilotDeploymentProfileDisplayName property.
func (m *DeviceManagementAutopilotEvent) SetWindowsAutopilotDeploymentProfileDisplayName(value *string)() {
    m.windowsAutopilotDeploymentProfileDisplayName = value
}
