package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceAppManagement singleton entity that acts as a container for all device app management functionality.
type DeviceAppManagement struct {
    Entity
}
// NewDeviceAppManagement instantiates a new deviceAppManagement and sets the default values.
func NewDeviceAppManagement()(*DeviceAppManagement) {
    m := &DeviceAppManagement{
        Entity: *NewEntity(),
    }
    return m
}
// CreateDeviceAppManagementFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceAppManagementFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceAppManagement(), nil
}
// GetAndroidManagedAppProtections gets the androidManagedAppProtections property value. Android managed app policies.
func (m *DeviceAppManagement) GetAndroidManagedAppProtections()([]AndroidManagedAppProtectionable) {
    val, err := m.GetBackingStore().Get("androidManagedAppProtections")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]AndroidManagedAppProtectionable)
    }
    return nil
}
// GetDefaultManagedAppProtections gets the defaultManagedAppProtections property value. Default managed app policies.
func (m *DeviceAppManagement) GetDefaultManagedAppProtections()([]DefaultManagedAppProtectionable) {
    val, err := m.GetBackingStore().Get("defaultManagedAppProtections")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]DefaultManagedAppProtectionable)
    }
    return nil
}
// GetDeviceAppManagementTasks gets the deviceAppManagementTasks property value. Device app management tasks.
func (m *DeviceAppManagement) GetDeviceAppManagementTasks()([]DeviceAppManagementTaskable) {
    val, err := m.GetBackingStore().Get("deviceAppManagementTasks")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]DeviceAppManagementTaskable)
    }
    return nil
}
// GetEnterpriseCodeSigningCertificates gets the enterpriseCodeSigningCertificates property value. The Windows Enterprise Code Signing Certificate.
func (m *DeviceAppManagement) GetEnterpriseCodeSigningCertificates()([]EnterpriseCodeSigningCertificateable) {
    val, err := m.GetBackingStore().Get("enterpriseCodeSigningCertificates")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]EnterpriseCodeSigningCertificateable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceAppManagement) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["androidManagedAppProtections"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAndroidManagedAppProtectionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AndroidManagedAppProtectionable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(AndroidManagedAppProtectionable)
                }
            }
            m.SetAndroidManagedAppProtections(res)
        }
        return nil
    }
    res["defaultManagedAppProtections"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDefaultManagedAppProtectionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DefaultManagedAppProtectionable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(DefaultManagedAppProtectionable)
                }
            }
            m.SetDefaultManagedAppProtections(res)
        }
        return nil
    }
    res["deviceAppManagementTasks"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceAppManagementTaskFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceAppManagementTaskable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(DeviceAppManagementTaskable)
                }
            }
            m.SetDeviceAppManagementTasks(res)
        }
        return nil
    }
    res["enterpriseCodeSigningCertificates"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateEnterpriseCodeSigningCertificateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]EnterpriseCodeSigningCertificateable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(EnterpriseCodeSigningCertificateable)
                }
            }
            m.SetEnterpriseCodeSigningCertificates(res)
        }
        return nil
    }
    res["iosLobAppProvisioningConfigurations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateIosLobAppProvisioningConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]IosLobAppProvisioningConfigurationable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(IosLobAppProvisioningConfigurationable)
                }
            }
            m.SetIosLobAppProvisioningConfigurations(res)
        }
        return nil
    }
    res["iosManagedAppProtections"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateIosManagedAppProtectionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]IosManagedAppProtectionable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(IosManagedAppProtectionable)
                }
            }
            m.SetIosManagedAppProtections(res)
        }
        return nil
    }
    res["isEnabledForMicrosoftStoreForBusiness"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsEnabledForMicrosoftStoreForBusiness(val)
        }
        return nil
    }
    res["managedAppPolicies"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateManagedAppPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagedAppPolicyable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ManagedAppPolicyable)
                }
            }
            m.SetManagedAppPolicies(res)
        }
        return nil
    }
    res["managedAppRegistrations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateManagedAppRegistrationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagedAppRegistrationable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ManagedAppRegistrationable)
                }
            }
            m.SetManagedAppRegistrations(res)
        }
        return nil
    }
    res["managedAppStatuses"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateManagedAppStatusFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagedAppStatusable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ManagedAppStatusable)
                }
            }
            m.SetManagedAppStatuses(res)
        }
        return nil
    }
    res["managedEBookCategories"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateManagedEBookCategoryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagedEBookCategoryable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ManagedEBookCategoryable)
                }
            }
            m.SetManagedEBookCategories(res)
        }
        return nil
    }
    res["managedEBooks"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateManagedEBookFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagedEBookable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ManagedEBookable)
                }
            }
            m.SetManagedEBooks(res)
        }
        return nil
    }
    res["mdmWindowsInformationProtectionPolicies"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMdmWindowsInformationProtectionPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MdmWindowsInformationProtectionPolicyable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(MdmWindowsInformationProtectionPolicyable)
                }
            }
            m.SetMdmWindowsInformationProtectionPolicies(res)
        }
        return nil
    }
    res["microsoftStoreForBusinessLanguage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMicrosoftStoreForBusinessLanguage(val)
        }
        return nil
    }
    res["microsoftStoreForBusinessLastCompletedApplicationSyncTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMicrosoftStoreForBusinessLastCompletedApplicationSyncTime(val)
        }
        return nil
    }
    res["microsoftStoreForBusinessLastSuccessfulSyncDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMicrosoftStoreForBusinessLastSuccessfulSyncDateTime(val)
        }
        return nil
    }
    res["microsoftStoreForBusinessPortalSelection"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseMicrosoftStoreForBusinessPortalSelectionOptions)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMicrosoftStoreForBusinessPortalSelection(val.(*MicrosoftStoreForBusinessPortalSelectionOptions))
        }
        return nil
    }
    res["mobileAppCategories"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMobileAppCategoryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MobileAppCategoryable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(MobileAppCategoryable)
                }
            }
            m.SetMobileAppCategories(res)
        }
        return nil
    }
    res["mobileAppConfigurations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateManagedDeviceMobileAppConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagedDeviceMobileAppConfigurationable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ManagedDeviceMobileAppConfigurationable)
                }
            }
            m.SetMobileAppConfigurations(res)
        }
        return nil
    }
    res["mobileApps"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMobileAppFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MobileAppable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(MobileAppable)
                }
            }
            m.SetMobileApps(res)
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
    res["policySets"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePolicySetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PolicySetable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(PolicySetable)
                }
            }
            m.SetPolicySets(res)
        }
        return nil
    }
    res["symantecCodeSigningCertificate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSymantecCodeSigningCertificateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSymantecCodeSigningCertificate(val.(SymantecCodeSigningCertificateable))
        }
        return nil
    }
    res["targetedManagedAppConfigurations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateTargetedManagedAppConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TargetedManagedAppConfigurationable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(TargetedManagedAppConfigurationable)
                }
            }
            m.SetTargetedManagedAppConfigurations(res)
        }
        return nil
    }
    res["vppTokens"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateVppTokenFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]VppTokenable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(VppTokenable)
                }
            }
            m.SetVppTokens(res)
        }
        return nil
    }
    res["wdacSupplementalPolicies"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWindowsDefenderApplicationControlSupplementalPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WindowsDefenderApplicationControlSupplementalPolicyable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(WindowsDefenderApplicationControlSupplementalPolicyable)
                }
            }
            m.SetWdacSupplementalPolicies(res)
        }
        return nil
    }
    res["windowsInformationProtectionDeviceRegistrations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWindowsInformationProtectionDeviceRegistrationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WindowsInformationProtectionDeviceRegistrationable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(WindowsInformationProtectionDeviceRegistrationable)
                }
            }
            m.SetWindowsInformationProtectionDeviceRegistrations(res)
        }
        return nil
    }
    res["windowsInformationProtectionPolicies"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWindowsInformationProtectionPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WindowsInformationProtectionPolicyable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(WindowsInformationProtectionPolicyable)
                }
            }
            m.SetWindowsInformationProtectionPolicies(res)
        }
        return nil
    }
    res["windowsInformationProtectionWipeActions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWindowsInformationProtectionWipeActionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WindowsInformationProtectionWipeActionable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(WindowsInformationProtectionWipeActionable)
                }
            }
            m.SetWindowsInformationProtectionWipeActions(res)
        }
        return nil
    }
    res["windowsManagedAppProtections"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWindowsManagedAppProtectionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WindowsManagedAppProtectionable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(WindowsManagedAppProtectionable)
                }
            }
            m.SetWindowsManagedAppProtections(res)
        }
        return nil
    }
    res["windowsManagementApp"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWindowsManagementAppFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWindowsManagementApp(val.(WindowsManagementAppable))
        }
        return nil
    }
    return res
}
// GetIosLobAppProvisioningConfigurations gets the iosLobAppProvisioningConfigurations property value. The IOS Lob App Provisioning Configurations.
func (m *DeviceAppManagement) GetIosLobAppProvisioningConfigurations()([]IosLobAppProvisioningConfigurationable) {
    val, err := m.GetBackingStore().Get("iosLobAppProvisioningConfigurations")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]IosLobAppProvisioningConfigurationable)
    }
    return nil
}
// GetIosManagedAppProtections gets the iosManagedAppProtections property value. iOS managed app policies.
func (m *DeviceAppManagement) GetIosManagedAppProtections()([]IosManagedAppProtectionable) {
    val, err := m.GetBackingStore().Get("iosManagedAppProtections")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]IosManagedAppProtectionable)
    }
    return nil
}
// GetIsEnabledForMicrosoftStoreForBusiness gets the isEnabledForMicrosoftStoreForBusiness property value. Whether the account is enabled for syncing applications from the Microsoft Store for Business.
func (m *DeviceAppManagement) GetIsEnabledForMicrosoftStoreForBusiness()(*bool) {
    val, err := m.GetBackingStore().Get("isEnabledForMicrosoftStoreForBusiness")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetManagedAppPolicies gets the managedAppPolicies property value. Managed app policies.
func (m *DeviceAppManagement) GetManagedAppPolicies()([]ManagedAppPolicyable) {
    val, err := m.GetBackingStore().Get("managedAppPolicies")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ManagedAppPolicyable)
    }
    return nil
}
// GetManagedAppRegistrations gets the managedAppRegistrations property value. The managed app registrations.
func (m *DeviceAppManagement) GetManagedAppRegistrations()([]ManagedAppRegistrationable) {
    val, err := m.GetBackingStore().Get("managedAppRegistrations")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ManagedAppRegistrationable)
    }
    return nil
}
// GetManagedAppStatuses gets the managedAppStatuses property value. The managed app statuses.
func (m *DeviceAppManagement) GetManagedAppStatuses()([]ManagedAppStatusable) {
    val, err := m.GetBackingStore().Get("managedAppStatuses")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ManagedAppStatusable)
    }
    return nil
}
// GetManagedEBookCategories gets the managedEBookCategories property value. The mobile eBook categories.
func (m *DeviceAppManagement) GetManagedEBookCategories()([]ManagedEBookCategoryable) {
    val, err := m.GetBackingStore().Get("managedEBookCategories")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ManagedEBookCategoryable)
    }
    return nil
}
// GetManagedEBooks gets the managedEBooks property value. The Managed eBook.
func (m *DeviceAppManagement) GetManagedEBooks()([]ManagedEBookable) {
    val, err := m.GetBackingStore().Get("managedEBooks")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ManagedEBookable)
    }
    return nil
}
// GetMdmWindowsInformationProtectionPolicies gets the mdmWindowsInformationProtectionPolicies property value. Windows information protection for apps running on devices which are MDM enrolled.
func (m *DeviceAppManagement) GetMdmWindowsInformationProtectionPolicies()([]MdmWindowsInformationProtectionPolicyable) {
    val, err := m.GetBackingStore().Get("mdmWindowsInformationProtectionPolicies")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]MdmWindowsInformationProtectionPolicyable)
    }
    return nil
}
// GetMicrosoftStoreForBusinessLanguage gets the microsoftStoreForBusinessLanguage property value. The locale information used to sync applications from the Microsoft Store for Business. Cultures that are specific to a country/region. The names of these cultures follow RFC 4646 (Windows Vista and later). The format is -<country/regioncode2>, where  is a lowercase two-letter code derived from ISO 639-1 and <country/regioncode2> is an uppercase two-letter code derived from ISO 3166. For example, en-US for English (United States) is a specific culture.
func (m *DeviceAppManagement) GetMicrosoftStoreForBusinessLanguage()(*string) {
    val, err := m.GetBackingStore().Get("microsoftStoreForBusinessLanguage")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetMicrosoftStoreForBusinessLastCompletedApplicationSyncTime gets the microsoftStoreForBusinessLastCompletedApplicationSyncTime property value. The last time an application sync from the Microsoft Store for Business was completed.
func (m *DeviceAppManagement) GetMicrosoftStoreForBusinessLastCompletedApplicationSyncTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("microsoftStoreForBusinessLastCompletedApplicationSyncTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetMicrosoftStoreForBusinessLastSuccessfulSyncDateTime gets the microsoftStoreForBusinessLastSuccessfulSyncDateTime property value. The last time the apps from the Microsoft Store for Business were synced successfully for the account.
func (m *DeviceAppManagement) GetMicrosoftStoreForBusinessLastSuccessfulSyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("microsoftStoreForBusinessLastSuccessfulSyncDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetMicrosoftStoreForBusinessPortalSelection gets the microsoftStoreForBusinessPortalSelection property value. Portal to which admin syncs available Microsoft Store for Business apps. This is available in the Intune Admin Console.
func (m *DeviceAppManagement) GetMicrosoftStoreForBusinessPortalSelection()(*MicrosoftStoreForBusinessPortalSelectionOptions) {
    val, err := m.GetBackingStore().Get("microsoftStoreForBusinessPortalSelection")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*MicrosoftStoreForBusinessPortalSelectionOptions)
    }
    return nil
}
// GetMobileAppCategories gets the mobileAppCategories property value. The mobile app categories.
func (m *DeviceAppManagement) GetMobileAppCategories()([]MobileAppCategoryable) {
    val, err := m.GetBackingStore().Get("mobileAppCategories")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]MobileAppCategoryable)
    }
    return nil
}
// GetMobileAppConfigurations gets the mobileAppConfigurations property value. The Managed Device Mobile Application Configurations.
func (m *DeviceAppManagement) GetMobileAppConfigurations()([]ManagedDeviceMobileAppConfigurationable) {
    val, err := m.GetBackingStore().Get("mobileAppConfigurations")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ManagedDeviceMobileAppConfigurationable)
    }
    return nil
}
// GetMobileApps gets the mobileApps property value. The mobile apps.
func (m *DeviceAppManagement) GetMobileApps()([]MobileAppable) {
    val, err := m.GetBackingStore().Get("mobileApps")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]MobileAppable)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *DeviceAppManagement) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPolicySets gets the policySets property value. The PolicySet of Policies and Applications
