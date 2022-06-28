package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CrossTenantAccessPolicy 
type CrossTenantAccessPolicy struct {
    TenantRelationshipAccessPolicyBase
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Used to specify which Microsoft clouds an organization would like to collaborate with. By default, this value is empty. Supported values for this field are: microsoftonline.com, microsoftonline.us, and partner.microsoftonline.cn.
    allowedCloudEndpoints []string
    // Defines the default configuration for how your organization interacts with external Azure Active Directory organizations.
    default_escaped CrossTenantAccessPolicyConfigurationDefaultable
    // Defines partner-specific configurations for external Azure Active Directory organizations.
    partners []CrossTenantAccessPolicyConfigurationPartnerable
}
// NewCrossTenantAccessPolicy instantiates a new CrossTenantAccessPolicy and sets the default values.
func NewCrossTenantAccessPolicy()(*CrossTenantAccessPolicy) {
    m := &CrossTenantAccessPolicy{
        TenantRelationshipAccessPolicyBase: *NewTenantRelationshipAccessPolicyBase(),
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateCrossTenantAccessPolicyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCrossTenantAccessPolicyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCrossTenantAccessPolicy(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CrossTenantAccessPolicy) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAllowedCloudEndpoints gets the allowedCloudEndpoints property value. Used to specify which Microsoft clouds an organization would like to collaborate with. By default, this value is empty. Supported values for this field are: microsoftonline.com, microsoftonline.us, and partner.microsoftonline.cn.
func (m *CrossTenantAccessPolicy) GetAllowedCloudEndpoints()([]string) {
    if m == nil {
        return nil
    } else {
        return m.allowedCloudEndpoints
    }
}
// GetDefault gets the default property value. Defines the default configuration for how your organization interacts with external Azure Active Directory organizations.
func (m *CrossTenantAccessPolicy) GetDefault()(CrossTenantAccessPolicyConfigurationDefaultable) {
    if m == nil {
        return nil
    } else {
        return m.default_escaped
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CrossTenantAccessPolicy) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.TenantRelationshipAccessPolicyBase.GetFieldDeserializers()
    res["allowedCloudEndpoints"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetAllowedCloudEndpoints(res)
        }
        return nil
    }
    res["default"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCrossTenantAccessPolicyConfigurationDefaultFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefault(val.(CrossTenantAccessPolicyConfigurationDefaultable))
        }
        return nil
    }
    res["partners"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCrossTenantAccessPolicyConfigurationPartnerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CrossTenantAccessPolicyConfigurationPartnerable, len(val))
            for i, v := range val {
                res[i] = v.(CrossTenantAccessPolicyConfigurationPartnerable)
            }
            m.SetPartners(res)
        }
        return nil
    }
    return res
}
// GetPartners gets the partners property value. Defines partner-specific configurations for external Azure Active Directory organizations.
func (m *CrossTenantAccessPolicy) GetPartners()([]CrossTenantAccessPolicyConfigurationPartnerable) {
    if m == nil {
        return nil
    } else {
        return m.partners
    }
}
// Serialize serializes information the current object
func (m *CrossTenantAccessPolicy) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.TenantRelationshipAccessPolicyBase.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAllowedCloudEndpoints() != nil {
        err = writer.WriteCollectionOfStringValues("allowedCloudEndpoints", m.GetAllowedCloudEndpoints())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("default", m.GetDefault())
        if err != nil {
            return err
        }
    }
    if m.GetPartners() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetPartners()))
        for i, v := range m.GetPartners() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("partners", cast)
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
func (m *CrossTenantAccessPolicy) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAllowedCloudEndpoints sets the allowedCloudEndpoints property value. Used to specify which Microsoft clouds an organization would like to collaborate with. By default, this value is empty. Supported values for this field are: microsoftonline.com, microsoftonline.us, and partner.microsoftonline.cn.
func (m *CrossTenantAccessPolicy) SetAllowedCloudEndpoints(value []string)() {
    if m != nil {
        m.allowedCloudEndpoints = value
    }
}
// SetDefault sets the default property value. Defines the default configuration for how your organization interacts with external Azure Active Directory organizations.
func (m *CrossTenantAccessPolicy) SetDefault(value CrossTenantAccessPolicyConfigurationDefaultable)() {
    if m != nil {
        m.default_escaped = value
    }
}
// SetPartners sets the partners property value. Defines partner-specific configurations for external Azure Active Directory organizations.
func (m *CrossTenantAccessPolicy) SetPartners(value []CrossTenantAccessPolicyConfigurationPartnerable)() {
    if m != nil {
        m.partners = value
    }
}
