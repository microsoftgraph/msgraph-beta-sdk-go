package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GroupPolicyPresentationValueMultiText 
type GroupPolicyPresentationValueMultiText struct {
    GroupPolicyPresentationValue
    // A collection of non-empty strings for the associated presentation.
    values []string
}
// NewGroupPolicyPresentationValueMultiText instantiates a new GroupPolicyPresentationValueMultiText and sets the default values.
func NewGroupPolicyPresentationValueMultiText()(*GroupPolicyPresentationValueMultiText) {
    m := &GroupPolicyPresentationValueMultiText{
        GroupPolicyPresentationValue: *NewGroupPolicyPresentationValue(),
    }
    odataTypeValue := "#microsoft.graph.groupPolicyPresentationValueMultiText";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateGroupPolicyPresentationValueMultiTextFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGroupPolicyPresentationValueMultiTextFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGroupPolicyPresentationValueMultiText(), nil
}
// GetFieldDeserializers the deserialization information for the current model
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
                res[i] = *(v.(*string))
            }
            m.SetValues(res)
        }
        return nil
    }
    return res
}
// GetValues gets the values property value. A collection of non-empty strings for the associated presentation.
func (m *GroupPolicyPresentationValueMultiText) GetValues()([]string) {
    return m.values
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
    m.values = value
}
