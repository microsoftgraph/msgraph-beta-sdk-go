package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Device casts the previous resource to device.
type Device struct {
    DirectoryObject
    // true if the account is enabled; otherwise, false. Default is true.  Supports $filter (eq, ne, not, in). Only callers in Global Administrator and Cloud Device Administrator roles can set this property.
    accountEnabled *bool
    // For internal use only. Not nullable. Supports $filter (eq, not, ge, le).
    alternativeSecurityIds []AlternativeSecurityIdable
    // The timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only. Supports $filter (eq, ne, not, ge, le, and eq on null values) and $orderBy.
    approximateLastSignInDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Set of commands sent to this device.
    commands []Commandable
    // The timestamp when the device is no longer deemed compliant. The timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
    complianceExpirationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // User-defined property set by Intune to automatically add devices to groups and simplify managing devices.
    deviceCategory *string
    // Identifier set by Azure Device Registration Service at the time of registration. Supports $filter (eq, ne, not, startsWith).
    deviceId *string
    // For internal use only. Set to null.
    deviceMetadata *string
    // Ownership of the device. This property is set by Intune. Possible values are: unknown, company, personal.
    deviceOwnership *string
    // For internal use only.
    deviceVersion *int32
    // The display name for the device. Required. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values), $search, and $orderBy.
    displayName *string
    // The on-premises domain name of Hybrid Azure AD joined devices. This property is set by Intune.
    domainName *string
    // Enrollment profile applied to the device. For example, Apple Device Enrollment Profile, Device enrollment - Corporate device identifiers, or Windows Autopilot profile name. This property is set by Intune.
    enrollmentProfileName *string
    // Enrollment type of the device. This property is set by Intune. Possible values are: unknown, userEnrollment, deviceEnrollmentManager, appleBulkWithUser, appleBulkWithoutUser, windowsAzureADJoin, windowsBulkUserless, windowsAutoEnrollment, windowsBulkAzureDomainJoin, windowsCoManagement.
    enrollmentType *string
    // Contains extension attributes 1-15 for the device. The individual extension attributes are not selectable. These properties are mastered in cloud and can be set during creation or update of a device object in Azure AD. Supports $filter (eq, not, startsWith, and eq on null values).
    extensionAttributes OnPremisesExtensionAttributesable
    // The collection of open extensions defined for the device. Read-only. Nullable.
    extensions []Extensionable
    // List of hostNames for the device.
    hostnames []string
    // true if the device complies with Mobile Device Management (MDM) policies; otherwise, false. Read-only. This can only be updated by Intune for any device OS type or by an approved MDM app for Windows OS devices. Supports $filter (eq, ne, not).
    isCompliant *bool
    // true if the device is managed by a Mobile Device Management (MDM) app; otherwise, false. This can only be updated by Intune for any device OS type or by an approved MDM app for Windows OS devices. Supports $filter (eq, ne, not).
    isManaged *bool
    // The isManagementRestricted property
    isManagementRestricted *bool
    // true if device is rooted; false if device is jail-broken. This can only be updated by Intune.
    isRooted *bool
    // Form factor of device. Only returned if user signs in with a Microsoft account as part of Project Rome.
    kind *string
    // Management channel of the device.  This property is set by Intune. Possible values are: eas, mdm, easMdm, intuneClient, easIntuneClient, configurationManagerClient, configurationManagerClientMdm, configurationManagerClientMdmEas, unknown, jamf, googleCloudDevicePolicyController.
    managementType *string
    // Manufacturer of the device. Read-only.
    manufacturer *string
    // Application identifier used to register device into MDM. Read-only. Supports $filter (eq, ne, not, startsWith).
    mdmAppId *string
    // Groups and administrative units that this device is a member of. Read-only. Nullable. Supports $expand.
    memberOf []DirectoryObjectable
    // Model of the device. Read-only.
    model *string
    // Friendly name of a device. Only returned if user signs in with a Microsoft account as part of Project Rome.
    name *string
    // The last time at which the object was synced with the on-premises directory. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z Read-only. Supports $filter (eq, ne, not, ge, le, in).
    onPremisesLastSyncDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // true if this object is synced from an on-premises directory; false if this object was originally synced from an on-premises directory but is no longer synced; null if this object has never been synced from an on-premises directory (default). Read-only. Supports $filter (eq, ne, not, in, and eq on null values).
    onPremisesSyncEnabled *bool
    // The type of operating system on the device. Required. Supports $filter (eq, ne, not, ge, le, startsWith, and eq on null values).
    operatingSystem *string
    // Operating system version of the device. Required. Supports $filter (eq, ne, not, ge, le, startsWith, and eq on null values).
    operatingSystemVersion *string
    // For internal use only. Not nullable. Supports $filter (eq, not, ge, le, startsWith, and counting empty collections).
    physicalIds []string
    // Platform of device. Only returned if user signs in with a Microsoft account as part of Project Rome. Only returned if user signs in with a Microsoft account as part of Project Rome.
    platform *string
    // The profile type of the device. Possible values: RegisteredDevice (default), SecureVM, Printer, Shared, IoT.
    profileType *string
    // The user that cloud joined the device or registered their personal device. The registered owner is set at the time of registration. Currently, there can be only one owner. Read-only. Nullable. Supports $expand.
    registeredOwners []DirectoryObjectable
    // Collection of registered users of the device. For cloud joined devices and registered personal devices, registered users are set to the same value as registered owners at the time of registration. Read-only. Nullable. Supports $expand.
    registeredUsers []DirectoryObjectable
    // Date and time of when the device was registered. The timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
    registrationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Device is online or offline. Only returned if user signs in with a Microsoft account as part of Project Rome.
    status *string
    // List of labels applied to the device by the system. Supports $filter (eq when counting empty collections).
    systemLabels []string
    // Groups and administrative units that this device is a member of. This operation is transitive. Supports $expand.
    transitiveMemberOf []DirectoryObjectable
    // Type of trust for the joined device. Read-only. Possible values: Workplace (indicates bring your own personal devices), AzureAd (Cloud only joined devices), ServerAd (on-premises domain joined devices joined to Azure AD). For more details, see Introduction to device management in Azure Active Directory
    trustType *string
    // Represents the usage rights a device has been granted.
    usageRights []UsageRightable
}
// NewDevice instantiates a new device and sets the default values.
func NewDevice()(*Device) {
    m := &Device{
        DirectoryObject: *NewDirectoryObject(),
    }
    odataTypeValue := "#microsoft.graph.device";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateDeviceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDevice(), nil
}
// GetAccountEnabled gets the accountEnabled property value. true if the account is enabled; otherwise, false. Default is true.  Supports $filter (eq, ne, not, in). Only callers in Global Administrator and Cloud Device Administrator roles can set this property.
func (m *Device) GetAccountEnabled()(*bool) {
    return m.accountEnabled
}
// GetAlternativeSecurityIds gets the alternativeSecurityIds property value. For internal use only. Not nullable. Supports $filter (eq, not, ge, le).
func (m *Device) GetAlternativeSecurityIds()([]AlternativeSecurityIdable) {
    return m.alternativeSecurityIds
}
// GetApproximateLastSignInDateTime gets the approximateLastSignInDateTime property value. The timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only. Supports $filter (eq, ne, not, ge, le, and eq on null values) and $orderBy.
func (m *Device) GetApproximateLastSignInDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.approximateLastSignInDateTime
}
// GetCommands gets the commands property value. Set of commands sent to this device.
func (m *Device) GetCommands()([]Commandable) {
    return m.commands
}
// GetComplianceExpirationDateTime gets the complianceExpirationDateTime property value. The timestamp when the device is no longer deemed compliant. The timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
func (m *Device) GetComplianceExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.complianceExpirationDateTime
}
// GetDeviceCategory gets the deviceCategory property value. User-defined property set by Intune to automatically add devices to groups and simplify managing devices.
func (m *Device) GetDeviceCategory()(*string) {
    return m.deviceCategory
}
// GetDeviceId gets the deviceId property value. Identifier set by Azure Device Registration Service at the time of registration. Supports $filter (eq, ne, not, startsWith).
func (m *Device) GetDeviceId()(*string) {
    return m.deviceId
}
// GetDeviceMetadata gets the deviceMetadata property value. For internal use only. Set to null.
func (m *Device) GetDeviceMetadata()(*string) {
    return m.deviceMetadata
}
// GetDeviceOwnership gets the deviceOwnership property value. Ownership of the device. This property is set by Intune. Possible values are: unknown, company, personal.
func (m *Device) GetDeviceOwnership()(*string) {
    return m.deviceOwnership
}
// GetDeviceVersion gets the deviceVersion property value. For internal use only.
func (m *Device) GetDeviceVersion()(*int32) {
    return m.deviceVersion
}
// GetDisplayName gets the displayName property value. The display name for the device. Required. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values), $search, and $orderBy.
func (m *Device) GetDisplayName()(*string) {
    return m.displayName
}
// GetDomainName gets the domainName property value. The on-premises domain name of Hybrid Azure AD joined devices. This property is set by Intune.
func (m *Device) GetDomainName()(*string) {
    return m.domainName
}
// GetEnrollmentProfileName gets the enrollmentProfileName property value. Enrollment profile applied to the device. For example, Apple Device Enrollment Profile, Device enrollment - Corporate device identifiers, or Windows Autopilot profile name. This property is set by Intune.
func (m *Device) GetEnrollmentProfileName()(*string) {
    return m.enrollmentProfileName
}
// GetEnrollmentType gets the enrollmentType property value. Enrollment type of the device. This property is set by Intune. Possible values are: unknown, userEnrollment, deviceEnrollmentManager, appleBulkWithUser, appleBulkWithoutUser, windowsAzureADJoin, windowsBulkUserless, windowsAutoEnrollment, windowsBulkAzureDomainJoin, windowsCoManagement.
func (m *Device) GetEnrollmentType()(*string) {
    return m.enrollmentType
}
// GetExtensionAttributes gets the extensionAttributes property value. Contains extension attributes 1-15 for the device. The individual extension attributes are not selectable. These properties are mastered in cloud and can be set during creation or update of a device object in Azure AD. Supports $filter (eq, not, startsWith, and eq on null values).
func (m *Device) GetExtensionAttributes()(OnPremisesExtensionAttributesable) {
    return m.extensionAttributes
}
// GetExtensions gets the extensions property value. The collection of open extensions defined for the device. Read-only. Nullable.
func (m *Device) GetExtensions()([]Extensionable) {
    return m.extensions
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Device) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DirectoryObject.GetFieldDeserializers()
    res["accountEnabled"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetAccountEnabled)
    res["alternativeSecurityIds"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateAlternativeSecurityIdFromDiscriminatorValue , m.SetAlternativeSecurityIds)
    res["approximateLastSignInDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetApproximateLastSignInDateTime)
    res["commands"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateCommandFromDiscriminatorValue , m.SetCommands)
    res["complianceExpirationDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetComplianceExpirationDateTime)
    res["deviceCategory"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDeviceCategory)
    res["deviceId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDeviceId)
    res["deviceMetadata"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDeviceMetadata)
    res["deviceOwnership"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDeviceOwnership)
    res["deviceVersion"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetDeviceVersion)
    res["displayName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDisplayName)
    res["domainName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDomainName)
    res["enrollmentProfileName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetEnrollmentProfileName)
    res["enrollmentType"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetEnrollmentType)
    res["extensionAttributes"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateOnPremisesExtensionAttributesFromDiscriminatorValue , m.SetExtensionAttributes)
    res["extensions"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateExtensionFromDiscriminatorValue , m.SetExtensions)
    res["hostnames"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetHostnames)
    res["isCompliant"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIsCompliant)
    res["isManaged"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIsManaged)
    res["isManagementRestricted"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIsManagementRestricted)
    res["isRooted"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIsRooted)
    res["kind"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetKind)
    res["managementType"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetManagementType)
    res["manufacturer"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetManufacturer)
    res["mdmAppId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetMdmAppId)
    res["memberOf"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateDirectoryObjectFromDiscriminatorValue , m.SetMemberOf)
    res["model"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetModel)
    res["name"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetName)
    res["onPremisesLastSyncDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetOnPremisesLastSyncDateTime)
    res["onPremisesSyncEnabled"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetOnPremisesSyncEnabled)
    res["operatingSystem"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOperatingSystem)
    res["operatingSystemVersion"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOperatingSystemVersion)
    res["physicalIds"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetPhysicalIds)
    res["platform"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetPlatform)
    res["profileType"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetProfileType)
    res["registeredOwners"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateDirectoryObjectFromDiscriminatorValue , m.SetRegisteredOwners)
    res["registeredUsers"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateDirectoryObjectFromDiscriminatorValue , m.SetRegisteredUsers)
    res["registrationDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetRegistrationDateTime)
    res["status"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetStatus)
    res["systemLabels"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetSystemLabels)
    res["transitiveMemberOf"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateDirectoryObjectFromDiscriminatorValue , m.SetTransitiveMemberOf)
    res["trustType"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetTrustType)
    res["usageRights"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateUsageRightFromDiscriminatorValue , m.SetUsageRights)
    return res
}
// GetHostnames gets the hostnames property value. List of hostNames for the device.
func (m *Device) GetHostnames()([]string) {
    return m.hostnames
}
// GetIsCompliant gets the isCompliant property value. true if the device complies with Mobile Device Management (MDM) policies; otherwise, false. Read-only. This can only be updated by Intune for any device OS type or by an approved MDM app for Windows OS devices. Supports $filter (eq, ne, not).
func (m *Device) GetIsCompliant()(*bool) {
    return m.isCompliant
}
// GetIsManaged gets the isManaged property value. true if the device is managed by a Mobile Device Management (MDM) app; otherwise, false. This can only be updated by Intune for any device OS type or by an approved MDM app for Windows OS devices. Supports $filter (eq, ne, not).
func (m *Device) GetIsManaged()(*bool) {
    return m.isManaged
}
// GetIsManagementRestricted gets the isManagementRestricted property value. The isManagementRestricted property
func (m *Device) GetIsManagementRestricted()(*bool) {
    return m.isManagementRestricted
}
// GetIsRooted gets the isRooted property value. true if device is rooted; false if device is jail-broken. This can only be updated by Intune.
func (m *Device) GetIsRooted()(*bool) {
    return m.isRooted
}
// GetKind gets the kind property value. Form factor of device. Only returned if user signs in with a Microsoft account as part of Project Rome.
func (m *Device) GetKind()(*string) {
    return m.kind
}
// GetManagementType gets the managementType property value. Management channel of the device.  This property is set by Intune. Possible values are: eas, mdm, easMdm, intuneClient, easIntuneClient, configurationManagerClient, configurationManagerClientMdm, configurationManagerClientMdmEas, unknown, jamf, googleCloudDevicePolicyController.
func (m *Device) GetManagementType()(*string) {
    return m.managementType
}
// GetManufacturer gets the manufacturer property value. Manufacturer of the device. Read-only.
func (m *Device) GetManufacturer()(*string) {
    return m.manufacturer
}
// GetMdmAppId gets the mdmAppId property value. Application identifier used to register device into MDM. Read-only. Supports $filter (eq, ne, not, startsWith).
func (m *Device) GetMdmAppId()(*string) {
    return m.mdmAppId
}
// GetMemberOf gets the memberOf property value. Groups and administrative units that this device is a member of. Read-only. Nullable. Supports $expand.
func (m *Device) GetMemberOf()([]DirectoryObjectable) {
    return m.memberOf
}
// GetModel gets the model property value. Model of the device. Read-only.
func (m *Device) GetModel()(*string) {
    return m.model
}
// GetName gets the name property value. Friendly name of a device. Only returned if user signs in with a Microsoft account as part of Project Rome.
func (m *Device) GetName()(*string) {
    return m.name
}
// GetOnPremisesLastSyncDateTime gets the onPremisesLastSyncDateTime property value. The last time at which the object was synced with the on-premises directory. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z Read-only. Supports $filter (eq, ne, not, ge, le, in).
func (m *Device) GetOnPremisesLastSyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.onPremisesLastSyncDateTime
}
// GetOnPremisesSyncEnabled gets the onPremisesSyncEnabled property value. true if this object is synced from an on-premises directory; false if this object was originally synced from an on-premises directory but is no longer synced; null if this object has never been synced from an on-premises directory (default). Read-only. Supports $filter (eq, ne, not, in, and eq on null values).
func (m *Device) GetOnPremisesSyncEnabled()(*bool) {
    return m.onPremisesSyncEnabled
}
// GetOperatingSystem gets the operatingSystem property value. The type of operating system on the device. Required. Supports $filter (eq, ne, not, ge, le, startsWith, and eq on null values).
func (m *Device) GetOperatingSystem()(*string) {
    return m.operatingSystem
}
// GetOperatingSystemVersion gets the operatingSystemVersion property value. Operating system version of the device. Required. Supports $filter (eq, ne, not, ge, le, startsWith, and eq on null values).
func (m *Device) GetOperatingSystemVersion()(*string) {
    return m.operatingSystemVersion
}
// GetPhysicalIds gets the physicalIds property value. For internal use only. Not nullable. Supports $filter (eq, not, ge, le, startsWith, and counting empty collections).
func (m *Device) GetPhysicalIds()([]string) {
    return m.physicalIds
}
// GetPlatform gets the platform property value. Platform of device. Only returned if user signs in with a Microsoft account as part of Project Rome. Only returned if user signs in with a Microsoft account as part of Project Rome.
func (m *Device) GetPlatform()(*string) {
    return m.platform
}
// GetProfileType gets the profileType property value. The profile type of the device. Possible values: RegisteredDevice (default), SecureVM, Printer, Shared, IoT.
func (m *Device) GetProfileType()(*string) {
    return m.profileType
}
// GetRegisteredOwners gets the registeredOwners property value. The user that cloud joined the device or registered their personal device. The registered owner is set at the time of registration. Currently, there can be only one owner. Read-only. Nullable. Supports $expand.
func (m *Device) GetRegisteredOwners()([]DirectoryObjectable) {
    return m.registeredOwners
}
// GetRegisteredUsers gets the registeredUsers property value. Collection of registered users of the device. For cloud joined devices and registered personal devices, registered users are set to the same value as registered owners at the time of registration. Read-only. Nullable. Supports $expand.
func (m *Device) GetRegisteredUsers()([]DirectoryObjectable) {
    return m.registeredUsers
}
// GetRegistrationDateTime gets the registrationDateTime property value. Date and time of when the device was registered. The timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
func (m *Device) GetRegistrationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.registrationDateTime
}
// GetStatus gets the status property value. Device is online or offline. Only returned if user signs in with a Microsoft account as part of Project Rome.
func (m *Device) GetStatus()(*string) {
    return m.status
}
// GetSystemLabels gets the systemLabels property value. List of labels applied to the device by the system. Supports $filter (eq when counting empty collections).
func (m *Device) GetSystemLabels()([]string) {
    return m.systemLabels
}
// GetTransitiveMemberOf gets the transitiveMemberOf property value. Groups and administrative units that this device is a member of. This operation is transitive. Supports $expand.
func (m *Device) GetTransitiveMemberOf()([]DirectoryObjectable) {
    return m.transitiveMemberOf
}
// GetTrustType gets the trustType property value. Type of trust for the joined device. Read-only. Possible values: Workplace (indicates bring your own personal devices), AzureAd (Cloud only joined devices), ServerAd (on-premises domain joined devices joined to Azure AD). For more details, see Introduction to device management in Azure Active Directory
func (m *Device) GetTrustType()(*string) {
    return m.trustType
}
// GetUsageRights gets the usageRights property value. Represents the usage rights a device has been granted.
func (m *Device) GetUsageRights()([]UsageRightable) {
    return m.usageRights
}
// Serialize serializes information the current object
func (m *Device) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DirectoryObject.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("accountEnabled", m.GetAccountEnabled())
        if err != nil {
            return err
        }
    }
    if m.GetAlternativeSecurityIds() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetAlternativeSecurityIds())
        err = writer.WriteCollectionOfObjectValues("alternativeSecurityIds", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("approximateLastSignInDateTime", m.GetApproximateLastSignInDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetCommands() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetCommands())
        err = writer.WriteCollectionOfObjectValues("commands", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("complianceExpirationDateTime", m.GetComplianceExpirationDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deviceCategory", m.GetDeviceCategory())
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
        err = writer.WriteStringValue("deviceMetadata", m.GetDeviceMetadata())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deviceOwnership", m.GetDeviceOwnership())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("deviceVersion", m.GetDeviceVersion())
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
    {
        err = writer.WriteStringValue("domainName", m.GetDomainName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("enrollmentProfileName", m.GetEnrollmentProfileName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("enrollmentType", m.GetEnrollmentType())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("extensionAttributes", m.GetExtensionAttributes())
        if err != nil {
            return err
        }
    }
    if m.GetExtensions() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetExtensions())
        err = writer.WriteCollectionOfObjectValues("extensions", cast)
        if err != nil {
            return err
        }
    }
    if m.GetHostnames() != nil {
        err = writer.WriteCollectionOfStringValues("hostnames", m.GetHostnames())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isCompliant", m.GetIsCompliant())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isManaged", m.GetIsManaged())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isManagementRestricted", m.GetIsManagementRestricted())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isRooted", m.GetIsRooted())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("kind", m.GetKind())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("managementType", m.GetManagementType())
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
        err = writer.WriteStringValue("mdmAppId", m.GetMdmAppId())
        if err != nil {
            return err
        }
    }
    if m.GetMemberOf() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetMemberOf())
        err = writer.WriteCollectionOfObjectValues("memberOf", cast)
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
        err = writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("onPremisesLastSyncDateTime", m.GetOnPremisesLastSyncDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("onPremisesSyncEnabled", m.GetOnPremisesSyncEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("operatingSystem", m.GetOperatingSystem())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("operatingSystemVersion", m.GetOperatingSystemVersion())
        if err != nil {
            return err
        }
    }
    if m.GetPhysicalIds() != nil {
        err = writer.WriteCollectionOfStringValues("physicalIds", m.GetPhysicalIds())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("platform", m.GetPlatform())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("profileType", m.GetProfileType())
        if err != nil {
            return err
        }
    }
    if m.GetRegisteredOwners() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetRegisteredOwners())
        err = writer.WriteCollectionOfObjectValues("registeredOwners", cast)
        if err != nil {
            return err
        }
    }
    if m.GetRegisteredUsers() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetRegisteredUsers())
        err = writer.WriteCollectionOfObjectValues("registeredUsers", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("registrationDateTime", m.GetRegistrationDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("status", m.GetStatus())
        if err != nil {
            return err
        }
    }
    if m.GetSystemLabels() != nil {
        err = writer.WriteCollectionOfStringValues("systemLabels", m.GetSystemLabels())
        if err != nil {
            return err
        }
    }
    if m.GetTransitiveMemberOf() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetTransitiveMemberOf())
        err = writer.WriteCollectionOfObjectValues("transitiveMemberOf", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("trustType", m.GetTrustType())
        if err != nil {
            return err
        }
    }
    if m.GetUsageRights() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetUsageRights())
        err = writer.WriteCollectionOfObjectValues("usageRights", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAccountEnabled sets the accountEnabled property value. true if the account is enabled; otherwise, false. Default is true.  Supports $filter (eq, ne, not, in). Only callers in Global Administrator and Cloud Device Administrator roles can set this property.
func (m *Device) SetAccountEnabled(value *bool)() {
    m.accountEnabled = value
}
// SetAlternativeSecurityIds sets the alternativeSecurityIds property value. For internal use only. Not nullable. Supports $filter (eq, not, ge, le).
func (m *Device) SetAlternativeSecurityIds(value []AlternativeSecurityIdable)() {
    m.alternativeSecurityIds = value
}
// SetApproximateLastSignInDateTime sets the approximateLastSignInDateTime property value. The timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only. Supports $filter (eq, ne, not, ge, le, and eq on null values) and $orderBy.
func (m *Device) SetApproximateLastSignInDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.approximateLastSignInDateTime = value
}
// SetCommands sets the commands property value. Set of commands sent to this device.
func (m *Device) SetCommands(value []Commandable)() {
    m.commands = value
}
// SetComplianceExpirationDateTime sets the complianceExpirationDateTime property value. The timestamp when the device is no longer deemed compliant. The timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
func (m *Device) SetComplianceExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.complianceExpirationDateTime = value
}
// SetDeviceCategory sets the deviceCategory property value. User-defined property set by Intune to automatically add devices to groups and simplify managing devices.
func (m *Device) SetDeviceCategory(value *string)() {
    m.deviceCategory = value
}
// SetDeviceId sets the deviceId property value. Identifier set by Azure Device Registration Service at the time of registration. Supports $filter (eq, ne, not, startsWith).
func (m *Device) SetDeviceId(value *string)() {
    m.deviceId = value
}
// SetDeviceMetadata sets the deviceMetadata property value. For internal use only. Set to null.
func (m *Device) SetDeviceMetadata(value *string)() {
    m.deviceMetadata = value
}
// SetDeviceOwnership sets the deviceOwnership property value. Ownership of the device. This property is set by Intune. Possible values are: unknown, company, personal.
func (m *Device) SetDeviceOwnership(value *string)() {
    m.deviceOwnership = value
}
// SetDeviceVersion sets the deviceVersion property value. For internal use only.
func (m *Device) SetDeviceVersion(value *int32)() {
    m.deviceVersion = value
}
// SetDisplayName sets the displayName property value. The display name for the device. Required. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values), $search, and $orderBy.
func (m *Device) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetDomainName sets the domainName property value. The on-premises domain name of Hybrid Azure AD joined devices. This property is set by Intune.
func (m *Device) SetDomainName(value *string)() {
    m.domainName = value
}
// SetEnrollmentProfileName sets the enrollmentProfileName property value. Enrollment profile applied to the device. For example, Apple Device Enrollment Profile, Device enrollment - Corporate device identifiers, or Windows Autopilot profile name. This property is set by Intune.
func (m *Device) SetEnrollmentProfileName(value *string)() {
    m.enrollmentProfileName = value
}
// SetEnrollmentType sets the enrollmentType property value. Enrollment type of the device. This property is set by Intune. Possible values are: unknown, userEnrollment, deviceEnrollmentManager, appleBulkWithUser, appleBulkWithoutUser, windowsAzureADJoin, windowsBulkUserless, windowsAutoEnrollment, windowsBulkAzureDomainJoin, windowsCoManagement.
func (m *Device) SetEnrollmentType(value *string)() {
    m.enrollmentType = value
}
// SetExtensionAttributes sets the extensionAttributes property value. Contains extension attributes 1-15 for the device. The individual extension attributes are not selectable. These properties are mastered in cloud and can be set during creation or update of a device object in Azure AD. Supports $filter (eq, not, startsWith, and eq on null values).
func (m *Device) SetExtensionAttributes(value OnPremisesExtensionAttributesable)() {
    m.extensionAttributes = value
}
// SetExtensions sets the extensions property value. The collection of open extensions defined for the device. Read-only. Nullable.
func (m *Device) SetExtensions(value []Extensionable)() {
    m.extensions = value
}
// SetHostnames sets the hostnames property value. List of hostNames for the device.
func (m *Device) SetHostnames(value []string)() {
    m.hostnames = value
}
// SetIsCompliant sets the isCompliant property value. true if the device complies with Mobile Device Management (MDM) policies; otherwise, false. Read-only. This can only be updated by Intune for any device OS type or by an approved MDM app for Windows OS devices. Supports $filter (eq, ne, not).
func (m *Device) SetIsCompliant(value *bool)() {
    m.isCompliant = value
}
// SetIsManaged sets the isManaged property value. true if the device is managed by a Mobile Device Management (MDM) app; otherwise, false. This can only be updated by Intune for any device OS type or by an approved MDM app for Windows OS devices. Supports $filter (eq, ne, not).
func (m *Device) SetIsManaged(value *bool)() {
    m.isManaged = value
}
// SetIsManagementRestricted sets the isManagementRestricted property value. The isManagementRestricted property
func (m *Device) SetIsManagementRestricted(value *bool)() {
    m.isManagementRestricted = value
}
// SetIsRooted sets the isRooted property value. true if device is rooted; false if device is jail-broken. This can only be updated by Intune.
func (m *Device) SetIsRooted(value *bool)() {
    m.isRooted = value
}
// SetKind sets the kind property value. Form factor of device. Only returned if user signs in with a Microsoft account as part of Project Rome.
func (m *Device) SetKind(value *string)() {
    m.kind = value
}
// SetManagementType sets the managementType property value. Management channel of the device.  This property is set by Intune. Possible values are: eas, mdm, easMdm, intuneClient, easIntuneClient, configurationManagerClient, configurationManagerClientMdm, configurationManagerClientMdmEas, unknown, jamf, googleCloudDevicePolicyController.
func (m *Device) SetManagementType(value *string)() {
    m.managementType = value
}
// SetManufacturer sets the manufacturer property value. Manufacturer of the device. Read-only.
func (m *Device) SetManufacturer(value *string)() {
    m.manufacturer = value
}
// SetMdmAppId sets the mdmAppId property value. Application identifier used to register device into MDM. Read-only. Supports $filter (eq, ne, not, startsWith).
func (m *Device) SetMdmAppId(value *string)() {
    m.mdmAppId = value
}
// SetMemberOf sets the memberOf property value. Groups and administrative units that this device is a member of. Read-only. Nullable. Supports $expand.
func (m *Device) SetMemberOf(value []DirectoryObjectable)() {
    m.memberOf = value
}
// SetModel sets the model property value. Model of the device. Read-only.
func (m *Device) SetModel(value *string)() {
    m.model = value
}
// SetName sets the name property value. Friendly name of a device. Only returned if user signs in with a Microsoft account as part of Project Rome.
func (m *Device) SetName(value *string)() {
    m.name = value
}
// SetOnPremisesLastSyncDateTime sets the onPremisesLastSyncDateTime property value. The last time at which the object was synced with the on-premises directory. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z Read-only. Supports $filter (eq, ne, not, ge, le, in).
func (m *Device) SetOnPremisesLastSyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.onPremisesLastSyncDateTime = value
}
// SetOnPremisesSyncEnabled sets the onPremisesSyncEnabled property value. true if this object is synced from an on-premises directory; false if this object was originally synced from an on-premises directory but is no longer synced; null if this object has never been synced from an on-premises directory (default). Read-only. Supports $filter (eq, ne, not, in, and eq on null values).
func (m *Device) SetOnPremisesSyncEnabled(value *bool)() {
    m.onPremisesSyncEnabled = value
}
// SetOperatingSystem sets the operatingSystem property value. The type of operating system on the device. Required. Supports $filter (eq, ne, not, ge, le, startsWith, and eq on null values).
func (m *Device) SetOperatingSystem(value *string)() {
    m.operatingSystem = value
}
// SetOperatingSystemVersion sets the operatingSystemVersion property value. Operating system version of the device. Required. Supports $filter (eq, ne, not, ge, le, startsWith, and eq on null values).
func (m *Device) SetOperatingSystemVersion(value *string)() {
    m.operatingSystemVersion = value
}
// SetPhysicalIds sets the physicalIds property value. For internal use only. Not nullable. Supports $filter (eq, not, ge, le, startsWith, and counting empty collections).
func (m *Device) SetPhysicalIds(value []string)() {
    m.physicalIds = value
}
// SetPlatform sets the platform property value. Platform of device. Only returned if user signs in with a Microsoft account as part of Project Rome. Only returned if user signs in with a Microsoft account as part of Project Rome.
func (m *Device) SetPlatform(value *string)() {
    m.platform = value
}
// SetProfileType sets the profileType property value. The profile type of the device. Possible values: RegisteredDevice (default), SecureVM, Printer, Shared, IoT.
func (m *Device) SetProfileType(value *string)() {
    m.profileType = value
}
// SetRegisteredOwners sets the registeredOwners property value. The user that cloud joined the device or registered their personal device. The registered owner is set at the time of registration. Currently, there can be only one owner. Read-only. Nullable. Supports $expand.
func (m *Device) SetRegisteredOwners(value []DirectoryObjectable)() {
    m.registeredOwners = value
}
// SetRegisteredUsers sets the registeredUsers property value. Collection of registered users of the device. For cloud joined devices and registered personal devices, registered users are set to the same value as registered owners at the time of registration. Read-only. Nullable. Supports $expand.
func (m *Device) SetRegisteredUsers(value []DirectoryObjectable)() {
    m.registeredUsers = value
}
// SetRegistrationDateTime sets the registrationDateTime property value. Date and time of when the device was registered. The timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
func (m *Device) SetRegistrationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.registrationDateTime = value
}
// SetStatus sets the status property value. Device is online or offline. Only returned if user signs in with a Microsoft account as part of Project Rome.
func (m *Device) SetStatus(value *string)() {
    m.status = value
}
// SetSystemLabels sets the systemLabels property value. List of labels applied to the device by the system. Supports $filter (eq when counting empty collections).
func (m *Device) SetSystemLabels(value []string)() {
    m.systemLabels = value
}
// SetTransitiveMemberOf sets the transitiveMemberOf property value. Groups and administrative units that this device is a member of. This operation is transitive. Supports $expand.
func (m *Device) SetTransitiveMemberOf(value []DirectoryObjectable)() {
    m.transitiveMemberOf = value
}
// SetTrustType sets the trustType property value. Type of trust for the joined device. Read-only. Possible values: Workplace (indicates bring your own personal devices), AzureAd (Cloud only joined devices), ServerAd (on-premises domain joined devices joined to Azure AD). For more details, see Introduction to device management in Azure Active Directory
func (m *Device) SetTrustType(value *string)() {
    m.trustType = value
}
// SetUsageRights sets the usageRights property value. Represents the usage rights a device has been granted.
func (m *Device) SetUsageRights(value []UsageRightable)() {
    m.usageRights = value
}
