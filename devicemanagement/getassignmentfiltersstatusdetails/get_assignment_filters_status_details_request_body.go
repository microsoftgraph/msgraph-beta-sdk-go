package getassignmentfiltersstatusdetails

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GetAssignmentFiltersStatusDetailsRequestBody provides operations to call the getAssignmentFiltersStatusDetails method.
type GetAssignmentFiltersStatusDetailsRequestBody struct {
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
// NewGetAssignmentFiltersStatusDetailsRequestBody instantiates a new getAssignmentFiltersStatusDetailsRequestBody and sets the default values.
func NewGetAssignmentFiltersStatusDetailsRequestBody()(*GetAssignmentFiltersStatusDetailsRequestBody) {
    m := &GetAssignmentFiltersStatusDetailsRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateGetAssignmentFiltersStatusDetailsRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGetAssignmentFiltersStatusDetailsRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGetAssignmentFiltersStatusDetailsRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *GetAssignmentFiltersStatusDetailsRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAssignmentFilterIds gets the assignmentFilterIds property value. The assignmentFilterIds property
func (m *GetAssignmentFiltersStatusDetailsRequestBody) GetAssignmentFilterIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.assignmentFilterIds
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GetAssignmentFiltersStatusDetailsRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
func (m *GetAssignmentFiltersStatusDetailsRequestBody) GetManagedDeviceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.managedDeviceId
    }
}
// GetPayloadId gets the payloadId property value. The payloadId property
func (m *GetAssignmentFiltersStatusDetailsRequestBody) GetPayloadId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.payloadId
    }
}
// GetSkip gets the skip property value. The skip property
func (m *GetAssignmentFiltersStatusDetailsRequestBody) GetSkip()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.skip
    }
}
// GetTop gets the top property value. The top property
func (m *GetAssignmentFiltersStatusDetailsRequestBody) GetTop()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.top
    }
}
// GetUserId gets the userId property value. The userId property
func (m *GetAssignmentFiltersStatusDetailsRequestBody) GetUserId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userId
    }
}
// Serialize serializes information the current object
func (m *GetAssignmentFiltersStatusDetailsRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *GetAssignmentFiltersStatusDetailsRequestBody) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAssignmentFilterIds sets the assignmentFilterIds property value. The assignmentFilterIds property
func (m *GetAssignmentFiltersStatusDetailsRequestBody) SetAssignmentFilterIds(value []string)() {
    if m != nil {
        m.assignmentFilterIds = value
    }
}
// SetManagedDeviceId sets the managedDeviceId property value. The managedDeviceId property
func (m *GetAssignmentFiltersStatusDetailsRequestBody) SetManagedDeviceId(value *string)() {
    if m != nil {
        m.managedDeviceId = value
    }
}
// SetPayloadId sets the payloadId property value. The payloadId property
func (m *GetAssignmentFiltersStatusDetailsRequestBody) SetPayloadId(value *string)() {
    if m != nil {
        m.payloadId = value
    }
}
// SetSkip sets the skip property value. The skip property
func (m *GetAssignmentFiltersStatusDetailsRequestBody) SetSkip(value *int32)() {
    if m != nil {
        m.skip = value
    }
}
// SetTop sets the top property value. The top property
func (m *GetAssignmentFiltersStatusDetailsRequestBody) SetTop(value *int32)() {
    if m != nil {
        m.top = value
    }
}
// SetUserId sets the userId property value. The userId property
func (m *GetAssignmentFiltersStatusDetailsRequestBody) SetUserId(value *string)() {
    if m != nil {
        m.userId = value
    }
}
