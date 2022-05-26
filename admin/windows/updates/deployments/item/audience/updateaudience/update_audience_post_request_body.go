package updateaudience

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c "github.com/microsoftgraph/msgraph-beta-sdk-go/models/windowsupdates"
)

// UpdateAudiencePostRequestBody provides operations to call the updateAudience method.
type UpdateAudiencePostRequestBody struct {
    // The addExclusions property
    addExclusions []i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.UpdatableAssetable
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The addMembers property
    addMembers []i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.UpdatableAssetable
    // The removeExclusions property
    removeExclusions []i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.UpdatableAssetable
    // The removeMembers property
    removeMembers []i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.UpdatableAssetable
}
// NewUpdateAudiencePostRequestBody instantiates a new updateAudiencePostRequestBody and sets the default values.
func NewUpdateAudiencePostRequestBody()(*UpdateAudiencePostRequestBody) {
    m := &UpdateAudiencePostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateUpdateAudiencePostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUpdateAudiencePostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUpdateAudiencePostRequestBody(), nil
}
// GetAddExclusions gets the addExclusions property value. The addExclusions property
func (m *UpdateAudiencePostRequestBody) GetAddExclusions()([]i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.UpdatableAssetable) {
    if m == nil {
        return nil
    } else {
        return m.addExclusions
    }
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UpdateAudiencePostRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAddMembers gets the addMembers property value. The addMembers property
func (m *UpdateAudiencePostRequestBody) GetAddMembers()([]i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.UpdatableAssetable) {
    if m == nil {
        return nil
    } else {
        return m.addMembers
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UpdateAudiencePostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["addExclusions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.CreateUpdatableAssetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.UpdatableAssetable, len(val))
            for i, v := range val {
                res[i] = v.(i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.UpdatableAssetable)
            }
            m.SetAddExclusions(res)
        }
        return nil
    }
    res["addMembers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.CreateUpdatableAssetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.UpdatableAssetable, len(val))
            for i, v := range val {
                res[i] = v.(i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.UpdatableAssetable)
            }
            m.SetAddMembers(res)
        }
        return nil
    }
    res["removeExclusions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.CreateUpdatableAssetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.UpdatableAssetable, len(val))
            for i, v := range val {
                res[i] = v.(i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.UpdatableAssetable)
            }
            m.SetRemoveExclusions(res)
        }
        return nil
    }
    res["removeMembers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.CreateUpdatableAssetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.UpdatableAssetable, len(val))
            for i, v := range val {
                res[i] = v.(i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.UpdatableAssetable)
            }
            m.SetRemoveMembers(res)
        }
        return nil
    }
    return res
}
// GetRemoveExclusions gets the removeExclusions property value. The removeExclusions property
func (m *UpdateAudiencePostRequestBody) GetRemoveExclusions()([]i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.UpdatableAssetable) {
    if m == nil {
        return nil
    } else {
        return m.removeExclusions
    }
}
// GetRemoveMembers gets the removeMembers property value. The removeMembers property
func (m *UpdateAudiencePostRequestBody) GetRemoveMembers()([]i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.UpdatableAssetable) {
    if m == nil {
        return nil
    } else {
        return m.removeMembers
    }
}
// Serialize serializes information the current object
func (m *UpdateAudiencePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAddExclusions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAddExclusions()))
        for i, v := range m.GetAddExclusions() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("addExclusions", cast)
        if err != nil {
            return err
        }
    }
    if m.GetAddMembers() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAddMembers()))
        for i, v := range m.GetAddMembers() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("addMembers", cast)
        if err != nil {
            return err
        }
    }
    if m.GetRemoveExclusions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRemoveExclusions()))
        for i, v := range m.GetRemoveExclusions() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("removeExclusions", cast)
        if err != nil {
            return err
        }
    }
    if m.GetRemoveMembers() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRemoveMembers()))
        for i, v := range m.GetRemoveMembers() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("removeMembers", cast)
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
// SetAddExclusions sets the addExclusions property value. The addExclusions property
func (m *UpdateAudiencePostRequestBody) SetAddExclusions(value []i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.UpdatableAssetable)() {
    if m != nil {
        m.addExclusions = value
    }
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UpdateAudiencePostRequestBody) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAddMembers sets the addMembers property value. The addMembers property
func (m *UpdateAudiencePostRequestBody) SetAddMembers(value []i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.UpdatableAssetable)() {
    if m != nil {
        m.addMembers = value
    }
}
// SetRemoveExclusions sets the removeExclusions property value. The removeExclusions property
func (m *UpdateAudiencePostRequestBody) SetRemoveExclusions(value []i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.UpdatableAssetable)() {
    if m != nil {
        m.removeExclusions = value
    }
}
// SetRemoveMembers sets the removeMembers property value. The removeMembers property
func (m *UpdateAudiencePostRequestBody) SetRemoveMembers(value []i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.UpdatableAssetable)() {
    if m != nil {
        m.removeMembers = value
    }
}
