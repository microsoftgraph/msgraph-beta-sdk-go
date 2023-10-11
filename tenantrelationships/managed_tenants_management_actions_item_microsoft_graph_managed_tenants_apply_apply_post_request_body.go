package tenantrelationships

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// ManagedTenantsManagementActionsItemMicrosoftGraphManagedTenantsApplyApplyPostRequestBody 
type ManagedTenantsManagementActionsItemMicrosoftGraphManagedTenantsApplyApplyPostRequestBody struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewManagedTenantsManagementActionsItemMicrosoftGraphManagedTenantsApplyApplyPostRequestBody instantiates a new ManagedTenantsManagementActionsItemMicrosoftGraphManagedTenantsApplyApplyPostRequestBody and sets the default values.
func NewManagedTenantsManagementActionsItemMicrosoftGraphManagedTenantsApplyApplyPostRequestBody()(*ManagedTenantsManagementActionsItemMicrosoftGraphManagedTenantsApplyApplyPostRequestBody) {
    m := &ManagedTenantsManagementActionsItemMicrosoftGraphManagedTenantsApplyApplyPostRequestBody{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateManagedTenantsManagementActionsItemMicrosoftGraphManagedTenantsApplyApplyPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateManagedTenantsManagementActionsItemMicrosoftGraphManagedTenantsApplyApplyPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewManagedTenantsManagementActionsItemMicrosoftGraphManagedTenantsApplyApplyPostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ManagedTenantsManagementActionsItemMicrosoftGraphManagedTenantsApplyApplyPostRequestBody) GetAdditionalData()(map[string]any) {
    val , err :=  m.backingStore.Get("additionalData")
    if err != nil {
        panic(err)
    }
    if val == nil {
        var value = make(map[string]any);
        m.SetAdditionalData(value);
    }
    return val.(map[string]any)
}
// GetBackingStore gets the BackingStore property value. Stores model information.
func (m *ManagedTenantsManagementActionsItemMicrosoftGraphManagedTenantsApplyApplyPostRequestBody) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetExcludeGroups gets the excludeGroups property value. The excludeGroups property
func (m *ManagedTenantsManagementActionsItemMicrosoftGraphManagedTenantsApplyApplyPostRequestBody) GetExcludeGroups()([]string) {
    val, err := m.GetBackingStore().Get("excludeGroups")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
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
                if v != nil {
                    res[i] = *(v.(*string))
                }
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
                if v != nil {
                    res[i] = *(v.(*string))
                }
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
    val, err := m.GetBackingStore().Get("includeAllUsers")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetIncludeGroups gets the includeGroups property value. The includeGroups property
func (m *ManagedTenantsManagementActionsItemMicrosoftGraphManagedTenantsApplyApplyPostRequestBody) GetIncludeGroups()([]string) {
    val, err := m.GetBackingStore().Get("includeGroups")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetManagementTemplateId gets the managementTemplateId property value. The managementTemplateId property
func (m *ManagedTenantsManagementActionsItemMicrosoftGraphManagedTenantsApplyApplyPostRequestBody) GetManagementTemplateId()(*string) {
    val, err := m.GetBackingStore().Get("managementTemplateId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetTenantGroupId gets the tenantGroupId property value. The tenantGroupId property
func (m *ManagedTenantsManagementActionsItemMicrosoftGraphManagedTenantsApplyApplyPostRequestBody) GetTenantGroupId()(*string) {
    val, err := m.GetBackingStore().Get("tenantGroupId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetTenantId gets the tenantId property value. The tenantId property
func (m *ManagedTenantsManagementActionsItemMicrosoftGraphManagedTenantsApplyApplyPostRequestBody) GetTenantId()(*string) {
    val, err := m.GetBackingStore().Get("tenantId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
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
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ManagedTenantsManagementActionsItemMicrosoftGraphManagedTenantsApplyApplyPostRequestBody) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *ManagedTenantsManagementActionsItemMicrosoftGraphManagedTenantsApplyApplyPostRequestBody) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetExcludeGroups sets the excludeGroups property value. The excludeGroups property
func (m *ManagedTenantsManagementActionsItemMicrosoftGraphManagedTenantsApplyApplyPostRequestBody) SetExcludeGroups(value []string)() {
    err := m.GetBackingStore().Set("excludeGroups", value)
    if err != nil {
        panic(err)
    }
}
// SetIncludeAllUsers sets the includeAllUsers property value. The includeAllUsers property
func (m *ManagedTenantsManagementActionsItemMicrosoftGraphManagedTenantsApplyApplyPostRequestBody) SetIncludeAllUsers(value *bool)() {
    err := m.GetBackingStore().Set("includeAllUsers", value)
    if err != nil {
        panic(err)
    }
}
// SetIncludeGroups sets the includeGroups property value. The includeGroups property
func (m *ManagedTenantsManagementActionsItemMicrosoftGraphManagedTenantsApplyApplyPostRequestBody) SetIncludeGroups(value []string)() {
    err := m.GetBackingStore().Set("includeGroups", value)
    if err != nil {
        panic(err)
    }
}
// SetManagementTemplateId sets the managementTemplateId property value. The managementTemplateId property
func (m *ManagedTenantsManagementActionsItemMicrosoftGraphManagedTenantsApplyApplyPostRequestBody) SetManagementTemplateId(value *string)() {
    err := m.GetBackingStore().Set("managementTemplateId", value)
    if err != nil {
        panic(err)
    }
}
// SetTenantGroupId sets the tenantGroupId property value. The tenantGroupId property
func (m *ManagedTenantsManagementActionsItemMicrosoftGraphManagedTenantsApplyApplyPostRequestBody) SetTenantGroupId(value *string)() {
    err := m.GetBackingStore().Set("tenantGroupId", value)
    if err != nil {
        panic(err)
    }
}
// SetTenantId sets the tenantId property value. The tenantId property
func (m *ManagedTenantsManagementActionsItemMicrosoftGraphManagedTenantsApplyApplyPostRequestBody) SetTenantId(value *string)() {
    err := m.GetBackingStore().Set("tenantId", value)
    if err != nil {
        panic(err)
    }
}
// ManagedTenantsManagementActionsItemMicrosoftGraphManagedTenantsApplyApplyPostRequestBodyable 
type ManagedTenantsManagementActionsItemMicrosoftGraphManagedTenantsApplyApplyPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetExcludeGroups()([]string)
    GetIncludeAllUsers()(*bool)
    GetIncludeGroups()([]string)
    GetManagementTemplateId()(*string)
    GetTenantGroupId()(*string)
    GetTenantId()(*string)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetExcludeGroups(value []string)()
    SetIncludeAllUsers(value *bool)()
    SetIncludeGroups(value []string)()
    SetManagementTemplateId(value *string)()
    SetTenantGroupId(value *string)()
    SetTenantId(value *string)()
}
