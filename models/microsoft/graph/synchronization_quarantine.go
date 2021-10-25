package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type SynchronizationQuarantine struct {
    additionalData map[string]interface{};
    currentBegan *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    error *SynchronizationError;
    nextAttempt *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    reason *QuarantineReason;
    seriesBegan *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    seriesCount *int64;
}
func NewSynchronizationQuarantine()(*SynchronizationQuarantine) {
    m := &SynchronizationQuarantine{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *SynchronizationQuarantine) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *SynchronizationQuarantine) GetCurrentBegan()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.currentBegan
    }
}
func (m *SynchronizationQuarantine) GetError()(*SynchronizationError) {
    if m == nil {
        return nil
    } else {
        return m.error
    }
}
func (m *SynchronizationQuarantine) GetNextAttempt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.nextAttempt
    }
}
func (m *SynchronizationQuarantine) GetReason()(*QuarantineReason) {
    if m == nil {
        return nil
    } else {
        return m.reason
    }
}
func (m *SynchronizationQuarantine) GetSeriesBegan()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.seriesBegan
    }
}
func (m *SynchronizationQuarantine) GetSeriesCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.seriesCount
    }
}
func (m *SynchronizationQuarantine) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["currentBegan"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetCurrentBegan(val)
        return nil
    }
    res["error"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSynchronizationError() })
        if err != nil {
            return err
        }
        m.SetError(val.(*SynchronizationError))
        return nil
    }
    res["nextAttempt"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetNextAttempt(val)
        return nil
    }
    res["reason"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseQuarantineReason)
        if err != nil {
            return err
        }
        cast := val.(QuarantineReason)
        m.SetReason(&cast)
        return nil
    }
    res["seriesBegan"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetSeriesBegan(val)
        return nil
    }
    res["seriesCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetSeriesCount(val)
        return nil
    }
    return res
}
func (m *SynchronizationQuarantine) IsNil()(bool) {
    return m == nil
}
func (m *SynchronizationQuarantine) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteTimeValue("currentBegan", m.GetCurrentBegan())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("error", m.GetError())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("nextAttempt", m.GetNextAttempt())
        if err != nil {
            return err
        }
    }
    if m.GetReason() != nil {
        cast := m.GetReason().String()
        err := writer.WriteStringValue("reason", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("seriesBegan", m.GetSeriesBegan())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("seriesCount", m.GetSeriesCount())
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
func (m *SynchronizationQuarantine) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *SynchronizationQuarantine) SetCurrentBegan(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.currentBegan = value
}
func (m *SynchronizationQuarantine) SetError(value *SynchronizationError)() {
    m.error = value
}
func (m *SynchronizationQuarantine) SetNextAttempt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.nextAttempt = value
}
func (m *SynchronizationQuarantine) SetReason(value *QuarantineReason)() {
    m.reason = value
}
func (m *SynchronizationQuarantine) SetSeriesBegan(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.seriesBegan = value
}
func (m *SynchronizationQuarantine) SetSeriesCount(value *int64)() {
    m.seriesCount = value
}
