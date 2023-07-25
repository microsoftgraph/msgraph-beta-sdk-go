package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MacOSKerberosSingleSignOnExtension represents a Kerberos-type Single Sign-On extension profile for macOS devices.
type MacOSKerberosSingleSignOnExtension struct {
    MacOSSingleSignOnExtension
}
// NewMacOSKerberosSingleSignOnExtension instantiates a new macOSKerberosSingleSignOnExtension and sets the default values.
func NewMacOSKerberosSingleSignOnExtension()(*MacOSKerberosSingleSignOnExtension) {
    m := &MacOSKerberosSingleSignOnExtension{
        MacOSSingleSignOnExtension: *NewMacOSSingleSignOnExtension(),
    }
    odataTypeValue := "#microsoft.graph.macOSKerberosSingleSignOnExtension"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateMacOSKerberosSingleSignOnExtensionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMacOSKerberosSingleSignOnExtensionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMacOSKerberosSingleSignOnExtension(), nil
}
// GetActiveDirectorySiteCode gets the activeDirectorySiteCode property value. Gets or sets the Active Directory site.
func (m *MacOSKerberosSingleSignOnExtension) GetActiveDirectorySiteCode()(*string) {
    val, err := m.GetBackingStore().Get("activeDirectorySiteCode")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetBlockActiveDirectorySiteAutoDiscovery gets the blockActiveDirectorySiteAutoDiscovery property value. Enables or disables whether the Kerberos extension can automatically determine its site name.
func (m *MacOSKerberosSingleSignOnExtension) GetBlockActiveDirectorySiteAutoDiscovery()(*bool) {
    val, err := m.GetBackingStore().Get("blockActiveDirectorySiteAutoDiscovery")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetBlockAutomaticLogin gets the blockAutomaticLogin property value. Enables or disables Keychain usage.
func (m *MacOSKerberosSingleSignOnExtension) GetBlockAutomaticLogin()(*bool) {
    val, err := m.GetBackingStore().Get("blockAutomaticLogin")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetCacheName gets the cacheName property value. Gets or sets the Generic Security Services name of the Kerberos cache to use for this profile.
func (m *MacOSKerberosSingleSignOnExtension) GetCacheName()(*string) {
    val, err := m.GetBackingStore().Get("cacheName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetCredentialBundleIdAccessControlList gets the credentialBundleIdAccessControlList property value. Gets or sets a list of app Bundle IDs allowed to access the Kerberos Ticket Granting Ticket.
func (m *MacOSKerberosSingleSignOnExtension) GetCredentialBundleIdAccessControlList()([]string) {
    val, err := m.GetBackingStore().Get("credentialBundleIdAccessControlList")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetCredentialsCacheMonitored gets the credentialsCacheMonitored property value. When set to True, the credential is requested on the next matching Kerberos challenge or network state change. When the credential is expired or missing, a new credential is created. Available for devices running macOS versions 12 and later.
func (m *MacOSKerberosSingleSignOnExtension) GetCredentialsCacheMonitored()(*bool) {
    val, err := m.GetBackingStore().Get("credentialsCacheMonitored")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetDomainRealms gets the domainRealms property value. Gets or sets a list of realms for custom domain-realm mapping. Realms are case sensitive.
func (m *MacOSKerberosSingleSignOnExtension) GetDomainRealms()([]string) {
    val, err := m.GetBackingStore().Get("domainRealms")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetDomains gets the domains property value. Gets or sets a list of hosts or domain names for which the app extension performs SSO.
func (m *MacOSKerberosSingleSignOnExtension) GetDomains()([]string) {
    val, err := m.GetBackingStore().Get("domains")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MacOSKerberosSingleSignOnExtension) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.MacOSSingleSignOnExtension.GetFieldDeserializers()
    res["activeDirectorySiteCode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActiveDirectorySiteCode(val)
        }
        return nil
    }
    res["blockActiveDirectorySiteAutoDiscovery"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBlockActiveDirectorySiteAutoDiscovery(val)
        }
        return nil
    }
    res["blockAutomaticLogin"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBlockAutomaticLogin(val)
        }
        return nil
    }
    res["cacheName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCacheName(val)
        }
        return nil
    }
    res["credentialBundleIdAccessControlList"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetCredentialBundleIdAccessControlList(res)
        }
        return nil
    }
    res["credentialsCacheMonitored"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCredentialsCacheMonitored(val)
        }
        return nil
    }
    res["domainRealms"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetDomainRealms(res)
        }
        return nil
    }
    res["domains"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetDomains(res)
        }
        return nil
    }
    res["isDefaultRealm"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsDefaultRealm(val)
        }
        return nil
    }
    res["kerberosAppsInBundleIdACLIncluded"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKerberosAppsInBundleIdACLIncluded(val)
        }
        return nil
    }
    res["managedAppsInBundleIdACLIncluded"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagedAppsInBundleIdACLIncluded(val)
        }
        return nil
    }
    res["modeCredentialUsed"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetModeCredentialUsed(val)
        }
        return nil
    }
    res["passwordBlockModification"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordBlockModification(val)
        }
        return nil
    }
    res["passwordChangeUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordChangeUrl(val)
        }
        return nil
    }
    res["passwordEnableLocalSync"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordEnableLocalSync(val)
        }
        return nil
    }
    res["passwordExpirationDays"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordExpirationDays(val)
        }
        return nil
    }
    res["passwordExpirationNotificationDays"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordExpirationNotificationDays(val)
        }
        return nil
    }
    res["passwordMinimumAgeDays"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordMinimumAgeDays(val)
        }
        return nil
    }
    res["passwordMinimumLength"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordMinimumLength(val)
        }
        return nil
    }
    res["passwordPreviousPasswordBlockCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordPreviousPasswordBlockCount(val)
        }
        return nil
    }
    res["passwordRequireActiveDirectoryComplexity"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordRequireActiveDirectoryComplexity(val)
        }
        return nil
    }
    res["passwordRequirementsDescription"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordRequirementsDescription(val)
        }
        return nil
    }
    res["preferredKDCs"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetPreferredKDCs(res)
        }
        return nil
    }
    res["realm"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRealm(val)
        }
        return nil
    }
    res["requireUserPresence"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequireUserPresence(val)
        }
        return nil
    }
    res["signInHelpText"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSignInHelpText(val)
        }
        return nil
    }
    res["tlsForLDAPRequired"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTlsForLDAPRequired(val)
        }
        return nil
    }
    res["usernameLabelCustom"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUsernameLabelCustom(val)
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
    res["userSetupDelayed"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserSetupDelayed(val)
        }
        return nil
    }
    return res
}
// GetIsDefaultRealm gets the isDefaultRealm property value. When true, this profile's realm will be selected as the default. Necessary if multiple Kerberos-type profiles are configured.
func (m *MacOSKerberosSingleSignOnExtension) GetIsDefaultRealm()(*bool) {
    val, err := m.GetBackingStore().Get("isDefaultRealm")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetKerberosAppsInBundleIdACLIncluded gets the kerberosAppsInBundleIdACLIncluded property value. When set to True, the Kerberos extension allows any apps entered with the app bundle ID, managed apps, and standard Kerberos utilities, such as TicketViewer and klist, to access and use the credential. Available for devices running macOS versions 12 and later.
func (m *MacOSKerberosSingleSignOnExtension) GetKerberosAppsInBundleIdACLIncluded()(*bool) {
    val, err := m.GetBackingStore().Get("kerberosAppsInBundleIdACLIncluded")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetManagedAppsInBundleIdACLIncluded gets the managedAppsInBundleIdACLIncluded property value. When set to True, the Kerberos extension allows managed apps, and any apps entered with the app bundle ID to access the credential. When set to False, the Kerberos extension allows all apps to access the credential. Available for devices running iOS and iPadOS versions 14 and later.
func (m *MacOSKerberosSingleSignOnExtension) GetManagedAppsInBundleIdACLIncluded()(*bool) {
    val, err := m.GetBackingStore().Get("managedAppsInBundleIdACLIncluded")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetModeCredentialUsed gets the modeCredentialUsed property value. Select how other processes use the Kerberos Extension credential.
func (m *MacOSKerberosSingleSignOnExtension) GetModeCredentialUsed()(*string) {
    val, err := m.GetBackingStore().Get("modeCredentialUsed")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPasswordBlockModification gets the passwordBlockModification property value. Enables or disables password changes.
func (m *MacOSKerberosSingleSignOnExtension) GetPasswordBlockModification()(*bool) {
    val, err := m.GetBackingStore().Get("passwordBlockModification")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetPasswordChangeUrl gets the passwordChangeUrl property value. Gets or sets the URL that the user will be sent to when they initiate a password change.
func (m *MacOSKerberosSingleSignOnExtension) GetPasswordChangeUrl()(*string) {
    val, err := m.GetBackingStore().Get("passwordChangeUrl")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPasswordEnableLocalSync gets the passwordEnableLocalSync property value. Enables or disables password syncing. This won't affect users logged in with a mobile account on macOS.
func (m *MacOSKerberosSingleSignOnExtension) GetPasswordEnableLocalSync()(*bool) {
    val, err := m.GetBackingStore().Get("passwordEnableLocalSync")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetPasswordExpirationDays gets the passwordExpirationDays property value. Overrides the default password expiration in days. For most domains, this value is calculated automatically.
func (m *MacOSKerberosSingleSignOnExtension) GetPasswordExpirationDays()(*int32) {
    val, err := m.GetBackingStore().Get("passwordExpirationDays")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetPasswordExpirationNotificationDays gets the passwordExpirationNotificationDays property value. Gets or sets the number of days until the user is notified that their password will expire (default is 15).
func (m *MacOSKerberosSingleSignOnExtension) GetPasswordExpirationNotificationDays()(*int32) {
    val, err := m.GetBackingStore().Get("passwordExpirationNotificationDays")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetPasswordMinimumAgeDays gets the passwordMinimumAgeDays property value. Gets or sets the minimum number of days until a user can change their password again.
func (m *MacOSKerberosSingleSignOnExtension) GetPasswordMinimumAgeDays()(*int32) {
    val, err := m.GetBackingStore().Get("passwordMinimumAgeDays")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetPasswordMinimumLength gets the passwordMinimumLength property value. Gets or sets the minimum length of a password.
func (m *MacOSKerberosSingleSignOnExtension) GetPasswordMinimumLength()(*int32) {
    val, err := m.GetBackingStore().Get("passwordMinimumLength")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetPasswordPreviousPasswordBlockCount gets the passwordPreviousPasswordBlockCount property value. Gets or sets the number of previous passwords to block.
func (m *MacOSKerberosSingleSignOnExtension) GetPasswordPreviousPasswordBlockCount()(*int32) {
    val, err := m.GetBackingStore().Get("passwordPreviousPasswordBlockCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetPasswordRequireActiveDirectoryComplexity gets the passwordRequireActiveDirectoryComplexity property value. Enables or disables whether passwords must meet Active Directory's complexity requirements.
func (m *MacOSKerberosSingleSignOnExtension) GetPasswordRequireActiveDirectoryComplexity()(*bool) {
    val, err := m.GetBackingStore().Get("passwordRequireActiveDirectoryComplexity")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetPasswordRequirementsDescription gets the passwordRequirementsDescription property value. Gets or sets a description of the password complexity requirements.
func (m *MacOSKerberosSingleSignOnExtension) GetPasswordRequirementsDescription()(*string) {
    val, err := m.GetBackingStore().Get("passwordRequirementsDescription")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPreferredKDCs gets the preferredKDCs property value. Add creates an ordered list of preferred Key Distribution Centers (KDCs) to use for Kerberos traffic. This list is used when the servers are not discoverable using DNS. When the servers are discoverable, the list is used for both connectivity checks, and used first for Kerberos traffic. If the servers don’t respond, then the device uses DNS discovery. Delete removes an existing list, and devices use DNS discovery. Available for devices running macOS versions 12 and later.
func (m *MacOSKerberosSingleSignOnExtension) GetPreferredKDCs()([]string) {
    val, err := m.GetBackingStore().Get("preferredKDCs")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetRealm gets the realm property value. Gets or sets the case-sensitive realm name for this profile.
func (m *MacOSKerberosSingleSignOnExtension) GetRealm()(*string) {
    val, err := m.GetBackingStore().Get("realm")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetRequireUserPresence gets the requireUserPresence property value. Gets or sets whether to require authentication via Touch ID, Face ID, or a passcode to access the keychain entry.
func (m *MacOSKerberosSingleSignOnExtension) GetRequireUserPresence()(*bool) {
    val, err := m.GetBackingStore().Get("requireUserPresence")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetSignInHelpText gets the signInHelpText property value. Text displayed to the user at the Kerberos sign in window. Available for devices running iOS and iPadOS versions 14 and later.
func (m *MacOSKerberosSingleSignOnExtension) GetSignInHelpText()(*string) {
    val, err := m.GetBackingStore().Get("signInHelpText")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetTlsForLDAPRequired gets the tlsForLDAPRequired property value. When set to True, LDAP connections are required to use Transport Layer Security (TLS). Available for devices running macOS versions 11 and later.
func (m *MacOSKerberosSingleSignOnExtension) GetTlsForLDAPRequired()(*bool) {
    val, err := m.GetBackingStore().Get("tlsForLDAPRequired")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetUsernameLabelCustom gets the usernameLabelCustom property value. This label replaces the user name shown in the Kerberos extension. You can enter a name to match the name of your company or organization. Available for devices running macOS versions 11 and later.
func (m *MacOSKerberosSingleSignOnExtension) GetUsernameLabelCustom()(*string) {
    val, err := m.GetBackingStore().Get("usernameLabelCustom")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetUserPrincipalName gets the userPrincipalName property value. Gets or sets the principle user name to use for this profile. The realm name does not need to be included.
func (m *MacOSKerberosSingleSignOnExtension) GetUserPrincipalName()(*string) {
    val, err := m.GetBackingStore().Get("userPrincipalName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetUserSetupDelayed gets the userSetupDelayed property value. When set to True, the user isn’t prompted to set up the Kerberos extension until the extension is enabled by the admin, or a Kerberos challenge is received. Available for devices running macOS versions 11 and later.
func (m *MacOSKerberosSingleSignOnExtension) GetUserSetupDelayed()(*bool) {
    val, err := m.GetBackingStore().Get("userSetupDelayed")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// Serialize serializes information the current object
func (m *MacOSKerberosSingleSignOnExtension) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.MacOSSingleSignOnExtension.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("activeDirectorySiteCode", m.GetActiveDirectorySiteCode())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("blockActiveDirectorySiteAutoDiscovery", m.GetBlockActiveDirectorySiteAutoDiscovery())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("blockAutomaticLogin", m.GetBlockAutomaticLogin())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("cacheName", m.GetCacheName())
        if err != nil {
            return err
        }
    }
    if m.GetCredentialBundleIdAccessControlList() != nil {
        err = writer.WriteCollectionOfStringValues("credentialBundleIdAccessControlList", m.GetCredentialBundleIdAccessControlList())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("credentialsCacheMonitored", m.GetCredentialsCacheMonitored())
        if err != nil {
            return err
        }
    }
    if m.GetDomainRealms() != nil {
        err = writer.WriteCollectionOfStringValues("domainRealms", m.GetDomainRealms())
        if err != nil {
            return err
        }
    }
    if m.GetDomains() != nil {
        err = writer.WriteCollectionOfStringValues("domains", m.GetDomains())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isDefaultRealm", m.GetIsDefaultRealm())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("kerberosAppsInBundleIdACLIncluded", m.GetKerberosAppsInBundleIdACLIncluded())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("managedAppsInBundleIdACLIncluded", m.GetManagedAppsInBundleIdACLIncluded())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("modeCredentialUsed", m.GetModeCredentialUsed())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("passwordBlockModification", m.GetPasswordBlockModification())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("passwordChangeUrl", m.GetPasswordChangeUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("passwordEnableLocalSync", m.GetPasswordEnableLocalSync())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("passwordExpirationDays", m.GetPasswordExpirationDays())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("passwordExpirationNotificationDays", m.GetPasswordExpirationNotificationDays())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("passwordMinimumAgeDays", m.GetPasswordMinimumAgeDays())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("passwordMinimumLength", m.GetPasswordMinimumLength())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("passwordPreviousPasswordBlockCount", m.GetPasswordPreviousPasswordBlockCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("passwordRequireActiveDirectoryComplexity", m.GetPasswordRequireActiveDirectoryComplexity())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("passwordRequirementsDescription", m.GetPasswordRequirementsDescription())
        if err != nil {
            return err
        }
    }
    if m.GetPreferredKDCs() != nil {
        err = writer.WriteCollectionOfStringValues("preferredKDCs", m.GetPreferredKDCs())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("realm", m.GetRealm())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("requireUserPresence", m.GetRequireUserPresence())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("signInHelpText", m.GetSignInHelpText())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("tlsForLDAPRequired", m.GetTlsForLDAPRequired())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("usernameLabelCustom", m.GetUsernameLabelCustom())
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
        err = writer.WriteBoolValue("userSetupDelayed", m.GetUserSetupDelayed())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetActiveDirectorySiteCode sets the activeDirectorySiteCode property value. Gets or sets the Active Directory site.
func (m *MacOSKerberosSingleSignOnExtension) SetActiveDirectorySiteCode(value *string)() {
    err := m.GetBackingStore().Set("activeDirectorySiteCode", value)
    if err != nil {
        panic(err)
    }
}
// SetBlockActiveDirectorySiteAutoDiscovery sets the blockActiveDirectorySiteAutoDiscovery property value. Enables or disables whether the Kerberos extension can automatically determine its site name.
func (m *MacOSKerberosSingleSignOnExtension) SetBlockActiveDirectorySiteAutoDiscovery(value *bool)() {
    err := m.GetBackingStore().Set("blockActiveDirectorySiteAutoDiscovery", value)
    if err != nil {
        panic(err)
    }
}
// SetBlockAutomaticLogin sets the blockAutomaticLogin property value. Enables or disables Keychain usage.
func (m *MacOSKerberosSingleSignOnExtension) SetBlockAutomaticLogin(value *bool)() {
    err := m.GetBackingStore().Set("blockAutomaticLogin", value)
    if err != nil {
        panic(err)
    }
}
// SetCacheName sets the cacheName property value. Gets or sets the Generic Security Services name of the Kerberos cache to use for this profile.
func (m *MacOSKerberosSingleSignOnExtension) SetCacheName(value *string)() {
    err := m.GetBackingStore().Set("cacheName", value)
    if err != nil {
        panic(err)
    }
}
// SetCredentialBundleIdAccessControlList sets the credentialBundleIdAccessControlList property value. Gets or sets a list of app Bundle IDs allowed to access the Kerberos Ticket Granting Ticket.
func (m *MacOSKerberosSingleSignOnExtension) SetCredentialBundleIdAccessControlList(value []string)() {
    err := m.GetBackingStore().Set("credentialBundleIdAccessControlList", value)
    if err != nil {
        panic(err)
    }
}
// SetCredentialsCacheMonitored sets the credentialsCacheMonitored property value. When set to True, the credential is requested on the next matching Kerberos challenge or network state change. When the credential is expired or missing, a new credential is created. Available for devices running macOS versions 12 and later.
func (m *MacOSKerberosSingleSignOnExtension) SetCredentialsCacheMonitored(value *bool)() {
    err := m.GetBackingStore().Set("credentialsCacheMonitored", value)
    if err != nil {
        panic(err)
    }
}
// SetDomainRealms sets the domainRealms property value. Gets or sets a list of realms for custom domain-realm mapping. Realms are case sensitive.
func (m *MacOSKerberosSingleSignOnExtension) SetDomainRealms(value []string)() {
    err := m.GetBackingStore().Set("domainRealms", value)
    if err != nil {
        panic(err)
    }
}
// SetDomains sets the domains property value. Gets or sets a list of hosts or domain names for which the app extension performs SSO.
func (m *MacOSKerberosSingleSignOnExtension) SetDomains(value []string)() {
    err := m.GetBackingStore().Set("domains", value)
    if err != nil {
        panic(err)
    }
}
// SetIsDefaultRealm sets the isDefaultRealm property value. When true, this profile's realm will be selected as the default. Necessary if multiple Kerberos-type profiles are configured.
func (m *MacOSKerberosSingleSignOnExtension) SetIsDefaultRealm(value *bool)() {
    err := m.GetBackingStore().Set("isDefaultRealm", value)
    if err != nil {
        panic(err)
    }
}
// SetKerberosAppsInBundleIdACLIncluded sets the kerberosAppsInBundleIdACLIncluded property value. When set to True, the Kerberos extension allows any apps entered with the app bundle ID, managed apps, and standard Kerberos utilities, such as TicketViewer and klist, to access and use the credential. Available for devices running macOS versions 12 and later.
func (m *MacOSKerberosSingleSignOnExtension) SetKerberosAppsInBundleIdACLIncluded(value *bool)() {
    err := m.GetBackingStore().Set("kerberosAppsInBundleIdACLIncluded", value)
    if err != nil {
        panic(err)
    }
}
// SetManagedAppsInBundleIdACLIncluded sets the managedAppsInBundleIdACLIncluded property value. When set to True, the Kerberos extension allows managed apps, and any apps entered with the app bundle ID to access the credential. When set to False, the Kerberos extension allows all apps to access the credential. Available for devices running iOS and iPadOS versions 14 and later.
func (m *MacOSKerberosSingleSignOnExtension) SetManagedAppsInBundleIdACLIncluded(value *bool)() {
    err := m.GetBackingStore().Set("managedAppsInBundleIdACLIncluded", value)
    if err != nil {
        panic(err)
    }
}
// SetModeCredentialUsed sets the modeCredentialUsed property value. Select how other processes use the Kerberos Extension credential.
func (m *MacOSKerberosSingleSignOnExtension) SetModeCredentialUsed(value *string)() {
    err := m.GetBackingStore().Set("modeCredentialUsed", value)
    if err != nil {
        panic(err)
    }
}
// SetPasswordBlockModification sets the passwordBlockModification property value. Enables or disables password changes.
func (m *MacOSKerberosSingleSignOnExtension) SetPasswordBlockModification(value *bool)() {
    err := m.GetBackingStore().Set("passwordBlockModification", value)
    if err != nil {
        panic(err)
    }
}
// SetPasswordChangeUrl sets the passwordChangeUrl property value. Gets or sets the URL that the user will be sent to when they initiate a password change.
func (m *MacOSKerberosSingleSignOnExtension) SetPasswordChangeUrl(value *string)() {
    err := m.GetBackingStore().Set("passwordChangeUrl", value)
    if err != nil {
        panic(err)
    }
}
// SetPasswordEnableLocalSync sets the passwordEnableLocalSync property value. Enables or disables password syncing. This won't affect users logged in with a mobile account on macOS.
func (m *MacOSKerberosSingleSignOnExtension) SetPasswordEnableLocalSync(value *bool)() {
    err := m.GetBackingStore().Set("passwordEnableLocalSync", value)
    if err != nil {
        panic(err)
    }
}
// SetPasswordExpirationDays sets the passwordExpirationDays property value. Overrides the default password expiration in days. For most domains, this value is calculated automatically.
func (m *MacOSKerberosSingleSignOnExtension) SetPasswordExpirationDays(value *int32)() {
    err := m.GetBackingStore().Set("passwordExpirationDays", value)
    if err != nil {
        panic(err)
    }
}
// SetPasswordExpirationNotificationDays sets the passwordExpirationNotificationDays property value. Gets or sets the number of days until the user is notified that their password will expire (default is 15).
func (m *MacOSKerberosSingleSignOnExtension) SetPasswordExpirationNotificationDays(value *int32)() {
    err := m.GetBackingStore().Set("passwordExpirationNotificationDays", value)
    if err != nil {
        panic(err)
    }
}
// SetPasswordMinimumAgeDays sets the passwordMinimumAgeDays property value. Gets or sets the minimum number of days until a user can change their password again.
func (m *MacOSKerberosSingleSignOnExtension) SetPasswordMinimumAgeDays(value *int32)() {
    err := m.GetBackingStore().Set("passwordMinimumAgeDays", value)
    if err != nil {
        panic(err)
    }
}
// SetPasswordMinimumLength sets the passwordMinimumLength property value. Gets or sets the minimum length of a password.
func (m *MacOSKerberosSingleSignOnExtension) SetPasswordMinimumLength(value *int32)() {
    err := m.GetBackingStore().Set("passwordMinimumLength", value)
    if err != nil {
        panic(err)
    }
}
// SetPasswordPreviousPasswordBlockCount sets the passwordPreviousPasswordBlockCount property value. Gets or sets the number of previous passwords to block.
func (m *MacOSKerberosSingleSignOnExtension) SetPasswordPreviousPasswordBlockCount(value *int32)() {
    err := m.GetBackingStore().Set("passwordPreviousPasswordBlockCount", value)
    if err != nil {
        panic(err)
    }
}
// SetPasswordRequireActiveDirectoryComplexity sets the passwordRequireActiveDirectoryComplexity property value. Enables or disables whether passwords must meet Active Directory's complexity requirements.
func (m *MacOSKerberosSingleSignOnExtension) SetPasswordRequireActiveDirectoryComplexity(value *bool)() {
    err := m.GetBackingStore().Set("passwordRequireActiveDirectoryComplexity", value)
    if err != nil {
        panic(err)
    }
}
// SetPasswordRequirementsDescription sets the passwordRequirementsDescription property value. Gets or sets a description of the password complexity requirements.
func (m *MacOSKerberosSingleSignOnExtension) SetPasswordRequirementsDescription(value *string)() {
    err := m.GetBackingStore().Set("passwordRequirementsDescription", value)
    if err != nil {
        panic(err)
    }
}
// SetPreferredKDCs sets the preferredKDCs property value. Add creates an ordered list of preferred Key Distribution Centers (KDCs) to use for Kerberos traffic. This list is used when the servers are not discoverable using DNS. When the servers are discoverable, the list is used for both connectivity checks, and used first for Kerberos traffic. If the servers don’t respond, then the device uses DNS discovery. Delete removes an existing list, and devices use DNS discovery. Available for devices running macOS versions 12 and later.
func (m *MacOSKerberosSingleSignOnExtension) SetPreferredKDCs(value []string)() {
    err := m.GetBackingStore().Set("preferredKDCs", value)
    if err != nil {
        panic(err)
    }
}
// SetRealm sets the realm property value. Gets or sets the case-sensitive realm name for this profile.
func (m *MacOSKerberosSingleSignOnExtension) SetRealm(value *string)() {
    err := m.GetBackingStore().Set("realm", value)
    if err != nil {
        panic(err)
    }
}
// SetRequireUserPresence sets the requireUserPresence property value. Gets or sets whether to require authentication via Touch ID, Face ID, or a passcode to access the keychain entry.
func (m *MacOSKerberosSingleSignOnExtension) SetRequireUserPresence(value *bool)() {
    err := m.GetBackingStore().Set("requireUserPresence", value)
    if err != nil {
        panic(err)
    }
}
// SetSignInHelpText sets the signInHelpText property value. Text displayed to the user at the Kerberos sign in window. Available for devices running iOS and iPadOS versions 14 and later.
func (m *MacOSKerberosSingleSignOnExtension) SetSignInHelpText(value *string)() {
    err := m.GetBackingStore().Set("signInHelpText", value)
    if err != nil {
        panic(err)
    }
}
// SetTlsForLDAPRequired sets the tlsForLDAPRequired property value. When set to True, LDAP connections are required to use Transport Layer Security (TLS). Available for devices running macOS versions 11 and later.
func (m *MacOSKerberosSingleSignOnExtension) SetTlsForLDAPRequired(value *bool)() {
    err := m.GetBackingStore().Set("tlsForLDAPRequired", value)
    if err != nil {
        panic(err)
    }
}
// SetUsernameLabelCustom sets the usernameLabelCustom property value. This label replaces the user name shown in the Kerberos extension. You can enter a name to match the name of your company or organization. Available for devices running macOS versions 11 and later.
func (m *MacOSKerberosSingleSignOnExtension) SetUsernameLabelCustom(value *string)() {
    err := m.GetBackingStore().Set("usernameLabelCustom", value)
    if err != nil {
        panic(err)
    }
}
// SetUserPrincipalName sets the userPrincipalName property value. Gets or sets the principle user name to use for this profile. The realm name does not need to be included.
func (m *MacOSKerberosSingleSignOnExtension) SetUserPrincipalName(value *string)() {
    err := m.GetBackingStore().Set("userPrincipalName", value)
    if err != nil {
        panic(err)
    }
}
// SetUserSetupDelayed sets the userSetupDelayed property value. When set to True, the user isn’t prompted to set up the Kerberos extension until the extension is enabled by the admin, or a Kerberos challenge is received. Available for devices running macOS versions 11 and later.
func (m *MacOSKerberosSingleSignOnExtension) SetUserSetupDelayed(value *bool)() {
    err := m.GetBackingStore().Set("userSetupDelayed", value)
    if err != nil {
        panic(err)
    }
}
// MacOSKerberosSingleSignOnExtensionable 
type MacOSKerberosSingleSignOnExtensionable interface {
    MacOSSingleSignOnExtensionable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActiveDirectorySiteCode()(*string)
    GetBlockActiveDirectorySiteAutoDiscovery()(*bool)
    GetBlockAutomaticLogin()(*bool)
    GetCacheName()(*string)
    GetCredentialBundleIdAccessControlList()([]string)
    GetCredentialsCacheMonitored()(*bool)
    GetDomainRealms()([]string)
    GetDomains()([]string)
    GetIsDefaultRealm()(*bool)
    GetKerberosAppsInBundleIdACLIncluded()(*bool)
    GetManagedAppsInBundleIdACLIncluded()(*bool)
    GetModeCredentialUsed()(*string)
    GetPasswordBlockModification()(*bool)
    GetPasswordChangeUrl()(*string)
    GetPasswordEnableLocalSync()(*bool)
    GetPasswordExpirationDays()(*int32)
    GetPasswordExpirationNotificationDays()(*int32)
    GetPasswordMinimumAgeDays()(*int32)
    GetPasswordMinimumLength()(*int32)
    GetPasswordPreviousPasswordBlockCount()(*int32)
    GetPasswordRequireActiveDirectoryComplexity()(*bool)
    GetPasswordRequirementsDescription()(*string)
    GetPreferredKDCs()([]string)
    GetRealm()(*string)
    GetRequireUserPresence()(*bool)
    GetSignInHelpText()(*string)
    GetTlsForLDAPRequired()(*bool)
    GetUsernameLabelCustom()(*string)
    GetUserPrincipalName()(*string)
    GetUserSetupDelayed()(*bool)
    SetActiveDirectorySiteCode(value *string)()
    SetBlockActiveDirectorySiteAutoDiscovery(value *bool)()
    SetBlockAutomaticLogin(value *bool)()
    SetCacheName(value *string)()
    SetCredentialBundleIdAccessControlList(value []string)()
    SetCredentialsCacheMonitored(value *bool)()
    SetDomainRealms(value []string)()
    SetDomains(value []string)()
    SetIsDefaultRealm(value *bool)()
    SetKerberosAppsInBundleIdACLIncluded(value *bool)()
    SetManagedAppsInBundleIdACLIncluded(value *bool)()
    SetModeCredentialUsed(value *string)()
    SetPasswordBlockModification(value *bool)()
    SetPasswordChangeUrl(value *string)()
    SetPasswordEnableLocalSync(value *bool)()
    SetPasswordExpirationDays(value *int32)()
    SetPasswordExpirationNotificationDays(value *int32)()
    SetPasswordMinimumAgeDays(value *int32)()
    SetPasswordMinimumLength(value *int32)()
    SetPasswordPreviousPasswordBlockCount(value *int32)()
    SetPasswordRequireActiveDirectoryComplexity(value *bool)()
    SetPasswordRequirementsDescription(value *string)()
    SetPreferredKDCs(value []string)()
    SetRealm(value *string)()
    SetRequireUserPresence(value *bool)()
    SetSignInHelpText(value *string)()
    SetTlsForLDAPRequired(value *bool)()
    SetUsernameLabelCustom(value *string)()
    SetUserPrincipalName(value *string)()
    SetUserSetupDelayed(value *bool)()
}
