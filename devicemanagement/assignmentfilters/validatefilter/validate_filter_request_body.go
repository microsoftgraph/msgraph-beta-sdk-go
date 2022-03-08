package validatefilter

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// ValidateFilterRequestBody provides operations to call the validateFilter method.
type ValidateFilterRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    deviceAndAppManagementAssignmentFilter i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceAndAppManagementAssignmentFilterable;
}
// NewValidateFilterRequestBody instantiates a new validateFilterRequestBody and sets the default values.
func NewValidateFilterRequestBody()(*ValidateFilterRequestBody) {
    m := &ValidateFilterRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateValidateFilterRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateValidateFilterRequestBodyFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewValidateFilterRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ValidateFilterRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDeviceAndAppManagementAssignmentFilter gets the deviceAndAppManagementAssignmentFilter property value. 
func (m *ValidateFilterRequestBody) GetDeviceAndAppManagementAssignmentFilter()(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceAndAppManagementAssignmentFilterable) {
    if m == nil {
        return nil
    } else {
        return m.deviceAndAppManagementAssignmentFilter
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ValidateFilterRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["deviceAndAppManagementAssignmentFilter"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateDeviceAndAppManagementAssignmentFilterFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceAndAppManagementAssignmentFilter(val.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceAndAppManagementAssignmentFilterable))
        }
        return nil
    }
    return res
}
func (m *ValidateFilterRequestBody) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ValidateFilterRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("deviceAndAppManagementAssignmentFilter", m.GetDeviceAndAppManagementAssignmentFilter())
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
func (m *ValidateFilterRequestBody) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetDeviceAndAppManagementAssignmentFilter sets the deviceAndAppManagementAssignmentFilter property value. 
func (m *ValidateFilterRequestBody) SetDeviceAndAppManagementAssignmentFilter(value i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceAndAppManagementAssignmentFilterable)() {
    if m != nil {
        m.deviceAndAppManagementAssignmentFilter = value
    }
}
