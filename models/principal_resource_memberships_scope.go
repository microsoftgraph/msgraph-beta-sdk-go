package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PrincipalResourceMembershipsScope 
type PrincipalResourceMembershipsScope struct {
    AccessReviewScope
}
// NewPrincipalResourceMembershipsScope instantiates a new principalResourceMembershipsScope and sets the default values.
func NewPrincipalResourceMembershipsScope()(*PrincipalResourceMembershipsScope) {
    m := &PrincipalResourceMembershipsScope{
        AccessReviewScope: *NewAccessReviewScope(),
    }
    odataTypeValue := "#microsoft.graph.principalResourceMembershipsScope"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreatePrincipalResourceMembershipsScopeFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePrincipalResourceMembershipsScopeFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPrincipalResourceMembershipsScope(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PrincipalResourceMembershipsScope) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AccessReviewScope.GetFieldDeserializers()
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    res["principalScopes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAccessReviewScopeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AccessReviewScopeable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(AccessReviewScopeable)
                }
            }
            m.SetPrincipalScopes(res)
        }
        return nil
    }
    res["resourceScopes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAccessReviewScopeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AccessReviewScopeable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(AccessReviewScopeable)
                }
            }
            m.SetResourceScopes(res)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *PrincipalResourceMembershipsScope) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPrincipalScopes gets the principalScopes property value. Defines the scopes of the principals whose access to resources are reviewed in the access review.
func (m *PrincipalResourceMembershipsScope) GetPrincipalScopes()([]AccessReviewScopeable) {
    val, err := m.GetBackingStore().Get("principalScopes")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]AccessReviewScopeable)
    }
    return nil
}
// GetResourceScopes gets the resourceScopes property value. Defines the scopes of the resources for which access is reviewed.
func (m *PrincipalResourceMembershipsScope) GetResourceScopes()([]AccessReviewScopeable) {
    val, err := m.GetBackingStore().Get("resourceScopes")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]AccessReviewScopeable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *PrincipalResourceMembershipsScope) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AccessReviewScope.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    if m.GetPrincipalScopes() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetPrincipalScopes()))
        for i, v := range m.GetPrincipalScopes() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("principalScopes", cast)
        if err != nil {
            return err
        }
    }
    if m.GetResourceScopes() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetResourceScopes()))
        for i, v := range m.GetResourceScopes() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("resourceScopes", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *PrincipalResourceMembershipsScope) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetPrincipalScopes sets the principalScopes property value. Defines the scopes of the principals whose access to resources are reviewed in the access review.
func (m *PrincipalResourceMembershipsScope) SetPrincipalScopes(value []AccessReviewScopeable)() {
    err := m.GetBackingStore().Set("principalScopes", value)
    if err != nil {
        panic(err)
    }
}
// SetResourceScopes sets the resourceScopes property value. Defines the scopes of the resources for which access is reviewed.
func (m *PrincipalResourceMembershipsScope) SetResourceScopes(value []AccessReviewScopeable)() {
    err := m.GetBackingStore().Set("resourceScopes", value)
    if err != nil {
        panic(err)
    }
}
// PrincipalResourceMembershipsScopeable 
type PrincipalResourceMembershipsScopeable interface {
    AccessReviewScopeable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetOdataType()(*string)
    GetPrincipalScopes()([]AccessReviewScopeable)
    GetResourceScopes()([]AccessReviewScopeable)
    SetOdataType(value *string)()
    SetPrincipalScopes(value []AccessReviewScopeable)()
    SetResourceScopes(value []AccessReviewScopeable)()
}
