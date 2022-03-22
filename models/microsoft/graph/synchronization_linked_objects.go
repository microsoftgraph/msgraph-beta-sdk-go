package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// SynchronizationLinkedObjects 
type SynchronizationLinkedObjects struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    manager SynchronizationJobSubjectable;
    // 
    members []SynchronizationJobSubjectable;
    // 
    owners []SynchronizationJobSubjectable;
}
// NewSynchronizationLinkedObjects instantiates a new synchronizationLinkedObjects and sets the default values.
func NewSynchronizationLinkedObjects()(*SynchronizationLinkedObjects) {
    m := &SynchronizationLinkedObjects{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateSynchronizationLinkedObjectsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSynchronizationLinkedObjectsFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewSynchronizationLinkedObjects(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SynchronizationLinkedObjects) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SynchronizationLinkedObjects) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["manager"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateSynchronizationJobSubjectFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManager(val.(SynchronizationJobSubjectable))
        }
        return nil
    }
    res["members"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSynchronizationJobSubjectFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SynchronizationJobSubjectable, len(val))
            for i, v := range val {
                res[i] = v.(SynchronizationJobSubjectable)
            }
            m.SetMembers(res)
        }
        return nil
    }
    res["owners"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSynchronizationJobSubjectFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SynchronizationJobSubjectable, len(val))
            for i, v := range val {
                res[i] = v.(SynchronizationJobSubjectable)
            }
            m.SetOwners(res)
        }
        return nil
    }
    return res
}
// GetManager gets the manager property value. 
func (m *SynchronizationLinkedObjects) GetManager()(SynchronizationJobSubjectable) {
    if m == nil {
        return nil
    } else {
        return m.manager
    }
}
// GetMembers gets the members property value. 
func (m *SynchronizationLinkedObjects) GetMembers()([]SynchronizationJobSubjectable) {
    if m == nil {
        return nil
    } else {
        return m.members
    }
}
// GetOwners gets the owners property value. 
func (m *SynchronizationLinkedObjects) GetOwners()([]SynchronizationJobSubjectable) {
    if m == nil {
        return nil
    } else {
        return m.owners
    }
}
// Serialize serializes information the current object
func (m *SynchronizationLinkedObjects) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("manager", m.GetManager())
        if err != nil {
            return err
        }
    }
    if m.GetMembers() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetMembers()))
        for i, v := range m.GetMembers() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("members", cast)
        if err != nil {
            return err
        }
    }
    if m.GetOwners() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetOwners()))
        for i, v := range m.GetOwners() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("owners", cast)
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SynchronizationLinkedObjects) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetManager sets the manager property value. 
func (m *SynchronizationLinkedObjects) SetManager(value SynchronizationJobSubjectable)() {
    if m != nil {
        m.manager = value
    }
}
// SetMembers sets the members property value. 
func (m *SynchronizationLinkedObjects) SetMembers(value []SynchronizationJobSubjectable)() {
    if m != nil {
        m.members = value
    }
}
// SetOwners sets the owners property value. 
func (m *SynchronizationLinkedObjects) SetOwners(value []SynchronizationJobSubjectable)() {
    if m != nil {
        m.owners = value
    }
}
