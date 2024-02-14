package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type PlannerSharedWithContainer struct {
    PlannerPlanContainer
}
// NewPlannerSharedWithContainer instantiates a new PlannerSharedWithContainer and sets the default values.
func NewPlannerSharedWithContainer()(*PlannerSharedWithContainer) {
    m := &PlannerSharedWithContainer{
        PlannerPlanContainer: *NewPlannerPlanContainer(),
    }
    odataTypeValue := "#microsoft.graph.plannerSharedWithContainer"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreatePlannerSharedWithContainerFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreatePlannerSharedWithContainerFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPlannerSharedWithContainer(), nil
}
// GetAccessLevel gets the accessLevel property value. The accessLevel property
// returns a *PlannerPlanAccessLevel when successful
func (m *PlannerSharedWithContainer) GetAccessLevel()(*PlannerPlanAccessLevel) {
    val, err := m.GetBackingStore().Get("accessLevel")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*PlannerPlanAccessLevel)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *PlannerSharedWithContainer) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.PlannerPlanContainer.GetFieldDeserializers()
    res["accessLevel"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParsePlannerPlanAccessLevel)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccessLevel(val.(*PlannerPlanAccessLevel))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *PlannerSharedWithContainer) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.PlannerPlanContainer.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAccessLevel() != nil {
        cast := (*m.GetAccessLevel()).String()
        err = writer.WriteStringValue("accessLevel", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAccessLevel sets the accessLevel property value. The accessLevel property
func (m *PlannerSharedWithContainer) SetAccessLevel(value *PlannerPlanAccessLevel)() {
    err := m.GetBackingStore().Set("accessLevel", value)
    if err != nil {
        panic(err)
    }
}
type PlannerSharedWithContainerable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    PlannerPlanContainerable
    GetAccessLevel()(*PlannerPlanAccessLevel)
    SetAccessLevel(value *PlannerPlanAccessLevel)()
}
