package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// SynchronizationJobable 
type SynchronizationJobable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetSchedule()(SynchronizationScheduleable)
    GetSchema()(SynchronizationSchemaable)
    GetStatus()(SynchronizationStatusable)
    GetSynchronizationJobSettings()([]KeyValuePairable)
    GetTemplateId()(*string)
    SetSchedule(value SynchronizationScheduleable)()
    SetSchema(value SynchronizationSchemaable)()
    SetStatus(value SynchronizationStatusable)()
    SetSynchronizationJobSettings(value []KeyValuePairable)()
    SetTemplateId(value *string)()
}
