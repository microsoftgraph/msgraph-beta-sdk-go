package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// PlannerBucketTaskBoardTaskFormat 
type PlannerBucketTaskBoardTaskFormat struct {
    PlannerDelta
    // Hint used to order tasks in the Bucket view of the Task Board. The format is defined as outlined here.
    orderHint *string;
}
// NewPlannerBucketTaskBoardTaskFormat instantiates a new plannerBucketTaskBoardTaskFormat and sets the default values.
func NewPlannerBucketTaskBoardTaskFormat()(*PlannerBucketTaskBoardTaskFormat) {
    m := &PlannerBucketTaskBoardTaskFormat{
        PlannerDelta: *NewPlannerDelta(),
    }
    return m
}
// GetOrderHint gets the orderHint property value. Hint used to order tasks in the Bucket view of the Task Board. The format is defined as outlined here.
func (m *PlannerBucketTaskBoardTaskFormat) GetOrderHint()(*string) {
    if m == nil {
        return nil
    } else {
        return m.orderHint
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PlannerBucketTaskBoardTaskFormat) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.PlannerDelta.GetFieldDeserializers()
    res["orderHint"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOrderHint(val)
        }
        return nil
    }
    return res
}
func (m *PlannerBucketTaskBoardTaskFormat) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *PlannerBucketTaskBoardTaskFormat) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.PlannerDelta.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("orderHint", m.GetOrderHint())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetOrderHint sets the orderHint property value. Hint used to order tasks in the Bucket view of the Task Board. The format is defined as outlined here.
func (m *PlannerBucketTaskBoardTaskFormat) SetOrderHint(value *string)() {
    if m != nil {
        m.orderHint = value
    }
}
