package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// ServerProcessedContent 
type ServerProcessedContent struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewServerProcessedContent instantiates a new serverProcessedContent and sets the default values.
func NewServerProcessedContent()(*ServerProcessedContent) {
    m := &ServerProcessedContent{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateServerProcessedContentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateServerProcessedContentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewServerProcessedContent(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ServerProcessedContent) GetAdditionalData()(map[string]any) {
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
// GetBackingStore gets the backingStore property value. Stores model information.
func (m *ServerProcessedContent) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetComponentDependencies gets the componentDependencies property value. A key-value map where keys are string identifiers and values are component ids. SharePoint servers might decide to use this hint to preload the script for corresponding components for performance boost.
func (m *ServerProcessedContent) GetComponentDependencies()([]MetaDataKeyStringPairable) {
    val, err := m.GetBackingStore().Get("componentDependencies")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]MetaDataKeyStringPairable)
    }
    return nil
}
// GetCustomMetadata gets the customMetadata property value. A key-value map where keys are string identifier and values are object of custom key-value pair.
func (m *ServerProcessedContent) GetCustomMetadata()([]MetaDataKeyValuePairable) {
    val, err := m.GetBackingStore().Get("customMetadata")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]MetaDataKeyValuePairable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ServerProcessedContent) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["componentDependencies"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMetaDataKeyStringPairFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MetaDataKeyStringPairable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(MetaDataKeyStringPairable)
                }
            }
            m.SetComponentDependencies(res)
        }
        return nil
    }
    res["customMetadata"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMetaDataKeyValuePairFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MetaDataKeyValuePairable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(MetaDataKeyValuePairable)
                }
            }
            m.SetCustomMetadata(res)
        }
        return nil
    }
    res["htmlStrings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMetaDataKeyStringPairFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MetaDataKeyStringPairable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(MetaDataKeyStringPairable)
                }
            }
            m.SetHtmlStrings(res)
        }
        return nil
    }
    res["imageSources"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMetaDataKeyStringPairFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MetaDataKeyStringPairable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(MetaDataKeyStringPairable)
                }
            }
            m.SetImageSources(res)
        }
        return nil
    }
    res["links"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMetaDataKeyStringPairFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MetaDataKeyStringPairable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(MetaDataKeyStringPairable)
                }
            }
            m.SetLinks(res)
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    res["searchablePlainTexts"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMetaDataKeyStringPairFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MetaDataKeyStringPairable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(MetaDataKeyStringPairable)
                }
            }
            m.SetSearchablePlainTexts(res)
        }
        return nil
    }
    return res
}
// GetHtmlStrings gets the htmlStrings property value. A key-value map where keys are string identifiers and values are rich text with HTML format. SharePoint servers treat the values as HTML content and run services like safety checks, search index and link fixup on them.
func (m *ServerProcessedContent) GetHtmlStrings()([]MetaDataKeyStringPairable) {
    val, err := m.GetBackingStore().Get("htmlStrings")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]MetaDataKeyStringPairable)
    }
    return nil
}
// GetImageSources gets the imageSources property value. A key-value map where keys are string identifiers and values are image sources. SharePoint servers treat the values as image sources and run services like search index and link fixup on them.
func (m *ServerProcessedContent) GetImageSources()([]MetaDataKeyStringPairable) {
    val, err := m.GetBackingStore().Get("imageSources")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]MetaDataKeyStringPairable)
    }
    return nil
}
// GetLinks gets the links property value. A key-value map where keys are string identifiers and values are links. SharePoint servers treat the values as links and run services like link fixup on them.
func (m *ServerProcessedContent) GetLinks()([]MetaDataKeyStringPairable) {
    val, err := m.GetBackingStore().Get("links")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]MetaDataKeyStringPairable)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *ServerProcessedContent) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSearchablePlainTexts gets the searchablePlainTexts property value. A key-value map where keys are string identifiers and values are strings that should be search indexed.
