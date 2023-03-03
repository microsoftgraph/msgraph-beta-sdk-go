package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsAutopilotDeploymentProfileCollectionResponse 
type WindowsAutopilotDeploymentProfileCollectionResponse struct {
    BaseCollectionPaginationCountResponse
}
// NewWindowsAutopilotDeploymentProfileCollectionResponse instantiates a new WindowsAutopilotDeploymentProfileCollectionResponse and sets the default values.
func NewWindowsAutopilotDeploymentProfileCollectionResponse()(*WindowsAutopilotDeploymentProfileCollectionResponse) {
    m := &WindowsAutopilotDeploymentProfileCollectionResponse{
        BaseCollectionPaginationCountResponse: *NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateWindowsAutopilotDeploymentProfileCollectionResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsAutopilotDeploymentProfileCollectionResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsAutopilotDeploymentProfileCollectionResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsAutopilotDeploymentProfileCollectionResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWindowsAutopilotDeploymentProfileFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WindowsAutopilotDeploymentProfileable, len(val))
            for i, v := range val {
                res[i] = v.(WindowsAutopilotDeploymentProfileable)
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *WindowsAutopilotDeploymentProfileCollectionResponse) GetValue()([]WindowsAutopilotDeploymentProfileable) {
    val, err := m.GetBackingStore().Get("value")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]WindowsAutopilotDeploymentProfileable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *WindowsAutopilotDeploymentProfileCollectionResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.BaseCollectionPaginationCountResponse.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetValue() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetValue()))
        for i, v := range m.GetValue() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("value", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetValue sets the value property value. The value property
func (m *WindowsAutopilotDeploymentProfileCollectionResponse) SetValue(value []WindowsAutopilotDeploymentProfileable)() {
    err := m.GetBackingStore().Set("value", value)
    if err != nil {
        panic(err)
    }
}
// WindowsAutopilotDeploymentProfileCollectionResponseable 
type WindowsAutopilotDeploymentProfileCollectionResponseable interface {
    BaseCollectionPaginationCountResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetValue()([]WindowsAutopilotDeploymentProfileable)
    SetValue(value []WindowsAutopilotDeploymentProfileable)()
}
