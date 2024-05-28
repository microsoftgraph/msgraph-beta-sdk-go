package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

type IntentsItemMigratetotemplateMigrateToTemplatePostRequestBody struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewIntentsItemMigratetotemplateMigrateToTemplatePostRequestBody instantiates a new IntentsItemMigratetotemplateMigrateToTemplatePostRequestBody and sets the default values.
func NewIntentsItemMigratetotemplateMigrateToTemplatePostRequestBody()(*IntentsItemMigratetotemplateMigrateToTemplatePostRequestBody) {
    m := &IntentsItemMigratetotemplateMigrateToTemplatePostRequestBody{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateIntentsItemMigratetotemplateMigrateToTemplatePostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateIntentsItemMigratetotemplateMigrateToTemplatePostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIntentsItemMigratetotemplateMigrateToTemplatePostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *IntentsItemMigratetotemplateMigrateToTemplatePostRequestBody) GetAdditionalData()(map[string]any) {
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
// GetBackingStore gets the BackingStore property value. Stores model information.
// returns a BackingStore when successful
func (m *IntentsItemMigratetotemplateMigrateToTemplatePostRequestBody) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *IntentsItemMigratetotemplateMigrateToTemplatePostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["newTemplateId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNewTemplateId(val)
        }
        return nil
    }
    res["preserveCustomValues"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPreserveCustomValues(val)
        }
        return nil
    }
    return res
}
// GetNewTemplateId gets the newTemplateId property value. The newTemplateId property
// returns a *string when successful
func (m *IntentsItemMigratetotemplateMigrateToTemplatePostRequestBody) GetNewTemplateId()(*string) {
    val, err := m.GetBackingStore().Get("newTemplateId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPreserveCustomValues gets the preserveCustomValues property value. The preserveCustomValues property
// returns a *bool when successful
func (m *IntentsItemMigratetotemplateMigrateToTemplatePostRequestBody) GetPreserveCustomValues()(*bool) {
    val, err := m.GetBackingStore().Get("preserveCustomValues")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// Serialize serializes information the current object
func (m *IntentsItemMigratetotemplateMigrateToTemplatePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("newTemplateId", m.GetNewTemplateId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("preserveCustomValues", m.GetPreserveCustomValues())
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
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *IntentsItemMigratetotemplateMigrateToTemplatePostRequestBody) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *IntentsItemMigratetotemplateMigrateToTemplatePostRequestBody) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetNewTemplateId sets the newTemplateId property value. The newTemplateId property
func (m *IntentsItemMigratetotemplateMigrateToTemplatePostRequestBody) SetNewTemplateId(value *string)() {
    err := m.GetBackingStore().Set("newTemplateId", value)
    if err != nil {
        panic(err)
    }
}
// SetPreserveCustomValues sets the preserveCustomValues property value. The preserveCustomValues property
func (m *IntentsItemMigratetotemplateMigrateToTemplatePostRequestBody) SetPreserveCustomValues(value *bool)() {
    err := m.GetBackingStore().Set("preserveCustomValues", value)
    if err != nil {
        panic(err)
    }
}
type IntentsItemMigratetotemplateMigrateToTemplatePostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetNewTemplateId()(*string)
    GetPreserveCustomValues()(*bool)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetNewTemplateId(value *string)()
    SetPreserveCustomValues(value *bool)()
}
