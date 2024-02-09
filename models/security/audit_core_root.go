package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

type AuditCoreRoot struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entity
}
// NewAuditCoreRoot instantiates a new AuditCoreRoot and sets the default values.
func NewAuditCoreRoot()(*AuditCoreRoot) {
    m := &AuditCoreRoot{
        Entity: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewEntity(),
    }
    return m
}
// CreateAuditCoreRootFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAuditCoreRootFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAuditCoreRoot(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AuditCoreRoot) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["queries"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAuditLogQueryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AuditLogQueryable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(AuditLogQueryable)
                }
            }
            m.SetQueries(res)
        }
        return nil
    }
    return res
}
// GetQueries gets the queries property value. The queries property
// returns a []AuditLogQueryable when successful
func (m *AuditCoreRoot) GetQueries()([]AuditLogQueryable) {
    val, err := m.GetBackingStore().Get("queries")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]AuditLogQueryable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *AuditCoreRoot) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetQueries() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetQueries()))
        for i, v := range m.GetQueries() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("queries", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetQueries sets the queries property value. The queries property
func (m *AuditCoreRoot) SetQueries(value []AuditLogQueryable)() {
    err := m.GetBackingStore().Set("queries", value)
    if err != nil {
        panic(err)
    }
}
type AuditCoreRootable interface {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetQueries()([]AuditLogQueryable)
    SetQueries(value []AuditLogQueryable)()
}
