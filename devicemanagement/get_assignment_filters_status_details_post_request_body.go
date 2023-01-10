package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GetAssignmentFiltersStatusDetailsPostRequestBody 
type GetAssignmentFiltersStatusDetailsPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The assignmentFilterIds property
    assignmentFilterIds []string
    // The managedDeviceId property
    managedDeviceId *string
    // The payloadId property
    payloadId *string
    // The skip property
    skip *int32
    // The top property
    top *int32
    // The userId property
    userId *string
}
// NewGetAssignmentFiltersStatusDetailsPostRequestBody instantiates a new GetAssignmentFiltersStatusDetailsPostRequestBody and sets the default values.
func NewGetAssignmentFiltersStatusDetailsPostRequestBody()(*GetAssignmentFiltersStatusDetailsPostRequestBody) {
    m := &GetAssignmentFiltersStatusDetailsPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateGetAssignmentFiltersStatusDetailsPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGetAssignmentFiltersStatusDetailsPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGetAssignmentFiltersStatusDetailsPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *GetAssignmentFiltersStatusDetailsPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetAssignmentFilterIds gets the assignmentFilterIds property value. The assignmentFilterIds property
func (m *GetAssignmentFiltersStatusDetailsPostRequestBody) GetAssignmentFilterIds()([]string) {
    return m.assignmentFilterIds
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GetAssignmentFiltersStatusDetailsPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["assignmentFilterIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetAssignmentFilterIds(res)
        }
        return nil
    }
    res["managedDeviceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagedDeviceId(val)
        }
        return nil
    }
    res["payloadId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPayloadId(val)
        }
        return nil
    }
    res["skip"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSkip(val)
        }
        return nil
    }
    res["top"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTop(val)
        }
        return nil
    }
    res["userId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserId(val)
        }
        return nil
    }
    return res
}
// GetManagedDeviceId gets the managedDeviceId property value. The managedDeviceId property
func (m *GetAssignmentFiltersStatusDetailsPostRequestBody) GetManagedDeviceId()(*string) {
    return m.managedDeviceId
}
// GetPayloadId gets the payloadId property value. The payloadId property
func (m *GetAssignmentFiltersStatusDetailsPostRequestBody) GetPayloadId()(*string) {
    return m.payloadId
}
// GetSkip gets the skip property value. The skip property
func (m *GetAssignmentFiltersStatusDetailsPostRequestBody) GetSkip()(*int32) {
    return m.skip
}
// GetTop gets the top property value. The top property
func (m *GetAssignmentFiltersStatusDetailsPostRequestBody) GetTop()(*int32) {
    return m.top
}
// GetUserId gets the userId property value. The userId property
func (m *GetAssignmentFiltersStatusDetailsPostRequestBody) GetUserId()(*string) {
    return m.userId
}
// Serialize serializes information the current object
func (m *GetAssignmentFiltersStatusDetailsPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAssignmentFilterIds() != nil {
        err := writer.WriteCollectionOfStringValues("assignmentFilterIds", m.GetAssignmentFilterIds())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("managedDeviceId", m.GetManagedDeviceId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("payloadId", m.GetPayloadId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("skip", m.GetSkip())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("top", m.GetTop())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("userId", m.GetUserId())
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
func (m *GetAssignmentFiltersStatusDetailsPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetAssignmentFilterIds sets the assignmentFilterIds property value. The assignmentFilterIds property
func (m *GetAssignmentFiltersStatusDetailsPostRequestBody) SetAssignmentFilterIds(value []string)() {
    m.assignmentFilterIds = value
}
// SetManagedDeviceId sets the managedDeviceId property value. The managedDeviceId property
func (m *GetAssignmentFiltersStatusDetailsPostRequestBody) SetManagedDeviceId(value *string)() {
    m.managedDeviceId = value
}
// SetPayloadId sets the payloadId property value. The payloadId property
func (m *GetAssignmentFiltersStatusDetailsPostRequestBody) SetPayloadId(value *string)() {
    m.payloadId = value
}
// SetSkip sets the skip property value. The skip property
func (m *GetAssignmentFiltersStatusDetailsPostRequestBody) SetSkip(value *int32)() {
    m.skip = value
}
// SetTop sets the top property value. The top property
func (m *GetAssignmentFiltersStatusDetailsPostRequestBody) SetTop(value *int32)() {
    m.top = value
}
// SetUserId sets the userId property value. The userId property
func (m *GetAssignmentFiltersStatusDetailsPostRequestBody) SetUserId(value *string)() {
    m.userId = value
}
