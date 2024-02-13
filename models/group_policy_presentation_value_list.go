package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GroupPolicyPresentationValueList the entity represents a collection of name/value pairs of a list box presentation on a policy definition.
type GroupPolicyPresentationValueList struct {
    GroupPolicyPresentationValue
}
// NewGroupPolicyPresentationValueList instantiates a new GroupPolicyPresentationValueList and sets the default values.
func NewGroupPolicyPresentationValueList()(*GroupPolicyPresentationValueList) {
    m := &GroupPolicyPresentationValueList{
        GroupPolicyPresentationValue: *NewGroupPolicyPresentationValue(),
    }
    return m
}
// CreateGroupPolicyPresentationValueListFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateGroupPolicyPresentationValueListFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGroupPolicyPresentationValueList(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *GroupPolicyPresentationValueList) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.GroupPolicyPresentationValue.GetFieldDeserializers()
    res["values"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateKeyValuePairFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]KeyValuePairable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(KeyValuePairable)
                }
            }
            m.SetValues(res)
        }
        return nil
    }
    return res
}
// GetValues gets the values property value. A list of pairs for the associated presentation.
// returns a []KeyValuePairable when successful
func (m *GroupPolicyPresentationValueList) GetValues()([]KeyValuePairable) {
    val, err := m.GetBackingStore().Get("values")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]KeyValuePairable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *GroupPolicyPresentationValueList) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.GroupPolicyPresentationValue.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetValues() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetValues()))
        for i, v := range m.GetValues() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("values", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetValues sets the values property value. A list of pairs for the associated presentation.
func (m *GroupPolicyPresentationValueList) SetValues(value []KeyValuePairable)() {
    err := m.GetBackingStore().Set("values", value)
    if err != nil {
        panic(err)
    }
}
type GroupPolicyPresentationValueListable interface {
    GroupPolicyPresentationValueable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetValues()([]KeyValuePairable)
    SetValues(value []KeyValuePairable)()
}
