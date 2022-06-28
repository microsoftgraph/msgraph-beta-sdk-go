package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PlannerDelta provides operations to manage the collection of accessReview entities.
type PlannerDelta struct {
    Entity
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
}
// NewPlannerDelta instantiates a new plannerDelta and sets the default values.
func NewPlannerDelta()(*PlannerDelta) {
    m := &PlannerDelta{
        Entity: *NewEntity(),
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreatePlannerDeltaFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePlannerDeltaFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                mappingStr := *mappingValue
                switch mappingStr {
                    case "#microsoft.graph.plannerAssignedToTaskBoardTaskFormat":
                        return NewPlannerAssignedToTaskBoardTaskFormat(), nil
                    case "#microsoft.graph.plannerBucket":
                        return NewPlannerBucket(), nil
                    case "#microsoft.graph.plannerBucketTaskBoardTaskFormat":
                        return NewPlannerBucketTaskBoardTaskFormat(), nil
                    case "#microsoft.graph.plannerPlan":
                        return NewPlannerPlan(), nil
                    case "#microsoft.graph.plannerPlanDetails":
                        return NewPlannerPlanDetails(), nil
                    case "#microsoft.graph.plannerProgressTaskBoardTaskFormat":
                        return NewPlannerProgressTaskBoardTaskFormat(), nil
                    case "#microsoft.graph.plannerTask":
                        return NewPlannerTask(), nil
                    case "#microsoft.graph.plannerTaskDetails":
                        return NewPlannerTaskDetails(), nil
                    case "#microsoft.graph.plannerUser":
                        return NewPlannerUser(), nil
                }
            }
        }
    }
    return NewPlannerDelta(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PlannerDelta) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PlannerDelta) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *PlannerDelta) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PlannerDelta) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
