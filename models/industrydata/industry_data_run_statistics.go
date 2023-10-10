package industrydata

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// IndustryDataRunStatistics 
type IndustryDataRunStatistics struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewIndustryDataRunStatistics instantiates a new industryDataRunStatistics and sets the default values.
func NewIndustryDataRunStatistics()(*IndustryDataRunStatistics) {
    m := &IndustryDataRunStatistics{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateIndustryDataRunStatisticsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateIndustryDataRunStatisticsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIndustryDataRunStatistics(), nil
}
// GetActivityStatistics gets the activityStatistics property value. The collection of statistics for each activity included in this run.
func (m *IndustryDataRunStatistics) GetActivityStatistics()([]IndustryDataActivityStatisticsable) {
    val, err := m.GetBackingStore().Get("activityStatistics")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]IndustryDataActivityStatisticsable)
    }
    return nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *IndustryDataRunStatistics) GetAdditionalData()(map[string]any) {
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
func (m *IndustryDataRunStatistics) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetFieldDeserializers the deserialization information for the current model
func (m *IndustryDataRunStatistics) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["activityStatistics"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateIndustryDataActivityStatisticsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]IndustryDataActivityStatisticsable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(IndustryDataActivityStatisticsable)
                }
            }
            m.SetActivityStatistics(res)
        }
        return nil
    }
    res["inboundTotals"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAggregatedInboundStatisticsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInboundTotals(val.(AggregatedInboundStatisticsable))
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
    res["runId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRunId(val)
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseIndustryDataRunStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*IndustryDataRunStatus))
        }
        return nil
    }
    return res
}
// GetInboundTotals gets the inboundTotals property value. The aggregate statistics for all inbound flows.
func (m *IndustryDataRunStatistics) GetInboundTotals()(AggregatedInboundStatisticsable) {
    val, err := m.GetBackingStore().Get("inboundTotals")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(AggregatedInboundStatisticsable)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *IndustryDataRunStatistics) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetRunId gets the runId property value. The ID of the underlying run for the statistics.
func (m *IndustryDataRunStatistics) GetRunId()(*string) {
    val, err := m.GetBackingStore().Get("runId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetStatus gets the status property value. The status property
func (m *IndustryDataRunStatistics) GetStatus()(*IndustryDataRunStatus) {
    val, err := m.GetBackingStore().Get("status")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*IndustryDataRunStatus)
    }
    return nil
}
// Serialize serializes information the current object
func (m *IndustryDataRunStatistics) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    if m.GetStatus() != nil {
        cast := (*m.GetStatus()).String()
        err := writer.WriteStringValue("status", &cast)
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
// SetActivityStatistics sets the activityStatistics property value. The collection of statistics for each activity included in this run.
func (m *IndustryDataRunStatistics) SetActivityStatistics(value []IndustryDataActivityStatisticsable)() {
    err := m.GetBackingStore().Set("activityStatistics", value)
    if err != nil {
        panic(err)
    }
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *IndustryDataRunStatistics) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *IndustryDataRunStatistics) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetInboundTotals sets the inboundTotals property value. The aggregate statistics for all inbound flows.
func (m *IndustryDataRunStatistics) SetInboundTotals(value AggregatedInboundStatisticsable)() {
    err := m.GetBackingStore().Set("inboundTotals", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *IndustryDataRunStatistics) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetRunId sets the runId property value. The ID of the underlying run for the statistics.
func (m *IndustryDataRunStatistics) SetRunId(value *string)() {
    err := m.GetBackingStore().Set("runId", value)
    if err != nil {
        panic(err)
    }
}
// SetStatus sets the status property value. The status property
func (m *IndustryDataRunStatistics) SetStatus(value *IndustryDataRunStatus)() {
    err := m.GetBackingStore().Set("status", value)
    if err != nil {
        panic(err)
    }
}
// IndustryDataRunStatisticsable 
type IndustryDataRunStatisticsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActivityStatistics()([]IndustryDataActivityStatisticsable)
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetInboundTotals()(AggregatedInboundStatisticsable)
    GetOdataType()(*string)
    GetRunId()(*string)
    GetStatus()(*IndustryDataRunStatus)
    SetActivityStatistics(value []IndustryDataActivityStatisticsable)()
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetInboundTotals(value AggregatedInboundStatisticsable)()
    SetOdataType(value *string)()
    SetRunId(value *string)()
    SetStatus(value *IndustryDataRunStatus)()
}
