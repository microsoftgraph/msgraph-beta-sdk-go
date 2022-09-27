package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ConditionalAccessExternalTenants 
type ConditionalAccessExternalTenants struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The membershipKind property
    membershipKind *ConditionalAccessExternalTenantsMembershipKind
    // The OdataType property
    odataType *string
}
// NewConditionalAccessExternalTenants instantiates a new conditionalAccessExternalTenants and sets the default values.
func NewConditionalAccessExternalTenants()(*ConditionalAccessExternalTenants) {
    m := &ConditionalAccessExternalTenants{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    odataTypeValue := "#microsoft.graph.conditionalAccessExternalTenants";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateConditionalAccessExternalTenantsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateConditionalAccessExternalTenantsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.conditionalAccessAllExternalTenants":
                        return NewConditionalAccessAllExternalTenants(), nil
                    case "#microsoft.graph.conditionalAccessEnumeratedExternalTenants":
                        return NewConditionalAccessEnumeratedExternalTenants(), nil
                }
            }
        }
    }
    return NewConditionalAccessExternalTenants(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ConditionalAccessExternalTenants) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ConditionalAccessExternalTenants) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["membershipKind"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseConditionalAccessExternalTenantsMembershipKind , m.SetMembershipKind)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    return res
}
// GetMembershipKind gets the membershipKind property value. The membershipKind property
func (m *ConditionalAccessExternalTenants) GetMembershipKind()(*ConditionalAccessExternalTenantsMembershipKind) {
    return m.membershipKind
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *ConditionalAccessExternalTenants) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *ConditionalAccessExternalTenants) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetMembershipKind() != nil {
        cast := (*m.GetMembershipKind()).String()
        err := writer.WriteStringValue("membershipKind", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ConditionalAccessExternalTenants) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetMembershipKind sets the membershipKind property value. The membershipKind property
func (m *ConditionalAccessExternalTenants) SetMembershipKind(value *ConditionalAccessExternalTenantsMembershipKind)() {
    m.membershipKind = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *ConditionalAccessExternalTenants) SetOdataType(value *string)() {
    m.odataType = value
}
