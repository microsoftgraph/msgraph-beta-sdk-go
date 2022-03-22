package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DlpPoliciesJobResult 
type DlpPoliciesJobResult struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    auditCorrelationId *string;
    // 
    evaluationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // 
    matchingRules []MatchingDlpRuleable;
}
// NewDlpPoliciesJobResult instantiates a new dlpPoliciesJobResult and sets the default values.
func NewDlpPoliciesJobResult()(*DlpPoliciesJobResult) {
    m := &DlpPoliciesJobResult{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDlpPoliciesJobResultFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDlpPoliciesJobResultFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewDlpPoliciesJobResult(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DlpPoliciesJobResult) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAuditCorrelationId gets the auditCorrelationId property value. 
func (m *DlpPoliciesJobResult) GetAuditCorrelationId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.auditCorrelationId
    }
}
// GetEvaluationDateTime gets the evaluationDateTime property value. 
func (m *DlpPoliciesJobResult) GetEvaluationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.evaluationDateTime
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DlpPoliciesJobResult) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["auditCorrelationId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuditCorrelationId(val)
        }
        return nil
    }
    res["evaluationDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEvaluationDateTime(val)
        }
        return nil
    }
    res["matchingRules"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMatchingDlpRuleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MatchingDlpRuleable, len(val))
            for i, v := range val {
                res[i] = v.(MatchingDlpRuleable)
            }
            m.SetMatchingRules(res)
        }
        return nil
    }
    return res
}
// GetMatchingRules gets the matchingRules property value. 
func (m *DlpPoliciesJobResult) GetMatchingRules()([]MatchingDlpRuleable) {
    if m == nil {
        return nil
    } else {
        return m.matchingRules
    }
}
// Serialize serializes information the current object
func (m *DlpPoliciesJobResult) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("auditCorrelationId", m.GetAuditCorrelationId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("evaluationDateTime", m.GetEvaluationDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetMatchingRules() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetMatchingRules()))
        for i, v := range m.GetMatchingRules() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("matchingRules", cast)
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
func (m *DlpPoliciesJobResult) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAuditCorrelationId sets the auditCorrelationId property value. 
func (m *DlpPoliciesJobResult) SetAuditCorrelationId(value *string)() {
    if m != nil {
        m.auditCorrelationId = value
    }
}
// SetEvaluationDateTime sets the evaluationDateTime property value. 
func (m *DlpPoliciesJobResult) SetEvaluationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.evaluationDateTime = value
    }
}
// SetMatchingRules sets the matchingRules property value. 
func (m *DlpPoliciesJobResult) SetMatchingRules(value []MatchingDlpRuleable)() {
    if m != nil {
        m.matchingRules = value
    }
}
