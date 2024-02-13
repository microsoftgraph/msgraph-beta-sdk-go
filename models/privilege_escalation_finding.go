package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type PrivilegeEscalationFinding struct {
    Finding
}
// NewPrivilegeEscalationFinding instantiates a new PrivilegeEscalationFinding and sets the default values.
func NewPrivilegeEscalationFinding()(*PrivilegeEscalationFinding) {
    m := &PrivilegeEscalationFinding{
        Finding: *NewFinding(),
    }
    return m
}
// CreatePrivilegeEscalationFindingFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreatePrivilegeEscalationFindingFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "#microsoft.graph.privilegeEscalationAwsResourceFinding":
                        return NewPrivilegeEscalationAwsResourceFinding(), nil
                    case "#microsoft.graph.privilegeEscalationAwsRoleFinding":
                        return NewPrivilegeEscalationAwsRoleFinding(), nil
                    case "#microsoft.graph.privilegeEscalationGcpServiceAccountFinding":
                        return NewPrivilegeEscalationGcpServiceAccountFinding(), nil
                    case "#microsoft.graph.privilegeEscalationUserFinding":
                        return NewPrivilegeEscalationUserFinding(), nil
                }
            }
        }
    }
    return NewPrivilegeEscalationFinding(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *PrivilegeEscalationFinding) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Finding.GetFieldDeserializers()
    res["identity"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAuthorizationSystemIdentityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIdentity(val.(AuthorizationSystemIdentityable))
        }
        return nil
    }
    res["identityDetails"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentityDetailsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIdentityDetails(val.(IdentityDetailsable))
        }
        return nil
    }
    res["permissionsCreepIndex"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePermissionsCreepIndexFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPermissionsCreepIndex(val.(PermissionsCreepIndexable))
        }
        return nil
    }
    res["privilegeEscalationDetails"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePrivilegeEscalationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PrivilegeEscalationable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(PrivilegeEscalationable)
                }
            }
            m.SetPrivilegeEscalationDetails(res)
        }
        return nil
    }
    return res
}
// GetIdentity gets the identity property value. The identity property
// returns a AuthorizationSystemIdentityable when successful
func (m *PrivilegeEscalationFinding) GetIdentity()(AuthorizationSystemIdentityable) {
    val, err := m.GetBackingStore().Get("identity")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(AuthorizationSystemIdentityable)
    }
    return nil
}
// GetIdentityDetails gets the identityDetails property value. An identity's information details. Inherited from finding.
// returns a IdentityDetailsable when successful
func (m *PrivilegeEscalationFinding) GetIdentityDetails()(IdentityDetailsable) {
    val, err := m.GetBackingStore().Get("identityDetails")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(IdentityDetailsable)
    }
    return nil
}
// GetPermissionsCreepIndex gets the permissionsCreepIndex property value. The permissionsCreepIndex property
// returns a PermissionsCreepIndexable when successful
func (m *PrivilegeEscalationFinding) GetPermissionsCreepIndex()(PermissionsCreepIndexable) {
    val, err := m.GetBackingStore().Get("permissionsCreepIndex")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(PermissionsCreepIndexable)
    }
    return nil
}
// GetPrivilegeEscalationDetails gets the privilegeEscalationDetails property value. The list of escalations that the identity is capable of performing.
// returns a []PrivilegeEscalationable when successful
func (m *PrivilegeEscalationFinding) GetPrivilegeEscalationDetails()([]PrivilegeEscalationable) {
    val, err := m.GetBackingStore().Get("privilegeEscalationDetails")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]PrivilegeEscalationable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *PrivilegeEscalationFinding) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Finding.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("identity", m.GetIdentity())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("identityDetails", m.GetIdentityDetails())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("permissionsCreepIndex", m.GetPermissionsCreepIndex())
        if err != nil {
            return err
        }
    }
    if m.GetPrivilegeEscalationDetails() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetPrivilegeEscalationDetails()))
        for i, v := range m.GetPrivilegeEscalationDetails() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("privilegeEscalationDetails", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetIdentity sets the identity property value. The identity property
func (m *PrivilegeEscalationFinding) SetIdentity(value AuthorizationSystemIdentityable)() {
    err := m.GetBackingStore().Set("identity", value)
    if err != nil {
        panic(err)
    }
}
// SetIdentityDetails sets the identityDetails property value. An identity's information details. Inherited from finding.
func (m *PrivilegeEscalationFinding) SetIdentityDetails(value IdentityDetailsable)() {
    err := m.GetBackingStore().Set("identityDetails", value)
    if err != nil {
        panic(err)
    }
}
// SetPermissionsCreepIndex sets the permissionsCreepIndex property value. The permissionsCreepIndex property
func (m *PrivilegeEscalationFinding) SetPermissionsCreepIndex(value PermissionsCreepIndexable)() {
    err := m.GetBackingStore().Set("permissionsCreepIndex", value)
    if err != nil {
        panic(err)
    }
}
// SetPrivilegeEscalationDetails sets the privilegeEscalationDetails property value. The list of escalations that the identity is capable of performing.
func (m *PrivilegeEscalationFinding) SetPrivilegeEscalationDetails(value []PrivilegeEscalationable)() {
    err := m.GetBackingStore().Set("privilegeEscalationDetails", value)
    if err != nil {
        panic(err)
    }
}
type PrivilegeEscalationFindingable interface {
    Findingable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetIdentity()(AuthorizationSystemIdentityable)
    GetIdentityDetails()(IdentityDetailsable)
    GetPermissionsCreepIndex()(PermissionsCreepIndexable)
    GetPrivilegeEscalationDetails()([]PrivilegeEscalationable)
    SetIdentity(value AuthorizationSystemIdentityable)()
    SetIdentityDetails(value IdentityDetailsable)()
    SetPermissionsCreepIndex(value PermissionsCreepIndexable)()
    SetPrivilegeEscalationDetails(value []PrivilegeEscalationable)()
}
