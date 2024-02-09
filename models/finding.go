package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type Finding struct {
    Entity
}
// NewFinding instantiates a new Finding and sets the default values.
func NewFinding()(*Finding) {
    m := &Finding{
        Entity: *NewEntity(),
    }
    return m
}
// CreateFindingFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateFindingFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.awsExternalSystemAccessFinding":
                        return NewAwsExternalSystemAccessFinding(), nil
                    case "#microsoft.graph.awsExternalSystemAccessRoleFinding":
                        return NewAwsExternalSystemAccessRoleFinding(), nil
                    case "#microsoft.graph.awsIdentityAccessManagementKeyAgeFinding":
                        return NewAwsIdentityAccessManagementKeyAgeFinding(), nil
                    case "#microsoft.graph.awsIdentityAccessManagementKeyUsageFinding":
                        return NewAwsIdentityAccessManagementKeyUsageFinding(), nil
                    case "#microsoft.graph.awsSecretInformationAccessFinding":
                        return NewAwsSecretInformationAccessFinding(), nil
                    case "#microsoft.graph.awsSecurityToolAdministrationFinding":
                        return NewAwsSecurityToolAdministrationFinding(), nil
                    case "#microsoft.graph.encryptedAwsStorageBucketFinding":
                        return NewEncryptedAwsStorageBucketFinding(), nil
                    case "#microsoft.graph.encryptedAzureStorageAccountFinding":
                        return NewEncryptedAzureStorageAccountFinding(), nil
                    case "#microsoft.graph.encryptedGcpStorageBucketFinding":
                        return NewEncryptedGcpStorageBucketFinding(), nil
                    case "#microsoft.graph.externallyAccessibleAwsStorageBucketFinding":
                        return NewExternallyAccessibleAwsStorageBucketFinding(), nil
                    case "#microsoft.graph.externallyAccessibleAzureBlobContainerFinding":
                        return NewExternallyAccessibleAzureBlobContainerFinding(), nil
                    case "#microsoft.graph.externallyAccessibleGcpStorageBucketFinding":
                        return NewExternallyAccessibleGcpStorageBucketFinding(), nil
                    case "#microsoft.graph.identityFinding":
                        return NewIdentityFinding(), nil
                    case "#microsoft.graph.inactiveAwsResourceFinding":
                        return NewInactiveAwsResourceFinding(), nil
                    case "#microsoft.graph.inactiveAwsRoleFinding":
                        return NewInactiveAwsRoleFinding(), nil
                    case "#microsoft.graph.inactiveAzureServicePrincipalFinding":
                        return NewInactiveAzureServicePrincipalFinding(), nil
                    case "#microsoft.graph.inactiveGcpServiceAccountFinding":
                        return NewInactiveGcpServiceAccountFinding(), nil
                    case "#microsoft.graph.inactiveGroupFinding":
                        return NewInactiveGroupFinding(), nil
                    case "#microsoft.graph.inactiveServerlessFunctionFinding":
                        return NewInactiveServerlessFunctionFinding(), nil
                    case "#microsoft.graph.inactiveUserFinding":
                        return NewInactiveUserFinding(), nil
                    case "#microsoft.graph.openAwsSecurityGroupFinding":
                        return NewOpenAwsSecurityGroupFinding(), nil
                    case "#microsoft.graph.openNetworkAzureSecurityGroupFinding":
                        return NewOpenNetworkAzureSecurityGroupFinding(), nil
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
                    case "#microsoft.graph.privilegeEscalationAwsResourceFinding":
                        return NewPrivilegeEscalationAwsResourceFinding(), nil
                    case "#microsoft.graph.privilegeEscalationAwsRoleFinding":
                        return NewPrivilegeEscalationAwsRoleFinding(), nil
                    case "#microsoft.graph.privilegeEscalationFinding":
                        return NewPrivilegeEscalationFinding(), nil
                    case "#microsoft.graph.privilegeEscalationGcpServiceAccountFinding":
                        return NewPrivilegeEscalationGcpServiceAccountFinding(), nil
                    case "#microsoft.graph.privilegeEscalationUserFinding":
                        return NewPrivilegeEscalationUserFinding(), nil
                    case "#microsoft.graph.secretInformationAccessAwsResourceFinding":
                        return NewSecretInformationAccessAwsResourceFinding(), nil
                    case "#microsoft.graph.secretInformationAccessAwsRoleFinding":
                        return NewSecretInformationAccessAwsRoleFinding(), nil
                    case "#microsoft.graph.secretInformationAccessAwsServerlessFunctionFinding":
                        return NewSecretInformationAccessAwsServerlessFunctionFinding(), nil
                    case "#microsoft.graph.secretInformationAccessAwsUserFinding":
                        return NewSecretInformationAccessAwsUserFinding(), nil
                    case "#microsoft.graph.securityToolAwsResourceAdministratorFinding":
                        return NewSecurityToolAwsResourceAdministratorFinding(), nil
                    case "#microsoft.graph.securityToolAwsRoleAdministratorFinding":
                        return NewSecurityToolAwsRoleAdministratorFinding(), nil
                    case "#microsoft.graph.securityToolAwsServerlessFunctionAdministratorFinding":
                        return NewSecurityToolAwsServerlessFunctionAdministratorFinding(), nil
                    case "#microsoft.graph.securityToolAwsUserAdministratorFinding":
                        return NewSecurityToolAwsUserAdministratorFinding(), nil
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
                    case "#microsoft.graph.virtualMachineWithAwsStorageBucketAccessFinding":
                        return NewVirtualMachineWithAwsStorageBucketAccessFinding(), nil
                }
            }
        }
    }
    return NewFinding(), nil
}
// GetCreatedDateTime gets the createdDateTime property value. Defines when the finding was created.
// returns a *Time when successful
func (m *Finding) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("createdDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Finding) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["createdDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *Finding) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCreatedDateTime sets the createdDateTime property value. Defines when the finding was created.
func (m *Finding) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("createdDateTime", value)
    if err != nil {
        panic(err)
    }
}
type Findingable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
}
