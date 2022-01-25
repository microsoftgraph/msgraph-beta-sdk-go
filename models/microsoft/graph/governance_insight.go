package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// GovernanceInsight 
type GovernanceInsight struct {
    Entity
    // Indicates when the insight was created.
    insightCreatedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
}
// NewGovernanceInsight instantiates a new governanceInsight and sets the default values.
func NewGovernanceInsight()(*GovernanceInsight) {
    m := &GovernanceInsight{
        Entity: *NewEntity(),
    }
    return m
}
// GetInsightCreatedDateTime gets the insightCreatedDateTime property value. Indicates when the insight was created.
func (m *GovernanceInsight) GetInsightCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.insightCreatedDateTime
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GovernanceInsight) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["insightCreatedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInsightCreatedDateTime(val)
        }
        return nil
    }
    return res
}
func (m *GovernanceInsight) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *GovernanceInsight) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("insightCreatedDateTime", m.GetInsightCreatedDateTime())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetInsightCreatedDateTime sets the insightCreatedDateTime property value. Indicates when the insight was created.
func (m *GovernanceInsight) SetInsightCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.insightCreatedDateTime = value
    }
}
