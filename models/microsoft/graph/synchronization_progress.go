package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// SynchronizationProgress 
type SynchronizationProgress struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The numerator of a progress ratio; the number of units of changes already processed.
    completedUnits *int64;
    // The time of a progress observation as an offset in minutes from UTC.
    progressObservationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The denominator of a progress ratio; a number of units of changes to be processed to accomplish synchronization.
    totalUnits *int64;
    // An optional description of the units.
    units *string;
}
// NewSynchronizationProgress instantiates a new synchronizationProgress and sets the default values.
func NewSynchronizationProgress()(*SynchronizationProgress) {
    m := &SynchronizationProgress{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SynchronizationProgress) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetCompletedUnits gets the completedUnits property value. The numerator of a progress ratio; the number of units of changes already processed.
func (m *SynchronizationProgress) GetCompletedUnits()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.completedUnits
    }
}
// GetProgressObservationDateTime gets the progressObservationDateTime property value. The time of a progress observation as an offset in minutes from UTC.
func (m *SynchronizationProgress) GetProgressObservationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.progressObservationDateTime
    }
}
// GetTotalUnits gets the totalUnits property value. The denominator of a progress ratio; a number of units of changes to be processed to accomplish synchronization.
func (m *SynchronizationProgress) GetTotalUnits()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.totalUnits
    }
}
// GetUnits gets the units property value. An optional description of the units.
func (m *SynchronizationProgress) GetUnits()(*string) {
    if m == nil {
        return nil
    } else {
        return m.units
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SynchronizationProgress) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["completedUnits"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCompletedUnits(val)
        }
        return nil
    }
    res["progressObservationDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProgressObservationDateTime(val)
        }
        return nil
    }
    res["totalUnits"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalUnits(val)
        }
        return nil
    }
    res["units"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUnits(val)
        }
        return nil
    }
    return res
}
func (m *SynchronizationProgress) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *SynchronizationProgress) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteInt64Value("completedUnits", m.GetCompletedUnits())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("progressObservationDateTime", m.GetProgressObservationDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("totalUnits", m.GetTotalUnits())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("units", m.GetUnits())
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
func (m *SynchronizationProgress) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetCompletedUnits sets the completedUnits property value. The numerator of a progress ratio; the number of units of changes already processed.
func (m *SynchronizationProgress) SetCompletedUnits(value *int64)() {
    if m != nil {
        m.completedUnits = value
    }
}
// SetProgressObservationDateTime sets the progressObservationDateTime property value. The time of a progress observation as an offset in minutes from UTC.
func (m *SynchronizationProgress) SetProgressObservationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.progressObservationDateTime = value
    }
}
// SetTotalUnits sets the totalUnits property value. The denominator of a progress ratio; a number of units of changes to be processed to accomplish synchronization.
func (m *SynchronizationProgress) SetTotalUnits(value *int64)() {
    if m != nil {
        m.totalUnits = value
    }
}
// SetUnits sets the units property value. An optional description of the units.
func (m *SynchronizationProgress) SetUnits(value *string)() {
    if m != nil {
        m.units = value
    }
}
