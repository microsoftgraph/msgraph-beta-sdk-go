package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PlannerRecurrenceSchedule 
type PlannerRecurrenceSchedule struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The nextOccurrenceDateTime property
    nextOccurrenceDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The OdataType property
    odataType *string
    // The pattern property
    pattern RecurrencePatternable
    // The patternStartDateTime property
    patternStartDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
}
// NewPlannerRecurrenceSchedule instantiates a new plannerRecurrenceSchedule and sets the default values.
func NewPlannerRecurrenceSchedule()(*PlannerRecurrenceSchedule) {
    m := &PlannerRecurrenceSchedule{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreatePlannerRecurrenceScheduleFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePlannerRecurrenceScheduleFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPlannerRecurrenceSchedule(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PlannerRecurrenceSchedule) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PlannerRecurrenceSchedule) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["nextOccurrenceDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNextOccurrenceDateTime(val)
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
    res["pattern"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateRecurrencePatternFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPattern(val.(RecurrencePatternable))
        }
        return nil
    }
    res["patternStartDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPatternStartDateTime(val)
        }
        return nil
    }
    return res
}
// GetNextOccurrenceDateTime gets the nextOccurrenceDateTime property value. The nextOccurrenceDateTime property
func (m *PlannerRecurrenceSchedule) GetNextOccurrenceDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.nextOccurrenceDateTime
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *PlannerRecurrenceSchedule) GetOdataType()(*string) {
    return m.odataType
}
// GetPattern gets the pattern property value. The pattern property
func (m *PlannerRecurrenceSchedule) GetPattern()(RecurrencePatternable) {
    return m.pattern
}
// GetPatternStartDateTime gets the patternStartDateTime property value. The patternStartDateTime property
func (m *PlannerRecurrenceSchedule) GetPatternStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.patternStartDateTime
}
// Serialize serializes information the current object
func (m *PlannerRecurrenceSchedule) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteTimeValue("nextOccurrenceDateTime", m.GetNextOccurrenceDateTime())
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
        err := writer.WriteObjectValue("pattern", m.GetPattern())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("patternStartDateTime", m.GetPatternStartDateTime())
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
func (m *PlannerRecurrenceSchedule) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetNextOccurrenceDateTime sets the nextOccurrenceDateTime property value. The nextOccurrenceDateTime property
func (m *PlannerRecurrenceSchedule) SetNextOccurrenceDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.nextOccurrenceDateTime = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *PlannerRecurrenceSchedule) SetOdataType(value *string)() {
    m.odataType = value
}
// SetPattern sets the pattern property value. The pattern property
func (m *PlannerRecurrenceSchedule) SetPattern(value RecurrencePatternable)() {
    m.pattern = value
}
// SetPatternStartDateTime sets the patternStartDateTime property value. The patternStartDateTime property
func (m *PlannerRecurrenceSchedule) SetPatternStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.patternStartDateTime = value
}
