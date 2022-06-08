package reports

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i254c294d6b358b3854c3ee45f61d16c1237e0477952e917b034bbb0d8a6802b1 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/usercredentialusagedetails"
    i27c74777b0efbee67b16cbd84335b5228b74a547d6ffe5e338e3c1c6cc55c61b "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/monthlyprintusagesummariesbyprinter"
    i330d7d86bcb93d59ef14fc1257795533ce121643f4e8becd8884b57bc63e7b4b "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/monthlyprintusagebyprinter"
    i3772e364c43656cba18386bc364e131c7943e3a0b0461e3a162fb5eae513ac0b "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/dailyprintusagesummariesbyuser"
    i40dc3bbc1c58ea07c4c8b16942aaeca4da00eb3514241eef66edcd213ba084f7 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/monthlyprintusagesummariesbyuser"
    i46fd699154cb21987f377181c37a09acabdc0e13ef2aec22fc2f1352022502ba "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/dailyprintusagebyuser"
    i5ba6893546e6aaea0f4a67b2dd033135639babe968232c6fcc5163ab253c7847 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/dailyprintusagesummariesbyprinter"
    i6532f3214bfaa5d2aac58a9d9ffe42fb56e5e2242405e98165287c807365bcda "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/monthlyprintusagebyuser"
    ia74d4f7cd8fb70eb99b3c94cccde51ec24c099b4b64a46531b40bf30619e07c4 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/authenticationmethods"
    ibf41f160c65ef1194ee9e34f2d2b3bbf34a32e8ceb9026adff36a62f7067a2ac "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/dailyprintusagebyprinter"
    icddea0ed347960de17929a877880af4ac0bf296fb68f5c0e35eebd457fc0734f "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/applicationsignindetailedsummary"
    ice018542cd139d6164abff37b45f7b6a1d3055a473b3658086e14293b59be16e "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/credentialuserregistrationdetails"
    i1453ad3549b1a27e47913c2d4a29d5625c5b5e04deb3a2a6f264f5cd24a2e8da "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/dailyprintusagesummariesbyprinter/item"
    i1702ca3f0c998e29512738841125456c0b4e81357e5f7ca9f4176443002d28d5 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/monthlyprintusagebyuser/item"
    i2268eb94e47ea5d81b666c3310292b49733c18b0196efe2957b868503a9da7ff "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/monthlyprintusagesummariesbyuser/item"
    i22a3d34b6ea91c582f9826d88dc039aae3e5f8f62146be9df6713c324659e396 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/credentialuserregistrationdetails/item"
    i3808d92a196a8e86ccfe0455bb6aaf56bf5956b11460f112fbc1dbcd8375551d "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/dailyprintusagebyprinter/item"
    i4ca23ea1fafc22df57a5d1150889ab294f6d1582a79d145b961145b60dbf4c5b "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/dailyprintusagebyuser/item"
    i4f1c6062b781770b431d75b41deb0a2ac227468563ebac0623bef14f3efe351d "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/usercredentialusagedetails/item"
    i55a12b578425084e14afca840f9f8ad56ad8296a76d84139beb67b9619a01ae2 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/monthlyprintusagebyprinter/item"
    i6581ea85ca9063d302a537e4fdb5af97c20fd90311aa357f5482935e84bcd517 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/applicationsignindetailedsummary/item"
    ic1e71dc755b571267687f63b13312b758bc574de61144bfa50cd81af3b6d3626 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/monthlyprintusagesummariesbyprinter/item"
    iffede79819a5c0828784daecd91c1db9829729eb545a0c3b1948096142482a96 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/dailyprintusagesummariesbyuser/item"
)

