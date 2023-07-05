package networkaccess

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// TenantStatus 
type TenantStatus struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entity
}
// NewTenantStatus instantiates a new TenantStatus and sets the default values.
func NewTenantStatus()(*TenantStatus) {
    m := &TenantStatus{
        Entity: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewEntity(),
    }
    return m
}
// CreateTenantStatusFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTenantStatusFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTenantStatus(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TenantStatus) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["onboardingErrorMessage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOnboardingErrorMessage(val)
        }
        return nil
    }
    res["onboardingStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseOnboardingStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOnboardingStatus(val.(*OnboardingStatus))
        }
        return nil
    }
    return res
}
// GetOnboardingErrorMessage gets the onboardingErrorMessage property value. The onboardingErrorMessage property
func (m *TenantStatus) GetOnboardingErrorMessage()(*string) {
    val, err := m.GetBackingStore().Get("onboardingErrorMessage")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetOnboardingStatus gets the onboardingStatus property value. The onboardingStatus property
func (m *TenantStatus) GetOnboardingStatus()(*OnboardingStatus) {
    val, err := m.GetBackingStore().Get("onboardingStatus")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*OnboardingStatus)
    }
    return nil
}
// Serialize serializes information the current object
func (m *TenantStatus) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("onboardingErrorMessage", m.GetOnboardingErrorMessage())
        if err != nil {
            return err
        }
    }
    if m.GetOnboardingStatus() != nil {
        cast := (*m.GetOnboardingStatus()).String()
        err = writer.WriteStringValue("onboardingStatus", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetOnboardingErrorMessage sets the onboardingErrorMessage property value. The onboardingErrorMessage property
func (m *TenantStatus) SetOnboardingErrorMessage(value *string)() {
    err := m.GetBackingStore().Set("onboardingErrorMessage", value)
    if err != nil {
        panic(err)
    }
}
// SetOnboardingStatus sets the onboardingStatus property value. The onboardingStatus property
func (m *TenantStatus) SetOnboardingStatus(value *OnboardingStatus)() {
    err := m.GetBackingStore().Set("onboardingStatus", value)
    if err != nil {
        panic(err)
    }
}
// TenantStatusable 
type TenantStatusable interface {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetOnboardingErrorMessage()(*string)
    GetOnboardingStatus()(*OnboardingStatus)
    SetOnboardingErrorMessage(value *string)()
    SetOnboardingStatus(value *OnboardingStatus)()
}
