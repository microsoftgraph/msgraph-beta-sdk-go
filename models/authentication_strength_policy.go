package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AuthenticationStrengthPolicy 
type AuthenticationStrengthPolicy struct {
    Entity
    // The allowedCombinations property
    allowedCombinations []AuthenticationMethodModes
    // The combinationConfigurations property
    combinationConfigurations []AuthenticationCombinationConfigurationable
    // The createdDateTime property
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The description property
    description *string
    // The displayName property
    displayName *string
    // The modifiedDateTime property
    modifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The policyType property
    policyType *AuthenticationStrengthPolicyType
    // The requirementsSatisfied property
    requirementsSatisfied *AuthenticationStrengthRequirements
}
// NewAuthenticationStrengthPolicy instantiates a new AuthenticationStrengthPolicy and sets the default values.
func NewAuthenticationStrengthPolicy()(*AuthenticationStrengthPolicy) {
    m := &AuthenticationStrengthPolicy{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.authenticationStrengthPolicy";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateAuthenticationStrengthPolicyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAuthenticationStrengthPolicyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAuthenticationStrengthPolicy(), nil
}
// GetAllowedCombinations gets the allowedCombinations property value. The allowedCombinations property
func (m *AuthenticationStrengthPolicy) GetAllowedCombinations()([]AuthenticationMethodModes) {
    return m.allowedCombinations
}
// GetCombinationConfigurations gets the combinationConfigurations property value. The combinationConfigurations property
func (m *AuthenticationStrengthPolicy) GetCombinationConfigurations()([]AuthenticationCombinationConfigurationable) {
    return m.combinationConfigurations
}
// GetCreatedDateTime gets the createdDateTime property value. The createdDateTime property
func (m *AuthenticationStrengthPolicy) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetDescription gets the description property value. The description property
func (m *AuthenticationStrengthPolicy) GetDescription()(*string) {
    return m.description
}
// GetDisplayName gets the displayName property value. The displayName property
func (m *AuthenticationStrengthPolicy) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AuthenticationStrengthPolicy) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["allowedCombinations"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfEnumValues(ParseAuthenticationMethodModes , m.SetAllowedCombinations)
    res["combinationConfigurations"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateAuthenticationCombinationConfigurationFromDiscriminatorValue , m.SetCombinationConfigurations)
    res["createdDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetCreatedDateTime)
    res["description"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDescription)
    res["displayName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDisplayName)
    res["modifiedDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetModifiedDateTime)
    res["policyType"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseAuthenticationStrengthPolicyType , m.SetPolicyType)
    res["requirementsSatisfied"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseAuthenticationStrengthRequirements , m.SetRequirementsSatisfied)
    return res
}
// GetModifiedDateTime gets the modifiedDateTime property value. The modifiedDateTime property
func (m *AuthenticationStrengthPolicy) GetModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.modifiedDateTime
}
// GetPolicyType gets the policyType property value. The policyType property
func (m *AuthenticationStrengthPolicy) GetPolicyType()(*AuthenticationStrengthPolicyType) {
    return m.policyType
}
// GetRequirementsSatisfied gets the requirementsSatisfied property value. The requirementsSatisfied property
func (m *AuthenticationStrengthPolicy) GetRequirementsSatisfied()(*AuthenticationStrengthRequirements) {
    return m.requirementsSatisfied
}
// Serialize serializes information the current object
func (m *AuthenticationStrengthPolicy) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAllowedCombinations() != nil {
        err = writer.WriteCollectionOfStringValues("allowedCombinations", SerializeAuthenticationMethodModes(m.GetAllowedCombinations()))
        if err != nil {
            return err
        }
    }
    if m.GetCombinationConfigurations() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetCombinationConfigurations())
        err = writer.WriteCollectionOfObjectValues("combinationConfigurations", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("modifiedDateTime", m.GetModifiedDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetPolicyType() != nil {
        cast := (*m.GetPolicyType()).String()
        err = writer.WriteStringValue("policyType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetRequirementsSatisfied() != nil {
        cast := (*m.GetRequirementsSatisfied()).String()
        err = writer.WriteStringValue("requirementsSatisfied", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAllowedCombinations sets the allowedCombinations property value. The allowedCombinations property
func (m *AuthenticationStrengthPolicy) SetAllowedCombinations(value []AuthenticationMethodModes)() {
    m.allowedCombinations = value
}
// SetCombinationConfigurations sets the combinationConfigurations property value. The combinationConfigurations property
func (m *AuthenticationStrengthPolicy) SetCombinationConfigurations(value []AuthenticationCombinationConfigurationable)() {
    m.combinationConfigurations = value
}
// SetCreatedDateTime sets the createdDateTime property value. The createdDateTime property
func (m *AuthenticationStrengthPolicy) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetDescription sets the description property value. The description property
func (m *AuthenticationStrengthPolicy) SetDescription(value *string)() {
    m.description = value
}
// SetDisplayName sets the displayName property value. The displayName property
func (m *AuthenticationStrengthPolicy) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetModifiedDateTime sets the modifiedDateTime property value. The modifiedDateTime property
func (m *AuthenticationStrengthPolicy) SetModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.modifiedDateTime = value
}
// SetPolicyType sets the policyType property value. The policyType property
func (m *AuthenticationStrengthPolicy) SetPolicyType(value *AuthenticationStrengthPolicyType)() {
    m.policyType = value
}
// SetRequirementsSatisfied sets the requirementsSatisfied property value. The requirementsSatisfied property
func (m *AuthenticationStrengthPolicy) SetRequirementsSatisfied(value *AuthenticationStrengthRequirements)() {
    m.requirementsSatisfied = value
}
