package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type SynchronizationProgress struct {
    additionalData map[string]interface{};
    completedUnits *int64;
    progressObservationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    totalUnits *int64;
    units *string;
}
func NewSynchronizationProgress()(*SynchronizationProgress) {
    m := &SynchronizationProgress{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *SynchronizationProgress) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *SynchronizationProgress) GetCompletedUnits()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.completedUnits
    }
}
func (m *SynchronizationProgress) GetProgressObservationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.progressObservationDateTime
    }
}
func (m *SynchronizationProgress) GetTotalUnits()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.totalUnits
    }
}
func (m *SynchronizationProgress) GetUnits()(*string) {
    if m == nil {
        return nil
    } else {
        return m.units
    }
}
func (m *SynchronizationProgress) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["completedUnits"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetCompletedUnits(val)
        return nil
    }
    res["progressObservationDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetProgressObservationDateTime(val)
        return nil
    }
    res["totalUnits"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetTotalUnits(val)
        return nil
    }
    res["units"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetUnits(val)
        return nil
    }
    return res
}
func (m *SynchronizationProgress) IsNil()(bool) {
    return m == nil
}
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
func (m *SynchronizationProgress) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *SynchronizationProgress) SetCompletedUnits(value *int64)() {
    m.completedUnits = value
}
func (m *SynchronizationProgress) SetProgressObservationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.progressObservationDateTime = value
}
func (m *SynchronizationProgress) SetTotalUnits(value *int64)() {
    m.totalUnits = value
}
func (m *SynchronizationProgress) SetUnits(value *string)() {
    m.units = value
}
