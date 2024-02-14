package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsAutopilotDeviceIdentity the windowsAutopilotDeviceIdentity resource represents a Windows Autopilot Device.
type WindowsAutopilotDeviceIdentity struct {
    Entity
}
// NewWindowsAutopilotDeviceIdentity instantiates a new WindowsAutopilotDeviceIdentity and sets the default values.
func NewWindowsAutopilotDeviceIdentity()(*WindowsAutopilotDeviceIdentity) {
    m := &WindowsAutopilotDeviceIdentity{
        Entity: *NewEntity(),
    }
    return m
}
// CreateWindowsAutopilotDeviceIdentityFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateWindowsAutopilotDeviceIdentityFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsAutopilotDeviceIdentity(), nil
}
// GetAddressableUserName gets the addressableUserName property value. Addressable user name.
// returns a *string when successful
func (m *WindowsAutopilotDeviceIdentity) GetAddressableUserName()(*string) {
    val, err := m.GetBackingStore().Get("addressableUserName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetAzureActiveDirectoryDeviceId gets the azureActiveDirectoryDeviceId property value. AAD Device ID - to be deprecated
// returns a *string when successful
func (m *WindowsAutopilotDeviceIdentity) GetAzureActiveDirectoryDeviceId()(*string) {
    val, err := m.GetBackingStore().Get("azureActiveDirectoryDeviceId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetAzureAdDeviceId gets the azureAdDeviceId property value. AAD Device ID
// returns a *string when successful
func (m *WindowsAutopilotDeviceIdentity) GetAzureAdDeviceId()(*string) {
    val, err := m.GetBackingStore().Get("azureAdDeviceId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDeploymentProfile gets the deploymentProfile property value. Deployment profile currently assigned to the Windows autopilot device.
// returns a WindowsAutopilotDeploymentProfileable when successful
func (m *WindowsAutopilotDeviceIdentity) GetDeploymentProfile()(WindowsAutopilotDeploymentProfileable) {
    val, err := m.GetBackingStore().Get("deploymentProfile")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(WindowsAutopilotDeploymentProfileable)
    }
    return nil
}
// GetDeploymentProfileAssignedDateTime gets the deploymentProfileAssignedDateTime property value. Profile set time of the Windows autopilot device.
// returns a *Time when successful
func (m *WindowsAutopilotDeviceIdentity) GetDeploymentProfileAssignedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("deploymentProfileAssignedDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetDeploymentProfileAssignmentDetailedStatus gets the deploymentProfileAssignmentDetailedStatus property value. The deploymentProfileAssignmentDetailedStatus property
// returns a *WindowsAutopilotProfileAssignmentDetailedStatus when successful
func (m *WindowsAutopilotDeviceIdentity) GetDeploymentProfileAssignmentDetailedStatus()(*WindowsAutopilotProfileAssignmentDetailedStatus) {
    val, err := m.GetBackingStore().Get("deploymentProfileAssignmentDetailedStatus")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*WindowsAutopilotProfileAssignmentDetailedStatus)
    }
    return nil
}
// GetDeploymentProfileAssignmentStatus gets the deploymentProfileAssignmentStatus property value. The deploymentProfileAssignmentStatus property
// returns a *WindowsAutopilotProfileAssignmentStatus when successful
func (m *WindowsAutopilotDeviceIdentity) GetDeploymentProfileAssignmentStatus()(*WindowsAutopilotProfileAssignmentStatus) {
    val, err := m.GetBackingStore().Get("deploymentProfileAssignmentStatus")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*WindowsAutopilotProfileAssignmentStatus)
    }
    return nil
}
// GetDeviceAccountPassword gets the deviceAccountPassword property value. Surface Hub Device Account Password
// returns a *string when successful
func (m *WindowsAutopilotDeviceIdentity) GetDeviceAccountPassword()(*string) {
    val, err := m.GetBackingStore().Get("deviceAccountPassword")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDeviceAccountUpn gets the deviceAccountUpn property value. Surface Hub Device Account Upn
// returns a *string when successful
func (m *WindowsAutopilotDeviceIdentity) GetDeviceAccountUpn()(*string) {
    val, err := m.GetBackingStore().Get("deviceAccountUpn")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDeviceFriendlyName gets the deviceFriendlyName property value. Surface Hub Device Friendly Name
// returns a *string when successful
func (m *WindowsAutopilotDeviceIdentity) GetDeviceFriendlyName()(*string) {
    val, err := m.GetBackingStore().Get("deviceFriendlyName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDisplayName gets the displayName property value. Display Name
// returns a *string when successful
func (m *WindowsAutopilotDeviceIdentity) GetDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("displayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetEnrollmentState gets the enrollmentState property value. The enrollmentState property
// returns a *EnrollmentState when successful
func (m *WindowsAutopilotDeviceIdentity) GetEnrollmentState()(*EnrollmentState) {
    val, err := m.GetBackingStore().Get("enrollmentState")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*EnrollmentState)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *WindowsAutopilotDeviceIdentity) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["addressableUserName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAddressableUserName(val)
        }
        return nil
    }
    res["azureActiveDirectoryDeviceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAzureActiveDirectoryDeviceId(val)
        }
        return nil
    }
    res["azureAdDeviceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAzureAdDeviceId(val)
        }
        return nil
    }
    res["deploymentProfile"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWindowsAutopilotDeploymentProfileFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeploymentProfile(val.(WindowsAutopilotDeploymentProfileable))
        }
        return nil
    }
    res["deploymentProfileAssignedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeploymentProfileAssignedDateTime(val)
        }
        return nil
    }
    res["deploymentProfileAssignmentDetailedStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseWindowsAutopilotProfileAssignmentDetailedStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeploymentProfileAssignmentDetailedStatus(val.(*WindowsAutopilotProfileAssignmentDetailedStatus))
        }
        return nil
    }
    res["deploymentProfileAssignmentStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseWindowsAutopilotProfileAssignmentStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeploymentProfileAssignmentStatus(val.(*WindowsAutopilotProfileAssignmentStatus))
        }
        return nil
    }
    res["deviceAccountPassword"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceAccountPassword(val)
        }
        return nil
    }
    res["deviceAccountUpn"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceAccountUpn(val)
        }
        return nil
    }
    res["deviceFriendlyName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceFriendlyName(val)
        }
        return nil
    }
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["enrollmentState"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEnrollmentState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnrollmentState(val.(*EnrollmentState))
        }
        return nil
    }
    res["groupTag"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGroupTag(val)
        }
        return nil
    }
    res["intendedDeploymentProfile"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWindowsAutopilotDeploymentProfileFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIntendedDeploymentProfile(val.(WindowsAutopilotDeploymentProfileable))
        }
        return nil
    }
    res["lastContactedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastContactedDateTime(val)
        }
        return nil
    }
    res["managedDeviceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagedDeviceId(val)
        }
        return nil
    }
    res["manufacturer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManufacturer(val)
        }
        return nil
    }
    res["model"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetModel(val)
        }
        return nil
    }
    res["productKey"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProductKey(val)
        }
        return nil
    }
    res["purchaseOrderIdentifier"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPurchaseOrderIdentifier(val)
        }
        return nil
    }
    res["remediationState"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseWindowsAutopilotDeviceRemediationState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRemediationState(val.(*WindowsAutopilotDeviceRemediationState))
        }
        return nil
    }
    res["remediationStateLastModifiedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRemediationStateLastModifiedDateTime(val)
        }
        return nil
    }
    res["resourceName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResourceName(val)
        }
        return nil
    }
    res["serialNumber"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSerialNumber(val)
        }
        return nil
    }
    res["skuNumber"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSkuNumber(val)
        }
        return nil
    }
    res["systemFamily"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSystemFamily(val)
        }
        return nil
    }
    res["userlessEnrollmentStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseWindowsAutopilotUserlessEnrollmentStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserlessEnrollmentStatus(val.(*WindowsAutopilotUserlessEnrollmentStatus))
        }
        return nil
    }
    res["userPrincipalName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
// GetGroupTag gets the groupTag property value. Group Tag of the Windows autopilot device.
// returns a *string when successful
func (m *WindowsAutopilotDeviceIdentity) GetGroupTag()(*string) {
    val, err := m.GetBackingStore().Get("groupTag")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetIntendedDeploymentProfile gets the intendedDeploymentProfile property value. Deployment profile intended to be assigned to the Windows autopilot device.
// returns a WindowsAutopilotDeploymentProfileable when successful
func (m *WindowsAutopilotDeviceIdentity) GetIntendedDeploymentProfile()(WindowsAutopilotDeploymentProfileable) {
    val, err := m.GetBackingStore().Get("intendedDeploymentProfile")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(WindowsAutopilotDeploymentProfileable)
    }
    return nil
}
// GetLastContactedDateTime gets the lastContactedDateTime property value. Intune Last Contacted Date Time of the Windows autopilot device.
// returns a *Time when successful
func (m *WindowsAutopilotDeviceIdentity) GetLastContactedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("lastContactedDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetManagedDeviceId gets the managedDeviceId property value. Managed Device ID
// returns a *string when successful
func (m *WindowsAutopilotDeviceIdentity) GetManagedDeviceId()(*string) {
    val, err := m.GetBackingStore().Get("managedDeviceId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetManufacturer gets the manufacturer property value. Oem manufacturer of the Windows autopilot device.
// returns a *string when successful
func (m *WindowsAutopilotDeviceIdentity) GetManufacturer()(*string) {
    val, err := m.GetBackingStore().Get("manufacturer")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetModel gets the model property value. Model name of the Windows autopilot device.
// returns a *string when successful
func (m *WindowsAutopilotDeviceIdentity) GetModel()(*string) {
    val, err := m.GetBackingStore().Get("model")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetProductKey gets the productKey property value. Product Key of the Windows autopilot device.
// returns a *string when successful
func (m *WindowsAutopilotDeviceIdentity) GetProductKey()(*string) {
    val, err := m.GetBackingStore().Get("productKey")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPurchaseOrderIdentifier gets the purchaseOrderIdentifier property value. Purchase Order Identifier of the Windows autopilot device.
// returns a *string when successful
func (m *WindowsAutopilotDeviceIdentity) GetPurchaseOrderIdentifier()(*string) {
    val, err := m.GetBackingStore().Get("purchaseOrderIdentifier")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetRemediationState gets the remediationState property value. Device remediation status, indicating whether or not hardware has been changed for an Autopilot-registered device.
// returns a *WindowsAutopilotDeviceRemediationState when successful
func (m *WindowsAutopilotDeviceIdentity) GetRemediationState()(*WindowsAutopilotDeviceRemediationState) {
    val, err := m.GetBackingStore().Get("remediationState")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*WindowsAutopilotDeviceRemediationState)
    }
    return nil
}
// GetRemediationStateLastModifiedDateTime gets the remediationStateLastModifiedDateTime property value. RemediationState set time of Autopilot device.
// returns a *Time when successful
func (m *WindowsAutopilotDeviceIdentity) GetRemediationStateLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("remediationStateLastModifiedDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetResourceName gets the resourceName property value. Resource Name.
// returns a *string when successful
func (m *WindowsAutopilotDeviceIdentity) GetResourceName()(*string) {
    val, err := m.GetBackingStore().Get("resourceName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSerialNumber gets the serialNumber property value. Serial number of the Windows autopilot device.
// returns a *string when successful
func (m *WindowsAutopilotDeviceIdentity) GetSerialNumber()(*string) {
    val, err := m.GetBackingStore().Get("serialNumber")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSkuNumber gets the skuNumber property value. SKU Number
// returns a *string when successful
func (m *WindowsAutopilotDeviceIdentity) GetSkuNumber()(*string) {
    val, err := m.GetBackingStore().Get("skuNumber")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSystemFamily gets the systemFamily property value. System Family
// returns a *string when successful
func (m *WindowsAutopilotDeviceIdentity) GetSystemFamily()(*string) {
    val, err := m.GetBackingStore().Get("systemFamily")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetUserlessEnrollmentStatus gets the userlessEnrollmentStatus property value. Userless enrollment block status, indicating whether the next device enrollment will be blocked.
// returns a *WindowsAutopilotUserlessEnrollmentStatus when successful
func (m *WindowsAutopilotDeviceIdentity) GetUserlessEnrollmentStatus()(*WindowsAutopilotUserlessEnrollmentStatus) {
    val, err := m.GetBackingStore().Get("userlessEnrollmentStatus")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*WindowsAutopilotUserlessEnrollmentStatus)
    }
    return nil
}
// GetUserPrincipalName gets the userPrincipalName property value. User Principal Name.
// returns a *string when successful
func (m *WindowsAutopilotDeviceIdentity) GetUserPrincipalName()(*string) {
    val, err := m.GetBackingStore().Get("userPrincipalName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *WindowsAutopilotDeviceIdentity) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("addressableUserName", m.GetAddressableUserName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("azureActiveDirectoryDeviceId", m.GetAzureActiveDirectoryDeviceId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("azureAdDeviceId", m.GetAzureAdDeviceId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("deploymentProfile", m.GetDeploymentProfile())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("deploymentProfileAssignedDateTime", m.GetDeploymentProfileAssignedDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetDeploymentProfileAssignmentDetailedStatus() != nil {
        cast := (*m.GetDeploymentProfileAssignmentDetailedStatus()).String()
        err = writer.WriteStringValue("deploymentProfileAssignmentDetailedStatus", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetDeploymentProfileAssignmentStatus() != nil {
        cast := (*m.GetDeploymentProfileAssignmentStatus()).String()
        err = writer.WriteStringValue("deploymentProfileAssignmentStatus", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deviceAccountPassword", m.GetDeviceAccountPassword())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deviceAccountUpn", m.GetDeviceAccountUpn())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deviceFriendlyName", m.GetDeviceFriendlyName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    if m.GetEnrollmentState() != nil {
        cast := (*m.GetEnrollmentState()).String()
        err = writer.WriteStringValue("enrollmentState", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("groupTag", m.GetGroupTag())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("intendedDeploymentProfile", m.GetIntendedDeploymentProfile())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastContactedDateTime", m.GetLastContactedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("managedDeviceId", m.GetManagedDeviceId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("manufacturer", m.GetManufacturer())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("model", m.GetModel())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("productKey", m.GetProductKey())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("purchaseOrderIdentifier", m.GetPurchaseOrderIdentifier())
        if err != nil {
            return err
        }
    }
    if m.GetRemediationState() != nil {
        cast := (*m.GetRemediationState()).String()
        err = writer.WriteStringValue("remediationState", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("remediationStateLastModifiedDateTime", m.GetRemediationStateLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("resourceName", m.GetResourceName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("serialNumber", m.GetSerialNumber())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("skuNumber", m.GetSkuNumber())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("systemFamily", m.GetSystemFamily())
        if err != nil {
            return err
        }
    }
    if m.GetUserlessEnrollmentStatus() != nil {
        cast := (*m.GetUserlessEnrollmentStatus()).String()
        err = writer.WriteStringValue("userlessEnrollmentStatus", &cast)
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
// SetAddressableUserName sets the addressableUserName property value. Addressable user name.
func (m *WindowsAutopilotDeviceIdentity) SetAddressableUserName(value *string)() {
    err := m.GetBackingStore().Set("addressableUserName", value)
    if err != nil {
        panic(err)
    }
}
// SetAzureActiveDirectoryDeviceId sets the azureActiveDirectoryDeviceId property value. AAD Device ID - to be deprecated
func (m *WindowsAutopilotDeviceIdentity) SetAzureActiveDirectoryDeviceId(value *string)() {
    err := m.GetBackingStore().Set("azureActiveDirectoryDeviceId", value)
    if err != nil {
        panic(err)
    }
}
// SetAzureAdDeviceId sets the azureAdDeviceId property value. AAD Device ID
func (m *WindowsAutopilotDeviceIdentity) SetAzureAdDeviceId(value *string)() {
    err := m.GetBackingStore().Set("azureAdDeviceId", value)
    if err != nil {
        panic(err)
    }
}
// SetDeploymentProfile sets the deploymentProfile property value. Deployment profile currently assigned to the Windows autopilot device.
func (m *WindowsAutopilotDeviceIdentity) SetDeploymentProfile(value WindowsAutopilotDeploymentProfileable)() {
    err := m.GetBackingStore().Set("deploymentProfile", value)
    if err != nil {
        panic(err)
    }
}
// SetDeploymentProfileAssignedDateTime sets the deploymentProfileAssignedDateTime property value. Profile set time of the Windows autopilot device.
func (m *WindowsAutopilotDeviceIdentity) SetDeploymentProfileAssignedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("deploymentProfileAssignedDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetDeploymentProfileAssignmentDetailedStatus sets the deploymentProfileAssignmentDetailedStatus property value. The deploymentProfileAssignmentDetailedStatus property
func (m *WindowsAutopilotDeviceIdentity) SetDeploymentProfileAssignmentDetailedStatus(value *WindowsAutopilotProfileAssignmentDetailedStatus)() {
    err := m.GetBackingStore().Set("deploymentProfileAssignmentDetailedStatus", value)
    if err != nil {
        panic(err)
    }
}
// SetDeploymentProfileAssignmentStatus sets the deploymentProfileAssignmentStatus property value. The deploymentProfileAssignmentStatus property
func (m *WindowsAutopilotDeviceIdentity) SetDeploymentProfileAssignmentStatus(value *WindowsAutopilotProfileAssignmentStatus)() {
    err := m.GetBackingStore().Set("deploymentProfileAssignmentStatus", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceAccountPassword sets the deviceAccountPassword property value. Surface Hub Device Account Password
func (m *WindowsAutopilotDeviceIdentity) SetDeviceAccountPassword(value *string)() {
    err := m.GetBackingStore().Set("deviceAccountPassword", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceAccountUpn sets the deviceAccountUpn property value. Surface Hub Device Account Upn
func (m *WindowsAutopilotDeviceIdentity) SetDeviceAccountUpn(value *string)() {
    err := m.GetBackingStore().Set("deviceAccountUpn", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceFriendlyName sets the deviceFriendlyName property value. Surface Hub Device Friendly Name
func (m *WindowsAutopilotDeviceIdentity) SetDeviceFriendlyName(value *string)() {
    err := m.GetBackingStore().Set("deviceFriendlyName", value)
    if err != nil {
        panic(err)
    }
}
// SetDisplayName sets the displayName property value. Display Name
func (m *WindowsAutopilotDeviceIdentity) SetDisplayName(value *string)() {
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// SetEnrollmentState sets the enrollmentState property value. The enrollmentState property
func (m *WindowsAutopilotDeviceIdentity) SetEnrollmentState(value *EnrollmentState)() {
    err := m.GetBackingStore().Set("enrollmentState", value)
    if err != nil {
        panic(err)
    }
}
// SetGroupTag sets the groupTag property value. Group Tag of the Windows autopilot device.
func (m *WindowsAutopilotDeviceIdentity) SetGroupTag(value *string)() {
    err := m.GetBackingStore().Set("groupTag", value)
    if err != nil {
        panic(err)
    }
}
// SetIntendedDeploymentProfile sets the intendedDeploymentProfile property value. Deployment profile intended to be assigned to the Windows autopilot device.
func (m *WindowsAutopilotDeviceIdentity) SetIntendedDeploymentProfile(value WindowsAutopilotDeploymentProfileable)() {
    err := m.GetBackingStore().Set("intendedDeploymentProfile", value)
    if err != nil {
        panic(err)
    }
}
// SetLastContactedDateTime sets the lastContactedDateTime property value. Intune Last Contacted Date Time of the Windows autopilot device.
func (m *WindowsAutopilotDeviceIdentity) SetLastContactedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("lastContactedDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetManagedDeviceId sets the managedDeviceId property value. Managed Device ID
func (m *WindowsAutopilotDeviceIdentity) SetManagedDeviceId(value *string)() {
    err := m.GetBackingStore().Set("managedDeviceId", value)
    if err != nil {
        panic(err)
    }
}
// SetManufacturer sets the manufacturer property value. Oem manufacturer of the Windows autopilot device.
func (m *WindowsAutopilotDeviceIdentity) SetManufacturer(value *string)() {
    err := m.GetBackingStore().Set("manufacturer", value)
    if err != nil {
        panic(err)
    }
}
// SetModel sets the model property value. Model name of the Windows autopilot device.
func (m *WindowsAutopilotDeviceIdentity) SetModel(value *string)() {
    err := m.GetBackingStore().Set("model", value)
    if err != nil {
        panic(err)
    }
}
// SetProductKey sets the productKey property value. Product Key of the Windows autopilot device.
func (m *WindowsAutopilotDeviceIdentity) SetProductKey(value *string)() {
    err := m.GetBackingStore().Set("productKey", value)
    if err != nil {
        panic(err)
    }
}
// SetPurchaseOrderIdentifier sets the purchaseOrderIdentifier property value. Purchase Order Identifier of the Windows autopilot device.
func (m *WindowsAutopilotDeviceIdentity) SetPurchaseOrderIdentifier(value *string)() {
    err := m.GetBackingStore().Set("purchaseOrderIdentifier", value)
    if err != nil {
        panic(err)
    }
}
// SetRemediationState sets the remediationState property value. Device remediation status, indicating whether or not hardware has been changed for an Autopilot-registered device.
func (m *WindowsAutopilotDeviceIdentity) SetRemediationState(value *WindowsAutopilotDeviceRemediationState)() {
    err := m.GetBackingStore().Set("remediationState", value)
    if err != nil {
        panic(err)
    }
}
// SetRemediationStateLastModifiedDateTime sets the remediationStateLastModifiedDateTime property value. RemediationState set time of Autopilot device.
func (m *WindowsAutopilotDeviceIdentity) SetRemediationStateLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("remediationStateLastModifiedDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetResourceName sets the resourceName property value. Resource Name.
func (m *WindowsAutopilotDeviceIdentity) SetResourceName(value *string)() {
    err := m.GetBackingStore().Set("resourceName", value)
    if err != nil {
        panic(err)
    }
}
// SetSerialNumber sets the serialNumber property value. Serial number of the Windows autopilot device.
func (m *WindowsAutopilotDeviceIdentity) SetSerialNumber(value *string)() {
    err := m.GetBackingStore().Set("serialNumber", value)
    if err != nil {
        panic(err)
    }
}
// SetSkuNumber sets the skuNumber property value. SKU Number
func (m *WindowsAutopilotDeviceIdentity) SetSkuNumber(value *string)() {
    err := m.GetBackingStore().Set("skuNumber", value)
    if err != nil {
        panic(err)
    }
}
// SetSystemFamily sets the systemFamily property value. System Family
func (m *WindowsAutopilotDeviceIdentity) SetSystemFamily(value *string)() {
    err := m.GetBackingStore().Set("systemFamily", value)
    if err != nil {
        panic(err)
    }
}
// SetUserlessEnrollmentStatus sets the userlessEnrollmentStatus property value. Userless enrollment block status, indicating whether the next device enrollment will be blocked.
func (m *WindowsAutopilotDeviceIdentity) SetUserlessEnrollmentStatus(value *WindowsAutopilotUserlessEnrollmentStatus)() {
    err := m.GetBackingStore().Set("userlessEnrollmentStatus", value)
    if err != nil {
        panic(err)
    }
}
// SetUserPrincipalName sets the userPrincipalName property value. User Principal Name.
func (m *WindowsAutopilotDeviceIdentity) SetUserPrincipalName(value *string)() {
    err := m.GetBackingStore().Set("userPrincipalName", value)
    if err != nil {
        panic(err)
    }
}
type WindowsAutopilotDeviceIdentityable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAddressableUserName()(*string)
    GetAzureActiveDirectoryDeviceId()(*string)
    GetAzureAdDeviceId()(*string)
    GetDeploymentProfile()(WindowsAutopilotDeploymentProfileable)
    GetDeploymentProfileAssignedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDeploymentProfileAssignmentDetailedStatus()(*WindowsAutopilotProfileAssignmentDetailedStatus)
    GetDeploymentProfileAssignmentStatus()(*WindowsAutopilotProfileAssignmentStatus)
    GetDeviceAccountPassword()(*string)
    GetDeviceAccountUpn()(*string)
    GetDeviceFriendlyName()(*string)
    GetDisplayName()(*string)
    GetEnrollmentState()(*EnrollmentState)
    GetGroupTag()(*string)
    GetIntendedDeploymentProfile()(WindowsAutopilotDeploymentProfileable)
    GetLastContactedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetManagedDeviceId()(*string)
    GetManufacturer()(*string)
    GetModel()(*string)
    GetProductKey()(*string)
    GetPurchaseOrderIdentifier()(*string)
    GetRemediationState()(*WindowsAutopilotDeviceRemediationState)
    GetRemediationStateLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetResourceName()(*string)
    GetSerialNumber()(*string)
    GetSkuNumber()(*string)
    GetSystemFamily()(*string)
    GetUserlessEnrollmentStatus()(*WindowsAutopilotUserlessEnrollmentStatus)
    GetUserPrincipalName()(*string)
    SetAddressableUserName(value *string)()
    SetAzureActiveDirectoryDeviceId(value *string)()
    SetAzureAdDeviceId(value *string)()
    SetDeploymentProfile(value WindowsAutopilotDeploymentProfileable)()
    SetDeploymentProfileAssignedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDeploymentProfileAssignmentDetailedStatus(value *WindowsAutopilotProfileAssignmentDetailedStatus)()
    SetDeploymentProfileAssignmentStatus(value *WindowsAutopilotProfileAssignmentStatus)()
    SetDeviceAccountPassword(value *string)()
    SetDeviceAccountUpn(value *string)()
    SetDeviceFriendlyName(value *string)()
    SetDisplayName(value *string)()
    SetEnrollmentState(value *EnrollmentState)()
    SetGroupTag(value *string)()
    SetIntendedDeploymentProfile(value WindowsAutopilotDeploymentProfileable)()
    SetLastContactedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetManagedDeviceId(value *string)()
    SetManufacturer(value *string)()
    SetModel(value *string)()
    SetProductKey(value *string)()
    SetPurchaseOrderIdentifier(value *string)()
    SetRemediationState(value *WindowsAutopilotDeviceRemediationState)()
    SetRemediationStateLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetResourceName(value *string)()
    SetSerialNumber(value *string)()
    SetSkuNumber(value *string)()
    SetSystemFamily(value *string)()
    SetUserlessEnrollmentStatus(value *WindowsAutopilotUserlessEnrollmentStatus)()
    SetUserPrincipalName(value *string)()
}
