package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceAppManagement 
type DeviceAppManagement struct {
    Entity
    // Android managed app policies.
    androidManagedAppProtections []AndroidManagedAppProtectionable
    // Default managed app policies.
    defaultManagedAppProtections []DefaultManagedAppProtectionable
    // Device app management tasks.
    deviceAppManagementTasks []DeviceAppManagementTaskable
    // The Windows Enterprise Code Signing Certificate.
    enterpriseCodeSigningCertificates []EnterpriseCodeSigningCertificateable
    // The IOS Lob App Provisioning Configurations.
    iosLobAppProvisioningConfigurations []IosLobAppProvisioningConfigurationable
    // iOS managed app policies.
    iosManagedAppProtections []IosManagedAppProtectionable
    // Whether the account is enabled for syncing applications from the Microsoft Store for Business.
    isEnabledForMicrosoftStoreForBusiness *bool
    // Managed app policies.
    managedAppPolicies []ManagedAppPolicyable
    // The managed app registrations.
    managedAppRegistrations []ManagedAppRegistrationable
    // The managed app statuses.
    managedAppStatuses []ManagedAppStatusable
    // The mobile eBook categories.
    managedEBookCategories []ManagedEBookCategoryable
    // The Managed eBook.
    managedEBooks []ManagedEBookable
    // Windows information protection for apps running on devices which are MDM enrolled.
    mdmWindowsInformationProtectionPolicies []MdmWindowsInformationProtectionPolicyable
    // The locale information used to sync applications from the Microsoft Store for Business. Cultures that are specific to a country/region. The names of these cultures follow RFC 4646 (Windows Vista and later). The format is -<country/regioncode2>, where  is a lowercase two-letter code derived from ISO 639-1 and <country/regioncode2> is an uppercase two-letter code derived from ISO 3166. For example, en-US for English (United States) is a specific culture.
    microsoftStoreForBusinessLanguage *string
    // The last time an application sync from the Microsoft Store for Business was completed.
    microsoftStoreForBusinessLastCompletedApplicationSyncTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The last time the apps from the Microsoft Store for Business were synced successfully for the account.
    microsoftStoreForBusinessLastSuccessfulSyncDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Portal to which admin syncs available Microsoft Store for Business apps. This is available in the Intune Admin Console.
    microsoftStoreForBusinessPortalSelection *MicrosoftStoreForBusinessPortalSelectionOptions
    // The mobile app categories.
    mobileAppCategories []MobileAppCategoryable
    // The Managed Device Mobile Application Configurations.
    mobileAppConfigurations []ManagedDeviceMobileAppConfigurationable
    // The mobile apps.
    mobileApps []MobileAppable
    // The PolicySet of Policies and Applications
    policySets []PolicySetable
    // Side Loading Keys that are required for the Windows 8 and 8.1 Apps installation.
    sideLoadingKeys []SideLoadingKeyable
    // The WinPhone Symantec Code Signing Certificate.
    symantecCodeSigningCertificate SymantecCodeSigningCertificateable
    // Targeted managed app configurations.
    targetedManagedAppConfigurations []TargetedManagedAppConfigurationable
    // List of Vpp tokens for this organization.
    vppTokens []VppTokenable
    // The collection of Windows Defender Application Control Supplemental Policies.
    wdacSupplementalPolicies []WindowsDefenderApplicationControlSupplementalPolicyable
    // Windows information protection device registrations that are not MDM enrolled.
    windowsInformationProtectionDeviceRegistrations []WindowsInformationProtectionDeviceRegistrationable
    // Windows information protection for apps running on devices which are not MDM enrolled.
    windowsInformationProtectionPolicies []WindowsInformationProtectionPolicyable
    // Windows information protection wipe actions.
    windowsInformationProtectionWipeActions []WindowsInformationProtectionWipeActionable
    // Windows managed app policies.
    windowsManagedAppProtections []WindowsManagedAppProtectionable
    // Windows management app.
    windowsManagementApp WindowsManagementAppable
}
// NewDeviceAppManagement instantiates a new DeviceAppManagement and sets the default values.
func NewDeviceAppManagement()(*DeviceAppManagement) {
    m := &DeviceAppManagement{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.deviceAppManagement";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateDeviceAppManagementFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceAppManagementFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceAppManagement(), nil
}
// GetAndroidManagedAppProtections gets the androidManagedAppProtections property value. Android managed app policies.
func (m *DeviceAppManagement) GetAndroidManagedAppProtections()([]AndroidManagedAppProtectionable) {
    return m.androidManagedAppProtections
}
// GetDefaultManagedAppProtections gets the defaultManagedAppProtections property value. Default managed app policies.
func (m *DeviceAppManagement) GetDefaultManagedAppProtections()([]DefaultManagedAppProtectionable) {
    return m.defaultManagedAppProtections
}
// GetDeviceAppManagementTasks gets the deviceAppManagementTasks property value. Device app management tasks.
func (m *DeviceAppManagement) GetDeviceAppManagementTasks()([]DeviceAppManagementTaskable) {
    return m.deviceAppManagementTasks
}
// GetEnterpriseCodeSigningCertificates gets the enterpriseCodeSigningCertificates property value. The Windows Enterprise Code Signing Certificate.
func (m *DeviceAppManagement) GetEnterpriseCodeSigningCertificates()([]EnterpriseCodeSigningCertificateable) {
    return m.enterpriseCodeSigningCertificates
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceAppManagement) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["androidManagedAppProtections"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateAndroidManagedAppProtectionFromDiscriminatorValue , m.SetAndroidManagedAppProtections)
    res["defaultManagedAppProtections"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateDefaultManagedAppProtectionFromDiscriminatorValue , m.SetDefaultManagedAppProtections)
    res["deviceAppManagementTasks"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateDeviceAppManagementTaskFromDiscriminatorValue , m.SetDeviceAppManagementTasks)
    res["enterpriseCodeSigningCertificates"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateEnterpriseCodeSigningCertificateFromDiscriminatorValue , m.SetEnterpriseCodeSigningCertificates)
    res["iosLobAppProvisioningConfigurations"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateIosLobAppProvisioningConfigurationFromDiscriminatorValue , m.SetIosLobAppProvisioningConfigurations)
    res["iosManagedAppProtections"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateIosManagedAppProtectionFromDiscriminatorValue , m.SetIosManagedAppProtections)
    res["isEnabledForMicrosoftStoreForBusiness"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIsEnabledForMicrosoftStoreForBusiness)
    res["managedAppPolicies"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateManagedAppPolicyFromDiscriminatorValue , m.SetManagedAppPolicies)
    res["managedAppRegistrations"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateManagedAppRegistrationFromDiscriminatorValue , m.SetManagedAppRegistrations)
    res["managedAppStatuses"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateManagedAppStatusFromDiscriminatorValue , m.SetManagedAppStatuses)
    res["managedEBookCategories"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateManagedEBookCategoryFromDiscriminatorValue , m.SetManagedEBookCategories)
    res["managedEBooks"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateManagedEBookFromDiscriminatorValue , m.SetManagedEBooks)
    res["mdmWindowsInformationProtectionPolicies"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateMdmWindowsInformationProtectionPolicyFromDiscriminatorValue , m.SetMdmWindowsInformationProtectionPolicies)
    res["microsoftStoreForBusinessLanguage"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetMicrosoftStoreForBusinessLanguage)
    res["microsoftStoreForBusinessLastCompletedApplicationSyncTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetMicrosoftStoreForBusinessLastCompletedApplicationSyncTime)
    res["microsoftStoreForBusinessLastSuccessfulSyncDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetMicrosoftStoreForBusinessLastSuccessfulSyncDateTime)
    res["microsoftStoreForBusinessPortalSelection"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseMicrosoftStoreForBusinessPortalSelectionOptions , m.SetMicrosoftStoreForBusinessPortalSelection)
    res["mobileAppCategories"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateMobileAppCategoryFromDiscriminatorValue , m.SetMobileAppCategories)
    res["mobileAppConfigurations"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateManagedDeviceMobileAppConfigurationFromDiscriminatorValue , m.SetMobileAppConfigurations)
    res["mobileApps"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateMobileAppFromDiscriminatorValue , m.SetMobileApps)
    res["policySets"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreatePolicySetFromDiscriminatorValue , m.SetPolicySets)
    res["sideLoadingKeys"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateSideLoadingKeyFromDiscriminatorValue , m.SetSideLoadingKeys)
    res["symantecCodeSigningCertificate"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateSymantecCodeSigningCertificateFromDiscriminatorValue , m.SetSymantecCodeSigningCertificate)
    res["targetedManagedAppConfigurations"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateTargetedManagedAppConfigurationFromDiscriminatorValue , m.SetTargetedManagedAppConfigurations)
    res["vppTokens"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateVppTokenFromDiscriminatorValue , m.SetVppTokens)
    res["wdacSupplementalPolicies"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateWindowsDefenderApplicationControlSupplementalPolicyFromDiscriminatorValue , m.SetWdacSupplementalPolicies)
    res["windowsInformationProtectionDeviceRegistrations"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateWindowsInformationProtectionDeviceRegistrationFromDiscriminatorValue , m.SetWindowsInformationProtectionDeviceRegistrations)
    res["windowsInformationProtectionPolicies"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateWindowsInformationProtectionPolicyFromDiscriminatorValue , m.SetWindowsInformationProtectionPolicies)
    res["windowsInformationProtectionWipeActions"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateWindowsInformationProtectionWipeActionFromDiscriminatorValue , m.SetWindowsInformationProtectionWipeActions)
    res["windowsManagedAppProtections"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateWindowsManagedAppProtectionFromDiscriminatorValue , m.SetWindowsManagedAppProtections)
    res["windowsManagementApp"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateWindowsManagementAppFromDiscriminatorValue , m.SetWindowsManagementApp)
    return res
}
// GetIosLobAppProvisioningConfigurations gets the iosLobAppProvisioningConfigurations property value. The IOS Lob App Provisioning Configurations.
func (m *DeviceAppManagement) GetIosLobAppProvisioningConfigurations()([]IosLobAppProvisioningConfigurationable) {
    return m.iosLobAppProvisioningConfigurations
}
// GetIosManagedAppProtections gets the iosManagedAppProtections property value. iOS managed app policies.
func (m *DeviceAppManagement) GetIosManagedAppProtections()([]IosManagedAppProtectionable) {
    return m.iosManagedAppProtections
}
// GetIsEnabledForMicrosoftStoreForBusiness gets the isEnabledForMicrosoftStoreForBusiness property value. Whether the account is enabled for syncing applications from the Microsoft Store for Business.
func (m *DeviceAppManagement) GetIsEnabledForMicrosoftStoreForBusiness()(*bool) {
    return m.isEnabledForMicrosoftStoreForBusiness
}
// GetManagedAppPolicies gets the managedAppPolicies property value. Managed app policies.
func (m *DeviceAppManagement) GetManagedAppPolicies()([]ManagedAppPolicyable) {
    return m.managedAppPolicies
}
// GetManagedAppRegistrations gets the managedAppRegistrations property value. The managed app registrations.
func (m *DeviceAppManagement) GetManagedAppRegistrations()([]ManagedAppRegistrationable) {
    return m.managedAppRegistrations
}
// GetManagedAppStatuses gets the managedAppStatuses property value. The managed app statuses.
func (m *DeviceAppManagement) GetManagedAppStatuses()([]ManagedAppStatusable) {
    return m.managedAppStatuses
}
// GetManagedEBookCategories gets the managedEBookCategories property value. The mobile eBook categories.
func (m *DeviceAppManagement) GetManagedEBookCategories()([]ManagedEBookCategoryable) {
    return m.managedEBookCategories
}
// GetManagedEBooks gets the managedEBooks property value. The Managed eBook.
func (m *DeviceAppManagement) GetManagedEBooks()([]ManagedEBookable) {
    return m.managedEBooks
}
// GetMdmWindowsInformationProtectionPolicies gets the mdmWindowsInformationProtectionPolicies property value. Windows information protection for apps running on devices which are MDM enrolled.
func (m *DeviceAppManagement) GetMdmWindowsInformationProtectionPolicies()([]MdmWindowsInformationProtectionPolicyable) {
    return m.mdmWindowsInformationProtectionPolicies
}
// GetMicrosoftStoreForBusinessLanguage gets the microsoftStoreForBusinessLanguage property value. The locale information used to sync applications from the Microsoft Store for Business. Cultures that are specific to a country/region. The names of these cultures follow RFC 4646 (Windows Vista and later). The format is -<country/regioncode2>, where  is a lowercase two-letter code derived from ISO 639-1 and <country/regioncode2> is an uppercase two-letter code derived from ISO 3166. For example, en-US for English (United States) is a specific culture.
func (m *DeviceAppManagement) GetMicrosoftStoreForBusinessLanguage()(*string) {
    return m.microsoftStoreForBusinessLanguage
}
// GetMicrosoftStoreForBusinessLastCompletedApplicationSyncTime gets the microsoftStoreForBusinessLastCompletedApplicationSyncTime property value. The last time an application sync from the Microsoft Store for Business was completed.
func (m *DeviceAppManagement) GetMicrosoftStoreForBusinessLastCompletedApplicationSyncTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.microsoftStoreForBusinessLastCompletedApplicationSyncTime
}
// GetMicrosoftStoreForBusinessLastSuccessfulSyncDateTime gets the microsoftStoreForBusinessLastSuccessfulSyncDateTime property value. The last time the apps from the Microsoft Store for Business were synced successfully for the account.
func (m *DeviceAppManagement) GetMicrosoftStoreForBusinessLastSuccessfulSyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.microsoftStoreForBusinessLastSuccessfulSyncDateTime
}
// GetMicrosoftStoreForBusinessPortalSelection gets the microsoftStoreForBusinessPortalSelection property value. Portal to which admin syncs available Microsoft Store for Business apps. This is available in the Intune Admin Console.
func (m *DeviceAppManagement) GetMicrosoftStoreForBusinessPortalSelection()(*MicrosoftStoreForBusinessPortalSelectionOptions) {
    return m.microsoftStoreForBusinessPortalSelection
}
// GetMobileAppCategories gets the mobileAppCategories property value. The mobile app categories.
func (m *DeviceAppManagement) GetMobileAppCategories()([]MobileAppCategoryable) {
    return m.mobileAppCategories
}
// GetMobileAppConfigurations gets the mobileAppConfigurations property value. The Managed Device Mobile Application Configurations.
func (m *DeviceAppManagement) GetMobileAppConfigurations()([]ManagedDeviceMobileAppConfigurationable) {
    return m.mobileAppConfigurations
}
// GetMobileApps gets the mobileApps property value. The mobile apps.
func (m *DeviceAppManagement) GetMobileApps()([]MobileAppable) {
    return m.mobileApps
}
// GetPolicySets gets the policySets property value. The PolicySet of Policies and Applications
func (m *DeviceAppManagement) GetPolicySets()([]PolicySetable) {
    return m.policySets
}
// GetSideLoadingKeys gets the sideLoadingKeys property value. Side Loading Keys that are required for the Windows 8 and 8.1 Apps installation.
func (m *DeviceAppManagement) GetSideLoadingKeys()([]SideLoadingKeyable) {
    return m.sideLoadingKeys
}
// GetSymantecCodeSigningCertificate gets the symantecCodeSigningCertificate property value. The WinPhone Symantec Code Signing Certificate.
func (m *DeviceAppManagement) GetSymantecCodeSigningCertificate()(SymantecCodeSigningCertificateable) {
    return m.symantecCodeSigningCertificate
}
// GetTargetedManagedAppConfigurations gets the targetedManagedAppConfigurations property value. Targeted managed app configurations.
func (m *DeviceAppManagement) GetTargetedManagedAppConfigurations()([]TargetedManagedAppConfigurationable) {
    return m.targetedManagedAppConfigurations
}
// GetVppTokens gets the vppTokens property value. List of Vpp tokens for this organization.
func (m *DeviceAppManagement) GetVppTokens()([]VppTokenable) {
    return m.vppTokens
}
// GetWdacSupplementalPolicies gets the wdacSupplementalPolicies property value. The collection of Windows Defender Application Control Supplemental Policies.
func (m *DeviceAppManagement) GetWdacSupplementalPolicies()([]WindowsDefenderApplicationControlSupplementalPolicyable) {
    return m.wdacSupplementalPolicies
}
// GetWindowsInformationProtectionDeviceRegistrations gets the windowsInformationProtectionDeviceRegistrations property value. Windows information protection device registrations that are not MDM enrolled.
func (m *DeviceAppManagement) GetWindowsInformationProtectionDeviceRegistrations()([]WindowsInformationProtectionDeviceRegistrationable) {
    return m.windowsInformationProtectionDeviceRegistrations
}
// GetWindowsInformationProtectionPolicies gets the windowsInformationProtectionPolicies property value. Windows information protection for apps running on devices which are not MDM enrolled.
func (m *DeviceAppManagement) GetWindowsInformationProtectionPolicies()([]WindowsInformationProtectionPolicyable) {
    return m.windowsInformationProtectionPolicies
}
// GetWindowsInformationProtectionWipeActions gets the windowsInformationProtectionWipeActions property value. Windows information protection wipe actions.
func (m *DeviceAppManagement) GetWindowsInformationProtectionWipeActions()([]WindowsInformationProtectionWipeActionable) {
    return m.windowsInformationProtectionWipeActions
}
// GetWindowsManagedAppProtections gets the windowsManagedAppProtections property value. Windows managed app policies.
func (m *DeviceAppManagement) GetWindowsManagedAppProtections()([]WindowsManagedAppProtectionable) {
    return m.windowsManagedAppProtections
}
// GetWindowsManagementApp gets the windowsManagementApp property value. Windows management app.
func (m *DeviceAppManagement) GetWindowsManagementApp()(WindowsManagementAppable) {
    return m.windowsManagementApp
}
// Serialize serializes information the current object
func (m *DeviceAppManagement) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAndroidManagedAppProtections() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetAndroidManagedAppProtections())
        err = writer.WriteCollectionOfObjectValues("androidManagedAppProtections", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDefaultManagedAppProtections() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetDefaultManagedAppProtections())
        err = writer.WriteCollectionOfObjectValues("defaultManagedAppProtections", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDeviceAppManagementTasks() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetDeviceAppManagementTasks())
        err = writer.WriteCollectionOfObjectValues("deviceAppManagementTasks", cast)
        if err != nil {
            return err
        }
    }
    if m.GetEnterpriseCodeSigningCertificates() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetEnterpriseCodeSigningCertificates())
        err = writer.WriteCollectionOfObjectValues("enterpriseCodeSigningCertificates", cast)
        if err != nil {
            return err
        }
    }
    if m.GetIosLobAppProvisioningConfigurations() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetIosLobAppProvisioningConfigurations())
        err = writer.WriteCollectionOfObjectValues("iosLobAppProvisioningConfigurations", cast)
        if err != nil {
            return err
        }
    }
    if m.GetIosManagedAppProtections() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetIosManagedAppProtections())
        err = writer.WriteCollectionOfObjectValues("iosManagedAppProtections", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isEnabledForMicrosoftStoreForBusiness", m.GetIsEnabledForMicrosoftStoreForBusiness())
        if err != nil {
            return err
        }
    }
    if m.GetManagedAppPolicies() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetManagedAppPolicies())
        err = writer.WriteCollectionOfObjectValues("managedAppPolicies", cast)
        if err != nil {
            return err
        }
    }
    if m.GetManagedAppRegistrations() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetManagedAppRegistrations())
        err = writer.WriteCollectionOfObjectValues("managedAppRegistrations", cast)
        if err != nil {
            return err
        }
    }
    if m.GetManagedAppStatuses() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetManagedAppStatuses())
        err = writer.WriteCollectionOfObjectValues("managedAppStatuses", cast)
        if err != nil {
            return err
        }
    }
    if m.GetManagedEBookCategories() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetManagedEBookCategories())
        err = writer.WriteCollectionOfObjectValues("managedEBookCategories", cast)
        if err != nil {
            return err
        }
    }
    if m.GetManagedEBooks() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetManagedEBooks())
        err = writer.WriteCollectionOfObjectValues("managedEBooks", cast)
        if err != nil {
            return err
        }
    }
    if m.GetMdmWindowsInformationProtectionPolicies() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetMdmWindowsInformationProtectionPolicies())
        err = writer.WriteCollectionOfObjectValues("mdmWindowsInformationProtectionPolicies", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("microsoftStoreForBusinessLanguage", m.GetMicrosoftStoreForBusinessLanguage())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("microsoftStoreForBusinessLastCompletedApplicationSyncTime", m.GetMicrosoftStoreForBusinessLastCompletedApplicationSyncTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("microsoftStoreForBusinessLastSuccessfulSyncDateTime", m.GetMicrosoftStoreForBusinessLastSuccessfulSyncDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetMicrosoftStoreForBusinessPortalSelection() != nil {
        cast := (*m.GetMicrosoftStoreForBusinessPortalSelection()).String()
        err = writer.WriteStringValue("microsoftStoreForBusinessPortalSelection", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetMobileAppCategories() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetMobileAppCategories())
        err = writer.WriteCollectionOfObjectValues("mobileAppCategories", cast)
        if err != nil {
            return err
        }
    }
    if m.GetMobileAppConfigurations() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetMobileAppConfigurations())
        err = writer.WriteCollectionOfObjectValues("mobileAppConfigurations", cast)
        if err != nil {
            return err
        }
    }
    if m.GetMobileApps() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetMobileApps())
        err = writer.WriteCollectionOfObjectValues("mobileApps", cast)
        if err != nil {
            return err
        }
    }
    if m.GetPolicySets() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetPolicySets())
        err = writer.WriteCollectionOfObjectValues("policySets", cast)
        if err != nil {
            return err
        }
    }
    if m.GetSideLoadingKeys() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetSideLoadingKeys())
        err = writer.WriteCollectionOfObjectValues("sideLoadingKeys", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("symantecCodeSigningCertificate", m.GetSymantecCodeSigningCertificate())
        if err != nil {
            return err
        }
    }
    if m.GetTargetedManagedAppConfigurations() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetTargetedManagedAppConfigurations())
        err = writer.WriteCollectionOfObjectValues("targetedManagedAppConfigurations", cast)
        if err != nil {
            return err
        }
    }
    if m.GetVppTokens() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetVppTokens())
        err = writer.WriteCollectionOfObjectValues("vppTokens", cast)
        if err != nil {
            return err
        }
    }
    if m.GetWdacSupplementalPolicies() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetWdacSupplementalPolicies())
        err = writer.WriteCollectionOfObjectValues("wdacSupplementalPolicies", cast)
        if err != nil {
            return err
        }
    }
    if m.GetWindowsInformationProtectionDeviceRegistrations() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetWindowsInformationProtectionDeviceRegistrations())
        err = writer.WriteCollectionOfObjectValues("windowsInformationProtectionDeviceRegistrations", cast)
        if err != nil {
            return err
        }
    }
    if m.GetWindowsInformationProtectionPolicies() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetWindowsInformationProtectionPolicies())
        err = writer.WriteCollectionOfObjectValues("windowsInformationProtectionPolicies", cast)
        if err != nil {
            return err
        }
    }
    if m.GetWindowsInformationProtectionWipeActions() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetWindowsInformationProtectionWipeActions())
        err = writer.WriteCollectionOfObjectValues("windowsInformationProtectionWipeActions", cast)
        if err != nil {
            return err
        }
    }
    if m.GetWindowsManagedAppProtections() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetWindowsManagedAppProtections())
        err = writer.WriteCollectionOfObjectValues("windowsManagedAppProtections", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("windowsManagementApp", m.GetWindowsManagementApp())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAndroidManagedAppProtections sets the androidManagedAppProtections property value. Android managed app policies.
func (m *DeviceAppManagement) SetAndroidManagedAppProtections(value []AndroidManagedAppProtectionable)() {
    m.androidManagedAppProtections = value
}
// SetDefaultManagedAppProtections sets the defaultManagedAppProtections property value. Default managed app policies.
func (m *DeviceAppManagement) SetDefaultManagedAppProtections(value []DefaultManagedAppProtectionable)() {
    m.defaultManagedAppProtections = value
}
// SetDeviceAppManagementTasks sets the deviceAppManagementTasks property value. Device app management tasks.
func (m *DeviceAppManagement) SetDeviceAppManagementTasks(value []DeviceAppManagementTaskable)() {
    m.deviceAppManagementTasks = value
}
// SetEnterpriseCodeSigningCertificates sets the enterpriseCodeSigningCertificates property value. The Windows Enterprise Code Signing Certificate.
func (m *DeviceAppManagement) SetEnterpriseCodeSigningCertificates(value []EnterpriseCodeSigningCertificateable)() {
    m.enterpriseCodeSigningCertificates = value
}
// SetIosLobAppProvisioningConfigurations sets the iosLobAppProvisioningConfigurations property value. The IOS Lob App Provisioning Configurations.
func (m *DeviceAppManagement) SetIosLobAppProvisioningConfigurations(value []IosLobAppProvisioningConfigurationable)() {
    m.iosLobAppProvisioningConfigurations = value
}
// SetIosManagedAppProtections sets the iosManagedAppProtections property value. iOS managed app policies.
func (m *DeviceAppManagement) SetIosManagedAppProtections(value []IosManagedAppProtectionable)() {
    m.iosManagedAppProtections = value
}
// SetIsEnabledForMicrosoftStoreForBusiness sets the isEnabledForMicrosoftStoreForBusiness property value. Whether the account is enabled for syncing applications from the Microsoft Store for Business.
func (m *DeviceAppManagement) SetIsEnabledForMicrosoftStoreForBusiness(value *bool)() {
    m.isEnabledForMicrosoftStoreForBusiness = value
}
// SetManagedAppPolicies sets the managedAppPolicies property value. Managed app policies.
func (m *DeviceAppManagement) SetManagedAppPolicies(value []ManagedAppPolicyable)() {
    m.managedAppPolicies = value
}
// SetManagedAppRegistrations sets the managedAppRegistrations property value. The managed app registrations.
func (m *DeviceAppManagement) SetManagedAppRegistrations(value []ManagedAppRegistrationable)() {
    m.managedAppRegistrations = value
}
// SetManagedAppStatuses sets the managedAppStatuses property value. The managed app statuses.
func (m *DeviceAppManagement) SetManagedAppStatuses(value []ManagedAppStatusable)() {
    m.managedAppStatuses = value
}
// SetManagedEBookCategories sets the managedEBookCategories property value. The mobile eBook categories.
func (m *DeviceAppManagement) SetManagedEBookCategories(value []ManagedEBookCategoryable)() {
    m.managedEBookCategories = value
}
// SetManagedEBooks sets the managedEBooks property value. The Managed eBook.
func (m *DeviceAppManagement) SetManagedEBooks(value []ManagedEBookable)() {
    m.managedEBooks = value
}
// SetMdmWindowsInformationProtectionPolicies sets the mdmWindowsInformationProtectionPolicies property value. Windows information protection for apps running on devices which are MDM enrolled.
func (m *DeviceAppManagement) SetMdmWindowsInformationProtectionPolicies(value []MdmWindowsInformationProtectionPolicyable)() {
    m.mdmWindowsInformationProtectionPolicies = value
}
// SetMicrosoftStoreForBusinessLanguage sets the microsoftStoreForBusinessLanguage property value. The locale information used to sync applications from the Microsoft Store for Business. Cultures that are specific to a country/region. The names of these cultures follow RFC 4646 (Windows Vista and later). The format is -<country/regioncode2>, where  is a lowercase two-letter code derived from ISO 639-1 and <country/regioncode2> is an uppercase two-letter code derived from ISO 3166. For example, en-US for English (United States) is a specific culture.
func (m *DeviceAppManagement) SetMicrosoftStoreForBusinessLanguage(value *string)() {
    m.microsoftStoreForBusinessLanguage = value
}
// SetMicrosoftStoreForBusinessLastCompletedApplicationSyncTime sets the microsoftStoreForBusinessLastCompletedApplicationSyncTime property value. The last time an application sync from the Microsoft Store for Business was completed.
func (m *DeviceAppManagement) SetMicrosoftStoreForBusinessLastCompletedApplicationSyncTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.microsoftStoreForBusinessLastCompletedApplicationSyncTime = value
}
// SetMicrosoftStoreForBusinessLastSuccessfulSyncDateTime sets the microsoftStoreForBusinessLastSuccessfulSyncDateTime property value. The last time the apps from the Microsoft Store for Business were synced successfully for the account.
func (m *DeviceAppManagement) SetMicrosoftStoreForBusinessLastSuccessfulSyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.microsoftStoreForBusinessLastSuccessfulSyncDateTime = value
}
// SetMicrosoftStoreForBusinessPortalSelection sets the microsoftStoreForBusinessPortalSelection property value. Portal to which admin syncs available Microsoft Store for Business apps. This is available in the Intune Admin Console.
func (m *DeviceAppManagement) SetMicrosoftStoreForBusinessPortalSelection(value *MicrosoftStoreForBusinessPortalSelectionOptions)() {
    m.microsoftStoreForBusinessPortalSelection = value
}
// SetMobileAppCategories sets the mobileAppCategories property value. The mobile app categories.
func (m *DeviceAppManagement) SetMobileAppCategories(value []MobileAppCategoryable)() {
    m.mobileAppCategories = value
}
// SetMobileAppConfigurations sets the mobileAppConfigurations property value. The Managed Device Mobile Application Configurations.
func (m *DeviceAppManagement) SetMobileAppConfigurations(value []ManagedDeviceMobileAppConfigurationable)() {
    m.mobileAppConfigurations = value
}
// SetMobileApps sets the mobileApps property value. The mobile apps.
func (m *DeviceAppManagement) SetMobileApps(value []MobileAppable)() {
    m.mobileApps = value
}
// SetPolicySets sets the policySets property value. The PolicySet of Policies and Applications
func (m *DeviceAppManagement) SetPolicySets(value []PolicySetable)() {
    m.policySets = value
}
// SetSideLoadingKeys sets the sideLoadingKeys property value. Side Loading Keys that are required for the Windows 8 and 8.1 Apps installation.
func (m *DeviceAppManagement) SetSideLoadingKeys(value []SideLoadingKeyable)() {
    m.sideLoadingKeys = value
}
// SetSymantecCodeSigningCertificate sets the symantecCodeSigningCertificate property value. The WinPhone Symantec Code Signing Certificate.
func (m *DeviceAppManagement) SetSymantecCodeSigningCertificate(value SymantecCodeSigningCertificateable)() {
    m.symantecCodeSigningCertificate = value
}
// SetTargetedManagedAppConfigurations sets the targetedManagedAppConfigurations property value. Targeted managed app configurations.
func (m *DeviceAppManagement) SetTargetedManagedAppConfigurations(value []TargetedManagedAppConfigurationable)() {
    m.targetedManagedAppConfigurations = value
}
// SetVppTokens sets the vppTokens property value. List of Vpp tokens for this organization.
func (m *DeviceAppManagement) SetVppTokens(value []VppTokenable)() {
    m.vppTokens = value
}
// SetWdacSupplementalPolicies sets the wdacSupplementalPolicies property value. The collection of Windows Defender Application Control Supplemental Policies.
func (m *DeviceAppManagement) SetWdacSupplementalPolicies(value []WindowsDefenderApplicationControlSupplementalPolicyable)() {
    m.wdacSupplementalPolicies = value
}
// SetWindowsInformationProtectionDeviceRegistrations sets the windowsInformationProtectionDeviceRegistrations property value. Windows information protection device registrations that are not MDM enrolled.
func (m *DeviceAppManagement) SetWindowsInformationProtectionDeviceRegistrations(value []WindowsInformationProtectionDeviceRegistrationable)() {
    m.windowsInformationProtectionDeviceRegistrations = value
}
// SetWindowsInformationProtectionPolicies sets the windowsInformationProtectionPolicies property value. Windows information protection for apps running on devices which are not MDM enrolled.
func (m *DeviceAppManagement) SetWindowsInformationProtectionPolicies(value []WindowsInformationProtectionPolicyable)() {
    m.windowsInformationProtectionPolicies = value
}
// SetWindowsInformationProtectionWipeActions sets the windowsInformationProtectionWipeActions property value. Windows information protection wipe actions.
func (m *DeviceAppManagement) SetWindowsInformationProtectionWipeActions(value []WindowsInformationProtectionWipeActionable)() {
    m.windowsInformationProtectionWipeActions = value
}
// SetWindowsManagedAppProtections sets the windowsManagedAppProtections property value. Windows managed app policies.
func (m *DeviceAppManagement) SetWindowsManagedAppProtections(value []WindowsManagedAppProtectionable)() {
    m.windowsManagedAppProtections = value
}
// SetWindowsManagementApp sets the windowsManagementApp property value. Windows management app.
func (m *DeviceAppManagement) SetWindowsManagementApp(value WindowsManagementAppable)() {
    m.windowsManagementApp = value
}
