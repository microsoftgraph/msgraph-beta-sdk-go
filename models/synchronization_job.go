package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SynchronizationJob 
type SynchronizationJob struct {
    Entity
    // Schedule used to run the job. Read-only.
    schedule SynchronizationScheduleable
    // The synchronization schema configured for the job.
    schema SynchronizationSchemaable
    // Status of the job, which includes when the job was last run, current job state, and errors.
    status SynchronizationStatusable
    // Settings associated with the job. Some settings are inherited from the template.
    synchronizationJobSettings []KeyValuePairable
    // Identifier of the synchronization template this job is based on.
    templateId *string
}
// NewSynchronizationJob instantiates a new synchronizationJob and sets the default values.
func NewSynchronizationJob()(*SynchronizationJob) {
    m := &SynchronizationJob{
        Entity: *NewEntity(),
    }
    return m
}
// CreateSynchronizationJobFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSynchronizationJobFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSynchronizationJob(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SynchronizationJob) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["schedule"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSynchronizationScheduleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSchedule(val.(SynchronizationScheduleable))
        }
        return nil
    }
    res["schema"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSynchronizationSchemaFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSchema(val.(SynchronizationSchemaable))
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSynchronizationStatusFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(SynchronizationStatusable))
        }
        return nil
    }
    res["synchronizationJobSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateKeyValuePairFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]KeyValuePairable, len(val))
            for i, v := range val {
                res[i] = v.(KeyValuePairable)
            }
            m.SetSynchronizationJobSettings(res)
        }
        return nil
    }
    res["templateId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTemplateId(val)
        }
        return nil
    }
    return res
}
// GetSchedule gets the schedule property value. Schedule used to run the job. Read-only.
func (m *SynchronizationJob) GetSchedule()(SynchronizationScheduleable) {
    if m == nil {
        return nil
    } else {
        return m.schedule
    }
}
// GetSchema gets the schema property value. The synchronization schema configured for the job.
func (m *SynchronizationJob) GetSchema()(SynchronizationSchemaable) {
    if m == nil {
        return nil
    } else {
        return m.schema
    }
}
// GetStatus gets the status property value. Status of the job, which includes when the job was last run, current job state, and errors.
func (m *SynchronizationJob) GetStatus()(SynchronizationStatusable) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
// GetSynchronizationJobSettings gets the synchronizationJobSettings property value. Settings associated with the job. Some settings are inherited from the template.
func (m *SynchronizationJob) GetSynchronizationJobSettings()([]KeyValuePairable) {
    if m == nil {
        return nil
    } else {
        return m.synchronizationJobSettings
    }
}
// GetTemplateId gets the templateId property value. Identifier of the synchronization template this job is based on.
func (m *SynchronizationJob) GetTemplateId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.templateId
    }
}
// Serialize serializes information the current object
func (m *SynchronizationJob) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    if m.GetSynchronizationJobSettings() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSynchronizationJobSettings()))
        for i, v := range m.GetSynchronizationJobSettings() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
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
// SetSchedule sets the schedule property value. Schedule used to run the job. Read-only.
func (m *SynchronizationJob) SetSchedule(value SynchronizationScheduleable)() {
    if m != nil {
        m.schedule = value
    }
}
// SetSchema sets the schema property value. The synchronization schema configured for the job.
func (m *SynchronizationJob) SetSchema(value SynchronizationSchemaable)() {
    if m != nil {
        m.schema = value
    }
}
// SetStatus sets the status property value. Status of the job, which includes when the job was last run, current job state, and errors.
func (m *SynchronizationJob) SetStatus(value SynchronizationStatusable)() {
    if m != nil {
        m.status = value
    }
}
// SetSynchronizationJobSettings sets the synchronizationJobSettings property value. Settings associated with the job. Some settings are inherited from the template.
func (m *SynchronizationJob) SetSynchronizationJobSettings(value []KeyValuePairable)() {
    if m != nil {
        m.synchronizationJobSettings = value
    }
}
// SetTemplateId sets the templateId property value. Identifier of the synchronization template this job is based on.
func (m *SynchronizationJob) SetTemplateId(value *string)() {
    if m != nil {
        m.templateId = value
    }
}
