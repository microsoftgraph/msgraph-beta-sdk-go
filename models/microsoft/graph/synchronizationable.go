package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Synchronizationable 
type Synchronizationable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetJobs()([]SynchronizationJobable)
    GetSecrets()([]SynchronizationSecretKeyStringValuePairable)
    GetTemplates()([]SynchronizationTemplateable)
    SetJobs(value []SynchronizationJobable)()
    SetSecrets(value []SynchronizationSecretKeyStringValuePairable)()
    SetTemplates(value []SynchronizationTemplateable)()
}
