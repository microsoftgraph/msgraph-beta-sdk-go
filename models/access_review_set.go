package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AccessReviewSet 
type AccessReviewSet struct {
    Entity
}
// NewAccessReviewSet instantiates a new accessReviewSet and sets the default values.
func NewAccessReviewSet()(*AccessReviewSet) {
    m := &AccessReviewSet{
        Entity: *NewEntity(),
    }
    return m
}
// CreateAccessReviewSetFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAccessReviewSetFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAccessReviewSet(), nil
}
// GetDecisions gets the decisions property value. Represents a Microsoft Entra access review decision on an instance of a review.
func (m *AccessReviewSet) GetDecisions()([]AccessReviewInstanceDecisionItemable) {
    val, err := m.GetBackingStore().Get("decisions")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]AccessReviewInstanceDecisionItemable)
    }
    return nil
}
// GetDefinitions gets the definitions property value. Represents the template and scheduling for an access review.
func (m *AccessReviewSet) GetDefinitions()([]AccessReviewScheduleDefinitionable) {
    val, err := m.GetBackingStore().Get("definitions")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]AccessReviewScheduleDefinitionable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AccessReviewSet) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["decisions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAccessReviewInstanceDecisionItemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AccessReviewInstanceDecisionItemable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(AccessReviewInstanceDecisionItemable)
                }
            }
            m.SetDecisions(res)
        }
        return nil
    }
    res["definitions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAccessReviewScheduleDefinitionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AccessReviewScheduleDefinitionable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(AccessReviewScheduleDefinitionable)
                }
            }
            m.SetDefinitions(res)
        }
        return nil
    }
    res["historyDefinitions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAccessReviewHistoryDefinitionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AccessReviewHistoryDefinitionable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(AccessReviewHistoryDefinitionable)
                }
            }
            m.SetHistoryDefinitions(res)
        }
        return nil
    }
    res["policy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
    val, err := m.GetBackingStore().Get("historyDefinitions")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]AccessReviewHistoryDefinitionable)
    }
    return nil
}
// GetPolicy gets the policy property value. Resource that enables administrators to manage directory-level access review policies in their tenant.
func (m *AccessReviewSet) GetPolicy()(AccessReviewPolicyable) {
    val, err := m.GetBackingStore().Get("policy")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(AccessReviewPolicyable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *AccessReviewSet) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetDecisions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDecisions()))
        for i, v := range m.GetDecisions() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("decisions", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDefinitions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDefinitions()))
        for i, v := range m.GetDefinitions() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("definitions", cast)
        if err != nil {
            return err
        }
    }
    if m.GetHistoryDefinitions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetHistoryDefinitions()))
        for i, v := range m.GetHistoryDefinitions() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
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
// SetDecisions sets the decisions property value. Represents a Microsoft Entra access review decision on an instance of a review.
func (m *AccessReviewSet) SetDecisions(value []AccessReviewInstanceDecisionItemable)() {
    err := m.GetBackingStore().Set("decisions", value)
    if err != nil {
        panic(err)
    }
}
// SetDefinitions sets the definitions property value. Represents the template and scheduling for an access review.
func (m *AccessReviewSet) SetDefinitions(value []AccessReviewScheduleDefinitionable)() {
    err := m.GetBackingStore().Set("definitions", value)
    if err != nil {
        panic(err)
    }
}
// SetHistoryDefinitions sets the historyDefinitions property value. Represents a collection of access review history data and the scopes used to collect that data.
func (m *AccessReviewSet) SetHistoryDefinitions(value []AccessReviewHistoryDefinitionable)() {
    err := m.GetBackingStore().Set("historyDefinitions", value)
    if err != nil {
        panic(err)
    }
}
// SetPolicy sets the policy property value. Resource that enables administrators to manage directory-level access review policies in their tenant.
func (m *AccessReviewSet) SetPolicy(value AccessReviewPolicyable)() {
    err := m.GetBackingStore().Set("policy", value)
    if err != nil {
        panic(err)
    }
}
// AccessReviewSetable 
type AccessReviewSetable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDecisions()([]AccessReviewInstanceDecisionItemable)
    GetDefinitions()([]AccessReviewScheduleDefinitionable)
    GetHistoryDefinitions()([]AccessReviewHistoryDefinitionable)
    GetPolicy()(AccessReviewPolicyable)
    SetDecisions(value []AccessReviewInstanceDecisionItemable)()
    SetDefinitions(value []AccessReviewScheduleDefinitionable)()
    SetHistoryDefinitions(value []AccessReviewHistoryDefinitionable)()
    SetPolicy(value AccessReviewPolicyable)()
}
