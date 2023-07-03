package managedtenants

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// ManagementAction 
type ManagementAction struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entity
}
// NewManagementAction instantiates a new ManagementAction and sets the default values.
func NewManagementAction()(*ManagementAction) {
    m := &ManagementAction{
        Entity: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewEntity(),
    }
    return m
}
// CreateManagementActionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateManagementActionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewManagementAction(), nil
}
// GetCategory gets the category property value. The category property
func (m *ManagementAction) GetCategory()(*ManagementCategory) {
    val, err := m.GetBackingStore().Get("category")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ManagementCategory)
    }
    return nil
}
// GetDescription gets the description property value. The description for the management action. Optional. Read-only.
func (m *ManagementAction) GetDescription()(*string) {
    val, err := m.GetBackingStore().Get("description")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDisplayName gets the displayName property value. The display name for the management action. Optional. Read-only.
func (m *ManagementAction) GetDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("displayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ManagementAction) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["category"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseManagementCategory)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCategory(val.(*ManagementCategory))
        }
        return nil
    }
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["referenceTemplateId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReferenceTemplateId(val)
        }
        return nil
    }
    res["referenceTemplateVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReferenceTemplateVersion(val)
        }
        return nil
    }
    res["workloadActions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWorkloadActionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WorkloadActionable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(WorkloadActionable)
                }
            }
            m.SetWorkloadActions(res)
        }
        return nil
    }
    return res
}
// GetReferenceTemplateId gets the referenceTemplateId property value. The reference for the management template used to generate the management action. Required. Read-only.
func (m *ManagementAction) GetReferenceTemplateId()(*string) {
    val, err := m.GetBackingStore().Get("referenceTemplateId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetReferenceTemplateVersion gets the referenceTemplateVersion property value. The referenceTemplateVersion property
func (m *ManagementAction) GetReferenceTemplateVersion()(*int32) {
    val, err := m.GetBackingStore().Get("referenceTemplateVersion")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetWorkloadActions gets the workloadActions property value. The collection of workload actions associated with the management action. Required. Read-only.
func (m *ManagementAction) GetWorkloadActions()([]WorkloadActionable) {
    val, err := m.GetBackingStore().Get("workloadActions")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]WorkloadActionable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *ManagementAction) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetCategory() != nil {
        cast := (*m.GetCategory()).String()
        err = writer.WriteStringValue("category", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("referenceTemplateId", m.GetReferenceTemplateId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("referenceTemplateVersion", m.GetReferenceTemplateVersion())
        if err != nil {
            return err
        }
    }
    if m.GetWorkloadActions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetWorkloadActions()))
        for i, v := range m.GetWorkloadActions() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("workloadActions", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCategory sets the category property value. The category property
func (m *ManagementAction) SetCategory(value *ManagementCategory)() {
    err := m.GetBackingStore().Set("category", value)
    if err != nil {
        panic(err)
    }
}
// SetDescription sets the description property value. The description for the management action. Optional. Read-only.
func (m *ManagementAction) SetDescription(value *string)() {
    err := m.GetBackingStore().Set("description", value)
    if err != nil {
        panic(err)
    }
}
// SetDisplayName sets the displayName property value. The display name for the management action. Optional. Read-only.
func (m *ManagementAction) SetDisplayName(value *string)() {
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// SetReferenceTemplateId sets the referenceTemplateId property value. The reference for the management template used to generate the management action. Required. Read-only.
func (m *ManagementAction) SetReferenceTemplateId(value *string)() {
    err := m.GetBackingStore().Set("referenceTemplateId", value)
    if err != nil {
        panic(err)
    }
}
// SetReferenceTemplateVersion sets the referenceTemplateVersion property value. The referenceTemplateVersion property
func (m *ManagementAction) SetReferenceTemplateVersion(value *int32)() {
    err := m.GetBackingStore().Set("referenceTemplateVersion", value)
    if err != nil {
        panic(err)
    }
}
// SetWorkloadActions sets the workloadActions property value. The collection of workload actions associated with the management action. Required. Read-only.
func (m *ManagementAction) SetWorkloadActions(value []WorkloadActionable)() {
    err := m.GetBackingStore().Set("workloadActions", value)
    if err != nil {
        panic(err)
    }
}
// ManagementActionable 
type ManagementActionable interface {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCategory()(*ManagementCategory)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetReferenceTemplateId()(*string)
    GetReferenceTemplateVersion()(*int32)
    GetWorkloadActions()([]WorkloadActionable)
    SetCategory(value *ManagementCategory)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetReferenceTemplateId(value *string)()
    SetReferenceTemplateVersion(value *int32)()
    SetWorkloadActions(value []WorkloadActionable)()
}
