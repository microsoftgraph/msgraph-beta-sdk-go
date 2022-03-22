package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// SynchronizationJobApplicationParameters 
type SynchronizationJobApplicationParameters struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The identifier of a the synchronizationRule to be applied.
    ruleId *string;
    // The identifiers of one or more objects to which a synchronizationJob is to be applied.
    subjects []SynchronizationJobSubjectable;
}
// NewSynchronizationJobApplicationParameters instantiates a new synchronizationJobApplicationParameters and sets the default values.
func NewSynchronizationJobApplicationParameters()(*SynchronizationJobApplicationParameters) {
    m := &SynchronizationJobApplicationParameters{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateSynchronizationJobApplicationParametersFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSynchronizationJobApplicationParametersFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewSynchronizationJobApplicationParameters(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SynchronizationJobApplicationParameters) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SynchronizationJobApplicationParameters) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["ruleId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRuleId(val)
        }
        return nil
    }
    res["subjects"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSynchronizationJobSubjectFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SynchronizationJobSubjectable, len(val))
            for i, v := range val {
                res[i] = v.(SynchronizationJobSubjectable)
            }
            m.SetSubjects(res)
        }
        return nil
    }
    return res
}
// GetRuleId gets the ruleId property value. The identifier of a the synchronizationRule to be applied.
func (m *SynchronizationJobApplicationParameters) GetRuleId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.ruleId
    }
}
// GetSubjects gets the subjects property value. The identifiers of one or more objects to which a synchronizationJob is to be applied.
func (m *SynchronizationJobApplicationParameters) GetSubjects()([]SynchronizationJobSubjectable) {
    if m == nil {
        return nil
    } else {
        return m.subjects
    }
}
// Serialize serializes information the current object
func (m *SynchronizationJobApplicationParameters) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("ruleId", m.GetRuleId())
        if err != nil {
            return err
        }
    }
    if m.GetSubjects() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetSubjects()))
        for i, v := range m.GetSubjects() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("subjects", cast)
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
func (m *SynchronizationJobApplicationParameters) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetRuleId sets the ruleId property value. The identifier of a the synchronizationRule to be applied.
func (m *SynchronizationJobApplicationParameters) SetRuleId(value *string)() {
    if m != nil {
        m.ruleId = value
    }
}
// SetSubjects sets the subjects property value. The identifiers of one or more objects to which a synchronizationJob is to be applied.
func (m *SynchronizationJobApplicationParameters) SetSubjects(value []SynchronizationJobSubjectable)() {
    if m != nil {
        m.subjects = value
    }
}
