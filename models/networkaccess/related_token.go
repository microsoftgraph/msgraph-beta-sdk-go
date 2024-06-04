package networkaccess

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type RelatedToken struct {
    RelatedResource
}
// NewRelatedToken instantiates a new RelatedToken and sets the default values.
func NewRelatedToken()(*RelatedToken) {
    m := &RelatedToken{
        RelatedResource: *NewRelatedResource(),
    }
    odataTypeValue := "#microsoft.graph.networkaccess.relatedToken"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateRelatedTokenFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRelatedTokenFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRelatedToken(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *RelatedToken) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.RelatedResource.GetFieldDeserializers()
    res["uniqueTokenIdentifier"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUniqueTokenIdentifier(val)
        }
        return nil
    }
    return res
}
// GetUniqueTokenIdentifier gets the uniqueTokenIdentifier property value. The uniqueTokenIdentifier property
// returns a *string when successful
func (m *RelatedToken) GetUniqueTokenIdentifier()(*string) {
    val, err := m.GetBackingStore().Get("uniqueTokenIdentifier")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *RelatedToken) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.RelatedResource.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("uniqueTokenIdentifier", m.GetUniqueTokenIdentifier())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetUniqueTokenIdentifier sets the uniqueTokenIdentifier property value. The uniqueTokenIdentifier property
func (m *RelatedToken) SetUniqueTokenIdentifier(value *string)() {
    err := m.GetBackingStore().Set("uniqueTokenIdentifier", value)
    if err != nil {
        panic(err)
    }
}
type RelatedTokenable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    RelatedResourceable
    GetUniqueTokenIdentifier()(*string)
    SetUniqueTokenIdentifier(value *string)()
}
