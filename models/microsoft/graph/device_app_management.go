package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type DeviceAppManagement struct {
    Entity
    // Android managed app policies.
    androidManagedAppProtections []AndroidManagedAppProtection;
    // Default managed app policies.
    defaultManagedAppProtections []DefaultManagedAppProtection;
    // Device app management tasks.
    deviceAppManagementTasks []DeviceAppManagementTask;
    // The Windows Enterprise Code Signing Certificate.
    enterpriseCodeSigningCertificates []EnterpriseCodeSigningCertificate;
    // The IOS Lob App Provisioning Configurations.
    iosLobAppProvisioningConfigurations []IosLobAppProvisioningConfiguration;
    // iOS managed app policies.
    iosManagedAppProtections []IosManagedAppProtection;
    // Whether the account is enabled for syncing applications from the Microsoft Store for Business.
    isEnabledForMicrosoftStoreForBusiness *bool;
    // Managed app policies.
    managedAppPolicies []ManagedAppPolicy;
    // The managed app registrations.
    managedAppRegistrations []ManagedAppRegistration;
    // The managed app statuses.
    managedAppStatuses []ManagedAppStatus;
    // The mobile eBook categories.
    managedEBookCategories []ManagedEBookCategory;
    // The Managed eBook.
    managedEBooks []ManagedEBook;
    // Windows information protection for apps running on devices which are MDM enrolled.
    mdmWindowsInformationProtectionPolicies []MdmWindowsInformationProtectionPolicy;
    // The locale information used to sync applications from the Microsoft Store for Business. Cultures that are specific to a country/region. The names of these cultures follow RFC 4646 (Windows Vista and later). The format is -<country/regioncode2>, where  is a lowercase two-letter code derived from ISO 639-1 and <country/regioncode2> is an uppercase two-letter code derived from ISO 3166. For example, en-US for English (United States) is a specific culture.
    microsoftStoreForBusinessLanguage *string;
    // The last time an application sync from the Microsoft Store for Business was completed.
    microsoftStoreForBusinessLastCompletedApplicationSyncTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The last time the apps from the Microsoft Store for Business were synced successfully for the account.
    microsoftStoreForBusinessLastSuccessfulSyncDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The end user portal information is used to sync applications from the Microsoft Store for Business to Intune Company Portal. There are three options to pick from ['Company portal only', 'Company portal and private store', 'Private store only']. Possible values are: none, companyPortal, privateStore.
    microsoftStoreForBusinessPortalSelection *MicrosoftStoreForBusinessPortalSelectionOptions;
    // The mobile app categories.
    mobileAppCategories []MobileAppCategory;
    // The Managed Device Mobile Application Configurations.
    mobileAppConfigurations []ManagedDeviceMobileAppConfiguration;
    // The mobile apps.
    mobileApps []MobileApp;
    // The PolicySet of Policies and Applications
    policySets []PolicySet;
    // Side Loading Keys that are required for the Windows 8 and 8.1 Apps installation.
    sideLoadingKeys []SideLoadingKey;
    // The WinPhone Symantec Code Signing Certificate.
    symantecCodeSigningCertificate *SymantecCodeSigningCertificate;
    // Targeted managed app configurations.
    targetedManagedAppConfigurations []TargetedManagedAppConfiguration;
    // List of Vpp tokens for this organization.
    vppTokens []VppToken;
    // The collection of Windows Defender Application Control Supplemental Policies.
    wdacSupplementalPolicies []WindowsDefenderApplicationControlSupplementalPolicy;
    // Windows information protection device registrations that are not MDM enrolled.
    windowsInformationProtectionDeviceRegistrations []WindowsInformationProtectionDeviceRegistration;
    // Windows information protection for apps running on devices which are not MDM enrolled.
    windowsInformationProtectionPolicies []WindowsInformationProtectionPolicy;
    // Windows information protection wipe actions.
    windowsInformationProtectionWipeActions []WindowsInformationProtectionWipeAction;
    // Windows management app.
    windowsManagementApp *WindowsManagementApp;
}
// Instantiates a new deviceAppManagement and sets the default values.
func NewDeviceAppManagement()(*DeviceAppManagement) {
    m := &DeviceAppManagement{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the androidManagedAppProtections property value. Android managed app policies.
func (m *DeviceAppManagement) GetAndroidManagedAppProtections()([]AndroidManagedAppProtection) {
    if m == nil {
        return nil
    } else {
        return m.androidManagedAppProtections
    }
}
// Gets the defaultManagedAppProtections property value. Default managed app policies.
func (m *DeviceAppManagement) GetDefaultManagedAppProtections()([]DefaultManagedAppProtection) {
    if m == nil {
        return nil
    } else {
        return m.defaultManagedAppProtections
    }
}
// Gets the deviceAppManagementTasks property value. Device app management tasks.
func (m *DeviceAppManagement) GetDeviceAppManagementTasks()([]DeviceAppManagementTask) {
    if m == nil {
        return nil
    } else {
        return m.deviceAppManagementTasks
    }
}
// Gets the enterpriseCodeSigningCertificates property value. The Windows Enterprise Code Signing Certificate.
func (m *DeviceAppManagement) GetEnterpriseCodeSigningCertificates()([]EnterpriseCodeSigningCertificate) {
    if m == nil {
        return nil
    } else {
        return m.enterpriseCodeSigningCertificates
    }
}
// Gets the iosLobAppProvisioningConfigurations property value. The IOS Lob App Provisioning Configurations.
func (m *DeviceAppManagement) GetIosLobAppProvisioningConfigurations()([]IosLobAppProvisioningConfiguration) {
    if m == nil {
        return nil
    } else {
        return m.iosLobAppProvisioningConfigurations
    }
}
// Gets the iosManagedAppProtections property value. iOS managed app policies.
func (m *DeviceAppManagement) GetIosManagedAppProtections()([]IosManagedAppProtection) {
    if m == nil {
        return nil
    } else {
        return m.iosManagedAppProtections
    }
}
// Gets the isEnabledForMicrosoftStoreForBusiness property value. Whether the account is enabled for syncing applications from the Microsoft Store for Business.
func (m *DeviceAppManagement) GetIsEnabledForMicrosoftStoreForBusiness()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isEnabledForMicrosoftStoreForBusiness
    }
}
// Gets the managedAppPolicies property value. Managed app policies.
func (m *DeviceAppManagement) GetManagedAppPolicies()([]ManagedAppPolicy) {
    if m == nil {
        return nil
    } else {
        return m.managedAppPolicies
    }
}
// Gets the managedAppRegistrations property value. The managed app registrations.
func (m *DeviceAppManagement) GetManagedAppRegistrations()([]ManagedAppRegistration) {
    if m == nil {
        return nil
    } else {
        return m.managedAppRegistrations
    }
}
// Gets the managedAppStatuses property value. The managed app statuses.
func (m *DeviceAppManagement) GetManagedAppStatuses()([]ManagedAppStatus) {
    if m == nil {
        return nil
    } else {
        return m.managedAppStatuses
    }
}
// Gets the managedEBookCategories property value. The mobile eBook categories.
func (m *DeviceAppManagement) GetManagedEBookCategories()([]ManagedEBookCategory) {
    if m == nil {
        return nil
    } else {
        return m.managedEBookCategories
    }
}
// Gets the managedEBooks property value. The Managed eBook.
func (m *DeviceAppManagement) GetManagedEBooks()([]ManagedEBook) {
    if m == nil {
        return nil
    } else {
        return m.managedEBooks
    }
}
// Gets the mdmWindowsInformationProtectionPolicies property value. Windows information protection for apps running on devices which are MDM enrolled.
func (m *DeviceAppManagement) GetMdmWindowsInformationProtectionPolicies()([]MdmWindowsInformationProtectionPolicy) {
    if m == nil {
        return nil
    } else {
        return m.mdmWindowsInformationProtectionPolicies
    }
}
// Gets the microsoftStoreForBusinessLanguage property value. The locale information used to sync applications from the Microsoft Store for Business. Cultures that are specific to a country/region. The names of these cultures follow RFC 4646 (Windows Vista and later). The format is -<country/regioncode2>, where  is a lowercase two-letter code derived from ISO 639-1 and <country/regioncode2> is an uppercase two-letter code derived from ISO 3166. For example, en-US for English (United States) is a specific culture.
func (m *DeviceAppManagement) GetMicrosoftStoreForBusinessLanguage()(*string) {
    if m == nil {
        return nil
    } else {
        return m.microsoftStoreForBusinessLanguage
    }
}
// Gets the microsoftStoreForBusinessLastCompletedApplicationSyncTime property value. The last time an application sync from the Microsoft Store for Business was completed.
func (m *DeviceAppManagement) GetMicrosoftStoreForBusinessLastCompletedApplicationSyncTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.microsoftStoreForBusinessLastCompletedApplicationSyncTime
    }
}
// Gets the microsoftStoreForBusinessLastSuccessfulSyncDateTime property value. The last time the apps from the Microsoft Store for Business were synced successfully for the account.
func (m *DeviceAppManagement) GetMicrosoftStoreForBusinessLastSuccessfulSyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.microsoftStoreForBusinessLastSuccessfulSyncDateTime
    }
}
// Gets the microsoftStoreForBusinessPortalSelection property value. The end user portal information is used to sync applications from the Microsoft Store for Business to Intune Company Portal. There are three options to pick from ['Company portal only', 'Company portal and private store', 'Private store only']. Possible values are: none, companyPortal, privateStore.
func (m *DeviceAppManagement) GetMicrosoftStoreForBusinessPortalSelection()(*MicrosoftStoreForBusinessPortalSelectionOptions) {
    if m == nil {
        return nil
    } else {
        return m.microsoftStoreForBusinessPortalSelection
    }
}
// Gets the mobileAppCategories property value. The mobile app categories.
func (m *DeviceAppManagement) GetMobileAppCategories()([]MobileAppCategory) {
    if m == nil {
        return nil
    } else {
        return m.mobileAppCategories
    }
}
// Gets the mobileAppConfigurations property value. The Managed Device Mobile Application Configurations.
func (m *DeviceAppManagement) GetMobileAppConfigurations()([]ManagedDeviceMobileAppConfiguration) {
    if m == nil {
        return nil
    } else {
        return m.mobileAppConfigurations
    }
}
// Gets the mobileApps property value. The mobile apps.
func (m *DeviceAppManagement) GetMobileApps()([]MobileApp) {
    if m == nil {
        return nil
    } else {
        return m.mobileApps
    }
}
// Gets the policySets property value. The PolicySet of Policies and Applications
func (m *DeviceAppManagement) GetPolicySets()([]PolicySet) {
    if m == nil {
        return nil
    } else {
        return m.policySets
    }
}
// Gets the sideLoadingKeys property value. Side Loading Keys that are required for the Windows 8 and 8.1 Apps installation.
func (m *DeviceAppManagement) GetSideLoadingKeys()([]SideLoadingKey) {
    if m == nil {
        return nil
    } else {
        return m.sideLoadingKeys
    }
}
// Gets the symantecCodeSigningCertificate property value. The WinPhone Symantec Code Signing Certificate.
func (m *DeviceAppManagement) GetSymantecCodeSigningCertificate()(*SymantecCodeSigningCertificate) {
    if m == nil {
        return nil
    } else {
        return m.symantecCodeSigningCertificate
    }
}
// Gets the targetedManagedAppConfigurations property value. Targeted managed app configurations.
func (m *DeviceAppManagement) GetTargetedManagedAppConfigurations()([]TargetedManagedAppConfiguration) {
    if m == nil {
        return nil
    } else {
        return m.targetedManagedAppConfigurations
    }
}
// Gets the vppTokens property value. List of Vpp tokens for this organization.
func (m *DeviceAppManagement) GetVppTokens()([]VppToken) {
    if m == nil {
        return nil
    } else {
        return m.vppTokens
    }
}
// Gets the wdacSupplementalPolicies property value. The collection of Windows Defender Application Control Supplemental Policies.
func (m *DeviceAppManagement) GetWdacSupplementalPolicies()([]WindowsDefenderApplicationControlSupplementalPolicy) {
    if m == nil {
        return nil
    } else {
        return m.wdacSupplementalPolicies
    }
}
// Gets the windowsInformationProtectionDeviceRegistrations property value. Windows information protection device registrations that are not MDM enrolled.
func (m *DeviceAppManagement) GetWindowsInformationProtectionDeviceRegistrations()([]WindowsInformationProtectionDeviceRegistration) {
    if m == nil {
        return nil
    } else {
        return m.windowsInformationProtectionDeviceRegistrations
    }
}
// Gets the windowsInformationProtectionPolicies property value. Windows information protection for apps running on devices which are not MDM enrolled.
func (m *DeviceAppManagement) GetWindowsInformationProtectionPolicies()([]WindowsInformationProtectionPolicy) {
    if m == nil {
        return nil
    } else {
        return m.windowsInformationProtectionPolicies
    }
}
// Gets the windowsInformationProtectionWipeActions property value. Windows information protection wipe actions.
func (m *DeviceAppManagement) GetWindowsInformationProtectionWipeActions()([]WindowsInformationProtectionWipeAction) {
    if m == nil {
        return nil
    } else {
        return m.windowsInformationProtectionWipeActions
    }
}
// Gets the windowsManagementApp property value. Windows management app.
func (m *DeviceAppManagement) GetWindowsManagementApp()(*WindowsManagementApp) {
    if m == nil {
        return nil
    } else {
        return m.windowsManagementApp
    }
}
// The deserialization information for the current model
func (m *DeviceAppManagement) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["androidManagedAppProtections"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAndroidManagedAppProtection() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AndroidManagedAppProtection, len(val))
            for i, v := range val {
                res[i] = *(v.(*AndroidManagedAppProtection))
            }
            m.SetAndroidManagedAppProtections(res)
        }
        return nil
    }
    res["defaultManagedAppProtections"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDefaultManagedAppProtection() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DefaultManagedAppProtection, len(val))
            for i, v := range val {
                res[i] = *(v.(*DefaultManagedAppProtection))
            }
            m.SetDefaultManagedAppProtections(res)
        }
        return nil
    }
    res["deviceAppManagementTasks"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceAppManagementTask() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceAppManagementTask, len(val))
            for i, v := range val {
                res[i] = *(v.(*DeviceAppManagementTask))
            }
            m.SetDeviceAppManagementTasks(res)
        }
        return nil
    }
    res["enterpriseCodeSigningCertificates"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewEnterpriseCodeSigningCertificate() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]EnterpriseCodeSigningCertificate, len(val))
            for i, v := range val {
                res[i] = *(v.(*EnterpriseCodeSigningCertificate))
            }
            m.SetEnterpriseCodeSigningCertificates(res)
        }
        return nil
    }
    res["iosLobAppProvisioningConfigurations"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewIosLobAppProvisioningConfiguration() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]IosLobAppProvisioningConfiguration, len(val))
            for i, v := range val {
                res[i] = *(v.(*IosLobAppProvisioningConfiguration))
            }
            m.SetIosLobAppProvisioningConfigurations(res)
        }
        return nil
    }
    res["iosManagedAppProtections"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewIosManagedAppProtection() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]IosManagedAppProtection, len(val))
            for i, v := range val {
                res[i] = *(v.(*IosManagedAppProtection))
            }
            m.SetIosManagedAppProtections(res)
        }
        return nil
    }
    res["isEnabledForMicrosoftStoreForBusiness"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsEnabledForMicrosoftStoreForBusiness(val)
        }
        return nil
    }
    res["managedAppPolicies"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewManagedAppPolicy() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagedAppPolicy, len(val))
            for i, v := range val {
                res[i] = *(v.(*ManagedAppPolicy))
            }
            m.SetManagedAppPolicies(res)
        }
        return nil
    }
    res["managedAppRegistrations"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewManagedAppRegistration() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagedAppRegistration, len(val))
            for i, v := range val {
                res[i] = *(v.(*ManagedAppRegistration))
            }
            m.SetManagedAppRegistrations(res)
        }
        return nil
    }
    res["managedAppStatuses"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewManagedAppStatus() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagedAppStatus, len(val))
            for i, v := range val {
                res[i] = *(v.(*ManagedAppStatus))
            }
            m.SetManagedAppStatuses(res)
        }
        return nil
    }
    res["managedEBookCategories"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewManagedEBookCategory() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagedEBookCategory, len(val))
            for i, v := range val {
                res[i] = *(v.(*ManagedEBookCategory))
            }
            m.SetManagedEBookCategories(res)
        }
        return nil
    }
    res["managedEBooks"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewManagedEBook() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagedEBook, len(val))
            for i, v := range val {
                res[i] = *(v.(*ManagedEBook))
            }
            m.SetManagedEBooks(res)
        }
        return nil
    }
    res["mdmWindowsInformationProtectionPolicies"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewMdmWindowsInformationProtectionPolicy() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MdmWindowsInformationProtectionPolicy, len(val))
            for i, v := range val {
                res[i] = *(v.(*MdmWindowsInformationProtectionPolicy))
            }
            m.SetMdmWindowsInformationProtectionPolicies(res)
        }
        return nil
    }
    res["microsoftStoreForBusinessLanguage"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMicrosoftStoreForBusinessLanguage(val)
        }
        return nil
    }
    res["microsoftStoreForBusinessLastCompletedApplicationSyncTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMicrosoftStoreForBusinessLastCompletedApplicationSyncTime(val)
        }
        return nil
    }
    res["microsoftStoreForBusinessLastSuccessfulSyncDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMicrosoftStoreForBusinessLastSuccessfulSyncDateTime(val)
        }
        return nil
    }
    res["microsoftStoreForBusinessPortalSelection"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseMicrosoftStoreForBusinessPortalSelectionOptions)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(MicrosoftStoreForBusinessPortalSelectionOptions)
            m.SetMicrosoftStoreForBusinessPortalSelection(&cast)
        }
        return nil
    }
    res["mobileAppCategories"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewMobileAppCategory() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MobileAppCategory, len(val))
            for i, v := range val {
                res[i] = *(v.(*MobileAppCategory))
            }
            m.SetMobileAppCategories(res)
        }
        return nil
    }
    res["mobileAppConfigurations"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewManagedDeviceMobileAppConfiguration() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagedDeviceMobileAppConfiguration, len(val))
            for i, v := range val {
                res[i] = *(v.(*ManagedDeviceMobileAppConfiguration))
            }
            m.SetMobileAppConfigurations(res)
        }
        return nil
    }
    res["mobileApps"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewMobileApp() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MobileApp, len(val))
            for i, v := range val {
                res[i] = *(v.(*MobileApp))
            }
            m.SetMobileApps(res)
        }
        return nil
    }
    res["policySets"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPolicySet() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PolicySet, len(val))
            for i, v := range val {
                res[i] = *(v.(*PolicySet))
            }
            m.SetPolicySets(res)
        }
        return nil
    }
    res["sideLoadingKeys"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSideLoadingKey() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SideLoadingKey, len(val))
            for i, v := range val {
                res[i] = *(v.(*SideLoadingKey))
            }
            m.SetSideLoadingKeys(res)
        }
        return nil
    }
    res["symantecCodeSigningCertificate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSymantecCodeSigningCertificate() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSymantecCodeSigningCertificate(val.(*SymantecCodeSigningCertificate))
        }
        return nil
    }
    res["targetedManagedAppConfigurations"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTargetedManagedAppConfiguration() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TargetedManagedAppConfiguration, len(val))
            for i, v := range val {
                res[i] = *(v.(*TargetedManagedAppConfiguration))
            }
            m.SetTargetedManagedAppConfigurations(res)
        }
        return nil
    }
    res["vppTokens"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewVppToken() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]VppToken, len(val))
            for i, v := range val {
                res[i] = *(v.(*VppToken))
            }
            m.SetVppTokens(res)
        }
        return nil
    }
    res["wdacSupplementalPolicies"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWindowsDefenderApplicationControlSupplementalPolicy() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WindowsDefenderApplicationControlSupplementalPolicy, len(val))
            for i, v := range val {
                res[i] = *(v.(*WindowsDefenderApplicationControlSupplementalPolicy))
            }
            m.SetWdacSupplementalPolicies(res)
        }
        return nil
    }
    res["windowsInformationProtectionDeviceRegistrations"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWindowsInformationProtectionDeviceRegistration() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WindowsInformationProtectionDeviceRegistration, len(val))
            for i, v := range val {
                res[i] = *(v.(*WindowsInformationProtectionDeviceRegistration))
            }
            m.SetWindowsInformationProtectionDeviceRegistrations(res)
        }
        return nil
    }
    res["windowsInformationProtectionPolicies"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWindowsInformationProtectionPolicy() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WindowsInformationProtectionPolicy, len(val))
            for i, v := range val {
                res[i] = *(v.(*WindowsInformationProtectionPolicy))
            }
            m.SetWindowsInformationProtectionPolicies(res)
        }
        return nil
    }
    res["windowsInformationProtectionWipeActions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWindowsInformationProtectionWipeAction() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WindowsInformationProtectionWipeAction, len(val))
            for i, v := range val {
                res[i] = *(v.(*WindowsInformationProtectionWipeAction))
            }
            m.SetWindowsInformationProtectionWipeActions(res)
        }
        return nil
    }
    res["windowsManagementApp"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWindowsManagementApp() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWindowsManagementApp(val.(*WindowsManagementApp))
        }
        return nil
    }
    return res
}
func (m *DeviceAppManagement) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *DeviceAppManagement) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAndroidManagedAppProtections()))
        for i, v := range m.GetAndroidManagedAppProtections() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("androidManagedAppProtections", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetDefaultManagedAppProtections()))
        for i, v := range m.GetDefaultManagedAppProtections() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("defaultManagedAppProtections", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetDeviceAppManagementTasks()))
        for i, v := range m.GetDeviceAppManagementTasks() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("deviceAppManagementTasks", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetEnterpriseCodeSigningCertificates()))
        for i, v := range m.GetEnterpriseCodeSigningCertificates() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("enterpriseCodeSigningCertificates", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetIosLobAppProvisioningConfigurations()))
        for i, v := range m.GetIosLobAppProvisioningConfigurations() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("iosLobAppProvisioningConfigurations", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetIosManagedAppProtections()))
        for i, v := range m.GetIosManagedAppProtections() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
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
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetManagedAppPolicies()))
        for i, v := range m.GetManagedAppPolicies() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("managedAppPolicies", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetManagedAppRegistrations()))
        for i, v := range m.GetManagedAppRegistrations() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("managedAppRegistrations", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetManagedAppStatuses()))
        for i, v := range m.GetManagedAppStatuses() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("managedAppStatuses", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetManagedEBookCategories()))
        for i, v := range m.GetManagedEBookCategories() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("managedEBookCategories", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetManagedEBooks()))
        for i, v := range m.GetManagedEBooks() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("managedEBooks", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetMdmWindowsInformationProtectionPolicies()))
        for i, v := range m.GetMdmWindowsInformationProtectionPolicies() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
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
        cast := m.GetMicrosoftStoreForBusinessPortalSelection().String()
        err = writer.WriteStringValue("microsoftStoreForBusinessPortalSelection", &cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetMobileAppCategories()))
        for i, v := range m.GetMobileAppCategories() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("mobileAppCategories", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetMobileAppConfigurations()))
        for i, v := range m.GetMobileAppConfigurations() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("mobileAppConfigurations", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetMobileApps()))
        for i, v := range m.GetMobileApps() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("mobileApps", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetPolicySets()))
        for i, v := range m.GetPolicySets() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("policySets", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetSideLoadingKeys()))
        for i, v := range m.GetSideLoadingKeys() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
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
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetTargetedManagedAppConfigurations()))
        for i, v := range m.GetTargetedManagedAppConfigurations() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("targetedManagedAppConfigurations", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetVppTokens()))
        for i, v := range m.GetVppTokens() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("vppTokens", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetWdacSupplementalPolicies()))
        for i, v := range m.GetWdacSupplementalPolicies() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("wdacSupplementalPolicies", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetWindowsInformationProtectionDeviceRegistrations()))
        for i, v := range m.GetWindowsInformationProtectionDeviceRegistrations() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("windowsInformationProtectionDeviceRegistrations", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetWindowsInformationProtectionPolicies()))
        for i, v := range m.GetWindowsInformationProtectionPolicies() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("windowsInformationProtectionPolicies", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetWindowsInformationProtectionWipeActions()))
        for i, v := range m.GetWindowsInformationProtectionWipeActions() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("windowsInformationProtectionWipeActions", cast)
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
// Sets the androidManagedAppProtections property value. Android managed app policies.
// Parameters:
//  - value : Value to set for the androidManagedAppProtections property.
func (m *DeviceAppManagement) SetAndroidManagedAppProtections(value []AndroidManagedAppProtection)() {
    m.androidManagedAppProtections = value
}
// Sets the defaultManagedAppProtections property value. Default managed app policies.
// Parameters:
//  - value : Value to set for the defaultManagedAppProtections property.
func (m *DeviceAppManagement) SetDefaultManagedAppProtections(value []DefaultManagedAppProtection)() {
    m.defaultManagedAppProtections = value
}
// Sets the deviceAppManagementTasks property value. Device app management tasks.
// Parameters:
//  - value : Value to set for the deviceAppManagementTasks property.
func (m *DeviceAppManagement) SetDeviceAppManagementTasks(value []DeviceAppManagementTask)() {
    m.deviceAppManagementTasks = value
}
// Sets the enterpriseCodeSigningCertificates property value. The Windows Enterprise Code Signing Certificate.
// Parameters:
//  - value : Value to set for the enterpriseCodeSigningCertificates property.
func (m *DeviceAppManagement) SetEnterpriseCodeSigningCertificates(value []EnterpriseCodeSigningCertificate)() {
    m.enterpriseCodeSigningCertificates = value
}
// Sets the iosLobAppProvisioningConfigurations property value. The IOS Lob App Provisioning Configurations.
// Parameters:
//  - value : Value to set for the iosLobAppProvisioningConfigurations property.
func (m *DeviceAppManagement) SetIosLobAppProvisioningConfigurations(value []IosLobAppProvisioningConfiguration)() {
    m.iosLobAppProvisioningConfigurations = value
}
// Sets the iosManagedAppProtections property value. iOS managed app policies.
// Parameters:
//  - value : Value to set for the iosManagedAppProtections property.
func (m *DeviceAppManagement) SetIosManagedAppProtections(value []IosManagedAppProtection)() {
    m.iosManagedAppProtections = value
}
// Sets the isEnabledForMicrosoftStoreForBusiness property value. Whether the account is enabled for syncing applications from the Microsoft Store for Business.
// Parameters:
//  - value : Value to set for the isEnabledForMicrosoftStoreForBusiness property.
func (m *DeviceAppManagement) SetIsEnabledForMicrosoftStoreForBusiness(value *bool)() {
    m.isEnabledForMicrosoftStoreForBusiness = value
}
// Sets the managedAppPolicies property value. Managed app policies.
// Parameters:
//  - value : Value to set for the managedAppPolicies property.
func (m *DeviceAppManagement) SetManagedAppPolicies(value []ManagedAppPolicy)() {
    m.managedAppPolicies = value
}
// Sets the managedAppRegistrations property value. The managed app registrations.
// Parameters:
//  - value : Value to set for the managedAppRegistrations property.
func (m *DeviceAppManagement) SetManagedAppRegistrations(value []ManagedAppRegistration)() {
    m.managedAppRegistrations = value
}
// Sets the managedAppStatuses property value. The managed app statuses.
// Parameters:
//  - value : Value to set for the managedAppStatuses property.
func (m *DeviceAppManagement) SetManagedAppStatuses(value []ManagedAppStatus)() {
    m.managedAppStatuses = value
}
// Sets the managedEBookCategories property value. The mobile eBook categories.
// Parameters:
//  - value : Value to set for the managedEBookCategories property.
func (m *DeviceAppManagement) SetManagedEBookCategories(value []ManagedEBookCategory)() {
    m.managedEBookCategories = value
}
// Sets the managedEBooks property value. The Managed eBook.
// Parameters:
//  - value : Value to set for the managedEBooks property.
func (m *DeviceAppManagement) SetManagedEBooks(value []ManagedEBook)() {
    m.managedEBooks = value
}
// Sets the mdmWindowsInformationProtectionPolicies property value. Windows information protection for apps running on devices which are MDM enrolled.
// Parameters:
//  - value : Value to set for the mdmWindowsInformationProtectionPolicies property.
func (m *DeviceAppManagement) SetMdmWindowsInformationProtectionPolicies(value []MdmWindowsInformationProtectionPolicy)() {
    m.mdmWindowsInformationProtectionPolicies = value
}
// Sets the microsoftStoreForBusinessLanguage property value. The locale information used to sync applications from the Microsoft Store for Business. Cultures that are specific to a country/region. The names of these cultures follow RFC 4646 (Windows Vista and later). The format is -<country/regioncode2>, where  is a lowercase two-letter code derived from ISO 639-1 and <country/regioncode2> is an uppercase two-letter code derived from ISO 3166. For example, en-US for English (United States) is a specific culture.
// Parameters:
//  - value : Value to set for the microsoftStoreForBusinessLanguage property.
func (m *DeviceAppManagement) SetMicrosoftStoreForBusinessLanguage(value *string)() {
    m.microsoftStoreForBusinessLanguage = value
}
// Sets the microsoftStoreForBusinessLastCompletedApplicationSyncTime property value. The last time an application sync from the Microsoft Store for Business was completed.
// Parameters:
//  - value : Value to set for the microsoftStoreForBusinessLastCompletedApplicationSyncTime property.
func (m *DeviceAppManagement) SetMicrosoftStoreForBusinessLastCompletedApplicationSyncTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.microsoftStoreForBusinessLastCompletedApplicationSyncTime = value
}
// Sets the microsoftStoreForBusinessLastSuccessfulSyncDateTime property value. The last time the apps from the Microsoft Store for Business were synced successfully for the account.
// Parameters:
//  - value : Value to set for the microsoftStoreForBusinessLastSuccessfulSyncDateTime property.
func (m *DeviceAppManagement) SetMicrosoftStoreForBusinessLastSuccessfulSyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.microsoftStoreForBusinessLastSuccessfulSyncDateTime = value
}
// Sets the microsoftStoreForBusinessPortalSelection property value. The end user portal information is used to sync applications from the Microsoft Store for Business to Intune Company Portal. There are three options to pick from ['Company portal only', 'Company portal and private store', 'Private store only']. Possible values are: none, companyPortal, privateStore.
// Parameters:
//  - value : Value to set for the microsoftStoreForBusinessPortalSelection property.
func (m *DeviceAppManagement) SetMicrosoftStoreForBusinessPortalSelection(value *MicrosoftStoreForBusinessPortalSelectionOptions)() {
    m.microsoftStoreForBusinessPortalSelection = value
}
// Sets the mobileAppCategories property value. The mobile app categories.
// Parameters:
//  - value : Value to set for the mobileAppCategories property.
func (m *DeviceAppManagement) SetMobileAppCategories(value []MobileAppCategory)() {
    m.mobileAppCategories = value
}
// Sets the mobileAppConfigurations property value. The Managed Device Mobile Application Configurations.
// Parameters:
//  - value : Value to set for the mobileAppConfigurations property.
func (m *DeviceAppManagement) SetMobileAppConfigurations(value []ManagedDeviceMobileAppConfiguration)() {
    m.mobileAppConfigurations = value
}
// Sets the mobileApps property value. The mobile apps.
// Parameters:
//  - value : Value to set for the mobileApps property.
func (m *DeviceAppManagement) SetMobileApps(value []MobileApp)() {
    m.mobileApps = value
}
// Sets the policySets property value. The PolicySet of Policies and Applications
// Parameters:
//  - value : Value to set for the policySets property.
func (m *DeviceAppManagement) SetPolicySets(value []PolicySet)() {
    m.policySets = value
}
// Sets the sideLoadingKeys property value. Side Loading Keys that are required for the Windows 8 and 8.1 Apps installation.
// Parameters:
//  - value : Value to set for the sideLoadingKeys property.
func (m *DeviceAppManagement) SetSideLoadingKeys(value []SideLoadingKey)() {
    m.sideLoadingKeys = value
}
// Sets the symantecCodeSigningCertificate property value. The WinPhone Symantec Code Signing Certificate.
// Parameters:
//  - value : Value to set for the symantecCodeSigningCertificate property.
func (m *DeviceAppManagement) SetSymantecCodeSigningCertificate(value *SymantecCodeSigningCertificate)() {
    m.symantecCodeSigningCertificate = value
}
// Sets the targetedManagedAppConfigurations property value. Targeted managed app configurations.
// Parameters:
//  - value : Value to set for the targetedManagedAppConfigurations property.
func (m *DeviceAppManagement) SetTargetedManagedAppConfigurations(value []TargetedManagedAppConfiguration)() {
    m.targetedManagedAppConfigurations = value
}
// Sets the vppTokens property value. List of Vpp tokens for this organization.
// Parameters:
//  - value : Value to set for the vppTokens property.
func (m *DeviceAppManagement) SetVppTokens(value []VppToken)() {
    m.vppTokens = value
}
// Sets the wdacSupplementalPolicies property value. The collection of Windows Defender Application Control Supplemental Policies.
// Parameters:
//  - value : Value to set for the wdacSupplementalPolicies property.
func (m *DeviceAppManagement) SetWdacSupplementalPolicies(value []WindowsDefenderApplicationControlSupplementalPolicy)() {
    m.wdacSupplementalPolicies = value
}
// Sets the windowsInformationProtectionDeviceRegistrations property value. Windows information protection device registrations that are not MDM enrolled.
// Parameters:
//  - value : Value to set for the windowsInformationProtectionDeviceRegistrations property.
func (m *DeviceAppManagement) SetWindowsInformationProtectionDeviceRegistrations(value []WindowsInformationProtectionDeviceRegistration)() {
    m.windowsInformationProtectionDeviceRegistrations = value
}
// Sets the windowsInformationProtectionPolicies property value. Windows information protection for apps running on devices which are not MDM enrolled.
// Parameters:
//  - value : Value to set for the windowsInformationProtectionPolicies property.
func (m *DeviceAppManagement) SetWindowsInformationProtectionPolicies(value []WindowsInformationProtectionPolicy)() {
    m.windowsInformationProtectionPolicies = value
}
// Sets the windowsInformationProtectionWipeActions property value. Windows information protection wipe actions.
// Parameters:
//  - value : Value to set for the windowsInformationProtectionWipeActions property.
func (m *DeviceAppManagement) SetWindowsInformationProtectionWipeActions(value []WindowsInformationProtectionWipeAction)() {
    m.windowsInformationProtectionWipeActions = value
}
// Sets the windowsManagementApp property value. Windows management app.
// Parameters:
//  - value : Value to set for the windowsManagementApp property.
func (m *DeviceAppManagement) SetWindowsManagementApp(value *WindowsManagementApp)() {
    m.windowsManagementApp = value
}
