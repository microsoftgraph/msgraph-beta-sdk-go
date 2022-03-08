package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// WindowsAssignedAccessProfileable 
type WindowsAssignedAccessProfileable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAppUserModelIds()([]string)
    GetDesktopAppPaths()([]string)
    GetProfileName()(*string)
    GetShowTaskBar()(*bool)
    GetStartMenuLayoutXml()([]byte)
    GetUserAccounts()([]string)
    SetAppUserModelIds(value []string)()
    SetDesktopAppPaths(value []string)()
    SetProfileName(value *string)()
    SetShowTaskBar(value *bool)()
    SetStartMenuLayoutXml(value []byte)()
    SetUserAccounts(value []string)()
}
