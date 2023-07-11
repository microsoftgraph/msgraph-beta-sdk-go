package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SharedEmailDomainInvitationCollectionResponse 
type SharedEmailDomainInvitationCollectionResponse struct {
    BaseCollectionPaginationCountResponse
}
// NewSharedEmailDomainInvitationCollectionResponse instantiates a new sharedEmailDomainInvitationCollectionResponse and sets the default values.
func NewSharedEmailDomainInvitationCollectionResponse()(*SharedEmailDomainInvitationCollectionResponse) {
    m := &SharedEmailDomainInvitationCollectionResponse{
        BaseCollectionPaginationCountResponse: *NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateSharedEmailDomainInvitationCollectionResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSharedEmailDomainInvitationCollectionResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSharedEmailDomainInvitationCollectionResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SharedEmailDomainInvitationCollectionResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSharedEmailDomainInvitationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SharedEmailDomainInvitationable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(SharedEmailDomainInvitationable)
                }
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *SharedEmailDomainInvitationCollectionResponse) GetValue()([]SharedEmailDomainInvitationable) {
    val, err := m.GetBackingStore().Get("value")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]SharedEmailDomainInvitationable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *SharedEmailDomainInvitationCollectionResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.BaseCollectionPaginationCountResponse.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetValue() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetValue()))
        for i, v := range m.GetValue() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("value", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetValue sets the value property value. The value property
func (m *SharedEmailDomainInvitationCollectionResponse) SetValue(value []SharedEmailDomainInvitationable)() {
    err := m.GetBackingStore().Set("value", value)
    if err != nil {
        panic(err)
    }
}
// SharedEmailDomainInvitationCollectionResponseable 
type SharedEmailDomainInvitationCollectionResponseable interface {
    BaseCollectionPaginationCountResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetValue()([]SharedEmailDomainInvitationable)
    SetValue(value []SharedEmailDomainInvitationable)()
}
