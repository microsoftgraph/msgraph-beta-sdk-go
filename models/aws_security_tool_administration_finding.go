package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type AwsSecurityToolAdministrationFinding struct {
    Finding
}
// NewAwsSecurityToolAdministrationFinding instantiates a new AwsSecurityToolAdministrationFinding and sets the default values.
func NewAwsSecurityToolAdministrationFinding()(*AwsSecurityToolAdministrationFinding) {
    m := &AwsSecurityToolAdministrationFinding{
        Finding: *NewFinding(),
    }
    return m
}
// CreateAwsSecurityToolAdministrationFindingFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAwsSecurityToolAdministrationFindingFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.securityToolAwsResourceAdministratorFinding":
                        return NewSecurityToolAwsResourceAdministratorFinding(), nil
                    case "#microsoft.graph.securityToolAwsRoleAdministratorFinding":
                        return NewSecurityToolAwsRoleAdministratorFinding(), nil
                    case "#microsoft.graph.securityToolAwsServerlessFunctionAdministratorFinding":
                        return NewSecurityToolAwsServerlessFunctionAdministratorFinding(), nil
                    case "#microsoft.graph.securityToolAwsUserAdministratorFinding":
                        return NewSecurityToolAwsUserAdministratorFinding(), nil
                }
            }
        }
    }
    return NewAwsSecurityToolAdministrationFinding(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AwsSecurityToolAdministrationFinding) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["securityTools"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAwsSecurityToolWebServices)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecurityTools(val.(*AwsSecurityToolWebServices))
        }
        return nil
    }
    return res
}
// GetIdentity gets the identity property value. The identity property
// returns a AuthorizationSystemIdentityable when successful
func (m *AwsSecurityToolAdministrationFinding) GetIdentity()(AuthorizationSystemIdentityable) {
    val, err := m.GetBackingStore().Get("identity")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(AuthorizationSystemIdentityable)
    }
    return nil
}
// GetIdentityDetails gets the identityDetails property value. The identityDetails property
// returns a IdentityDetailsable when successful
func (m *AwsSecurityToolAdministrationFinding) GetIdentityDetails()(IdentityDetailsable) {
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
func (m *AwsSecurityToolAdministrationFinding) GetPermissionsCreepIndex()(PermissionsCreepIndexable) {
    val, err := m.GetBackingStore().Get("permissionsCreepIndex")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(PermissionsCreepIndexable)
    }
    return nil
}
// GetSecurityTools gets the securityTools property value. The securityTools property
// returns a *AwsSecurityToolWebServices when successful
func (m *AwsSecurityToolAdministrationFinding) GetSecurityTools()(*AwsSecurityToolWebServices) {
    val, err := m.GetBackingStore().Get("securityTools")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*AwsSecurityToolWebServices)
    }
    return nil
}
// Serialize serializes information the current object
func (m *AwsSecurityToolAdministrationFinding) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    if m.GetSecurityTools() != nil {
        cast := (*m.GetSecurityTools()).String()
        err = writer.WriteStringValue("securityTools", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetIdentity sets the identity property value. The identity property
func (m *AwsSecurityToolAdministrationFinding) SetIdentity(value AuthorizationSystemIdentityable)() {
    err := m.GetBackingStore().Set("identity", value)
    if err != nil {
        panic(err)
    }
}
// SetIdentityDetails sets the identityDetails property value. The identityDetails property
func (m *AwsSecurityToolAdministrationFinding) SetIdentityDetails(value IdentityDetailsable)() {
    err := m.GetBackingStore().Set("identityDetails", value)
    if err != nil {
        panic(err)
    }
}
// SetPermissionsCreepIndex sets the permissionsCreepIndex property value. The permissionsCreepIndex property
func (m *AwsSecurityToolAdministrationFinding) SetPermissionsCreepIndex(value PermissionsCreepIndexable)() {
    err := m.GetBackingStore().Set("permissionsCreepIndex", value)
    if err != nil {
        panic(err)
    }
}
// SetSecurityTools sets the securityTools property value. The securityTools property
func (m *AwsSecurityToolAdministrationFinding) SetSecurityTools(value *AwsSecurityToolWebServices)() {
    err := m.GetBackingStore().Set("securityTools", value)
    if err != nil {
        panic(err)
    }
}
type AwsSecurityToolAdministrationFindingable interface {
    Findingable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetIdentity()(AuthorizationSystemIdentityable)
    GetIdentityDetails()(IdentityDetailsable)
    GetPermissionsCreepIndex()(PermissionsCreepIndexable)
    GetSecurityTools()(*AwsSecurityToolWebServices)
    SetIdentity(value AuthorizationSystemIdentityable)()
    SetIdentityDetails(value IdentityDetailsable)()
    SetPermissionsCreepIndex(value PermissionsCreepIndexable)()
    SetSecurityTools(value *AwsSecurityToolWebServices)()
}
