package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SynchronizationJobApplicationParameters 
type SynchronizationJobApplicationParameters struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The OdataType property
    odataType *string
    // The identifier of the synchronizationRule to be applied. This rule ID is defined in the schema for a given synchronization job or template.
    ruleId *string
    // The identifiers of one or more objects to which a synchronizationJob is to be applied.
    subjects []SynchronizationJobSubjectable
}
// NewSynchronizationJobApplicationParameters instantiates a new synchronizationJobApplicationParameters and sets the default values.
func NewSynchronizationJobApplicationParameters()(*SynchronizationJobApplicationParameters) {
    m := &SynchronizationJobApplicationParameters{
    }
    m.SetAdditionalData(make(map[string]any));
    return m
}
// CreateSynchronizationJobApplicationParametersFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSynchronizationJobApplicationParametersFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSynchronizationJobApplicationParameters(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SynchronizationJobApplicationParameters) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SynchronizationJobApplicationParameters) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    res["ruleId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRuleId(val)
        }
        return nil
    }
    res["subjects"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *SynchronizationJobApplicationParameters) GetOdataType()(*string) {
    return m.odataType
}
// GetRuleId gets the ruleId property value. The identifier of the synchronizationRule to be applied. This rule ID is defined in the schema for a given synchronization job or template.
func (m *SynchronizationJobApplicationParameters) GetRuleId()(*string) {
    return m.ruleId
}
// GetSubjects gets the subjects property value. The identifiers of one or more objects to which a synchronizationJob is to be applied.
func (m *SynchronizationJobApplicationParameters) GetSubjects()([]SynchronizationJobSubjectable) {
    return m.subjects
}
// Serialize serializes information the current object
func (m *SynchronizationJobApplicationParameters) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("ruleId", m.GetRuleId())
        if err != nil {
            return err
        }
    }
    if m.GetSubjects() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSubjects()))
        for i, v := range m.GetSubjects() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
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
func (m *SynchronizationJobApplicationParameters) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *SynchronizationJobApplicationParameters) SetOdataType(value *string)() {
    m.odataType = value
}
// SetRuleId sets the ruleId property value. The identifier of the synchronizationRule to be applied. This rule ID is defined in the schema for a given synchronization job or template.
func (m *SynchronizationJobApplicationParameters) SetRuleId(value *string)() {
    m.ruleId = value
}
// SetSubjects sets the subjects property value. The identifiers of one or more objects to which a synchronizationJob is to be applied.
func (m *SynchronizationJobApplicationParameters) SetSubjects(value []SynchronizationJobSubjectable)() {
    m.subjects = value
}
