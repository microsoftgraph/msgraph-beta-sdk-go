package networkaccess

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// TrafficDistributionPoint 
type TrafficDistributionPoint struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewTrafficDistributionPoint instantiates a new trafficDistributionPoint and sets the default values.
func NewTrafficDistributionPoint()(*TrafficDistributionPoint) {
    m := &TrafficDistributionPoint{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateTrafficDistributionPointFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTrafficDistributionPointFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTrafficDistributionPoint(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TrafficDistributionPoint) GetAdditionalData()(map[string]any) {
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
func (m *TrafficDistributionPoint) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TrafficDistributionPoint) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["internetAccessTrafficCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInternetAccessTrafficCount(val)
        }
        return nil
    }
    res["microsoft365AccessTrafficCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMicrosoft365AccessTrafficCount(val)
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
    res["privateAccessTrafficCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrivateAccessTrafficCount(val)
        }
        return nil
    }
    res["timeStampDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTimeStampDateTime(val)
        }
        return nil
    }
    res["totalTrafficCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalTrafficCount(val)
        }
        return nil
    }
    return res
}
// GetInternetAccessTrafficCount gets the internetAccessTrafficCount property value. The internetAccessTrafficCount property
func (m *TrafficDistributionPoint) GetInternetAccessTrafficCount()(*int64) {
    val, err := m.GetBackingStore().Get("internetAccessTrafficCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// GetMicrosoft365AccessTrafficCount gets the microsoft365AccessTrafficCount property value. The microsoft365AccessTrafficCount property
func (m *TrafficDistributionPoint) GetMicrosoft365AccessTrafficCount()(*int64) {
    val, err := m.GetBackingStore().Get("microsoft365AccessTrafficCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *TrafficDistributionPoint) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPrivateAccessTrafficCount gets the privateAccessTrafficCount property value. The privateAccessTrafficCount property
func (m *TrafficDistributionPoint) GetPrivateAccessTrafficCount()(*int64) {
    val, err := m.GetBackingStore().Get("privateAccessTrafficCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// GetTimeStampDateTime gets the timeStampDateTime property value. The timeStampDateTime property
func (m *TrafficDistributionPoint) GetTimeStampDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("timeStampDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetTotalTrafficCount gets the totalTrafficCount property value. The totalTrafficCount property
func (m *TrafficDistributionPoint) GetTotalTrafficCount()(*int64) {
    val, err := m.GetBackingStore().Get("totalTrafficCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// Serialize serializes information the current object
func (m *TrafficDistributionPoint) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt64Value("internetAccessTrafficCount", m.GetInternetAccessTrafficCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("microsoft365AccessTrafficCount", m.GetMicrosoft365AccessTrafficCount())
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
        err := writer.WriteInt64Value("privateAccessTrafficCount", m.GetPrivateAccessTrafficCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("timeStampDateTime", m.GetTimeStampDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("totalTrafficCount", m.GetTotalTrafficCount())
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
func (m *TrafficDistributionPoint) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *TrafficDistributionPoint) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetInternetAccessTrafficCount sets the internetAccessTrafficCount property value. The internetAccessTrafficCount property
func (m *TrafficDistributionPoint) SetInternetAccessTrafficCount(value *int64)() {
    err := m.GetBackingStore().Set("internetAccessTrafficCount", value)
    if err != nil {
        panic(err)
    }
}
// SetMicrosoft365AccessTrafficCount sets the microsoft365AccessTrafficCount property value. The microsoft365AccessTrafficCount property
func (m *TrafficDistributionPoint) SetMicrosoft365AccessTrafficCount(value *int64)() {
    err := m.GetBackingStore().Set("microsoft365AccessTrafficCount", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *TrafficDistributionPoint) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetPrivateAccessTrafficCount sets the privateAccessTrafficCount property value. The privateAccessTrafficCount property
func (m *TrafficDistributionPoint) SetPrivateAccessTrafficCount(value *int64)() {
    err := m.GetBackingStore().Set("privateAccessTrafficCount", value)
    if err != nil {
        panic(err)
    }
}
// SetTimeStampDateTime sets the timeStampDateTime property value. The timeStampDateTime property
func (m *TrafficDistributionPoint) SetTimeStampDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("timeStampDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetTotalTrafficCount sets the totalTrafficCount property value. The totalTrafficCount property
func (m *TrafficDistributionPoint) SetTotalTrafficCount(value *int64)() {
    err := m.GetBackingStore().Set("totalTrafficCount", value)
    if err != nil {
        panic(err)
    }
}
// TrafficDistributionPointable 
type TrafficDistributionPointable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetInternetAccessTrafficCount()(*int64)
    GetMicrosoft365AccessTrafficCount()(*int64)
    GetOdataType()(*string)
    GetPrivateAccessTrafficCount()(*int64)
    GetTimeStampDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetTotalTrafficCount()(*int64)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetInternetAccessTrafficCount(value *int64)()
    SetMicrosoft365AccessTrafficCount(value *int64)()
    SetOdataType(value *string)()
    SetPrivateAccessTrafficCount(value *int64)()
    SetTimeStampDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetTotalTrafficCount(value *int64)()
}
