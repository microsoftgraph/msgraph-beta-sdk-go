package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AccessReviewSet provides operations to manage the identityGovernance singleton.
type AccessReviewSet struct {
    Entity
    // Represents an Azure AD access review decision on an instance of a review.
    decisions []AccessReviewInstanceDecisionItemable;
    // Represents the template and scheduling for an access review.
    definitions []AccessReviewScheduleDefinitionable;
    // Represents a collection of access review history data and the scopes used to collect that data.
    historyDefinitions []AccessReviewHistoryDefinitionable;
    // Resource that enables administrators to manage directory-level access review policies in their tenant.
    policy AccessReviewPolicyable;
}
// NewAccessReviewSet instantiates a new accessReviewSet and sets the default values.
func NewAccessReviewSet()(*AccessReviewSet) {
    m := &AccessReviewSet{
        Entity: *NewEntity(),
    }
    return m
}
// CreateAccessReviewSetFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAccessReviewSetFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewAccessReviewSet(), nil
}
// GetDecisions gets the decisions property value. Represents an Azure AD access review decision on an instance of a review.
func (m *AccessReviewSet) GetDecisions()([]AccessReviewInstanceDecisionItemable) {
    if m == nil {
        return nil
    } else {
        return m.decisions
    }
}
// GetDefinitions gets the definitions property value. Represents the template and scheduling for an access review.
func (m *AccessReviewSet) GetDefinitions()([]AccessReviewScheduleDefinitionable) {
    if m == nil {
        return nil
    } else {
        return m.definitions
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AccessReviewSet) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["decisions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAccessReviewInstanceDecisionItemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AccessReviewInstanceDecisionItemable, len(val))
            for i, v := range val {
                res[i] = v.(AccessReviewInstanceDecisionItemable)
            }
            m.SetDecisions(res)
        }
        return nil
    }
    res["definitions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAccessReviewScheduleDefinitionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AccessReviewScheduleDefinitionable, len(val))
            for i, v := range val {
                res[i] = v.(AccessReviewScheduleDefinitionable)
            }
            m.SetDefinitions(res)
        }
        return nil
    }
    res["historyDefinitions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAccessReviewHistoryDefinitionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AccessReviewHistoryDefinitionable, len(val))
            for i, v := range val {
                res[i] = v.(AccessReviewHistoryDefinitionable)
            }
            m.SetHistoryDefinitions(res)
        }
        return nil
    }
    res["policy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateAccessReviewPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPolicy(val.(AccessReviewPolicyable))
        }
        return nil
    }
    return res
}
// GetHistoryDefinitions gets the historyDefinitions property value. Represents a collection of access review history data and the scopes used to collect that data.
func (m *AccessReviewSet) GetHistoryDefinitions()([]AccessReviewHistoryDefinitionable) {
    if m == nil {
        return nil
    } else {
        return m.historyDefinitions
    }
}
// GetPolicy gets the policy property value. Resource that enables administrators to manage directory-level access review policies in their tenant.
func (m *AccessReviewSet) GetPolicy()(AccessReviewPolicyable) {
    if m == nil {
        return nil
    } else {
        return m.policy
    }
}
func (m *AccessReviewSet) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *AccessReviewSet) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetDecisions() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetDecisions()))
        for i, v := range m.GetDecisions() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("decisions", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDefinitions() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetDefinitions()))
        for i, v := range m.GetDefinitions() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("definitions", cast)
        if err != nil {
            return err
        }
    }
    if m.GetHistoryDefinitions() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetHistoryDefinitions()))
        for i, v := range m.GetHistoryDefinitions() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("historyDefinitions", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("policy", m.GetPolicy())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDecisions sets the decisions property value. Represents an Azure AD access review decision on an instance of a review.
func (m *AccessReviewSet) SetDecisions(value []AccessReviewInstanceDecisionItemable)() {
    if m != nil {
        m.decisions = value
    }
}
// SetDefinitions sets the definitions property value. Represents the template and scheduling for an access review.
func (m *AccessReviewSet) SetDefinitions(value []AccessReviewScheduleDefinitionable)() {
    if m != nil {
        m.definitions = value
    }
}
// SetHistoryDefinitions sets the historyDefinitions property value. Represents a collection of access review history data and the scopes used to collect that data.
func (m *AccessReviewSet) SetHistoryDefinitions(value []AccessReviewHistoryDefinitionable)() {
    if m != nil {
        m.historyDefinitions = value
    }
}
// SetPolicy sets the policy property value. Resource that enables administrators to manage directory-level access review policies in their tenant.
func (m *AccessReviewSet) SetPolicy(value AccessReviewPolicyable)() {
    if m != nil {
        m.policy = value
    }
}
