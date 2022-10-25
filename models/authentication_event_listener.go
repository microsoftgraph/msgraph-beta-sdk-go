package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AuthenticationEventListener provides operations to manage the collection of accessReviewDecision entities.
type AuthenticationEventListener struct {
    Entity
    // The authenticationEventsFlowId property
    authenticationEventsFlowId *string
    // The conditions property
    conditions AuthenticationConditionsable
    // The priority property
    priority *int32
    // The tags property
    tags []KeyValuePairable
}
// NewAuthenticationEventListener instantiates a new authenticationEventListener and sets the default values.
func NewAuthenticationEventListener()(*AuthenticationEventListener) {
    m := &AuthenticationEventListener{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.authenticationEventListener";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateAuthenticationEventListenerFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAuthenticationEventListenerFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.onTokenIssuanceStartListener":
                        return NewOnTokenIssuanceStartListener(), nil
                }
            }
        }
    }
    return NewAuthenticationEventListener(), nil
}
// GetAuthenticationEventsFlowId gets the authenticationEventsFlowId property value. The authenticationEventsFlowId property
func (m *AuthenticationEventListener) GetAuthenticationEventsFlowId()(*string) {
    return m.authenticationEventsFlowId
}
// GetConditions gets the conditions property value. The conditions property
func (m *AuthenticationEventListener) GetConditions()(AuthenticationConditionsable) {
    return m.conditions
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AuthenticationEventListener) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["authenticationEventsFlowId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetAuthenticationEventsFlowId)
    res["conditions"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateAuthenticationConditionsFromDiscriminatorValue , m.SetConditions)
    res["priority"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetPriority)
    res["tags"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateKeyValuePairFromDiscriminatorValue , m.SetTags)
    return res
}
// GetPriority gets the priority property value. The priority property
func (m *AuthenticationEventListener) GetPriority()(*int32) {
    return m.priority
}
// GetTags gets the tags property value. The tags property
func (m *AuthenticationEventListener) GetTags()([]KeyValuePairable) {
    return m.tags
}
// Serialize serializes information the current object
func (m *AuthenticationEventListener) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("authenticationEventsFlowId", m.GetAuthenticationEventsFlowId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("conditions", m.GetConditions())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("priority", m.GetPriority())
        if err != nil {
            return err
        }
    }
    if m.GetTags() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetTags())
        err = writer.WriteCollectionOfObjectValues("tags", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAuthenticationEventsFlowId sets the authenticationEventsFlowId property value. The authenticationEventsFlowId property
func (m *AuthenticationEventListener) SetAuthenticationEventsFlowId(value *string)() {
    m.authenticationEventsFlowId = value
}
// SetConditions sets the conditions property value. The conditions property
func (m *AuthenticationEventListener) SetConditions(value AuthenticationConditionsable)() {
    m.conditions = value
}
// SetPriority sets the priority property value. The priority property
func (m *AuthenticationEventListener) SetPriority(value *int32)() {
    m.priority = value
}
// SetTags sets the tags property value. The tags property
func (m *AuthenticationEventListener) SetTags(value []KeyValuePairable)() {
    m.tags = value
}
