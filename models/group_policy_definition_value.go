package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GroupPolicyDefinitionValue the definition value entity stores the value for a single group policy definition.
type GroupPolicyDefinitionValue struct {
    Entity
    // Group Policy Configuration Type
    configurationType *GroupPolicyConfigurationType
    // The date and time the object was created.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The associated group policy definition with the value.
    definition GroupPolicyDefinitionable
    // Enables or disables the associated group policy definition.
    enabled *bool
    // The date and time the entity was last modified.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The associated group policy presentation values with the definition value.
    presentationValues []GroupPolicyPresentationValueable
}
// NewGroupPolicyDefinitionValue instantiates a new groupPolicyDefinitionValue and sets the default values.
func NewGroupPolicyDefinitionValue()(*GroupPolicyDefinitionValue) {
    m := &GroupPolicyDefinitionValue{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.groupPolicyDefinitionValue";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateGroupPolicyDefinitionValueFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGroupPolicyDefinitionValueFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGroupPolicyDefinitionValue(), nil
}
// GetConfigurationType gets the configurationType property value. Group Policy Configuration Type
func (m *GroupPolicyDefinitionValue) GetConfigurationType()(*GroupPolicyConfigurationType) {
    return m.configurationType
}
// GetCreatedDateTime gets the createdDateTime property value. The date and time the object was created.
func (m *GroupPolicyDefinitionValue) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetDefinition gets the definition property value. The associated group policy definition with the value.
func (m *GroupPolicyDefinitionValue) GetDefinition()(GroupPolicyDefinitionable) {
    return m.definition
}
// GetEnabled gets the enabled property value. Enables or disables the associated group policy definition.
func (m *GroupPolicyDefinitionValue) GetEnabled()(*bool) {
    return m.enabled
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GroupPolicyDefinitionValue) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["configurationType"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseGroupPolicyConfigurationType , m.SetConfigurationType)
    res["createdDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetCreatedDateTime)
    res["definition"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateGroupPolicyDefinitionFromDiscriminatorValue , m.SetDefinition)
    res["enabled"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetEnabled)
    res["lastModifiedDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetLastModifiedDateTime)
    res["presentationValues"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateGroupPolicyPresentationValueFromDiscriminatorValue , m.SetPresentationValues)
    return res
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The date and time the entity was last modified.
func (m *GroupPolicyDefinitionValue) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
}
// GetPresentationValues gets the presentationValues property value. The associated group policy presentation values with the definition value.
func (m *GroupPolicyDefinitionValue) GetPresentationValues()([]GroupPolicyPresentationValueable) {
    return m.presentationValues
}
// Serialize serializes information the current object
func (m *GroupPolicyDefinitionValue) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetConfigurationType() != nil {
        cast := (*m.GetConfigurationType()).String()
        err = writer.WriteStringValue("configurationType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("definition", m.GetDefinition())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("enabled", m.GetEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetPresentationValues() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetPresentationValues())
        err = writer.WriteCollectionOfObjectValues("presentationValues", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetConfigurationType sets the configurationType property value. Group Policy Configuration Type
func (m *GroupPolicyDefinitionValue) SetConfigurationType(value *GroupPolicyConfigurationType)() {
    m.configurationType = value
}
// SetCreatedDateTime sets the createdDateTime property value. The date and time the object was created.
func (m *GroupPolicyDefinitionValue) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetDefinition sets the definition property value. The associated group policy definition with the value.
func (m *GroupPolicyDefinitionValue) SetDefinition(value GroupPolicyDefinitionable)() {
    m.definition = value
}
// SetEnabled sets the enabled property value. Enables or disables the associated group policy definition.
func (m *GroupPolicyDefinitionValue) SetEnabled(value *bool)() {
    m.enabled = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The date and time the entity was last modified.
func (m *GroupPolicyDefinitionValue) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// SetPresentationValues sets the presentationValues property value. The associated group policy presentation values with the definition value.
func (m *GroupPolicyDefinitionValue) SetPresentationValues(value []GroupPolicyPresentationValueable)() {
    m.presentationValues = value
}
