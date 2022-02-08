package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Device 
type Device struct {
    DirectoryObject
    // true if the account is enabled; otherwise, false. Required. Default is true.  Supports $filter (eq, ne, not, in). Only callers in Global Administrator and Cloud Device Administrator roles can set this property.
    accountEnabled *bool;
    // For internal use only. Not nullable. Supports $filter (eq, not, ge, le).
    alternativeSecurityIds []AlternativeSecurityId;
    // The timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only. Supports $filter (eq, ne, not, ge, le, and eq on null values) and $orderBy.
    approximateLastSignInDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Set of commands sent to this device.
    commands []Command;
    // The timestamp when the device is no longer deemed compliant. The timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
    complianceExpirationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // User-defined property set by Intune to automatically add devices to groups and simplify managing devices.
    deviceCategory *string;
    // Unique identifier set by Azure Device Registration Service at the time of registration. Supports $filter (eq, ne, not, startsWith).
    deviceId *string;
    // For internal use only. Set to null.
    deviceMetadata *string;
    // Ownership of the device. This property is set by Intune. Possible values are: unknown, company, personal.
    deviceOwnership *string;
    // For internal use only.
    deviceVersion *int32;
    // The display name for the device. Required. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values), $search, and $orderBy.
    displayName *string;
    // The on-premises domain name of Hybrid Azure AD joined devices. This property is set by Intune.
    domainName *string;
    // Enrollment profile applied to the device. For example, Apple Device Enrollment Profile, Device enrollment - Corporate device identifiers, or Windows Autopilot profile name. This property is set by Intune.
    enrollmentProfileName *string;
    // Enrollment type of the device. This property is set by Intune. Possible values are: unknown, userEnrollment, deviceEnrollmentManager, appleBulkWithUser, appleBulkWithoutUser, windowsAzureADJoin, windowsBulkUserless, windowsAutoEnrollment, windowsBulkAzureDomainJoin, windowsCoManagement.
    enrollmentType *string;
    // Contains extension attributes 1-15 for the device. The individual extension attributes are not selectable. These properties are mastered in cloud and can be set during creation or update of a device object in Azure AD. Supports $filter (eq, not, startsWith, and eq on null values).
    extensionAttributes *OnPremisesExtensionAttributes;
    // The collection of open extensions defined for the device. Read-only. Nullable.
    extensions []Extension;
    // List of hostNames for the device.
    hostnames []string;
    // true if the device complies with Mobile Device Management (MDM) policies; otherwise, false. Read-only. This can only be updated by Intune for any device OS type or by an approved MDM app for Windows OS devices. Supports $filter (eq, ne, not).
    isCompliant *bool;
    // true if the device is managed by a Mobile Device Management (MDM) app; otherwise, false. This can only be updated by Intune for any device OS type or by an approved MDM app for Windows OS devices. Supports $filter (eq, ne, not).
    isManaged *bool;
    // true if device is rooted; false if device is jail-broken. This can only be updated by Intune.
    isRooted *bool;
    // Form factor of device. Only returned if user signs in with a Microsoft account as part of Project Rome.
    kind *string;
    // Management channel of the device.  This property is set by Intune. Possible values are: eas, mdm, easMdm, intuneClient, easIntuneClient, configurationManagerClient, configurationManagerClientMdm, configurationManagerClientMdmEas, unknown, jamf, googleCloudDevicePolicyController.
    managementType *string;
    // Manufacturer of the device. Read-only.
    manufacturer *string;
    // Application identifier used to register device into MDM. Read-only. Supports $filter (eq, ne, not, startsWith).
    mdmAppId *string;
    // Groups that this device is a member of. Read-only. Nullable. Supports $expand.
    memberOf []DirectoryObject;
    // Model of the device. Read-only.
    model *string;
    // Friendly name of a device. Only returned if user signs in with a Microsoft account as part of Project Rome.
    name *string;
    // The last time at which the object was synced with the on-premises directory. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z Read-only. Supports $filter (eq, ne, not, ge, le, in).
    onPremisesLastSyncDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // true if this object is synced from an on-premises directory; false if this object was originally synced from an on-premises directory but is no longer synced; null if this object has never been synced from an on-premises directory (default). Read-only. Supports $filter (eq, ne, not, in, and eq on null values).
    onPremisesSyncEnabled *bool;
    // The type of operating system on the device. Required. Supports $filter (eq, ne, not, ge, le, startsWith, and eq on null values).
    operatingSystem *string;
    // The version of the operating system on the device. Required. Supports $filter (eq, ne, not, ge, le, startsWith, and eq on null values).
    operatingSystemVersion *string;
    // For internal use only. Not nullable. Supports $filter (eq, not, ge, le, startsWith).
    physicalIds []string;
    // Platform of device. Only returned if user signs in with a Microsoft account as part of Project Rome. Only returned if user signs in with a Microsoft account as part of Project Rome.
    platform *string;
    // The profile type of the device. Possible values: RegisteredDevice (default), SecureVM, Printer, Shared, IoT.
    profileType *string;
    // The user that cloud joined the device or registered their personal device. The registered owner is set at the time of registration. Currently, there can be only one owner. Read-only. Nullable. Supports $expand.
    registeredOwners []DirectoryObject;
    // Collection of registered users of the device. For cloud joined devices and registered personal devices, registered users are set to the same value as registered owners at the time of registration. Read-only. Nullable. Supports $expand.
    registeredUsers []DirectoryObject;
    // Date and time of when the device was registered. The timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
    registrationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Device is online or offline. Only returned if user signs in with a Microsoft account as part of Project Rome.
    status *string;
    // List of labels applied to the device by the system.
    systemLabels []string;
    // Groups that the device is a member of. This operation is transitive. Supports $expand.
    transitiveMemberOf []DirectoryObject;
    // Type of trust for the joined device. Read-only. Possible values:  Workplace (indicates bring your own personal devices), AzureAd (Cloud only joined devices), ServerAd (on-premises domain joined devices joined to Azure AD). For more details, see Introduction to device management in Azure Active Directory
    trustType *string;
    // Represents the usage rights a device has been granted.
    usageRights []UsageRight;
}
// NewDevice instantiates a new device and sets the default values.
func NewDevice()(*Device) {
    m := &Device{
        DirectoryObject: *NewDirectoryObject(),
    }
    return m
}
// GetAccountEnabled gets the accountEnabled property value. true if the account is enabled; otherwise, false. Required. Default is true.  Supports $filter (eq, ne, not, in). Only callers in Global Administrator and Cloud Device Administrator roles can set this property.
func (m *Device) GetAccountEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.accountEnabled
    }
}
// GetAlternativeSecurityIds gets the alternativeSecurityIds property value. For internal use only. Not nullable. Supports $filter (eq, not, ge, le).
func (m *Device) GetAlternativeSecurityIds()([]AlternativeSecurityId) {
    if m == nil {
        return nil
    } else {
        return m.alternativeSecurityIds
    }
}
// GetApproximateLastSignInDateTime gets the approximateLastSignInDateTime property value. The timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only. Supports $filter (eq, ne, not, ge, le, and eq on null values) and $orderBy.
func (m *Device) GetApproximateLastSignInDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.approximateLastSignInDateTime
    }
}
// GetCommands gets the commands property value. Set of commands sent to this device.
func (m *Device) GetCommands()([]Command) {
    if m == nil {
        return nil
    } else {
        return m.commands
    }
}
// GetComplianceExpirationDateTime gets the complianceExpirationDateTime property value. The timestamp when the device is no longer deemed compliant. The timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
func (m *Device) GetComplianceExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.complianceExpirationDateTime
    }
}
// GetDeviceCategory gets the deviceCategory property value. User-defined property set by Intune to automatically add devices to groups and simplify managing devices.
func (m *Device) GetDeviceCategory()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceCategory
    }
}
// GetDeviceId gets the deviceId property value. Unique identifier set by Azure Device Registration Service at the time of registration. Supports $filter (eq, ne, not, startsWith).
func (m *Device) GetDeviceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceId
    }
}
// GetDeviceMetadata gets the deviceMetadata property value. For internal use only. Set to null.
func (m *Device) GetDeviceMetadata()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceMetadata
    }
}
// GetDeviceOwnership gets the deviceOwnership property value. Ownership of the device. This property is set by Intune. Possible values are: unknown, company, personal.
func (m *Device) GetDeviceOwnership()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceOwnership
    }
}
// GetDeviceVersion gets the deviceVersion property value. For internal use only.
func (m *Device) GetDeviceVersion()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.deviceVersion
    }
}
// GetDisplayName gets the displayName property value. The display name for the device. Required. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values), $search, and $orderBy.
func (m *Device) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetDomainName gets the domainName property value. The on-premises domain name of Hybrid Azure AD joined devices. This property is set by Intune.
func (m *Device) GetDomainName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.domainName
    }
}
// GetEnrollmentProfileName gets the enrollmentProfileName property value. Enrollment profile applied to the device. For example, Apple Device Enrollment Profile, Device enrollment - Corporate device identifiers, or Windows Autopilot profile name. This property is set by Intune.
func (m *Device) GetEnrollmentProfileName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.enrollmentProfileName
    }
}
// GetEnrollmentType gets the enrollmentType property value. Enrollment type of the device. This property is set by Intune. Possible values are: unknown, userEnrollment, deviceEnrollmentManager, appleBulkWithUser, appleBulkWithoutUser, windowsAzureADJoin, windowsBulkUserless, windowsAutoEnrollment, windowsBulkAzureDomainJoin, windowsCoManagement.
func (m *Device) GetEnrollmentType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.enrollmentType
    }
}
// GetExtensionAttributes gets the extensionAttributes property value. Contains extension attributes 1-15 for the device. The individual extension attributes are not selectable. These properties are mastered in cloud and can be set during creation or update of a device object in Azure AD. Supports $filter (eq, not, startsWith, and eq on null values).
func (m *Device) GetExtensionAttributes()(*OnPremisesExtensionAttributes) {
    if m == nil {
        return nil
    } else {
        return m.extensionAttributes
    }
}
// GetExtensions gets the extensions property value. The collection of open extensions defined for the device. Read-only. Nullable.
func (m *Device) GetExtensions()([]Extension) {
    if m == nil {
        return nil
    } else {
        return m.extensions
    }
}
// GetHostnames gets the hostnames property value. List of hostNames for the device.
func (m *Device) GetHostnames()([]string) {
    if m == nil {
        return nil
    } else {
        return m.hostnames
    }
}
// GetIsCompliant gets the isCompliant property value. true if the device complies with Mobile Device Management (MDM) policies; otherwise, false. Read-only. This can only be updated by Intune for any device OS type or by an approved MDM app for Windows OS devices. Supports $filter (eq, ne, not).
func (m *Device) GetIsCompliant()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isCompliant
    }
}
// GetIsManaged gets the isManaged property value. true if the device is managed by a Mobile Device Management (MDM) app; otherwise, false. This can only be updated by Intune for any device OS type or by an approved MDM app for Windows OS devices. Supports $filter (eq, ne, not).
func (m *Device) GetIsManaged()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isManaged
    }
}
// GetIsRooted gets the isRooted property value. true if device is rooted; false if device is jail-broken. This can only be updated by Intune.
func (m *Device) GetIsRooted()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isRooted
    }
}
// GetKind gets the kind property value. Form factor of device. Only returned if user signs in with a Microsoft account as part of Project Rome.
func (m *Device) GetKind()(*string) {
    if m == nil {
        return nil
    } else {
        return m.kind
    }
}
// GetManagementType gets the managementType property value. Management channel of the device.  This property is set by Intune. Possible values are: eas, mdm, easMdm, intuneClient, easIntuneClient, configurationManagerClient, configurationManagerClientMdm, configurationManagerClientMdmEas, unknown, jamf, googleCloudDevicePolicyController.
func (m *Device) GetManagementType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.managementType
    }
}
// GetManufacturer gets the manufacturer property value. Manufacturer of the device. Read-only.
func (m *Device) GetManufacturer()(*string) {
    if m == nil {
        return nil
    } else {
        return m.manufacturer
    }
}
// GetMdmAppId gets the mdmAppId property value. Application identifier used to register device into MDM. Read-only. Supports $filter (eq, ne, not, startsWith).
func (m *Device) GetMdmAppId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.mdmAppId
    }
}
// GetMemberOf gets the memberOf property value. Groups that this device is a member of. Read-only. Nullable. Supports $expand.
func (m *Device) GetMemberOf()([]DirectoryObject) {
    if m == nil {
        return nil
    } else {
        return m.memberOf
    }
}
// GetModel gets the model property value. Model of the device. Read-only.
func (m *Device) GetModel()(*string) {
    if m == nil {
        return nil
    } else {
        return m.model
    }
}
// GetName gets the name property value. Friendly name of a device. Only returned if user signs in with a Microsoft account as part of Project Rome.
func (m *Device) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
// GetOnPremisesLastSyncDateTime gets the onPremisesLastSyncDateTime property value. The last time at which the object was synced with the on-premises directory. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z Read-only. Supports $filter (eq, ne, not, ge, le, in).
func (m *Device) GetOnPremisesLastSyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.onPremisesLastSyncDateTime
    }
}
// GetOnPremisesSyncEnabled gets the onPremisesSyncEnabled property value. true if this object is synced from an on-premises directory; false if this object was originally synced from an on-premises directory but is no longer synced; null if this object has never been synced from an on-premises directory (default). Read-only. Supports $filter (eq, ne, not, in, and eq on null values).
func (m *Device) GetOnPremisesSyncEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.onPremisesSyncEnabled
    }
}
// GetOperatingSystem gets the operatingSystem property value. The type of operating system on the device. Required. Supports $filter (eq, ne, not, ge, le, startsWith, and eq on null values).
func (m *Device) GetOperatingSystem()(*string) {
    if m == nil {
        return nil
    } else {
        return m.operatingSystem
    }
}
// GetOperatingSystemVersion gets the operatingSystemVersion property value. The version of the operating system on the device. Required. Supports $filter (eq, ne, not, ge, le, startsWith, and eq on null values).
func (m *Device) GetOperatingSystemVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.operatingSystemVersion
    }
}
// GetPhysicalIds gets the physicalIds property value. For internal use only. Not nullable. Supports $filter (eq, not, ge, le, startsWith).
func (m *Device) GetPhysicalIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.physicalIds
    }
}
// GetPlatform gets the platform property value. Platform of device. Only returned if user signs in with a Microsoft account as part of Project Rome. Only returned if user signs in with a Microsoft account as part of Project Rome.
func (m *Device) GetPlatform()(*string) {
    if m == nil {
        return nil
    } else {
        return m.platform
    }
}
// GetProfileType gets the profileType property value. The profile type of the device. Possible values: RegisteredDevice (default), SecureVM, Printer, Shared, IoT.
func (m *Device) GetProfileType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.profileType
    }
}
// GetRegisteredOwners gets the registeredOwners property value. The user that cloud joined the device or registered their personal device. The registered owner is set at the time of registration. Currently, there can be only one owner. Read-only. Nullable. Supports $expand.
func (m *Device) GetRegisteredOwners()([]DirectoryObject) {
    if m == nil {
        return nil
    } else {
        return m.registeredOwners
    }
}
// GetRegisteredUsers gets the registeredUsers property value. Collection of registered users of the device. For cloud joined devices and registered personal devices, registered users are set to the same value as registered owners at the time of registration. Read-only. Nullable. Supports $expand.
func (m *Device) GetRegisteredUsers()([]DirectoryObject) {
    if m == nil {
        return nil
    } else {
        return m.registeredUsers
    }
}
// GetRegistrationDateTime gets the registrationDateTime property value. Date and time of when the device was registered. The timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
func (m *Device) GetRegistrationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.registrationDateTime
    }
}
// GetStatus gets the status property value. Device is online or offline. Only returned if user signs in with a Microsoft account as part of Project Rome.
func (m *Device) GetStatus()(*string) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
// GetSystemLabels gets the systemLabels property value. List of labels applied to the device by the system.
func (m *Device) GetSystemLabels()([]string) {
    if m == nil {
        return nil
    } else {
        return m.systemLabels
    }
}
// GetTransitiveMemberOf gets the transitiveMemberOf property value. Groups that the device is a member of. This operation is transitive. Supports $expand.
func (m *Device) GetTransitiveMemberOf()([]DirectoryObject) {
    if m == nil {
        return nil
    } else {
        return m.transitiveMemberOf
    }
}
// GetTrustType gets the trustType property value. Type of trust for the joined device. Read-only. Possible values:  Workplace (indicates bring your own personal devices), AzureAd (Cloud only joined devices), ServerAd (on-premises domain joined devices joined to Azure AD). For more details, see Introduction to device management in Azure Active Directory
func (m *Device) GetTrustType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.trustType
    }
}
// GetUsageRights gets the usageRights property value. Represents the usage rights a device has been granted.
func (m *Device) GetUsageRights()([]UsageRight) {
    if m == nil {
        return nil
    } else {
        return m.usageRights
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Device) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.DirectoryObject.GetFieldDeserializers()
    res["accountEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccountEnabled(val)
        }
        return nil
    }
    res["alternativeSecurityIds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAlternativeSecurityId() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AlternativeSecurityId, len(val))
            for i, v := range val {
                res[i] = *(v.(*AlternativeSecurityId))
            }
            m.SetAlternativeSecurityIds(res)
        }
        return nil
    }
    res["approximateLastSignInDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApproximateLastSignInDateTime(val)
        }
        return nil
    }
    res["commands"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCommand() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Command, len(val))
            for i, v := range val {
                res[i] = *(v.(*Command))
            }
            m.SetCommands(res)
        }
        return nil
    }
    res["complianceExpirationDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetComplianceExpirationDateTime(val)
        }
        return nil
    }
    res["deviceCategory"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceCategory(val)
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
    res["deviceMetadata"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceMetadata(val)
        }
        return nil
    }
    res["deviceOwnership"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceOwnership(val)
        }
        return nil
    }
    res["deviceVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceVersion(val)
        }
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["domainName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDomainName(val)
        }
        return nil
    }
    res["enrollmentProfileName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnrollmentProfileName(val)
        }
        return nil
    }
    res["enrollmentType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnrollmentType(val)
        }
        return nil
    }
    res["extensionAttributes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewOnPremisesExtensionAttributes() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExtensionAttributes(val.(*OnPremisesExtensionAttributes))
        }
        return nil
    }
    res["extensions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewExtension() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Extension, len(val))
            for i, v := range val {
                res[i] = *(v.(*Extension))
            }
            m.SetExtensions(res)
        }
        return nil
    }
    res["hostnames"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetHostnames(res)
        }
        return nil
    }
    res["isCompliant"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsCompliant(val)
        }
        return nil
    }
    res["isManaged"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsManaged(val)
        }
        return nil
    }
    res["isRooted"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsRooted(val)
        }
        return nil
    }
    res["kind"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKind(val)
        }
        return nil
    }
    res["managementType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagementType(val)
        }
        return nil
    }
    res["manufacturer"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManufacturer(val)
        }
        return nil
    }
    res["mdmAppId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMdmAppId(val)
        }
        return nil
    }
    res["memberOf"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDirectoryObject() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DirectoryObject, len(val))
            for i, v := range val {
                res[i] = *(v.(*DirectoryObject))
            }
            m.SetMemberOf(res)
        }
        return nil
    }
    res["model"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetModel(val)
        }
        return nil
    }
    res["name"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["onPremisesLastSyncDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOnPremisesLastSyncDateTime(val)
        }
        return nil
    }
    res["onPremisesSyncEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOnPremisesSyncEnabled(val)
        }
        return nil
    }
    res["operatingSystem"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOperatingSystem(val)
        }
        return nil
    }
    res["operatingSystemVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOperatingSystemVersion(val)
        }
        return nil
    }
    res["physicalIds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetPhysicalIds(res)
        }
        return nil
    }
    res["platform"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPlatform(val)
        }
        return nil
    }
    res["profileType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProfileType(val)
        }
        return nil
    }
    res["registeredOwners"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDirectoryObject() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DirectoryObject, len(val))
            for i, v := range val {
                res[i] = *(v.(*DirectoryObject))
            }
            m.SetRegisteredOwners(res)
        }
        return nil
    }
    res["registeredUsers"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDirectoryObject() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DirectoryObject, len(val))
            for i, v := range val {
                res[i] = *(v.(*DirectoryObject))
            }
            m.SetRegisteredUsers(res)
        }
        return nil
    }
    res["registrationDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRegistrationDateTime(val)
        }
        return nil
    }
    res["status"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val)
        }
        return nil
    }
    res["systemLabels"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetSystemLabels(res)
        }
        return nil
    }
    res["transitiveMemberOf"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDirectoryObject() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DirectoryObject, len(val))
            for i, v := range val {
                res[i] = *(v.(*DirectoryObject))
            }
            m.SetTransitiveMemberOf(res)
        }
        return nil
    }
    res["trustType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTrustType(val)
        }
        return nil
    }
    res["usageRights"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUsageRight() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UsageRight, len(val))
            for i, v := range val {
                res[i] = *(v.(*UsageRight))
            }
            m.SetUsageRights(res)
        }
        return nil
    }
    return res
}
func (m *Device) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *Device) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAlternativeSecurityIds()))
        for i, v := range m.GetAlternativeSecurityIds() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
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
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetCommands()))
        for i, v := range m.GetCommands() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
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
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetExtensions()))
        for i, v := range m.GetExtensions() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
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
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetMemberOf()))
        for i, v := range m.GetMemberOf() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
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
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetRegisteredOwners()))
        for i, v := range m.GetRegisteredOwners() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("registeredOwners", cast)
        if err != nil {
            return err
        }
    }
    if m.GetRegisteredUsers() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetRegisteredUsers()))
        for i, v := range m.GetRegisteredUsers() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
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
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetTransitiveMemberOf()))
        for i, v := range m.GetTransitiveMemberOf() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
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
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetUsageRights()))
        for i, v := range m.GetUsageRights() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("usageRights", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAccountEnabled sets the accountEnabled property value. true if the account is enabled; otherwise, false. Required. Default is true.  Supports $filter (eq, ne, not, in). Only callers in Global Administrator and Cloud Device Administrator roles can set this property.
func (m *Device) SetAccountEnabled(value *bool)() {
    if m != nil {
        m.accountEnabled = value
    }
}
// SetAlternativeSecurityIds sets the alternativeSecurityIds property value. For internal use only. Not nullable. Supports $filter (eq, not, ge, le).
func (m *Device) SetAlternativeSecurityIds(value []AlternativeSecurityId)() {
    if m != nil {
        m.alternativeSecurityIds = value
    }
}
// SetApproximateLastSignInDateTime sets the approximateLastSignInDateTime property value. The timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only. Supports $filter (eq, ne, not, ge, le, and eq on null values) and $orderBy.
func (m *Device) SetApproximateLastSignInDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.approximateLastSignInDateTime = value
    }
}
// SetCommands sets the commands property value. Set of commands sent to this device.
func (m *Device) SetCommands(value []Command)() {
    if m != nil {
        m.commands = value
    }
}
// SetComplianceExpirationDateTime sets the complianceExpirationDateTime property value. The timestamp when the device is no longer deemed compliant. The timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
func (m *Device) SetComplianceExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.complianceExpirationDateTime = value
    }
}
// SetDeviceCategory sets the deviceCategory property value. User-defined property set by Intune to automatically add devices to groups and simplify managing devices.
func (m *Device) SetDeviceCategory(value *string)() {
    if m != nil {
        m.deviceCategory = value
    }
}
// SetDeviceId sets the deviceId property value. Unique identifier set by Azure Device Registration Service at the time of registration. Supports $filter (eq, ne, not, startsWith).
func (m *Device) SetDeviceId(value *string)() {
    if m != nil {
        m.deviceId = value
    }
}
// SetDeviceMetadata sets the deviceMetadata property value. For internal use only. Set to null.
func (m *Device) SetDeviceMetadata(value *string)() {
    if m != nil {
        m.deviceMetadata = value
    }
}
// SetDeviceOwnership sets the deviceOwnership property value. Ownership of the device. This property is set by Intune. Possible values are: unknown, company, personal.
func (m *Device) SetDeviceOwnership(value *string)() {
    if m != nil {
        m.deviceOwnership = value
    }
}
// SetDeviceVersion sets the deviceVersion property value. For internal use only.
func (m *Device) SetDeviceVersion(value *int32)() {
    if m != nil {
        m.deviceVersion = value
    }
}
// SetDisplayName sets the displayName property value. The display name for the device. Required. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values), $search, and $orderBy.
func (m *Device) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetDomainName sets the domainName property value. The on-premises domain name of Hybrid Azure AD joined devices. This property is set by Intune.
func (m *Device) SetDomainName(value *string)() {
    if m != nil {
        m.domainName = value
    }
}
// SetEnrollmentProfileName sets the enrollmentProfileName property value. Enrollment profile applied to the device. For example, Apple Device Enrollment Profile, Device enrollment - Corporate device identifiers, or Windows Autopilot profile name. This property is set by Intune.
func (m *Device) SetEnrollmentProfileName(value *string)() {
    if m != nil {
        m.enrollmentProfileName = value
    }
}
// SetEnrollmentType sets the enrollmentType property value. Enrollment type of the device. This property is set by Intune. Possible values are: unknown, userEnrollment, deviceEnrollmentManager, appleBulkWithUser, appleBulkWithoutUser, windowsAzureADJoin, windowsBulkUserless, windowsAutoEnrollment, windowsBulkAzureDomainJoin, windowsCoManagement.
func (m *Device) SetEnrollmentType(value *string)() {
    if m != nil {
        m.enrollmentType = value
    }
}
// SetExtensionAttributes sets the extensionAttributes property value. Contains extension attributes 1-15 for the device. The individual extension attributes are not selectable. These properties are mastered in cloud and can be set during creation or update of a device object in Azure AD. Supports $filter (eq, not, startsWith, and eq on null values).
func (m *Device) SetExtensionAttributes(value *OnPremisesExtensionAttributes)() {
    if m != nil {
        m.extensionAttributes = value
    }
}
// SetExtensions sets the extensions property value. The collection of open extensions defined for the device. Read-only. Nullable.
func (m *Device) SetExtensions(value []Extension)() {
    if m != nil {
        m.extensions = value
    }
}
// SetHostnames sets the hostnames property value. List of hostNames for the device.
func (m *Device) SetHostnames(value []string)() {
    if m != nil {
        m.hostnames = value
    }
}
// SetIsCompliant sets the isCompliant property value. true if the device complies with Mobile Device Management (MDM) policies; otherwise, false. Read-only. This can only be updated by Intune for any device OS type or by an approved MDM app for Windows OS devices. Supports $filter (eq, ne, not).
func (m *Device) SetIsCompliant(value *bool)() {
    if m != nil {
        m.isCompliant = value
    }
}
// SetIsManaged sets the isManaged property value. true if the device is managed by a Mobile Device Management (MDM) app; otherwise, false. This can only be updated by Intune for any device OS type or by an approved MDM app for Windows OS devices. Supports $filter (eq, ne, not).
func (m *Device) SetIsManaged(value *bool)() {
    if m != nil {
        m.isManaged = value
    }
}
// SetIsRooted sets the isRooted property value. true if device is rooted; false if device is jail-broken. This can only be updated by Intune.
func (m *Device) SetIsRooted(value *bool)() {
    if m != nil {
        m.isRooted = value
    }
}
// SetKind sets the kind property value. Form factor of device. Only returned if user signs in with a Microsoft account as part of Project Rome.
func (m *Device) SetKind(value *string)() {
    if m != nil {
        m.kind = value
    }
}
// SetManagementType sets the managementType property value. Management channel of the device.  This property is set by Intune. Possible values are: eas, mdm, easMdm, intuneClient, easIntuneClient, configurationManagerClient, configurationManagerClientMdm, configurationManagerClientMdmEas, unknown, jamf, googleCloudDevicePolicyController.
func (m *Device) SetManagementType(value *string)() {
    if m != nil {
        m.managementType = value
    }
}
// SetManufacturer sets the manufacturer property value. Manufacturer of the device. Read-only.
func (m *Device) SetManufacturer(value *string)() {
    if m != nil {
        m.manufacturer = value
    }
}
// SetMdmAppId sets the mdmAppId property value. Application identifier used to register device into MDM. Read-only. Supports $filter (eq, ne, not, startsWith).
func (m *Device) SetMdmAppId(value *string)() {
    if m != nil {
        m.mdmAppId = value
    }
}
// SetMemberOf sets the memberOf property value. Groups that this device is a member of. Read-only. Nullable. Supports $expand.
func (m *Device) SetMemberOf(value []DirectoryObject)() {
    if m != nil {
        m.memberOf = value
    }
}
// SetModel sets the model property value. Model of the device. Read-only.
func (m *Device) SetModel(value *string)() {
    if m != nil {
        m.model = value
    }
}
// SetName sets the name property value. Friendly name of a device. Only returned if user signs in with a Microsoft account as part of Project Rome.
func (m *Device) SetName(value *string)() {
    if m != nil {
        m.name = value
    }
}
// SetOnPremisesLastSyncDateTime sets the onPremisesLastSyncDateTime property value. The last time at which the object was synced with the on-premises directory. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z Read-only. Supports $filter (eq, ne, not, ge, le, in).
func (m *Device) SetOnPremisesLastSyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.onPremisesLastSyncDateTime = value
    }
}
// SetOnPremisesSyncEnabled sets the onPremisesSyncEnabled property value. true if this object is synced from an on-premises directory; false if this object was originally synced from an on-premises directory but is no longer synced; null if this object has never been synced from an on-premises directory (default). Read-only. Supports $filter (eq, ne, not, in, and eq on null values).
func (m *Device) SetOnPremisesSyncEnabled(value *bool)() {
    if m != nil {
        m.onPremisesSyncEnabled = value
    }
}
// SetOperatingSystem sets the operatingSystem property value. The type of operating system on the device. Required. Supports $filter (eq, ne, not, ge, le, startsWith, and eq on null values).
func (m *Device) SetOperatingSystem(value *string)() {
    if m != nil {
        m.operatingSystem = value
    }
}
// SetOperatingSystemVersion sets the operatingSystemVersion property value. The version of the operating system on the device. Required. Supports $filter (eq, ne, not, ge, le, startsWith, and eq on null values).
func (m *Device) SetOperatingSystemVersion(value *string)() {
    if m != nil {
        m.operatingSystemVersion = value
    }
}
// SetPhysicalIds sets the physicalIds property value. For internal use only. Not nullable. Supports $filter (eq, not, ge, le, startsWith).
func (m *Device) SetPhysicalIds(value []string)() {
    if m != nil {
        m.physicalIds = value
    }
}
// SetPlatform sets the platform property value. Platform of device. Only returned if user signs in with a Microsoft account as part of Project Rome. Only returned if user signs in with a Microsoft account as part of Project Rome.
func (m *Device) SetPlatform(value *string)() {
    if m != nil {
        m.platform = value
    }
}
// SetProfileType sets the profileType property value. The profile type of the device. Possible values: RegisteredDevice (default), SecureVM, Printer, Shared, IoT.
func (m *Device) SetProfileType(value *string)() {
    if m != nil {
        m.profileType = value
    }
}
// SetRegisteredOwners sets the registeredOwners property value. The user that cloud joined the device or registered their personal device. The registered owner is set at the time of registration. Currently, there can be only one owner. Read-only. Nullable. Supports $expand.
func (m *Device) SetRegisteredOwners(value []DirectoryObject)() {
    if m != nil {
        m.registeredOwners = value
    }
}
// SetRegisteredUsers sets the registeredUsers property value. Collection of registered users of the device. For cloud joined devices and registered personal devices, registered users are set to the same value as registered owners at the time of registration. Read-only. Nullable. Supports $expand.
func (m *Device) SetRegisteredUsers(value []DirectoryObject)() {
    if m != nil {
        m.registeredUsers = value
    }
}
// SetRegistrationDateTime sets the registrationDateTime property value. Date and time of when the device was registered. The timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
func (m *Device) SetRegistrationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.registrationDateTime = value
    }
}
// SetStatus sets the status property value. Device is online or offline. Only returned if user signs in with a Microsoft account as part of Project Rome.
func (m *Device) SetStatus(value *string)() {
    if m != nil {
        m.status = value
    }
}
// SetSystemLabels sets the systemLabels property value. List of labels applied to the device by the system.
func (m *Device) SetSystemLabels(value []string)() {
    if m != nil {
        m.systemLabels = value
    }
}
// SetTransitiveMemberOf sets the transitiveMemberOf property value. Groups that the device is a member of. This operation is transitive. Supports $expand.
func (m *Device) SetTransitiveMemberOf(value []DirectoryObject)() {
    if m != nil {
        m.transitiveMemberOf = value
    }
}
// SetTrustType sets the trustType property value. Type of trust for the joined device. Read-only. Possible values:  Workplace (indicates bring your own personal devices), AzureAd (Cloud only joined devices), ServerAd (on-premises domain joined devices joined to Azure AD). For more details, see Introduction to device management in Azure Active Directory
func (m *Device) SetTrustType(value *string)() {
    if m != nil {
        m.trustType = value
    }
}
// SetUsageRights sets the usageRights property value. Represents the usage rights a device has been granted.
func (m *Device) SetUsageRights(value []UsageRight)() {
    if m != nil {
        m.usageRights = value
    }
}
