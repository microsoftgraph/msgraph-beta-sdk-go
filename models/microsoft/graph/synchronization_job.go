package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type SynchronizationJob struct {
    Entity
    // Schedule used to run the job. Read-only.
    schedule *SynchronizationSchedule;
    // The synchronization schema configured for the job.
    schema *SynchronizationSchema;
    // Status of the job, which includes when the job was last run, current job state, and errors.
    status *SynchronizationStatus;
    // Settings associated with the job. Some settings are inherited from the template.
    synchronizationJobSettings []KeyValuePair;
    // Identifier of the synchronization template this job is based on.
    templateId *string;
}
// Instantiates a new synchronizationJob and sets the default values.
func NewSynchronizationJob()(*SynchronizationJob) {
    m := &SynchronizationJob{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the schedule property value. Schedule used to run the job. Read-only.
func (m *SynchronizationJob) GetSchedule()(*SynchronizationSchedule) {
    if m == nil {
        return nil
    } else {
        return m.schedule
    }
}
// Gets the schema property value. The synchronization schema configured for the job.
func (m *SynchronizationJob) GetSchema()(*SynchronizationSchema) {
    if m == nil {
        return nil
    } else {
        return m.schema
    }
}
// Gets the status property value. Status of the job, which includes when the job was last run, current job state, and errors.
func (m *SynchronizationJob) GetStatus()(*SynchronizationStatus) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
// Gets the synchronizationJobSettings property value. Settings associated with the job. Some settings are inherited from the template.
func (m *SynchronizationJob) GetSynchronizationJobSettings()([]KeyValuePair) {
    if m == nil {
        return nil
    } else {
        return m.synchronizationJobSettings
    }
}
// Gets the templateId property value. Identifier of the synchronization template this job is based on.
func (m *SynchronizationJob) GetTemplateId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.templateId
    }
}
// The deserialization information for the current model
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
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the schedule property value. Schedule used to run the job. Read-only.
// Parameters:
//  - value : Value to set for the schedule property.
func (m *SynchronizationJob) SetSchedule(value *SynchronizationSchedule)() {
    m.schedule = value
}
// Sets the schema property value. The synchronization schema configured for the job.
// Parameters:
//  - value : Value to set for the schema property.
func (m *SynchronizationJob) SetSchema(value *SynchronizationSchema)() {
    m.schema = value
}
// Sets the status property value. Status of the job, which includes when the job was last run, current job state, and errors.
// Parameters:
//  - value : Value to set for the status property.
func (m *SynchronizationJob) SetStatus(value *SynchronizationStatus)() {
    m.status = value
}
// Sets the synchronizationJobSettings property value. Settings associated with the job. Some settings are inherited from the template.
// Parameters:
//  - value : Value to set for the synchronizationJobSettings property.
func (m *SynchronizationJob) SetSynchronizationJobSettings(value []KeyValuePair)() {
    m.synchronizationJobSettings = value
}
// Sets the templateId property value. Identifier of the synchronization template this job is based on.
// Parameters:
//  - value : Value to set for the templateId property.
func (m *SynchronizationJob) SetTemplateId(value *string)() {
    m.templateId = value
}
