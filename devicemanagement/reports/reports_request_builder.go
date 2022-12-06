package reports

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i08ed1bcbc2935e740bb1a75a82536d462055f7eb6c3b9f8c54f292796af881e9 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/reports/getmalwaresummaryreport"
    i0fcc2fe709581d2d74944b1f48febf27a4779c937c3a15442be80066edd36e63 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/reports/getgrouppolicysettingsdevicesettingsreport"
    i10798401aaaa8bec22bd9808818b05c618c65b4aa8de7cb7241d06e2eb24fd23 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/reports/getenrollmentconfigurationpoliciesbydevice"
    i17c7cf0accb5d03c972835746749cfe0eb895bec6658e210ece13d82213fd070 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/reports/getcompliancepoliciesreportfordevice"
    i1948951f1fadee39cec5840b7fb1388d50eceac2036ce74a8ece8724792a1636 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/reports/getcompliancepolicydevicesummaryreport"
    i1c7e958d1d2cd9f9a07df9fde55f97df047b41716fe43eeddc02616d301d9405 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/reports/getcompliancepolicydevicesreport"
    i1d42b2baf1a1b776afa9e69081dc91559e4cdd90109e8ba39f1c3cfd8c36b3b3 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/reports/getcachedreport"
    i246fa2e435cd7622048a0471ca5df979d86e83142f57f108d53701f4d8c17c18 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/reports/getconfigurationsettingdetailsreport"
    i2660e5353accb6f61768bb2883c08d64660fe2c8c67503cb521191b4c4e33e41 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/reports/getconfigurationpolicynoncompliancesummaryreport"
    i2a7e95d28bcab38f85f05fe3f131357d5c408f5fbd47fb213b7aac8dbd71328a "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/reports/cachedreportconfigurations"
    i2abf29284e813349f3083b2dc57f6553c2b467c8f68ad9c82e8950ce8d0afd2d "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/reports/getconfigurationsettingsreport"
    i2cbf3d93d72f9416ce29a9ac2774ddf6c3cceadcccbd462303a2ff5215c59622 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/reports/getpolicynoncompliancemetadata"
    i2d1fa8bbec4edb544879f637766d18e684b3ebd9bc8ab7dc1f62547a9cc9f6a0 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/reports/getcompliancesettingdetailsreport"
    i2fd8397e3f98434ca086fb32193c27fc2a226132ab5d78ecbbc4c3348c1e72fa "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/reports/getzebrafotadeploymentreport"
    i30c617cf06bedeb418c01e5099d3b64972ae3750fc00223f1abec08903667d2e "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/reports/getmobileapplicationmanagementappconfigurationreport"
    i3e013f91ac3d040684b27abcd0dc56839c52d2faa6d92ef0d2295a7fefa96a5b "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/reports/getappstatusoverviewreport"
    i3f11f18462beffdea839437cba5610e05b58902433c635e68fdfdfb6e5f1c4b7 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/reports/getconfigurationpolicydevicesummaryreport"
    i45e758acbf0a61d8e0b4ab34b380296a837b301cde6c8d13b415ce8dfb415464 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/reports/getencryptionreportfordevices"
    i45eba99273cfc2518cc265806b0cba99d31c259eee99db3fa9a93ca6fbd11d28 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/reports/getcompliancepolicynoncompliancereport"
    i460f7dd518397a3ddaedfec14fee80b46cfdb6c8bcad4e1a01d71973f2624fa1 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/reports/getcompliancesettingsreport"
    i4d73fdd41ba9bd99a3822503da6e6d8345d38f32a116b7860fb94fbdc8a9c79d "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/reports/getactivemalwarereport"
    i5284d7a78de859429101f202a37b95a2897b7a115a509f7ea9be6db8a70482a3 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/reports/getcertificatesreport"
    i5b82ca25c293ac194e35ceb34a2eb3523dd2a1568f2e61c9458b2d7fe4f1d970 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/reports/getcompliancepolicynoncompliancesummaryreport"
    i5d960953aeefdfe382eaaeee66b54c5138188e0487e4a969bc66342437b3efc3 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/reports/getpolicynoncompliancereport"
    i71062c76ab59540ed61b75db2e8378ccdfa322ba04cfe5afa641522cd26d9367 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/reports/getdeviceconfigurationpolicysettingssummaryreport"
    i711aceb383b63eb4ad8db059fe173d8b78d6bb1bdcfaf386735b20a8e54c2748 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/reports/getcompliancesettingnoncompliancereport"
    i71903f4f495c2ba2d1bec9d9d52c924a6bd43444cea991fb60e1874f824ed54e "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/reports/getfailedmobileappssummaryreport"
    i729c6211b1234e1904eb6e668f5a3a56598a9ecedbc4f360676e24e43bfaffdb "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/reports/getdeviceconfigurationpolicystatussummary"
    i75076cf3e3d6371cc0bb5d2810006d6e49e086cb781c7867b1093cfc1cf69827 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/reports/getappsinstallsummaryreport"
    i7611b2ea51c277728d8bca812a4642adb44f343f51f28de5c5e087ec400d0645 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/reports/getdevicemanagementintentsettingsreport"
    i788f436053d26e542373ad7c025d3a1a70768b96972d545d5aec67d0881993fb "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/reports/getquiettimepolicyusersummaryreport"
    i8ad2db41d58439c4405d08911118b14826289f800859f884ca28da1387d489c2 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/reports/getsettingnoncompliancereport"
    i95054c49fc60448e61ba13d5075f71dd6ae5b29fcc1f0307e8b5e38c80345642 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/reports/getwindowsupdatealertsummaryreport"
    i9819883d2f88a17d06e7a89c99985f254bca57d86a82a07526740f4057982566 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/reports/getconfigurationpolicynoncompliancereport"
    ia038b54cc9cc6b16a2ce646f288d86bbc4f28df40b41525d8397a84fa2a524ab "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/reports/getremoteassistancesessionsreport"
    ia7100ea2edcc786fe9054729a54d7600f750ebcc9be73a96446bf4b9bc605356 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/reports/getdeviceinstallstatusreport"
    iaa2bd508f982ae7d95dc62bbaefe9eb13e0c3fc9ad77d9389f8d82f3fb1dd151 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/reports/getrelatedappsstatusreport"
    iad468b4335f9e19b6476bd9d2cb864bb58150160cddf3825611adcbb6a864f65 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/reports/getreportfilters"
    iaf13e8679fcb6c2a8e3dd6221c05721a160c4ae39dc431f95147f9ceb536c688 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/reports/getdevicemanagementintentpersettingcontributingprofiles"
    ib57779f05872bd3be1d691d1b254dbe109c8ac7c77e7a265ef928aedbba3d32d "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/reports/getunhealthydefenderagentsreport"
    ib6b7e73b832065adf3e2ef8550d271d24c62fdb5128ea16eae3808e57561d401 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/reports/getallcertificatesreport"
    ib6f9fa86536428e22220c67df5e0bf322aa3124cb40d39ca6a96b97728e8551f "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/reports/gethistoricalreport"
    ibb77da97a1e31c9bb16ce4681732ad7475ae450918c13097176fab6ef841fcdf "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/reports/getconfigurationpolicysettingsdevicesummaryreport"
    ic029e50a08072e00092e7667a1fc68f437d1be340acce2c03ee7a004dea3e5f1 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/reports/getconfigurationpolicydevicesreport"
    ic09bfed61d7acf9bcd44dfdcb91600937d60cff331615b3b8a208a31e277cd7d "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/reports/getconfigmanagerdevicepolicystatusreport"
    ic72a13840aa7734c378a9cc2064090c5894fdd482f7d082fc5d0ebe8829cc479 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/reports/getquiettimepolicyusersreport"
    icd46e1255a1aee0bc5f9216ddd5fb29e365c5848ea3edecfd78d864191ca9748 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/reports/getdevicenoncompliancereport"
    icef6de0280daff5ea2bb1854119e785ff95cb81abf14b3546d97c564e7d27123 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/reports/getwindowsqualityupdatealertsperpolicyperdevicereport"
    id1df61af1e3874f797c91e9607c99dfd3cb6f0dbf128a90ab44536b912adf005 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/reports/getconfigurationpoliciesreportfordevice"
    id7b1135fd80c87263a81483b65064413c7251dd0137f48118d1471f835eefab4 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/reports/getfailedmobileappsreport"
    idb8bea4a603e0e6d6c5b6a880a8fa55f58151f4ad0fc9465122a375eec8aa927 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/reports/getmobileapplicationmanagementappregistrationsummaryreport"
    idc692bbd177fe7be2a1d1da6b8436f4433611cc7abff49f069c1a2e0e027e466 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/reports/getunhealthyfirewallsummaryreport"
    idf207024cf64a316ff9c7fa94616461d390fe5977d0687432c3309d78ba029ac "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/reports/getpolicynoncompliancesummaryreport"
    ie6deaed0072be9659d87a74fe95df383a715951d0aea4be5ba203ca890b3c041 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/reports/getactivemalwaresummaryreport"
    ieab9e3b30d90701370eeeb662f2e677617ef77e81a7f0a2a9d2c1d94fab9b995 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/reports/getwindowsupdatealertsperpolicyperdevicereport"
    if4b9775de84c27e94434be0280aaa9a055f85419e5b5857a66bb9b9eed3280f4 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/reports/getunhealthyfirewallreport"
    if7d41ece7783a25f185dc8071dfbd2363227c4fb6d8e08d2a9412f206c7279a9 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/reports/exportjobs"
    if7f28186646455aa9e0df362fc86fe6afcf4c8f20b7f65732c9e4126cdc6c8e0 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/reports/getwindowsqualityupdatealertsummaryreport"
    if876c79e09c00c1578bab3da9f4fcb0cc6d4ae3f9b2e792a2412130b488fcdc0 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/reports/getuserinstallstatusreport"
    if8e67aebead4149dc709bc45370408435c48c5beafd2a598a3552c78692904c1 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/reports/getconfigurationsettingnoncompliancereport"
    ifd91eed2e9e34bace676422bed143e114cf0500c85fd495dbb81c9d6c7393570 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/reports/getnoncompliantdevicesandsettingsreport"
    i2bf9ce7fb7bd0fa9fc75119c61d3de5e2e31a446f3ebba40eaf7f289544d9943 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/reports/cachedreportconfigurations/item"
    if607fcbf4c23ed3b14db77b88b46a29896a853e746acfffd003d215b0fb83a02 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/reports/exportjobs/item"
)

