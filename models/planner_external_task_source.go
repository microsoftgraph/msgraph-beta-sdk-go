package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PlannerExternalTaskSource 
type PlannerExternalTaskSource struct {
    PlannerTaskCreation
}
// NewPlannerExternalTaskSource instantiates a new PlannerExternalTaskSource and sets the default values.
func NewPlannerExternalTaskSource()(*PlannerExternalTaskSource) {
    m := &PlannerExternalTaskSource{
        PlannerTaskCreation: *NewPlannerTaskCreation(),
    }
    odataTypeValue := "#microsoft.graph.plannerExternalTaskSource"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreatePlannerExternalTaskSourceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePlannerExternalTaskSourceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPlannerExternalTaskSource(), nil
}
// GetContextScenarioId gets the contextScenarioId property value. Nullable. An identifier for the scenario associated with this external source. This should be in reverse DNS format. For example, Contoso company owned application for customer support would have a value like 'com.constoso.customerSupport'.
func (m *PlannerExternalTaskSource) GetContextScenarioId()(*string) {
    val, err := m.GetBackingStore().Get("contextScenarioId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDisplayLinkType gets the displayLinkType property value. Specifies how an application should display the link to the associated plannerExternalTaskSource. The possible values are: none, default.
func (m *PlannerExternalTaskSource) GetDisplayLinkType()(*PlannerExternalTaskSourceDisplayType) {
    val, err := m.GetBackingStore().Get("displayLinkType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*PlannerExternalTaskSourceDisplayType)
    }
    return nil
}
// GetDisplayNameSegments gets the displayNameSegments property value. The segments of the name of the external experience. Segments represent a hierarchical structure that allows other apps to display the relationship.
func (m *PlannerExternalTaskSource) GetDisplayNameSegments()([]string) {
    val, err := m.GetBackingStore().Get("displayNameSegments")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetExternalContextId gets the externalContextId property value. Nullable. The id of the external entity's containing entity or context.
func (m *PlannerExternalTaskSource) GetExternalContextId()(*string) {
    val, err := m.GetBackingStore().Get("externalContextId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetExternalObjectId gets the externalObjectId property value. Nullable. The id of the entity that an external service associates with a task.
func (m *PlannerExternalTaskSource) GetExternalObjectId()(*string) {
    val, err := m.GetBackingStore().Get("externalObjectId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetExternalObjectVersion gets the externalObjectVersion property value. Nullable. The external Item Version for the object specified by the externalObjectId.
func (m *PlannerExternalTaskSource) GetExternalObjectVersion()(*string) {
    val, err := m.GetBackingStore().Get("externalObjectVersion")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PlannerExternalTaskSource) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.PlannerTaskCreation.GetFieldDeserializers()
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
    res["displayLinkType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParsePlannerExternalTaskSourceDisplayType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayLinkType(val.(*PlannerExternalTaskSourceDisplayType))
        }
        return nil
    }
    res["displayNameSegments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetDisplayNameSegments(res)
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
    res["externalObjectVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExternalObjectVersion(val)
        }
        return nil
    }
    res["webUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWebUrl(val)
        }
        return nil
    }
    return res
}
// GetWebUrl gets the webUrl property value. Nullable. URL of the user experience represented by the associated plannerExternalTaskSource.
func (m *PlannerExternalTaskSource) GetWebUrl()(*string) {
    val, err := m.GetBackingStore().Get("webUrl")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *PlannerExternalTaskSource) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.PlannerTaskCreation.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("contextScenarioId", m.GetContextScenarioId())
        if err != nil {
            return err
        }
    }
    if m.GetDisplayLinkType() != nil {
        cast := (*m.GetDisplayLinkType()).String()
        err = writer.WriteStringValue("displayLinkType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetDisplayNameSegments() != nil {
        err = writer.WriteCollectionOfStringValues("displayNameSegments", m.GetDisplayNameSegments())
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
    {
        err = writer.WriteStringValue("externalObjectVersion", m.GetExternalObjectVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("webUrl", m.GetWebUrl())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetContextScenarioId sets the contextScenarioId property value. Nullable. An identifier for the scenario associated with this external source. This should be in reverse DNS format. For example, Contoso company owned application for customer support would have a value like 'com.constoso.customerSupport'.
func (m *PlannerExternalTaskSource) SetContextScenarioId(value *string)() {
    err := m.GetBackingStore().Set("contextScenarioId", value)
    if err != nil {
        panic(err)
    }
}
// SetDisplayLinkType sets the displayLinkType property value. Specifies how an application should display the link to the associated plannerExternalTaskSource. The possible values are: none, default.
func (m *PlannerExternalTaskSource) SetDisplayLinkType(value *PlannerExternalTaskSourceDisplayType)() {
    err := m.GetBackingStore().Set("displayLinkType", value)
    if err != nil {
        panic(err)
    }
}
// SetDisplayNameSegments sets the displayNameSegments property value. The segments of the name of the external experience. Segments represent a hierarchical structure that allows other apps to display the relationship.
func (m *PlannerExternalTaskSource) SetDisplayNameSegments(value []string)() {
    err := m.GetBackingStore().Set("displayNameSegments", value)
    if err != nil {
        panic(err)
    }
}
// SetExternalContextId sets the externalContextId property value. Nullable. The id of the external entity's containing entity or context.
func (m *PlannerExternalTaskSource) SetExternalContextId(value *string)() {
    err := m.GetBackingStore().Set("externalContextId", value)
    if err != nil {
        panic(err)
    }
}
// SetExternalObjectId sets the externalObjectId property value. Nullable. The id of the entity that an external service associates with a task.
func (m *PlannerExternalTaskSource) SetExternalObjectId(value *string)() {
    err := m.GetBackingStore().Set("externalObjectId", value)
    if err != nil {
        panic(err)
    }
}
// SetExternalObjectVersion sets the externalObjectVersion property value. Nullable. The external Item Version for the object specified by the externalObjectId.
func (m *PlannerExternalTaskSource) SetExternalObjectVersion(value *string)() {
    err := m.GetBackingStore().Set("externalObjectVersion", value)
    if err != nil {
        panic(err)
    }
}
// SetWebUrl sets the webUrl property value. Nullable. URL of the user experience represented by the associated plannerExternalTaskSource.
func (m *PlannerExternalTaskSource) SetWebUrl(value *string)() {
    err := m.GetBackingStore().Set("webUrl", value)
    if err != nil {
        panic(err)
    }
}
// PlannerExternalTaskSourceable 
type PlannerExternalTaskSourceable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    PlannerTaskCreationable
    GetContextScenarioId()(*string)
    GetDisplayLinkType()(*PlannerExternalTaskSourceDisplayType)
    GetDisplayNameSegments()([]string)
    GetExternalContextId()(*string)
    GetExternalObjectId()(*string)
    GetExternalObjectVersion()(*string)
    GetWebUrl()(*string)
    SetContextScenarioId(value *string)()
    SetDisplayLinkType(value *PlannerExternalTaskSourceDisplayType)()
    SetDisplayNameSegments(value []string)()
    SetExternalContextId(value *string)()
    SetExternalObjectId(value *string)()
    SetExternalObjectVersion(value *string)()
    SetWebUrl(value *string)()
}
