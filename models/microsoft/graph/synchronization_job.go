package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type SynchronizationJob struct {
    Entity
    schedule *SynchronizationSchedule;
    schema *SynchronizationSchema;
    status *SynchronizationStatus;
    synchronizationJobSettings []KeyValuePair;
    templateId *string;
}
func NewSynchronizationJob()(*SynchronizationJob) {
    m := &SynchronizationJob{
        Entity: *NewEntity(),
    }
    return m
}
func (m *SynchronizationJob) GetSchedule()(*SynchronizationSchedule) {
    if m == nil {
        return nil
    } else {
        return m.schedule
    }
}
func (m *SynchronizationJob) GetSchema()(*SynchronizationSchema) {
    if m == nil {
        return nil
    } else {
        return m.schema
    }
}
func (m *SynchronizationJob) GetStatus()(*SynchronizationStatus) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
func (m *SynchronizationJob) GetSynchronizationJobSettings()([]KeyValuePair) {
    if m == nil {
        return nil
    } else {
        return m.synchronizationJobSettings
    }
}
func (m *SynchronizationJob) GetTemplateId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.templateId
    }
}
func (m *SynchronizationJob) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["schedule"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSynchronizationSchedule() })
        if err != nil {
            return err
        }
        m.SetSchedule(val.(*SynchronizationSchedule))
        return nil
    }
    res["schema"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSynchronizationSchema() })
        if err != nil {
            return err
        }
        m.SetSchema(val.(*SynchronizationSchema))
        return nil
    }
    res["status"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSynchronizationStatus() })
        if err != nil {
            return err
        }
        m.SetStatus(val.(*SynchronizationStatus))
        return nil
    }
    res["synchronizationJobSettings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewKeyValuePair() })
        if err != nil {
            return err
        }
        res := make([]KeyValuePair, len(val))
        for i, v := range val {
            res[i] = *(v.(*KeyValuePair))
        }
        m.SetSynchronizationJobSettings(res)
        return nil
    }
    res["templateId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetTemplateId(val)
        return nil
    }
    return res
}
func (m *SynchronizationJob) IsNil()(bool) {
    return m == nil
}
func (m *SynchronizationJob) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("schedule", m.GetSchedule())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("schema", m.GetSchema())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("status", m.GetStatus())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetSynchronizationJobSettings()))
        for i, v := range m.GetSynchronizationJobSettings() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("synchronizationJobSettings", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("templateId", m.GetTemplateId())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *SynchronizationJob) SetSchedule(value *SynchronizationSchedule)() {
    m.schedule = value
}
func (m *SynchronizationJob) SetSchema(value *SynchronizationSchema)() {
    m.schema = value
}
func (m *SynchronizationJob) SetStatus(value *SynchronizationStatus)() {
    m.status = value
}
func (m *SynchronizationJob) SetSynchronizationJobSettings(value []KeyValuePair)() {
    m.synchronizationJobSettings = value
}
func (m *SynchronizationJob) SetTemplateId(value *string)() {
    m.templateId = value
}
