package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UnifiedRoleManagementPolicyAuthenticationContextRule 
type UnifiedRoleManagementPolicyAuthenticationContextRule struct {
    UnifiedRoleManagementPolicyRule
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The value of the authentication context claim.
    claimValue *string
    // Whether this rule is enabled.
    isEnabled *bool
}
// NewUnifiedRoleManagementPolicyAuthenticationContextRule instantiates a new UnifiedRoleManagementPolicyAuthenticationContextRule and sets the default values.
func NewUnifiedRoleManagementPolicyAuthenticationContextRule()(*UnifiedRoleManagementPolicyAuthenticationContextRule) {
    m := &UnifiedRoleManagementPolicyAuthenticationContextRule{
        UnifiedRoleManagementPolicyRule: *NewUnifiedRoleManagementPolicyRule(),
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateUnifiedRoleManagementPolicyAuthenticationContextRuleFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUnifiedRoleManagementPolicyAuthenticationContextRuleFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUnifiedRoleManagementPolicyAuthenticationContextRule(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UnifiedRoleManagementPolicyAuthenticationContextRule) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetClaimValue gets the claimValue property value. The value of the authentication context claim.
func (m *UnifiedRoleManagementPolicyAuthenticationContextRule) GetClaimValue()(*string) {
    if m == nil {
        return nil
    } else {
        return m.claimValue
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UnifiedRoleManagementPolicyAuthenticationContextRule) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.UnifiedRoleManagementPolicyRule.GetFieldDeserializers()
    res["claimValue"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClaimValue(val)
        }
        return nil
    }
    res["isEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsEnabled(val)
        }
        return nil
    }
    return res
}
// GetIsEnabled gets the isEnabled property value. Whether this rule is enabled.
func (m *UnifiedRoleManagementPolicyAuthenticationContextRule) GetIsEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isEnabled
    }
}
// Serialize serializes information the current object
func (m *UnifiedRoleManagementPolicyAuthenticationContextRule) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.UnifiedRoleManagementPolicyRule.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("claimValue", m.GetClaimValue())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isEnabled", m.GetIsEnabled())
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
func (m *UnifiedRoleManagementPolicyAuthenticationContextRule) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetClaimValue sets the claimValue property value. The value of the authentication context claim.
func (m *UnifiedRoleManagementPolicyAuthenticationContextRule) SetClaimValue(value *string)() {
    if m != nil {
        m.claimValue = value
    }
}
// SetIsEnabled sets the isEnabled property value. Whether this rule is enabled.
func (m *UnifiedRoleManagementPolicyAuthenticationContextRule) SetIsEnabled(value *bool)() {
    if m != nil {
        m.isEnabled = value
    }
}
