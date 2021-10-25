package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type DeviceManagementTroubleshootingErrorDetails struct {
    additionalData map[string]interface{};
    context *string;
    failure *string;
    failureDetails *string;
    remediation *string;
    resources []DeviceManagementTroubleshootingErrorResource;
}
func NewDeviceManagementTroubleshootingErrorDetails()(*DeviceManagementTroubleshootingErrorDetails) {
    m := &DeviceManagementTroubleshootingErrorDetails{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *DeviceManagementTroubleshootingErrorDetails) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *DeviceManagementTroubleshootingErrorDetails) GetContext()(*string) {
    if m == nil {
        return nil
    } else {
        return m.context
    }
}
func (m *DeviceManagementTroubleshootingErrorDetails) GetFailure()(*string) {
    if m == nil {
        return nil
    } else {
        return m.failure
    }
}
func (m *DeviceManagementTroubleshootingErrorDetails) GetFailureDetails()(*string) {
    if m == nil {
        return nil
    } else {
        return m.failureDetails
    }
}
func (m *DeviceManagementTroubleshootingErrorDetails) GetRemediation()(*string) {
    if m == nil {
        return nil
    } else {
        return m.remediation
    }
}
func (m *DeviceManagementTroubleshootingErrorDetails) GetResources()([]DeviceManagementTroubleshootingErrorResource) {
    if m == nil {
        return nil
    } else {
        return m.resources
    }
}
func (m *DeviceManagementTroubleshootingErrorDetails) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["context"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetContext(val)
        return nil
    }
    res["failure"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetFailure(val)
        return nil
    }
    res["failureDetails"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetFailureDetails(val)
        return nil
    }
    res["remediation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetRemediation(val)
        return nil
    }
    res["resources"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceManagementTroubleshootingErrorResource() })
        if err != nil {
            return err
        }
        res := make([]DeviceManagementTroubleshootingErrorResource, len(val))
        for i, v := range val {
            res[i] = *(v.(*DeviceManagementTroubleshootingErrorResource))
        }
        m.SetResources(res)
        return nil
    }
    return res
}
func (m *DeviceManagementTroubleshootingErrorDetails) IsNil()(bool) {
    return m == nil
}
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
func (m *DeviceManagementTroubleshootingErrorDetails) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *DeviceManagementTroubleshootingErrorDetails) SetContext(value *string)() {
    m.context = value
}
func (m *DeviceManagementTroubleshootingErrorDetails) SetFailure(value *string)() {
    m.failure = value
}
func (m *DeviceManagementTroubleshootingErrorDetails) SetFailureDetails(value *string)() {
    m.failureDetails = value
}
func (m *DeviceManagementTroubleshootingErrorDetails) SetRemediation(value *string)() {
    m.remediation = value
}
func (m *DeviceManagementTroubleshootingErrorDetails) SetResources(value []DeviceManagementTroubleshootingErrorResource)() {
    m.resources = value
}
