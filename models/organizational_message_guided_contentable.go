package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OrganizationalMessageGuidedContentable 
type OrganizationalMessageGuidedContentable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetLogo()(OrganizationalMessageLogoGuideable)
    GetPlacementDetails()([]OrganizationalMessagePlacementDetailable)
    GetScenario()(*OrganizationalMessageScenario)
    GetSurface()(*OrganizationalMessageSurface)
    GetTheme()(*OrganizationalMessageTheme)
    SetLogo(value OrganizationalMessageLogoGuideable)()
    SetPlacementDetails(value []OrganizationalMessagePlacementDetailable)()
    SetScenario(value *OrganizationalMessageScenario)()
    SetSurface(value *OrganizationalMessageSurface)()
    SetTheme(value *OrganizationalMessageTheme)()
}
