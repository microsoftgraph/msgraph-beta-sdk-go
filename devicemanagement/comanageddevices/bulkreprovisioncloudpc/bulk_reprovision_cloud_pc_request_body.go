package bulkreprovisioncloudpc

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type BulkReprovisionCloudPcRequestBody struct {
    additionalData map[string]interface{};
    managedDeviceIds []string;
}
func NewBulkReprovisionCloudPcRequestBody()(*BulkReprovisionCloudPcRequestBody) {
    m := &BulkReprovisionCloudPcRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *BulkReprovisionCloudPcRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *BulkReprovisionCloudPcRequestBody) GetManagedDeviceIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.managedDeviceIds
    }
}
func (m *BulkReprovisionCloudPcRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["managedDeviceIds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetManagedDeviceIds(res)
        return nil
    }
    return res
}
func (m *BulkReprovisionCloudPcRequestBody) IsNil()(bool) {
    return m == nil
}
func (m *BulkReprovisionCloudPcRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteCollectionOfStringValues("managedDeviceIds", m.GetManagedDeviceIds())
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
func (m *BulkReprovisionCloudPcRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *BulkReprovisionCloudPcRequestBody) SetManagedDeviceIds(value []string)() {
    m.managedDeviceIds = value
}
