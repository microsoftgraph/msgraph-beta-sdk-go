// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type SharePointMigrationJobProgressEvent struct {
    SharePointMigrationEvent
}
// NewSharePointMigrationJobProgressEvent instantiates a new SharePointMigrationJobProgressEvent and sets the default values.
func NewSharePointMigrationJobProgressEvent()(*SharePointMigrationJobProgressEvent) {
    m := &SharePointMigrationJobProgressEvent{
        SharePointMigrationEvent: *NewSharePointMigrationEvent(),
    }
    return m
}
// CreateSharePointMigrationJobProgressEventFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSharePointMigrationJobProgressEventFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSharePointMigrationJobProgressEvent(), nil
}
// GetBytesProcessed gets the bytesProcessed property value. The bytesProcessed property
// returns a *int64 when successful
func (m *SharePointMigrationJobProgressEvent) GetBytesProcessed()(*int64) {
    val, err := m.GetBackingStore().Get("bytesProcessed")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// GetBytesProcessedOnlyCurrentVersion gets the bytesProcessedOnlyCurrentVersion property value. The bytesProcessedOnlyCurrentVersion property
// returns a *int64 when successful
func (m *SharePointMigrationJobProgressEvent) GetBytesProcessedOnlyCurrentVersion()(*int64) {
    val, err := m.GetBackingStore().Get("bytesProcessedOnlyCurrentVersion")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// GetCpuDurationMs gets the cpuDurationMs property value. The cpuDurationMs property
// returns a *int64 when successful
func (m *SharePointMigrationJobProgressEvent) GetCpuDurationMs()(*int64) {
    val, err := m.GetBackingStore().Get("cpuDurationMs")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *SharePointMigrationJobProgressEvent) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.SharePointMigrationEvent.GetFieldDeserializers()
    res["bytesProcessed"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBytesProcessed(val)
        }
        return nil
    }
    res["bytesProcessedOnlyCurrentVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBytesProcessedOnlyCurrentVersion(val)
        }
        return nil
    }
    res["cpuDurationMs"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCpuDurationMs(val)
        }
        return nil
    }
    res["filesProcessed"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFilesProcessed(val)
        }
        return nil
    }
    res["filesProcessedOnlyCurrentVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFilesProcessedOnlyCurrentVersion(val)
        }
        return nil
    }
    res["isCompleted"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsCompleted(val)
        }
        return nil
    }
    res["lastProcessedObjectId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastProcessedObjectId(val)
        }
        return nil
    }
    res["objectsProcessed"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetObjectsProcessed(val)
        }
        return nil
    }
    res["sqlDurationMs"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSqlDurationMs(val)
        }
        return nil
    }
    res["sqlQueryCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSqlQueryCount(val)
        }
        return nil
    }
    res["totalDurationMs"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalDurationMs(val)
        }
        return nil
    }
    res["totalErrors"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalErrors(val)
        }
        return nil
    }
    res["totalExpectedBytes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalExpectedBytes(val)
        }
        return nil
    }
    res["totalExpectedObjects"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalExpectedObjects(val)
        }
        return nil
    }
    res["totalRetryCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalRetryCount(val)
        }
        return nil
    }
    res["totalWarnings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalWarnings(val)
        }
        return nil
    }
    res["waitTimeOnSqlThrottlingMs"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWaitTimeOnSqlThrottlingMs(val)
        }
        return nil
    }
    return res
}
// GetFilesProcessed gets the filesProcessed property value. The filesProcessed property
// returns a *int64 when successful
func (m *SharePointMigrationJobProgressEvent) GetFilesProcessed()(*int64) {
    val, err := m.GetBackingStore().Get("filesProcessed")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// GetFilesProcessedOnlyCurrentVersion gets the filesProcessedOnlyCurrentVersion property value. The filesProcessedOnlyCurrentVersion property
// returns a *int64 when successful
func (m *SharePointMigrationJobProgressEvent) GetFilesProcessedOnlyCurrentVersion()(*int64) {
    val, err := m.GetBackingStore().Get("filesProcessedOnlyCurrentVersion")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// GetIsCompleted gets the isCompleted property value. The isCompleted property
// returns a *bool when successful
func (m *SharePointMigrationJobProgressEvent) GetIsCompleted()(*bool) {
    val, err := m.GetBackingStore().Get("isCompleted")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetLastProcessedObjectId gets the lastProcessedObjectId property value. The lastProcessedObjectId property
// returns a *string when successful
func (m *SharePointMigrationJobProgressEvent) GetLastProcessedObjectId()(*string) {
    val, err := m.GetBackingStore().Get("lastProcessedObjectId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetObjectsProcessed gets the objectsProcessed property value. The objectsProcessed property
// returns a *int64 when successful
func (m *SharePointMigrationJobProgressEvent) GetObjectsProcessed()(*int64) {
    val, err := m.GetBackingStore().Get("objectsProcessed")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// GetSqlDurationMs gets the sqlDurationMs property value. The sqlDurationMs property
// returns a *int64 when successful
func (m *SharePointMigrationJobProgressEvent) GetSqlDurationMs()(*int64) {
    val, err := m.GetBackingStore().Get("sqlDurationMs")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// GetSqlQueryCount gets the sqlQueryCount property value. The sqlQueryCount property
// returns a *int64 when successful
func (m *SharePointMigrationJobProgressEvent) GetSqlQueryCount()(*int64) {
    val, err := m.GetBackingStore().Get("sqlQueryCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// GetTotalDurationMs gets the totalDurationMs property value. The totalDurationMs property
// returns a *int64 when successful
func (m *SharePointMigrationJobProgressEvent) GetTotalDurationMs()(*int64) {
    val, err := m.GetBackingStore().Get("totalDurationMs")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// GetTotalErrors gets the totalErrors property value. The totalErrors property
// returns a *int64 when successful
func (m *SharePointMigrationJobProgressEvent) GetTotalErrors()(*int64) {
    val, err := m.GetBackingStore().Get("totalErrors")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// GetTotalExpectedBytes gets the totalExpectedBytes property value. The totalExpectedBytes property
// returns a *int64 when successful
func (m *SharePointMigrationJobProgressEvent) GetTotalExpectedBytes()(*int64) {
    val, err := m.GetBackingStore().Get("totalExpectedBytes")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// GetTotalExpectedObjects gets the totalExpectedObjects property value. The totalExpectedObjects property
// returns a *int64 when successful
func (m *SharePointMigrationJobProgressEvent) GetTotalExpectedObjects()(*int64) {
    val, err := m.GetBackingStore().Get("totalExpectedObjects")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// GetTotalRetryCount gets the totalRetryCount property value. The totalRetryCount property
// returns a *int32 when successful
func (m *SharePointMigrationJobProgressEvent) GetTotalRetryCount()(*int32) {
    val, err := m.GetBackingStore().Get("totalRetryCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetTotalWarnings gets the totalWarnings property value. The totalWarnings property
// returns a *int64 when successful
func (m *SharePointMigrationJobProgressEvent) GetTotalWarnings()(*int64) {
    val, err := m.GetBackingStore().Get("totalWarnings")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// GetWaitTimeOnSqlThrottlingMs gets the waitTimeOnSqlThrottlingMs property value. The waitTimeOnSqlThrottlingMs property
// returns a *int64 when successful
func (m *SharePointMigrationJobProgressEvent) GetWaitTimeOnSqlThrottlingMs()(*int64) {
    val, err := m.GetBackingStore().Get("waitTimeOnSqlThrottlingMs")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// Serialize serializes information the current object
func (m *SharePointMigrationJobProgressEvent) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.SharePointMigrationEvent.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt64Value("bytesProcessed", m.GetBytesProcessed())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("bytesProcessedOnlyCurrentVersion", m.GetBytesProcessedOnlyCurrentVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("cpuDurationMs", m.GetCpuDurationMs())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("filesProcessed", m.GetFilesProcessed())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("filesProcessedOnlyCurrentVersion", m.GetFilesProcessedOnlyCurrentVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isCompleted", m.GetIsCompleted())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("lastProcessedObjectId", m.GetLastProcessedObjectId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("objectsProcessed", m.GetObjectsProcessed())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("sqlDurationMs", m.GetSqlDurationMs())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("sqlQueryCount", m.GetSqlQueryCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("totalDurationMs", m.GetTotalDurationMs())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("totalErrors", m.GetTotalErrors())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("totalExpectedBytes", m.GetTotalExpectedBytes())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("totalExpectedObjects", m.GetTotalExpectedObjects())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("totalRetryCount", m.GetTotalRetryCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("totalWarnings", m.GetTotalWarnings())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("waitTimeOnSqlThrottlingMs", m.GetWaitTimeOnSqlThrottlingMs())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetBytesProcessed sets the bytesProcessed property value. The bytesProcessed property
func (m *SharePointMigrationJobProgressEvent) SetBytesProcessed(value *int64)() {
    err := m.GetBackingStore().Set("bytesProcessed", value)
    if err != nil {
        panic(err)
    }
}
// SetBytesProcessedOnlyCurrentVersion sets the bytesProcessedOnlyCurrentVersion property value. The bytesProcessedOnlyCurrentVersion property
func (m *SharePointMigrationJobProgressEvent) SetBytesProcessedOnlyCurrentVersion(value *int64)() {
    err := m.GetBackingStore().Set("bytesProcessedOnlyCurrentVersion", value)
    if err != nil {
        panic(err)
    }
}
// SetCpuDurationMs sets the cpuDurationMs property value. The cpuDurationMs property
func (m *SharePointMigrationJobProgressEvent) SetCpuDurationMs(value *int64)() {
    err := m.GetBackingStore().Set("cpuDurationMs", value)
    if err != nil {
        panic(err)
    }
}
// SetFilesProcessed sets the filesProcessed property value. The filesProcessed property
func (m *SharePointMigrationJobProgressEvent) SetFilesProcessed(value *int64)() {
    err := m.GetBackingStore().Set("filesProcessed", value)
    if err != nil {
        panic(err)
    }
}
// SetFilesProcessedOnlyCurrentVersion sets the filesProcessedOnlyCurrentVersion property value. The filesProcessedOnlyCurrentVersion property
func (m *SharePointMigrationJobProgressEvent) SetFilesProcessedOnlyCurrentVersion(value *int64)() {
    err := m.GetBackingStore().Set("filesProcessedOnlyCurrentVersion", value)
    if err != nil {
        panic(err)
    }
}
// SetIsCompleted sets the isCompleted property value. The isCompleted property
func (m *SharePointMigrationJobProgressEvent) SetIsCompleted(value *bool)() {
    err := m.GetBackingStore().Set("isCompleted", value)
    if err != nil {
        panic(err)
    }
}
// SetLastProcessedObjectId sets the lastProcessedObjectId property value. The lastProcessedObjectId property
func (m *SharePointMigrationJobProgressEvent) SetLastProcessedObjectId(value *string)() {
    err := m.GetBackingStore().Set("lastProcessedObjectId", value)
    if err != nil {
        panic(err)
    }
}
// SetObjectsProcessed sets the objectsProcessed property value. The objectsProcessed property
func (m *SharePointMigrationJobProgressEvent) SetObjectsProcessed(value *int64)() {
    err := m.GetBackingStore().Set("objectsProcessed", value)
    if err != nil {
        panic(err)
    }
}
// SetSqlDurationMs sets the sqlDurationMs property value. The sqlDurationMs property
func (m *SharePointMigrationJobProgressEvent) SetSqlDurationMs(value *int64)() {
    err := m.GetBackingStore().Set("sqlDurationMs", value)
    if err != nil {
        panic(err)
    }
}
// SetSqlQueryCount sets the sqlQueryCount property value. The sqlQueryCount property
func (m *SharePointMigrationJobProgressEvent) SetSqlQueryCount(value *int64)() {
    err := m.GetBackingStore().Set("sqlQueryCount", value)
    if err != nil {
        panic(err)
    }
}
// SetTotalDurationMs sets the totalDurationMs property value. The totalDurationMs property
func (m *SharePointMigrationJobProgressEvent) SetTotalDurationMs(value *int64)() {
    err := m.GetBackingStore().Set("totalDurationMs", value)
    if err != nil {
        panic(err)
    }
}
// SetTotalErrors sets the totalErrors property value. The totalErrors property
func (m *SharePointMigrationJobProgressEvent) SetTotalErrors(value *int64)() {
    err := m.GetBackingStore().Set("totalErrors", value)
    if err != nil {
        panic(err)
    }
}
// SetTotalExpectedBytes sets the totalExpectedBytes property value. The totalExpectedBytes property
func (m *SharePointMigrationJobProgressEvent) SetTotalExpectedBytes(value *int64)() {
    err := m.GetBackingStore().Set("totalExpectedBytes", value)
    if err != nil {
        panic(err)
    }
}
// SetTotalExpectedObjects sets the totalExpectedObjects property value. The totalExpectedObjects property
func (m *SharePointMigrationJobProgressEvent) SetTotalExpectedObjects(value *int64)() {
    err := m.GetBackingStore().Set("totalExpectedObjects", value)
    if err != nil {
        panic(err)
    }
}
// SetTotalRetryCount sets the totalRetryCount property value. The totalRetryCount property
func (m *SharePointMigrationJobProgressEvent) SetTotalRetryCount(value *int32)() {
    err := m.GetBackingStore().Set("totalRetryCount", value)
    if err != nil {
        panic(err)
    }
}
// SetTotalWarnings sets the totalWarnings property value. The totalWarnings property
func (m *SharePointMigrationJobProgressEvent) SetTotalWarnings(value *int64)() {
    err := m.GetBackingStore().Set("totalWarnings", value)
    if err != nil {
        panic(err)
    }
}
// SetWaitTimeOnSqlThrottlingMs sets the waitTimeOnSqlThrottlingMs property value. The waitTimeOnSqlThrottlingMs property
func (m *SharePointMigrationJobProgressEvent) SetWaitTimeOnSqlThrottlingMs(value *int64)() {
    err := m.GetBackingStore().Set("waitTimeOnSqlThrottlingMs", value)
    if err != nil {
        panic(err)
    }
}
type SharePointMigrationJobProgressEventable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    SharePointMigrationEventable
    GetBytesProcessed()(*int64)
    GetBytesProcessedOnlyCurrentVersion()(*int64)
    GetCpuDurationMs()(*int64)
    GetFilesProcessed()(*int64)
    GetFilesProcessedOnlyCurrentVersion()(*int64)
    GetIsCompleted()(*bool)
    GetLastProcessedObjectId()(*string)
    GetObjectsProcessed()(*int64)
    GetSqlDurationMs()(*int64)
    GetSqlQueryCount()(*int64)
    GetTotalDurationMs()(*int64)
    GetTotalErrors()(*int64)
    GetTotalExpectedBytes()(*int64)
    GetTotalExpectedObjects()(*int64)
    GetTotalRetryCount()(*int32)
    GetTotalWarnings()(*int64)
    GetWaitTimeOnSqlThrottlingMs()(*int64)
    SetBytesProcessed(value *int64)()
    SetBytesProcessedOnlyCurrentVersion(value *int64)()
    SetCpuDurationMs(value *int64)()
    SetFilesProcessed(value *int64)()
    SetFilesProcessedOnlyCurrentVersion(value *int64)()
    SetIsCompleted(value *bool)()
    SetLastProcessedObjectId(value *string)()
    SetObjectsProcessed(value *int64)()
    SetSqlDurationMs(value *int64)()
    SetSqlQueryCount(value *int64)()
    SetTotalDurationMs(value *int64)()
    SetTotalErrors(value *int64)()
    SetTotalExpectedBytes(value *int64)()
    SetTotalExpectedObjects(value *int64)()
    SetTotalRetryCount(value *int32)()
    SetTotalWarnings(value *int64)()
    SetWaitTimeOnSqlThrottlingMs(value *int64)()
}
