package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type CustomClaimCondition struct {
    CustomClaimConditionBase
}
// NewCustomClaimCondition instantiates a new CustomClaimCondition and sets the default values.
func NewCustomClaimCondition()(*CustomClaimCondition) {
    m := &CustomClaimCondition{
        CustomClaimConditionBase: *NewCustomClaimConditionBase(),
    }
    odataTypeValue := "#microsoft.graph.customClaimCondition"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateCustomClaimConditionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCustomClaimConditionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCustomClaimCondition(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CustomClaimCondition) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.CustomClaimConditionBase.GetFieldDeserializers()
    res["memberOf"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetMemberOf(res)
        }
        return nil
    }
    res["userType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseClaimConditionUserType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserType(val.(*ClaimConditionUserType))
        }
        return nil
    }
    return res
}
// GetMemberOf gets the memberOf property value. The memberOf property
// returns a []string when successful
func (m *CustomClaimCondition) GetMemberOf()([]string) {
    val, err := m.GetBackingStore().Get("memberOf")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetUserType gets the userType property value. The userType property
// returns a *ClaimConditionUserType when successful
func (m *CustomClaimCondition) GetUserType()(*ClaimConditionUserType) {
    val, err := m.GetBackingStore().Get("userType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ClaimConditionUserType)
    }
    return nil
}
// Serialize serializes information the current object
func (m *CustomClaimCondition) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.CustomClaimConditionBase.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetMemberOf() != nil {
        err = writer.WriteCollectionOfStringValues("memberOf", m.GetMemberOf())
        if err != nil {
            return err
        }
    }
    if m.GetUserType() != nil {
        cast := (*m.GetUserType()).String()
        err = writer.WriteStringValue("userType", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetMemberOf sets the memberOf property value. The memberOf property
func (m *CustomClaimCondition) SetMemberOf(value []string)() {
    err := m.GetBackingStore().Set("memberOf", value)
    if err != nil {
        panic(err)
    }
}
// SetUserType sets the userType property value. The userType property
func (m *CustomClaimCondition) SetUserType(value *ClaimConditionUserType)() {
    err := m.GetBackingStore().Set("userType", value)
    if err != nil {
        panic(err)
    }
}
type CustomClaimConditionable interface {
    CustomClaimConditionBaseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetMemberOf()([]string)
    GetUserType()(*ClaimConditionUserType)
    SetMemberOf(value []string)()
    SetUserType(value *ClaimConditionUserType)()
}
