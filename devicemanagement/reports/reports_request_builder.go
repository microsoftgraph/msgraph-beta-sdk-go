package reports

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i08ed1bcbc2935e740bb1a75a82536d462055f7eb6c3b9f8c54f292796af881e9 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/reports/getmalwaresummaryreport"
    i1d42b2baf1a1b776afa9e69081dc91559e4cdd90109e8ba39f1c3cfd8c36b3b3 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/reports/getcachedreport"
    i246fa2e435cd7622048a0471ca5df979d86e83142f57f108d53701f4d8c17c18 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/reports/getconfigurationsettingdetailsreport"
    i2660e5353accb6f61768bb2883c08d64660fe2c8c67503cb521191b4c4e33e41 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/reports/getconfigurationpolicynoncompliancesummaryreport"
    i2a7e95d28bcab38f85f05fe3f131357d5c408f5fbd47fb213b7aac8dbd71328a "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/reports/cachedreportconfigurations"
    i2abf29284e813349f3083b2dc57f6553c2b467c8f68ad9c82e8950ce8d0afd2d "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/reports/getconfigurationsettingsreport"
    i2cbf3d93d72f9416ce29a9ac2774ddf6c3cceadcccbd462303a2ff5215c59622 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/reports/getpolicynoncompliancemetadata"
    i3e013f91ac3d040684b27abcd0dc56839c52d2faa6d92ef0d2295a7fefa96a5b "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/reports/getappstatusoverviewreport"
    i3f11f18462beffdea839437cba5610e05b58902433c635e68fdfdfb6e5f1c4b7 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/reports/getconfigurationpolicydevicesummaryreport"
    i45eba99273cfc2518cc265806b0cba99d31c259eee99db3fa9a93ca6fbd11d28 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/reports/getcompliancepolicynoncompliancereport"
    i4d73fdd41ba9bd99a3822503da6e6d8345d38f32a116b7860fb94fbdc8a9c79d "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/reports/getactivemalwarereport"
    i5284d7a78de859429101f202a37b95a2897b7a115a509f7ea9be6db8a70482a3 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/reports/getcertificatesreport"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i5b82ca25c293ac194e35ceb34a2eb3523dd2a1568f2e61c9458b2d7fe4f1d970 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/reports/getcompliancepolicynoncompliancesummaryreport"
    i5d960953aeefdfe382eaaeee66b54c5138188e0487e4a969bc66342437b3efc3 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/reports/getpolicynoncompliancereport"
    i711aceb383b63eb4ad8db059fe173d8b78d6bb1bdcfaf386735b20a8e54c2748 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/reports/getcompliancesettingnoncompliancereport"
    i71903f4f495c2ba2d1bec9d9d52c924a6bd43444cea991fb60e1874f824ed54e "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/reports/getfailedmobileappssummaryreport"
    i75076cf3e3d6371cc0bb5d2810006d6e49e086cb781c7867b1093cfc1cf69827 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/reports/getappsinstallsummaryreport"
    i7611b2ea51c277728d8bca812a4642adb44f343f51f28de5c5e087ec400d0645 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/reports/getdevicemanagementintentsettingsreport"
    i8ad2db41d58439c4405d08911118b14826289f800859f884ca28da1387d489c2 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/reports/getsettingnoncompliancereport"
    i95054c49fc60448e61ba13d5075f71dd6ae5b29fcc1f0307e8b5e38c80345642 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/reports/getwindowsupdatealertsummaryreport"
    i9819883d2f88a17d06e7a89c99985f254bca57d86a82a07526740f4057982566 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/reports/getconfigurationpolicynoncompliancereport"
    ia7100ea2edcc786fe9054729a54d7600f750ebcc9be73a96446bf4b9bc605356 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/reports/getdeviceinstallstatusreport"
    iaa2bd508f982ae7d95dc62bbaefe9eb13e0c3fc9ad77d9389f8d82f3fb1dd151 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/reports/getrelatedappsstatusreport"
    iad468b4335f9e19b6476bd9d2cb864bb58150160cddf3825611adcbb6a864f65 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/reports/getreportfilters"
    ib57779f05872bd3be1d691d1b254dbe109c8ac7c77e7a265ef928aedbba3d32d "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/reports/getunhealthydefenderagentsreport"
    ib6b7e73b832065adf3e2ef8550d271d24c62fdb5128ea16eae3808e57561d401 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/reports/getallcertificatesreport"
    ib6f9fa86536428e22220c67df5e0bf322aa3124cb40d39ca6a96b97728e8551f "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/reports/gethistoricalreport"
    ibb77da97a1e31c9bb16ce4681732ad7475ae450918c13097176fab6ef841fcdf "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/reports/getconfigurationpolicysettingsdevicesummaryreport"
    ic029e50a08072e00092e7667a1fc68f437d1be340acce2c03ee7a004dea3e5f1 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/reports/getconfigurationpolicydevicesreport"
    ic09bfed61d7acf9bcd44dfdcb91600937d60cff331615b3b8a208a31e277cd7d "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/reports/getconfigmanagerdevicepolicystatusreport"
    icd46e1255a1aee0bc5f9216ddd5fb29e365c5848ea3edecfd78d864191ca9748 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/reports/getdevicenoncompliancereport"
    icef6de0280daff5ea2bb1854119e785ff95cb81abf14b3546d97c564e7d27123 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/reports/getwindowsqualityupdatealertsperpolicyperdevicereport"
    id1df61af1e3874f797c91e9607c99dfd3cb6f0dbf128a90ab44536b912adf005 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/reports/getconfigurationpoliciesreportfordevice"
    id7b1135fd80c87263a81483b65064413c7251dd0137f48118d1471f835eefab4 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/reports/getfailedmobileappsreport"
    idc692bbd177fe7be2a1d1da6b8436f4433611cc7abff49f069c1a2e0e027e466 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/reports/getunhealthyfirewallsummaryreport"
    idf207024cf64a316ff9c7fa94616461d390fe5977d0687432c3309d78ba029ac "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/reports/getpolicynoncompliancesummaryreport"
    ie6deaed0072be9659d87a74fe95df383a715951d0aea4be5ba203ca890b3c041 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/reports/getactivemalwaresummaryreport"
    iea8e9eceb6829c23694025f220d78170445e17ec05d7949319503a9dfef6f2dc "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/reports/reportschedules"
    ieab9e3b30d90701370eeeb662f2e677617ef77e81a7f0a2a9d2c1d94fab9b995 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/reports/getwindowsupdatealertsperpolicyperdevicereport"
    if4b9775de84c27e94434be0280aaa9a055f85419e5b5857a66bb9b9eed3280f4 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/reports/getunhealthyfirewallreport"
    if7d41ece7783a25f185dc8071dfbd2363227c4fb6d8e08d2a9412f206c7279a9 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/reports/exportjobs"
    if7f28186646455aa9e0df362fc86fe6afcf4c8f20b7f65732c9e4126cdc6c8e0 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/reports/getwindowsqualityupdatealertsummaryreport"
    if876c79e09c00c1578bab3da9f4fcb0cc6d4ae3f9b2e792a2412130b488fcdc0 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/reports/getuserinstallstatusreport"
    if8e67aebead4149dc709bc45370408435c48c5beafd2a598a3552c78692904c1 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/reports/getconfigurationsettingnoncompliancereport"
    i2bf9ce7fb7bd0fa9fc75119c61d3de5e2e31a446f3ebba40eaf7f289544d9943 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/reports/cachedreportconfigurations/item"
    id6b331ae2f0cb2af60ba8b44c9282d7a2ff2e031c38004db15739c1402e76714 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/reports/reportschedules/item"
    if607fcbf4c23ed3b14db77b88b46a29896a853e746acfffd003d215b0fb83a02 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/reports/exportjobs/item"
)

type ReportsRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type ReportsRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Expand []string;
    Select_escpaped []string;
}
func (m *ReportsRequestBuilder) CachedReportConfigurations()(*i2a7e95d28bcab38f85f05fe3f131357d5c408f5fbd47fb213b7aac8dbd71328a.CachedReportConfigurationsRequestBuilder) {
    return i2a7e95d28bcab38f85f05fe3f131357d5c408f5fbd47fb213b7aac8dbd71328a.NewCachedReportConfigurationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ReportsRequestBuilder) CachedReportConfigurationsById(id string)(*i2bf9ce7fb7bd0fa9fc75119c61d3de5e2e31a446f3ebba40eaf7f289544d9943.DeviceManagementCachedReportConfigurationRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["deviceManagementCachedReportConfiguration_id"] = id
    }
    return i2bf9ce7fb7bd0fa9fc75119c61d3de5e2e31a446f3ebba40eaf7f289544d9943.NewDeviceManagementCachedReportConfigurationRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func NewReportsRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ReportsRequestBuilder) {
    m := &ReportsRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/deviceManagement/reports{?select,expand}";
    urlTplParams := make(map[string]string)
    if pathParameters != nil {
        for idx, item := range pathParameters {
            urlTplParams[idx] = item
        }
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
func NewReportsRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ReportsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewReportsRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *ReportsRequestBuilder) CreateDeleteRequestInformation(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.DELETE
    if h != nil {
        err := h(requestInfo.Headers)
        if err != nil {
            return nil, err
        }
    }
    if o != nil {
        err := requestInfo.AddRequestOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
func (m *ReportsRequestBuilder) CreateGetRequestInformation(q func (value *ReportsRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(ReportsRequestBuilderGetQueryParameters)
        err := q(qParams)
        if err != nil {
            return nil, err
        }
        err = qParams.AddQueryParameters(requestInfo.QueryParameters)
        if err != nil {
            return nil, err
        }
    }
    if h != nil {
        err := h(requestInfo.Headers)
        if err != nil {
            return nil, err
        }
    }
    if o != nil {
        err := requestInfo.AddRequestOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
func (m *ReportsRequestBuilder) CreatePatchRequestInformation(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceManagementReports, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if h != nil {
        err := h(requestInfo.Headers)
        if err != nil {
            return nil, err
        }
    }
    if o != nil {
        err := requestInfo.AddRequestOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
func (m *ReportsRequestBuilder) Delete(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(h, o);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, responseHandler)
    if err != nil {
        return err
    }
    return nil
}
func (m *ReportsRequestBuilder) ExportJobs()(*if7d41ece7783a25f185dc8071dfbd2363227c4fb6d8e08d2a9412f206c7279a9.ExportJobsRequestBuilder) {
    return if7d41ece7783a25f185dc8071dfbd2363227c4fb6d8e08d2a9412f206c7279a9.NewExportJobsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ReportsRequestBuilder) ExportJobsById(id string)(*if607fcbf4c23ed3b14db77b88b46a29896a853e746acfffd003d215b0fb83a02.DeviceManagementExportJobRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["deviceManagementExportJob_id"] = id
    }
    return if607fcbf4c23ed3b14db77b88b46a29896a853e746acfffd003d215b0fb83a02.NewDeviceManagementExportJobRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ReportsRequestBuilder) Get(q func (value *ReportsRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceManagementReports, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewDeviceManagementReports() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceManagementReports), nil
}
func (m *ReportsRequestBuilder) GetActiveMalwareReport()(*i4d73fdd41ba9bd99a3822503da6e6d8345d38f32a116b7860fb94fbdc8a9c79d.GetActiveMalwareReportRequestBuilder) {
    return i4d73fdd41ba9bd99a3822503da6e6d8345d38f32a116b7860fb94fbdc8a9c79d.NewGetActiveMalwareReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ReportsRequestBuilder) GetActiveMalwareSummaryReport()(*ie6deaed0072be9659d87a74fe95df383a715951d0aea4be5ba203ca890b3c041.GetActiveMalwareSummaryReportRequestBuilder) {
    return ie6deaed0072be9659d87a74fe95df383a715951d0aea4be5ba203ca890b3c041.NewGetActiveMalwareSummaryReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ReportsRequestBuilder) GetAllCertificatesReport()(*ib6b7e73b832065adf3e2ef8550d271d24c62fdb5128ea16eae3808e57561d401.GetAllCertificatesReportRequestBuilder) {
    return ib6b7e73b832065adf3e2ef8550d271d24c62fdb5128ea16eae3808e57561d401.NewGetAllCertificatesReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ReportsRequestBuilder) GetAppsInstallSummaryReport()(*i75076cf3e3d6371cc0bb5d2810006d6e49e086cb781c7867b1093cfc1cf69827.GetAppsInstallSummaryReportRequestBuilder) {
    return i75076cf3e3d6371cc0bb5d2810006d6e49e086cb781c7867b1093cfc1cf69827.NewGetAppsInstallSummaryReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ReportsRequestBuilder) GetAppStatusOverviewReport()(*i3e013f91ac3d040684b27abcd0dc56839c52d2faa6d92ef0d2295a7fefa96a5b.GetAppStatusOverviewReportRequestBuilder) {
    return i3e013f91ac3d040684b27abcd0dc56839c52d2faa6d92ef0d2295a7fefa96a5b.NewGetAppStatusOverviewReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ReportsRequestBuilder) GetCachedReport()(*i1d42b2baf1a1b776afa9e69081dc91559e4cdd90109e8ba39f1c3cfd8c36b3b3.GetCachedReportRequestBuilder) {
    return i1d42b2baf1a1b776afa9e69081dc91559e4cdd90109e8ba39f1c3cfd8c36b3b3.NewGetCachedReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ReportsRequestBuilder) GetCertificatesReport()(*i5284d7a78de859429101f202a37b95a2897b7a115a509f7ea9be6db8a70482a3.GetCertificatesReportRequestBuilder) {
    return i5284d7a78de859429101f202a37b95a2897b7a115a509f7ea9be6db8a70482a3.NewGetCertificatesReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ReportsRequestBuilder) GetCompliancePolicyNonComplianceReport()(*i45eba99273cfc2518cc265806b0cba99d31c259eee99db3fa9a93ca6fbd11d28.GetCompliancePolicyNonComplianceReportRequestBuilder) {
    return i45eba99273cfc2518cc265806b0cba99d31c259eee99db3fa9a93ca6fbd11d28.NewGetCompliancePolicyNonComplianceReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ReportsRequestBuilder) GetCompliancePolicyNonComplianceSummaryReport()(*i5b82ca25c293ac194e35ceb34a2eb3523dd2a1568f2e61c9458b2d7fe4f1d970.GetCompliancePolicyNonComplianceSummaryReportRequestBuilder) {
    return i5b82ca25c293ac194e35ceb34a2eb3523dd2a1568f2e61c9458b2d7fe4f1d970.NewGetCompliancePolicyNonComplianceSummaryReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ReportsRequestBuilder) GetComplianceSettingNonComplianceReport()(*i711aceb383b63eb4ad8db059fe173d8b78d6bb1bdcfaf386735b20a8e54c2748.GetComplianceSettingNonComplianceReportRequestBuilder) {
    return i711aceb383b63eb4ad8db059fe173d8b78d6bb1bdcfaf386735b20a8e54c2748.NewGetComplianceSettingNonComplianceReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ReportsRequestBuilder) GetConfigManagerDevicePolicyStatusReport()(*ic09bfed61d7acf9bcd44dfdcb91600937d60cff331615b3b8a208a31e277cd7d.GetConfigManagerDevicePolicyStatusReportRequestBuilder) {
    return ic09bfed61d7acf9bcd44dfdcb91600937d60cff331615b3b8a208a31e277cd7d.NewGetConfigManagerDevicePolicyStatusReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ReportsRequestBuilder) GetConfigurationPoliciesReportForDevice()(*id1df61af1e3874f797c91e9607c99dfd3cb6f0dbf128a90ab44536b912adf005.GetConfigurationPoliciesReportForDeviceRequestBuilder) {
    return id1df61af1e3874f797c91e9607c99dfd3cb6f0dbf128a90ab44536b912adf005.NewGetConfigurationPoliciesReportForDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ReportsRequestBuilder) GetConfigurationPolicyDevicesReport()(*ic029e50a08072e00092e7667a1fc68f437d1be340acce2c03ee7a004dea3e5f1.GetConfigurationPolicyDevicesReportRequestBuilder) {
    return ic029e50a08072e00092e7667a1fc68f437d1be340acce2c03ee7a004dea3e5f1.NewGetConfigurationPolicyDevicesReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ReportsRequestBuilder) GetConfigurationPolicyDeviceSummaryReport()(*i3f11f18462beffdea839437cba5610e05b58902433c635e68fdfdfb6e5f1c4b7.GetConfigurationPolicyDeviceSummaryReportRequestBuilder) {
    return i3f11f18462beffdea839437cba5610e05b58902433c635e68fdfdfb6e5f1c4b7.NewGetConfigurationPolicyDeviceSummaryReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ReportsRequestBuilder) GetConfigurationPolicyNonComplianceReport()(*i9819883d2f88a17d06e7a89c99985f254bca57d86a82a07526740f4057982566.GetConfigurationPolicyNonComplianceReportRequestBuilder) {
    return i9819883d2f88a17d06e7a89c99985f254bca57d86a82a07526740f4057982566.NewGetConfigurationPolicyNonComplianceReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ReportsRequestBuilder) GetConfigurationPolicyNonComplianceSummaryReport()(*i2660e5353accb6f61768bb2883c08d64660fe2c8c67503cb521191b4c4e33e41.GetConfigurationPolicyNonComplianceSummaryReportRequestBuilder) {
    return i2660e5353accb6f61768bb2883c08d64660fe2c8c67503cb521191b4c4e33e41.NewGetConfigurationPolicyNonComplianceSummaryReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ReportsRequestBuilder) GetConfigurationPolicySettingsDeviceSummaryReport()(*ibb77da97a1e31c9bb16ce4681732ad7475ae450918c13097176fab6ef841fcdf.GetConfigurationPolicySettingsDeviceSummaryReportRequestBuilder) {
    return ibb77da97a1e31c9bb16ce4681732ad7475ae450918c13097176fab6ef841fcdf.NewGetConfigurationPolicySettingsDeviceSummaryReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ReportsRequestBuilder) GetConfigurationSettingDetailsReport()(*i246fa2e435cd7622048a0471ca5df979d86e83142f57f108d53701f4d8c17c18.GetConfigurationSettingDetailsReportRequestBuilder) {
    return i246fa2e435cd7622048a0471ca5df979d86e83142f57f108d53701f4d8c17c18.NewGetConfigurationSettingDetailsReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ReportsRequestBuilder) GetConfigurationSettingNonComplianceReport()(*if8e67aebead4149dc709bc45370408435c48c5beafd2a598a3552c78692904c1.GetConfigurationSettingNonComplianceReportRequestBuilder) {
    return if8e67aebead4149dc709bc45370408435c48c5beafd2a598a3552c78692904c1.NewGetConfigurationSettingNonComplianceReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ReportsRequestBuilder) GetConfigurationSettingsReport()(*i2abf29284e813349f3083b2dc57f6553c2b467c8f68ad9c82e8950ce8d0afd2d.GetConfigurationSettingsReportRequestBuilder) {
    return i2abf29284e813349f3083b2dc57f6553c2b467c8f68ad9c82e8950ce8d0afd2d.NewGetConfigurationSettingsReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ReportsRequestBuilder) GetDeviceInstallStatusReport()(*ia7100ea2edcc786fe9054729a54d7600f750ebcc9be73a96446bf4b9bc605356.GetDeviceInstallStatusReportRequestBuilder) {
    return ia7100ea2edcc786fe9054729a54d7600f750ebcc9be73a96446bf4b9bc605356.NewGetDeviceInstallStatusReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ReportsRequestBuilder) GetDeviceManagementIntentSettingsReport()(*i7611b2ea51c277728d8bca812a4642adb44f343f51f28de5c5e087ec400d0645.GetDeviceManagementIntentSettingsReportRequestBuilder) {
    return i7611b2ea51c277728d8bca812a4642adb44f343f51f28de5c5e087ec400d0645.NewGetDeviceManagementIntentSettingsReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ReportsRequestBuilder) GetDeviceNonComplianceReport()(*icd46e1255a1aee0bc5f9216ddd5fb29e365c5848ea3edecfd78d864191ca9748.GetDeviceNonComplianceReportRequestBuilder) {
    return icd46e1255a1aee0bc5f9216ddd5fb29e365c5848ea3edecfd78d864191ca9748.NewGetDeviceNonComplianceReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ReportsRequestBuilder) GetFailedMobileAppsReport()(*id7b1135fd80c87263a81483b65064413c7251dd0137f48118d1471f835eefab4.GetFailedMobileAppsReportRequestBuilder) {
    return id7b1135fd80c87263a81483b65064413c7251dd0137f48118d1471f835eefab4.NewGetFailedMobileAppsReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ReportsRequestBuilder) GetFailedMobileAppsSummaryReport()(*i71903f4f495c2ba2d1bec9d9d52c924a6bd43444cea991fb60e1874f824ed54e.GetFailedMobileAppsSummaryReportRequestBuilder) {
    return i71903f4f495c2ba2d1bec9d9d52c924a6bd43444cea991fb60e1874f824ed54e.NewGetFailedMobileAppsSummaryReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ReportsRequestBuilder) GetHistoricalReport()(*ib6f9fa86536428e22220c67df5e0bf322aa3124cb40d39ca6a96b97728e8551f.GetHistoricalReportRequestBuilder) {
    return ib6f9fa86536428e22220c67df5e0bf322aa3124cb40d39ca6a96b97728e8551f.NewGetHistoricalReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ReportsRequestBuilder) GetMalwareSummaryReport()(*i08ed1bcbc2935e740bb1a75a82536d462055f7eb6c3b9f8c54f292796af881e9.GetMalwareSummaryReportRequestBuilder) {
    return i08ed1bcbc2935e740bb1a75a82536d462055f7eb6c3b9f8c54f292796af881e9.NewGetMalwareSummaryReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ReportsRequestBuilder) GetPolicyNonComplianceMetadata()(*i2cbf3d93d72f9416ce29a9ac2774ddf6c3cceadcccbd462303a2ff5215c59622.GetPolicyNonComplianceMetadataRequestBuilder) {
    return i2cbf3d93d72f9416ce29a9ac2774ddf6c3cceadcccbd462303a2ff5215c59622.NewGetPolicyNonComplianceMetadataRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ReportsRequestBuilder) GetPolicyNonComplianceReport()(*i5d960953aeefdfe382eaaeee66b54c5138188e0487e4a969bc66342437b3efc3.GetPolicyNonComplianceReportRequestBuilder) {
    return i5d960953aeefdfe382eaaeee66b54c5138188e0487e4a969bc66342437b3efc3.NewGetPolicyNonComplianceReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ReportsRequestBuilder) GetPolicyNonComplianceSummaryReport()(*idf207024cf64a316ff9c7fa94616461d390fe5977d0687432c3309d78ba029ac.GetPolicyNonComplianceSummaryReportRequestBuilder) {
    return idf207024cf64a316ff9c7fa94616461d390fe5977d0687432c3309d78ba029ac.NewGetPolicyNonComplianceSummaryReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ReportsRequestBuilder) GetRelatedAppsStatusReport()(*iaa2bd508f982ae7d95dc62bbaefe9eb13e0c3fc9ad77d9389f8d82f3fb1dd151.GetRelatedAppsStatusReportRequestBuilder) {
    return iaa2bd508f982ae7d95dc62bbaefe9eb13e0c3fc9ad77d9389f8d82f3fb1dd151.NewGetRelatedAppsStatusReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ReportsRequestBuilder) GetReportFilters()(*iad468b4335f9e19b6476bd9d2cb864bb58150160cddf3825611adcbb6a864f65.GetReportFiltersRequestBuilder) {
    return iad468b4335f9e19b6476bd9d2cb864bb58150160cddf3825611adcbb6a864f65.NewGetReportFiltersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ReportsRequestBuilder) GetSettingNonComplianceReport()(*i8ad2db41d58439c4405d08911118b14826289f800859f884ca28da1387d489c2.GetSettingNonComplianceReportRequestBuilder) {
    return i8ad2db41d58439c4405d08911118b14826289f800859f884ca28da1387d489c2.NewGetSettingNonComplianceReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ReportsRequestBuilder) GetUnhealthyDefenderAgentsReport()(*ib57779f05872bd3be1d691d1b254dbe109c8ac7c77e7a265ef928aedbba3d32d.GetUnhealthyDefenderAgentsReportRequestBuilder) {
    return ib57779f05872bd3be1d691d1b254dbe109c8ac7c77e7a265ef928aedbba3d32d.NewGetUnhealthyDefenderAgentsReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ReportsRequestBuilder) GetUnhealthyFirewallReport()(*if4b9775de84c27e94434be0280aaa9a055f85419e5b5857a66bb9b9eed3280f4.GetUnhealthyFirewallReportRequestBuilder) {
    return if4b9775de84c27e94434be0280aaa9a055f85419e5b5857a66bb9b9eed3280f4.NewGetUnhealthyFirewallReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ReportsRequestBuilder) GetUnhealthyFirewallSummaryReport()(*idc692bbd177fe7be2a1d1da6b8436f4433611cc7abff49f069c1a2e0e027e466.GetUnhealthyFirewallSummaryReportRequestBuilder) {
    return idc692bbd177fe7be2a1d1da6b8436f4433611cc7abff49f069c1a2e0e027e466.NewGetUnhealthyFirewallSummaryReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ReportsRequestBuilder) GetUserInstallStatusReport()(*if876c79e09c00c1578bab3da9f4fcb0cc6d4ae3f9b2e792a2412130b488fcdc0.GetUserInstallStatusReportRequestBuilder) {
    return if876c79e09c00c1578bab3da9f4fcb0cc6d4ae3f9b2e792a2412130b488fcdc0.NewGetUserInstallStatusReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ReportsRequestBuilder) GetWindowsQualityUpdateAlertsPerPolicyPerDeviceReport()(*icef6de0280daff5ea2bb1854119e785ff95cb81abf14b3546d97c564e7d27123.GetWindowsQualityUpdateAlertsPerPolicyPerDeviceReportRequestBuilder) {
    return icef6de0280daff5ea2bb1854119e785ff95cb81abf14b3546d97c564e7d27123.NewGetWindowsQualityUpdateAlertsPerPolicyPerDeviceReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ReportsRequestBuilder) GetWindowsQualityUpdateAlertSummaryReport()(*if7f28186646455aa9e0df362fc86fe6afcf4c8f20b7f65732c9e4126cdc6c8e0.GetWindowsQualityUpdateAlertSummaryReportRequestBuilder) {
    return if7f28186646455aa9e0df362fc86fe6afcf4c8f20b7f65732c9e4126cdc6c8e0.NewGetWindowsQualityUpdateAlertSummaryReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ReportsRequestBuilder) GetWindowsUpdateAlertsPerPolicyPerDeviceReport()(*ieab9e3b30d90701370eeeb662f2e677617ef77e81a7f0a2a9d2c1d94fab9b995.GetWindowsUpdateAlertsPerPolicyPerDeviceReportRequestBuilder) {
    return ieab9e3b30d90701370eeeb662f2e677617ef77e81a7f0a2a9d2c1d94fab9b995.NewGetWindowsUpdateAlertsPerPolicyPerDeviceReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ReportsRequestBuilder) GetWindowsUpdateAlertSummaryReport()(*i95054c49fc60448e61ba13d5075f71dd6ae5b29fcc1f0307e8b5e38c80345642.GetWindowsUpdateAlertSummaryReportRequestBuilder) {
    return i95054c49fc60448e61ba13d5075f71dd6ae5b29fcc1f0307e8b5e38c80345642.NewGetWindowsUpdateAlertSummaryReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ReportsRequestBuilder) Patch(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceManagementReports, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(body, h, o);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, responseHandler)
    if err != nil {
        return err
    }
    return nil
}
func (m *ReportsRequestBuilder) ReportSchedules()(*iea8e9eceb6829c23694025f220d78170445e17ec05d7949319503a9dfef6f2dc.ReportSchedulesRequestBuilder) {
    return iea8e9eceb6829c23694025f220d78170445e17ec05d7949319503a9dfef6f2dc.NewReportSchedulesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ReportsRequestBuilder) ReportSchedulesById(id string)(*id6b331ae2f0cb2af60ba8b44c9282d7a2ff2e031c38004db15739c1402e76714.DeviceManagementReportScheduleRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["deviceManagementReportSchedule_id"] = id
    }
    return id6b331ae2f0cb2af60ba8b44c9282d7a2ff2e031c38004db15739c1402e76714.NewDeviceManagementReportScheduleRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
