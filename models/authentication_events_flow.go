package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AuthenticationEventsFlow 
type AuthenticationEventsFlow struct {
    Entity
    // The OdataType property
    OdataType *string
}
// NewAuthenticationEventsFlow instantiates a new authenticationEventsFlow and sets the default values.
func NewAuthenticationEventsFlow()(*AuthenticationEventsFlow) {
    m := &AuthenticationEventsFlow{
        Entity: *NewEntity(),
    }
    return m
}
// CreateAuthenticationEventsFlowFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAuthenticationEventsFlowFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.externalUsersSelfServiceSignUpEventsFlow":
                        return NewExternalUsersSelfServiceSignUpEventsFlow(), nil
                }
            }
        }
    }
    return NewAuthenticationEventsFlow(), nil
}
// GetConditions gets the conditions property value. The conditions representing the context of the authentication request which will be used to decide whether the events policy will be invoked.
func (m *AuthenticationEventsFlow) GetConditions()(AuthenticationConditionsable) {
    val, err := m.GetBackingStore().Get("conditions")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(AuthenticationConditionsable)
    }
    return nil
}
// GetDescription gets the description property value. The description of the events policy.
func (m *AuthenticationEventsFlow) GetDescription()(*string) {
    val, err := m.GetBackingStore().Get("description")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDisplayName gets the displayName property value. Required. The display name for the events policy.
func (m *AuthenticationEventsFlow) GetDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("displayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AuthenticationEventsFlow) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["conditions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAuthenticationConditionsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConditions(val.(AuthenticationConditionsable))
        }
        return nil
    }
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["priority"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPriority(val)
        }
        return nil
    }
    return res
}
// GetPriority gets the priority property value. The priority to use for each individual event of the events policy. If multiple competing listeners for an event have the same priority, one is chosen and an error is silently logged. Defaults to 500.
func (m *AuthenticationEventsFlow) GetPriority()(*int32) {
    val, err := m.GetBackingStore().Get("priority")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// Serialize serializes information the current object
func (m *AuthenticationEventsFlow) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("conditions", m.GetConditions())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
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
    return nil
}
// SetConditions sets the conditions property value. The conditions representing the context of the authentication request which will be used to decide whether the events policy will be invoked.
func (m *AuthenticationEventsFlow) SetConditions(value AuthenticationConditionsable)() {
    err := m.GetBackingStore().Set("conditions", value)
    if err != nil {
        panic(err)
    }
}
// SetDescription sets the description property value. The description of the events policy.
func (m *AuthenticationEventsFlow) SetDescription(value *string)() {
    err := m.GetBackingStore().Set("description", value)
    if err != nil {
        panic(err)
    }
}
// SetDisplayName sets the displayName property value. Required. The display name for the events policy.
func (m *AuthenticationEventsFlow) SetDisplayName(value *string)() {
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// SetPriority sets the priority property value. The priority to use for each individual event of the events policy. If multiple competing listeners for an event have the same priority, one is chosen and an error is silently logged. Defaults to 500.
func (m *AuthenticationEventsFlow) SetPriority(value *int32)() {
    err := m.GetBackingStore().Set("priority", value)
    if err != nil {
        panic(err)
    }
}
// AuthenticationEventsFlowable 
type AuthenticationEventsFlowable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetConditions()(AuthenticationConditionsable)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetPriority()(*int32)
    SetConditions(value AuthenticationConditionsable)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetPriority(value *int32)()
}