func (m *ServerProcessedContent) GetSearchablePlainTexts()([]MetaDataKeyStringPairable) {
    val, err := m.GetBackingStore().Get("searchablePlainTexts")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]MetaDataKeyStringPairable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *ServerProcessedContent) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetComponentDependencies() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetComponentDependencies()))
        for i, v := range m.GetComponentDependencies() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("componentDependencies", cast)
        if err != nil {
            return err
        }
    }
    if m.GetCustomMetadata() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCustomMetadata()))
        for i, v := range m.GetCustomMetadata() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("customMetadata", cast)
        if err != nil {
            return err
        }
    }
    if m.GetHtmlStrings() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetHtmlStrings()))
        for i, v := range m.GetHtmlStrings() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("htmlStrings", cast)
        if err != nil {
            return err
        }
    }
    if m.GetImageSources() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetImageSources()))
        for i, v := range m.GetImageSources() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("imageSources", cast)
        if err != nil {
            return err
        }
    }
    if m.GetLinks() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetLinks()))
        for i, v := range m.GetLinks() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("links", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    if m.GetSearchablePlainTexts() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSearchablePlainTexts()))
        for i, v := range m.GetSearchablePlainTexts() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("searchablePlainTexts", cast)
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ServerProcessedContent) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the backingStore property value. Stores model information.
func (m *ServerProcessedContent) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetComponentDependencies sets the componentDependencies property value. A key-value map where keys are string identifiers and values are component ids. SharePoint servers might decide to use this hint to preload the script for corresponding components for performance boost.
func (m *ServerProcessedContent) SetComponentDependencies(value []MetaDataKeyStringPairable)() {
    err := m.GetBackingStore().Set("componentDependencies", value)
    if err != nil {
        panic(err)
    }
}
// SetCustomMetadata sets the customMetadata property value. A key-value map where keys are string identifier and values are object of custom key-value pair.
func (m *ServerProcessedContent) SetCustomMetadata(value []MetaDataKeyValuePairable)() {
    err := m.GetBackingStore().Set("customMetadata", value)
    if err != nil {
        panic(err)
    }
}
// SetHtmlStrings sets the htmlStrings property value. A key-value map where keys are string identifiers and values are rich text with HTML format. SharePoint servers treat the values as HTML content and run services like safety checks, search index and link fixup on them.
func (m *ServerProcessedContent) SetHtmlStrings(value []MetaDataKeyStringPairable)() {
    err := m.GetBackingStore().Set("htmlStrings", value)
    if err != nil {
        panic(err)
    }
}
// SetImageSources sets the imageSources property value. A key-value map where keys are string identifiers and values are image sources. SharePoint servers treat the values as image sources and run services like search index and link fixup on them.
func (m *ServerProcessedContent) SetImageSources(value []MetaDataKeyStringPairable)() {
    err := m.GetBackingStore().Set("imageSources", value)
    if err != nil {
        panic(err)
    }
}
// SetLinks sets the links property value. A key-value map where keys are string identifiers and values are links. SharePoint servers treat the values as links and run services like link fixup on them.
func (m *ServerProcessedContent) SetLinks(value []MetaDataKeyStringPairable)() {
    err := m.GetBackingStore().Set("links", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *ServerProcessedContent) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetSearchablePlainTexts sets the searchablePlainTexts property value. A key-value map where keys are string identifiers and values are strings that should be search indexed.
func (m *ServerProcessedContent) SetSearchablePlainTexts(value []MetaDataKeyStringPairable)() {
    err := m.GetBackingStore().Set("searchablePlainTexts", value)
    if err != nil {
        panic(err)
    }
}
// ServerProcessedContentable 
type ServerProcessedContentable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetComponentDependencies()([]MetaDataKeyStringPairable)
    GetCustomMetadata()([]MetaDataKeyValuePairable)
    GetHtmlStrings()([]MetaDataKeyStringPairable)
    GetImageSources()([]MetaDataKeyStringPairable)
    GetLinks()([]MetaDataKeyStringPairable)
    GetOdataType()(*string)
    GetSearchablePlainTexts()([]MetaDataKeyStringPairable)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetComponentDependencies(value []MetaDataKeyStringPairable)()
    SetCustomMetadata(value []MetaDataKeyValuePairable)()
    SetHtmlStrings(value []MetaDataKeyStringPairable)()
    SetImageSources(value []MetaDataKeyStringPairable)()
    SetLinks(value []MetaDataKeyStringPairable)()
    SetOdataType(value *string)()
    SetSearchablePlainTexts(value []MetaDataKeyStringPairable)()
}
