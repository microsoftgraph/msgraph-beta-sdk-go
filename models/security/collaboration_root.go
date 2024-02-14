package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

type CollaborationRoot struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entity
}
// NewCollaborationRoot instantiates a new CollaborationRoot and sets the default values.
func NewCollaborationRoot()(*CollaborationRoot) {
    m := &CollaborationRoot{
        Entity: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewEntity(),
    }
    return m
}
// CreateCollaborationRootFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCollaborationRootFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCollaborationRoot(), nil
}
// GetAnalyzedEmails gets the analyzedEmails property value. The analyzedEmails property
// returns a []AnalyzedEmailable when successful
func (m *CollaborationRoot) GetAnalyzedEmails()([]AnalyzedEmailable) {
    val, err := m.GetBackingStore().Get("analyzedEmails")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]AnalyzedEmailable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CollaborationRoot) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["analyzedEmails"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAnalyzedEmailFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AnalyzedEmailable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(AnalyzedEmailable)
                }
            }
            m.SetAnalyzedEmails(res)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *CollaborationRoot) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAnalyzedEmails() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAnalyzedEmails()))
        for i, v := range m.GetAnalyzedEmails() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("analyzedEmails", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAnalyzedEmails sets the analyzedEmails property value. The analyzedEmails property
func (m *CollaborationRoot) SetAnalyzedEmails(value []AnalyzedEmailable)() {
    err := m.GetBackingStore().Set("analyzedEmails", value)
    if err != nil {
        panic(err)
    }
}
type CollaborationRootable interface {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAnalyzedEmails()([]AnalyzedEmailable)
    SetAnalyzedEmails(value []AnalyzedEmailable)()
}
