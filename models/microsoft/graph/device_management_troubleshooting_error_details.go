package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DeviceManagementTroubleshootingErrorDetails provides operations to manage the deviceManagement singleton.
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
    resources []DeviceManagementTroubleshootingErrorResourceable;
}
// NewDeviceManagementTroubleshootingErrorDetails instantiates a new deviceManagementTroubleshootingErrorDetails and sets the default values.
func NewDeviceManagementTroubleshootingErrorDetails()(*DeviceManagementTroubleshootingErrorDetails) {
    m := &DeviceManagementTroubleshootingErrorDetails{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDeviceManagementTroubleshootingErrorDetailsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementTroubleshootingErrorDetailsFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewDeviceManagementTroubleshootingErrorDetails(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementTroubleshootingErrorDetails) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetContext gets the context property value. Not yet documented
func (m *DeviceManagementTroubleshootingErrorDetails) GetContext()(*string) {
    if m == nil {
        return nil
    } else {
        return m.context
    }
}
// GetFailure gets the failure property value. Not yet documented
func (m *DeviceManagementTroubleshootingErrorDetails) GetFailure()(*string) {
    if m == nil {
        return nil
    } else {
        return m.failure
    }
}
// GetFailureDetails gets the failureDetails property value. The detailed description of what went wrong.
func (m *DeviceManagementTroubleshootingErrorDetails) GetFailureDetails()(*string) {
    if m == nil {
        return nil
    } else {
        return m.failureDetails
    }
}
// GetFieldDeserializers the deserialization information for the current model
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
        val, err := n.GetCollectionOfObjectValues(CreateDeviceManagementTroubleshootingErrorResourceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementTroubleshootingErrorResourceable, len(val))
            for i, v := range val {
                res[i] = v.(DeviceManagementTroubleshootingErrorResourceable)
            }
            m.SetResources(res)
        }
        return nil
    }
    return res
}
// GetRemediation gets the remediation property value. The detailed description of how to remediate this issue.
func (m *DeviceManagementTroubleshootingErrorDetails) GetRemediation()(*string) {
    if m == nil {
        return nil
    } else {
        return m.remediation
    }
}
// GetResources gets the resources property value. Links to helpful documentation about this failure.
func (m *DeviceManagementTroubleshootingErrorDetails) GetResources()([]DeviceManagementTroubleshootingErrorResourceable) {
    if m == nil {
        return nil
    } else {
        return m.resources
    }
}
func (m *DeviceManagementTroubleshootingErrorDetails) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
    if m.GetResources() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetResources()))
        for i, v := range m.GetResources() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementTroubleshootingErrorDetails) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetContext sets the context property value. Not yet documented
func (m *DeviceManagementTroubleshootingErrorDetails) SetContext(value *string)() {
    if m != nil {
        m.context = value
    }
}
// SetFailure sets the failure property value. Not yet documented
func (m *DeviceManagementTroubleshootingErrorDetails) SetFailure(value *string)() {
    if m != nil {
        m.failure = value
    }
}
// SetFailureDetails sets the failureDetails property value. The detailed description of what went wrong.
func (m *DeviceManagementTroubleshootingErrorDetails) SetFailureDetails(value *string)() {
    if m != nil {
        m.failureDetails = value
    }
}
// SetRemediation sets the remediation property value. The detailed description of how to remediate this issue.
func (m *DeviceManagementTroubleshootingErrorDetails) SetRemediation(value *string)() {
    if m != nil {
        m.remediation = value
    }
}
// SetResources sets the resources property value. Links to helpful documentation about this failure.
func (m *DeviceManagementTroubleshootingErrorDetails) SetResources(value []DeviceManagementTroubleshootingErrorResourceable)() {
    if m != nil {
        m.resources = value
    }
}
