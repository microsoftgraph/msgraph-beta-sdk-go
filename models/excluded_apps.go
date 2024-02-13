package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// ExcludedApps contains properties for Excluded Office365 Apps.
type ExcludedApps struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewExcludedApps instantiates a new ExcludedApps and sets the default values.
func NewExcludedApps()(*ExcludedApps) {
    m := &ExcludedApps{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateExcludedAppsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateExcludedAppsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewExcludedApps(), nil
}
// GetAccess gets the access property value. The value for if MS Office Access should be excluded or not.
// returns a *bool when successful
func (m *ExcludedApps) GetAccess()(*bool) {
    val, err := m.GetBackingStore().Get("access")
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
func (m *ExcludedApps) GetAdditionalData()(map[string]any) {
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
func (m *ExcludedApps) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetBing gets the bing property value. The value for if Microsoft Search as default should be excluded or not.
// returns a *bool when successful
func (m *ExcludedApps) GetBing()(*bool) {
    val, err := m.GetBackingStore().Get("bing")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetExcel gets the excel property value. The value for if MS Office Excel should be excluded or not.
// returns a *bool when successful
func (m *ExcludedApps) GetExcel()(*bool) {
    val, err := m.GetBackingStore().Get("excel")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ExcludedApps) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["access"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccess(val)
        }
        return nil
    }
    res["bing"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBing(val)
        }
        return nil
    }
    res["excel"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExcel(val)
        }
        return nil
    }
    res["groove"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGroove(val)
        }
        return nil
    }
    res["infoPath"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInfoPath(val)
        }
        return nil
    }
    res["lync"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLync(val)
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
    res["oneDrive"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOneDrive(val)
        }
        return nil
    }
    res["oneNote"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOneNote(val)
        }
        return nil
    }
    res["outlook"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOutlook(val)
        }
        return nil
    }
    res["powerPoint"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPowerPoint(val)
        }
        return nil
    }
    res["publisher"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPublisher(val)
        }
        return nil
    }
    res["sharePointDesigner"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSharePointDesigner(val)
        }
        return nil
    }
    res["teams"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTeams(val)
        }
        return nil
    }
    res["visio"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVisio(val)
        }
        return nil
    }
    res["word"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWord(val)
        }
        return nil
    }
    return res
}
// GetGroove gets the groove property value. The value for if MS Office OneDrive for Business - Groove should be excluded or not.
// returns a *bool when successful
func (m *ExcludedApps) GetGroove()(*bool) {
    val, err := m.GetBackingStore().Get("groove")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetInfoPath gets the infoPath property value. The value for if MS Office InfoPath should be excluded or not.
// returns a *bool when successful
func (m *ExcludedApps) GetInfoPath()(*bool) {
    val, err := m.GetBackingStore().Get("infoPath")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetLync gets the lync property value. The value for if MS Office Skype for Business - Lync should be excluded or not.
// returns a *bool when successful
func (m *ExcludedApps) GetLync()(*bool) {
    val, err := m.GetBackingStore().Get("lync")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *ExcludedApps) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetOneDrive gets the oneDrive property value. The value for if MS Office OneDrive should be excluded or not.
// returns a *bool when successful
func (m *ExcludedApps) GetOneDrive()(*bool) {
    val, err := m.GetBackingStore().Get("oneDrive")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetOneNote gets the oneNote property value. The value for if MS Office OneNote should be excluded or not.
// returns a *bool when successful
func (m *ExcludedApps) GetOneNote()(*bool) {
    val, err := m.GetBackingStore().Get("oneNote")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetOutlook gets the outlook property value. The value for if MS Office Outlook should be excluded or not.
// returns a *bool when successful
func (m *ExcludedApps) GetOutlook()(*bool) {
    val, err := m.GetBackingStore().Get("outlook")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetPowerPoint gets the powerPoint property value. The value for if MS Office PowerPoint should be excluded or not.
// returns a *bool when successful
func (m *ExcludedApps) GetPowerPoint()(*bool) {
    val, err := m.GetBackingStore().Get("powerPoint")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetPublisher gets the publisher property value. The value for if MS Office Publisher should be excluded or not.
// returns a *bool when successful
func (m *ExcludedApps) GetPublisher()(*bool) {
    val, err := m.GetBackingStore().Get("publisher")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetSharePointDesigner gets the sharePointDesigner property value. The value for if MS Office SharePointDesigner should be excluded or not.
// returns a *bool when successful
func (m *ExcludedApps) GetSharePointDesigner()(*bool) {
    val, err := m.GetBackingStore().Get("sharePointDesigner")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetTeams gets the teams property value. The value for if MS Office Teams should be excluded or not.
// returns a *bool when successful
func (m *ExcludedApps) GetTeams()(*bool) {
    val, err := m.GetBackingStore().Get("teams")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetVisio gets the visio property value. The value for if MS Office Visio should be excluded or not.
// returns a *bool when successful
func (m *ExcludedApps) GetVisio()(*bool) {
    val, err := m.GetBackingStore().Get("visio")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetWord gets the word property value. The value for if MS Office Word should be excluded or not.
// returns a *bool when successful
func (m *ExcludedApps) GetWord()(*bool) {
    val, err := m.GetBackingStore().Get("word")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// Serialize serializes information the current object
func (m *ExcludedApps) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("access", m.GetAccess())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("bing", m.GetBing())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("excel", m.GetExcel())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("groove", m.GetGroove())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("infoPath", m.GetInfoPath())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("lync", m.GetLync())
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
    {
        err := writer.WriteBoolValue("oneDrive", m.GetOneDrive())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("oneNote", m.GetOneNote())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("outlook", m.GetOutlook())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("powerPoint", m.GetPowerPoint())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("publisher", m.GetPublisher())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("sharePointDesigner", m.GetSharePointDesigner())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("teams", m.GetTeams())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("visio", m.GetVisio())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("word", m.GetWord())
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
// SetAccess sets the access property value. The value for if MS Office Access should be excluded or not.
func (m *ExcludedApps) SetAccess(value *bool)() {
    err := m.GetBackingStore().Set("access", value)
    if err != nil {
        panic(err)
    }
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ExcludedApps) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *ExcludedApps) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetBing sets the bing property value. The value for if Microsoft Search as default should be excluded or not.
func (m *ExcludedApps) SetBing(value *bool)() {
    err := m.GetBackingStore().Set("bing", value)
    if err != nil {
        panic(err)
    }
}
// SetExcel sets the excel property value. The value for if MS Office Excel should be excluded or not.
func (m *ExcludedApps) SetExcel(value *bool)() {
    err := m.GetBackingStore().Set("excel", value)
    if err != nil {
        panic(err)
    }
}
// SetGroove sets the groove property value. The value for if MS Office OneDrive for Business - Groove should be excluded or not.
func (m *ExcludedApps) SetGroove(value *bool)() {
    err := m.GetBackingStore().Set("groove", value)
    if err != nil {
        panic(err)
    }
}
// SetInfoPath sets the infoPath property value. The value for if MS Office InfoPath should be excluded or not.
func (m *ExcludedApps) SetInfoPath(value *bool)() {
    err := m.GetBackingStore().Set("infoPath", value)
    if err != nil {
        panic(err)
    }
}
// SetLync sets the lync property value. The value for if MS Office Skype for Business - Lync should be excluded or not.
func (m *ExcludedApps) SetLync(value *bool)() {
    err := m.GetBackingStore().Set("lync", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *ExcludedApps) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetOneDrive sets the oneDrive property value. The value for if MS Office OneDrive should be excluded or not.
func (m *ExcludedApps) SetOneDrive(value *bool)() {
    err := m.GetBackingStore().Set("oneDrive", value)
    if err != nil {
        panic(err)
    }
}
// SetOneNote sets the oneNote property value. The value for if MS Office OneNote should be excluded or not.
func (m *ExcludedApps) SetOneNote(value *bool)() {
    err := m.GetBackingStore().Set("oneNote", value)
    if err != nil {
        panic(err)
    }
}
// SetOutlook sets the outlook property value. The value for if MS Office Outlook should be excluded or not.
func (m *ExcludedApps) SetOutlook(value *bool)() {
    err := m.GetBackingStore().Set("outlook", value)
    if err != nil {
        panic(err)
    }
}
// SetPowerPoint sets the powerPoint property value. The value for if MS Office PowerPoint should be excluded or not.
func (m *ExcludedApps) SetPowerPoint(value *bool)() {
    err := m.GetBackingStore().Set("powerPoint", value)
    if err != nil {
        panic(err)
    }
}
// SetPublisher sets the publisher property value. The value for if MS Office Publisher should be excluded or not.
func (m *ExcludedApps) SetPublisher(value *bool)() {
    err := m.GetBackingStore().Set("publisher", value)
    if err != nil {
        panic(err)
    }
}
// SetSharePointDesigner sets the sharePointDesigner property value. The value for if MS Office SharePointDesigner should be excluded or not.
func (m *ExcludedApps) SetSharePointDesigner(value *bool)() {
    err := m.GetBackingStore().Set("sharePointDesigner", value)
    if err != nil {
        panic(err)
    }
}
// SetTeams sets the teams property value. The value for if MS Office Teams should be excluded or not.
func (m *ExcludedApps) SetTeams(value *bool)() {
    err := m.GetBackingStore().Set("teams", value)
    if err != nil {
        panic(err)
    }
}
// SetVisio sets the visio property value. The value for if MS Office Visio should be excluded or not.
func (m *ExcludedApps) SetVisio(value *bool)() {
    err := m.GetBackingStore().Set("visio", value)
    if err != nil {
        panic(err)
    }
}
// SetWord sets the word property value. The value for if MS Office Word should be excluded or not.
func (m *ExcludedApps) SetWord(value *bool)() {
    err := m.GetBackingStore().Set("word", value)
    if err != nil {
        panic(err)
    }
}
type ExcludedAppsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccess()(*bool)
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetBing()(*bool)
    GetExcel()(*bool)
    GetGroove()(*bool)
    GetInfoPath()(*bool)
    GetLync()(*bool)
    GetOdataType()(*string)
    GetOneDrive()(*bool)
    GetOneNote()(*bool)
    GetOutlook()(*bool)
    GetPowerPoint()(*bool)
    GetPublisher()(*bool)
    GetSharePointDesigner()(*bool)
    GetTeams()(*bool)
    GetVisio()(*bool)
    GetWord()(*bool)
    SetAccess(value *bool)()
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetBing(value *bool)()
    SetExcel(value *bool)()
    SetGroove(value *bool)()
    SetInfoPath(value *bool)()
    SetLync(value *bool)()
    SetOdataType(value *string)()
    SetOneDrive(value *bool)()
    SetOneNote(value *bool)()
    SetOutlook(value *bool)()
    SetPowerPoint(value *bool)()
    SetPublisher(value *bool)()
    SetSharePointDesigner(value *bool)()
    SetTeams(value *bool)()
    SetVisio(value *bool)()
    SetWord(value *bool)()
}
