package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// BulkDriverActionResult 
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
// NewBulkDriverActionResult instantiates a new bulkDriverActionResult and sets the default values.
func NewBulkDriverActionResult()(*BulkDriverActionResult) {
    m := &BulkDriverActionResult{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *BulkDriverActionResult) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFailedDriverIds gets the failedDriverIds property value. List of driver Ids where the action is failed.
func (m *BulkDriverActionResult) GetFailedDriverIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.failedDriverIds
    }
}
// GetNotFoundDriverIds gets the notFoundDriverIds property value. List of driver Ids that are not found.
func (m *BulkDriverActionResult) GetNotFoundDriverIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.notFoundDriverIds
    }
}
// GetSuccessfulDriverIds gets the successfulDriverIds property value. List of driver Ids where the action is successful.
func (m *BulkDriverActionResult) GetSuccessfulDriverIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.successfulDriverIds
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *BulkDriverActionResult) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["failedDriverIds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetFailedDriverIds(res)
        }
        return nil
    }
    res["notFoundDriverIds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetNotFoundDriverIds(res)
        }
        return nil
    }
    res["successfulDriverIds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetSuccessfulDriverIds(res)
        }
        return nil
    }
    return res
}
func (m *BulkDriverActionResult) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *BulkDriverActionResult) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetFailedDriverIds() != nil {
        err := writer.WriteCollectionOfStringValues("failedDriverIds", m.GetFailedDriverIds())
        if err != nil {
            return err
        }
    }
    if m.GetNotFoundDriverIds() != nil {
        err := writer.WriteCollectionOfStringValues("notFoundDriverIds", m.GetNotFoundDriverIds())
        if err != nil {
            return err
        }
    }
    if m.GetSuccessfulDriverIds() != nil {
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *BulkDriverActionResult) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetFailedDriverIds sets the failedDriverIds property value. List of driver Ids where the action is failed.
func (m *BulkDriverActionResult) SetFailedDriverIds(value []string)() {
    if m != nil {
        m.failedDriverIds = value
    }
}
// SetNotFoundDriverIds sets the notFoundDriverIds property value. List of driver Ids that are not found.
func (m *BulkDriverActionResult) SetNotFoundDriverIds(value []string)() {
    if m != nil {
        m.notFoundDriverIds = value
    }
}
// SetSuccessfulDriverIds sets the successfulDriverIds property value. List of driver Ids where the action is successful.
func (m *BulkDriverActionResult) SetSuccessfulDriverIds(value []string)() {
    if m != nil {
        m.successfulDriverIds = value
    }
}
