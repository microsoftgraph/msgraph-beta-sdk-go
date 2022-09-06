package updateallowedcombinations

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UpdateAllowedCombinationsPostRequestBody provides operations to call the updateAllowedCombinations method.
type UpdateAllowedCombinationsPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The allowedCombinations property
    allowedCombinations []string
}
// NewUpdateAllowedCombinationsPostRequestBody instantiates a new updateAllowedCombinationsPostRequestBody and sets the default values.
func NewUpdateAllowedCombinationsPostRequestBody()(*UpdateAllowedCombinationsPostRequestBody) {
    m := &UpdateAllowedCombinationsPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateUpdateAllowedCombinationsPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUpdateAllowedCombinationsPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUpdateAllowedCombinationsPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UpdateAllowedCombinationsPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetAllowedCombinations gets the allowedCombinations property value. The allowedCombinations property
func (m *UpdateAllowedCombinationsPostRequestBody) GetAllowedCombinations()([]string) {
    return m.allowedCombinations
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UpdateAllowedCombinationsPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["allowedCombinations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetAllowedCombinations(res)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *UpdateAllowedCombinationsPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAllowedCombinations() != nil {
        err := writer.WriteCollectionOfStringValues("allowedCombinations", m.GetAllowedCombinations())
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
func (m *UpdateAllowedCombinationsPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetAllowedCombinations sets the allowedCombinations property value. The allowedCombinations property
func (m *UpdateAllowedCombinationsPostRequestBody) SetAllowedCombinations(value []string)() {
    m.allowedCombinations = value
}
