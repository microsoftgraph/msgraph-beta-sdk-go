package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type CustomSecurityAttributeExemption struct {
    Entity
}
// NewCustomSecurityAttributeExemption instantiates a new CustomSecurityAttributeExemption and sets the default values.
func NewCustomSecurityAttributeExemption()(*CustomSecurityAttributeExemption) {
    m := &CustomSecurityAttributeExemption{
        Entity: *NewEntity(),
    }
    return m
}
// CreateCustomSecurityAttributeExemptionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCustomSecurityAttributeExemptionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.customSecurityAttributeStringValueExemption":
                        return NewCustomSecurityAttributeStringValueExemption(), nil
                }
            }
        }
    }
    return NewCustomSecurityAttributeExemption(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CustomSecurityAttributeExemption) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["operator"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCustomSecurityAttributeComparisonOperator)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOperator(val.(*CustomSecurityAttributeComparisonOperator))
        }
        return nil
    }
    return res
}
// GetOperator gets the operator property value. The operator property
// returns a *CustomSecurityAttributeComparisonOperator when successful
func (m *CustomSecurityAttributeExemption) GetOperator()(*CustomSecurityAttributeComparisonOperator) {
    val, err := m.GetBackingStore().Get("operator")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*CustomSecurityAttributeComparisonOperator)
    }
    return nil
}
// Serialize serializes information the current object
func (m *CustomSecurityAttributeExemption) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetOperator() != nil {
        cast := (*m.GetOperator()).String()
        err = writer.WriteStringValue("operator", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetOperator sets the operator property value. The operator property
func (m *CustomSecurityAttributeExemption) SetOperator(value *CustomSecurityAttributeComparisonOperator)() {
    err := m.GetBackingStore().Set("operator", value)
    if err != nil {
        panic(err)
    }
}
type CustomSecurityAttributeExemptionable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetOperator()(*CustomSecurityAttributeComparisonOperator)
    SetOperator(value *CustomSecurityAttributeComparisonOperator)()
}
