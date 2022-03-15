package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// CloudPcStatusDetails provides operations to manage the compliance singleton.
type CloudPcStatusDetails struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Any additional information about the Cloud PC status.
    additionalInformation []KeyValuePairable;
    // The code associated with the Cloud PC status.
    code *string;
    // The status message.
    message *string;
}
// NewCloudPcStatusDetails instantiates a new cloudPcStatusDetails and sets the default values.
func NewCloudPcStatusDetails()(*CloudPcStatusDetails) {
    m := &CloudPcStatusDetails{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateCloudPcStatusDetailsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCloudPcStatusDetailsFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewCloudPcStatusDetails(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CloudPcStatusDetails) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAdditionalInformation gets the additionalInformation property value. Any additional information about the Cloud PC status.
func (m *CloudPcStatusDetails) GetAdditionalInformation()([]KeyValuePairable) {
    if m == nil {
        return nil
    } else {
        return m.additionalInformation
    }
}
// GetCode gets the code property value. The code associated with the Cloud PC status.
func (m *CloudPcStatusDetails) GetCode()(*string) {
    if m == nil {
        return nil
    } else {
        return m.code
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CloudPcStatusDetails) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["additionalInformation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateKeyValuePairFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]KeyValuePairable, len(val))
            for i, v := range val {
                res[i] = v.(KeyValuePairable)
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
// GetMessage gets the message property value. The status message.
func (m *CloudPcStatusDetails) GetMessage()(*string) {
    if m == nil {
        return nil
    } else {
        return m.message
    }
}
func (m *CloudPcStatusDetails) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *CloudPcStatusDetails) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetAdditionalInformation() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAdditionalInformation()))
        for i, v := range m.GetAdditionalInformation() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CloudPcStatusDetails) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAdditionalInformation sets the additionalInformation property value. Any additional information about the Cloud PC status.
func (m *CloudPcStatusDetails) SetAdditionalInformation(value []KeyValuePairable)() {
    if m != nil {
        m.additionalInformation = value
    }
}
// SetCode sets the code property value. The code associated with the Cloud PC status.
func (m *CloudPcStatusDetails) SetCode(value *string)() {
    if m != nil {
        m.code = value
    }
}
// SetMessage sets the message property value. The status message.
func (m *CloudPcStatusDetails) SetMessage(value *string)() {
    if m != nil {
        m.message = value
    }
}
