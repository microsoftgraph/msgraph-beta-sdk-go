package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type SynchronizationSchedule struct {
    additionalData map[string]interface{};
    expiration *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    interval *string;
    state *SynchronizationScheduleState;
}
func NewSynchronizationSchedule()(*SynchronizationSchedule) {
    m := &SynchronizationSchedule{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *SynchronizationSchedule) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *SynchronizationSchedule) GetExpiration()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.expiration
    }
}
func (m *SynchronizationSchedule) GetInterval()(*string) {
    if m == nil {
        return nil
    } else {
        return m.interval
    }
}
func (m *SynchronizationSchedule) GetState()(*SynchronizationScheduleState) {
    if m == nil {
        return nil
    } else {
        return m.state
    }
}
func (m *SynchronizationSchedule) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["expiration"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetExpiration(val)
        return nil
    }
    res["interval"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetInterval(val)
        return nil
    }
    res["state"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseSynchronizationScheduleState)
        if err != nil {
            return err
        }
        cast := val.(SynchronizationScheduleState)
        m.SetState(&cast)
        return nil
    }
    return res
}
func (m *SynchronizationSchedule) IsNil()(bool) {
    return m == nil
}
func (m *SynchronizationSchedule) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteTimeValue("expiration", m.GetExpiration())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("interval", m.GetInterval())
        if err != nil {
            return err
        }
    }
    if m.GetState() != nil {
        cast := m.GetState().String()
        err := writer.WriteStringValue("state", &cast)
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
func (m *SynchronizationSchedule) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *SynchronizationSchedule) SetExpiration(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.expiration = value
}
func (m *SynchronizationSchedule) SetInterval(value *string)() {
    m.interval = value
}
func (m *SynchronizationSchedule) SetState(value *SynchronizationScheduleState)() {
    m.state = value
}
