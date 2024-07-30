package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

type PositionDetail struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewPositionDetail instantiates a new PositionDetail and sets the default values.
func NewPositionDetail()(*PositionDetail) {
    m := &PositionDetail{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreatePositionDetailFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreatePositionDetailFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPositionDetail(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *PositionDetail) GetAdditionalData()(map[string]any) {
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
func (m *PositionDetail) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetCompany gets the company property value. Detail about the company or employer.
// returns a CompanyDetailable when successful
func (m *PositionDetail) GetCompany()(CompanyDetailable) {
    val, err := m.GetBackingStore().Get("company")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(CompanyDetailable)
    }
    return nil
}
// GetDescription gets the description property value. Description of the position in question.
// returns a *string when successful
func (m *PositionDetail) GetDescription()(*string) {
    val, err := m.GetBackingStore().Get("description")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetEndMonthYear gets the endMonthYear property value. When the position ended.
// returns a *DateOnly when successful
func (m *PositionDetail) GetEndMonthYear()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
    val, err := m.GetBackingStore().Get("endMonthYear")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *PositionDetail) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["company"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCompanyDetailFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCompany(val.(CompanyDetailable))
        }
        return nil
    }
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["endMonthYear"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetDateOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEndMonthYear(val)
        }
        return nil
    }
    res["jobTitle"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetJobTitle(val)
        }
        return nil
    }
    res["layer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLayer(val)
        }
        return nil
    }
    res["level"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLevel(val)
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
    res["role"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRole(val)
        }
        return nil
    }
    res["secondaryJobTitle"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecondaryJobTitle(val)
        }
        return nil
    }
    res["secondaryRole"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecondaryRole(val)
        }
        return nil
    }
    res["startMonthYear"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetDateOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartMonthYear(val)
        }
        return nil
    }
    res["summary"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSummary(val)
        }
        return nil
    }
    return res
}
// GetJobTitle gets the jobTitle property value. The title held when in that position.
// returns a *string when successful
func (m *PositionDetail) GetJobTitle()(*string) {
    val, err := m.GetBackingStore().Get("jobTitle")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetLayer gets the layer property value. The place where the employee is within the organizational hierarchy.
// returns a *int32 when successful
func (m *PositionDetail) GetLayer()(*int32) {
    val, err := m.GetBackingStore().Get("layer")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetLevel gets the level property value. The employee’s experience or management level.
// returns a *string when successful
func (m *PositionDetail) GetLevel()(*string) {
    val, err := m.GetBackingStore().Get("level")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *PositionDetail) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetRole gets the role property value. The role the position entailed.
// returns a *string when successful
func (m *PositionDetail) GetRole()(*string) {
    val, err := m.GetBackingStore().Get("role")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSecondaryJobTitle gets the secondaryJobTitle property value. The secondaryJobTitle property
// returns a *string when successful
func (m *PositionDetail) GetSecondaryJobTitle()(*string) {
    val, err := m.GetBackingStore().Get("secondaryJobTitle")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSecondaryRole gets the secondaryRole property value. The secondaryRole property
// returns a *string when successful
func (m *PositionDetail) GetSecondaryRole()(*string) {
    val, err := m.GetBackingStore().Get("secondaryRole")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetStartMonthYear gets the startMonthYear property value. The start month and year of the position.
// returns a *DateOnly when successful
func (m *PositionDetail) GetStartMonthYear()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
    val, err := m.GetBackingStore().Get("startMonthYear")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
    }
    return nil
}
// GetSummary gets the summary property value. summary of the position.
// returns a *string when successful
func (m *PositionDetail) GetSummary()(*string) {
    val, err := m.GetBackingStore().Get("summary")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *PositionDetail) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("company", m.GetCompany())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteDateOnlyValue("endMonthYear", m.GetEndMonthYear())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("jobTitle", m.GetJobTitle())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("layer", m.GetLayer())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("level", m.GetLevel())
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
        err := writer.WriteStringValue("role", m.GetRole())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("secondaryJobTitle", m.GetSecondaryJobTitle())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("secondaryRole", m.GetSecondaryRole())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteDateOnlyValue("startMonthYear", m.GetStartMonthYear())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("summary", m.GetSummary())
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
func (m *PositionDetail) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *PositionDetail) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetCompany sets the company property value. Detail about the company or employer.
func (m *PositionDetail) SetCompany(value CompanyDetailable)() {
    err := m.GetBackingStore().Set("company", value)
    if err != nil {
        panic(err)
    }
}
// SetDescription sets the description property value. Description of the position in question.
func (m *PositionDetail) SetDescription(value *string)() {
    err := m.GetBackingStore().Set("description", value)
    if err != nil {
        panic(err)
    }
}
// SetEndMonthYear sets the endMonthYear property value. When the position ended.
func (m *PositionDetail) SetEndMonthYear(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    err := m.GetBackingStore().Set("endMonthYear", value)
    if err != nil {
        panic(err)
    }
}
// SetJobTitle sets the jobTitle property value. The title held when in that position.
func (m *PositionDetail) SetJobTitle(value *string)() {
    err := m.GetBackingStore().Set("jobTitle", value)
    if err != nil {
        panic(err)
    }
}
// SetLayer sets the layer property value. The place where the employee is within the organizational hierarchy.
func (m *PositionDetail) SetLayer(value *int32)() {
    err := m.GetBackingStore().Set("layer", value)
    if err != nil {
        panic(err)
    }
}
// SetLevel sets the level property value. The employee’s experience or management level.
func (m *PositionDetail) SetLevel(value *string)() {
    err := m.GetBackingStore().Set("level", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *PositionDetail) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetRole sets the role property value. The role the position entailed.
func (m *PositionDetail) SetRole(value *string)() {
    err := m.GetBackingStore().Set("role", value)
    if err != nil {
        panic(err)
    }
}
// SetSecondaryJobTitle sets the secondaryJobTitle property value. The secondaryJobTitle property
func (m *PositionDetail) SetSecondaryJobTitle(value *string)() {
    err := m.GetBackingStore().Set("secondaryJobTitle", value)
    if err != nil {
        panic(err)
    }
}
// SetSecondaryRole sets the secondaryRole property value. The secondaryRole property
func (m *PositionDetail) SetSecondaryRole(value *string)() {
    err := m.GetBackingStore().Set("secondaryRole", value)
    if err != nil {
        panic(err)
    }
}
// SetStartMonthYear sets the startMonthYear property value. The start month and year of the position.
func (m *PositionDetail) SetStartMonthYear(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    err := m.GetBackingStore().Set("startMonthYear", value)
    if err != nil {
        panic(err)
    }
}
// SetSummary sets the summary property value. summary of the position.
func (m *PositionDetail) SetSummary(value *string)() {
    err := m.GetBackingStore().Set("summary", value)
    if err != nil {
        panic(err)
    }
}
type PositionDetailable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetCompany()(CompanyDetailable)
    GetDescription()(*string)
    GetEndMonthYear()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
    GetJobTitle()(*string)
    GetLayer()(*int32)
    GetLevel()(*string)
    GetOdataType()(*string)
    GetRole()(*string)
    GetSecondaryJobTitle()(*string)
    GetSecondaryRole()(*string)
    GetStartMonthYear()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
    GetSummary()(*string)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetCompany(value CompanyDetailable)()
    SetDescription(value *string)()
    SetEndMonthYear(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)()
    SetJobTitle(value *string)()
    SetLayer(value *int32)()
    SetLevel(value *string)()
    SetOdataType(value *string)()
    SetRole(value *string)()
    SetSecondaryJobTitle(value *string)()
    SetSecondaryRole(value *string)()
    SetStartMonthYear(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)()
    SetSummary(value *string)()
}
