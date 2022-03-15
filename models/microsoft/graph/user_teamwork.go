package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// UserTeamwork provides operations to manage the deviceManagement singleton.
type UserTeamwork struct {
    Entity
    // 
    associatedTeams []AssociatedTeamInfoable;
    // The apps installed in the personal scope of this user.
    installedApps []UserScopeTeamsAppInstallationable;
}
// NewUserTeamwork instantiates a new userTeamwork and sets the default values.
func NewUserTeamwork()(*UserTeamwork) {
    m := &UserTeamwork{
        Entity: *NewEntity(),
    }
    return m
}
// CreateUserTeamworkFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUserTeamworkFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewUserTeamwork(), nil
}
// GetAssociatedTeams gets the associatedTeams property value. 
func (m *UserTeamwork) GetAssociatedTeams()([]AssociatedTeamInfoable) {
    if m == nil {
        return nil
    } else {
        return m.associatedTeams
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserTeamwork) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["associatedTeams"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAssociatedTeamInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AssociatedTeamInfoable, len(val))
            for i, v := range val {
                res[i] = v.(AssociatedTeamInfoable)
            }
            m.SetAssociatedTeams(res)
        }
        return nil
    }
    res["installedApps"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserScopeTeamsAppInstallationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserScopeTeamsAppInstallationable, len(val))
            for i, v := range val {
                res[i] = v.(UserScopeTeamsAppInstallationable)
            }
            m.SetInstalledApps(res)
        }
        return nil
    }
    return res
}
// GetInstalledApps gets the installedApps property value. The apps installed in the personal scope of this user.
func (m *UserTeamwork) GetInstalledApps()([]UserScopeTeamsAppInstallationable) {
    if m == nil {
        return nil
    } else {
        return m.installedApps
    }
}
func (m *UserTeamwork) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *UserTeamwork) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAssociatedTeams() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAssociatedTeams()))
        for i, v := range m.GetAssociatedTeams() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("associatedTeams", cast)
        if err != nil {
            return err
        }
    }
    if m.GetInstalledApps() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetInstalledApps()))
        for i, v := range m.GetInstalledApps() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("installedApps", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAssociatedTeams sets the associatedTeams property value. 
func (m *UserTeamwork) SetAssociatedTeams(value []AssociatedTeamInfoable)() {
    if m != nil {
        m.associatedTeams = value
    }
}
// SetInstalledApps sets the installedApps property value. The apps installed in the personal scope of this user.
func (m *UserTeamwork) SetInstalledApps(value []UserScopeTeamsAppInstallationable)() {
    if m != nil {
        m.installedApps = value
    }
}
