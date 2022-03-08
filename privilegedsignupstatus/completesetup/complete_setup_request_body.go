package completesetup

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// CompleteSetupRequestBody provides operations to call the completeSetup method.
type CompleteSetupRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    tenantSetupInfo i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.TenantSetupInfoable;
}
// NewCompleteSetupRequestBody instantiates a new completeSetupRequestBody and sets the default values.
func NewCompleteSetupRequestBody()(*CompleteSetupRequestBody) {
    m := &CompleteSetupRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateCompleteSetupRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCompleteSetupRequestBodyFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewCompleteSetupRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CompleteSetupRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CompleteSetupRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["tenantSetupInfo"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateTenantSetupInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTenantSetupInfo(val.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.TenantSetupInfoable))
        }
        return nil
    }
    return res
}
// GetTenantSetupInfo gets the tenantSetupInfo property value. 
func (m *CompleteSetupRequestBody) GetTenantSetupInfo()(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.TenantSetupInfoable) {
    if m == nil {
        return nil
    } else {
        return m.tenantSetupInfo
    }
}
func (m *CompleteSetupRequestBody) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *CompleteSetupRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("tenantSetupInfo", m.GetTenantSetupInfo())
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
func (m *CompleteSetupRequestBody) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetTenantSetupInfo sets the tenantSetupInfo property value. 
func (m *CompleteSetupRequestBody) SetTenantSetupInfo(value i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.TenantSetupInfoable)() {
    if m != nil {
        m.tenantSetupInfo = value
    }
}
