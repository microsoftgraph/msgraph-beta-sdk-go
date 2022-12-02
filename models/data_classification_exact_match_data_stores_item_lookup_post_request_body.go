package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DataClassificationExactMatchDataStoresItemLookupPostRequestBody provides operations to call the lookup method.
type DataClassificationExactMatchDataStoresItemLookupPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The key property
    key *string
    // The resultColumnNames property
    resultColumnNames []string
    // The values property
    values []string
}
// NewDataClassificationExactMatchDataStoresItemLookupPostRequestBody instantiates a new DataClassificationExactMatchDataStoresItemLookupPostRequestBody and sets the default values.
func NewDataClassificationExactMatchDataStoresItemLookupPostRequestBody()(*DataClassificationExactMatchDataStoresItemLookupPostRequestBody) {
    m := &DataClassificationExactMatchDataStoresItemLookupPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDataClassificationExactMatchDataStoresItemLookupPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDataClassificationExactMatchDataStoresItemLookupPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDataClassificationExactMatchDataStoresItemLookupPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DataClassificationExactMatchDataStoresItemLookupPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DataClassificationExactMatchDataStoresItemLookupPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["key"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKey(val)
        }
        return nil
    }
    res["resultColumnNames"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetResultColumnNames(res)
        }
        return nil
    }
    res["values"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetValues(res)
        }
        return nil
    }
    return res
}
// GetKey gets the key property value. The key property
func (m *DataClassificationExactMatchDataStoresItemLookupPostRequestBody) GetKey()(*string) {
    return m.key
}
// GetResultColumnNames gets the resultColumnNames property value. The resultColumnNames property
func (m *DataClassificationExactMatchDataStoresItemLookupPostRequestBody) GetResultColumnNames()([]string) {
    return m.resultColumnNames
}
// GetValues gets the values property value. The values property
func (m *DataClassificationExactMatchDataStoresItemLookupPostRequestBody) GetValues()([]string) {
    return m.values
}
// Serialize serializes information the current object
func (m *DataClassificationExactMatchDataStoresItemLookupPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("key", m.GetKey())
        if err != nil {
            return err
        }
    }
    if m.GetResultColumnNames() != nil {
        err := writer.WriteCollectionOfStringValues("resultColumnNames", m.GetResultColumnNames())
        if err != nil {
            return err
        }
    }
    if m.GetValues() != nil {
        err := writer.WriteCollectionOfStringValues("values", m.GetValues())
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
func (m *DataClassificationExactMatchDataStoresItemLookupPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetKey sets the key property value. The key property
func (m *DataClassificationExactMatchDataStoresItemLookupPostRequestBody) SetKey(value *string)() {
    m.key = value
}
// SetResultColumnNames sets the resultColumnNames property value. The resultColumnNames property
func (m *DataClassificationExactMatchDataStoresItemLookupPostRequestBody) SetResultColumnNames(value []string)() {
    m.resultColumnNames = value
}
// SetValues sets the values property value. The values property
func (m *DataClassificationExactMatchDataStoresItemLookupPostRequestBody) SetValues(value []string)() {
    m.values = value
}
