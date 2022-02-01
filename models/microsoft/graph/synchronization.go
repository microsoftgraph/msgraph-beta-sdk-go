package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Synchronization 
type Synchronization struct {
    Entity
    // Performs synchronization by periodically running in the background, polling for changes in one directory, and pushing them to another directory.
    jobs []SynchronizationJob;
    // Represents a collection of credentials to access provisioned cloud applications.
    secrets []SynchronizationSecretKeyStringValuePair;
    // Pre-configured synchronization settings for a particular application.
    templates []SynchronizationTemplate;
}
// NewSynchronization instantiates a new synchronization and sets the default values.
func NewSynchronization()(*Synchronization) {
    m := &Synchronization{
        Entity: *NewEntity(),
    }
    return m
}
// GetJobs gets the jobs property value. Performs synchronization by periodically running in the background, polling for changes in one directory, and pushing them to another directory.
func (m *Synchronization) GetJobs()([]SynchronizationJob) {
    if m == nil {
        return nil
    } else {
        return m.jobs
    }
}
// GetSecrets gets the secrets property value. Represents a collection of credentials to access provisioned cloud applications.
func (m *Synchronization) GetSecrets()([]SynchronizationSecretKeyStringValuePair) {
    if m == nil {
        return nil
    } else {
        return m.secrets
    }
}
// GetTemplates gets the templates property value. Pre-configured synchronization settings for a particular application.
func (m *Synchronization) GetTemplates()([]SynchronizationTemplate) {
    if m == nil {
        return nil
    } else {
        return m.templates
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Synchronization) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["jobs"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSynchronizationJob() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SynchronizationJob, len(val))
            for i, v := range val {
                res[i] = *(v.(*SynchronizationJob))
            }
            m.SetJobs(res)
        }
        return nil
    }
    res["secrets"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSynchronizationSecretKeyStringValuePair() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SynchronizationSecretKeyStringValuePair, len(val))
            for i, v := range val {
                res[i] = *(v.(*SynchronizationSecretKeyStringValuePair))
            }
            m.SetSecrets(res)
        }
        return nil
    }
    res["templates"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSynchronizationTemplate() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SynchronizationTemplate, len(val))
            for i, v := range val {
                res[i] = *(v.(*SynchronizationTemplate))
            }
            m.SetTemplates(res)
        }
        return nil
    }
    return res
}
func (m *Synchronization) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *Synchronization) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetJobs() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetJobs()))
        for i, v := range m.GetJobs() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("jobs", cast)
        if err != nil {
            return err
        }
    }
    if m.GetSecrets() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetSecrets()))
        for i, v := range m.GetSecrets() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("secrets", cast)
        if err != nil {
            return err
        }
    }
    if m.GetTemplates() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetTemplates()))
        for i, v := range m.GetTemplates() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("templates", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetJobs sets the jobs property value. Performs synchronization by periodically running in the background, polling for changes in one directory, and pushing them to another directory.
func (m *Synchronization) SetJobs(value []SynchronizationJob)() {
    if m != nil {
        m.jobs = value
    }
}
// SetSecrets sets the secrets property value. Represents a collection of credentials to access provisioned cloud applications.
func (m *Synchronization) SetSecrets(value []SynchronizationSecretKeyStringValuePair)() {
    if m != nil {
        m.secrets = value
    }
}
// SetTemplates sets the templates property value. Pre-configured synchronization settings for a particular application.
func (m *Synchronization) SetTemplates(value []SynchronizationTemplate)() {
    if m != nil {
        m.templates = value
    }
}
