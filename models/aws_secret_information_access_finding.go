package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type AwsSecretInformationAccessFinding struct {
    Finding
}
// NewAwsSecretInformationAccessFinding instantiates a new AwsSecretInformationAccessFinding and sets the default values.
func NewAwsSecretInformationAccessFinding()(*AwsSecretInformationAccessFinding) {
    m := &AwsSecretInformationAccessFinding{
        Finding: *NewFinding(),
    }
    return m
}
// CreateAwsSecretInformationAccessFindingFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAwsSecretInformationAccessFindingFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.secretInformationAccessAwsResourceFinding":
                        return NewSecretInformationAccessAwsResourceFinding(), nil
                    case "#microsoft.graph.secretInformationAccessAwsRoleFinding":
                        return NewSecretInformationAccessAwsRoleFinding(), nil
                    case "#microsoft.graph.secretInformationAccessAwsServerlessFunctionFinding":
                        return NewSecretInformationAccessAwsServerlessFunctionFinding(), nil
                    case "#microsoft.graph.secretInformationAccessAwsUserFinding":
                        return NewSecretInformationAccessAwsUserFinding(), nil
                }
            }
        }
    }
    return NewAwsSecretInformationAccessFinding(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AwsSecretInformationAccessFinding) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["secretInformationWebServices"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAwsSecretInformationWebServices)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecretInformationWebServices(val.(*AwsSecretInformationWebServices))
        }
        return nil
    }
    return res
}
// GetIdentity gets the identity property value. The identity property
// returns a AuthorizationSystemIdentityable when successful
func (m *AwsSecretInformationAccessFinding) GetIdentity()(AuthorizationSystemIdentityable) {
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
func (m *AwsSecretInformationAccessFinding) GetIdentityDetails()(IdentityDetailsable) {
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
func (m *AwsSecretInformationAccessFinding) GetPermissionsCreepIndex()(PermissionsCreepIndexable) {
    val, err := m.GetBackingStore().Get("permissionsCreepIndex")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(PermissionsCreepIndexable)
    }
    return nil
}
// GetSecretInformationWebServices gets the secretInformationWebServices property value. The secretInformationWebServices property
// returns a *AwsSecretInformationWebServices when successful
func (m *AwsSecretInformationAccessFinding) GetSecretInformationWebServices()(*AwsSecretInformationWebServices) {
    val, err := m.GetBackingStore().Get("secretInformationWebServices")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*AwsSecretInformationWebServices)
    }
    return nil
}
// Serialize serializes information the current object
func (m *AwsSecretInformationAccessFinding) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    if m.GetSecretInformationWebServices() != nil {
        cast := (*m.GetSecretInformationWebServices()).String()
        err = writer.WriteStringValue("secretInformationWebServices", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetIdentity sets the identity property value. The identity property
func (m *AwsSecretInformationAccessFinding) SetIdentity(value AuthorizationSystemIdentityable)() {
    err := m.GetBackingStore().Set("identity", value)
    if err != nil {
        panic(err)
    }
}
// SetIdentityDetails sets the identityDetails property value. The identityDetails property
func (m *AwsSecretInformationAccessFinding) SetIdentityDetails(value IdentityDetailsable)() {
    err := m.GetBackingStore().Set("identityDetails", value)
    if err != nil {
        panic(err)
    }
}
// SetPermissionsCreepIndex sets the permissionsCreepIndex property value. The permissionsCreepIndex property
func (m *AwsSecretInformationAccessFinding) SetPermissionsCreepIndex(value PermissionsCreepIndexable)() {
    err := m.GetBackingStore().Set("permissionsCreepIndex", value)
    if err != nil {
        panic(err)
    }
}
// SetSecretInformationWebServices sets the secretInformationWebServices property value. The secretInformationWebServices property
func (m *AwsSecretInformationAccessFinding) SetSecretInformationWebServices(value *AwsSecretInformationWebServices)() {
    err := m.GetBackingStore().Set("secretInformationWebServices", value)
    if err != nil {
        panic(err)
    }
}
type AwsSecretInformationAccessFindingable interface {
    Findingable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetIdentity()(AuthorizationSystemIdentityable)
    GetIdentityDetails()(IdentityDetailsable)
    GetPermissionsCreepIndex()(PermissionsCreepIndexable)
    GetSecretInformationWebServices()(*AwsSecretInformationWebServices)
    SetIdentity(value AuthorizationSystemIdentityable)()
    SetIdentityDetails(value IdentityDetailsable)()
    SetPermissionsCreepIndex(value PermissionsCreepIndexable)()
    SetSecretInformationWebServices(value *AwsSecretInformationWebServices)()
}
