package apply

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ApplyPostRequestBody provides operations to call the apply method.
type ApplyPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
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
// NewApplyPostRequestBody instantiates a new applyPostRequestBody and sets the default values.
func NewApplyPostRequestBody()(*ApplyPostRequestBody) {
    m := &ApplyPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateApplyPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateApplyPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewApplyPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ApplyPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetExcludeGroups gets the excludeGroups property value. The excludeGroups property
func (m *ApplyPostRequestBody) GetExcludeGroups()([]string) {
    return m.excludeGroups
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ApplyPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["excludeGroups"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetExcludeGroups)
    res["includeAllUsers"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIncludeAllUsers)
    res["includeGroups"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetIncludeGroups)
    res["managementTemplateId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetManagementTemplateId)
    res["tenantGroupId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetTenantGroupId)
    res["tenantId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetTenantId)
    return res
}
// GetIncludeAllUsers gets the includeAllUsers property value. The includeAllUsers property
func (m *ApplyPostRequestBody) GetIncludeAllUsers()(*bool) {
    return m.includeAllUsers
}
// GetIncludeGroups gets the includeGroups property value. The includeGroups property
func (m *ApplyPostRequestBody) GetIncludeGroups()([]string) {
    return m.includeGroups
}
// GetManagementTemplateId gets the managementTemplateId property value. The managementTemplateId property
func (m *ApplyPostRequestBody) GetManagementTemplateId()(*string) {
    return m.managementTemplateId
}
// GetTenantGroupId gets the tenantGroupId property value. The tenantGroupId property
func (m *ApplyPostRequestBody) GetTenantGroupId()(*string) {
    return m.tenantGroupId
}
// GetTenantId gets the tenantId property value. The tenantId property
func (m *ApplyPostRequestBody) GetTenantId()(*string) {
    return m.tenantId
}
// Serialize serializes information the current object
func (m *ApplyPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *ApplyPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetExcludeGroups sets the excludeGroups property value. The excludeGroups property
func (m *ApplyPostRequestBody) SetExcludeGroups(value []string)() {
    m.excludeGroups = value
}
// SetIncludeAllUsers sets the includeAllUsers property value. The includeAllUsers property
func (m *ApplyPostRequestBody) SetIncludeAllUsers(value *bool)() {
    m.includeAllUsers = value
}
// SetIncludeGroups sets the includeGroups property value. The includeGroups property
func (m *ApplyPostRequestBody) SetIncludeGroups(value []string)() {
    m.includeGroups = value
}
// SetManagementTemplateId sets the managementTemplateId property value. The managementTemplateId property
func (m *ApplyPostRequestBody) SetManagementTemplateId(value *string)() {
    m.managementTemplateId = value
}
// SetTenantGroupId sets the tenantGroupId property value. The tenantGroupId property
func (m *ApplyPostRequestBody) SetTenantGroupId(value *string)() {
    m.tenantGroupId = value
}
// SetTenantId sets the tenantId property value. The tenantId property
func (m *ApplyPostRequestBody) SetTenantId(value *string)() {
    m.tenantId = value
}