// ReportsRequestBuilder provides operations to manage the reports property of the microsoft.graph.deviceManagement entity.
type ReportsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ReportsRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ReportsRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ReportsRequestBuilderGetQueryParameters reports singleton
type ReportsRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ReportsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ReportsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ReportsRequestBuilderGetQueryParameters
}
// ReportsRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ReportsRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// CachedReportConfigurations provides operations to manage the cachedReportConfigurations property of the microsoft.graph.deviceManagementReports entity.
func (m *ReportsRequestBuilder) CachedReportConfigurations()(*i2a7e95d28bcab38f85f05fe3f131357d5c408f5fbd47fb213b7aac8dbd71328a.CachedReportConfigurationsRequestBuilder) {
    return i2a7e95d28bcab38f85f05fe3f131357d5c408f5fbd47fb213b7aac8dbd71328a.NewCachedReportConfigurationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CachedReportConfigurationsById provides operations to manage the cachedReportConfigurations property of the microsoft.graph.deviceManagementReports entity.
func (m *ReportsRequestBuilder) CachedReportConfigurationsById(id string)(*i2bf9ce7fb7bd0fa9fc75119c61d3de5e2e31a446f3ebba40eaf7f289544d9943.DeviceManagementCachedReportConfigurationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementCachedReportConfiguration%2Did"] = id
    }
    return i2bf9ce7fb7bd0fa9fc75119c61d3de5e2e31a446f3ebba40eaf7f289544d9943.NewDeviceManagementCachedReportConfigurationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewReportsRequestBuilderInternal instantiates a new ReportsRequestBuilder and sets the default values.
func NewReportsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ReportsRequestBuilder) {
    m := &ReportsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/reports{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewReportsRequestBuilder instantiates a new ReportsRequestBuilder and sets the default values.
func NewReportsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ReportsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewReportsRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property reports for deviceManagement
func (m *ReportsRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *ReportsRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateGetRequestInformation reports singleton
func (m *ReportsRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *ReportsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property reports in deviceManagement
func (m *ReportsRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementReportsable, requestConfiguration *ReportsRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete navigation property reports for deviceManagement
func (m *ReportsRequestBuilder) Delete(ctx context.Context, requestConfiguration *ReportsRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ExportJobs provides operations to manage the exportJobs property of the microsoft.graph.deviceManagementReports entity.
func (m *ReportsRequestBuilder) ExportJobs()(*if7d41ece7783a25f185dc8071dfbd2363227c4fb6d8e08d2a9412f206c7279a9.ExportJobsRequestBuilder) {
    return if7d41ece7783a25f185dc8071dfbd2363227c4fb6d8e08d2a9412f206c7279a9.NewExportJobsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExportJobsById provides operations to manage the exportJobs property of the microsoft.graph.deviceManagementReports entity.
func (m *ReportsRequestBuilder) ExportJobsById(id string)(*if607fcbf4c23ed3b14db77b88b46a29896a853e746acfffd003d215b0fb83a02.DeviceManagementExportJobItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementExportJob%2Did"] = id
    }
    return if607fcbf4c23ed3b14db77b88b46a29896a853e746acfffd003d215b0fb83a02.NewDeviceManagementExportJobItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get reports singleton
func (m *ReportsRequestBuilder) Get(ctx context.Context, requestConfiguration *ReportsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementReportsable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceManagementReportsFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementReportsable), nil
}
// GetActiveMalwareReport provides operations to call the getActiveMalwareReport method.
func (m *ReportsRequestBuilder) GetActiveMalwareReport()(*i4d73fdd41ba9bd99a3822503da6e6d8345d38f32a116b7860fb94fbdc8a9c79d.GetActiveMalwareReportRequestBuilder) {
    return i4d73fdd41ba9bd99a3822503da6e6d8345d38f32a116b7860fb94fbdc8a9c79d.NewGetActiveMalwareReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetActiveMalwareSummaryReport provides operations to call the getActiveMalwareSummaryReport method.
func (m *ReportsRequestBuilder) GetActiveMalwareSummaryReport()(*ie6deaed0072be9659d87a74fe95df383a715951d0aea4be5ba203ca890b3c041.GetActiveMalwareSummaryReportRequestBuilder) {
    return ie6deaed0072be9659d87a74fe95df383a715951d0aea4be5ba203ca890b3c041.NewGetActiveMalwareSummaryReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetAllCertificatesReport provides operations to call the getAllCertificatesReport method.
func (m *ReportsRequestBuilder) GetAllCertificatesReport()(*ib6b7e73b832065adf3e2ef8550d271d24c62fdb5128ea16eae3808e57561d401.GetAllCertificatesReportRequestBuilder) {
    return ib6b7e73b832065adf3e2ef8550d271d24c62fdb5128ea16eae3808e57561d401.NewGetAllCertificatesReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetAppsInstallSummaryReport provides operations to call the getAppsInstallSummaryReport method.
func (m *ReportsRequestBuilder) GetAppsInstallSummaryReport()(*i75076cf3e3d6371cc0bb5d2810006d6e49e086cb781c7867b1093cfc1cf69827.GetAppsInstallSummaryReportRequestBuilder) {
    return i75076cf3e3d6371cc0bb5d2810006d6e49e086cb781c7867b1093cfc1cf69827.NewGetAppsInstallSummaryReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetAppStatusOverviewReport provides operations to call the getAppStatusOverviewReport method.
func (m *ReportsRequestBuilder) GetAppStatusOverviewReport()(*i3e013f91ac3d040684b27abcd0dc56839c52d2faa6d92ef0d2295a7fefa96a5b.GetAppStatusOverviewReportRequestBuilder) {
    return i3e013f91ac3d040684b27abcd0dc56839c52d2faa6d92ef0d2295a7fefa96a5b.NewGetAppStatusOverviewReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetCachedReport provides operations to call the getCachedReport method.
func (m *ReportsRequestBuilder) GetCachedReport()(*i1d42b2baf1a1b776afa9e69081dc91559e4cdd90109e8ba39f1c3cfd8c36b3b3.GetCachedReportRequestBuilder) {
    return i1d42b2baf1a1b776afa9e69081dc91559e4cdd90109e8ba39f1c3cfd8c36b3b3.NewGetCachedReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetCertificatesReport provides operations to call the getCertificatesReport method.
func (m *ReportsRequestBuilder) GetCertificatesReport()(*i5284d7a78de859429101f202a37b95a2897b7a115a509f7ea9be6db8a70482a3.GetCertificatesReportRequestBuilder) {
    return i5284d7a78de859429101f202a37b95a2897b7a115a509f7ea9be6db8a70482a3.NewGetCertificatesReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetCompliancePoliciesReportForDevice provides operations to call the getCompliancePoliciesReportForDevice method.
func (m *ReportsRequestBuilder) GetCompliancePoliciesReportForDevice()(*i17c7cf0accb5d03c972835746749cfe0eb895bec6658e210ece13d82213fd070.GetCompliancePoliciesReportForDeviceRequestBuilder) {
    return i17c7cf0accb5d03c972835746749cfe0eb895bec6658e210ece13d82213fd070.NewGetCompliancePoliciesReportForDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetCompliancePolicyDevicesReport provides operations to call the getCompliancePolicyDevicesReport method.
func (m *ReportsRequestBuilder) GetCompliancePolicyDevicesReport()(*i1c7e958d1d2cd9f9a07df9fde55f97df047b41716fe43eeddc02616d301d9405.GetCompliancePolicyDevicesReportRequestBuilder) {
    return i1c7e958d1d2cd9f9a07df9fde55f97df047b41716fe43eeddc02616d301d9405.NewGetCompliancePolicyDevicesReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetCompliancePolicyDeviceSummaryReport provides operations to call the getCompliancePolicyDeviceSummaryReport method.
func (m *ReportsRequestBuilder) GetCompliancePolicyDeviceSummaryReport()(*i1948951f1fadee39cec5840b7fb1388d50eceac2036ce74a8ece8724792a1636.GetCompliancePolicyDeviceSummaryReportRequestBuilder) {
    return i1948951f1fadee39cec5840b7fb1388d50eceac2036ce74a8ece8724792a1636.NewGetCompliancePolicyDeviceSummaryReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetCompliancePolicyNonComplianceReport provides operations to call the getCompliancePolicyNonComplianceReport method.
func (m *ReportsRequestBuilder) GetCompliancePolicyNonComplianceReport()(*i45eba99273cfc2518cc265806b0cba99d31c259eee99db3fa9a93ca6fbd11d28.GetCompliancePolicyNonComplianceReportRequestBuilder) {
    return i45eba99273cfc2518cc265806b0cba99d31c259eee99db3fa9a93ca6fbd11d28.NewGetCompliancePolicyNonComplianceReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetCompliancePolicyNonComplianceSummaryReport provides operations to call the getCompliancePolicyNonComplianceSummaryReport method.
func (m *ReportsRequestBuilder) GetCompliancePolicyNonComplianceSummaryReport()(*i5b82ca25c293ac194e35ceb34a2eb3523dd2a1568f2e61c9458b2d7fe4f1d970.GetCompliancePolicyNonComplianceSummaryReportRequestBuilder) {
    return i5b82ca25c293ac194e35ceb34a2eb3523dd2a1568f2e61c9458b2d7fe4f1d970.NewGetCompliancePolicyNonComplianceSummaryReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetComplianceSettingDetailsReport provides operations to call the getComplianceSettingDetailsReport method.
func (m *ReportsRequestBuilder) GetComplianceSettingDetailsReport()(*i2d1fa8bbec4edb544879f637766d18e684b3ebd9bc8ab7dc1f62547a9cc9f6a0.GetComplianceSettingDetailsReportRequestBuilder) {
    return i2d1fa8bbec4edb544879f637766d18e684b3ebd9bc8ab7dc1f62547a9cc9f6a0.NewGetComplianceSettingDetailsReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetComplianceSettingNonComplianceReport provides operations to call the getComplianceSettingNonComplianceReport method.
func (m *ReportsRequestBuilder) GetComplianceSettingNonComplianceReport()(*i711aceb383b63eb4ad8db059fe173d8b78d6bb1bdcfaf386735b20a8e54c2748.GetComplianceSettingNonComplianceReportRequestBuilder) {
    return i711aceb383b63eb4ad8db059fe173d8b78d6bb1bdcfaf386735b20a8e54c2748.NewGetComplianceSettingNonComplianceReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetComplianceSettingsReport provides operations to call the getComplianceSettingsReport method.
func (m *ReportsRequestBuilder) GetComplianceSettingsReport()(*i460f7dd518397a3ddaedfec14fee80b46cfdb6c8bcad4e1a01d71973f2624fa1.GetComplianceSettingsReportRequestBuilder) {
    return i460f7dd518397a3ddaedfec14fee80b46cfdb6c8bcad4e1a01d71973f2624fa1.NewGetComplianceSettingsReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetConfigManagerDevicePolicyStatusReport provides operations to call the getConfigManagerDevicePolicyStatusReport method.
func (m *ReportsRequestBuilder) GetConfigManagerDevicePolicyStatusReport()(*ic09bfed61d7acf9bcd44dfdcb91600937d60cff331615b3b8a208a31e277cd7d.GetConfigManagerDevicePolicyStatusReportRequestBuilder) {
    return ic09bfed61d7acf9bcd44dfdcb91600937d60cff331615b3b8a208a31e277cd7d.NewGetConfigManagerDevicePolicyStatusReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetConfigurationPoliciesReportForDevice provides operations to call the getConfigurationPoliciesReportForDevice method.
func (m *ReportsRequestBuilder) GetConfigurationPoliciesReportForDevice()(*id1df61af1e3874f797c91e9607c99dfd3cb6f0dbf128a90ab44536b912adf005.GetConfigurationPoliciesReportForDeviceRequestBuilder) {
    return id1df61af1e3874f797c91e9607c99dfd3cb6f0dbf128a90ab44536b912adf005.NewGetConfigurationPoliciesReportForDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetConfigurationPolicyDevicesReport provides operations to call the getConfigurationPolicyDevicesReport method.
func (m *ReportsRequestBuilder) GetConfigurationPolicyDevicesReport()(*ic029e50a08072e00092e7667a1fc68f437d1be340acce2c03ee7a004dea3e5f1.GetConfigurationPolicyDevicesReportRequestBuilder) {
    return ic029e50a08072e00092e7667a1fc68f437d1be340acce2c03ee7a004dea3e5f1.NewGetConfigurationPolicyDevicesReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetConfigurationPolicyDeviceSummaryReport provides operations to call the getConfigurationPolicyDeviceSummaryReport method.
func (m *ReportsRequestBuilder) GetConfigurationPolicyDeviceSummaryReport()(*i3f11f18462beffdea839437cba5610e05b58902433c635e68fdfdfb6e5f1c4b7.GetConfigurationPolicyDeviceSummaryReportRequestBuilder) {
    return i3f11f18462beffdea839437cba5610e05b58902433c635e68fdfdfb6e5f1c4b7.NewGetConfigurationPolicyDeviceSummaryReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetConfigurationPolicyNonComplianceReport provides operations to call the getConfigurationPolicyNonComplianceReport method.
func (m *ReportsRequestBuilder) GetConfigurationPolicyNonComplianceReport()(*i9819883d2f88a17d06e7a89c99985f254bca57d86a82a07526740f4057982566.GetConfigurationPolicyNonComplianceReportRequestBuilder) {
    return i9819883d2f88a17d06e7a89c99985f254bca57d86a82a07526740f4057982566.NewGetConfigurationPolicyNonComplianceReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetConfigurationPolicyNonComplianceSummaryReport provides operations to call the getConfigurationPolicyNonComplianceSummaryReport method.
func (m *ReportsRequestBuilder) GetConfigurationPolicyNonComplianceSummaryReport()(*i2660e5353accb6f61768bb2883c08d64660fe2c8c67503cb521191b4c4e33e41.GetConfigurationPolicyNonComplianceSummaryReportRequestBuilder) {
    return i2660e5353accb6f61768bb2883c08d64660fe2c8c67503cb521191b4c4e33e41.NewGetConfigurationPolicyNonComplianceSummaryReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetConfigurationPolicySettingsDeviceSummaryReport provides operations to call the getConfigurationPolicySettingsDeviceSummaryReport method.
func (m *ReportsRequestBuilder) GetConfigurationPolicySettingsDeviceSummaryReport()(*ibb77da97a1e31c9bb16ce4681732ad7475ae450918c13097176fab6ef841fcdf.GetConfigurationPolicySettingsDeviceSummaryReportRequestBuilder) {
    return ibb77da97a1e31c9bb16ce4681732ad7475ae450918c13097176fab6ef841fcdf.NewGetConfigurationPolicySettingsDeviceSummaryReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetConfigurationSettingDetailsReport provides operations to call the getConfigurationSettingDetailsReport method.
func (m *ReportsRequestBuilder) GetConfigurationSettingDetailsReport()(*i246fa2e435cd7622048a0471ca5df979d86e83142f57f108d53701f4d8c17c18.GetConfigurationSettingDetailsReportRequestBuilder) {
    return i246fa2e435cd7622048a0471ca5df979d86e83142f57f108d53701f4d8c17c18.NewGetConfigurationSettingDetailsReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetConfigurationSettingNonComplianceReport provides operations to call the getConfigurationSettingNonComplianceReport method.
func (m *ReportsRequestBuilder) GetConfigurationSettingNonComplianceReport()(*if8e67aebead4149dc709bc45370408435c48c5beafd2a598a3552c78692904c1.GetConfigurationSettingNonComplianceReportRequestBuilder) {
    return if8e67aebead4149dc709bc45370408435c48c5beafd2a598a3552c78692904c1.NewGetConfigurationSettingNonComplianceReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetConfigurationSettingsReport provides operations to call the getConfigurationSettingsReport method.
func (m *ReportsRequestBuilder) GetConfigurationSettingsReport()(*i2abf29284e813349f3083b2dc57f6553c2b467c8f68ad9c82e8950ce8d0afd2d.GetConfigurationSettingsReportRequestBuilder) {
    return i2abf29284e813349f3083b2dc57f6553c2b467c8f68ad9c82e8950ce8d0afd2d.NewGetConfigurationSettingsReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetDeviceConfigurationPolicySettingsSummaryReport provides operations to call the getDeviceConfigurationPolicySettingsSummaryReport method.
func (m *ReportsRequestBuilder) GetDeviceConfigurationPolicySettingsSummaryReport()(*i71062c76ab59540ed61b75db2e8378ccdfa322ba04cfe5afa641522cd26d9367.GetDeviceConfigurationPolicySettingsSummaryReportRequestBuilder) {
    return i71062c76ab59540ed61b75db2e8378ccdfa322ba04cfe5afa641522cd26d9367.NewGetDeviceConfigurationPolicySettingsSummaryReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetDeviceConfigurationPolicyStatusSummary provides operations to call the getDeviceConfigurationPolicyStatusSummary method.
func (m *ReportsRequestBuilder) GetDeviceConfigurationPolicyStatusSummary()(*i729c6211b1234e1904eb6e668f5a3a56598a9ecedbc4f360676e24e43bfaffdb.GetDeviceConfigurationPolicyStatusSummaryRequestBuilder) {
    return i729c6211b1234e1904eb6e668f5a3a56598a9ecedbc4f360676e24e43bfaffdb.NewGetDeviceConfigurationPolicyStatusSummaryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetDeviceInstallStatusReport provides operations to call the getDeviceInstallStatusReport method.
func (m *ReportsRequestBuilder) GetDeviceInstallStatusReport()(*ia7100ea2edcc786fe9054729a54d7600f750ebcc9be73a96446bf4b9bc605356.GetDeviceInstallStatusReportRequestBuilder) {
    return ia7100ea2edcc786fe9054729a54d7600f750ebcc9be73a96446bf4b9bc605356.NewGetDeviceInstallStatusReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetDeviceManagementIntentPerSettingContributingProfiles provides operations to call the getDeviceManagementIntentPerSettingContributingProfiles method.
func (m *ReportsRequestBuilder) GetDeviceManagementIntentPerSettingContributingProfiles()(*iaf13e8679fcb6c2a8e3dd6221c05721a160c4ae39dc431f95147f9ceb536c688.GetDeviceManagementIntentPerSettingContributingProfilesRequestBuilder) {
    return iaf13e8679fcb6c2a8e3dd6221c05721a160c4ae39dc431f95147f9ceb536c688.NewGetDeviceManagementIntentPerSettingContributingProfilesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetDeviceManagementIntentSettingsReport provides operations to call the getDeviceManagementIntentSettingsReport method.
func (m *ReportsRequestBuilder) GetDeviceManagementIntentSettingsReport()(*i7611b2ea51c277728d8bca812a4642adb44f343f51f28de5c5e087ec400d0645.GetDeviceManagementIntentSettingsReportRequestBuilder) {
    return i7611b2ea51c277728d8bca812a4642adb44f343f51f28de5c5e087ec400d0645.NewGetDeviceManagementIntentSettingsReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetDeviceNonComplianceReport provides operations to call the getDeviceNonComplianceReport method.
func (m *ReportsRequestBuilder) GetDeviceNonComplianceReport()(*icd46e1255a1aee0bc5f9216ddd5fb29e365c5848ea3edecfd78d864191ca9748.GetDeviceNonComplianceReportRequestBuilder) {
    return icd46e1255a1aee0bc5f9216ddd5fb29e365c5848ea3edecfd78d864191ca9748.NewGetDeviceNonComplianceReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetEncryptionReportForDevices provides operations to call the getEncryptionReportForDevices method.
func (m *ReportsRequestBuilder) GetEncryptionReportForDevices()(*i45e758acbf0a61d8e0b4ab34b380296a837b301cde6c8d13b415ce8dfb415464.GetEncryptionReportForDevicesRequestBuilder) {
    return i45e758acbf0a61d8e0b4ab34b380296a837b301cde6c8d13b415ce8dfb415464.NewGetEncryptionReportForDevicesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetEnrollmentConfigurationPoliciesByDevice provides operations to call the getEnrollmentConfigurationPoliciesByDevice method.
func (m *ReportsRequestBuilder) GetEnrollmentConfigurationPoliciesByDevice()(*i10798401aaaa8bec22bd9808818b05c618c65b4aa8de7cb7241d06e2eb24fd23.GetEnrollmentConfigurationPoliciesByDeviceRequestBuilder) {
    return i10798401aaaa8bec22bd9808818b05c618c65b4aa8de7cb7241d06e2eb24fd23.NewGetEnrollmentConfigurationPoliciesByDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetFailedMobileAppsReport provides operations to call the getFailedMobileAppsReport method.
func (m *ReportsRequestBuilder) GetFailedMobileAppsReport()(*id7b1135fd80c87263a81483b65064413c7251dd0137f48118d1471f835eefab4.GetFailedMobileAppsReportRequestBuilder) {
    return id7b1135fd80c87263a81483b65064413c7251dd0137f48118d1471f835eefab4.NewGetFailedMobileAppsReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetFailedMobileAppsSummaryReport provides operations to call the getFailedMobileAppsSummaryReport method.
func (m *ReportsRequestBuilder) GetFailedMobileAppsSummaryReport()(*i71903f4f495c2ba2d1bec9d9d52c924a6bd43444cea991fb60e1874f824ed54e.GetFailedMobileAppsSummaryReportRequestBuilder) {
    return i71903f4f495c2ba2d1bec9d9d52c924a6bd43444cea991fb60e1874f824ed54e.NewGetFailedMobileAppsSummaryReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetGroupPolicySettingsDeviceSettingsReport provides operations to call the getGroupPolicySettingsDeviceSettingsReport method.
func (m *ReportsRequestBuilder) GetGroupPolicySettingsDeviceSettingsReport()(*i0fcc2fe709581d2d74944b1f48febf27a4779c937c3a15442be80066edd36e63.GetGroupPolicySettingsDeviceSettingsReportRequestBuilder) {
    return i0fcc2fe709581d2d74944b1f48febf27a4779c937c3a15442be80066edd36e63.NewGetGroupPolicySettingsDeviceSettingsReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetHistoricalReport provides operations to call the getHistoricalReport method.
func (m *ReportsRequestBuilder) GetHistoricalReport()(*ib6f9fa86536428e22220c67df5e0bf322aa3124cb40d39ca6a96b97728e8551f.GetHistoricalReportRequestBuilder) {
    return ib6f9fa86536428e22220c67df5e0bf322aa3124cb40d39ca6a96b97728e8551f.NewGetHistoricalReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMalwareSummaryReport provides operations to call the getMalwareSummaryReport method.
func (m *ReportsRequestBuilder) GetMalwareSummaryReport()(*i08ed1bcbc2935e740bb1a75a82536d462055f7eb6c3b9f8c54f292796af881e9.GetMalwareSummaryReportRequestBuilder) {
    return i08ed1bcbc2935e740bb1a75a82536d462055f7eb6c3b9f8c54f292796af881e9.NewGetMalwareSummaryReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMobileApplicationManagementAppConfigurationReport provides operations to call the getMobileApplicationManagementAppConfigurationReport method.
func (m *ReportsRequestBuilder) GetMobileApplicationManagementAppConfigurationReport()(*i30c617cf06bedeb418c01e5099d3b64972ae3750fc00223f1abec08903667d2e.GetMobileApplicationManagementAppConfigurationReportRequestBuilder) {
    return i30c617cf06bedeb418c01e5099d3b64972ae3750fc00223f1abec08903667d2e.NewGetMobileApplicationManagementAppConfigurationReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMobileApplicationManagementAppRegistrationSummaryReport provides operations to call the getMobileApplicationManagementAppRegistrationSummaryReport method.
func (m *ReportsRequestBuilder) GetMobileApplicationManagementAppRegistrationSummaryReport()(*idb8bea4a603e0e6d6c5b6a880a8fa55f58151f4ad0fc9465122a375eec8aa927.GetMobileApplicationManagementAppRegistrationSummaryReportRequestBuilder) {
    return idb8bea4a603e0e6d6c5b6a880a8fa55f58151f4ad0fc9465122a375eec8aa927.NewGetMobileApplicationManagementAppRegistrationSummaryReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetNoncompliantDevicesAndSettingsReport provides operations to call the getNoncompliantDevicesAndSettingsReport method.
func (m *ReportsRequestBuilder) GetNoncompliantDevicesAndSettingsReport()(*ifd91eed2e9e34bace676422bed143e114cf0500c85fd495dbb81c9d6c7393570.GetNoncompliantDevicesAndSettingsReportRequestBuilder) {
    return ifd91eed2e9e34bace676422bed143e114cf0500c85fd495dbb81c9d6c7393570.NewGetNoncompliantDevicesAndSettingsReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetPolicyNonComplianceMetadata provides operations to call the getPolicyNonComplianceMetadata method.
func (m *ReportsRequestBuilder) GetPolicyNonComplianceMetadata()(*i2cbf3d93d72f9416ce29a9ac2774ddf6c3cceadcccbd462303a2ff5215c59622.GetPolicyNonComplianceMetadataRequestBuilder) {
    return i2cbf3d93d72f9416ce29a9ac2774ddf6c3cceadcccbd462303a2ff5215c59622.NewGetPolicyNonComplianceMetadataRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetPolicyNonComplianceReport provides operations to call the getPolicyNonComplianceReport method.
func (m *ReportsRequestBuilder) GetPolicyNonComplianceReport()(*i5d960953aeefdfe382eaaeee66b54c5138188e0487e4a969bc66342437b3efc3.GetPolicyNonComplianceReportRequestBuilder) {
    return i5d960953aeefdfe382eaaeee66b54c5138188e0487e4a969bc66342437b3efc3.NewGetPolicyNonComplianceReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetPolicyNonComplianceSummaryReport provides operations to call the getPolicyNonComplianceSummaryReport method.
func (m *ReportsRequestBuilder) GetPolicyNonComplianceSummaryReport()(*idf207024cf64a316ff9c7fa94616461d390fe5977d0687432c3309d78ba029ac.GetPolicyNonComplianceSummaryReportRequestBuilder) {
    return idf207024cf64a316ff9c7fa94616461d390fe5977d0687432c3309d78ba029ac.NewGetPolicyNonComplianceSummaryReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetQuietTimePolicyUsersReport provides operations to call the getQuietTimePolicyUsersReport method.
func (m *ReportsRequestBuilder) GetQuietTimePolicyUsersReport()(*ic72a13840aa7734c378a9cc2064090c5894fdd482f7d082fc5d0ebe8829cc479.GetQuietTimePolicyUsersReportRequestBuilder) {
    return ic72a13840aa7734c378a9cc2064090c5894fdd482f7d082fc5d0ebe8829cc479.NewGetQuietTimePolicyUsersReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetQuietTimePolicyUserSummaryReport provides operations to call the getQuietTimePolicyUserSummaryReport method.
func (m *ReportsRequestBuilder) GetQuietTimePolicyUserSummaryReport()(*i788f436053d26e542373ad7c025d3a1a70768b96972d545d5aec67d0881993fb.GetQuietTimePolicyUserSummaryReportRequestBuilder) {
    return i788f436053d26e542373ad7c025d3a1a70768b96972d545d5aec67d0881993fb.NewGetQuietTimePolicyUserSummaryReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetRelatedAppsStatusReport provides operations to call the getRelatedAppsStatusReport method.
func (m *ReportsRequestBuilder) GetRelatedAppsStatusReport()(*iaa2bd508f982ae7d95dc62bbaefe9eb13e0c3fc9ad77d9389f8d82f3fb1dd151.GetRelatedAppsStatusReportRequestBuilder) {
    return iaa2bd508f982ae7d95dc62bbaefe9eb13e0c3fc9ad77d9389f8d82f3fb1dd151.NewGetRelatedAppsStatusReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetRemoteAssistanceSessionsReport provides operations to call the getRemoteAssistanceSessionsReport method.
func (m *ReportsRequestBuilder) GetRemoteAssistanceSessionsReport()(*ia038b54cc9cc6b16a2ce646f288d86bbc4f28df40b41525d8397a84fa2a524ab.GetRemoteAssistanceSessionsReportRequestBuilder) {
    return ia038b54cc9cc6b16a2ce646f288d86bbc4f28df40b41525d8397a84fa2a524ab.NewGetRemoteAssistanceSessionsReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetReportFilters provides operations to call the getReportFilters method.
func (m *ReportsRequestBuilder) GetReportFilters()(*iad468b4335f9e19b6476bd9d2cb864bb58150160cddf3825611adcbb6a864f65.GetReportFiltersRequestBuilder) {
    return iad468b4335f9e19b6476bd9d2cb864bb58150160cddf3825611adcbb6a864f65.NewGetReportFiltersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetSettingNonComplianceReport provides operations to call the getSettingNonComplianceReport method.
func (m *ReportsRequestBuilder) GetSettingNonComplianceReport()(*i8ad2db41d58439c4405d08911118b14826289f800859f884ca28da1387d489c2.GetSettingNonComplianceReportRequestBuilder) {
    return i8ad2db41d58439c4405d08911118b14826289f800859f884ca28da1387d489c2.NewGetSettingNonComplianceReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetUnhealthyDefenderAgentsReport provides operations to call the getUnhealthyDefenderAgentsReport method.
func (m *ReportsRequestBuilder) GetUnhealthyDefenderAgentsReport()(*ib57779f05872bd3be1d691d1b254dbe109c8ac7c77e7a265ef928aedbba3d32d.GetUnhealthyDefenderAgentsReportRequestBuilder) {
    return ib57779f05872bd3be1d691d1b254dbe109c8ac7c77e7a265ef928aedbba3d32d.NewGetUnhealthyDefenderAgentsReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetUnhealthyFirewallReport provides operations to call the getUnhealthyFirewallReport method.
func (m *ReportsRequestBuilder) GetUnhealthyFirewallReport()(*if4b9775de84c27e94434be0280aaa9a055f85419e5b5857a66bb9b9eed3280f4.GetUnhealthyFirewallReportRequestBuilder) {
    return if4b9775de84c27e94434be0280aaa9a055f85419e5b5857a66bb9b9eed3280f4.NewGetUnhealthyFirewallReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetUnhealthyFirewallSummaryReport provides operations to call the getUnhealthyFirewallSummaryReport method.
func (m *ReportsRequestBuilder) GetUnhealthyFirewallSummaryReport()(*idc692bbd177fe7be2a1d1da6b8436f4433611cc7abff49f069c1a2e0e027e466.GetUnhealthyFirewallSummaryReportRequestBuilder) {
    return idc692bbd177fe7be2a1d1da6b8436f4433611cc7abff49f069c1a2e0e027e466.NewGetUnhealthyFirewallSummaryReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetUserInstallStatusReport provides operations to call the getUserInstallStatusReport method.
func (m *ReportsRequestBuilder) GetUserInstallStatusReport()(*if876c79e09c00c1578bab3da9f4fcb0cc6d4ae3f9b2e792a2412130b488fcdc0.GetUserInstallStatusReportRequestBuilder) {
    return if876c79e09c00c1578bab3da9f4fcb0cc6d4ae3f9b2e792a2412130b488fcdc0.NewGetUserInstallStatusReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetWindowsQualityUpdateAlertsPerPolicyPerDeviceReport provides operations to call the getWindowsQualityUpdateAlertsPerPolicyPerDeviceReport method.
func (m *ReportsRequestBuilder) GetWindowsQualityUpdateAlertsPerPolicyPerDeviceReport()(*icef6de0280daff5ea2bb1854119e785ff95cb81abf14b3546d97c564e7d27123.GetWindowsQualityUpdateAlertsPerPolicyPerDeviceReportRequestBuilder) {
    return icef6de0280daff5ea2bb1854119e785ff95cb81abf14b3546d97c564e7d27123.NewGetWindowsQualityUpdateAlertsPerPolicyPerDeviceReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetWindowsQualityUpdateAlertSummaryReport provides operations to call the getWindowsQualityUpdateAlertSummaryReport method.
func (m *ReportsRequestBuilder) GetWindowsQualityUpdateAlertSummaryReport()(*if7f28186646455aa9e0df362fc86fe6afcf4c8f20b7f65732c9e4126cdc6c8e0.GetWindowsQualityUpdateAlertSummaryReportRequestBuilder) {
    return if7f28186646455aa9e0df362fc86fe6afcf4c8f20b7f65732c9e4126cdc6c8e0.NewGetWindowsQualityUpdateAlertSummaryReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetWindowsUpdateAlertsPerPolicyPerDeviceReport provides operations to call the getWindowsUpdateAlertsPerPolicyPerDeviceReport method.
func (m *ReportsRequestBuilder) GetWindowsUpdateAlertsPerPolicyPerDeviceReport()(*ieab9e3b30d90701370eeeb662f2e677617ef77e81a7f0a2a9d2c1d94fab9b995.GetWindowsUpdateAlertsPerPolicyPerDeviceReportRequestBuilder) {
    return ieab9e3b30d90701370eeeb662f2e677617ef77e81a7f0a2a9d2c1d94fab9b995.NewGetWindowsUpdateAlertsPerPolicyPerDeviceReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetWindowsUpdateAlertSummaryReport provides operations to call the getWindowsUpdateAlertSummaryReport method.
func (m *ReportsRequestBuilder) GetWindowsUpdateAlertSummaryReport()(*i95054c49fc60448e61ba13d5075f71dd6ae5b29fcc1f0307e8b5e38c80345642.GetWindowsUpdateAlertSummaryReportRequestBuilder) {
    return i95054c49fc60448e61ba13d5075f71dd6ae5b29fcc1f0307e8b5e38c80345642.NewGetWindowsUpdateAlertSummaryReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetZebraFotaDeploymentReport provides operations to call the getZebraFotaDeploymentReport method.
func (m *ReportsRequestBuilder) GetZebraFotaDeploymentReport()(*i2fd8397e3f98434ca086fb32193c27fc2a226132ab5d78ecbbc4c3348c1e72fa.GetZebraFotaDeploymentReportRequestBuilder) {
    return i2fd8397e3f98434ca086fb32193c27fc2a226132ab5d78ecbbc4c3348c1e72fa.NewGetZebraFotaDeploymentReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property reports in deviceManagement
func (m *ReportsRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementReportsable, requestConfiguration *ReportsRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementReportsable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceManagementReportsFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementReportsable), nil
}
