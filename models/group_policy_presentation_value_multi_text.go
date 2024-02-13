package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GroupPolicyPresentationValueMultiText the entity represents a string value of a multi-line text box presentation on a policy definition.
type GroupPolicyPresentationValueMultiText struct {
    GroupPolicyPresentationValue
}
// NewGroupPolicyPresentationValueMultiText instantiates a new GroupPolicyPresentationValueMultiText and sets the default values.
func NewGroupPolicyPresentationValueMultiText()(*GroupPolicyPresentationValueMultiText) {
    m := &GroupPolicyPresentationValueMultiText{
        GroupPolicyPresentationValue: *NewGroupPolicyPresentationValue(),
    }
    return m
}
// CreateGroupPolicyPresentationValueMultiTextFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateGroupPolicyPresentationValueMultiTextFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGroupPolicyPresentationValueMultiText(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *GroupPolicyPresentationValueMultiText) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.GroupPolicyPresentationValue.GetFieldDeserializers()
    res["values"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetValues(res)
        }
        return nil
    }
    return res
}
// GetValues gets the values property value. A collection of non-empty strings for the associated presentation.
// returns a []string when successful
func (m *GroupPolicyPresentationValueMultiText) GetValues()([]string) {
    val, err := m.GetBackingStore().Get("values")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *GroupPolicyPresentationValueMultiText) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.GroupPolicyPresentationValue.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetValues() != nil {
        err = writer.WriteCollectionOfStringValues("values", m.GetValues())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetValues sets the values property value. A collection of non-empty strings for the associated presentation.
func (m *GroupPolicyPresentationValueMultiText) SetValues(value []string)() {
    err := m.GetBackingStore().Set("values", value)
    if err != nil {
        panic(err)
    }
}
type GroupPolicyPresentationValueMultiTextable interface {
    GroupPolicyPresentationValueable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetValues()([]string)
    SetValues(value []string)()
}
