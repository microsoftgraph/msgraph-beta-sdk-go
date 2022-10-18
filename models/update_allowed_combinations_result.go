package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UpdateAllowedCombinationsResult 
type UpdateAllowedCombinationsResult struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Information about why the updateAllowedCombinations action was successful or failed.
    additionalInformation *string
    // References to existing Conditional Access policies that use this authentication strength.
    conditionalAccessReferences []string
    // The list of current authentication method combinations allowed by the authentication strength.
    currentCombinations []AuthenticationMethodModes
    // The OdataType property
    odataType *string
    // The list of former authentication method combinations allowed by the authentication strength before they were updated through the updateAllowedCombinations action.
    previousCombinations []AuthenticationMethodModes
}
// NewUpdateAllowedCombinationsResult instantiates a new updateAllowedCombinationsResult and sets the default values.
func NewUpdateAllowedCombinationsResult()(*UpdateAllowedCombinationsResult) {
    m := &UpdateAllowedCombinationsResult{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    odataTypeValue := "#microsoft.graph.updateAllowedCombinationsResult";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateUpdateAllowedCombinationsResultFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUpdateAllowedCombinationsResultFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUpdateAllowedCombinationsResult(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UpdateAllowedCombinationsResult) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetAdditionalInformation gets the additionalInformation property value. Information about why the updateAllowedCombinations action was successful or failed.
func (m *UpdateAllowedCombinationsResult) GetAdditionalInformation()(*string) {
    return m.additionalInformation
}
// GetConditionalAccessReferences gets the conditionalAccessReferences property value. References to existing Conditional Access policies that use this authentication strength.
func (m *UpdateAllowedCombinationsResult) GetConditionalAccessReferences()([]string) {
    return m.conditionalAccessReferences
}
// GetCurrentCombinations gets the currentCombinations property value. The list of current authentication method combinations allowed by the authentication strength.
func (m *UpdateAllowedCombinationsResult) GetCurrentCombinations()([]AuthenticationMethodModes) {
    return m.currentCombinations
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UpdateAllowedCombinationsResult) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["additionalInformation"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetAdditionalInformation)
    res["conditionalAccessReferences"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetConditionalAccessReferences)
    res["currentCombinations"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfEnumValues(ParseAuthenticationMethodModes , m.SetCurrentCombinations)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    res["previousCombinations"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfEnumValues(ParseAuthenticationMethodModes , m.SetPreviousCombinations)
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *UpdateAllowedCombinationsResult) GetOdataType()(*string) {
    return m.odataType
}
// GetPreviousCombinations gets the previousCombinations property value. The list of former authentication method combinations allowed by the authentication strength before they were updated through the updateAllowedCombinations action.
func (m *UpdateAllowedCombinationsResult) GetPreviousCombinations()([]AuthenticationMethodModes) {
    return m.previousCombinations
}
// Serialize serializes information the current object
func (m *UpdateAllowedCombinationsResult) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("additionalInformation", m.GetAdditionalInformation())
        if err != nil {
            return err
        }
    }
    if m.GetConditionalAccessReferences() != nil {
        err := writer.WriteCollectionOfStringValues("conditionalAccessReferences", m.GetConditionalAccessReferences())
        if err != nil {
            return err
        }
    }
    if m.GetCurrentCombinations() != nil {
        err := writer.WriteCollectionOfStringValues("currentCombinations", SerializeAuthenticationMethodModes(m.GetCurrentCombinations()))
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
    if m.GetPreviousCombinations() != nil {
        err := writer.WriteCollectionOfStringValues("previousCombinations", SerializeAuthenticationMethodModes(m.GetPreviousCombinations()))
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
func (m *UpdateAllowedCombinationsResult) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetAdditionalInformation sets the additionalInformation property value. Information about why the updateAllowedCombinations action was successful or failed.
func (m *UpdateAllowedCombinationsResult) SetAdditionalInformation(value *string)() {
    m.additionalInformation = value
}
// SetConditionalAccessReferences sets the conditionalAccessReferences property value. References to existing Conditional Access policies that use this authentication strength.
func (m *UpdateAllowedCombinationsResult) SetConditionalAccessReferences(value []string)() {
    m.conditionalAccessReferences = value
}
// SetCurrentCombinations sets the currentCombinations property value. The list of current authentication method combinations allowed by the authentication strength.
func (m *UpdateAllowedCombinationsResult) SetCurrentCombinations(value []AuthenticationMethodModes)() {
    m.currentCombinations = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *UpdateAllowedCombinationsResult) SetOdataType(value *string)() {
    m.odataType = value
}
// SetPreviousCombinations sets the previousCombinations property value. The list of former authentication method combinations allowed by the authentication strength before they were updated through the updateAllowedCombinations action.
func (m *UpdateAllowedCombinationsResult) SetPreviousCombinations(value []AuthenticationMethodModes)() {
    m.previousCombinations = value
}
