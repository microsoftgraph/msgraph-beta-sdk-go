package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DepMacOSEnrollmentProfile 
type DepMacOSEnrollmentProfile struct {
    DepEnrollmentBaseProfile
}
// NewDepMacOSEnrollmentProfile instantiates a new DepMacOSEnrollmentProfile and sets the default values.
func NewDepMacOSEnrollmentProfile()(*DepMacOSEnrollmentProfile) {
    m := &DepMacOSEnrollmentProfile{
        DepEnrollmentBaseProfile: *NewDepEnrollmentBaseProfile(),
    }
    odataTypeValue := "#microsoft.graph.depMacOSEnrollmentProfile"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateDepMacOSEnrollmentProfileFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDepMacOSEnrollmentProfileFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDepMacOSEnrollmentProfile(), nil
}
// GetAccessibilityScreenDisabled gets the accessibilityScreenDisabled property value. Indicates if Accessibility screen is disabled
func (m *DepMacOSEnrollmentProfile) GetAccessibilityScreenDisabled()(*bool) {
    val, err := m.GetBackingStore().Get("accessibilityScreenDisabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetAdminAccountFullName gets the adminAccountFullName property value. Indicates what the full name for the admin account is
func (m *DepMacOSEnrollmentProfile) GetAdminAccountFullName()(*string) {
    val, err := m.GetBackingStore().Get("adminAccountFullName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetAdminAccountPassword gets the adminAccountPassword property value. Indicates what the password for the admin account is
func (m *DepMacOSEnrollmentProfile) GetAdminAccountPassword()(*string) {
    val, err := m.GetBackingStore().Get("adminAccountPassword")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetAdminAccountUserName gets the adminAccountUserName property value. Indicates what the user name for the admin account is
func (m *DepMacOSEnrollmentProfile) GetAdminAccountUserName()(*string) {
    val, err := m.GetBackingStore().Get("adminAccountUserName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetAutoAdvanceSetupEnabled gets the autoAdvanceSetupEnabled property value. Indicates if Setup Assistant will automatically advance through its screen
func (m *DepMacOSEnrollmentProfile) GetAutoAdvanceSetupEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("autoAdvanceSetupEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetAutoUnlockWithWatchDisabled gets the autoUnlockWithWatchDisabled property value. Indicates if UnlockWithWatch screen is disabled
func (m *DepMacOSEnrollmentProfile) GetAutoUnlockWithWatchDisabled()(*bool) {
    val, err := m.GetBackingStore().Get("autoUnlockWithWatchDisabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetChooseYourLockScreenDisabled gets the chooseYourLockScreenDisabled property value. Indicates if iCloud Documents and Desktop screen is disabled
func (m *DepMacOSEnrollmentProfile) GetChooseYourLockScreenDisabled()(*bool) {
    val, err := m.GetBackingStore().Get("chooseYourLockScreenDisabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetDontAutoPopulatePrimaryAccountInfo gets the dontAutoPopulatePrimaryAccountInfo property value. Indicates whether Setup Assistant will auto populate the primary account information
func (m *DepMacOSEnrollmentProfile) GetDontAutoPopulatePrimaryAccountInfo()(*bool) {
    val, err := m.GetBackingStore().Get("dontAutoPopulatePrimaryAccountInfo")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetEnableRestrictEditing gets the enableRestrictEditing property value. Indicates whether the user will enable blockediting
func (m *DepMacOSEnrollmentProfile) GetEnableRestrictEditing()(*bool) {
    val, err := m.GetBackingStore().Get("enableRestrictEditing")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DepMacOSEnrollmentProfile) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DepEnrollmentBaseProfile.GetFieldDeserializers()
    res["accessibilityScreenDisabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccessibilityScreenDisabled(val)
        }
        return nil
    }
    res["adminAccountFullName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdminAccountFullName(val)
        }
        return nil
    }
    res["adminAccountPassword"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdminAccountPassword(val)
        }
        return nil
    }
    res["adminAccountUserName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdminAccountUserName(val)
        }
        return nil
    }
    res["autoAdvanceSetupEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAutoAdvanceSetupEnabled(val)
        }
        return nil
    }
    res["autoUnlockWithWatchDisabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAutoUnlockWithWatchDisabled(val)
        }
        return nil
    }
    res["chooseYourLockScreenDisabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetChooseYourLockScreenDisabled(val)
        }
        return nil
    }
    res["dontAutoPopulatePrimaryAccountInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDontAutoPopulatePrimaryAccountInfo(val)
        }
        return nil
    }
    res["enableRestrictEditing"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnableRestrictEditing(val)
        }
        return nil
    }
    res["fileVaultDisabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFileVaultDisabled(val)
        }
        return nil
    }
    res["hideAdminAccount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHideAdminAccount(val)
        }
        return nil
    }
    res["iCloudDiagnosticsDisabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetICloudDiagnosticsDisabled(val)
        }
        return nil
    }
    res["iCloudStorageDisabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetICloudStorageDisabled(val)
        }
        return nil
    }
    res["passCodeDisabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPassCodeDisabled(val)
        }
        return nil
    }
    res["primaryAccountFullName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrimaryAccountFullName(val)
        }
        return nil
    }
    res["primaryAccountUserName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrimaryAccountUserName(val)
        }
        return nil
    }
    res["registrationDisabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRegistrationDisabled(val)
        }
        return nil
    }
    res["requestRequiresNetworkTether"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequestRequiresNetworkTether(val)
        }
        return nil
    }
    res["setPrimarySetupAccountAsRegularUser"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSetPrimarySetupAccountAsRegularUser(val)
        }
        return nil
    }
    res["skipPrimarySetupAccountCreation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSkipPrimarySetupAccountCreation(val)
        }
        return nil
    }
    res["zoomDisabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetZoomDisabled(val)
        }
        return nil
    }
    return res
}
// GetFileVaultDisabled gets the fileVaultDisabled property value. Indicates if file vault is disabled
func (m *DepMacOSEnrollmentProfile) GetFileVaultDisabled()(*bool) {
    val, err := m.GetBackingStore().Get("fileVaultDisabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetHideAdminAccount gets the hideAdminAccount property value. Indicates whether the admin account should be hidded or not
func (m *DepMacOSEnrollmentProfile) GetHideAdminAccount()(*bool) {
    val, err := m.GetBackingStore().Get("hideAdminAccount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetICloudDiagnosticsDisabled gets the iCloudDiagnosticsDisabled property value. Indicates if iCloud Analytics screen is disabled
func (m *DepMacOSEnrollmentProfile) GetICloudDiagnosticsDisabled()(*bool) {
    val, err := m.GetBackingStore().Get("iCloudDiagnosticsDisabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetICloudStorageDisabled gets the iCloudStorageDisabled property value. Indicates if iCloud Documents and Desktop screen is disabled
func (m *DepMacOSEnrollmentProfile) GetICloudStorageDisabled()(*bool) {
    val, err := m.GetBackingStore().Get("iCloudStorageDisabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetPassCodeDisabled gets the passCodeDisabled property value. Indicates if Passcode setup pane is disabled
func (m *DepMacOSEnrollmentProfile) GetPassCodeDisabled()(*bool) {
    val, err := m.GetBackingStore().Get("passCodeDisabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetPrimaryAccountFullName gets the primaryAccountFullName property value. Indicates what the full name for the primary account is
func (m *DepMacOSEnrollmentProfile) GetPrimaryAccountFullName()(*string) {
    val, err := m.GetBackingStore().Get("primaryAccountFullName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPrimaryAccountUserName gets the primaryAccountUserName property value. Indicates what the account name for the primary account is
func (m *DepMacOSEnrollmentProfile) GetPrimaryAccountUserName()(*string) {
    val, err := m.GetBackingStore().Get("primaryAccountUserName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetRegistrationDisabled gets the registrationDisabled property value. Indicates if registration is disabled
func (m *DepMacOSEnrollmentProfile) GetRegistrationDisabled()(*bool) {
    val, err := m.GetBackingStore().Get("registrationDisabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetRequestRequiresNetworkTether gets the requestRequiresNetworkTether property value. Indicates if the device is network-tethered to run the command
func (m *DepMacOSEnrollmentProfile) GetRequestRequiresNetworkTether()(*bool) {
    val, err := m.GetBackingStore().Get("requestRequiresNetworkTether")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetSetPrimarySetupAccountAsRegularUser gets the setPrimarySetupAccountAsRegularUser property value. Indicates whether Setup Assistant will set the account as a regular user
func (m *DepMacOSEnrollmentProfile) GetSetPrimarySetupAccountAsRegularUser()(*bool) {
    val, err := m.GetBackingStore().Get("setPrimarySetupAccountAsRegularUser")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetSkipPrimarySetupAccountCreation gets the skipPrimarySetupAccountCreation property value. Indicates whether Setup Assistant will skip the user interface for primary account setup
func (m *DepMacOSEnrollmentProfile) GetSkipPrimarySetupAccountCreation()(*bool) {
    val, err := m.GetBackingStore().Get("skipPrimarySetupAccountCreation")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetZoomDisabled gets the zoomDisabled property value. Indicates if zoom setup pane is disabled
func (m *DepMacOSEnrollmentProfile) GetZoomDisabled()(*bool) {
    val, err := m.GetBackingStore().Get("zoomDisabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// Serialize serializes information the current object
func (m *DepMacOSEnrollmentProfile) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DepEnrollmentBaseProfile.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("accessibilityScreenDisabled", m.GetAccessibilityScreenDisabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("adminAccountFullName", m.GetAdminAccountFullName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("adminAccountPassword", m.GetAdminAccountPassword())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("adminAccountUserName", m.GetAdminAccountUserName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("autoAdvanceSetupEnabled", m.GetAutoAdvanceSetupEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("autoUnlockWithWatchDisabled", m.GetAutoUnlockWithWatchDisabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("chooseYourLockScreenDisabled", m.GetChooseYourLockScreenDisabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("dontAutoPopulatePrimaryAccountInfo", m.GetDontAutoPopulatePrimaryAccountInfo())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("enableRestrictEditing", m.GetEnableRestrictEditing())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("fileVaultDisabled", m.GetFileVaultDisabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("hideAdminAccount", m.GetHideAdminAccount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("iCloudDiagnosticsDisabled", m.GetICloudDiagnosticsDisabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("iCloudStorageDisabled", m.GetICloudStorageDisabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("passCodeDisabled", m.GetPassCodeDisabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("primaryAccountFullName", m.GetPrimaryAccountFullName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("primaryAccountUserName", m.GetPrimaryAccountUserName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("registrationDisabled", m.GetRegistrationDisabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("requestRequiresNetworkTether", m.GetRequestRequiresNetworkTether())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("setPrimarySetupAccountAsRegularUser", m.GetSetPrimarySetupAccountAsRegularUser())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("skipPrimarySetupAccountCreation", m.GetSkipPrimarySetupAccountCreation())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("zoomDisabled", m.GetZoomDisabled())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAccessibilityScreenDisabled sets the accessibilityScreenDisabled property value. Indicates if Accessibility screen is disabled
func (m *DepMacOSEnrollmentProfile) SetAccessibilityScreenDisabled(value *bool)() {
    err := m.GetBackingStore().Set("accessibilityScreenDisabled", value)
    if err != nil {
        panic(err)
    }
}
// SetAdminAccountFullName sets the adminAccountFullName property value. Indicates what the full name for the admin account is
func (m *DepMacOSEnrollmentProfile) SetAdminAccountFullName(value *string)() {
    err := m.GetBackingStore().Set("adminAccountFullName", value)
    if err != nil {
        panic(err)
    }
}
// SetAdminAccountPassword sets the adminAccountPassword property value. Indicates what the password for the admin account is
func (m *DepMacOSEnrollmentProfile) SetAdminAccountPassword(value *string)() {
    err := m.GetBackingStore().Set("adminAccountPassword", value)
    if err != nil {
        panic(err)
    }
}
// SetAdminAccountUserName sets the adminAccountUserName property value. Indicates what the user name for the admin account is
func (m *DepMacOSEnrollmentProfile) SetAdminAccountUserName(value *string)() {
    err := m.GetBackingStore().Set("adminAccountUserName", value)
    if err != nil {
        panic(err)
    }
}
// SetAutoAdvanceSetupEnabled sets the autoAdvanceSetupEnabled property value. Indicates if Setup Assistant will automatically advance through its screen
func (m *DepMacOSEnrollmentProfile) SetAutoAdvanceSetupEnabled(value *bool)() {
    err := m.GetBackingStore().Set("autoAdvanceSetupEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetAutoUnlockWithWatchDisabled sets the autoUnlockWithWatchDisabled property value. Indicates if UnlockWithWatch screen is disabled
func (m *DepMacOSEnrollmentProfile) SetAutoUnlockWithWatchDisabled(value *bool)() {
    err := m.GetBackingStore().Set("autoUnlockWithWatchDisabled", value)
    if err != nil {
        panic(err)
    }
}
// SetChooseYourLockScreenDisabled sets the chooseYourLockScreenDisabled property value. Indicates if iCloud Documents and Desktop screen is disabled
func (m *DepMacOSEnrollmentProfile) SetChooseYourLockScreenDisabled(value *bool)() {
    err := m.GetBackingStore().Set("chooseYourLockScreenDisabled", value)
    if err != nil {
        panic(err)
    }
}
// SetDontAutoPopulatePrimaryAccountInfo sets the dontAutoPopulatePrimaryAccountInfo property value. Indicates whether Setup Assistant will auto populate the primary account information
func (m *DepMacOSEnrollmentProfile) SetDontAutoPopulatePrimaryAccountInfo(value *bool)() {
    err := m.GetBackingStore().Set("dontAutoPopulatePrimaryAccountInfo", value)
    if err != nil {
        panic(err)
    }
}
// SetEnableRestrictEditing sets the enableRestrictEditing property value. Indicates whether the user will enable blockediting
func (m *DepMacOSEnrollmentProfile) SetEnableRestrictEditing(value *bool)() {
    err := m.GetBackingStore().Set("enableRestrictEditing", value)
    if err != nil {
        panic(err)
    }
}
// SetFileVaultDisabled sets the fileVaultDisabled property value. Indicates if file vault is disabled
func (m *DepMacOSEnrollmentProfile) SetFileVaultDisabled(value *bool)() {
    err := m.GetBackingStore().Set("fileVaultDisabled", value)
    if err != nil {
        panic(err)
    }
}
// SetHideAdminAccount sets the hideAdminAccount property value. Indicates whether the admin account should be hidded or not
func (m *DepMacOSEnrollmentProfile) SetHideAdminAccount(value *bool)() {
    err := m.GetBackingStore().Set("hideAdminAccount", value)
    if err != nil {
        panic(err)
    }
}
// SetICloudDiagnosticsDisabled sets the iCloudDiagnosticsDisabled property value. Indicates if iCloud Analytics screen is disabled
func (m *DepMacOSEnrollmentProfile) SetICloudDiagnosticsDisabled(value *bool)() {
    err := m.GetBackingStore().Set("iCloudDiagnosticsDisabled", value)
    if err != nil {
        panic(err)
    }
}
// SetICloudStorageDisabled sets the iCloudStorageDisabled property value. Indicates if iCloud Documents and Desktop screen is disabled
func (m *DepMacOSEnrollmentProfile) SetICloudStorageDisabled(value *bool)() {
    err := m.GetBackingStore().Set("iCloudStorageDisabled", value)
    if err != nil {
        panic(err)
    }
}
// SetPassCodeDisabled sets the passCodeDisabled property value. Indicates if Passcode setup pane is disabled
func (m *DepMacOSEnrollmentProfile) SetPassCodeDisabled(value *bool)() {
    err := m.GetBackingStore().Set("passCodeDisabled", value)
    if err != nil {
        panic(err)
    }
}
// SetPrimaryAccountFullName sets the primaryAccountFullName property value. Indicates what the full name for the primary account is
func (m *DepMacOSEnrollmentProfile) SetPrimaryAccountFullName(value *string)() {
    err := m.GetBackingStore().Set("primaryAccountFullName", value)
    if err != nil {
        panic(err)
    }
}
// SetPrimaryAccountUserName sets the primaryAccountUserName property value. Indicates what the account name for the primary account is
func (m *DepMacOSEnrollmentProfile) SetPrimaryAccountUserName(value *string)() {
    err := m.GetBackingStore().Set("primaryAccountUserName", value)
    if err != nil {
        panic(err)
    }
}
// SetRegistrationDisabled sets the registrationDisabled property value. Indicates if registration is disabled
func (m *DepMacOSEnrollmentProfile) SetRegistrationDisabled(value *bool)() {
    err := m.GetBackingStore().Set("registrationDisabled", value)
    if err != nil {
        panic(err)
    }
}
// SetRequestRequiresNetworkTether sets the requestRequiresNetworkTether property value. Indicates if the device is network-tethered to run the command
func (m *DepMacOSEnrollmentProfile) SetRequestRequiresNetworkTether(value *bool)() {
    err := m.GetBackingStore().Set("requestRequiresNetworkTether", value)
    if err != nil {
        panic(err)
    }
}
// SetSetPrimarySetupAccountAsRegularUser sets the setPrimarySetupAccountAsRegularUser property value. Indicates whether Setup Assistant will set the account as a regular user
func (m *DepMacOSEnrollmentProfile) SetSetPrimarySetupAccountAsRegularUser(value *bool)() {
    err := m.GetBackingStore().Set("setPrimarySetupAccountAsRegularUser", value)
    if err != nil {
        panic(err)
    }
}
// SetSkipPrimarySetupAccountCreation sets the skipPrimarySetupAccountCreation property value. Indicates whether Setup Assistant will skip the user interface for primary account setup
func (m *DepMacOSEnrollmentProfile) SetSkipPrimarySetupAccountCreation(value *bool)() {
    err := m.GetBackingStore().Set("skipPrimarySetupAccountCreation", value)
    if err != nil {
        panic(err)
    }
}
// SetZoomDisabled sets the zoomDisabled property value. Indicates if zoom setup pane is disabled
func (m *DepMacOSEnrollmentProfile) SetZoomDisabled(value *bool)() {
    err := m.GetBackingStore().Set("zoomDisabled", value)
    if err != nil {
        panic(err)
    }
}
// DepMacOSEnrollmentProfileable 
type DepMacOSEnrollmentProfileable interface {
    DepEnrollmentBaseProfileable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccessibilityScreenDisabled()(*bool)
    GetAdminAccountFullName()(*string)
    GetAdminAccountPassword()(*string)
    GetAdminAccountUserName()(*string)
    GetAutoAdvanceSetupEnabled()(*bool)
    GetAutoUnlockWithWatchDisabled()(*bool)
    GetChooseYourLockScreenDisabled()(*bool)
    GetDontAutoPopulatePrimaryAccountInfo()(*bool)
    GetEnableRestrictEditing()(*bool)
    GetFileVaultDisabled()(*bool)
    GetHideAdminAccount()(*bool)
    GetICloudDiagnosticsDisabled()(*bool)
    GetICloudStorageDisabled()(*bool)
    GetPassCodeDisabled()(*bool)
    GetPrimaryAccountFullName()(*string)
    GetPrimaryAccountUserName()(*string)
    GetRegistrationDisabled()(*bool)
    GetRequestRequiresNetworkTether()(*bool)
    GetSetPrimarySetupAccountAsRegularUser()(*bool)
    GetSkipPrimarySetupAccountCreation()(*bool)
    GetZoomDisabled()(*bool)
    SetAccessibilityScreenDisabled(value *bool)()
    SetAdminAccountFullName(value *string)()
    SetAdminAccountPassword(value *string)()
    SetAdminAccountUserName(value *string)()
    SetAutoAdvanceSetupEnabled(value *bool)()
    SetAutoUnlockWithWatchDisabled(value *bool)()
    SetChooseYourLockScreenDisabled(value *bool)()
    SetDontAutoPopulatePrimaryAccountInfo(value *bool)()
    SetEnableRestrictEditing(value *bool)()
    SetFileVaultDisabled(value *bool)()
    SetHideAdminAccount(value *bool)()
    SetICloudDiagnosticsDisabled(value *bool)()
    SetICloudStorageDisabled(value *bool)()
    SetPassCodeDisabled(value *bool)()
    SetPrimaryAccountFullName(value *string)()
    SetPrimaryAccountUserName(value *string)()
    SetRegistrationDisabled(value *bool)()
    SetRequestRequiresNetworkTether(value *bool)()
    SetSetPrimarySetupAccountAsRegularUser(value *bool)()
    SetSkipPrimarySetupAccountCreation(value *bool)()
    SetZoomDisabled(value *bool)()
}
