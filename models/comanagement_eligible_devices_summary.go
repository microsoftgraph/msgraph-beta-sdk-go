package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// ComanagementEligibleDevicesSummary 
type ComanagementEligibleDevicesSummary struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewComanagementEligibleDevicesSummary instantiates a new comanagementEligibleDevicesSummary and sets the default values.
func NewComanagementEligibleDevicesSummary()(*ComanagementEligibleDevicesSummary) {
    m := &ComanagementEligibleDevicesSummary{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateComanagementEligibleDevicesSummaryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateComanagementEligibleDevicesSummaryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewComanagementEligibleDevicesSummary(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ComanagementEligibleDevicesSummary) GetAdditionalData()(map[string]any) {
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
// GetBackingStore gets the BackingStore property value. Stores model information.
func (m *ComanagementEligibleDevicesSummary) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetComanagedCount gets the comanagedCount property value. Count of devices already Co-Managed
func (m *ComanagementEligibleDevicesSummary) GetComanagedCount()(*int32) {
    val, err := m.GetBackingStore().Get("comanagedCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetEligibleButNotAzureAdJoinedCount gets the eligibleButNotAzureAdJoinedCount property value. Count of devices eligible for Co-Management but not yet joined to Azure Active Directory
func (m *ComanagementEligibleDevicesSummary) GetEligibleButNotAzureAdJoinedCount()(*int32) {
    val, err := m.GetBackingStore().Get("eligibleButNotAzureAdJoinedCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetEligibleCount gets the eligibleCount property value. Count of devices fully eligible for Co-Management
func (m *ComanagementEligibleDevicesSummary) GetEligibleCount()(*int32) {
    val, err := m.GetBackingStore().Get("eligibleCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ComanagementEligibleDevicesSummary) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["comanagedCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetComanagedCount(val)
        }
        return nil
    }
    res["eligibleButNotAzureAdJoinedCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEligibleButNotAzureAdJoinedCount(val)
        }
        return nil
    }
    res["eligibleCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEligibleCount(val)
        }
        return nil
    }
    res["ineligibleCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIneligibleCount(val)
        }
        return nil
    }
    res["needsOsUpdateCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNeedsOsUpdateCount(val)
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
    res["scheduledForEnrollmentCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetScheduledForEnrollmentCount(val)
        }
        return nil
    }
    return res
}
// GetIneligibleCount gets the ineligibleCount property value. Count of devices ineligible for Co-Management
func (m *ComanagementEligibleDevicesSummary) GetIneligibleCount()(*int32) {
    val, err := m.GetBackingStore().Get("ineligibleCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetNeedsOsUpdateCount gets the needsOsUpdateCount property value. Count of devices that will be eligible for Co-Management after an OS update
func (m *ComanagementEligibleDevicesSummary) GetNeedsOsUpdateCount()(*int32) {
    val, err := m.GetBackingStore().Get("needsOsUpdateCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *ComanagementEligibleDevicesSummary) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetScheduledForEnrollmentCount gets the scheduledForEnrollmentCount property value. Count of devices scheduled for Co-Management enrollment. Valid values 0 to 9999999
func (m *ComanagementEligibleDevicesSummary) GetScheduledForEnrollmentCount()(*int32) {
    val, err := m.GetBackingStore().Get("scheduledForEnrollmentCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// Serialize serializes information the current object
func (m *ComanagementEligibleDevicesSummary) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("comanagedCount", m.GetComanagedCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("eligibleButNotAzureAdJoinedCount", m.GetEligibleButNotAzureAdJoinedCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("eligibleCount", m.GetEligibleCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("ineligibleCount", m.GetIneligibleCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("needsOsUpdateCount", m.GetNeedsOsUpdateCount())
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
        err := writer.WriteInt32Value("scheduledForEnrollmentCount", m.GetScheduledForEnrollmentCount())
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
func (m *ComanagementEligibleDevicesSummary) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *ComanagementEligibleDevicesSummary) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetComanagedCount sets the comanagedCount property value. Count of devices already Co-Managed
func (m *ComanagementEligibleDevicesSummary) SetComanagedCount(value *int32)() {
    err := m.GetBackingStore().Set("comanagedCount", value)
    if err != nil {
        panic(err)
    }
}
// SetEligibleButNotAzureAdJoinedCount sets the eligibleButNotAzureAdJoinedCount property value. Count of devices eligible for Co-Management but not yet joined to Azure Active Directory
func (m *ComanagementEligibleDevicesSummary) SetEligibleButNotAzureAdJoinedCount(value *int32)() {
    err := m.GetBackingStore().Set("eligibleButNotAzureAdJoinedCount", value)
    if err != nil {
        panic(err)
    }
}
// SetEligibleCount sets the eligibleCount property value. Count of devices fully eligible for Co-Management
func (m *ComanagementEligibleDevicesSummary) SetEligibleCount(value *int32)() {
    err := m.GetBackingStore().Set("eligibleCount", value)
    if err != nil {
        panic(err)
    }
}
// SetIneligibleCount sets the ineligibleCount property value. Count of devices ineligible for Co-Management
func (m *ComanagementEligibleDevicesSummary) SetIneligibleCount(value *int32)() {
    err := m.GetBackingStore().Set("ineligibleCount", value)
    if err != nil {
        panic(err)
    }
}
// SetNeedsOsUpdateCount sets the needsOsUpdateCount property value. Count of devices that will be eligible for Co-Management after an OS update
func (m *ComanagementEligibleDevicesSummary) SetNeedsOsUpdateCount(value *int32)() {
    err := m.GetBackingStore().Set("needsOsUpdateCount", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *ComanagementEligibleDevicesSummary) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetScheduledForEnrollmentCount sets the scheduledForEnrollmentCount property value. Count of devices scheduled for Co-Management enrollment. Valid values 0 to 9999999
func (m *ComanagementEligibleDevicesSummary) SetScheduledForEnrollmentCount(value *int32)() {
    err := m.GetBackingStore().Set("scheduledForEnrollmentCount", value)
    if err != nil {
        panic(err)
    }
}
// ComanagementEligibleDevicesSummaryable 
type ComanagementEligibleDevicesSummaryable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetComanagedCount()(*int32)
    GetEligibleButNotAzureAdJoinedCount()(*int32)
    GetEligibleCount()(*int32)
    GetIneligibleCount()(*int32)
    GetNeedsOsUpdateCount()(*int32)
    GetOdataType()(*string)
    GetScheduledForEnrollmentCount()(*int32)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetComanagedCount(value *int32)()
    SetEligibleButNotAzureAdJoinedCount(value *int32)()
    SetEligibleCount(value *int32)()
    SetIneligibleCount(value *int32)()
    SetNeedsOsUpdateCount(value *int32)()
    SetOdataType(value *string)()
    SetScheduledForEnrollmentCount(value *int32)()
}
