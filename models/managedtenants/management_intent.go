package managedtenants

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// ManagementIntent provides operations to manage the collection of activityStatistics entities.
type ManagementIntent struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entity
    // The display name for the management intent. Optional. Read-only.
    displayName *string
    // A flag indicating whether the management intent is global. Required. Read-only.
    isGlobal *bool
    // The collection of management templates associated with the management intent. Optional. Read-only.
    managementTemplates []ManagementTemplateDetailedInfoable
}
// NewManagementIntent instantiates a new managementIntent and sets the default values.
func NewManagementIntent()(*ManagementIntent) {
    m := &ManagementIntent{
        Entity: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.managedTenants.managementIntent";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateManagementIntentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateManagementIntentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewManagementIntent(), nil
}
// GetDisplayName gets the displayName property value. The display name for the management intent. Optional. Read-only.
func (m *ManagementIntent) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ManagementIntent) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["displayName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDisplayName)
    res["isGlobal"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIsGlobal)
    res["managementTemplates"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateManagementTemplateDetailedInfoFromDiscriminatorValue , m.SetManagementTemplates)
    return res
}
// GetIsGlobal gets the isGlobal property value. A flag indicating whether the management intent is global. Required. Read-only.
func (m *ManagementIntent) GetIsGlobal()(*bool) {
    return m.isGlobal
}
// GetManagementTemplates gets the managementTemplates property value. The collection of management templates associated with the management intent. Optional. Read-only.
func (m *ManagementIntent) GetManagementTemplates()([]ManagementTemplateDetailedInfoable) {
    return m.managementTemplates
}
// Serialize serializes information the current object
func (m *ManagementIntent) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isGlobal", m.GetIsGlobal())
        if err != nil {
            return err
        }
    }
    if m.GetManagementTemplates() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetManagementTemplates())
        err = writer.WriteCollectionOfObjectValues("managementTemplates", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDisplayName sets the displayName property value. The display name for the management intent. Optional. Read-only.
func (m *ManagementIntent) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetIsGlobal sets the isGlobal property value. A flag indicating whether the management intent is global. Required. Read-only.
func (m *ManagementIntent) SetIsGlobal(value *bool)() {
    m.isGlobal = value
}
// SetManagementTemplates sets the managementTemplates property value. The collection of management templates associated with the management intent. Optional. Read-only.
func (m *ManagementIntent) SetManagementTemplates(value []ManagementTemplateDetailedInfoable)() {
    m.managementTemplates = value
}
