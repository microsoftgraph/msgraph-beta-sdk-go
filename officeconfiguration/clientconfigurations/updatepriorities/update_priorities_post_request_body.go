package updatepriorities

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UpdatePrioritiesPostRequestBody provides operations to call the updatePriorities method.
type UpdatePrioritiesPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The officeConfigurationPolicyIds property
    officeConfigurationPolicyIds []string
    // The officeConfigurationPriorities property
    officeConfigurationPriorities []int32
}
// NewUpdatePrioritiesPostRequestBody instantiates a new updatePrioritiesPostRequestBody and sets the default values.
func NewUpdatePrioritiesPostRequestBody()(*UpdatePrioritiesPostRequestBody) {
    m := &UpdatePrioritiesPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateUpdatePrioritiesPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUpdatePrioritiesPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUpdatePrioritiesPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UpdatePrioritiesPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UpdatePrioritiesPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["officeConfigurationPolicyIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetOfficeConfigurationPolicyIds(res)
        }
        return nil
    }
    res["officeConfigurationPriorities"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("int32")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]int32, len(val))
            for i, v := range val {
                res[i] = *(v.(*int32))
            }
            m.SetOfficeConfigurationPriorities(res)
        }
        return nil
    }
    return res
}
// GetOfficeConfigurationPolicyIds gets the officeConfigurationPolicyIds property value. The officeConfigurationPolicyIds property
func (m *UpdatePrioritiesPostRequestBody) GetOfficeConfigurationPolicyIds()([]string) {
    return m.officeConfigurationPolicyIds
}
// GetOfficeConfigurationPriorities gets the officeConfigurationPriorities property value. The officeConfigurationPriorities property
func (m *UpdatePrioritiesPostRequestBody) GetOfficeConfigurationPriorities()([]int32) {
    return m.officeConfigurationPriorities
}
// Serialize serializes information the current object
func (m *UpdatePrioritiesPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetOfficeConfigurationPolicyIds() != nil {
        err := writer.WriteCollectionOfStringValues("officeConfigurationPolicyIds", m.GetOfficeConfigurationPolicyIds())
        if err != nil {
            return err
        }
    }
    if m.GetOfficeConfigurationPriorities() != nil {
        err := writer.WriteCollectionOfInt32Values("officeConfigurationPriorities", m.GetOfficeConfigurationPriorities())
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
func (m *UpdatePrioritiesPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetOfficeConfigurationPolicyIds sets the officeConfigurationPolicyIds property value. The officeConfigurationPolicyIds property
func (m *UpdatePrioritiesPostRequestBody) SetOfficeConfigurationPolicyIds(value []string)() {
    m.officeConfigurationPolicyIds = value
}
// SetOfficeConfigurationPriorities sets the officeConfigurationPriorities property value. The officeConfigurationPriorities property
func (m *UpdatePrioritiesPostRequestBody) SetOfficeConfigurationPriorities(value []int32)() {
    m.officeConfigurationPriorities = value
}
