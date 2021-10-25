package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type WindowsAutopilotDeviceIdentity struct {
    Entity
    addressableUserName *string;
    azureActiveDirectoryDeviceId *string;
    azureAdDeviceId *string;
    deploymentProfile *WindowsAutopilotDeploymentProfile;
    deploymentProfileAssignedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    deploymentProfileAssignmentDetailedStatus *WindowsAutopilotProfileAssignmentDetailedStatus;
    deploymentProfileAssignmentStatus *WindowsAutopilotProfileAssignmentStatus;
    displayName *string;
    enrollmentState *EnrollmentState;
    groupTag *string;
    intendedDeploymentProfile *WindowsAutopilotDeploymentProfile;
    lastContactedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    managedDeviceId *string;
    manufacturer *string;
    model *string;
    productKey *string;
    purchaseOrderIdentifier *string;
    resourceName *string;
    serialNumber *string;
    skuNumber *string;
    systemFamily *string;
    userPrincipalName *string;
}
func NewWindowsAutopilotDeviceIdentity()(*WindowsAutopilotDeviceIdentity) {
    m := &WindowsAutopilotDeviceIdentity{
        Entity: *NewEntity(),
    }
    return m
}
func (m *WindowsAutopilotDeviceIdentity) GetAddressableUserName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.addressableUserName
    }
}
func (m *WindowsAutopilotDeviceIdentity) GetAzureActiveDirectoryDeviceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.azureActiveDirectoryDeviceId
    }
}
func (m *WindowsAutopilotDeviceIdentity) GetAzureAdDeviceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.azureAdDeviceId
    }
}
func (m *WindowsAutopilotDeviceIdentity) GetDeploymentProfile()(*WindowsAutopilotDeploymentProfile) {
    if m == nil {
        return nil
    } else {
        return m.deploymentProfile
    }
}
func (m *WindowsAutopilotDeviceIdentity) GetDeploymentProfileAssignedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.deploymentProfileAssignedDateTime
    }
}
func (m *WindowsAutopilotDeviceIdentity) GetDeploymentProfileAssignmentDetailedStatus()(*WindowsAutopilotProfileAssignmentDetailedStatus) {
    if m == nil {
        return nil
    } else {
        return m.deploymentProfileAssignmentDetailedStatus
    }
}
func (m *WindowsAutopilotDeviceIdentity) GetDeploymentProfileAssignmentStatus()(*WindowsAutopilotProfileAssignmentStatus) {
    if m == nil {
        return nil
    } else {
        return m.deploymentProfileAssignmentStatus
    }
}
func (m *WindowsAutopilotDeviceIdentity) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
func (m *WindowsAutopilotDeviceIdentity) GetEnrollmentState()(*EnrollmentState) {
    if m == nil {
        return nil
    } else {
        return m.enrollmentState
    }
}
func (m *WindowsAutopilotDeviceIdentity) GetGroupTag()(*string) {
    if m == nil {
        return nil
    } else {
        return m.groupTag
    }
}
func (m *WindowsAutopilotDeviceIdentity) GetIntendedDeploymentProfile()(*WindowsAutopilotDeploymentProfile) {
    if m == nil {
        return nil
    } else {
        return m.intendedDeploymentProfile
    }
}
func (m *WindowsAutopilotDeviceIdentity) GetLastContactedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastContactedDateTime
    }
}
func (m *WindowsAutopilotDeviceIdentity) GetManagedDeviceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.managedDeviceId
    }
}
func (m *WindowsAutopilotDeviceIdentity) GetManufacturer()(*string) {
    if m == nil {
        return nil
    } else {
        return m.manufacturer
    }
}
func (m *WindowsAutopilotDeviceIdentity) GetModel()(*string) {
    if m == nil {
        return nil
    } else {
        return m.model
    }
}
func (m *WindowsAutopilotDeviceIdentity) GetProductKey()(*string) {
    if m == nil {
        return nil
    } else {
        return m.productKey
    }
}
func (m *WindowsAutopilotDeviceIdentity) GetPurchaseOrderIdentifier()(*string) {
    if m == nil {
        return nil
    } else {
        return m.purchaseOrderIdentifier
    }
}
func (m *WindowsAutopilotDeviceIdentity) GetResourceName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.resourceName
    }
}
func (m *WindowsAutopilotDeviceIdentity) GetSerialNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.serialNumber
    }
}
func (m *WindowsAutopilotDeviceIdentity) GetSkuNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.skuNumber
    }
}
func (m *WindowsAutopilotDeviceIdentity) GetSystemFamily()(*string) {
    if m == nil {
        return nil
    } else {
        return m.systemFamily
    }
}
func (m *WindowsAutopilotDeviceIdentity) GetUserPrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userPrincipalName
    }
}
func (m *WindowsAutopilotDeviceIdentity) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["addressableUserName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAddressableUserName(val)
        return nil
    }
    res["azureActiveDirectoryDeviceId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAzureActiveDirectoryDeviceId(val)
        return nil
    }
    res["azureAdDeviceId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAzureAdDeviceId(val)
        return nil
    }
    res["deploymentProfile"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWindowsAutopilotDeploymentProfile() })
        if err != nil {
            return err
        }
        m.SetDeploymentProfile(val.(*WindowsAutopilotDeploymentProfile))
        return nil
    }
    res["deploymentProfileAssignedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetDeploymentProfileAssignedDateTime(val)
        return nil
    }
    res["deploymentProfileAssignmentDetailedStatus"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseWindowsAutopilotProfileAssignmentDetailedStatus)
        if err != nil {
            return err
        }
        cast := val.(WindowsAutopilotProfileAssignmentDetailedStatus)
        m.SetDeploymentProfileAssignmentDetailedStatus(&cast)
        return nil
    }
    res["deploymentProfileAssignmentStatus"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseWindowsAutopilotProfileAssignmentStatus)
        if err != nil {
            return err
        }
        cast := val.(WindowsAutopilotProfileAssignmentStatus)
        m.SetDeploymentProfileAssignmentStatus(&cast)
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDisplayName(val)
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
    res["groupTag"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetGroupTag(val)
        return nil
    }
    res["intendedDeploymentProfile"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWindowsAutopilotDeploymentProfile() })
        if err != nil {
            return err
        }
        m.SetIntendedDeploymentProfile(val.(*WindowsAutopilotDeploymentProfile))
        return nil
    }
    res["lastContactedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetLastContactedDateTime(val)
        return nil
    }
    res["managedDeviceId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetManagedDeviceId(val)
        return nil
    }
    res["manufacturer"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetManufacturer(val)
        return nil
    }
    res["model"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetModel(val)
        return nil
    }
    res["productKey"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetProductKey(val)
        return nil
    }
    res["purchaseOrderIdentifier"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetPurchaseOrderIdentifier(val)
        return nil
    }
    res["resourceName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetResourceName(val)
        return nil
    }
    res["serialNumber"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetSerialNumber(val)
        return nil
    }
    res["skuNumber"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetSkuNumber(val)
        return nil
    }
    res["systemFamily"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetSystemFamily(val)
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
func (m *WindowsAutopilotDeviceIdentity) IsNil()(bool) {
    return m == nil
}
func (m *WindowsAutopilotDeviceIdentity) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
        cast := m.GetDeploymentProfileAssignmentDetailedStatus().String()
        err = writer.WriteStringValue("deploymentProfileAssignmentDetailedStatus", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetDeploymentProfileAssignmentStatus() != nil {
        cast := m.GetDeploymentProfileAssignmentStatus().String()
        err = writer.WriteStringValue("deploymentProfileAssignmentStatus", &cast)
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
        cast := m.GetEnrollmentState().String()
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
    {
        err = writer.WriteStringValue("userPrincipalName", m.GetUserPrincipalName())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *WindowsAutopilotDeviceIdentity) SetAddressableUserName(value *string)() {
    m.addressableUserName = value
}
func (m *WindowsAutopilotDeviceIdentity) SetAzureActiveDirectoryDeviceId(value *string)() {
    m.azureActiveDirectoryDeviceId = value
}
func (m *WindowsAutopilotDeviceIdentity) SetAzureAdDeviceId(value *string)() {
    m.azureAdDeviceId = value
}
func (m *WindowsAutopilotDeviceIdentity) SetDeploymentProfile(value *WindowsAutopilotDeploymentProfile)() {
    m.deploymentProfile = value
}
func (m *WindowsAutopilotDeviceIdentity) SetDeploymentProfileAssignedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.deploymentProfileAssignedDateTime = value
}
func (m *WindowsAutopilotDeviceIdentity) SetDeploymentProfileAssignmentDetailedStatus(value *WindowsAutopilotProfileAssignmentDetailedStatus)() {
    m.deploymentProfileAssignmentDetailedStatus = value
}
func (m *WindowsAutopilotDeviceIdentity) SetDeploymentProfileAssignmentStatus(value *WindowsAutopilotProfileAssignmentStatus)() {
    m.deploymentProfileAssignmentStatus = value
}
func (m *WindowsAutopilotDeviceIdentity) SetDisplayName(value *string)() {
    m.displayName = value
}
func (m *WindowsAutopilotDeviceIdentity) SetEnrollmentState(value *EnrollmentState)() {
    m.enrollmentState = value
}
func (m *WindowsAutopilotDeviceIdentity) SetGroupTag(value *string)() {
    m.groupTag = value
}
func (m *WindowsAutopilotDeviceIdentity) SetIntendedDeploymentProfile(value *WindowsAutopilotDeploymentProfile)() {
    m.intendedDeploymentProfile = value
}
func (m *WindowsAutopilotDeviceIdentity) SetLastContactedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastContactedDateTime = value
}
func (m *WindowsAutopilotDeviceIdentity) SetManagedDeviceId(value *string)() {
    m.managedDeviceId = value
}
func (m *WindowsAutopilotDeviceIdentity) SetManufacturer(value *string)() {
    m.manufacturer = value
}
func (m *WindowsAutopilotDeviceIdentity) SetModel(value *string)() {
    m.model = value
}
func (m *WindowsAutopilotDeviceIdentity) SetProductKey(value *string)() {
    m.productKey = value
}
func (m *WindowsAutopilotDeviceIdentity) SetPurchaseOrderIdentifier(value *string)() {
    m.purchaseOrderIdentifier = value
}
func (m *WindowsAutopilotDeviceIdentity) SetResourceName(value *string)() {
    m.resourceName = value
}
func (m *WindowsAutopilotDeviceIdentity) SetSerialNumber(value *string)() {
    m.serialNumber = value
}
func (m *WindowsAutopilotDeviceIdentity) SetSkuNumber(value *string)() {
    m.skuNumber = value
}
func (m *WindowsAutopilotDeviceIdentity) SetSystemFamily(value *string)() {
    m.systemFamily = value
}
func (m *WindowsAutopilotDeviceIdentity) SetUserPrincipalName(value *string)() {
    m.userPrincipalName = value
}
