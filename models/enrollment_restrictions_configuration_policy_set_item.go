package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EnrollmentRestrictionsConfigurationPolicySetItem 
type EnrollmentRestrictionsConfigurationPolicySetItem struct {
    PolicySetItem
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Limit of the EnrollmentRestrictionsConfigurationPolicySetItem.
    limit *int32
    // Priority of the EnrollmentRestrictionsConfigurationPolicySetItem.
    priority *int32
}
// NewEnrollmentRestrictionsConfigurationPolicySetItem instantiates a new EnrollmentRestrictionsConfigurationPolicySetItem and sets the default values.
func NewEnrollmentRestrictionsConfigurationPolicySetItem()(*EnrollmentRestrictionsConfigurationPolicySetItem) {
    m := &EnrollmentRestrictionsConfigurationPolicySetItem{
        PolicySetItem: *NewPolicySetItem(),
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateEnrollmentRestrictionsConfigurationPolicySetItemFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEnrollmentRestrictionsConfigurationPolicySetItemFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEnrollmentRestrictionsConfigurationPolicySetItem(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *EnrollmentRestrictionsConfigurationPolicySetItem) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EnrollmentRestrictionsConfigurationPolicySetItem) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.PolicySetItem.GetFieldDeserializers()
    res["limit"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLimit(val)
        }
        return nil
    }
    res["priority"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPriority(val)
        }
        return nil
    }
    return res
}
// GetLimit gets the limit property value. Limit of the EnrollmentRestrictionsConfigurationPolicySetItem.
func (m *EnrollmentRestrictionsConfigurationPolicySetItem) GetLimit()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.limit
    }
}
// GetPriority gets the priority property value. Priority of the EnrollmentRestrictionsConfigurationPolicySetItem.
func (m *EnrollmentRestrictionsConfigurationPolicySetItem) GetPriority()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.priority
    }
}
// Serialize serializes information the current object
func (m *EnrollmentRestrictionsConfigurationPolicySetItem) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.PolicySetItem.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt32Value("limit", m.GetLimit())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("priority", m.GetPriority())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *EnrollmentRestrictionsConfigurationPolicySetItem) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetLimit sets the limit property value. Limit of the EnrollmentRestrictionsConfigurationPolicySetItem.
func (m *EnrollmentRestrictionsConfigurationPolicySetItem) SetLimit(value *int32)() {
    if m != nil {
        m.limit = value
    }
}
// SetPriority sets the priority property value. Priority of the EnrollmentRestrictionsConfigurationPolicySetItem.
func (m *EnrollmentRestrictionsConfigurationPolicySetItem) SetPriority(value *int32)() {
    if m != nil {
        m.priority = value
    }
}
