package restorecloudpc

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// RestoreCloudPcRequestBody 
type RestoreCloudPcRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    cloudPcSnapshotId *string;
}
// NewRestoreCloudPcRequestBody instantiates a new restoreCloudPcRequestBody and sets the default values.
func NewRestoreCloudPcRequestBody()(*RestoreCloudPcRequestBody) {
    m := &RestoreCloudPcRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RestoreCloudPcRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetCloudPcSnapshotId gets the cloudPcSnapshotId property value. 
func (m *RestoreCloudPcRequestBody) GetCloudPcSnapshotId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.cloudPcSnapshotId
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RestoreCloudPcRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["cloudPcSnapshotId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCloudPcSnapshotId(val)
        }
        return nil
    }
    return res
}
func (m *RestoreCloudPcRequestBody) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *RestoreCloudPcRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("cloudPcSnapshotId", m.GetCloudPcSnapshotId())
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
func (m *RestoreCloudPcRequestBody) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetCloudPcSnapshotId sets the cloudPcSnapshotId property value. 
func (m *RestoreCloudPcRequestBody) SetCloudPcSnapshotId(value *string)() {
    if m != nil {
        m.cloudPcSnapshotId = value
    }
}
