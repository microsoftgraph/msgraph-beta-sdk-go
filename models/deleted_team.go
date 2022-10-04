package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeletedTeam provides operations to manage the collection of accessReviewDecision entities.
type DeletedTeam struct {
    Entity
    // The channels those are either shared with this deleted team or created in this deleted team.
    channels []Channelable
}
// NewDeletedTeam instantiates a new deletedTeam and sets the default values.
func NewDeletedTeam()(*DeletedTeam) {
    m := &DeletedTeam{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.deletedTeam";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateDeletedTeamFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeletedTeamFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeletedTeam(), nil
}
// GetChannels gets the channels property value. The channels those are either shared with this deleted team or created in this deleted team.
func (m *DeletedTeam) GetChannels()([]Channelable) {
    return m.channels
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeletedTeam) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["channels"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateChannelFromDiscriminatorValue , m.SetChannels)
    return res
}
// Serialize serializes information the current object
func (m *DeletedTeam) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetChannels() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetChannels())
        err = writer.WriteCollectionOfObjectValues("channels", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetChannels sets the channels property value. The channels those are either shared with this deleted team or created in this deleted team.
func (m *DeletedTeam) SetChannels(value []Channelable)() {
    m.channels = value
}
