package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TargetPolicyEndpoints 
type TargetPolicyEndpoints struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Use to filter the notification distribution to a specific platform or platforms. Valid values are Windows, iOS, Android and WebPush. By default, all push endpoint types (Windows, iOS, Android and WebPush) are enabled.
    platformTypes []string
}
// NewTargetPolicyEndpoints instantiates a new targetPolicyEndpoints and sets the default values.
func NewTargetPolicyEndpoints()(*TargetPolicyEndpoints) {
    m := &TargetPolicyEndpoints{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateTargetPolicyEndpointsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTargetPolicyEndpointsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTargetPolicyEndpoints(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TargetPolicyEndpoints) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TargetPolicyEndpoints) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["platformTypes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetPlatformTypes(res)
        }
        return nil
    }
    return res
}
// GetPlatformTypes gets the platformTypes property value. Use to filter the notification distribution to a specific platform or platforms. Valid values are Windows, iOS, Android and WebPush. By default, all push endpoint types (Windows, iOS, Android and WebPush) are enabled.
func (m *TargetPolicyEndpoints) GetPlatformTypes()([]string) {
    if m == nil {
        return nil
    } else {
        return m.platformTypes
    }
}
// Serialize serializes information the current object
func (m *TargetPolicyEndpoints) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetPlatformTypes() != nil {
        err := writer.WriteCollectionOfStringValues("platformTypes", m.GetPlatformTypes())
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
func (m *TargetPolicyEndpoints) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetPlatformTypes sets the platformTypes property value. Use to filter the notification distribution to a specific platform or platforms. Valid values are Windows, iOS, Android and WebPush. By default, all push endpoint types (Windows, iOS, Android and WebPush) are enabled.
func (m *TargetPolicyEndpoints) SetPlatformTypes(value []string)() {
    if m != nil {
        m.platformTypes = value
    }
}
