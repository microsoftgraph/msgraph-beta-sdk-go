package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// CrossTenantAccessPolicyB2BSetting provides operations to manage the policyRoot singleton.
type CrossTenantAccessPolicyB2BSetting struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The list of applications targeted with your cross-tenant access policy.
    applications CrossTenantAccessPolicyTargetConfigurationable;
    // The list of users and groups targeted with your cross-tenant access policy.
    usersAndGroups CrossTenantAccessPolicyTargetConfigurationable;
}
// NewCrossTenantAccessPolicyB2BSetting instantiates a new crossTenantAccessPolicyB2BSetting and sets the default values.
func NewCrossTenantAccessPolicyB2BSetting()(*CrossTenantAccessPolicyB2BSetting) {
    m := &CrossTenantAccessPolicyB2BSetting{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateCrossTenantAccessPolicyB2BSettingFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCrossTenantAccessPolicyB2BSettingFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewCrossTenantAccessPolicyB2BSetting(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CrossTenantAccessPolicyB2BSetting) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetApplications gets the applications property value. The list of applications targeted with your cross-tenant access policy.
func (m *CrossTenantAccessPolicyB2BSetting) GetApplications()(CrossTenantAccessPolicyTargetConfigurationable) {
    if m == nil {
        return nil
    } else {
        return m.applications
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CrossTenantAccessPolicyB2BSetting) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["applications"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateCrossTenantAccessPolicyTargetConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplications(val.(CrossTenantAccessPolicyTargetConfigurationable))
        }
        return nil
    }
    res["usersAndGroups"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateCrossTenantAccessPolicyTargetConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUsersAndGroups(val.(CrossTenantAccessPolicyTargetConfigurationable))
        }
        return nil
    }
    return res
}
// GetUsersAndGroups gets the usersAndGroups property value. The list of users and groups targeted with your cross-tenant access policy.
func (m *CrossTenantAccessPolicyB2BSetting) GetUsersAndGroups()(CrossTenantAccessPolicyTargetConfigurationable) {
    if m == nil {
        return nil
    } else {
        return m.usersAndGroups
    }
}
func (m *CrossTenantAccessPolicyB2BSetting) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *CrossTenantAccessPolicyB2BSetting) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("applications", m.GetApplications())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("usersAndGroups", m.GetUsersAndGroups())
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
func (m *CrossTenantAccessPolicyB2BSetting) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetApplications sets the applications property value. The list of applications targeted with your cross-tenant access policy.
func (m *CrossTenantAccessPolicyB2BSetting) SetApplications(value CrossTenantAccessPolicyTargetConfigurationable)() {
    if m != nil {
        m.applications = value
    }
}
// SetUsersAndGroups sets the usersAndGroups property value. The list of users and groups targeted with your cross-tenant access policy.
func (m *CrossTenantAccessPolicyB2BSetting) SetUsersAndGroups(value CrossTenantAccessPolicyTargetConfigurationable)() {
    if m != nil {
        m.usersAndGroups = value
    }
}
