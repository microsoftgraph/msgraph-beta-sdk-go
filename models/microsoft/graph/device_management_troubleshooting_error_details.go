package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type DeviceManagementTroubleshootingErrorDetails struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Not yet documented
    context *string;
    // Not yet documented
    failure *string;
    // The detailed description of what went wrong.
    failureDetails *string;
    // The detailed description of how to remediate this issue.
    remediation *string;
    // Links to helpful documentation about this failure.
    resources []DeviceManagementTroubleshootingErrorResource;
}
// Instantiates a new deviceManagementTroubleshootingErrorDetails and sets the default values.
func NewDeviceManagementTroubleshootingErrorDetails()(*DeviceManagementTroubleshootingErrorDetails) {
    m := &DeviceManagementTroubleshootingErrorDetails{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementTroubleshootingErrorDetails) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the context property value. Not yet documented
func (m *DeviceManagementTroubleshootingErrorDetails) GetContext()(*string) {
    if m == nil {
        return nil
    } else {
        return m.context
    }
}
// Gets the failure property value. Not yet documented
func (m *DeviceManagementTroubleshootingErrorDetails) GetFailure()(*string) {
    if m == nil {
        return nil
    } else {
        return m.failure
    }
}
// Gets the failureDetails property value. The detailed description of what went wrong.
func (m *DeviceManagementTroubleshootingErrorDetails) GetFailureDetails()(*string) {
    if m == nil {
        return nil
    } else {
        return m.failureDetails
    }
}
// Gets the remediation property value. The detailed description of how to remediate this issue.
func (m *DeviceManagementTroubleshootingErrorDetails) GetRemediation()(*string) {
    if m == nil {
        return nil
    } else {
        return m.remediation
    }
}
// Gets the resources property value. Links to helpful documentation about this failure.
func (m *DeviceManagementTroubleshootingErrorDetails) GetResources()([]DeviceManagementTroubleshootingErrorResource) {
    if m == nil {
        return nil
    } else {
        return m.resources
    }
}
// The deserialization information for the current model
func (m *DeviceManagementTroubleshootingErrorDetails) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["context"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContext(val)
        }
        return nil
    }
    res["failure"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFailure(val)
        }
        return nil
    }
    res["failureDetails"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFailureDetails(val)
        }
        return nil
    }
    res["remediation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRemediation(val)
        }
        return nil
    }
    res["resources"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceManagementTroubleshootingErrorResource() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementTroubleshootingErrorResource, len(val))
            for i, v := range val {
                res[i] = *(v.(*DeviceManagementTroubleshootingErrorResource))
            }
            m.SetResources(res)
        }
        return nil
    }
    return res
}
func (m *DeviceManagementTroubleshootingErrorDetails) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *DeviceManagementTroubleshootingErrorDetails) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("context", m.GetContext())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("failure", m.GetFailure())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("failureDetails", m.GetFailureDetails())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("remediation", m.GetRemediation())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetResources()))
        for i, v := range m.GetResources() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("resources", cast)
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
func (m *DeviceManagementTroubleshootingErrorDetails) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the context property value. Not yet documented
// Parameters:
//  - value : Value to set for the context property.
func (m *DeviceManagementTroubleshootingErrorDetails) SetContext(value *string)() {
    m.context = value
}
// Sets the failure property value. Not yet documented
// Parameters:
//  - value : Value to set for the failure property.
func (m *DeviceManagementTroubleshootingErrorDetails) SetFailure(value *string)() {
    m.failure = value
}
// Sets the failureDetails property value. The detailed description of what went wrong.
// Parameters:
//  - value : Value to set for the failureDetails property.
func (m *DeviceManagementTroubleshootingErrorDetails) SetFailureDetails(value *string)() {
    m.failureDetails = value
}
// Sets the remediation property value. The detailed description of how to remediate this issue.
// Parameters:
//  - value : Value to set for the remediation property.
func (m *DeviceManagementTroubleshootingErrorDetails) SetRemediation(value *string)() {
    m.remediation = value
}
// Sets the resources property value. Links to helpful documentation about this failure.
// Parameters:
//  - value : Value to set for the resources property.
func (m *DeviceManagementTroubleshootingErrorDetails) SetResources(value []DeviceManagementTroubleshootingErrorResource)() {
    m.resources = value
}
