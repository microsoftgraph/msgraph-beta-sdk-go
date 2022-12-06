package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PlannerExternalTaskSource 
type PlannerExternalTaskSource struct {
    PlannerTaskCreation
    // The contextScenarioId property
    contextScenarioId *string
    // The displayLinkType property
    displayLinkType *PlannerExternalTaskSourceDisplayType
    // The displayNameSegments property
    displayNameSegments []string
    // The externalContextId property
    externalContextId *string
    // The externalObjectId property
    externalObjectId *string
    // The externalObjectVersion property
    externalObjectVersion *string
    // The webUrl property
    webUrl *string
}
// NewPlannerExternalTaskSource instantiates a new PlannerExternalTaskSource and sets the default values.
func NewPlannerExternalTaskSource()(*PlannerExternalTaskSource) {
    m := &PlannerExternalTaskSource{
        PlannerTaskCreation: *NewPlannerTaskCreation(),
    }
    odataTypeValue := "#microsoft.graph.plannerExternalTaskSource";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreatePlannerExternalTaskSourceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePlannerExternalTaskSourceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPlannerExternalTaskSource(), nil
}
// GetContextScenarioId gets the contextScenarioId property value. The contextScenarioId property
func (m *PlannerExternalTaskSource) GetContextScenarioId()(*string) {
    return m.contextScenarioId
}
// GetDisplayLinkType gets the displayLinkType property value. The displayLinkType property
func (m *PlannerExternalTaskSource) GetDisplayLinkType()(*PlannerExternalTaskSourceDisplayType) {
    return m.displayLinkType
}
// GetDisplayNameSegments gets the displayNameSegments property value. The displayNameSegments property
func (m *PlannerExternalTaskSource) GetDisplayNameSegments()([]string) {
    return m.displayNameSegments
}
// GetExternalContextId gets the externalContextId property value. The externalContextId property
func (m *PlannerExternalTaskSource) GetExternalContextId()(*string) {
    return m.externalContextId
}
// GetExternalObjectId gets the externalObjectId property value. The externalObjectId property
func (m *PlannerExternalTaskSource) GetExternalObjectId()(*string) {
    return m.externalObjectId
}
// GetExternalObjectVersion gets the externalObjectVersion property value. The externalObjectVersion property
func (m *PlannerExternalTaskSource) GetExternalObjectVersion()(*string) {
    return m.externalObjectVersion
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PlannerExternalTaskSource) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.PlannerTaskCreation.GetFieldDeserializers()
    res["contextScenarioId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetContextScenarioId)
    res["displayLinkType"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParsePlannerExternalTaskSourceDisplayType , m.SetDisplayLinkType)
    res["displayNameSegments"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetDisplayNameSegments)
    res["externalContextId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetExternalContextId)
    res["externalObjectId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetExternalObjectId)
    res["externalObjectVersion"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetExternalObjectVersion)
    res["webUrl"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetWebUrl)
    return res
}
// GetWebUrl gets the webUrl property value. The webUrl property
func (m *PlannerExternalTaskSource) GetWebUrl()(*string) {
    return m.webUrl
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
// SetContextScenarioId sets the contextScenarioId property value. The contextScenarioId property
func (m *PlannerExternalTaskSource) SetContextScenarioId(value *string)() {
    m.contextScenarioId = value
}
// SetDisplayLinkType sets the displayLinkType property value. The displayLinkType property
func (m *PlannerExternalTaskSource) SetDisplayLinkType(value *PlannerExternalTaskSourceDisplayType)() {
    m.displayLinkType = value
}
// SetDisplayNameSegments sets the displayNameSegments property value. The displayNameSegments property
func (m *PlannerExternalTaskSource) SetDisplayNameSegments(value []string)() {
    m.displayNameSegments = value
}
// SetExternalContextId sets the externalContextId property value. The externalContextId property
func (m *PlannerExternalTaskSource) SetExternalContextId(value *string)() {
    m.externalContextId = value
}
// SetExternalObjectId sets the externalObjectId property value. The externalObjectId property
func (m *PlannerExternalTaskSource) SetExternalObjectId(value *string)() {
    m.externalObjectId = value
}
// SetExternalObjectVersion sets the externalObjectVersion property value. The externalObjectVersion property
func (m *PlannerExternalTaskSource) SetExternalObjectVersion(value *string)() {
    m.externalObjectVersion = value
}
// SetWebUrl sets the webUrl property value. The webUrl property
func (m *PlannerExternalTaskSource) SetWebUrl(value *string)() {
    m.webUrl = value
}
