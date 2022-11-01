package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ProfileCardProperty provides operations to manage the collection of accessReview entities.
type ProfileCardProperty struct {
    Entity
    // Allows an administrator to set a custom display label for the directory property and localize it for the users in their tenant.
    annotations []ProfileCardAnnotationable
    // Identifies a profileCardProperty resource in Get, Update, or Delete operations. Allows an administrator to surface hidden Azure Active Directory (Azure AD) properties on the Microsoft 365 profile card within their tenant. When present, the Azure AD field referenced in this field will be visible to all users in your tenant on the contact pane of the profile card. Allowed values for this field are: UserPrincipalName, Fax, StreetAddress, PostalCode, StateOrProvince, Alias, CustomAttribute1,  CustomAttribute2, CustomAttribute3, CustomAttribute4, CustomAttribute5, CustomAttribute6, CustomAttribute7, CustomAttribute8, CustomAttribute9, CustomAttribute10, CustomAttribute11, CustomAttribute12, CustomAttribute13, CustomAttribute14, CustomAttribute15.
    directoryPropertyName *string
}
// NewProfileCardProperty instantiates a new profileCardProperty and sets the default values.
func NewProfileCardProperty()(*ProfileCardProperty) {
    m := &ProfileCardProperty{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.profileCardProperty";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateProfileCardPropertyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateProfileCardPropertyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewProfileCardProperty(), nil
}
// GetAnnotations gets the annotations property value. Allows an administrator to set a custom display label for the directory property and localize it for the users in their tenant.
func (m *ProfileCardProperty) GetAnnotations()([]ProfileCardAnnotationable) {
    return m.annotations
}
// GetDirectoryPropertyName gets the directoryPropertyName property value. Identifies a profileCardProperty resource in Get, Update, or Delete operations. Allows an administrator to surface hidden Azure Active Directory (Azure AD) properties on the Microsoft 365 profile card within their tenant. When present, the Azure AD field referenced in this field will be visible to all users in your tenant on the contact pane of the profile card. Allowed values for this field are: UserPrincipalName, Fax, StreetAddress, PostalCode, StateOrProvince, Alias, CustomAttribute1,  CustomAttribute2, CustomAttribute3, CustomAttribute4, CustomAttribute5, CustomAttribute6, CustomAttribute7, CustomAttribute8, CustomAttribute9, CustomAttribute10, CustomAttribute11, CustomAttribute12, CustomAttribute13, CustomAttribute14, CustomAttribute15.
func (m *ProfileCardProperty) GetDirectoryPropertyName()(*string) {
    return m.directoryPropertyName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ProfileCardProperty) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["annotations"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateProfileCardAnnotationFromDiscriminatorValue , m.SetAnnotations)
    res["directoryPropertyName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDirectoryPropertyName)
    return res
}
// Serialize serializes information the current object
func (m *ProfileCardProperty) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAnnotations() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetAnnotations())
        err = writer.WriteCollectionOfObjectValues("annotations", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("directoryPropertyName", m.GetDirectoryPropertyName())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAnnotations sets the annotations property value. Allows an administrator to set a custom display label for the directory property and localize it for the users in their tenant.
func (m *ProfileCardProperty) SetAnnotations(value []ProfileCardAnnotationable)() {
    m.annotations = value
}
// SetDirectoryPropertyName sets the directoryPropertyName property value. Identifies a profileCardProperty resource in Get, Update, or Delete operations. Allows an administrator to surface hidden Azure Active Directory (Azure AD) properties on the Microsoft 365 profile card within their tenant. When present, the Azure AD field referenced in this field will be visible to all users in your tenant on the contact pane of the profile card. Allowed values for this field are: UserPrincipalName, Fax, StreetAddress, PostalCode, StateOrProvince, Alias, CustomAttribute1,  CustomAttribute2, CustomAttribute3, CustomAttribute4, CustomAttribute5, CustomAttribute6, CustomAttribute7, CustomAttribute8, CustomAttribute9, CustomAttribute10, CustomAttribute11, CustomAttribute12, CustomAttribute13, CustomAttribute14, CustomAttribute15.
func (m *ProfileCardProperty) SetDirectoryPropertyName(value *string)() {
    m.directoryPropertyName = value
}
