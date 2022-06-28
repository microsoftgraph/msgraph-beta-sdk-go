package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WebApp 
type WebApp struct {
    MobileApp
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The web app URL. This property cannot be PATCHed.
    appUrl *string
    // Whether or not to use managed browser. This property is only applicable for Android and IOS.
    useManagedBrowser *bool
}
// NewWebApp instantiates a new WebApp and sets the default values.
func NewWebApp()(*WebApp) {
    m := &WebApp{
        MobileApp: *NewMobileApp(),
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateWebAppFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWebAppFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWebApp(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WebApp) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAppUrl gets the appUrl property value. The web app URL. This property cannot be PATCHed.
func (m *WebApp) GetAppUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.appUrl
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WebApp) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.MobileApp.GetFieldDeserializers()
    res["appUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppUrl(val)
        }
        return nil
    }
    res["useManagedBrowser"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUseManagedBrowser(val)
        }
        return nil
    }
    return res
}
// GetUseManagedBrowser gets the useManagedBrowser property value. Whether or not to use managed browser. This property is only applicable for Android and IOS.
func (m *WebApp) GetUseManagedBrowser()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.useManagedBrowser
    }
}
// Serialize serializes information the current object
func (m *WebApp) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.MobileApp.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("appUrl", m.GetAppUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("useManagedBrowser", m.GetUseManagedBrowser())
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
func (m *WebApp) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAppUrl sets the appUrl property value. The web app URL. This property cannot be PATCHed.
func (m *WebApp) SetAppUrl(value *string)() {
    if m != nil {
        m.appUrl = value
    }
}
// SetUseManagedBrowser sets the useManagedBrowser property value. Whether or not to use managed browser. This property is only applicable for Android and IOS.
func (m *WebApp) SetUseManagedBrowser(value *bool)() {
    if m != nil {
        m.useManagedBrowser = value
    }
}
