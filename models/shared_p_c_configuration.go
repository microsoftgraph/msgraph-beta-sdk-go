package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SharedPCConfiguration this topic provides descriptions of the declared methods, properties and relationships exposed by the sharedPCConfiguration resource.
type SharedPCConfiguration struct {
    DeviceConfiguration
}
// NewSharedPCConfiguration instantiates a new SharedPCConfiguration and sets the default values.
func NewSharedPCConfiguration()(*SharedPCConfiguration) {
    m := &SharedPCConfiguration{
        DeviceConfiguration: *NewDeviceConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.sharedPCConfiguration"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateSharedPCConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSharedPCConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSharedPCConfiguration(), nil
}
// GetAccountManagerPolicy gets the accountManagerPolicy property value. Specifies how accounts are managed on a shared PC. Only applies when disableAccountManager is false.
// returns a SharedPCAccountManagerPolicyable when successful
func (m *SharedPCConfiguration) GetAccountManagerPolicy()(SharedPCAccountManagerPolicyable) {
    val, err := m.GetBackingStore().Get("accountManagerPolicy")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(SharedPCAccountManagerPolicyable)
    }
    return nil
}
// GetAllowedAccounts gets the allowedAccounts property value. Type of accounts that are allowed to share the PC.
// returns a *SharedPCAllowedAccountType when successful
func (m *SharedPCConfiguration) GetAllowedAccounts()(*SharedPCAllowedAccountType) {
    val, err := m.GetBackingStore().Get("allowedAccounts")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*SharedPCAllowedAccountType)
    }
    return nil
}
// GetAllowLocalStorage gets the allowLocalStorage property value. Specifies whether local storage is allowed on a shared PC.
// returns a *bool when successful
func (m *SharedPCConfiguration) GetAllowLocalStorage()(*bool) {
    val, err := m.GetBackingStore().Get("allowLocalStorage")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetDisableAccountManager gets the disableAccountManager property value. Disables the account manager for shared PC mode.
// returns a *bool when successful
func (m *SharedPCConfiguration) GetDisableAccountManager()(*bool) {
    val, err := m.GetBackingStore().Get("disableAccountManager")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetDisableEduPolicies gets the disableEduPolicies property value. Specifies whether the default shared PC education environment policies should be disabled. For Windows 10 RS2 and later, this policy will be applied without setting Enabled to true.
// returns a *bool when successful
func (m *SharedPCConfiguration) GetDisableEduPolicies()(*bool) {
    val, err := m.GetBackingStore().Get("disableEduPolicies")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetDisablePowerPolicies gets the disablePowerPolicies property value. Specifies whether the default shared PC power policies should be disabled.
// returns a *bool when successful
func (m *SharedPCConfiguration) GetDisablePowerPolicies()(*bool) {
    val, err := m.GetBackingStore().Get("disablePowerPolicies")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetDisableSignInOnResume gets the disableSignInOnResume property value. Disables the requirement to sign in whenever the device wakes up from sleep mode.
// returns a *bool when successful
func (m *SharedPCConfiguration) GetDisableSignInOnResume()(*bool) {
    val, err := m.GetBackingStore().Get("disableSignInOnResume")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetEnabled gets the enabled property value. Enables shared PC mode and applies the shared pc policies.
// returns a *bool when successful
func (m *SharedPCConfiguration) GetEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("enabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetFastFirstSignIn gets the fastFirstSignIn property value. Possible values of a property
// returns a *Enablement when successful
func (m *SharedPCConfiguration) GetFastFirstSignIn()(*Enablement) {
    val, err := m.GetBackingStore().Get("fastFirstSignIn")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*Enablement)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *SharedPCConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceConfiguration.GetFieldDeserializers()
    res["accountManagerPolicy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSharedPCAccountManagerPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccountManagerPolicy(val.(SharedPCAccountManagerPolicyable))
        }
        return nil
    }
    res["allowedAccounts"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseSharedPCAllowedAccountType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowedAccounts(val.(*SharedPCAllowedAccountType))
        }
        return nil
    }
    res["allowLocalStorage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowLocalStorage(val)
        }
        return nil
    }
    res["disableAccountManager"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisableAccountManager(val)
        }
        return nil
    }
    res["disableEduPolicies"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisableEduPolicies(val)
        }
        return nil
    }
    res["disablePowerPolicies"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisablePowerPolicies(val)
        }
        return nil
    }
    res["disableSignInOnResume"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisableSignInOnResume(val)
        }
        return nil
    }
    res["enabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnabled(val)
        }
        return nil
    }
    res["fastFirstSignIn"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEnablement)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFastFirstSignIn(val.(*Enablement))
        }
        return nil
    }
    res["idleTimeBeforeSleepInSeconds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIdleTimeBeforeSleepInSeconds(val)
        }
        return nil
    }
    res["kioskAppDisplayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKioskAppDisplayName(val)
        }
        return nil
    }
    res["kioskAppUserModelId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKioskAppUserModelId(val)
        }
        return nil
    }
    res["localStorage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEnablement)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocalStorage(val.(*Enablement))
        }
        return nil
    }
    res["maintenanceStartTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaintenanceStartTime(val)
        }
        return nil
    }
    res["setAccountManager"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEnablement)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSetAccountManager(val.(*Enablement))
        }
        return nil
    }
    res["setEduPolicies"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEnablement)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSetEduPolicies(val.(*Enablement))
        }
        return nil
    }
    res["setPowerPolicies"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEnablement)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSetPowerPolicies(val.(*Enablement))
        }
        return nil
    }
    res["signInOnResume"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEnablement)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSignInOnResume(val.(*Enablement))
        }
        return nil
    }
    return res
}
// GetIdleTimeBeforeSleepInSeconds gets the idleTimeBeforeSleepInSeconds property value. Specifies the time in seconds that a device must sit idle before the PC goes to sleep. Setting this value to 0 prevents the sleep timeout from occurring.
// returns a *int32 when successful
func (m *SharedPCConfiguration) GetIdleTimeBeforeSleepInSeconds()(*int32) {
    val, err := m.GetBackingStore().Get("idleTimeBeforeSleepInSeconds")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetKioskAppDisplayName gets the kioskAppDisplayName property value. Specifies the display text for the account shown on the sign-in screen which launches the app specified by SetKioskAppUserModelId. Only applies when KioskAppUserModelId is set.
// returns a *string when successful
func (m *SharedPCConfiguration) GetKioskAppDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("kioskAppDisplayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetKioskAppUserModelId gets the kioskAppUserModelId property value. Specifies the application user model ID of the app to use with assigned access.
// returns a *string when successful
func (m *SharedPCConfiguration) GetKioskAppUserModelId()(*string) {
    val, err := m.GetBackingStore().Get("kioskAppUserModelId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetLocalStorage gets the localStorage property value. Possible values of a property
// returns a *Enablement when successful
func (m *SharedPCConfiguration) GetLocalStorage()(*Enablement) {
    val, err := m.GetBackingStore().Get("localStorage")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*Enablement)
    }
    return nil
}
// GetMaintenanceStartTime gets the maintenanceStartTime property value. Specifies the daily start time of maintenance hour.
// returns a *TimeOnly when successful
func (m *SharedPCConfiguration) GetMaintenanceStartTime()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly) {
    val, err := m.GetBackingStore().Get("maintenanceStartTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly)
    }
    return nil
}
// GetSetAccountManager gets the setAccountManager property value. Possible values of a property
// returns a *Enablement when successful
func (m *SharedPCConfiguration) GetSetAccountManager()(*Enablement) {
    val, err := m.GetBackingStore().Get("setAccountManager")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*Enablement)
    }
    return nil
}
// GetSetEduPolicies gets the setEduPolicies property value. Possible values of a property
// returns a *Enablement when successful
func (m *SharedPCConfiguration) GetSetEduPolicies()(*Enablement) {
    val, err := m.GetBackingStore().Get("setEduPolicies")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*Enablement)
    }
    return nil
}
// GetSetPowerPolicies gets the setPowerPolicies property value. Possible values of a property
// returns a *Enablement when successful
func (m *SharedPCConfiguration) GetSetPowerPolicies()(*Enablement) {
    val, err := m.GetBackingStore().Get("setPowerPolicies")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*Enablement)
    }
    return nil
}
// GetSignInOnResume gets the signInOnResume property value. Possible values of a property
// returns a *Enablement when successful
func (m *SharedPCConfiguration) GetSignInOnResume()(*Enablement) {
    val, err := m.GetBackingStore().Get("signInOnResume")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*Enablement)
    }
    return nil
}
// Serialize serializes information the current object
func (m *SharedPCConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceConfiguration.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("accountManagerPolicy", m.GetAccountManagerPolicy())
        if err != nil {
            return err
        }
    }
    if m.GetAllowedAccounts() != nil {
        cast := (*m.GetAllowedAccounts()).String()
        err = writer.WriteStringValue("allowedAccounts", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("allowLocalStorage", m.GetAllowLocalStorage())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("disableAccountManager", m.GetDisableAccountManager())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("disableEduPolicies", m.GetDisableEduPolicies())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("disablePowerPolicies", m.GetDisablePowerPolicies())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("disableSignInOnResume", m.GetDisableSignInOnResume())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("enabled", m.GetEnabled())
        if err != nil {
            return err
        }
    }
    if m.GetFastFirstSignIn() != nil {
        cast := (*m.GetFastFirstSignIn()).String()
        err = writer.WriteStringValue("fastFirstSignIn", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("idleTimeBeforeSleepInSeconds", m.GetIdleTimeBeforeSleepInSeconds())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("kioskAppDisplayName", m.GetKioskAppDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("kioskAppUserModelId", m.GetKioskAppUserModelId())
        if err != nil {
            return err
        }
    }
    if m.GetLocalStorage() != nil {
        cast := (*m.GetLocalStorage()).String()
        err = writer.WriteStringValue("localStorage", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeOnlyValue("maintenanceStartTime", m.GetMaintenanceStartTime())
        if err != nil {
            return err
        }
    }
    if m.GetSetAccountManager() != nil {
        cast := (*m.GetSetAccountManager()).String()
        err = writer.WriteStringValue("setAccountManager", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetSetEduPolicies() != nil {
        cast := (*m.GetSetEduPolicies()).String()
        err = writer.WriteStringValue("setEduPolicies", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetSetPowerPolicies() != nil {
        cast := (*m.GetSetPowerPolicies()).String()
        err = writer.WriteStringValue("setPowerPolicies", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetSignInOnResume() != nil {
        cast := (*m.GetSignInOnResume()).String()
        err = writer.WriteStringValue("signInOnResume", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAccountManagerPolicy sets the accountManagerPolicy property value. Specifies how accounts are managed on a shared PC. Only applies when disableAccountManager is false.
func (m *SharedPCConfiguration) SetAccountManagerPolicy(value SharedPCAccountManagerPolicyable)() {
    err := m.GetBackingStore().Set("accountManagerPolicy", value)
    if err != nil {
        panic(err)
    }
}
// SetAllowedAccounts sets the allowedAccounts property value. Type of accounts that are allowed to share the PC.
func (m *SharedPCConfiguration) SetAllowedAccounts(value *SharedPCAllowedAccountType)() {
    err := m.GetBackingStore().Set("allowedAccounts", value)
    if err != nil {
        panic(err)
    }
}
// SetAllowLocalStorage sets the allowLocalStorage property value. Specifies whether local storage is allowed on a shared PC.
func (m *SharedPCConfiguration) SetAllowLocalStorage(value *bool)() {
    err := m.GetBackingStore().Set("allowLocalStorage", value)
    if err != nil {
        panic(err)
    }
}
// SetDisableAccountManager sets the disableAccountManager property value. Disables the account manager for shared PC mode.
func (m *SharedPCConfiguration) SetDisableAccountManager(value *bool)() {
    err := m.GetBackingStore().Set("disableAccountManager", value)
    if err != nil {
        panic(err)
    }
}
// SetDisableEduPolicies sets the disableEduPolicies property value. Specifies whether the default shared PC education environment policies should be disabled. For Windows 10 RS2 and later, this policy will be applied without setting Enabled to true.
func (m *SharedPCConfiguration) SetDisableEduPolicies(value *bool)() {
    err := m.GetBackingStore().Set("disableEduPolicies", value)
    if err != nil {
        panic(err)
    }
}
// SetDisablePowerPolicies sets the disablePowerPolicies property value. Specifies whether the default shared PC power policies should be disabled.
func (m *SharedPCConfiguration) SetDisablePowerPolicies(value *bool)() {
    err := m.GetBackingStore().Set("disablePowerPolicies", value)
    if err != nil {
        panic(err)
    }
}
// SetDisableSignInOnResume sets the disableSignInOnResume property value. Disables the requirement to sign in whenever the device wakes up from sleep mode.
func (m *SharedPCConfiguration) SetDisableSignInOnResume(value *bool)() {
    err := m.GetBackingStore().Set("disableSignInOnResume", value)
    if err != nil {
        panic(err)
    }
}
// SetEnabled sets the enabled property value. Enables shared PC mode and applies the shared pc policies.
func (m *SharedPCConfiguration) SetEnabled(value *bool)() {
    err := m.GetBackingStore().Set("enabled", value)
    if err != nil {
        panic(err)
    }
}
// SetFastFirstSignIn sets the fastFirstSignIn property value. Possible values of a property
func (m *SharedPCConfiguration) SetFastFirstSignIn(value *Enablement)() {
    err := m.GetBackingStore().Set("fastFirstSignIn", value)
    if err != nil {
        panic(err)
    }
}
// SetIdleTimeBeforeSleepInSeconds sets the idleTimeBeforeSleepInSeconds property value. Specifies the time in seconds that a device must sit idle before the PC goes to sleep. Setting this value to 0 prevents the sleep timeout from occurring.
func (m *SharedPCConfiguration) SetIdleTimeBeforeSleepInSeconds(value *int32)() {
    err := m.GetBackingStore().Set("idleTimeBeforeSleepInSeconds", value)
    if err != nil {
        panic(err)
    }
}
// SetKioskAppDisplayName sets the kioskAppDisplayName property value. Specifies the display text for the account shown on the sign-in screen which launches the app specified by SetKioskAppUserModelId. Only applies when KioskAppUserModelId is set.
func (m *SharedPCConfiguration) SetKioskAppDisplayName(value *string)() {
    err := m.GetBackingStore().Set("kioskAppDisplayName", value)
    if err != nil {
        panic(err)
    }
}
// SetKioskAppUserModelId sets the kioskAppUserModelId property value. Specifies the application user model ID of the app to use with assigned access.
func (m *SharedPCConfiguration) SetKioskAppUserModelId(value *string)() {
    err := m.GetBackingStore().Set("kioskAppUserModelId", value)
    if err != nil {
        panic(err)
    }
}
// SetLocalStorage sets the localStorage property value. Possible values of a property
func (m *SharedPCConfiguration) SetLocalStorage(value *Enablement)() {
    err := m.GetBackingStore().Set("localStorage", value)
    if err != nil {
        panic(err)
    }
}
// SetMaintenanceStartTime sets the maintenanceStartTime property value. Specifies the daily start time of maintenance hour.
func (m *SharedPCConfiguration) SetMaintenanceStartTime(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly)() {
    err := m.GetBackingStore().Set("maintenanceStartTime", value)
    if err != nil {
        panic(err)
    }
}
// SetSetAccountManager sets the setAccountManager property value. Possible values of a property
func (m *SharedPCConfiguration) SetSetAccountManager(value *Enablement)() {
    err := m.GetBackingStore().Set("setAccountManager", value)
    if err != nil {
        panic(err)
    }
}
// SetSetEduPolicies sets the setEduPolicies property value. Possible values of a property
func (m *SharedPCConfiguration) SetSetEduPolicies(value *Enablement)() {
    err := m.GetBackingStore().Set("setEduPolicies", value)
    if err != nil {
        panic(err)
    }
}
// SetSetPowerPolicies sets the setPowerPolicies property value. Possible values of a property
func (m *SharedPCConfiguration) SetSetPowerPolicies(value *Enablement)() {
    err := m.GetBackingStore().Set("setPowerPolicies", value)
    if err != nil {
        panic(err)
    }
}
// SetSignInOnResume sets the signInOnResume property value. Possible values of a property
func (m *SharedPCConfiguration) SetSignInOnResume(value *Enablement)() {
    err := m.GetBackingStore().Set("signInOnResume", value)
    if err != nil {
        panic(err)
    }
}
type SharedPCConfigurationable interface {
    DeviceConfigurationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccountManagerPolicy()(SharedPCAccountManagerPolicyable)
    GetAllowedAccounts()(*SharedPCAllowedAccountType)
    GetAllowLocalStorage()(*bool)
    GetDisableAccountManager()(*bool)
    GetDisableEduPolicies()(*bool)
    GetDisablePowerPolicies()(*bool)
    GetDisableSignInOnResume()(*bool)
    GetEnabled()(*bool)
    GetFastFirstSignIn()(*Enablement)
    GetIdleTimeBeforeSleepInSeconds()(*int32)
    GetKioskAppDisplayName()(*string)
    GetKioskAppUserModelId()(*string)
    GetLocalStorage()(*Enablement)
    GetMaintenanceStartTime()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly)
    GetSetAccountManager()(*Enablement)
    GetSetEduPolicies()(*Enablement)
    GetSetPowerPolicies()(*Enablement)
    GetSignInOnResume()(*Enablement)
    SetAccountManagerPolicy(value SharedPCAccountManagerPolicyable)()
    SetAllowedAccounts(value *SharedPCAllowedAccountType)()
    SetAllowLocalStorage(value *bool)()
    SetDisableAccountManager(value *bool)()
    SetDisableEduPolicies(value *bool)()
    SetDisablePowerPolicies(value *bool)()
    SetDisableSignInOnResume(value *bool)()
    SetEnabled(value *bool)()
    SetFastFirstSignIn(value *Enablement)()
    SetIdleTimeBeforeSleepInSeconds(value *int32)()
    SetKioskAppDisplayName(value *string)()
    SetKioskAppUserModelId(value *string)()
    SetLocalStorage(value *Enablement)()
    SetMaintenanceStartTime(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly)()
    SetSetAccountManager(value *Enablement)()
    SetSetEduPolicies(value *Enablement)()
    SetSetPowerPolicies(value *Enablement)()
    SetSignInOnResume(value *Enablement)()
}
