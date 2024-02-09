package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

type OnPremisesDirectorySynchronizationConfiguration struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewOnPremisesDirectorySynchronizationConfiguration instantiates a new OnPremisesDirectorySynchronizationConfiguration and sets the default values.
func NewOnPremisesDirectorySynchronizationConfiguration()(*OnPremisesDirectorySynchronizationConfiguration) {
    m := &OnPremisesDirectorySynchronizationConfiguration{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateOnPremisesDirectorySynchronizationConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateOnPremisesDirectorySynchronizationConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOnPremisesDirectorySynchronizationConfiguration(), nil
}
// GetAccidentalDeletionPrevention gets the accidentalDeletionPrevention property value. Contains the accidental deletion prevention configuration for a tenant.
// returns a OnPremisesAccidentalDeletionPreventionable when successful
func (m *OnPremisesDirectorySynchronizationConfiguration) GetAccidentalDeletionPrevention()(OnPremisesAccidentalDeletionPreventionable) {
    val, err := m.GetBackingStore().Get("accidentalDeletionPrevention")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(OnPremisesAccidentalDeletionPreventionable)
    }
    return nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *OnPremisesDirectorySynchronizationConfiguration) GetAdditionalData()(map[string]any) {
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
// GetAnchorAttribute gets the anchorAttribute property value. The anchor attribute allows customers to customize the property used to create source anchors for synchronization enabled objects.
// returns a *string when successful
func (m *OnPremisesDirectorySynchronizationConfiguration) GetAnchorAttribute()(*string) {
    val, err := m.GetBackingStore().Get("anchorAttribute")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetApplicationId gets the applicationId property value. The identifier of the on-premises directory synchronization client application that is configured for the tenant.
// returns a *string when successful
func (m *OnPremisesDirectorySynchronizationConfiguration) GetApplicationId()(*string) {
    val, err := m.GetBackingStore().Get("applicationId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetBackingStore gets the BackingStore property value. Stores model information.
// returns a BackingStore when successful
func (m *OnPremisesDirectorySynchronizationConfiguration) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetCurrentExportData gets the currentExportData property value. Data for the current export run.
// returns a OnPremisesCurrentExportDataable when successful
func (m *OnPremisesDirectorySynchronizationConfiguration) GetCurrentExportData()(OnPremisesCurrentExportDataable) {
    val, err := m.GetBackingStore().Get("currentExportData")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(OnPremisesCurrentExportDataable)
    }
    return nil
}
// GetCustomerRequestedSynchronizationInterval gets the customerRequestedSynchronizationInterval property value. Interval of time that the customer requested the sync client waits between sync cycles.
// returns a *ISODuration when successful
func (m *OnPremisesDirectorySynchronizationConfiguration) GetCustomerRequestedSynchronizationInterval()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration) {
    val, err := m.GetBackingStore().Get("customerRequestedSynchronizationInterval")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *OnPremisesDirectorySynchronizationConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["accidentalDeletionPrevention"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateOnPremisesAccidentalDeletionPreventionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccidentalDeletionPrevention(val.(OnPremisesAccidentalDeletionPreventionable))
        }
        return nil
    }
    res["anchorAttribute"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAnchorAttribute(val)
        }
        return nil
    }
    res["applicationId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplicationId(val)
        }
        return nil
    }
    res["currentExportData"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateOnPremisesCurrentExportDataFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCurrentExportData(val.(OnPremisesCurrentExportDataable))
        }
        return nil
    }
    res["customerRequestedSynchronizationInterval"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetISODurationValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomerRequestedSynchronizationInterval(val)
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
    res["synchronizationClientVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSynchronizationClientVersion(val)
        }
        return nil
    }
    res["synchronizationInterval"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetISODurationValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSynchronizationInterval(val)
        }
        return nil
    }
    res["writebackConfiguration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateOnPremisesWritebackConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWritebackConfiguration(val.(OnPremisesWritebackConfigurationable))
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *OnPremisesDirectorySynchronizationConfiguration) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSynchronizationClientVersion gets the synchronizationClientVersion property value. Indicates the version of the on-premises directory synchronization application.
// returns a *string when successful
func (m *OnPremisesDirectorySynchronizationConfiguration) GetSynchronizationClientVersion()(*string) {
    val, err := m.GetBackingStore().Get("synchronizationClientVersion")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSynchronizationInterval gets the synchronizationInterval property value. Interval of time the sync client should honor between sync cycles
// returns a *ISODuration when successful
func (m *OnPremisesDirectorySynchronizationConfiguration) GetSynchronizationInterval()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration) {
    val, err := m.GetBackingStore().Get("synchronizationInterval")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)
    }
    return nil
}
// GetWritebackConfiguration gets the writebackConfiguration property value. Configuration to control how cloud created or owned objects are synchronized back to the on-premises directory.
// returns a OnPremisesWritebackConfigurationable when successful
func (m *OnPremisesDirectorySynchronizationConfiguration) GetWritebackConfiguration()(OnPremisesWritebackConfigurationable) {
    val, err := m.GetBackingStore().Get("writebackConfiguration")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(OnPremisesWritebackConfigurationable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *OnPremisesDirectorySynchronizationConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("accidentalDeletionPrevention", m.GetAccidentalDeletionPrevention())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("anchorAttribute", m.GetAnchorAttribute())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("applicationId", m.GetApplicationId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("currentExportData", m.GetCurrentExportData())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteISODurationValue("customerRequestedSynchronizationInterval", m.GetCustomerRequestedSynchronizationInterval())
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
        err := writer.WriteStringValue("synchronizationClientVersion", m.GetSynchronizationClientVersion())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteISODurationValue("synchronizationInterval", m.GetSynchronizationInterval())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("writebackConfiguration", m.GetWritebackConfiguration())
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
// SetAccidentalDeletionPrevention sets the accidentalDeletionPrevention property value. Contains the accidental deletion prevention configuration for a tenant.
func (m *OnPremisesDirectorySynchronizationConfiguration) SetAccidentalDeletionPrevention(value OnPremisesAccidentalDeletionPreventionable)() {
    err := m.GetBackingStore().Set("accidentalDeletionPrevention", value)
    if err != nil {
        panic(err)
    }
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *OnPremisesDirectorySynchronizationConfiguration) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetAnchorAttribute sets the anchorAttribute property value. The anchor attribute allows customers to customize the property used to create source anchors for synchronization enabled objects.
func (m *OnPremisesDirectorySynchronizationConfiguration) SetAnchorAttribute(value *string)() {
    err := m.GetBackingStore().Set("anchorAttribute", value)
    if err != nil {
        panic(err)
    }
}
// SetApplicationId sets the applicationId property value. The identifier of the on-premises directory synchronization client application that is configured for the tenant.
func (m *OnPremisesDirectorySynchronizationConfiguration) SetApplicationId(value *string)() {
    err := m.GetBackingStore().Set("applicationId", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *OnPremisesDirectorySynchronizationConfiguration) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetCurrentExportData sets the currentExportData property value. Data for the current export run.
func (m *OnPremisesDirectorySynchronizationConfiguration) SetCurrentExportData(value OnPremisesCurrentExportDataable)() {
    err := m.GetBackingStore().Set("currentExportData", value)
    if err != nil {
        panic(err)
    }
}
// SetCustomerRequestedSynchronizationInterval sets the customerRequestedSynchronizationInterval property value. Interval of time that the customer requested the sync client waits between sync cycles.
func (m *OnPremisesDirectorySynchronizationConfiguration) SetCustomerRequestedSynchronizationInterval(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)() {
    err := m.GetBackingStore().Set("customerRequestedSynchronizationInterval", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *OnPremisesDirectorySynchronizationConfiguration) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetSynchronizationClientVersion sets the synchronizationClientVersion property value. Indicates the version of the on-premises directory synchronization application.
func (m *OnPremisesDirectorySynchronizationConfiguration) SetSynchronizationClientVersion(value *string)() {
    err := m.GetBackingStore().Set("synchronizationClientVersion", value)
    if err != nil {
        panic(err)
    }
}
// SetSynchronizationInterval sets the synchronizationInterval property value. Interval of time the sync client should honor between sync cycles
func (m *OnPremisesDirectorySynchronizationConfiguration) SetSynchronizationInterval(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)() {
    err := m.GetBackingStore().Set("synchronizationInterval", value)
    if err != nil {
        panic(err)
    }
}
// SetWritebackConfiguration sets the writebackConfiguration property value. Configuration to control how cloud created or owned objects are synchronized back to the on-premises directory.
func (m *OnPremisesDirectorySynchronizationConfiguration) SetWritebackConfiguration(value OnPremisesWritebackConfigurationable)() {
    err := m.GetBackingStore().Set("writebackConfiguration", value)
    if err != nil {
        panic(err)
    }
}
type OnPremisesDirectorySynchronizationConfigurationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccidentalDeletionPrevention()(OnPremisesAccidentalDeletionPreventionable)
    GetAnchorAttribute()(*string)
    GetApplicationId()(*string)
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetCurrentExportData()(OnPremisesCurrentExportDataable)
    GetCustomerRequestedSynchronizationInterval()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)
    GetOdataType()(*string)
    GetSynchronizationClientVersion()(*string)
    GetSynchronizationInterval()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)
    GetWritebackConfiguration()(OnPremisesWritebackConfigurationable)
    SetAccidentalDeletionPrevention(value OnPremisesAccidentalDeletionPreventionable)()
    SetAnchorAttribute(value *string)()
    SetApplicationId(value *string)()
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetCurrentExportData(value OnPremisesCurrentExportDataable)()
    SetCustomerRequestedSynchronizationInterval(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)()
    SetOdataType(value *string)()
    SetSynchronizationClientVersion(value *string)()
    SetSynchronizationInterval(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)()
    SetWritebackConfiguration(value OnPremisesWritebackConfigurationable)()
}
