package directory

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

type TemplatesDeviceTemplatesItemCreateDeviceFromTemplatePostRequestBody struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewTemplatesDeviceTemplatesItemCreateDeviceFromTemplatePostRequestBody instantiates a new TemplatesDeviceTemplatesItemCreateDeviceFromTemplatePostRequestBody and sets the default values.
func NewTemplatesDeviceTemplatesItemCreateDeviceFromTemplatePostRequestBody()(*TemplatesDeviceTemplatesItemCreateDeviceFromTemplatePostRequestBody) {
    m := &TemplatesDeviceTemplatesItemCreateDeviceFromTemplatePostRequestBody{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateTemplatesDeviceTemplatesItemCreateDeviceFromTemplatePostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTemplatesDeviceTemplatesItemCreateDeviceFromTemplatePostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTemplatesDeviceTemplatesItemCreateDeviceFromTemplatePostRequestBody(), nil
}
// GetAccountEnabled gets the accountEnabled property value. The accountEnabled property
// returns a *bool when successful
func (m *TemplatesDeviceTemplatesItemCreateDeviceFromTemplatePostRequestBody) GetAccountEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("accountEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *TemplatesDeviceTemplatesItemCreateDeviceFromTemplatePostRequestBody) GetAdditionalData()(map[string]any) {
    val , err :=  m.backingStore.Get("additionalData")
    if err != nil {
        panic(err)
    }
    if val == nil {
        var value = make(map[string]any);
        m.SetAdditionalData(value);
    }
    return val.(map[string]any)
}
// GetAlternativeNames gets the alternativeNames property value. The alternativeNames property
// returns a []string when successful
func (m *TemplatesDeviceTemplatesItemCreateDeviceFromTemplatePostRequestBody) GetAlternativeNames()([]string) {
    val, err := m.GetBackingStore().Get("alternativeNames")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetBackingStore gets the BackingStore property value. Stores model information.
// returns a BackingStore when successful
func (m *TemplatesDeviceTemplatesItemCreateDeviceFromTemplatePostRequestBody) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetExternalDeviceId gets the externalDeviceId property value. The externalDeviceId property
// returns a *string when successful
func (m *TemplatesDeviceTemplatesItemCreateDeviceFromTemplatePostRequestBody) GetExternalDeviceId()(*string) {
    val, err := m.GetBackingStore().Get("externalDeviceId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetExternalSourceName gets the externalSourceName property value. The externalSourceName property
// returns a *string when successful
func (m *TemplatesDeviceTemplatesItemCreateDeviceFromTemplatePostRequestBody) GetExternalSourceName()(*string) {
    val, err := m.GetBackingStore().Get("externalSourceName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *TemplatesDeviceTemplatesItemCreateDeviceFromTemplatePostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["accountEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccountEnabled(val)
        }
        return nil
    }
    res["alternativeNames"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetAlternativeNames(res)
        }
        return nil
    }
    res["externalDeviceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExternalDeviceId(val)
        }
        return nil
    }
    res["externalSourceName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExternalSourceName(val)
        }
        return nil
    }
    res["keyCredential"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateKeyCredentialFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKeyCredential(val.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.KeyCredentialable))
        }
        return nil
    }
    res["operatingSystemVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOperatingSystemVersion(val)
        }
        return nil
    }
    return res
}
// GetKeyCredential gets the keyCredential property value. The keyCredential property
// returns a KeyCredentialable when successful
func (m *TemplatesDeviceTemplatesItemCreateDeviceFromTemplatePostRequestBody) GetKeyCredential()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.KeyCredentialable) {
    val, err := m.GetBackingStore().Get("keyCredential")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.KeyCredentialable)
    }
    return nil
}
// GetOperatingSystemVersion gets the operatingSystemVersion property value. The operatingSystemVersion property
// returns a *string when successful
func (m *TemplatesDeviceTemplatesItemCreateDeviceFromTemplatePostRequestBody) GetOperatingSystemVersion()(*string) {
    val, err := m.GetBackingStore().Get("operatingSystemVersion")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *TemplatesDeviceTemplatesItemCreateDeviceFromTemplatePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("accountEnabled", m.GetAccountEnabled())
        if err != nil {
            return err
        }
    }
    if m.GetAlternativeNames() != nil {
        err := writer.WriteCollectionOfStringValues("alternativeNames", m.GetAlternativeNames())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("externalDeviceId", m.GetExternalDeviceId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("externalSourceName", m.GetExternalSourceName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("keyCredential", m.GetKeyCredential())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("operatingSystemVersion", m.GetOperatingSystemVersion())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAccountEnabled sets the accountEnabled property value. The accountEnabled property
func (m *TemplatesDeviceTemplatesItemCreateDeviceFromTemplatePostRequestBody) SetAccountEnabled(value *bool)() {
    err := m.GetBackingStore().Set("accountEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TemplatesDeviceTemplatesItemCreateDeviceFromTemplatePostRequestBody) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetAlternativeNames sets the alternativeNames property value. The alternativeNames property
func (m *TemplatesDeviceTemplatesItemCreateDeviceFromTemplatePostRequestBody) SetAlternativeNames(value []string)() {
    err := m.GetBackingStore().Set("alternativeNames", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *TemplatesDeviceTemplatesItemCreateDeviceFromTemplatePostRequestBody) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetExternalDeviceId sets the externalDeviceId property value. The externalDeviceId property
func (m *TemplatesDeviceTemplatesItemCreateDeviceFromTemplatePostRequestBody) SetExternalDeviceId(value *string)() {
    err := m.GetBackingStore().Set("externalDeviceId", value)
    if err != nil {
        panic(err)
    }
}
// SetExternalSourceName sets the externalSourceName property value. The externalSourceName property
func (m *TemplatesDeviceTemplatesItemCreateDeviceFromTemplatePostRequestBody) SetExternalSourceName(value *string)() {
    err := m.GetBackingStore().Set("externalSourceName", value)
    if err != nil {
        panic(err)
    }
}
// SetKeyCredential sets the keyCredential property value. The keyCredential property
func (m *TemplatesDeviceTemplatesItemCreateDeviceFromTemplatePostRequestBody) SetKeyCredential(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.KeyCredentialable)() {
    err := m.GetBackingStore().Set("keyCredential", value)
    if err != nil {
        panic(err)
    }
}
// SetOperatingSystemVersion sets the operatingSystemVersion property value. The operatingSystemVersion property
func (m *TemplatesDeviceTemplatesItemCreateDeviceFromTemplatePostRequestBody) SetOperatingSystemVersion(value *string)() {
    err := m.GetBackingStore().Set("operatingSystemVersion", value)
    if err != nil {
        panic(err)
    }
}
type TemplatesDeviceTemplatesItemCreateDeviceFromTemplatePostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccountEnabled()(*bool)
    GetAlternativeNames()([]string)
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetExternalDeviceId()(*string)
    GetExternalSourceName()(*string)
    GetKeyCredential()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.KeyCredentialable)
    GetOperatingSystemVersion()(*string)
    SetAccountEnabled(value *bool)()
    SetAlternativeNames(value []string)()
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetExternalDeviceId(value *string)()
    SetExternalSourceName(value *string)()
    SetKeyCredential(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.KeyCredentialable)()
    SetOperatingSystemVersion(value *string)()
}
