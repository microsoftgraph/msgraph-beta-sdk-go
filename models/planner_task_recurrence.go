package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PlannerTaskRecurrence 
type PlannerTaskRecurrence struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The nextInSeriesTaskId property
    nextInSeriesTaskId *string
    // The occurrenceId property
    occurrenceId *int32
    // The OdataType property
    odataType *string
    // The previousInSeriesTaskId property
    previousInSeriesTaskId *string
    // The recurrenceStartDateTime property
    recurrenceStartDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The schedule property
    schedule PlannerRecurrenceScheduleable
    // The seriesId property
    seriesId *string
}
// NewPlannerTaskRecurrence instantiates a new plannerTaskRecurrence and sets the default values.
func NewPlannerTaskRecurrence()(*PlannerTaskRecurrence) {
    m := &PlannerTaskRecurrence{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreatePlannerTaskRecurrenceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePlannerTaskRecurrenceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPlannerTaskRecurrence(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PlannerTaskRecurrence) GetAdditionalData()(map[string]any) {
    return m.additionalData
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
    return m.nextInSeriesTaskId
}
// GetOccurrenceId gets the occurrenceId property value. The occurrenceId property
func (m *PlannerTaskRecurrence) GetOccurrenceId()(*int32) {
    return m.occurrenceId
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *PlannerTaskRecurrence) GetOdataType()(*string) {
    return m.odataType
}
// GetPreviousInSeriesTaskId gets the previousInSeriesTaskId property value. The previousInSeriesTaskId property
func (m *PlannerTaskRecurrence) GetPreviousInSeriesTaskId()(*string) {
    return m.previousInSeriesTaskId
}
// GetRecurrenceStartDateTime gets the recurrenceStartDateTime property value. The recurrenceStartDateTime property
func (m *PlannerTaskRecurrence) GetRecurrenceStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.recurrenceStartDateTime
}
// GetSchedule gets the schedule property value. The schedule property
func (m *PlannerTaskRecurrence) GetSchedule()(PlannerRecurrenceScheduleable) {
    return m.schedule
}
// GetSeriesId gets the seriesId property value. The seriesId property
func (m *PlannerTaskRecurrence) GetSeriesId()(*string) {
    return m.seriesId
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
    m.additionalData = value
}
// SetNextInSeriesTaskId sets the nextInSeriesTaskId property value. The nextInSeriesTaskId property
func (m *PlannerTaskRecurrence) SetNextInSeriesTaskId(value *string)() {
    m.nextInSeriesTaskId = value
}
// SetOccurrenceId sets the occurrenceId property value. The occurrenceId property
func (m *PlannerTaskRecurrence) SetOccurrenceId(value *int32)() {
    m.occurrenceId = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *PlannerTaskRecurrence) SetOdataType(value *string)() {
    m.odataType = value
}
// SetPreviousInSeriesTaskId sets the previousInSeriesTaskId property value. The previousInSeriesTaskId property
func (m *PlannerTaskRecurrence) SetPreviousInSeriesTaskId(value *string)() {
    m.previousInSeriesTaskId = value
}
// SetRecurrenceStartDateTime sets the recurrenceStartDateTime property value. The recurrenceStartDateTime property
func (m *PlannerTaskRecurrence) SetRecurrenceStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.recurrenceStartDateTime = value
}
// SetSchedule sets the schedule property value. The schedule property
func (m *PlannerTaskRecurrence) SetSchedule(value PlannerRecurrenceScheduleable)() {
    m.schedule = value
}
// SetSeriesId sets the seriesId property value. The seriesId property
func (m *PlannerTaskRecurrence) SetSeriesId(value *string)() {
    m.seriesId = value
}
