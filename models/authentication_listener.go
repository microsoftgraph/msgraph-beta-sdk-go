package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AuthenticationListener provides operations to manage the collection of accessReviewDecision entities.
type AuthenticationListener struct {
    Entity
    // The priority of the listener. Determines the order of evaluation when an event has multiple listeners. The priority is evaluated from low to high.
    priority *int32
    // Filter based on the source of the authentication that is used to determine whether the listener is evaluated. This is currently limited to evaluations based on application the user is authenticating to.
    sourceFilter AuthenticationSourceFilterable
}
// NewAuthenticationListener instantiates a new authenticationListener and sets the default values.
func NewAuthenticationListener()(*AuthenticationListener) {
    m := &AuthenticationListener{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.authenticationListener";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateAuthenticationListenerFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAuthenticationListenerFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "#microsoft.graph.invokeUserFlowListener":
                        return NewInvokeUserFlowListener(), nil
                }
            }
        }
    }
    return NewAuthenticationListener(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AuthenticationListener) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["priority"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetPriority)
    res["sourceFilter"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateAuthenticationSourceFilterFromDiscriminatorValue , m.SetSourceFilter)
    return res
}
// GetPriority gets the priority property value. The priority of the listener. Determines the order of evaluation when an event has multiple listeners. The priority is evaluated from low to high.
func (m *AuthenticationListener) GetPriority()(*int32) {
    return m.priority
}
// GetSourceFilter gets the sourceFilter property value. Filter based on the source of the authentication that is used to determine whether the listener is evaluated. This is currently limited to evaluations based on application the user is authenticating to.
func (m *AuthenticationListener) GetSourceFilter()(AuthenticationSourceFilterable) {
    return m.sourceFilter
}
// Serialize serializes information the current object
func (m *AuthenticationListener) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt32Value("priority", m.GetPriority())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("sourceFilter", m.GetSourceFilter())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetPriority sets the priority property value. The priority of the listener. Determines the order of evaluation when an event has multiple listeners. The priority is evaluated from low to high.
func (m *AuthenticationListener) SetPriority(value *int32)() {
    m.priority = value
}
// SetSourceFilter sets the sourceFilter property value. Filter based on the source of the authentication that is used to determine whether the listener is evaluated. This is currently limited to evaluations based on application the user is authenticating to.
func (m *AuthenticationListener) SetSourceFilter(value AuthenticationSourceFilterable)() {
    m.sourceFilter = value
}
