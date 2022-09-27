package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GroupPolicyPresentationValue the base presentation value entity that stores the value for a single group policy presentation.
type GroupPolicyPresentationValue struct {
    Entity
    // The date and time the object was created.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The group policy definition value associated with the presentation value.
    definitionValue GroupPolicyDefinitionValueable
    // The date and time the object was last modified.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The group policy presentation associated with the presentation value.
    presentation GroupPolicyPresentationable
}
// NewGroupPolicyPresentationValue instantiates a new groupPolicyPresentationValue and sets the default values.
func NewGroupPolicyPresentationValue()(*GroupPolicyPresentationValue) {
    m := &GroupPolicyPresentationValue{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.groupPolicyPresentationValue";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateGroupPolicyPresentationValueFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGroupPolicyPresentationValueFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                switch *mappingValue {
                    case "#microsoft.graph.groupPolicyPresentationValueBoolean":
                        return NewGroupPolicyPresentationValueBoolean(), nil
                    case "#microsoft.graph.groupPolicyPresentationValueDecimal":
                        return NewGroupPolicyPresentationValueDecimal(), nil
                    case "#microsoft.graph.groupPolicyPresentationValueList":
                        return NewGroupPolicyPresentationValueList(), nil
                    case "#microsoft.graph.groupPolicyPresentationValueLongDecimal":
                        return NewGroupPolicyPresentationValueLongDecimal(), nil
                    case "#microsoft.graph.groupPolicyPresentationValueMultiText":
                        return NewGroupPolicyPresentationValueMultiText(), nil
                    case "#microsoft.graph.groupPolicyPresentationValueText":
                        return NewGroupPolicyPresentationValueText(), nil
                }
            }
        }
    }
    return NewGroupPolicyPresentationValue(), nil
}
// GetCreatedDateTime gets the createdDateTime property value. The date and time the object was created.
func (m *GroupPolicyPresentationValue) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetDefinitionValue gets the definitionValue property value. The group policy definition value associated with the presentation value.
func (m *GroupPolicyPresentationValue) GetDefinitionValue()(GroupPolicyDefinitionValueable) {
    return m.definitionValue
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GroupPolicyPresentationValue) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["createdDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetCreatedDateTime)
    res["definitionValue"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateGroupPolicyDefinitionValueFromDiscriminatorValue , m.SetDefinitionValue)
    res["lastModifiedDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetLastModifiedDateTime)
    res["presentation"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateGroupPolicyPresentationFromDiscriminatorValue , m.SetPresentation)
    return res
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The date and time the object was last modified.
func (m *GroupPolicyPresentationValue) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
}
// GetPresentation gets the presentation property value. The group policy presentation associated with the presentation value.
func (m *GroupPolicyPresentationValue) GetPresentation()(GroupPolicyPresentationable) {
    return m.presentation
}
// Serialize serializes information the current object
func (m *GroupPolicyPresentationValue) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("definitionValue", m.GetDefinitionValue())
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
    {
        err = writer.WriteObjectValue("presentation", m.GetPresentation())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCreatedDateTime sets the createdDateTime property value. The date and time the object was created.
func (m *GroupPolicyPresentationValue) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetDefinitionValue sets the definitionValue property value. The group policy definition value associated with the presentation value.
func (m *GroupPolicyPresentationValue) SetDefinitionValue(value GroupPolicyDefinitionValueable)() {
    m.definitionValue = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The date and time the object was last modified.
func (m *GroupPolicyPresentationValue) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// SetPresentation sets the presentation property value. The group policy presentation associated with the presentation value.
func (m *GroupPolicyPresentationValue) SetPresentation(value GroupPolicyPresentationable)() {
    m.presentation = value
}
