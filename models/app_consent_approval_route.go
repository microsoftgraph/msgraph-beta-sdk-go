package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AppConsentApprovalRoute 
type AppConsentApprovalRoute struct {
    Entity
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // A collection of userConsentRequest objects for a specific application.
    appConsentRequests []AppConsentRequestable
}
// NewAppConsentApprovalRoute instantiates a new AppConsentApprovalRoute and sets the default values.
func NewAppConsentApprovalRoute()(*AppConsentApprovalRoute) {
    m := &AppConsentApprovalRoute{
        Entity: *NewEntity(),
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateAppConsentApprovalRouteFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAppConsentApprovalRouteFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAppConsentApprovalRoute(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AppConsentApprovalRoute) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAppConsentRequests gets the appConsentRequests property value. A collection of userConsentRequest objects for a specific application.
func (m *AppConsentApprovalRoute) GetAppConsentRequests()([]AppConsentRequestable) {
    if m == nil {
        return nil
    } else {
        return m.appConsentRequests
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AppConsentApprovalRoute) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["appConsentRequests"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAppConsentRequestFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AppConsentRequestable, len(val))
            for i, v := range val {
                res[i] = v.(AppConsentRequestable)
            }
            m.SetAppConsentRequests(res)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *AppConsentApprovalRoute) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAppConsentRequests() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAppConsentRequests()))
        for i, v := range m.GetAppConsentRequests() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("appConsentRequests", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AppConsentApprovalRoute) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAppConsentRequests sets the appConsentRequests property value. A collection of userConsentRequest objects for a specific application.
func (m *AppConsentApprovalRoute) SetAppConsentRequests(value []AppConsentRequestable)() {
    if m != nil {
        m.appConsentRequests = value
    }
}
