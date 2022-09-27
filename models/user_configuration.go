package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserConfiguration provides operations to manage the collection of accessReviewDecision entities.
type UserConfiguration struct {
    Entity
    // The binaryData property
    binaryData []byte
}
// NewUserConfiguration instantiates a new userConfiguration and sets the default values.
func NewUserConfiguration()(*UserConfiguration) {
    m := &UserConfiguration{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.userConfiguration";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateUserConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUserConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUserConfiguration(), nil
}
// GetBinaryData gets the binaryData property value. The binaryData property
func (m *UserConfiguration) GetBinaryData()([]byte) {
    return m.binaryData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["binaryData"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetByteArrayValue(m.SetBinaryData)
    return res
}
// Serialize serializes information the current object
func (m *UserConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteByteArrayValue("binaryData", m.GetBinaryData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetBinaryData sets the binaryData property value. The binaryData property
func (m *UserConfiguration) SetBinaryData(value []byte)() {
    m.binaryData = value
}
