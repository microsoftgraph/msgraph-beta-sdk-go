package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type PlannerExternalBucketSource struct {
    PlannerBucketCreation
}
// NewPlannerExternalBucketSource instantiates a new PlannerExternalBucketSource and sets the default values.
func NewPlannerExternalBucketSource()(*PlannerExternalBucketSource) {
    m := &PlannerExternalBucketSource{
        PlannerBucketCreation: *NewPlannerBucketCreation(),
    }
    odataTypeValue := "#microsoft.graph.plannerExternalBucketSource"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreatePlannerExternalBucketSourceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreatePlannerExternalBucketSourceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPlannerExternalBucketSource(), nil
}
// GetContextScenarioId gets the contextScenarioId property value. Nullable. An identifier for the scenario associated with this external source. This should be in reverse DNS format. For example, Contoso company owned application for customer support would have a value like 'com.constoso.customerSupport'.
// returns a *string when successful
func (m *PlannerExternalBucketSource) GetContextScenarioId()(*string) {
    val, err := m.GetBackingStore().Get("contextScenarioId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetExternalContextId gets the externalContextId property value. Nullable. The ID of the external entity's containing entity or context.
// returns a *string when successful
func (m *PlannerExternalBucketSource) GetExternalContextId()(*string) {
    val, err := m.GetBackingStore().Get("externalContextId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetExternalObjectId gets the externalObjectId property value. Nullable. The ID of the entity that an external service associates with a bucket.
// returns a *string when successful
func (m *PlannerExternalBucketSource) GetExternalObjectId()(*string) {
    val, err := m.GetBackingStore().Get("externalObjectId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *PlannerExternalBucketSource) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.PlannerBucketCreation.GetFieldDeserializers()
    res["contextScenarioId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContextScenarioId(val)
        }
        return nil
    }
    res["externalContextId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExternalContextId(val)
        }
        return nil
    }
    res["externalObjectId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExternalObjectId(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *PlannerExternalBucketSource) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.PlannerBucketCreation.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("contextScenarioId", m.GetContextScenarioId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("externalContextId", m.GetExternalContextId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("externalObjectId", m.GetExternalObjectId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetContextScenarioId sets the contextScenarioId property value. Nullable. An identifier for the scenario associated with this external source. This should be in reverse DNS format. For example, Contoso company owned application for customer support would have a value like 'com.constoso.customerSupport'.
func (m *PlannerExternalBucketSource) SetContextScenarioId(value *string)() {
    err := m.GetBackingStore().Set("contextScenarioId", value)
    if err != nil {
        panic(err)
    }
}
// SetExternalContextId sets the externalContextId property value. Nullable. The ID of the external entity's containing entity or context.
func (m *PlannerExternalBucketSource) SetExternalContextId(value *string)() {
    err := m.GetBackingStore().Set("externalContextId", value)
    if err != nil {
        panic(err)
    }
}
// SetExternalObjectId sets the externalObjectId property value. Nullable. The ID of the entity that an external service associates with a bucket.
func (m *PlannerExternalBucketSource) SetExternalObjectId(value *string)() {
    err := m.GetBackingStore().Set("externalObjectId", value)
    if err != nil {
        panic(err)
    }
}
type PlannerExternalBucketSourceable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    PlannerBucketCreationable
    GetContextScenarioId()(*string)
    GetExternalContextId()(*string)
    GetExternalObjectId()(*string)
    SetContextScenarioId(value *string)()
    SetExternalContextId(value *string)()
    SetExternalObjectId(value *string)()
}
