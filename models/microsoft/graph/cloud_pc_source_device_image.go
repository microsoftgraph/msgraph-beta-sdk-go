package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// CloudPcSourceDeviceImage provides operations to call the getSourceImages method.
type CloudPcSourceDeviceImage struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The display name for the source image.
    displayName *string;
    // The ID of the source image.
    id *string;
    // 
    subscriptionDisplayName *string;
    // 
    subscriptionId *string;
}
// NewCloudPcSourceDeviceImage instantiates a new cloudPcSourceDeviceImage and sets the default values.
func NewCloudPcSourceDeviceImage()(*CloudPcSourceDeviceImage) {
    m := &CloudPcSourceDeviceImage{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateCloudPcSourceDeviceImageFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCloudPcSourceDeviceImageFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewCloudPcSourceDeviceImage(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CloudPcSourceDeviceImage) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDisplayName gets the displayName property value. The display name for the source image.
func (m *CloudPcSourceDeviceImage) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CloudPcSourceDeviceImage) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["id"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
        }
        return nil
    }
    res["subscriptionDisplayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubscriptionDisplayName(val)
        }
        return nil
    }
    res["subscriptionId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubscriptionId(val)
        }
        return nil
    }
    return res
}
// GetId gets the id property value. The ID of the source image.
func (m *CloudPcSourceDeviceImage) GetId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.id
    }
}
// GetSubscriptionDisplayName gets the subscriptionDisplayName property value. 
func (m *CloudPcSourceDeviceImage) GetSubscriptionDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.subscriptionDisplayName
    }
}
// GetSubscriptionId gets the subscriptionId property value. 
func (m *CloudPcSourceDeviceImage) GetSubscriptionId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.subscriptionId
    }
}
func (m *CloudPcSourceDeviceImage) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *CloudPcSourceDeviceImage) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("subscriptionDisplayName", m.GetSubscriptionDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("subscriptionId", m.GetSubscriptionId())
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
func (m *CloudPcSourceDeviceImage) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetDisplayName sets the displayName property value. The display name for the source image.
func (m *CloudPcSourceDeviceImage) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetId sets the id property value. The ID of the source image.
func (m *CloudPcSourceDeviceImage) SetId(value *string)() {
    if m != nil {
        m.id = value
    }
}
// SetSubscriptionDisplayName sets the subscriptionDisplayName property value. 
func (m *CloudPcSourceDeviceImage) SetSubscriptionDisplayName(value *string)() {
    if m != nil {
        m.subscriptionDisplayName = value
    }
}
// SetSubscriptionId sets the subscriptionId property value. 
func (m *CloudPcSourceDeviceImage) SetSubscriptionId(value *string)() {
    if m != nil {
        m.subscriptionId = value
    }
}
