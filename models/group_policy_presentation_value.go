package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GroupPolicyPresentationValue the base presentation value entity that stores the value for a single group policy presentation.
type GroupPolicyPresentationValue struct {
    Entity
}
// NewGroupPolicyPresentationValue instantiates a new groupPolicyPresentationValue and sets the default values.
func NewGroupPolicyPresentationValue()(*GroupPolicyPresentationValue) {
    m := &GroupPolicyPresentationValue{
        Entity: *NewEntity(),
    }
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
    val, err := m.GetBackingStore().Get("createdDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetDefinitionValue gets the definitionValue property value. The group policy definition value associated with the presentation value.
func (m *GroupPolicyPresentationValue) GetDefinitionValue()(GroupPolicyDefinitionValueable) {
    val, err := m.GetBackingStore().Get("definitionValue")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(GroupPolicyDefinitionValueable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GroupPolicyPresentationValue) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["createdDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
        }
        return nil
    }
    res["definitionValue"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateGroupPolicyDefinitionValueFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefinitionValue(val.(GroupPolicyDefinitionValueable))
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
    res["presentation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateGroupPolicyPresentationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPresentation(val.(GroupPolicyPresentationable))
        }
        return nil
    }
    return res
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The date and time the object was last modified.
func (m *GroupPolicyPresentationValue) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("lastModifiedDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetPresentation gets the presentation property value. The group policy presentation associated with the presentation value.
func (m *GroupPolicyPresentationValue) GetPresentation()(GroupPolicyPresentationable) {
    val, err := m.GetBackingStore().Get("presentation")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(GroupPolicyPresentationable)
    }
    return nil
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
    err := m.GetBackingStore().Set("createdDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetDefinitionValue sets the definitionValue property value. The group policy definition value associated with the presentation value.
func (m *GroupPolicyPresentationValue) SetDefinitionValue(value GroupPolicyDefinitionValueable)() {
    err := m.GetBackingStore().Set("definitionValue", value)
    if err != nil {
        panic(err)
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The date and time the object was last modified.
func (m *GroupPolicyPresentationValue) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("lastModifiedDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetPresentation sets the presentation property value. The group policy presentation associated with the presentation value.
func (m *GroupPolicyPresentationValue) SetPresentation(value GroupPolicyPresentationable)() {
    err := m.GetBackingStore().Set("presentation", value)
    if err != nil {
        panic(err)
    }
}
// GroupPolicyPresentationValueable 
type GroupPolicyPresentationValueable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDefinitionValue()(GroupPolicyDefinitionValueable)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetPresentation()(GroupPolicyPresentationable)
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDefinitionValue(value GroupPolicyDefinitionValueable)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetPresentation(value GroupPolicyPresentationable)()
}
