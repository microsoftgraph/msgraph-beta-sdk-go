package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type Synchronization struct {
    Entity
    // 
    jobs []SynchronizationJob;
    // 
    secrets []SynchronizationSecretKeyStringValuePair;
    // 
    templates []SynchronizationTemplate;
}
// Instantiates a new synchronization and sets the default values.
func NewSynchronization()(*Synchronization) {
    m := &Synchronization{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the jobs property value. 
func (m *Synchronization) GetJobs()([]SynchronizationJob) {
    if m == nil {
        return nil
    } else {
        return m.jobs
    }
}
// Gets the secrets property value. 
func (m *Synchronization) GetSecrets()([]SynchronizationSecretKeyStringValuePair) {
    if m == nil {
        return nil
    } else {
        return m.secrets
    }
}
// Gets the templates property value. 
func (m *Synchronization) GetTemplates()([]SynchronizationTemplate) {
    if m == nil {
        return nil
    } else {
        return m.templates
    }
}
// The deserialization information for the current model
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
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *Synchronization) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
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
    {
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
    {
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
// Sets the jobs property value. 
// Parameters:
//  - value : Value to set for the jobs property.
func (m *Synchronization) SetJobs(value []SynchronizationJob)() {
    m.jobs = value
}
// Sets the secrets property value. 
// Parameters:
//  - value : Value to set for the secrets property.
func (m *Synchronization) SetSecrets(value []SynchronizationSecretKeyStringValuePair)() {
    m.secrets = value
}
// Sets the templates property value. 
// Parameters:
//  - value : Value to set for the templates property.
func (m *Synchronization) SetTemplates(value []SynchronizationTemplate)() {
    m.templates = value
}
