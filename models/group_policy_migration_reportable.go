package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2bacd9b8d8db2e77ee2b5c5ccb19d679c36f920b8fee9d786a0adafff458afcd "github.com/google/UUID"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GroupPolicyMigrationReportable 
type GroupPolicyMigrationReportable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDisplayName()(*string)
    GetGroupPolicyCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetGroupPolicyLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetGroupPolicyObjectId()(*UUID)
    GetGroupPolicySettingMappings()([]GroupPolicySettingMappingable)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetMigrationReadiness()(*GroupPolicyMigrationReadiness)
    GetOuDistinguishedName()(*string)
    GetRoleScopeTagIds()([]string)
    GetSupportedSettingsCount()(*int32)
    GetSupportedSettingsPercent()(*int32)
    GetTargetedInActiveDirectory()(*bool)
    GetTotalSettingsCount()(*int32)
    GetUnsupportedGroupPolicyExtensions()([]UnsupportedGroupPolicyExtensionable)
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDisplayName(value *string)()
    SetGroupPolicyCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetGroupPolicyLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetGroupPolicyObjectId(value *UUID)()
    SetGroupPolicySettingMappings(value []GroupPolicySettingMappingable)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetMigrationReadiness(value *GroupPolicyMigrationReadiness)()
    SetOuDistinguishedName(value *string)()
    SetRoleScopeTagIds(value []string)()
    SetSupportedSettingsCount(value *int32)()
    SetSupportedSettingsPercent(value *int32)()
    SetTargetedInActiveDirectory(value *bool)()
    SetTotalSettingsCount(value *int32)()
    SetUnsupportedGroupPolicyExtensions(value []UnsupportedGroupPolicyExtensionable)()
}