func (m *DeviceAppManagement) GetPolicySets()([]PolicySetable) {
    val, err := m.GetBackingStore().Get("policySets")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]PolicySetable)
    }
    return nil
}
// GetSymantecCodeSigningCertificate gets the symantecCodeSigningCertificate property value. The WinPhone Symantec Code Signing Certificate.
func (m *DeviceAppManagement) GetSymantecCodeSigningCertificate()(SymantecCodeSigningCertificateable) {
    val, err := m.GetBackingStore().Get("symantecCodeSigningCertificate")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(SymantecCodeSigningCertificateable)
    }
    return nil
}
// GetTargetedManagedAppConfigurations gets the targetedManagedAppConfigurations property value. Targeted managed app configurations.
func (m *DeviceAppManagement) GetTargetedManagedAppConfigurations()([]TargetedManagedAppConfigurationable) {
    val, err := m.GetBackingStore().Get("targetedManagedAppConfigurations")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]TargetedManagedAppConfigurationable)
    }
    return nil
}
// GetVppTokens gets the vppTokens property value. List of Vpp tokens for this organization.
func (m *DeviceAppManagement) GetVppTokens()([]VppTokenable) {
    val, err := m.GetBackingStore().Get("vppTokens")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]VppTokenable)
    }
    return nil
}
// GetWdacSupplementalPolicies gets the wdacSupplementalPolicies property value. The collection of Windows Defender Application Control Supplemental Policies.
func (m *DeviceAppManagement) GetWdacSupplementalPolicies()([]WindowsDefenderApplicationControlSupplementalPolicyable) {
    val, err := m.GetBackingStore().Get("wdacSupplementalPolicies")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]WindowsDefenderApplicationControlSupplementalPolicyable)
    }
    return nil
}
// GetWindowsInformationProtectionDeviceRegistrations gets the windowsInformationProtectionDeviceRegistrations property value. Windows information protection device registrations that are not MDM enrolled.
func (m *DeviceAppManagement) GetWindowsInformationProtectionDeviceRegistrations()([]WindowsInformationProtectionDeviceRegistrationable) {
    val, err := m.GetBackingStore().Get("windowsInformationProtectionDeviceRegistrations")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]WindowsInformationProtectionDeviceRegistrationable)
    }
    return nil
}
// GetWindowsInformationProtectionPolicies gets the windowsInformationProtectionPolicies property value. Windows information protection for apps running on devices which are not MDM enrolled.
func (m *DeviceAppManagement) GetWindowsInformationProtectionPolicies()([]WindowsInformationProtectionPolicyable) {
    val, err := m.GetBackingStore().Get("windowsInformationProtectionPolicies")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]WindowsInformationProtectionPolicyable)
    }
    return nil
}
// GetWindowsInformationProtectionWipeActions gets the windowsInformationProtectionWipeActions property value. Windows information protection wipe actions.
func (m *DeviceAppManagement) GetWindowsInformationProtectionWipeActions()([]WindowsInformationProtectionWipeActionable) {
    val, err := m.GetBackingStore().Get("windowsInformationProtectionWipeActions")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]WindowsInformationProtectionWipeActionable)
    }
    return nil
}
// GetWindowsManagedAppProtections gets the windowsManagedAppProtections property value. Windows managed app policies.
func (m *DeviceAppManagement) GetWindowsManagedAppProtections()([]WindowsManagedAppProtectionable) {
    val, err := m.GetBackingStore().Get("windowsManagedAppProtections")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]WindowsManagedAppProtectionable)
    }
    return nil
}
// GetWindowsManagementApp gets the windowsManagementApp property value. Windows management app.
func (m *DeviceAppManagement) GetWindowsManagementApp()(WindowsManagementAppable) {
    val, err := m.GetBackingStore().Get("windowsManagementApp")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(WindowsManagementAppable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *DeviceAppManagement) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAndroidManagedAppProtections() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAndroidManagedAppProtections()))
        for i, v := range m.GetAndroidManagedAppProtections() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("androidManagedAppProtections", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDefaultManagedAppProtections() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDefaultManagedAppProtections()))
        for i, v := range m.GetDefaultManagedAppProtections() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("defaultManagedAppProtections", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDeviceAppManagementTasks() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDeviceAppManagementTasks()))
        for i, v := range m.GetDeviceAppManagementTasks() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("deviceAppManagementTasks", cast)
        if err != nil {
            return err
        }
    }
    if m.GetEnterpriseCodeSigningCertificates() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetEnterpriseCodeSigningCertificates()))
        for i, v := range m.GetEnterpriseCodeSigningCertificates() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("enterpriseCodeSigningCertificates", cast)
        if err != nil {
            return err
        }
    }
    if m.GetIosLobAppProvisioningConfigurations() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetIosLobAppProvisioningConfigurations()))
        for i, v := range m.GetIosLobAppProvisioningConfigurations() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("iosLobAppProvisioningConfigurations", cast)
        if err != nil {
            return err
        }
    }
    if m.GetIosManagedAppProtections() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetIosManagedAppProtections()))
        for i, v := range m.GetIosManagedAppProtections() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
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
    if m.GetManagedAppPolicies() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetManagedAppPolicies()))
        for i, v := range m.GetManagedAppPolicies() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("managedAppPolicies", cast)
        if err != nil {
            return err
        }
    }
    if m.GetManagedAppRegistrations() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetManagedAppRegistrations()))
        for i, v := range m.GetManagedAppRegistrations() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("managedAppRegistrations", cast)
        if err != nil {
            return err
        }
    }
    if m.GetManagedAppStatuses() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetManagedAppStatuses()))
        for i, v := range m.GetManagedAppStatuses() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("managedAppStatuses", cast)
        if err != nil {
            return err
        }
    }
    if m.GetManagedEBookCategories() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetManagedEBookCategories()))
        for i, v := range m.GetManagedEBookCategories() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("managedEBookCategories", cast)
        if err != nil {
            return err
        }
    }
    if m.GetManagedEBooks() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetManagedEBooks()))
        for i, v := range m.GetManagedEBooks() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("managedEBooks", cast)
        if err != nil {
            return err
        }
    }
    if m.GetMdmWindowsInformationProtectionPolicies() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetMdmWindowsInformationProtectionPolicies()))
        for i, v := range m.GetMdmWindowsInformationProtectionPolicies() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
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
        cast := (*m.GetMicrosoftStoreForBusinessPortalSelection()).String()
        err = writer.WriteStringValue("microsoftStoreForBusinessPortalSelection", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetMobileAppCategories() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetMobileAppCategories()))
        for i, v := range m.GetMobileAppCategories() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("mobileAppCategories", cast)
        if err != nil {
            return err
        }
    }
    if m.GetMobileAppConfigurations() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetMobileAppConfigurations()))
        for i, v := range m.GetMobileAppConfigurations() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("mobileAppConfigurations", cast)
        if err != nil {
            return err
        }
    }
    if m.GetMobileApps() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetMobileApps()))
        for i, v := range m.GetMobileApps() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("mobileApps", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    if m.GetPolicySets() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetPolicySets()))
        for i, v := range m.GetPolicySets() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("policySets", cast)
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
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTargetedManagedAppConfigurations()))
        for i, v := range m.GetTargetedManagedAppConfigurations() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("targetedManagedAppConfigurations", cast)
        if err != nil {
            return err
        }
    }
    if m.GetVppTokens() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetVppTokens()))
        for i, v := range m.GetVppTokens() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("vppTokens", cast)
        if err != nil {
            return err
        }
    }
    if m.GetWdacSupplementalPolicies() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetWdacSupplementalPolicies()))
        for i, v := range m.GetWdacSupplementalPolicies() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("wdacSupplementalPolicies", cast)
        if err != nil {
            return err
        }
    }
    if m.GetWindowsInformationProtectionDeviceRegistrations() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetWindowsInformationProtectionDeviceRegistrations()))
        for i, v := range m.GetWindowsInformationProtectionDeviceRegistrations() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("windowsInformationProtectionDeviceRegistrations", cast)
        if err != nil {
            return err
        }
    }
    if m.GetWindowsInformationProtectionPolicies() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetWindowsInformationProtectionPolicies()))
        for i, v := range m.GetWindowsInformationProtectionPolicies() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("windowsInformationProtectionPolicies", cast)
        if err != nil {
            return err
        }
    }
    if m.GetWindowsInformationProtectionWipeActions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetWindowsInformationProtectionWipeActions()))
        for i, v := range m.GetWindowsInformationProtectionWipeActions() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("windowsInformationProtectionWipeActions", cast)
        if err != nil {
            return err
        }
    }
    if m.GetWindowsManagedAppProtections() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetWindowsManagedAppProtections()))
        for i, v := range m.GetWindowsManagedAppProtections() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
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
    err := m.GetBackingStore().Set("androidManagedAppProtections", value)
    if err != nil {
        panic(err)
    }
}
// SetDefaultManagedAppProtections sets the defaultManagedAppProtections property value. Default managed app policies.
func (m *DeviceAppManagement) SetDefaultManagedAppProtections(value []DefaultManagedAppProtectionable)() {
    err := m.GetBackingStore().Set("defaultManagedAppProtections", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceAppManagementTasks sets the deviceAppManagementTasks property value. Device app management tasks.
func (m *DeviceAppManagement) SetDeviceAppManagementTasks(value []DeviceAppManagementTaskable)() {
    err := m.GetBackingStore().Set("deviceAppManagementTasks", value)
    if err != nil {
        panic(err)
    }
}
// SetEnterpriseCodeSigningCertificates sets the enterpriseCodeSigningCertificates property value. The Windows Enterprise Code Signing Certificate.
func (m *DeviceAppManagement) SetEnterpriseCodeSigningCertificates(value []EnterpriseCodeSigningCertificateable)() {
    err := m.GetBackingStore().Set("enterpriseCodeSigningCertificates", value)
    if err != nil {
        panic(err)
    }
}
// SetIosLobAppProvisioningConfigurations sets the iosLobAppProvisioningConfigurations property value. The IOS Lob App Provisioning Configurations.
func (m *DeviceAppManagement) SetIosLobAppProvisioningConfigurations(value []IosLobAppProvisioningConfigurationable)() {
    err := m.GetBackingStore().Set("iosLobAppProvisioningConfigurations", value)
    if err != nil {
        panic(err)
    }
}
// SetIosManagedAppProtections sets the iosManagedAppProtections property value. iOS managed app policies.
func (m *DeviceAppManagement) SetIosManagedAppProtections(value []IosManagedAppProtectionable)() {
    err := m.GetBackingStore().Set("iosManagedAppProtections", value)
    if err != nil {
        panic(err)
    }
}
// SetIsEnabledForMicrosoftStoreForBusiness sets the isEnabledForMicrosoftStoreForBusiness property value. Whether the account is enabled for syncing applications from the Microsoft Store for Business.
func (m *DeviceAppManagement) SetIsEnabledForMicrosoftStoreForBusiness(value *bool)() {
    err := m.GetBackingStore().Set("isEnabledForMicrosoftStoreForBusiness", value)
    if err != nil {
        panic(err)
    }
}
// SetManagedAppPolicies sets the managedAppPolicies property value. Managed app policies.
func (m *DeviceAppManagement) SetManagedAppPolicies(value []ManagedAppPolicyable)() {
    err := m.GetBackingStore().Set("managedAppPolicies", value)
    if err != nil {
        panic(err)
    }
}
// SetManagedAppRegistrations sets the managedAppRegistrations property value. The managed app registrations.
func (m *DeviceAppManagement) SetManagedAppRegistrations(value []ManagedAppRegistrationable)() {
    err := m.GetBackingStore().Set("managedAppRegistrations", value)
    if err != nil {
        panic(err)
    }
}
// SetManagedAppStatuses sets the managedAppStatuses property value. The managed app statuses.
func (m *DeviceAppManagement) SetManagedAppStatuses(value []ManagedAppStatusable)() {
    err := m.GetBackingStore().Set("managedAppStatuses", value)
    if err != nil {
        panic(err)
    }
}
// SetManagedEBookCategories sets the managedEBookCategories property value. The mobile eBook categories.
func (m *DeviceAppManagement) SetManagedEBookCategories(value []ManagedEBookCategoryable)() {
    err := m.GetBackingStore().Set("managedEBookCategories", value)
    if err != nil {
        panic(err)
    }
}
// SetManagedEBooks sets the managedEBooks property value. The Managed eBook.
func (m *DeviceAppManagement) SetManagedEBooks(value []ManagedEBookable)() {
    err := m.GetBackingStore().Set("managedEBooks", value)
    if err != nil {
        panic(err)
    }
}
// SetMdmWindowsInformationProtectionPolicies sets the mdmWindowsInformationProtectionPolicies property value. Windows information protection for apps running on devices which are MDM enrolled.
func (m *DeviceAppManagement) SetMdmWindowsInformationProtectionPolicies(value []MdmWindowsInformationProtectionPolicyable)() {
    err := m.GetBackingStore().Set("mdmWindowsInformationProtectionPolicies", value)
    if err != nil {
        panic(err)
    }
}
// SetMicrosoftStoreForBusinessLanguage sets the microsoftStoreForBusinessLanguage property value. The locale information used to sync applications from the Microsoft Store for Business. Cultures that are specific to a country/region. The names of these cultures follow RFC 4646 (Windows Vista and later). The format is -<country/regioncode2>, where  is a lowercase two-letter code derived from ISO 639-1 and <country/regioncode2> is an uppercase two-letter code derived from ISO 3166. For example, en-US for English (United States) is a specific culture.
func (m *DeviceAppManagement) SetMicrosoftStoreForBusinessLanguage(value *string)() {
    err := m.GetBackingStore().Set("microsoftStoreForBusinessLanguage", value)
    if err != nil {
        panic(err)
    }
}
// SetMicrosoftStoreForBusinessLastCompletedApplicationSyncTime sets the microsoftStoreForBusinessLastCompletedApplicationSyncTime property value. The last time an application sync from the Microsoft Store for Business was completed.
func (m *DeviceAppManagement) SetMicrosoftStoreForBusinessLastCompletedApplicationSyncTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("microsoftStoreForBusinessLastCompletedApplicationSyncTime", value)
    if err != nil {
        panic(err)
    }
}
// SetMicrosoftStoreForBusinessLastSuccessfulSyncDateTime sets the microsoftStoreForBusinessLastSuccessfulSyncDateTime property value. The last time the apps from the Microsoft Store for Business were synced successfully for the account.
func (m *DeviceAppManagement) SetMicrosoftStoreForBusinessLastSuccessfulSyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("microsoftStoreForBusinessLastSuccessfulSyncDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetMicrosoftStoreForBusinessPortalSelection sets the microsoftStoreForBusinessPortalSelection property value. Portal to which admin syncs available Microsoft Store for Business apps. This is available in the Intune Admin Console.
func (m *DeviceAppManagement) SetMicrosoftStoreForBusinessPortalSelection(value *MicrosoftStoreForBusinessPortalSelectionOptions)() {
    err := m.GetBackingStore().Set("microsoftStoreForBusinessPortalSelection", value)
    if err != nil {
        panic(err)
    }
}
// SetMobileAppCategories sets the mobileAppCategories property value. The mobile app categories.
func (m *DeviceAppManagement) SetMobileAppCategories(value []MobileAppCategoryable)() {
    err := m.GetBackingStore().Set("mobileAppCategories", value)
    if err != nil {
        panic(err)
    }
}
// SetMobileAppConfigurations sets the mobileAppConfigurations property value. The Managed Device Mobile Application Configurations.
func (m *DeviceAppManagement) SetMobileAppConfigurations(value []ManagedDeviceMobileAppConfigurationable)() {
    err := m.GetBackingStore().Set("mobileAppConfigurations", value)
    if err != nil {
        panic(err)
    }
}
// SetMobileApps sets the mobileApps property value. The mobile apps.
func (m *DeviceAppManagement) SetMobileApps(value []MobileAppable)() {
    err := m.GetBackingStore().Set("mobileApps", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *DeviceAppManagement) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetPolicySets sets the policySets property value. The PolicySet of Policies and Applications
func (m *DeviceAppManagement) SetPolicySets(value []PolicySetable)() {
    err := m.GetBackingStore().Set("policySets", value)
    if err != nil {
        panic(err)
    }
}
// SetSymantecCodeSigningCertificate sets the symantecCodeSigningCertificate property value. The WinPhone Symantec Code Signing Certificate.
func (m *DeviceAppManagement) SetSymantecCodeSigningCertificate(value SymantecCodeSigningCertificateable)() {
    err := m.GetBackingStore().Set("symantecCodeSigningCertificate", value)
    if err != nil {
        panic(err)
    }
}
// SetTargetedManagedAppConfigurations sets the targetedManagedAppConfigurations property value. Targeted managed app configurations.
func (m *DeviceAppManagement) SetTargetedManagedAppConfigurations(value []TargetedManagedAppConfigurationable)() {
    err := m.GetBackingStore().Set("targetedManagedAppConfigurations", value)
    if err != nil {
        panic(err)
    }
}
// SetVppTokens sets the vppTokens property value. List of Vpp tokens for this organization.
func (m *DeviceAppManagement) SetVppTokens(value []VppTokenable)() {
    err := m.GetBackingStore().Set("vppTokens", value)
    if err != nil {
        panic(err)
    }
}
// SetWdacSupplementalPolicies sets the wdacSupplementalPolicies property value. The collection of Windows Defender Application Control Supplemental Policies.
func (m *DeviceAppManagement) SetWdacSupplementalPolicies(value []WindowsDefenderApplicationControlSupplementalPolicyable)() {
    err := m.GetBackingStore().Set("wdacSupplementalPolicies", value)
    if err != nil {
        panic(err)
    }
}
// SetWindowsInformationProtectionDeviceRegistrations sets the windowsInformationProtectionDeviceRegistrations property value. Windows information protection device registrations that are not MDM enrolled.
func (m *DeviceAppManagement) SetWindowsInformationProtectionDeviceRegistrations(value []WindowsInformationProtectionDeviceRegistrationable)() {
    err := m.GetBackingStore().Set("windowsInformationProtectionDeviceRegistrations", value)
    if err != nil {
        panic(err)
    }
}
// SetWindowsInformationProtectionPolicies sets the windowsInformationProtectionPolicies property value. Windows information protection for apps running on devices which are not MDM enrolled.
func (m *DeviceAppManagement) SetWindowsInformationProtectionPolicies(value []WindowsInformationProtectionPolicyable)() {
    err := m.GetBackingStore().Set("windowsInformationProtectionPolicies", value)
    if err != nil {
        panic(err)
    }
}
// SetWindowsInformationProtectionWipeActions sets the windowsInformationProtectionWipeActions property value. Windows information protection wipe actions.
func (m *DeviceAppManagement) SetWindowsInformationProtectionWipeActions(value []WindowsInformationProtectionWipeActionable)() {
    err := m.GetBackingStore().Set("windowsInformationProtectionWipeActions", value)
    if err != nil {
        panic(err)
    }
}
// SetWindowsManagedAppProtections sets the windowsManagedAppProtections property value. Windows managed app policies.
func (m *DeviceAppManagement) SetWindowsManagedAppProtections(value []WindowsManagedAppProtectionable)() {
    err := m.GetBackingStore().Set("windowsManagedAppProtections", value)
    if err != nil {
        panic(err)
    }
}
// SetWindowsManagementApp sets the windowsManagementApp property value. Windows management app.
func (m *DeviceAppManagement) SetWindowsManagementApp(value WindowsManagementAppable)() {
    err := m.GetBackingStore().Set("windowsManagementApp", value)
    if err != nil {
        panic(err)
    }
}
// DeviceAppManagementable 
type DeviceAppManagementable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAndroidManagedAppProtections()([]AndroidManagedAppProtectionable)
    GetDefaultManagedAppProtections()([]DefaultManagedAppProtectionable)
    GetDeviceAppManagementTasks()([]DeviceAppManagementTaskable)
    GetEnterpriseCodeSigningCertificates()([]EnterpriseCodeSigningCertificateable)
    GetIosLobAppProvisioningConfigurations()([]IosLobAppProvisioningConfigurationable)
    GetIosManagedAppProtections()([]IosManagedAppProtectionable)
    GetIsEnabledForMicrosoftStoreForBusiness()(*bool)
    GetManagedAppPolicies()([]ManagedAppPolicyable)
    GetManagedAppRegistrations()([]ManagedAppRegistrationable)
    GetManagedAppStatuses()([]ManagedAppStatusable)
    GetManagedEBookCategories()([]ManagedEBookCategoryable)
    GetManagedEBooks()([]ManagedEBookable)
    GetMdmWindowsInformationProtectionPolicies()([]MdmWindowsInformationProtectionPolicyable)
    GetMicrosoftStoreForBusinessLanguage()(*string)
    GetMicrosoftStoreForBusinessLastCompletedApplicationSyncTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetMicrosoftStoreForBusinessLastSuccessfulSyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetMicrosoftStoreForBusinessPortalSelection()(*MicrosoftStoreForBusinessPortalSelectionOptions)
    GetMobileAppCategories()([]MobileAppCategoryable)
    GetMobileAppConfigurations()([]ManagedDeviceMobileAppConfigurationable)
    GetMobileApps()([]MobileAppable)
    GetOdataType()(*string)
    GetPolicySets()([]PolicySetable)
    GetSymantecCodeSigningCertificate()(SymantecCodeSigningCertificateable)
    GetTargetedManagedAppConfigurations()([]TargetedManagedAppConfigurationable)
    GetVppTokens()([]VppTokenable)
    GetWdacSupplementalPolicies()([]WindowsDefenderApplicationControlSupplementalPolicyable)
    GetWindowsInformationProtectionDeviceRegistrations()([]WindowsInformationProtectionDeviceRegistrationable)
    GetWindowsInformationProtectionPolicies()([]WindowsInformationProtectionPolicyable)
    GetWindowsInformationProtectionWipeActions()([]WindowsInformationProtectionWipeActionable)
    GetWindowsManagedAppProtections()([]WindowsManagedAppProtectionable)
    GetWindowsManagementApp()(WindowsManagementAppable)
    SetAndroidManagedAppProtections(value []AndroidManagedAppProtectionable)()
    SetDefaultManagedAppProtections(value []DefaultManagedAppProtectionable)()
    SetDeviceAppManagementTasks(value []DeviceAppManagementTaskable)()
    SetEnterpriseCodeSigningCertificates(value []EnterpriseCodeSigningCertificateable)()
    SetIosLobAppProvisioningConfigurations(value []IosLobAppProvisioningConfigurationable)()
    SetIosManagedAppProtections(value []IosManagedAppProtectionable)()
    SetIsEnabledForMicrosoftStoreForBusiness(value *bool)()
    SetManagedAppPolicies(value []ManagedAppPolicyable)()
    SetManagedAppRegistrations(value []ManagedAppRegistrationable)()
    SetManagedAppStatuses(value []ManagedAppStatusable)()
    SetManagedEBookCategories(value []ManagedEBookCategoryable)()
    SetManagedEBooks(value []ManagedEBookable)()
    SetMdmWindowsInformationProtectionPolicies(value []MdmWindowsInformationProtectionPolicyable)()
    SetMicrosoftStoreForBusinessLanguage(value *string)()
    SetMicrosoftStoreForBusinessLastCompletedApplicationSyncTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetMicrosoftStoreForBusinessLastSuccessfulSyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetMicrosoftStoreForBusinessPortalSelection(value *MicrosoftStoreForBusinessPortalSelectionOptions)()
    SetMobileAppCategories(value []MobileAppCategoryable)()
    SetMobileAppConfigurations(value []ManagedDeviceMobileAppConfigurationable)()
    SetMobileApps(value []MobileAppable)()
    SetOdataType(value *string)()
    SetPolicySets(value []PolicySetable)()
    SetSymantecCodeSigningCertificate(value SymantecCodeSigningCertificateable)()
    SetTargetedManagedAppConfigurations(value []TargetedManagedAppConfigurationable)()
    SetVppTokens(value []VppTokenable)()
    SetWdacSupplementalPolicies(value []WindowsDefenderApplicationControlSupplementalPolicyable)()
    SetWindowsInformationProtectionDeviceRegistrations(value []WindowsInformationProtectionDeviceRegistrationable)()
    SetWindowsInformationProtectionPolicies(value []WindowsInformationProtectionPolicyable)()
    SetWindowsInformationProtectionWipeActions(value []WindowsInformationProtectionWipeActionable)()
    SetWindowsManagedAppProtections(value []WindowsManagedAppProtectionable)()
    SetWindowsManagementApp(value WindowsManagementAppable)()
}
