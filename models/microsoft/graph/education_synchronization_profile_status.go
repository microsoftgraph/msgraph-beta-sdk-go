package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type EducationSynchronizationProfileStatus struct {
    Entity
    lastActivityDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    lastSynchronizationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    status *EducationSynchronizationStatus;
}
func NewEducationSynchronizationProfileStatus()(*EducationSynchronizationProfileStatus) {
    m := &EducationSynchronizationProfileStatus{
        Entity: *NewEntity(),
    }
    return m
}
func (m *EducationSynchronizationProfileStatus) GetLastActivityDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastActivityDateTime
    }
}
func (m *EducationSynchronizationProfileStatus) GetLastSynchronizationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastSynchronizationDateTime
    }
}
func (m *EducationSynchronizationProfileStatus) GetStatus()(*EducationSynchronizationStatus) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
func (m *EducationSynchronizationProfileStatus) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["lastActivityDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetLastActivityDateTime(val)
        return nil
    }
    res["lastSynchronizationDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetLastSynchronizationDateTime(val)
        return nil
    }
    res["status"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseEducationSynchronizationStatus)
        if err != nil {
            return err
        }
        cast := val.(EducationSynchronizationStatus)
        m.SetStatus(&cast)
        return nil
    }
    return res
}
func (m *EducationSynchronizationProfileStatus) IsNil()(bool) {
    return m == nil
}
func (m *EducationSynchronizationProfileStatus) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("lastActivityDateTime", m.GetLastActivityDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastSynchronizationDateTime", m.GetLastSynchronizationDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetStatus() != nil {
        cast := m.GetStatus().String()
        err = writer.WriteStringValue("status", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *EducationSynchronizationProfileStatus) SetLastActivityDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastActivityDateTime = value
}
func (m *EducationSynchronizationProfileStatus) SetLastSynchronizationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastSynchronizationDateTime = value
}
func (m *EducationSynchronizationProfileStatus) SetStatus(value *EducationSynchronizationStatus)() {
    m.status = value
}
