package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// CloudPcLaunchInfo provides operations to call the getCloudPcLaunchInfo method.
type CloudPcLaunchInfo struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    cloudPcId *string;
    // 
    cloudPcLaunchUrl *string;
}
// NewCloudPcLaunchInfo instantiates a new cloudPcLaunchInfo and sets the default values.
func NewCloudPcLaunchInfo()(*CloudPcLaunchInfo) {
    m := &CloudPcLaunchInfo{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateCloudPcLaunchInfoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCloudPcLaunchInfoFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewCloudPcLaunchInfo(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CloudPcLaunchInfo) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetCloudPcId gets the cloudPcId property value. 
func (m *CloudPcLaunchInfo) GetCloudPcId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.cloudPcId
    }
}
// GetCloudPcLaunchUrl gets the cloudPcLaunchUrl property value. 
func (m *CloudPcLaunchInfo) GetCloudPcLaunchUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.cloudPcLaunchUrl
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CloudPcLaunchInfo) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["cloudPcId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCloudPcId(val)
        }
        return nil
    }
    res["cloudPcLaunchUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCloudPcLaunchUrl(val)
        }
        return nil
    }
    return res
}
func (m *CloudPcLaunchInfo) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *CloudPcLaunchInfo) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("cloudPcId", m.GetCloudPcId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("cloudPcLaunchUrl", m.GetCloudPcLaunchUrl())
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
func (m *CloudPcLaunchInfo) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetCloudPcId sets the cloudPcId property value. 
func (m *CloudPcLaunchInfo) SetCloudPcId(value *string)() {
    if m != nil {
        m.cloudPcId = value
    }
}
// SetCloudPcLaunchUrl sets the cloudPcLaunchUrl property value. 
func (m *CloudPcLaunchInfo) SetCloudPcLaunchUrl(value *string)() {
    if m != nil {
        m.cloudPcLaunchUrl = value
    }
}
