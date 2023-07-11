package windowsupdates

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// DeploymentAudience 
type DeploymentAudience struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entity
    // The OdataType property
    OdataType *string
}
// NewDeploymentAudience instantiates a new deploymentAudience and sets the default values.
func NewDeploymentAudience()(*DeploymentAudience) {
    m := &DeploymentAudience{
        Entity: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewEntity(),
    }
    return m
}
// CreateDeploymentAudienceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeploymentAudienceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeploymentAudience(), nil
}
// GetApplicableContent gets the applicableContent property value. Content eligible to deploy to devices in the audience. Not nullable. Read-only.
func (m *DeploymentAudience) GetApplicableContent()([]ApplicableContentable) {
    val, err := m.GetBackingStore().Get("applicableContent")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ApplicableContentable)
    }
    return nil
}
// GetExclusions gets the exclusions property value. Specifies the assets to exclude from the audience.
func (m *DeploymentAudience) GetExclusions()([]UpdatableAssetable) {
    val, err := m.GetBackingStore().Get("exclusions")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]UpdatableAssetable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeploymentAudience) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["applicableContent"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateApplicableContentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ApplicableContentable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ApplicableContentable)
                }
            }
            m.SetApplicableContent(res)
        }
        return nil
    }
    res["exclusions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUpdatableAssetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UpdatableAssetable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(UpdatableAssetable)
                }
            }
            m.SetExclusions(res)
        }
        return nil
    }
    res["members"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUpdatableAssetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UpdatableAssetable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(UpdatableAssetable)
                }
            }
            m.SetMembers(res)
        }
        return nil
    }
    return res
}
// GetMembers gets the members property value. Specifies the assets to include in the audience.
func (m *DeploymentAudience) GetMembers()([]UpdatableAssetable) {
    val, err := m.GetBackingStore().Get("members")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]UpdatableAssetable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *DeploymentAudience) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetApplicableContent() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetApplicableContent()))
        for i, v := range m.GetApplicableContent() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("applicableContent", cast)
        if err != nil {
            return err
        }
    }
    if m.GetExclusions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetExclusions()))
        for i, v := range m.GetExclusions() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("exclusions", cast)
        if err != nil {
            return err
        }
    }
    if m.GetMembers() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetMembers()))
        for i, v := range m.GetMembers() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("members", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetApplicableContent sets the applicableContent property value. Content eligible to deploy to devices in the audience. Not nullable. Read-only.
func (m *DeploymentAudience) SetApplicableContent(value []ApplicableContentable)() {
    err := m.GetBackingStore().Set("applicableContent", value)
    if err != nil {
        panic(err)
    }
}
// SetExclusions sets the exclusions property value. Specifies the assets to exclude from the audience.
func (m *DeploymentAudience) SetExclusions(value []UpdatableAssetable)() {
    err := m.GetBackingStore().Set("exclusions", value)
    if err != nil {
        panic(err)
    }
}
// SetMembers sets the members property value. Specifies the assets to include in the audience.
func (m *DeploymentAudience) SetMembers(value []UpdatableAssetable)() {
    err := m.GetBackingStore().Set("members", value)
    if err != nil {
        panic(err)
    }
}
// DeploymentAudienceable 
type DeploymentAudienceable interface {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetApplicableContent()([]ApplicableContentable)
    GetExclusions()([]UpdatableAssetable)
    GetMembers()([]UpdatableAssetable)
    SetApplicableContent(value []ApplicableContentable)()
    SetExclusions(value []UpdatableAssetable)()
    SetMembers(value []UpdatableAssetable)()
}
