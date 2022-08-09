package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OrganizationalMessageDetailable 
type OrganizationalMessageDetailable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetContent()(OrganizationalMessageContentable)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetEndDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetFrequency()(*OrganizationalMessageFrequency)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetScenario()(*OrganizationalMessageScenario)
    GetStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetStatus()(*OrganizationalMessageStatus)
    GetSurface()(*OrganizationalMessageSurface)
    GetTargeting()(OrganizationalMessageTargetingable)
    GetTheme()(*OrganizationalMessageTheme)
    GetUserEngagementStatistics()(OrganizationalMessageInsightsable)
    GetVariant()(*string)
    SetContent(value OrganizationalMessageContentable)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetEndDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetFrequency(value *OrganizationalMessageFrequency)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetScenario(value *OrganizationalMessageScenario)()
    SetStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetStatus(value *OrganizationalMessageStatus)()
    SetSurface(value *OrganizationalMessageSurface)()
    SetTargeting(value OrganizationalMessageTargetingable)()
    SetTheme(value *OrganizationalMessageTheme)()
    SetUserEngagementStatistics(value OrganizationalMessageInsightsable)()
    SetVariant(value *string)()
}
