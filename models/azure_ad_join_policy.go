package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AzureAdJoinPolicy 
type AzureAdJoinPolicy struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The identifiers of the groups that are in the scope of the policy. Required when the appliesTo property is set to selected.
    allowedGroups []string
    // The identifiers of users that are in the scope of the policy. Required when the appliesTo property is set to selected.
    allowedUsers []string
    // Specifies whether to block or allow fine-grained control of the policy scope. The possible values are: 0 (meaning none), 1 (meaning all), 2 (meaning selected), 3 (meaning unknownFutureValue). The default value is 1. When set to 2, at least one user or group identifier must be specified in either allowedUsers or allowedGroups.  Setting this property to 0 or 1 removes all identifiers in both allowedUsers and allowedGroups.
    appliesTo *PolicyScope
    // Specifies whether this policy scope is configurable by the admin. The default value is false. When an admin has enabled Intune (MEM) to manage devices, this property is set to false and appliesTo defaults to 1 (meaning all).
    isAdminConfigurable *bool
    // The OdataType property
    odataType *string
}
// NewAzureAdJoinPolicy instantiates a new azureAdJoinPolicy and sets the default values.
func NewAzureAdJoinPolicy()(*AzureAdJoinPolicy) {
    m := &AzureAdJoinPolicy{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    odataTypeValue := "#microsoft.graph.azureAdJoinPolicy";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateAzureAdJoinPolicyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAzureAdJoinPolicyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAzureAdJoinPolicy(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AzureAdJoinPolicy) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetAllowedGroups gets the allowedGroups property value. The identifiers of the groups that are in the scope of the policy. Required when the appliesTo property is set to selected.
func (m *AzureAdJoinPolicy) GetAllowedGroups()([]string) {
    return m.allowedGroups
}
// GetAllowedUsers gets the allowedUsers property value. The identifiers of users that are in the scope of the policy. Required when the appliesTo property is set to selected.
func (m *AzureAdJoinPolicy) GetAllowedUsers()([]string) {
    return m.allowedUsers
}
// GetAppliesTo gets the appliesTo property value. Specifies whether to block or allow fine-grained control of the policy scope. The possible values are: 0 (meaning none), 1 (meaning all), 2 (meaning selected), 3 (meaning unknownFutureValue). The default value is 1. When set to 2, at least one user or group identifier must be specified in either allowedUsers or allowedGroups.  Setting this property to 0 or 1 removes all identifiers in both allowedUsers and allowedGroups.
func (m *AzureAdJoinPolicy) GetAppliesTo()(*PolicyScope) {
    return m.appliesTo
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AzureAdJoinPolicy) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["allowedGroups"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetAllowedGroups)
    res["allowedUsers"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetAllowedUsers)
    res["appliesTo"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParsePolicyScope , m.SetAppliesTo)
    res["isAdminConfigurable"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIsAdminConfigurable)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    return res
}
// GetIsAdminConfigurable gets the isAdminConfigurable property value. Specifies whether this policy scope is configurable by the admin. The default value is false. When an admin has enabled Intune (MEM) to manage devices, this property is set to false and appliesTo defaults to 1 (meaning all).
func (m *AzureAdJoinPolicy) GetIsAdminConfigurable()(*bool) {
    return m.isAdminConfigurable
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *AzureAdJoinPolicy) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *AzureAdJoinPolicy) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAllowedGroups() != nil {
        err := writer.WriteCollectionOfStringValues("allowedGroups", m.GetAllowedGroups())
        if err != nil {
            return err
        }
    }
    if m.GetAllowedUsers() != nil {
        err := writer.WriteCollectionOfStringValues("allowedUsers", m.GetAllowedUsers())
        if err != nil {
            return err
        }
    }
    if m.GetAppliesTo() != nil {
        cast := (*m.GetAppliesTo()).String()
        err := writer.WriteStringValue("appliesTo", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isAdminConfigurable", m.GetIsAdminConfigurable())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AzureAdJoinPolicy) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetAllowedGroups sets the allowedGroups property value. The identifiers of the groups that are in the scope of the policy. Required when the appliesTo property is set to selected.
func (m *AzureAdJoinPolicy) SetAllowedGroups(value []string)() {
    m.allowedGroups = value
}
// SetAllowedUsers sets the allowedUsers property value. The identifiers of users that are in the scope of the policy. Required when the appliesTo property is set to selected.
func (m *AzureAdJoinPolicy) SetAllowedUsers(value []string)() {
    m.allowedUsers = value
}
// SetAppliesTo sets the appliesTo property value. Specifies whether to block or allow fine-grained control of the policy scope. The possible values are: 0 (meaning none), 1 (meaning all), 2 (meaning selected), 3 (meaning unknownFutureValue). The default value is 1. When set to 2, at least one user or group identifier must be specified in either allowedUsers or allowedGroups.  Setting this property to 0 or 1 removes all identifiers in both allowedUsers and allowedGroups.
func (m *AzureAdJoinPolicy) SetAppliesTo(value *PolicyScope)() {
    m.appliesTo = value
}
// SetIsAdminConfigurable sets the isAdminConfigurable property value. Specifies whether this policy scope is configurable by the admin. The default value is false. When an admin has enabled Intune (MEM) to manage devices, this property is set to false and appliesTo defaults to 1 (meaning all).
func (m *AzureAdJoinPolicy) SetIsAdminConfigurable(value *bool)() {
    m.isAdminConfigurable = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AzureAdJoinPolicy) SetOdataType(value *string)() {
    m.odataType = value
}
