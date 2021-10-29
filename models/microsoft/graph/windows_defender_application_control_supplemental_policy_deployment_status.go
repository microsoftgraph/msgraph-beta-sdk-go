package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
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
// Instantiates a new windowsDefenderApplicationControlSupplementalPolicyDeploymentStatus and sets the default values.
func NewWindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus()(*WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus) {
    m := &WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the deploymentStatus property value. The deployment state of the policy. Possible values are: unknown, success, tokenError, notAuthorizedByToken, policyNotFound.
func (m *WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus) GetDeploymentStatus()(*WindowsDefenderApplicationControlSupplementalPolicyStatuses) {
    if m == nil {
        return nil
    } else {
        return m.deploymentStatus
    }
}
// Gets the deviceId property value. Device ID.
func (m *WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus) GetDeviceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceId
    }
}
// Gets the deviceName property value. Device name.
func (m *WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus) GetDeviceName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceName
    }
}
// Gets the lastSyncDateTime property value. Last sync date time.
func (m *WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus) GetLastSyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastSyncDateTime
    }
}
// Gets the osDescription property value. Windows OS Version Description.
func (m *WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus) GetOsDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.osDescription
    }
}
// Gets the osVersion property value. Windows OS Version.
func (m *WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus) GetOsVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.osVersion
    }
}
// Gets the policy property value. The navigation link to the WindowsDefenderApplicationControl supplemental policy.
func (m *WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus) GetPolicy()(*WindowsDefenderApplicationControlSupplementalPolicy) {
    if m == nil {
        return nil
    } else {
        return m.policy
    }
}
// Gets the policyVersion property value. Human readable version of the WindowsDefenderApplicationControl supplemental policy.
func (m *WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus) GetPolicyVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.policyVersion
    }
}
// Gets the userName property value. The name of the user of this device.
func (m *WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus) GetUserName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userName
    }
}
// Gets the userPrincipalName property value. User Principal Name.
func (m *WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus) GetUserPrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userPrincipalName
    }
}
// The deserialization information for the current model
func (m *WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["deploymentStatus"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseWindowsDefenderApplicationControlSupplementalPolicyStatuses)
        if err != nil {
            return err
        }
        cast := val.(WindowsDefenderApplicationControlSupplementalPolicyStatuses)
        m.SetDeploymentStatus(&cast)
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
    res["deviceName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDeviceName(val)
        return nil
    }
    res["lastSyncDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetLastSyncDateTime(val)
        return nil
    }
    res["osDescription"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetOsDescription(val)
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
    res["policy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWindowsDefenderApplicationControlSupplementalPolicy() })
        if err != nil {
            return err
        }
        m.SetPolicy(val.(*WindowsDefenderApplicationControlSupplementalPolicy))
        return nil
    }
    res["policyVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetPolicyVersion(val)
        return nil
    }
    res["userName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetUserName(val)
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
    return res
}
func (m *WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetDeploymentStatus() != nil {
        cast := m.GetDeploymentStatus().String()
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
// Sets the deploymentStatus property value. The deployment state of the policy. Possible values are: unknown, success, tokenError, notAuthorizedByToken, policyNotFound.
// Parameters:
//  - value : Value to set for the deploymentStatus property.
func (m *WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus) SetDeploymentStatus(value *WindowsDefenderApplicationControlSupplementalPolicyStatuses)() {
    m.deploymentStatus = value
}
// Sets the deviceId property value. Device ID.
// Parameters:
//  - value : Value to set for the deviceId property.
func (m *WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus) SetDeviceId(value *string)() {
    m.deviceId = value
}
// Sets the deviceName property value. Device name.
// Parameters:
//  - value : Value to set for the deviceName property.
func (m *WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus) SetDeviceName(value *string)() {
    m.deviceName = value
}
// Sets the lastSyncDateTime property value. Last sync date time.
// Parameters:
//  - value : Value to set for the lastSyncDateTime property.
func (m *WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus) SetLastSyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastSyncDateTime = value
}
// Sets the osDescription property value. Windows OS Version Description.
// Parameters:
//  - value : Value to set for the osDescription property.
func (m *WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus) SetOsDescription(value *string)() {
    m.osDescription = value
}
// Sets the osVersion property value. Windows OS Version.
// Parameters:
//  - value : Value to set for the osVersion property.
func (m *WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus) SetOsVersion(value *string)() {
    m.osVersion = value
}
// Sets the policy property value. The navigation link to the WindowsDefenderApplicationControl supplemental policy.
// Parameters:
//  - value : Value to set for the policy property.
func (m *WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus) SetPolicy(value *WindowsDefenderApplicationControlSupplementalPolicy)() {
    m.policy = value
}
// Sets the policyVersion property value. Human readable version of the WindowsDefenderApplicationControl supplemental policy.
// Parameters:
//  - value : Value to set for the policyVersion property.
func (m *WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus) SetPolicyVersion(value *string)() {
    m.policyVersion = value
}
// Sets the userName property value. The name of the user of this device.
// Parameters:
//  - value : Value to set for the userName property.
func (m *WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus) SetUserName(value *string)() {
    m.userName = value
}
// Sets the userPrincipalName property value. User Principal Name.
// Parameters:
//  - value : Value to set for the userPrincipalName property.
func (m *WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus) SetUserPrincipalName(value *string)() {
    m.userPrincipalName = value
}
