package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus 
type WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus struct {
    Entity
    // The deployment state of the policy. Possible values are: unknown, success, tokenError, notAuthorizedByToken, policyNotFound.
    deploymentStatus *WindowsDefenderApplicationControlSupplementalPolicyStatuses;
    // Device ID.
    deviceId *string;
    // Device name.
    deviceName *string;
    // Last sync date time.
    lastSyncDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Windows OS Version Description.
    osDescription *string;
    // Windows OS Version.
    osVersion *string;
    // The navigation link to the WindowsDefenderApplicationControl supplemental policy.
    policy *WindowsDefenderApplicationControlSupplementalPolicy;
    // Human readable version of the WindowsDefenderApplicationControl supplemental policy.
    policyVersion *string;
    // The name of the user of this device.
    userName *string;
    // User Principal Name.
    userPrincipalName *string;
}
// NewWindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus instantiates a new windowsDefenderApplicationControlSupplementalPolicyDeploymentStatus and sets the default values.
func NewWindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus()(*WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus) {
    m := &WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus{
        Entity: *NewEntity(),
    }
    return m
}
// GetDeploymentStatus gets the deploymentStatus property value. The deployment state of the policy. Possible values are: unknown, success, tokenError, notAuthorizedByToken, policyNotFound.
func (m *WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus) GetDeploymentStatus()(*WindowsDefenderApplicationControlSupplementalPolicyStatuses) {
    if m == nil {
        return nil
    } else {
        return m.deploymentStatus
    }
}
// GetDeviceId gets the deviceId property value. Device ID.
func (m *WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus) GetDeviceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceId
    }
}
// GetDeviceName gets the deviceName property value. Device name.
func (m *WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus) GetDeviceName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceName
    }
}
// GetLastSyncDateTime gets the lastSyncDateTime property value. Last sync date time.
func (m *WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus) GetLastSyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastSyncDateTime
    }
}
// GetOsDescription gets the osDescription property value. Windows OS Version Description.
func (m *WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus) GetOsDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.osDescription
    }
}
// GetOsVersion gets the osVersion property value. Windows OS Version.
func (m *WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus) GetOsVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.osVersion
    }
}
// GetPolicy gets the policy property value. The navigation link to the WindowsDefenderApplicationControl supplemental policy.
func (m *WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus) GetPolicy()(*WindowsDefenderApplicationControlSupplementalPolicy) {
    if m == nil {
        return nil
    } else {
        return m.policy
    }
}
// GetPolicyVersion gets the policyVersion property value. Human readable version of the WindowsDefenderApplicationControl supplemental policy.
func (m *WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus) GetPolicyVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.policyVersion
    }
}
// GetUserName gets the userName property value. The name of the user of this device.
func (m *WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus) GetUserName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userName
    }
}
// GetUserPrincipalName gets the userPrincipalName property value. User Principal Name.
func (m *WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus) GetUserPrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userPrincipalName
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["deploymentStatus"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseWindowsDefenderApplicationControlSupplementalPolicyStatuses)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeploymentStatus(val.(*WindowsDefenderApplicationControlSupplementalPolicyStatuses))
        }
        return nil
    }
    res["deviceId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceId(val)
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
    res["lastSyncDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastSyncDateTime(val)
        }
        return nil
    }
    res["osDescription"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOsDescription(val)
        }
        return nil
    }
    res["osVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOsVersion(val)
        }
        return nil
    }
    res["policy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWindowsDefenderApplicationControlSupplementalPolicy() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPolicy(val.(*WindowsDefenderApplicationControlSupplementalPolicy))
        }
        return nil
    }
    res["policyVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPolicyVersion(val)
        }
        return nil
    }
    res["userName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserName(val)
        }
        return nil
    }
    res["userPrincipalName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserPrincipalName(val)
        }
        return nil
    }
    return res
}
func (m *WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetDeploymentStatus() != nil {
        cast := (*m.GetDeploymentStatus()).String()
        err = writer.WriteStringValue("deploymentStatus", &cast)
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
        err = writer.WriteStringValue("deviceName", m.GetDeviceName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastSyncDateTime", m.GetLastSyncDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("osDescription", m.GetOsDescription())
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
        err = writer.WriteObjectValue("policy", m.GetPolicy())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("policyVersion", m.GetPolicyVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userName", m.GetUserName())
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
    return nil
}
// SetDeploymentStatus sets the deploymentStatus property value. The deployment state of the policy. Possible values are: unknown, success, tokenError, notAuthorizedByToken, policyNotFound.
func (m *WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus) SetDeploymentStatus(value *WindowsDefenderApplicationControlSupplementalPolicyStatuses)() {
    if m != nil {
        m.deploymentStatus = value
    }
}
// SetDeviceId sets the deviceId property value. Device ID.
func (m *WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus) SetDeviceId(value *string)() {
    if m != nil {
        m.deviceId = value
    }
}
// SetDeviceName sets the deviceName property value. Device name.
func (m *WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus) SetDeviceName(value *string)() {
    if m != nil {
        m.deviceName = value
    }
}
// SetLastSyncDateTime sets the lastSyncDateTime property value. Last sync date time.
func (m *WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus) SetLastSyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastSyncDateTime = value
    }
}
// SetOsDescription sets the osDescription property value. Windows OS Version Description.
func (m *WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus) SetOsDescription(value *string)() {
    if m != nil {
        m.osDescription = value
    }
}
// SetOsVersion sets the osVersion property value. Windows OS Version.
func (m *WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus) SetOsVersion(value *string)() {
    if m != nil {
        m.osVersion = value
    }
}
// SetPolicy sets the policy property value. The navigation link to the WindowsDefenderApplicationControl supplemental policy.
func (m *WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus) SetPolicy(value *WindowsDefenderApplicationControlSupplementalPolicy)() {
    if m != nil {
        m.policy = value
    }
}
// SetPolicyVersion sets the policyVersion property value. Human readable version of the WindowsDefenderApplicationControl supplemental policy.
func (m *WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus) SetPolicyVersion(value *string)() {
    if m != nil {
        m.policyVersion = value
    }
}
// SetUserName sets the userName property value. The name of the user of this device.
func (m *WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus) SetUserName(value *string)() {
    if m != nil {
        m.userName = value
    }
}
// SetUserPrincipalName sets the userPrincipalName property value. User Principal Name.
func (m *WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus) SetUserPrincipalName(value *string)() {
    if m != nil {
        m.userPrincipalName = value
    }
}
