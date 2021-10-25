package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type BulkManagedDeviceActionResult struct {
    additionalData map[string]interface{};
    failedDeviceIds []string;
    notFoundDeviceIds []string;
    notSupportedDeviceIds []string;
    successfulDeviceIds []string;
}
func NewBulkManagedDeviceActionResult()(*BulkManagedDeviceActionResult) {
    m := &BulkManagedDeviceActionResult{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *BulkManagedDeviceActionResult) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *BulkManagedDeviceActionResult) GetFailedDeviceIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.failedDeviceIds
    }
}
func (m *BulkManagedDeviceActionResult) GetNotFoundDeviceIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.notFoundDeviceIds
    }
}
func (m *BulkManagedDeviceActionResult) GetNotSupportedDeviceIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.notSupportedDeviceIds
    }
}
func (m *BulkManagedDeviceActionResult) GetSuccessfulDeviceIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.successfulDeviceIds
    }
}
func (m *BulkManagedDeviceActionResult) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["failedDeviceIds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetFailedDeviceIds(res)
        return nil
    }
    res["notFoundDeviceIds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetNotFoundDeviceIds(res)
        return nil
    }
    res["notSupportedDeviceIds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetNotSupportedDeviceIds(res)
        return nil
    }
    res["successfulDeviceIds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetSuccessfulDeviceIds(res)
        return nil
    }
    return res
}
func (m *BulkManagedDeviceActionResult) IsNil()(bool) {
    return m == nil
}
func (m *BulkManagedDeviceActionResult) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteCollectionOfStringValues("failedDeviceIds", m.GetFailedDeviceIds())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfStringValues("notFoundDeviceIds", m.GetNotFoundDeviceIds())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfStringValues("notSupportedDeviceIds", m.GetNotSupportedDeviceIds())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfStringValues("successfulDeviceIds", m.GetSuccessfulDeviceIds())
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
func (m *BulkManagedDeviceActionResult) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *BulkManagedDeviceActionResult) SetFailedDeviceIds(value []string)() {
    m.failedDeviceIds = value
}
func (m *BulkManagedDeviceActionResult) SetNotFoundDeviceIds(value []string)() {
    m.notFoundDeviceIds = value
}
func (m *BulkManagedDeviceActionResult) SetNotSupportedDeviceIds(value []string)() {
    m.notSupportedDeviceIds = value
}
func (m *BulkManagedDeviceActionResult) SetSuccessfulDeviceIds(value []string)() {
    m.successfulDeviceIds = value
}
