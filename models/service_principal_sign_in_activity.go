package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ServicePrincipalSignInActivity struct {
    Entity
}
// NewServicePrincipalSignInActivity instantiates a new ServicePrincipalSignInActivity and sets the default values.
func NewServicePrincipalSignInActivity()(*ServicePrincipalSignInActivity) {
    m := &ServicePrincipalSignInActivity{
        Entity: *NewEntity(),
    }
    return m
}
// CreateServicePrincipalSignInActivityFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateServicePrincipalSignInActivityFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewServicePrincipalSignInActivity(), nil
}
// GetAppId gets the appId property value. The globally unique appId (also called client ID on the Microsoft Entra admin center) of the credentialed resource application.
// returns a *string when successful
func (m *ServicePrincipalSignInActivity) GetAppId()(*string) {
    val, err := m.GetBackingStore().Get("appId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetApplicationAuthenticationClientSignInActivity gets the applicationAuthenticationClientSignInActivity property value. The sign-in activity of the application in a app-only authentication flow (app-to-app tokens) where the application acts like a client.
// returns a SignInActivityable when successful
func (m *ServicePrincipalSignInActivity) GetApplicationAuthenticationClientSignInActivity()(SignInActivityable) {
    val, err := m.GetBackingStore().Get("applicationAuthenticationClientSignInActivity")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(SignInActivityable)
    }
    return nil
}
// GetApplicationAuthenticationResourceSignInActivity gets the applicationAuthenticationResourceSignInActivity property value. The sign-in activity of the application in a app-only authentication flow (app-to-app tokens) where the application acts like a resource.
// returns a SignInActivityable when successful
func (m *ServicePrincipalSignInActivity) GetApplicationAuthenticationResourceSignInActivity()(SignInActivityable) {
    val, err := m.GetBackingStore().Get("applicationAuthenticationResourceSignInActivity")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(SignInActivityable)
    }
    return nil
}
// GetDelegatedClientSignInActivity gets the delegatedClientSignInActivity property value. The sign-in activity of the application in a delegated flow (user sign-in) where the application acts like a client.
// returns a SignInActivityable when successful
func (m *ServicePrincipalSignInActivity) GetDelegatedClientSignInActivity()(SignInActivityable) {
    val, err := m.GetBackingStore().Get("delegatedClientSignInActivity")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(SignInActivityable)
    }
    return nil
}
// GetDelegatedResourceSignInActivity gets the delegatedResourceSignInActivity property value. The sign-in activity of the application in a delegated flow (user sign-in) where the application acts like a resource.
// returns a SignInActivityable when successful
func (m *ServicePrincipalSignInActivity) GetDelegatedResourceSignInActivity()(SignInActivityable) {
    val, err := m.GetBackingStore().Get("delegatedResourceSignInActivity")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(SignInActivityable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ServicePrincipalSignInActivity) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["appId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppId(val)
        }
        return nil
    }
    res["applicationAuthenticationClientSignInActivity"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSignInActivityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplicationAuthenticationClientSignInActivity(val.(SignInActivityable))
        }
        return nil
    }
    res["applicationAuthenticationResourceSignInActivity"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSignInActivityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplicationAuthenticationResourceSignInActivity(val.(SignInActivityable))
        }
        return nil
    }
    res["delegatedClientSignInActivity"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSignInActivityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDelegatedClientSignInActivity(val.(SignInActivityable))
        }
        return nil
    }
    res["delegatedResourceSignInActivity"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSignInActivityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDelegatedResourceSignInActivity(val.(SignInActivityable))
        }
        return nil
    }
    res["lastSignInActivity"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSignInActivityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastSignInActivity(val.(SignInActivityable))
        }
        return nil
    }
    return res
}
// GetLastSignInActivity gets the lastSignInActivity property value. The most recent sign-in activity of the application across delegated or app-only flows where the application is used either as a client or resource.
// returns a SignInActivityable when successful
func (m *ServicePrincipalSignInActivity) GetLastSignInActivity()(SignInActivityable) {
    val, err := m.GetBackingStore().Get("lastSignInActivity")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(SignInActivityable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *ServicePrincipalSignInActivity) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("appId", m.GetAppId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("applicationAuthenticationClientSignInActivity", m.GetApplicationAuthenticationClientSignInActivity())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("applicationAuthenticationResourceSignInActivity", m.GetApplicationAuthenticationResourceSignInActivity())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("delegatedClientSignInActivity", m.GetDelegatedClientSignInActivity())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("delegatedResourceSignInActivity", m.GetDelegatedResourceSignInActivity())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("lastSignInActivity", m.GetLastSignInActivity())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAppId sets the appId property value. The globally unique appId (also called client ID on the Microsoft Entra admin center) of the credentialed resource application.
func (m *ServicePrincipalSignInActivity) SetAppId(value *string)() {
    err := m.GetBackingStore().Set("appId", value)
    if err != nil {
        panic(err)
    }
}
// SetApplicationAuthenticationClientSignInActivity sets the applicationAuthenticationClientSignInActivity property value. The sign-in activity of the application in a app-only authentication flow (app-to-app tokens) where the application acts like a client.
func (m *ServicePrincipalSignInActivity) SetApplicationAuthenticationClientSignInActivity(value SignInActivityable)() {
    err := m.GetBackingStore().Set("applicationAuthenticationClientSignInActivity", value)
    if err != nil {
        panic(err)
    }
}
// SetApplicationAuthenticationResourceSignInActivity sets the applicationAuthenticationResourceSignInActivity property value. The sign-in activity of the application in a app-only authentication flow (app-to-app tokens) where the application acts like a resource.
func (m *ServicePrincipalSignInActivity) SetApplicationAuthenticationResourceSignInActivity(value SignInActivityable)() {
    err := m.GetBackingStore().Set("applicationAuthenticationResourceSignInActivity", value)
    if err != nil {
        panic(err)
    }
}
// SetDelegatedClientSignInActivity sets the delegatedClientSignInActivity property value. The sign-in activity of the application in a delegated flow (user sign-in) where the application acts like a client.
func (m *ServicePrincipalSignInActivity) SetDelegatedClientSignInActivity(value SignInActivityable)() {
    err := m.GetBackingStore().Set("delegatedClientSignInActivity", value)
    if err != nil {
        panic(err)
    }
}
// SetDelegatedResourceSignInActivity sets the delegatedResourceSignInActivity property value. The sign-in activity of the application in a delegated flow (user sign-in) where the application acts like a resource.
func (m *ServicePrincipalSignInActivity) SetDelegatedResourceSignInActivity(value SignInActivityable)() {
    err := m.GetBackingStore().Set("delegatedResourceSignInActivity", value)
    if err != nil {
        panic(err)
    }
}
// SetLastSignInActivity sets the lastSignInActivity property value. The most recent sign-in activity of the application across delegated or app-only flows where the application is used either as a client or resource.
func (m *ServicePrincipalSignInActivity) SetLastSignInActivity(value SignInActivityable)() {
    err := m.GetBackingStore().Set("lastSignInActivity", value)
    if err != nil {
        panic(err)
    }
}
type ServicePrincipalSignInActivityable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAppId()(*string)
    GetApplicationAuthenticationClientSignInActivity()(SignInActivityable)
    GetApplicationAuthenticationResourceSignInActivity()(SignInActivityable)
    GetDelegatedClientSignInActivity()(SignInActivityable)
    GetDelegatedResourceSignInActivity()(SignInActivityable)
    GetLastSignInActivity()(SignInActivityable)
    SetAppId(value *string)()
    SetApplicationAuthenticationClientSignInActivity(value SignInActivityable)()
    SetApplicationAuthenticationResourceSignInActivity(value SignInActivityable)()
    SetDelegatedClientSignInActivity(value SignInActivityable)()
    SetDelegatedResourceSignInActivity(value SignInActivityable)()
    SetLastSignInActivity(value SignInActivityable)()
}
