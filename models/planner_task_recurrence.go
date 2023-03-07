package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// PlannerTaskRecurrence 
type PlannerTaskRecurrence struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewPlannerTaskRecurrence instantiates a new plannerTaskRecurrence and sets the default values.
func NewPlannerTaskRecurrence()(*PlannerTaskRecurrence) {
    m := &PlannerTaskRecurrence{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreatePlannerTaskRecurrenceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePlannerTaskRecurrenceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPlannerTaskRecurrence(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PlannerTaskRecurrence) GetAdditionalData()(map[string]any) {
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
// GetBackingStore gets the backingStore property value. Stores model information.
func (m *PlannerTaskRecurrence) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PlannerTaskRecurrence) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["nextInSeriesTaskId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNextInSeriesTaskId(val)
        }
        return nil
    }
    res["occurrenceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOccurrenceId(val)
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
    res["previousInSeriesTaskId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPreviousInSeriesTaskId(val)
        }
        return nil
    }
    res["recurrenceStartDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRecurrenceStartDateTime(val)
        }
        return nil
    }
    res["schedule"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePlannerRecurrenceScheduleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSchedule(val.(PlannerRecurrenceScheduleable))
        }
        return nil
    }
    res["seriesId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSeriesId(val)
        }
        return nil
    }
    return res
}
// GetNextInSeriesTaskId gets the nextInSeriesTaskId property value. The nextInSeriesTaskId property
func (m *PlannerTaskRecurrence) GetNextInSeriesTaskId()(*string) {
    val, err := m.GetBackingStore().Get("nextInSeriesTaskId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetOccurrenceId gets the occurrenceId property value. The occurrenceId property
func (m *PlannerTaskRecurrence) GetOccurrenceId()(*int32) {
    val, err := m.GetBackingStore().Get("occurrenceId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *PlannerTaskRecurrence) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPreviousInSeriesTaskId gets the previousInSeriesTaskId property value. The previousInSeriesTaskId property
func (m *PlannerTaskRecurrence) GetPreviousInSeriesTaskId()(*string) {
    val, err := m.GetBackingStore().Get("previousInSeriesTaskId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetRecurrenceStartDateTime gets the recurrenceStartDateTime property value. The recurrenceStartDateTime property
func (m *PlannerTaskRecurrence) GetRecurrenceStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("recurrenceStartDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetSchedule gets the schedule property value. The schedule property
func (m *PlannerTaskRecurrence) GetSchedule()(PlannerRecurrenceScheduleable) {
    val, err := m.GetBackingStore().Get("schedule")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(PlannerRecurrenceScheduleable)
    }
    return nil
}
// GetSeriesId gets the seriesId property value. The seriesId property
func (m *PlannerTaskRecurrence) GetSeriesId()(*string) {
    val, err := m.GetBackingStore().Get("seriesId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *PlannerTaskRecurrence) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("nextInSeriesTaskId", m.GetNextInSeriesTaskId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("occurrenceId", m.GetOccurrenceId())
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
        err := writer.WriteStringValue("previousInSeriesTaskId", m.GetPreviousInSeriesTaskId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("recurrenceStartDateTime", m.GetRecurrenceStartDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("schedule", m.GetSchedule())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("seriesId", m.GetSeriesId())
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
func (m *PlannerTaskRecurrence) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the backingStore property value. Stores model information.
func (m *PlannerTaskRecurrence) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetNextInSeriesTaskId sets the nextInSeriesTaskId property value. The nextInSeriesTaskId property
func (m *PlannerTaskRecurrence) SetNextInSeriesTaskId(value *string)() {
    err := m.GetBackingStore().Set("nextInSeriesTaskId", value)
    if err != nil {
        panic(err)
    }
}
// SetOccurrenceId sets the occurrenceId property value. The occurrenceId property
func (m *PlannerTaskRecurrence) SetOccurrenceId(value *int32)() {
    err := m.GetBackingStore().Set("occurrenceId", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *PlannerTaskRecurrence) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetPreviousInSeriesTaskId sets the previousInSeriesTaskId property value. The previousInSeriesTaskId property
func (m *PlannerTaskRecurrence) SetPreviousInSeriesTaskId(value *string)() {
    err := m.GetBackingStore().Set("previousInSeriesTaskId", value)
    if err != nil {
        panic(err)
    }
}
// SetRecurrenceStartDateTime sets the recurrenceStartDateTime property value. The recurrenceStartDateTime property
func (m *PlannerTaskRecurrence) SetRecurrenceStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("recurrenceStartDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetSchedule sets the schedule property value. The schedule property
func (m *PlannerTaskRecurrence) SetSchedule(value PlannerRecurrenceScheduleable)() {
    err := m.GetBackingStore().Set("schedule", value)
    if err != nil {
        panic(err)
    }
}
// SetSeriesId sets the seriesId property value. The seriesId property
func (m *PlannerTaskRecurrence) SetSeriesId(value *string)() {
    err := m.GetBackingStore().Set("seriesId", value)
    if err != nil {
        panic(err)
    }
}
// PlannerTaskRecurrenceable 
type PlannerTaskRecurrenceable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetNextInSeriesTaskId()(*string)
    GetOccurrenceId()(*int32)
    GetOdataType()(*string)
    GetPreviousInSeriesTaskId()(*string)
    GetRecurrenceStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetSchedule()(PlannerRecurrenceScheduleable)
    GetSeriesId()(*string)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetNextInSeriesTaskId(value *string)()
    SetOccurrenceId(value *int32)()
    SetOdataType(value *string)()
    SetPreviousInSeriesTaskId(value *string)()
    SetRecurrenceStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetSchedule(value PlannerRecurrenceScheduleable)()
    SetSeriesId(value *string)()
}
