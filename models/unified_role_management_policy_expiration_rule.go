package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UnifiedRoleManagementPolicyExpirationRule 
type UnifiedRoleManagementPolicyExpirationRule struct {
    UnifiedRoleManagementPolicyRule
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Indicates whether expiration is required or if it's a permanently active assignment or eligibility.
    isExpirationRequired *bool
    // The maximum duration allowed for eligibility or assignment which is not permanent. Required when isExpirationRequired is true.
    maximumDuration *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration
}
// NewUnifiedRoleManagementPolicyExpirationRule instantiates a new UnifiedRoleManagementPolicyExpirationRule and sets the default values.
func NewUnifiedRoleManagementPolicyExpirationRule()(*UnifiedRoleManagementPolicyExpirationRule) {
    m := &UnifiedRoleManagementPolicyExpirationRule{
        UnifiedRoleManagementPolicyRule: *NewUnifiedRoleManagementPolicyRule(),
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateUnifiedRoleManagementPolicyExpirationRuleFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUnifiedRoleManagementPolicyExpirationRuleFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUnifiedRoleManagementPolicyExpirationRule(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UnifiedRoleManagementPolicyExpirationRule) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UnifiedRoleManagementPolicyExpirationRule) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.UnifiedRoleManagementPolicyRule.GetFieldDeserializers()
    res["isExpirationRequired"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsExpirationRequired(val)
        }
        return nil
    }
    res["maximumDuration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetISODurationValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaximumDuration(val)
        }
        return nil
    }
    return res
}
// GetIsExpirationRequired gets the isExpirationRequired property value. Indicates whether expiration is required or if it's a permanently active assignment or eligibility.
func (m *UnifiedRoleManagementPolicyExpirationRule) GetIsExpirationRequired()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isExpirationRequired
    }
}
// GetMaximumDuration gets the maximumDuration property value. The maximum duration allowed for eligibility or assignment which is not permanent. Required when isExpirationRequired is true.
func (m *UnifiedRoleManagementPolicyExpirationRule) GetMaximumDuration()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration) {
    if m == nil {
        return nil
    } else {
        return m.maximumDuration
    }
}
// Serialize serializes information the current object
func (m *UnifiedRoleManagementPolicyExpirationRule) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.UnifiedRoleManagementPolicyRule.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("isExpirationRequired", m.GetIsExpirationRequired())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteISODurationValue("maximumDuration", m.GetMaximumDuration())
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
func (m *UnifiedRoleManagementPolicyExpirationRule) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetIsExpirationRequired sets the isExpirationRequired property value. Indicates whether expiration is required or if it's a permanently active assignment or eligibility.
func (m *UnifiedRoleManagementPolicyExpirationRule) SetIsExpirationRequired(value *bool)() {
    if m != nil {
        m.isExpirationRequired = value
    }
}
// SetMaximumDuration sets the maximumDuration property value. The maximum duration allowed for eligibility or assignment which is not permanent. Required when isExpirationRequired is true.
func (m *UnifiedRoleManagementPolicyExpirationRule) SetMaximumDuration(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)() {
    if m != nil {
        m.maximumDuration = value
    }
}
