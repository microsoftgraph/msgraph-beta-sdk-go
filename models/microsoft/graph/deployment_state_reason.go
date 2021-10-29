package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    ifded49a845bbaa9057da6e2cf565863ac34eb797e99b129c3e0659166af6b7e2 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/windowsupdates"
)

// 
type DeploymentStateReason struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Specifies a reason for the deployment state. Possible values are: scheduledByOfferWindow, offeringByRequest, pausedByRequest, pausedByMonitoring. Read-only.
    value *ifded49a845bbaa9057da6e2cf565863ac34eb797e99b129c3e0659166af6b7e2.DeploymentStateReasonValue;
}
// Instantiates a new deploymentStateReason and sets the default values.
func NewDeploymentStateReason()(*DeploymentStateReason) {
    m := &DeploymentStateReason{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeploymentStateReason) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the value property value. Specifies a reason for the deployment state. Possible values are: scheduledByOfferWindow, offeringByRequest, pausedByRequest, pausedByMonitoring. Read-only.
func (m *DeploymentStateReason) GetValue()(*ifded49a845bbaa9057da6e2cf565863ac34eb797e99b129c3e0659166af6b7e2.DeploymentStateReasonValue) {
    if m == nil {
        return nil
    } else {
        return m.value
    }
}
// The deserialization information for the current model
func (m *DeploymentStateReason) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["value"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ifded49a845bbaa9057da6e2cf565863ac34eb797e99b129c3e0659166af6b7e2.ParseDeploymentStateReasonValue)
        if err != nil {
            return err
        }
        cast := val.(ifded49a845bbaa9057da6e2cf565863ac34eb797e99b129c3e0659166af6b7e2.DeploymentStateReasonValue)
        m.SetValue(&cast)
        return nil
    }
    return res
}
func (m *DeploymentStateReason) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *DeploymentStateReason) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetValue() != nil {
        cast := m.GetValue().String()
        err := writer.WriteStringValue("value", &cast)
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *DeploymentStateReason) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the value property value. Specifies a reason for the deployment state. Possible values are: scheduledByOfferWindow, offeringByRequest, pausedByRequest, pausedByMonitoring. Read-only.
// Parameters:
//  - value : Value to set for the value property.
func (m *DeploymentStateReason) SetValue(value *ifded49a845bbaa9057da6e2cf565863ac34eb797e99b129c3e0659166af6b7e2.DeploymentStateReasonValue)() {
    m.value = value
}
