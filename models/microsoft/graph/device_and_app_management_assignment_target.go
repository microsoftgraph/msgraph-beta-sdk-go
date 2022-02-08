package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DeviceAndAppManagementAssignmentTarget 
type DeviceAndAppManagementAssignmentTarget struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The Id of the filter for the target assignment.
    deviceAndAppManagementAssignmentFilterId *string;
    // The type of filter of the target assignment i.e. Exclude or Include. Possible values are: none, include, exclude.
    deviceAndAppManagementAssignmentFilterType *DeviceAndAppManagementAssignmentFilterType;
}
// NewDeviceAndAppManagementAssignmentTarget instantiates a new deviceAndAppManagementAssignmentTarget and sets the default values.
func NewDeviceAndAppManagementAssignmentTarget()(*DeviceAndAppManagementAssignmentTarget) {
    m := &DeviceAndAppManagementAssignmentTarget{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceAndAppManagementAssignmentTarget) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDeviceAndAppManagementAssignmentFilterId gets the deviceAndAppManagementAssignmentFilterId property value. The Id of the filter for the target assignment.
func (m *DeviceAndAppManagementAssignmentTarget) GetDeviceAndAppManagementAssignmentFilterId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceAndAppManagementAssignmentFilterId
    }
}
// GetDeviceAndAppManagementAssignmentFilterType gets the deviceAndAppManagementAssignmentFilterType property value. The type of filter of the target assignment i.e. Exclude or Include. Possible values are: none, include, exclude.
func (m *DeviceAndAppManagementAssignmentTarget) GetDeviceAndAppManagementAssignmentFilterType()(*DeviceAndAppManagementAssignmentFilterType) {
    if m == nil {
        return nil
    } else {
        return m.deviceAndAppManagementAssignmentFilterType
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceAndAppManagementAssignmentTarget) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["deviceAndAppManagementAssignmentFilterId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceAndAppManagementAssignmentFilterId(val)
        }
        return nil
    }
    res["deviceAndAppManagementAssignmentFilterType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceAndAppManagementAssignmentFilterType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceAndAppManagementAssignmentFilterType(val.(*DeviceAndAppManagementAssignmentFilterType))
        }
        return nil
    }
    return res
}
func (m *DeviceAndAppManagementAssignmentTarget) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *DeviceAndAppManagementAssignmentTarget) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("deviceAndAppManagementAssignmentFilterId", m.GetDeviceAndAppManagementAssignmentFilterId())
        if err != nil {
            return err
        }
    }
    if m.GetDeviceAndAppManagementAssignmentFilterType() != nil {
        cast := (*m.GetDeviceAndAppManagementAssignmentFilterType()).String()
        err := writer.WriteStringValue("deviceAndAppManagementAssignmentFilterType", &cast)
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
func (m *DeviceAndAppManagementAssignmentTarget) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetDeviceAndAppManagementAssignmentFilterId sets the deviceAndAppManagementAssignmentFilterId property value. The Id of the filter for the target assignment.
func (m *DeviceAndAppManagementAssignmentTarget) SetDeviceAndAppManagementAssignmentFilterId(value *string)() {
    if m != nil {
        m.deviceAndAppManagementAssignmentFilterId = value
    }
}
// SetDeviceAndAppManagementAssignmentFilterType sets the deviceAndAppManagementAssignmentFilterType property value. The type of filter of the target assignment i.e. Exclude or Include. Possible values are: none, include, exclude.
func (m *DeviceAndAppManagementAssignmentTarget) SetDeviceAndAppManagementAssignmentFilterType(value *DeviceAndAppManagementAssignmentFilterType)() {
    if m != nil {
        m.deviceAndAppManagementAssignmentFilterType = value
    }
}
