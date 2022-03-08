package getoffice365groupsactivitygroupcountswithperiod

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// GetOffice365GroupsActivityGroupCountsWithPeriodResponse provides operations to call the getOffice365GroupsActivityGroupCounts method.
type GetOffice365GroupsActivityGroupCountsWithPeriodResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    value []i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Office365GroupsActivityGroupCountsable;
}
// NewGetOffice365GroupsActivityGroupCountsWithPeriodResponse instantiates a new getOffice365GroupsActivityGroupCountsWithPeriodResponse and sets the default values.
func NewGetOffice365GroupsActivityGroupCountsWithPeriodResponse()(*GetOffice365GroupsActivityGroupCountsWithPeriodResponse) {
    m := &GetOffice365GroupsActivityGroupCountsWithPeriodResponse{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateGetOffice365GroupsActivityGroupCountsWithPeriodResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGetOffice365GroupsActivityGroupCountsWithPeriodResponseFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewGetOffice365GroupsActivityGroupCountsWithPeriodResponse(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *GetOffice365GroupsActivityGroupCountsWithPeriodResponse) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GetOffice365GroupsActivityGroupCountsWithPeriodResponse) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["value"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateOffice365GroupsActivityGroupCountsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Office365GroupsActivityGroupCountsable, len(val))
            for i, v := range val {
                res[i] = v.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Office365GroupsActivityGroupCountsable)
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. 
func (m *GetOffice365GroupsActivityGroupCountsWithPeriodResponse) GetValue()([]i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Office365GroupsActivityGroupCountsable) {
    if m == nil {
        return nil
    } else {
        return m.value
    }
}
func (m *GetOffice365GroupsActivityGroupCountsWithPeriodResponse) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *GetOffice365GroupsActivityGroupCountsWithPeriodResponse) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetValue() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetValue()))
        for i, v := range m.GetValue() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("value", cast)
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
func (m *GetOffice365GroupsActivityGroupCountsWithPeriodResponse) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetValue sets the value property value. 
func (m *GetOffice365GroupsActivityGroupCountsWithPeriodResponse) SetValue(value []i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Office365GroupsActivityGroupCountsable)() {
    if m != nil {
        m.value = value
    }
}
