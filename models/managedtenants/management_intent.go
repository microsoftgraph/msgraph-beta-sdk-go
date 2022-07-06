package managedtenants

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// ManagementIntent provides operations to manage the collection of accessReviewDecision entities.
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
    return m
}
// CreateManagementIntentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateManagementIntentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewManagementIntent(), nil
}
// GetDisplayName gets the displayName property value. The display name for the management intent. Optional. Read-only.
func (m *ManagementIntent) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ManagementIntent) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
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
    res["isGlobal"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsGlobal(val)
        }
        return nil
    }
    res["managementTemplates"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateManagementTemplateDetailedInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagementTemplateDetailedInfoable, len(val))
            for i, v := range val {
                res[i] = v.(ManagementTemplateDetailedInfoable)
            }
            m.SetManagementTemplates(res)
        }
        return nil
    }
    return res
}
// GetIsGlobal gets the isGlobal property value. A flag indicating whether the management intent is global. Required. Read-only.
func (m *ManagementIntent) GetIsGlobal()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isGlobal
    }
}
// GetManagementTemplates gets the managementTemplates property value. The collection of management templates associated with the management intent. Optional. Read-only.
func (m *ManagementIntent) GetManagementTemplates()([]ManagementTemplateDetailedInfoable) {
    if m == nil {
        return nil
    } else {
        return m.managementTemplates
    }
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
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetManagementTemplates()))
        for i, v := range m.GetManagementTemplates() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("managementTemplates", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDisplayName sets the displayName property value. The display name for the management intent. Optional. Read-only.
func (m *ManagementIntent) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetIsGlobal sets the isGlobal property value. A flag indicating whether the management intent is global. Required. Read-only.
func (m *ManagementIntent) SetIsGlobal(value *bool)() {
    if m != nil {
        m.isGlobal = value
    }
}
// SetManagementTemplates sets the managementTemplates property value. The collection of management templates associated with the management intent. Optional. Read-only.
func (m *ManagementIntent) SetManagementTemplates(value []ManagementTemplateDetailedInfoable)() {
    if m != nil {
        m.managementTemplates = value
    }
}
