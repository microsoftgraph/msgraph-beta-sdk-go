package tenantrelationships

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ManagedTenantsManagementActionsItemMicrosoftGraphManagedTenantsApplyApplyPostRequestBody 
type ManagedTenantsManagementActionsItemMicrosoftGraphManagedTenantsApplyApplyPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The excludeGroups property
    excludeGroups []string
    // The includeAllUsers property
    includeAllUsers *bool
    // The includeGroups property
    includeGroups []string
    // The managementTemplateId property
    managementTemplateId *string
    // The tenantGroupId property
    tenantGroupId *string
    // The tenantId property
    tenantId *string
}
// NewManagedTenantsManagementActionsItemMicrosoftGraphManagedTenantsApplyApplyPostRequestBody instantiates a new ManagedTenantsManagementActionsItemMicrosoftGraphManagedTenantsApplyApplyPostRequestBody and sets the default values.
func NewManagedTenantsManagementActionsItemMicrosoftGraphManagedTenantsApplyApplyPostRequestBody()(*ManagedTenantsManagementActionsItemMicrosoftGraphManagedTenantsApplyApplyPostRequestBody) {
    m := &ManagedTenantsManagementActionsItemMicrosoftGraphManagedTenantsApplyApplyPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateManagedTenantsManagementActionsItemMicrosoftGraphManagedTenantsApplyApplyPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateManagedTenantsManagementActionsItemMicrosoftGraphManagedTenantsApplyApplyPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewManagedTenantsManagementActionsItemMicrosoftGraphManagedTenantsApplyApplyPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ManagedTenantsManagementActionsItemMicrosoftGraphManagedTenantsApplyApplyPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetExcludeGroups gets the excludeGroups property value. The excludeGroups property
func (m *ManagedTenantsManagementActionsItemMicrosoftGraphManagedTenantsApplyApplyPostRequestBody) GetExcludeGroups()([]string) {
    return m.excludeGroups
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ManagedTenantsManagementActionsItemMicrosoftGraphManagedTenantsApplyApplyPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["excludeGroups"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetExcludeGroups(res)
        }
        return nil
    }
    res["includeAllUsers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIncludeAllUsers(val)
        }
        return nil
    }
    res["includeGroups"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetIncludeGroups(res)
        }
        return nil
    }
    res["managementTemplateId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagementTemplateId(val)
        }
        return nil
    }
    res["tenantGroupId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTenantGroupId(val)
        }
        return nil
    }
    res["tenantId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTenantId(val)
        }
        return nil
    }
    return res
}
// GetIncludeAllUsers gets the includeAllUsers property value. The includeAllUsers property
func (m *ManagedTenantsManagementActionsItemMicrosoftGraphManagedTenantsApplyApplyPostRequestBody) GetIncludeAllUsers()(*bool) {
    return m.includeAllUsers
}
// GetIncludeGroups gets the includeGroups property value. The includeGroups property
func (m *ManagedTenantsManagementActionsItemMicrosoftGraphManagedTenantsApplyApplyPostRequestBody) GetIncludeGroups()([]string) {
    return m.includeGroups
}
// GetManagementTemplateId gets the managementTemplateId property value. The managementTemplateId property
func (m *ManagedTenantsManagementActionsItemMicrosoftGraphManagedTenantsApplyApplyPostRequestBody) GetManagementTemplateId()(*string) {
    return m.managementTemplateId
}
// GetTenantGroupId gets the tenantGroupId property value. The tenantGroupId property
func (m *ManagedTenantsManagementActionsItemMicrosoftGraphManagedTenantsApplyApplyPostRequestBody) GetTenantGroupId()(*string) {
    return m.tenantGroupId
}
// GetTenantId gets the tenantId property value. The tenantId property
func (m *ManagedTenantsManagementActionsItemMicrosoftGraphManagedTenantsApplyApplyPostRequestBody) GetTenantId()(*string) {
    return m.tenantId
}
// Serialize serializes information the current object
func (m *ManagedTenantsManagementActionsItemMicrosoftGraphManagedTenantsApplyApplyPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetExcludeGroups() != nil {
        err := writer.WriteCollectionOfStringValues("excludeGroups", m.GetExcludeGroups())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("includeAllUsers", m.GetIncludeAllUsers())
        if err != nil {
            return err
        }
    }
    if m.GetIncludeGroups() != nil {
        err := writer.WriteCollectionOfStringValues("includeGroups", m.GetIncludeGroups())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("managementTemplateId", m.GetManagementTemplateId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("tenantGroupId", m.GetTenantGroupId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("tenantId", m.GetTenantId())
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
func (m *ManagedTenantsManagementActionsItemMicrosoftGraphManagedTenantsApplyApplyPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetExcludeGroups sets the excludeGroups property value. The excludeGroups property
func (m *ManagedTenantsManagementActionsItemMicrosoftGraphManagedTenantsApplyApplyPostRequestBody) SetExcludeGroups(value []string)() {
    m.excludeGroups = value
}
// SetIncludeAllUsers sets the includeAllUsers property value. The includeAllUsers property
func (m *ManagedTenantsManagementActionsItemMicrosoftGraphManagedTenantsApplyApplyPostRequestBody) SetIncludeAllUsers(value *bool)() {
    m.includeAllUsers = value
}
// SetIncludeGroups sets the includeGroups property value. The includeGroups property
func (m *ManagedTenantsManagementActionsItemMicrosoftGraphManagedTenantsApplyApplyPostRequestBody) SetIncludeGroups(value []string)() {
    m.includeGroups = value
}
// SetManagementTemplateId sets the managementTemplateId property value. The managementTemplateId property
func (m *ManagedTenantsManagementActionsItemMicrosoftGraphManagedTenantsApplyApplyPostRequestBody) SetManagementTemplateId(value *string)() {
    m.managementTemplateId = value
}
// SetTenantGroupId sets the tenantGroupId property value. The tenantGroupId property
func (m *ManagedTenantsManagementActionsItemMicrosoftGraphManagedTenantsApplyApplyPostRequestBody) SetTenantGroupId(value *string)() {
    m.tenantGroupId = value
}
// SetTenantId sets the tenantId property value. The tenantId property
func (m *ManagedTenantsManagementActionsItemMicrosoftGraphManagedTenantsApplyApplyPostRequestBody) SetTenantId(value *string)() {
    m.tenantId = value
}
