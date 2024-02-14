package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type IdentityFinding struct {
    Finding
}
// NewIdentityFinding instantiates a new IdentityFinding and sets the default values.
func NewIdentityFinding()(*IdentityFinding) {
    m := &IdentityFinding{
        Finding: *NewFinding(),
    }
    return m
}
// CreateIdentityFindingFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateIdentityFindingFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.inactiveAwsResourceFinding":
                        return NewInactiveAwsResourceFinding(), nil
                    case "#microsoft.graph.inactiveAwsRoleFinding":
                        return NewInactiveAwsRoleFinding(), nil
                    case "#microsoft.graph.inactiveAzureServicePrincipalFinding":
                        return NewInactiveAzureServicePrincipalFinding(), nil
                    case "#microsoft.graph.inactiveGcpServiceAccountFinding":
                        return NewInactiveGcpServiceAccountFinding(), nil
                    case "#microsoft.graph.inactiveServerlessFunctionFinding":
                        return NewInactiveServerlessFunctionFinding(), nil
                    case "#microsoft.graph.inactiveUserFinding":
                        return NewInactiveUserFinding(), nil
                    case "#microsoft.graph.overprovisionedAwsResourceFinding":
                        return NewOverprovisionedAwsResourceFinding(), nil
                    case "#microsoft.graph.overprovisionedAwsRoleFinding":
                        return NewOverprovisionedAwsRoleFinding(), nil
                    case "#microsoft.graph.overprovisionedAzureServicePrincipalFinding":
                        return NewOverprovisionedAzureServicePrincipalFinding(), nil
                    case "#microsoft.graph.overprovisionedGcpServiceAccountFinding":
                        return NewOverprovisionedGcpServiceAccountFinding(), nil
                    case "#microsoft.graph.overprovisionedServerlessFunctionFinding":
                        return NewOverprovisionedServerlessFunctionFinding(), nil
                    case "#microsoft.graph.overprovisionedUserFinding":
                        return NewOverprovisionedUserFinding(), nil
                    case "#microsoft.graph.superAwsResourceFinding":
                        return NewSuperAwsResourceFinding(), nil
                    case "#microsoft.graph.superAwsRoleFinding":
                        return NewSuperAwsRoleFinding(), nil
                    case "#microsoft.graph.superAzureServicePrincipalFinding":
                        return NewSuperAzureServicePrincipalFinding(), nil
                    case "#microsoft.graph.superGcpServiceAccountFinding":
                        return NewSuperGcpServiceAccountFinding(), nil
                    case "#microsoft.graph.superServerlessFunctionFinding":
                        return NewSuperServerlessFunctionFinding(), nil
                    case "#microsoft.graph.superUserFinding":
                        return NewSuperUserFinding(), nil
                    case "#microsoft.graph.unenforcedMfaAwsUserFinding":
                        return NewUnenforcedMfaAwsUserFinding(), nil
                }
            }
        }
    }
    return NewIdentityFinding(), nil
}
// GetActionSummary gets the actionSummary property value. The actionSummary property
// returns a ActionSummaryable when successful
func (m *IdentityFinding) GetActionSummary()(ActionSummaryable) {
    val, err := m.GetBackingStore().Get("actionSummary")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(ActionSummaryable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *IdentityFinding) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Finding.GetFieldDeserializers()
    res["actionSummary"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateActionSummaryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActionSummary(val.(ActionSummaryable))
        }
        return nil
    }
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
    return res
}
// GetIdentity gets the identity property value. The identity property
// returns a AuthorizationSystemIdentityable when successful
func (m *IdentityFinding) GetIdentity()(AuthorizationSystemIdentityable) {
    val, err := m.GetBackingStore().Get("identity")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(AuthorizationSystemIdentityable)
    }
    return nil
}
// GetIdentityDetails gets the identityDetails property value. An identity's information details.
// returns a IdentityDetailsable when successful
func (m *IdentityFinding) GetIdentityDetails()(IdentityDetailsable) {
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
func (m *IdentityFinding) GetPermissionsCreepIndex()(PermissionsCreepIndexable) {
    val, err := m.GetBackingStore().Get("permissionsCreepIndex")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(PermissionsCreepIndexable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *IdentityFinding) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Finding.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("actionSummary", m.GetActionSummary())
        if err != nil {
            return err
        }
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
    return nil
}
// SetActionSummary sets the actionSummary property value. The actionSummary property
func (m *IdentityFinding) SetActionSummary(value ActionSummaryable)() {
    err := m.GetBackingStore().Set("actionSummary", value)
    if err != nil {
        panic(err)
    }
}
// SetIdentity sets the identity property value. The identity property
func (m *IdentityFinding) SetIdentity(value AuthorizationSystemIdentityable)() {
    err := m.GetBackingStore().Set("identity", value)
    if err != nil {
        panic(err)
    }
}
// SetIdentityDetails sets the identityDetails property value. An identity's information details.
func (m *IdentityFinding) SetIdentityDetails(value IdentityDetailsable)() {
    err := m.GetBackingStore().Set("identityDetails", value)
    if err != nil {
        panic(err)
    }
}
// SetPermissionsCreepIndex sets the permissionsCreepIndex property value. The permissionsCreepIndex property
func (m *IdentityFinding) SetPermissionsCreepIndex(value PermissionsCreepIndexable)() {
    err := m.GetBackingStore().Set("permissionsCreepIndex", value)
    if err != nil {
        panic(err)
    }
}
type IdentityFindingable interface {
    Findingable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActionSummary()(ActionSummaryable)
    GetIdentity()(AuthorizationSystemIdentityable)
    GetIdentityDetails()(IdentityDetailsable)
    GetPermissionsCreepIndex()(PermissionsCreepIndexable)
    SetActionSummary(value ActionSummaryable)()
    SetIdentity(value AuthorizationSystemIdentityable)()
    SetIdentityDetails(value IdentityDetailsable)()
    SetPermissionsCreepIndex(value PermissionsCreepIndexable)()
}