// ReportsRequestBuilder provides operations to manage the reports property of the microsoft.graph.print entity.
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
// ReportsRequestBuilderGetQueryParameters get reports from print
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
// ApplicationSignInDetailedSummary the applicationSignInDetailedSummary property
func (m *ReportsRequestBuilder) ApplicationSignInDetailedSummary()(*icddea0ed347960de17929a877880af4ac0bf296fb68f5c0e35eebd457fc0734f.ApplicationSignInDetailedSummaryRequestBuilder) {
    return icddea0ed347960de17929a877880af4ac0bf296fb68f5c0e35eebd457fc0734f.NewApplicationSignInDetailedSummaryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ApplicationSignInDetailedSummaryById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.print.reports.applicationSignInDetailedSummary.item collection
func (m *ReportsRequestBuilder) ApplicationSignInDetailedSummaryById(id string)(*i6581ea85ca9063d302a537e4fdb5af97c20fd90311aa357f5482935e84bcd517.ApplicationSignInDetailedSummaryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["applicationSignInDetailedSummary%2Did"] = id
    }
    return i6581ea85ca9063d302a537e4fdb5af97c20fd90311aa357f5482935e84bcd517.NewApplicationSignInDetailedSummaryItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// AuthenticationMethods the authenticationMethods property
func (m *ReportsRequestBuilder) AuthenticationMethods()(*ia74d4f7cd8fb70eb99b3c94cccde51ec24c099b4b64a46531b40bf30619e07c4.AuthenticationMethodsRequestBuilder) {
    return ia74d4f7cd8fb70eb99b3c94cccde51ec24c099b4b64a46531b40bf30619e07c4.NewAuthenticationMethodsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewReportsRequestBuilderInternal instantiates a new ReportsRequestBuilder and sets the default values.
func NewReportsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ReportsRequestBuilder) {
    m := &ReportsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/print/reports{?%24select,%24expand}";
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
// CreateDeleteRequestInformation delete navigation property reports for print
func (m *ReportsRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property reports for print
func (m *ReportsRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration *ReportsRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation get reports from print
func (m *ReportsRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration get reports from print
func (m *ReportsRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *ReportsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property reports in print
func (m *ReportsRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ReportRootable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property reports in print
func (m *ReportsRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ReportRootable, requestConfiguration *ReportsRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CredentialUserRegistrationDetails the credentialUserRegistrationDetails property
func (m *ReportsRequestBuilder) CredentialUserRegistrationDetails()(*ice018542cd139d6164abff37b45f7b6a1d3055a473b3658086e14293b59be16e.CredentialUserRegistrationDetailsRequestBuilder) {
    return ice018542cd139d6164abff37b45f7b6a1d3055a473b3658086e14293b59be16e.NewCredentialUserRegistrationDetailsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CredentialUserRegistrationDetailsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.print.reports.credentialUserRegistrationDetails.item collection
func (m *ReportsRequestBuilder) CredentialUserRegistrationDetailsById(id string)(*i22a3d34b6ea91c582f9826d88dc039aae3e5f8f62146be9df6713c324659e396.CredentialUserRegistrationDetailsItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["credentialUserRegistrationDetails%2Did"] = id
    }
    return i22a3d34b6ea91c582f9826d88dc039aae3e5f8f62146be9df6713c324659e396.NewCredentialUserRegistrationDetailsItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// DailyPrintUsageByPrinter the dailyPrintUsageByPrinter property
func (m *ReportsRequestBuilder) DailyPrintUsageByPrinter()(*ibf41f160c65ef1194ee9e34f2d2b3bbf34a32e8ceb9026adff36a62f7067a2ac.DailyPrintUsageByPrinterRequestBuilder) {
    return ibf41f160c65ef1194ee9e34f2d2b3bbf34a32e8ceb9026adff36a62f7067a2ac.NewDailyPrintUsageByPrinterRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DailyPrintUsageByPrinterById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.print.reports.dailyPrintUsageByPrinter.item collection
func (m *ReportsRequestBuilder) DailyPrintUsageByPrinterById(id string)(*i3808d92a196a8e86ccfe0455bb6aaf56bf5956b11460f112fbc1dbcd8375551d.PrintUsageByPrinterItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["printUsageByPrinter%2Did"] = id
    }
    return i3808d92a196a8e86ccfe0455bb6aaf56bf5956b11460f112fbc1dbcd8375551d.NewPrintUsageByPrinterItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// DailyPrintUsageByUser the dailyPrintUsageByUser property
func (m *ReportsRequestBuilder) DailyPrintUsageByUser()(*i46fd699154cb21987f377181c37a09acabdc0e13ef2aec22fc2f1352022502ba.DailyPrintUsageByUserRequestBuilder) {
    return i46fd699154cb21987f377181c37a09acabdc0e13ef2aec22fc2f1352022502ba.NewDailyPrintUsageByUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DailyPrintUsageByUserById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.print.reports.dailyPrintUsageByUser.item collection
func (m *ReportsRequestBuilder) DailyPrintUsageByUserById(id string)(*i4ca23ea1fafc22df57a5d1150889ab294f6d1582a79d145b961145b60dbf4c5b.PrintUsageByUserItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["printUsageByUser%2Did"] = id
    }
    return i4ca23ea1fafc22df57a5d1150889ab294f6d1582a79d145b961145b60dbf4c5b.NewPrintUsageByUserItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// DailyPrintUsageSummariesByPrinter the dailyPrintUsageSummariesByPrinter property
func (m *ReportsRequestBuilder) DailyPrintUsageSummariesByPrinter()(*i5ba6893546e6aaea0f4a67b2dd033135639babe968232c6fcc5163ab253c7847.DailyPrintUsageSummariesByPrinterRequestBuilder) {
    return i5ba6893546e6aaea0f4a67b2dd033135639babe968232c6fcc5163ab253c7847.NewDailyPrintUsageSummariesByPrinterRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DailyPrintUsageSummariesByPrinterById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.print.reports.dailyPrintUsageSummariesByPrinter.item collection
func (m *ReportsRequestBuilder) DailyPrintUsageSummariesByPrinterById(id string)(*i1453ad3549b1a27e47913c2d4a29d5625c5b5e04deb3a2a6f264f5cd24a2e8da.PrintUsageByPrinterItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["printUsageByPrinter%2Did"] = id
    }
    return i1453ad3549b1a27e47913c2d4a29d5625c5b5e04deb3a2a6f264f5cd24a2e8da.NewPrintUsageByPrinterItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// DailyPrintUsageSummariesByUser the dailyPrintUsageSummariesByUser property
func (m *ReportsRequestBuilder) DailyPrintUsageSummariesByUser()(*i3772e364c43656cba18386bc364e131c7943e3a0b0461e3a162fb5eae513ac0b.DailyPrintUsageSummariesByUserRequestBuilder) {
    return i3772e364c43656cba18386bc364e131c7943e3a0b0461e3a162fb5eae513ac0b.NewDailyPrintUsageSummariesByUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DailyPrintUsageSummariesByUserById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.print.reports.dailyPrintUsageSummariesByUser.item collection
func (m *ReportsRequestBuilder) DailyPrintUsageSummariesByUserById(id string)(*iffede79819a5c0828784daecd91c1db9829729eb545a0c3b1948096142482a96.PrintUsageByUserItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["printUsageByUser%2Did"] = id
    }
    return iffede79819a5c0828784daecd91c1db9829729eb545a0c3b1948096142482a96.NewPrintUsageByUserItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Delete delete navigation property reports for print
func (m *ReportsRequestBuilder) Delete()(error) {
    return m.DeleteWithRequestConfigurationAndResponseHandler(nil, nil);
}
// DeleteWithRequestConfigurationAndResponseHandler delete navigation property reports for print
func (m *ReportsRequestBuilder) DeleteWithRequestConfigurationAndResponseHandler(requestConfiguration *ReportsRequestBuilderDeleteRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
    requestInfo, err := m.CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, responseHandler, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get get reports from print
func (m *ReportsRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ReportRootable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler get reports from print
func (m *ReportsRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *ReportsRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ReportRootable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateReportRootFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ReportRootable), nil
}
// MonthlyPrintUsageByPrinter the monthlyPrintUsageByPrinter property
func (m *ReportsRequestBuilder) MonthlyPrintUsageByPrinter()(*i330d7d86bcb93d59ef14fc1257795533ce121643f4e8becd8884b57bc63e7b4b.MonthlyPrintUsageByPrinterRequestBuilder) {
    return i330d7d86bcb93d59ef14fc1257795533ce121643f4e8becd8884b57bc63e7b4b.NewMonthlyPrintUsageByPrinterRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MonthlyPrintUsageByPrinterById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.print.reports.monthlyPrintUsageByPrinter.item collection
func (m *ReportsRequestBuilder) MonthlyPrintUsageByPrinterById(id string)(*i55a12b578425084e14afca840f9f8ad56ad8296a76d84139beb67b9619a01ae2.PrintUsageByPrinterItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["printUsageByPrinter%2Did"] = id
    }
    return i55a12b578425084e14afca840f9f8ad56ad8296a76d84139beb67b9619a01ae2.NewPrintUsageByPrinterItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// MonthlyPrintUsageByUser the monthlyPrintUsageByUser property
func (m *ReportsRequestBuilder) MonthlyPrintUsageByUser()(*i6532f3214bfaa5d2aac58a9d9ffe42fb56e5e2242405e98165287c807365bcda.MonthlyPrintUsageByUserRequestBuilder) {
    return i6532f3214bfaa5d2aac58a9d9ffe42fb56e5e2242405e98165287c807365bcda.NewMonthlyPrintUsageByUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MonthlyPrintUsageByUserById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.print.reports.monthlyPrintUsageByUser.item collection
func (m *ReportsRequestBuilder) MonthlyPrintUsageByUserById(id string)(*i1702ca3f0c998e29512738841125456c0b4e81357e5f7ca9f4176443002d28d5.PrintUsageByUserItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["printUsageByUser%2Did"] = id
    }
    return i1702ca3f0c998e29512738841125456c0b4e81357e5f7ca9f4176443002d28d5.NewPrintUsageByUserItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// MonthlyPrintUsageSummariesByPrinter the monthlyPrintUsageSummariesByPrinter property
func (m *ReportsRequestBuilder) MonthlyPrintUsageSummariesByPrinter()(*i27c74777b0efbee67b16cbd84335b5228b74a547d6ffe5e338e3c1c6cc55c61b.MonthlyPrintUsageSummariesByPrinterRequestBuilder) {
    return i27c74777b0efbee67b16cbd84335b5228b74a547d6ffe5e338e3c1c6cc55c61b.NewMonthlyPrintUsageSummariesByPrinterRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MonthlyPrintUsageSummariesByPrinterById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.print.reports.monthlyPrintUsageSummariesByPrinter.item collection
func (m *ReportsRequestBuilder) MonthlyPrintUsageSummariesByPrinterById(id string)(*ic1e71dc755b571267687f63b13312b758bc574de61144bfa50cd81af3b6d3626.PrintUsageByPrinterItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["printUsageByPrinter%2Did"] = id
    }
    return ic1e71dc755b571267687f63b13312b758bc574de61144bfa50cd81af3b6d3626.NewPrintUsageByPrinterItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// MonthlyPrintUsageSummariesByUser the monthlyPrintUsageSummariesByUser property
func (m *ReportsRequestBuilder) MonthlyPrintUsageSummariesByUser()(*i40dc3bbc1c58ea07c4c8b16942aaeca4da00eb3514241eef66edcd213ba084f7.MonthlyPrintUsageSummariesByUserRequestBuilder) {
    return i40dc3bbc1c58ea07c4c8b16942aaeca4da00eb3514241eef66edcd213ba084f7.NewMonthlyPrintUsageSummariesByUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MonthlyPrintUsageSummariesByUserById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.print.reports.monthlyPrintUsageSummariesByUser.item collection
func (m *ReportsRequestBuilder) MonthlyPrintUsageSummariesByUserById(id string)(*i2268eb94e47ea5d81b666c3310292b49733c18b0196efe2957b868503a9da7ff.PrintUsageByUserItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["printUsageByUser%2Did"] = id
    }
    return i2268eb94e47ea5d81b666c3310292b49733c18b0196efe2957b868503a9da7ff.NewPrintUsageByUserItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property reports in print
func (m *ReportsRequestBuilder) Patch(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ReportRootable)(error) {
    return m.PatchWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PatchWithRequestConfigurationAndResponseHandler update the navigation property reports in print
func (m *ReportsRequestBuilder) PatchWithRequestConfigurationAndResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ReportRootable, requestConfiguration *ReportsRequestBuilderPatchRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
    requestInfo, err := m.CreatePatchRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, responseHandler, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// UserCredentialUsageDetails the userCredentialUsageDetails property
func (m *ReportsRequestBuilder) UserCredentialUsageDetails()(*i254c294d6b358b3854c3ee45f61d16c1237e0477952e917b034bbb0d8a6802b1.UserCredentialUsageDetailsRequestBuilder) {
    return i254c294d6b358b3854c3ee45f61d16c1237e0477952e917b034bbb0d8a6802b1.NewUserCredentialUsageDetailsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserCredentialUsageDetailsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.print.reports.userCredentialUsageDetails.item collection
func (m *ReportsRequestBuilder) UserCredentialUsageDetailsById(id string)(*i4f1c6062b781770b431d75b41deb0a2ac227468563ebac0623bef14f3efe351d.UserCredentialUsageDetailsItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userCredentialUsageDetails%2Did"] = id
    }
    return i4f1c6062b781770b431d75b41deb0a2ac227468563ebac0623bef14f3efe351d.NewUserCredentialUsageDetailsItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
