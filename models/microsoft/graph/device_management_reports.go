package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type DeviceManagementReports struct {
    Entity
    cachedReportConfigurations []DeviceManagementCachedReportConfiguration;
    exportJobs []DeviceManagementExportJob;
    reportSchedules []DeviceManagementReportSchedule;
}
func NewDeviceManagementReports()(*DeviceManagementReports) {
    m := &DeviceManagementReports{
        Entity: *NewEntity(),
    }
    return m
}
func (m *DeviceManagementReports) GetCachedReportConfigurations()([]DeviceManagementCachedReportConfiguration) {
    if m == nil {
        return nil
    } else {
        return m.cachedReportConfigurations
    }
}
func (m *DeviceManagementReports) GetExportJobs()([]DeviceManagementExportJob) {
    if m == nil {
        return nil
    } else {
        return m.exportJobs
    }
}
func (m *DeviceManagementReports) GetReportSchedules()([]DeviceManagementReportSchedule) {
    if m == nil {
        return nil
    } else {
        return m.reportSchedules
    }
}
func (m *DeviceManagementReports) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["cachedReportConfigurations"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceManagementCachedReportConfiguration() })
        if err != nil {
            return err
        }
        res := make([]DeviceManagementCachedReportConfiguration, len(val))
        for i, v := range val {
            res[i] = *(v.(*DeviceManagementCachedReportConfiguration))
        }
        m.SetCachedReportConfigurations(res)
        return nil
    }
    res["exportJobs"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceManagementExportJob() })
        if err != nil {
            return err
        }
        res := make([]DeviceManagementExportJob, len(val))
        for i, v := range val {
            res[i] = *(v.(*DeviceManagementExportJob))
        }
        m.SetExportJobs(res)
        return nil
    }
    res["reportSchedules"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceManagementReportSchedule() })
        if err != nil {
            return err
        }
        res := make([]DeviceManagementReportSchedule, len(val))
        for i, v := range val {
            res[i] = *(v.(*DeviceManagementReportSchedule))
        }
        m.SetReportSchedules(res)
        return nil
    }
    return res
}
func (m *DeviceManagementReports) IsNil()(bool) {
    return m == nil
}
func (m *DeviceManagementReports) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetCachedReportConfigurations()))
        for i, v := range m.GetCachedReportConfigurations() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("cachedReportConfigurations", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetExportJobs()))
        for i, v := range m.GetExportJobs() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("exportJobs", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetReportSchedules()))
        for i, v := range m.GetReportSchedules() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("reportSchedules", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *DeviceManagementReports) SetCachedReportConfigurations(value []DeviceManagementCachedReportConfiguration)() {
    m.cachedReportConfigurations = value
}
func (m *DeviceManagementReports) SetExportJobs(value []DeviceManagementExportJob)() {
    m.exportJobs = value
}
func (m *DeviceManagementReports) SetReportSchedules(value []DeviceManagementReportSchedule)() {
    m.reportSchedules = value
}
