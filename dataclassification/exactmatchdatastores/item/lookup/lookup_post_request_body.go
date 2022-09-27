package lookup

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// LookupPostRequestBody provides operations to call the lookup method.
type LookupPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The key property
    key *string
    // The resultColumnNames property
    resultColumnNames []string
    // The values property
    values []string
}
// NewLookupPostRequestBody instantiates a new lookupPostRequestBody and sets the default values.
func NewLookupPostRequestBody()(*LookupPostRequestBody) {
    m := &LookupPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateLookupPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateLookupPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewLookupPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *LookupPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *LookupPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["key"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetKey)
    res["resultColumnNames"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetResultColumnNames)
    res["values"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetValues)
    return res
}
// GetKey gets the key property value. The key property
func (m *LookupPostRequestBody) GetKey()(*string) {
    return m.key
}
// GetResultColumnNames gets the resultColumnNames property value. The resultColumnNames property
func (m *LookupPostRequestBody) GetResultColumnNames()([]string) {
    return m.resultColumnNames
}
// GetValues gets the values property value. The values property
func (m *LookupPostRequestBody) GetValues()([]string) {
    return m.values
}
// Serialize serializes information the current object
func (m *LookupPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *LookupPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetKey sets the key property value. The key property
func (m *LookupPostRequestBody) SetKey(value *string)() {
    m.key = value
}
// SetResultColumnNames sets the resultColumnNames property value. The resultColumnNames property
func (m *LookupPostRequestBody) SetResultColumnNames(value []string)() {
    m.resultColumnNames = value
}
// SetValues sets the values property value. The values property
func (m *LookupPostRequestBody) SetValues(value []string)() {
    m.values = value
}
