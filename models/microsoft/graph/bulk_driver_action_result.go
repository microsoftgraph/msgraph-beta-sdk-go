package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type BulkDriverActionResult struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // List of driver Ids where the action is failed.
    failedDriverIds []string;
    // List of driver Ids that are not found.
    notFoundDriverIds []string;
    // List of driver Ids where the action is successful.
    successfulDriverIds []string;
}
// Instantiates a new bulkDriverActionResult and sets the default values.
func NewBulkDriverActionResult()(*BulkDriverActionResult) {
    m := &BulkDriverActionResult{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *BulkDriverActionResult) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the failedDriverIds property value. List of driver Ids where the action is failed.
func (m *BulkDriverActionResult) GetFailedDriverIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.failedDriverIds
    }
}
// Gets the notFoundDriverIds property value. List of driver Ids that are not found.
func (m *BulkDriverActionResult) GetNotFoundDriverIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.notFoundDriverIds
    }
}
// Gets the successfulDriverIds property value. List of driver Ids where the action is successful.
func (m *BulkDriverActionResult) GetSuccessfulDriverIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.successfulDriverIds
    }
}
// The deserialization information for the current model
func (m *BulkDriverActionResult) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["failedDriverIds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetFailedDriverIds(res)
        return nil
    }
    res["notFoundDriverIds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetNotFoundDriverIds(res)
        return nil
    }
    res["successfulDriverIds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetSuccessfulDriverIds(res)
        return nil
    }
    return res
}
func (m *BulkDriverActionResult) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *BulkDriverActionResult) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteCollectionOfStringValues("failedDriverIds", m.GetFailedDriverIds())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfStringValues("notFoundDriverIds", m.GetNotFoundDriverIds())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfStringValues("successfulDriverIds", m.GetSuccessfulDriverIds())
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
func (m *BulkDriverActionResult) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the failedDriverIds property value. List of driver Ids where the action is failed.
// Parameters:
//  - value : Value to set for the failedDriverIds property.
func (m *BulkDriverActionResult) SetFailedDriverIds(value []string)() {
    m.failedDriverIds = value
}
// Sets the notFoundDriverIds property value. List of driver Ids that are not found.
// Parameters:
//  - value : Value to set for the notFoundDriverIds property.
func (m *BulkDriverActionResult) SetNotFoundDriverIds(value []string)() {
    m.notFoundDriverIds = value
}
// Sets the successfulDriverIds property value. List of driver Ids where the action is successful.
// Parameters:
//  - value : Value to set for the successfulDriverIds property.
func (m *BulkDriverActionResult) SetSuccessfulDriverIds(value []string)() {
    m.successfulDriverIds = value
}
