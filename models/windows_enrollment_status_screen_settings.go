package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// WindowsEnrollmentStatusScreenSettings enrollment status screen setting
type WindowsEnrollmentStatusScreenSettings struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewWindowsEnrollmentStatusScreenSettings instantiates a new WindowsEnrollmentStatusScreenSettings and sets the default values.
func NewWindowsEnrollmentStatusScreenSettings()(*WindowsEnrollmentStatusScreenSettings) {
    m := &WindowsEnrollmentStatusScreenSettings{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateWindowsEnrollmentStatusScreenSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateWindowsEnrollmentStatusScreenSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsEnrollmentStatusScreenSettings(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *WindowsEnrollmentStatusScreenSettings) GetAdditionalData()(map[string]any) {
    val , err :=  m.backingStore.Get("additionalData")
    if err != nil {
        panic(err)
    }
    if val == nil {
        var value = make(map[string]any);
        m.SetAdditionalData(value);
    }
    return val.(map[string]any)
}
// GetAllowDeviceUseBeforeProfileAndAppInstallComplete gets the allowDeviceUseBeforeProfileAndAppInstallComplete property value. Allow or block user to use device before profile and app installation complete
// returns a *bool when successful
func (m *WindowsEnrollmentStatusScreenSettings) GetAllowDeviceUseBeforeProfileAndAppInstallComplete()(*bool) {
    val, err := m.GetBackingStore().Get("allowDeviceUseBeforeProfileAndAppInstallComplete")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetAllowDeviceUseOnInstallFailure gets the allowDeviceUseOnInstallFailure property value. Allow the user to continue using the device on installation failure
// returns a *bool when successful
func (m *WindowsEnrollmentStatusScreenSettings) GetAllowDeviceUseOnInstallFailure()(*bool) {
    val, err := m.GetBackingStore().Get("allowDeviceUseOnInstallFailure")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetAllowLogCollectionOnInstallFailure gets the allowLogCollectionOnInstallFailure property value. Allow or block log collection on installation failure
// returns a *bool when successful
func (m *WindowsEnrollmentStatusScreenSettings) GetAllowLogCollectionOnInstallFailure()(*bool) {
    val, err := m.GetBackingStore().Get("allowLogCollectionOnInstallFailure")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetBackingStore gets the BackingStore property value. Stores model information.
// returns a BackingStore when successful
func (m *WindowsEnrollmentStatusScreenSettings) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetBlockDeviceSetupRetryByUser gets the blockDeviceSetupRetryByUser property value. Allow the user to retry the setup on installation failure
// returns a *bool when successful
func (m *WindowsEnrollmentStatusScreenSettings) GetBlockDeviceSetupRetryByUser()(*bool) {
    val, err := m.GetBackingStore().Get("blockDeviceSetupRetryByUser")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetCustomErrorMessage gets the customErrorMessage property value. Set custom error message to show upon installation failure
// returns a *string when successful
func (m *WindowsEnrollmentStatusScreenSettings) GetCustomErrorMessage()(*string) {
    val, err := m.GetBackingStore().Get("customErrorMessage")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *WindowsEnrollmentStatusScreenSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["allowDeviceUseBeforeProfileAndAppInstallComplete"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowDeviceUseBeforeProfileAndAppInstallComplete(val)
        }
        return nil
    }
    res["allowDeviceUseOnInstallFailure"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowDeviceUseOnInstallFailure(val)
        }
        return nil
    }
    res["allowLogCollectionOnInstallFailure"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowLogCollectionOnInstallFailure(val)
        }
        return nil
    }
    res["blockDeviceSetupRetryByUser"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBlockDeviceSetupRetryByUser(val)
        }
        return nil
    }
    res["customErrorMessage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomErrorMessage(val)
        }
        return nil
    }
    res["hideInstallationProgress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHideInstallationProgress(val)
        }
        return nil
    }
    res["installProgressTimeoutInMinutes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInstallProgressTimeoutInMinutes(val)
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
    return res
}
// GetHideInstallationProgress gets the hideInstallationProgress property value. Show or hide installation progress to user
// returns a *bool when successful
func (m *WindowsEnrollmentStatusScreenSettings) GetHideInstallationProgress()(*bool) {
    val, err := m.GetBackingStore().Get("hideInstallationProgress")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetInstallProgressTimeoutInMinutes gets the installProgressTimeoutInMinutes property value. Set installation progress timeout in minutes
// returns a *int32 when successful
func (m *WindowsEnrollmentStatusScreenSettings) GetInstallProgressTimeoutInMinutes()(*int32) {
    val, err := m.GetBackingStore().Get("installProgressTimeoutInMinutes")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *WindowsEnrollmentStatusScreenSettings) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *WindowsEnrollmentStatusScreenSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("allowDeviceUseBeforeProfileAndAppInstallComplete", m.GetAllowDeviceUseBeforeProfileAndAppInstallComplete())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("allowDeviceUseOnInstallFailure", m.GetAllowDeviceUseOnInstallFailure())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("allowLogCollectionOnInstallFailure", m.GetAllowLogCollectionOnInstallFailure())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("blockDeviceSetupRetryByUser", m.GetBlockDeviceSetupRetryByUser())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("customErrorMessage", m.GetCustomErrorMessage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("hideInstallationProgress", m.GetHideInstallationProgress())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("installProgressTimeoutInMinutes", m.GetInstallProgressTimeoutInMinutes())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WindowsEnrollmentStatusScreenSettings) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetAllowDeviceUseBeforeProfileAndAppInstallComplete sets the allowDeviceUseBeforeProfileAndAppInstallComplete property value. Allow or block user to use device before profile and app installation complete
func (m *WindowsEnrollmentStatusScreenSettings) SetAllowDeviceUseBeforeProfileAndAppInstallComplete(value *bool)() {
    err := m.GetBackingStore().Set("allowDeviceUseBeforeProfileAndAppInstallComplete", value)
    if err != nil {
        panic(err)
    }
}
// SetAllowDeviceUseOnInstallFailure sets the allowDeviceUseOnInstallFailure property value. Allow the user to continue using the device on installation failure
func (m *WindowsEnrollmentStatusScreenSettings) SetAllowDeviceUseOnInstallFailure(value *bool)() {
    err := m.GetBackingStore().Set("allowDeviceUseOnInstallFailure", value)
    if err != nil {
        panic(err)
    }
}
// SetAllowLogCollectionOnInstallFailure sets the allowLogCollectionOnInstallFailure property value. Allow or block log collection on installation failure
func (m *WindowsEnrollmentStatusScreenSettings) SetAllowLogCollectionOnInstallFailure(value *bool)() {
    err := m.GetBackingStore().Set("allowLogCollectionOnInstallFailure", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *WindowsEnrollmentStatusScreenSettings) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetBlockDeviceSetupRetryByUser sets the blockDeviceSetupRetryByUser property value. Allow the user to retry the setup on installation failure
func (m *WindowsEnrollmentStatusScreenSettings) SetBlockDeviceSetupRetryByUser(value *bool)() {
    err := m.GetBackingStore().Set("blockDeviceSetupRetryByUser", value)
    if err != nil {
        panic(err)
    }
}
// SetCustomErrorMessage sets the customErrorMessage property value. Set custom error message to show upon installation failure
func (m *WindowsEnrollmentStatusScreenSettings) SetCustomErrorMessage(value *string)() {
    err := m.GetBackingStore().Set("customErrorMessage", value)
    if err != nil {
        panic(err)
    }
}
// SetHideInstallationProgress sets the hideInstallationProgress property value. Show or hide installation progress to user
func (m *WindowsEnrollmentStatusScreenSettings) SetHideInstallationProgress(value *bool)() {
    err := m.GetBackingStore().Set("hideInstallationProgress", value)
    if err != nil {
        panic(err)
    }
}
// SetInstallProgressTimeoutInMinutes sets the installProgressTimeoutInMinutes property value. Set installation progress timeout in minutes
func (m *WindowsEnrollmentStatusScreenSettings) SetInstallProgressTimeoutInMinutes(value *int32)() {
    err := m.GetBackingStore().Set("installProgressTimeoutInMinutes", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *WindowsEnrollmentStatusScreenSettings) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
type WindowsEnrollmentStatusScreenSettingsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAllowDeviceUseBeforeProfileAndAppInstallComplete()(*bool)
    GetAllowDeviceUseOnInstallFailure()(*bool)
    GetAllowLogCollectionOnInstallFailure()(*bool)
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetBlockDeviceSetupRetryByUser()(*bool)
    GetCustomErrorMessage()(*string)
    GetHideInstallationProgress()(*bool)
    GetInstallProgressTimeoutInMinutes()(*int32)
    GetOdataType()(*string)
    SetAllowDeviceUseBeforeProfileAndAppInstallComplete(value *bool)()
    SetAllowDeviceUseOnInstallFailure(value *bool)()
    SetAllowLogCollectionOnInstallFailure(value *bool)()
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetBlockDeviceSetupRetryByUser(value *bool)()
    SetCustomErrorMessage(value *string)()
    SetHideInstallationProgress(value *bool)()
    SetInstallProgressTimeoutInMinutes(value *int32)()
    SetOdataType(value *string)()
}
