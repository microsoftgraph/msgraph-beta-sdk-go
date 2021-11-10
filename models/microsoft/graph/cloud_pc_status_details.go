package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type CloudPcStatusDetails struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Any additional information about the cloud PC status.
    additionalInformation []KeyValuePair;
    // The code associated with the cloud PC status.
    code *string;
    // The status message.
    message *string;
}
// Instantiates a new cloudPcStatusDetails and sets the default values.
func NewCloudPcStatusDetails()(*CloudPcStatusDetails) {
    m := &CloudPcStatusDetails{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CloudPcStatusDetails) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the additionalInformation property value. Any additional information about the cloud PC status.
func (m *CloudPcStatusDetails) GetAdditionalInformation()([]KeyValuePair) {
    if m == nil {
        return nil
    } else {
        return m.additionalInformation
    }
}
// Gets the code property value. The code associated with the cloud PC status.
func (m *CloudPcStatusDetails) GetCode()(*string) {
    if m == nil {
        return nil
    } else {
        return m.code
    }
}
// Gets the message property value. The status message.
func (m *CloudPcStatusDetails) GetMessage()(*string) {
    if m == nil {
        return nil
    } else {
        return m.message
    }
}
// The deserialization information for the current model
func (m *CloudPcStatusDetails) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["additionalInformation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewKeyValuePair() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]KeyValuePair, len(val))
            for i, v := range val {
                res[i] = *(v.(*KeyValuePair))
            }
            m.SetAdditionalInformation(res)
        }
        return nil
    }
    res["code"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCode(val)
        }
        return nil
    }
    res["message"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMessage(val)
        }
        return nil
    }
    return res
}
func (m *CloudPcStatusDetails) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *CloudPcStatusDetails) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAdditionalInformation()))
        for i, v := range m.GetAdditionalInformation() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("additionalInformation", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("code", m.GetCode())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("message", m.GetMessage())
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
func (m *CloudPcStatusDetails) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the additionalInformation property value. Any additional information about the cloud PC status.
// Parameters:
//  - value : Value to set for the additionalInformation property.
func (m *CloudPcStatusDetails) SetAdditionalInformation(value []KeyValuePair)() {
    m.additionalInformation = value
}
// Sets the code property value. The code associated with the cloud PC status.
// Parameters:
//  - value : Value to set for the code property.
func (m *CloudPcStatusDetails) SetCode(value *string)() {
    m.code = value
}
// Sets the message property value. The status message.
// Parameters:
//  - value : Value to set for the message property.
func (m *CloudPcStatusDetails) SetMessage(value *string)() {
    m.message = value
}
