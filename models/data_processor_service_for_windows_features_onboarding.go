package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// DataProcessorServiceForWindowsFeaturesOnboarding a configuration entity for MEM features that utilize Data Processor Service for Windows (DPSW) data.
type DataProcessorServiceForWindowsFeaturesOnboarding struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewDataProcessorServiceForWindowsFeaturesOnboarding instantiates a new dataProcessorServiceForWindowsFeaturesOnboarding and sets the default values.
func NewDataProcessorServiceForWindowsFeaturesOnboarding()(*DataProcessorServiceForWindowsFeaturesOnboarding) {
    m := &DataProcessorServiceForWindowsFeaturesOnboarding{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateDataProcessorServiceForWindowsFeaturesOnboardingFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDataProcessorServiceForWindowsFeaturesOnboardingFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDataProcessorServiceForWindowsFeaturesOnboarding(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DataProcessorServiceForWindowsFeaturesOnboarding) GetAdditionalData()(map[string]any) {
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
// GetAreDataProcessorServiceForWindowsFeaturesEnabled gets the areDataProcessorServiceForWindowsFeaturesEnabled property value. Indicates whether the tenant has enabled MEM features utilizing Data Processor Service for Windows (DPSW) data. When TRUE, the tenant has enabled MEM features utilizing Data Processor Service for Windows (DPSW) data. When FALSE, the tenant has not enabled MEM features utilizing Data Processor Service for Windows (DPSW) data. Default value is FALSE.
func (m *DataProcessorServiceForWindowsFeaturesOnboarding) GetAreDataProcessorServiceForWindowsFeaturesEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("areDataProcessorServiceForWindowsFeaturesEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetBackingStore gets the backingStore property value. Stores model information.
func (m *DataProcessorServiceForWindowsFeaturesOnboarding) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DataProcessorServiceForWindowsFeaturesOnboarding) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["areDataProcessorServiceForWindowsFeaturesEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAreDataProcessorServiceForWindowsFeaturesEnabled(val)
        }
        return nil
    }
    res["hasValidWindowsLicense"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHasValidWindowsLicense(val)
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
// GetHasValidWindowsLicense gets the hasValidWindowsLicense property value. Indicates whether the tenant has required Windows license. When TRUE, the tenant has the required Windows license. When FALSE, the tenant does not have the required Windows license. Default value is FALSE.
func (m *DataProcessorServiceForWindowsFeaturesOnboarding) GetHasValidWindowsLicense()(*bool) {
    val, err := m.GetBackingStore().Get("hasValidWindowsLicense")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *DataProcessorServiceForWindowsFeaturesOnboarding) GetOdataType()(*string) {
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
func (m *DataProcessorServiceForWindowsFeaturesOnboarding) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("areDataProcessorServiceForWindowsFeaturesEnabled", m.GetAreDataProcessorServiceForWindowsFeaturesEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("hasValidWindowsLicense", m.GetHasValidWindowsLicense())
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DataProcessorServiceForWindowsFeaturesOnboarding) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetAreDataProcessorServiceForWindowsFeaturesEnabled sets the areDataProcessorServiceForWindowsFeaturesEnabled property value. Indicates whether the tenant has enabled MEM features utilizing Data Processor Service for Windows (DPSW) data. When TRUE, the tenant has enabled MEM features utilizing Data Processor Service for Windows (DPSW) data. When FALSE, the tenant has not enabled MEM features utilizing Data Processor Service for Windows (DPSW) data. Default value is FALSE.
func (m *DataProcessorServiceForWindowsFeaturesOnboarding) SetAreDataProcessorServiceForWindowsFeaturesEnabled(value *bool)() {
    err := m.GetBackingStore().Set("areDataProcessorServiceForWindowsFeaturesEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the backingStore property value. Stores model information.
func (m *DataProcessorServiceForWindowsFeaturesOnboarding) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetHasValidWindowsLicense sets the hasValidWindowsLicense property value. Indicates whether the tenant has required Windows license. When TRUE, the tenant has the required Windows license. When FALSE, the tenant does not have the required Windows license. Default value is FALSE.
func (m *DataProcessorServiceForWindowsFeaturesOnboarding) SetHasValidWindowsLicense(value *bool)() {
    err := m.GetBackingStore().Set("hasValidWindowsLicense", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *DataProcessorServiceForWindowsFeaturesOnboarding) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// DataProcessorServiceForWindowsFeaturesOnboardingable 
type DataProcessorServiceForWindowsFeaturesOnboardingable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAreDataProcessorServiceForWindowsFeaturesEnabled()(*bool)
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetHasValidWindowsLicense()(*bool)
    GetOdataType()(*string)
    SetAreDataProcessorServiceForWindowsFeaturesEnabled(value *bool)()
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetHasValidWindowsLicense(value *bool)()
    SetOdataType(value *string)()
}
