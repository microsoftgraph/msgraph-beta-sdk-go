package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GroupPolicyPresentation the base entity for the display presentation of any of the additional options in a group policy definition.
type GroupPolicyPresentation struct {
    Entity
}
// NewGroupPolicyPresentation instantiates a new GroupPolicyPresentation and sets the default values.
func NewGroupPolicyPresentation()(*GroupPolicyPresentation) {
    m := &GroupPolicyPresentation{
        Entity: *NewEntity(),
    }
    return m
}
// CreateGroupPolicyPresentationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateGroupPolicyPresentationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.groupPolicyPresentationCheckBox":
                        return NewGroupPolicyPresentationCheckBox(), nil
                    case "#microsoft.graph.groupPolicyPresentationComboBox":
                        return NewGroupPolicyPresentationComboBox(), nil
                    case "#microsoft.graph.groupPolicyPresentationDecimalTextBox":
                        return NewGroupPolicyPresentationDecimalTextBox(), nil
                    case "#microsoft.graph.groupPolicyPresentationDropdownList":
                        return NewGroupPolicyPresentationDropdownList(), nil
                    case "#microsoft.graph.groupPolicyPresentationListBox":
                        return NewGroupPolicyPresentationListBox(), nil
                    case "#microsoft.graph.groupPolicyPresentationLongDecimalTextBox":
                        return NewGroupPolicyPresentationLongDecimalTextBox(), nil
                    case "#microsoft.graph.groupPolicyPresentationMultiTextBox":
                        return NewGroupPolicyPresentationMultiTextBox(), nil
                    case "#microsoft.graph.groupPolicyPresentationText":
                        return NewGroupPolicyPresentationText(), nil
                    case "#microsoft.graph.groupPolicyPresentationTextBox":
                        return NewGroupPolicyPresentationTextBox(), nil
                    case "#microsoft.graph.groupPolicyUploadedPresentation":
                        return NewGroupPolicyUploadedPresentation(), nil
                }
            }
        }
    }
    return NewGroupPolicyPresentation(), nil
}
// GetDefinition gets the definition property value. The group policy definition associated with the presentation.
// returns a GroupPolicyDefinitionable when successful
func (m *GroupPolicyPresentation) GetDefinition()(GroupPolicyDefinitionable) {
    val, err := m.GetBackingStore().Get("definition")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(GroupPolicyDefinitionable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *GroupPolicyPresentation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["definition"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateGroupPolicyDefinitionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefinition(val.(GroupPolicyDefinitionable))
        }
        return nil
    }
    res["label"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLabel(val)
        }
        return nil
    }
    res["lastModifiedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedDateTime(val)
        }
        return nil
    }
    return res
}
// GetLabel gets the label property value. Localized text label for any presentation entity. The default value is empty.
// returns a *string when successful
func (m *GroupPolicyPresentation) GetLabel()(*string) {
    val, err := m.GetBackingStore().Get("label")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The date and time the entity was last modified.
// returns a *Time when successful
func (m *GroupPolicyPresentation) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("lastModifiedDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// Serialize serializes information the current object
func (m *GroupPolicyPresentation) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("definition", m.GetDefinition())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("label", m.GetLabel())
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
    return nil
}
// SetDefinition sets the definition property value. The group policy definition associated with the presentation.
func (m *GroupPolicyPresentation) SetDefinition(value GroupPolicyDefinitionable)() {
    err := m.GetBackingStore().Set("definition", value)
    if err != nil {
        panic(err)
    }
}
// SetLabel sets the label property value. Localized text label for any presentation entity. The default value is empty.
func (m *GroupPolicyPresentation) SetLabel(value *string)() {
    err := m.GetBackingStore().Set("label", value)
    if err != nil {
        panic(err)
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The date and time the entity was last modified.
func (m *GroupPolicyPresentation) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("lastModifiedDateTime", value)
    if err != nil {
        panic(err)
    }
}
type GroupPolicyPresentationable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDefinition()(GroupPolicyDefinitionable)
    GetLabel()(*string)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    SetDefinition(value GroupPolicyDefinitionable)()
    SetLabel(value *string)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
}
