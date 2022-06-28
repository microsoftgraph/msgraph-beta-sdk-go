package windowsupdates

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UpdatableAssetGroup 
type UpdatableAssetGroup struct {
    UpdatableAsset
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Members of the group. Read-only.
    members []UpdatableAssetable
}
// NewUpdatableAssetGroup instantiates a new UpdatableAssetGroup and sets the default values.
func NewUpdatableAssetGroup()(*UpdatableAssetGroup) {
    m := &UpdatableAssetGroup{
        UpdatableAsset: *NewUpdatableAsset(),
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateUpdatableAssetGroupFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUpdatableAssetGroupFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUpdatableAssetGroup(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UpdatableAssetGroup) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UpdatableAssetGroup) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.UpdatableAsset.GetFieldDeserializers()
    res["members"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUpdatableAssetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UpdatableAssetable, len(val))
            for i, v := range val {
                res[i] = v.(UpdatableAssetable)
            }
            m.SetMembers(res)
        }
        return nil
    }
    return res
}
// GetMembers gets the members property value. Members of the group. Read-only.
func (m *UpdatableAssetGroup) GetMembers()([]UpdatableAssetable) {
    if m == nil {
        return nil
    } else {
        return m.members
    }
}
// Serialize serializes information the current object
func (m *UpdatableAssetGroup) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.UpdatableAsset.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetMembers() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetMembers()))
        for i, v := range m.GetMembers() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("members", cast)
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
func (m *UpdatableAssetGroup) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetMembers sets the members property value. Members of the group. Read-only.
func (m *UpdatableAssetGroup) SetMembers(value []UpdatableAssetable)() {
    if m != nil {
        m.members = value
    }
}
