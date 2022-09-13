package monitoring

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i2edb12705e6a63a8a0fb3f8c7a11f4ab12f4be764e61fa1094f401595fb171bf "github.com/microsoftgraph/msgraph-beta-sdk-go/models/devicemanagement"
    idba65f28d93adc7778f8fe81d4f189ab183b31b1afdd469eaefc6bc32666e664 "github.com/microsoftgraph/msgraph-beta-sdk-go/monitoring/alertrecords"
    ie61fb3019adb9dd0420256aa6aacdb175005e870f509b997fb24e6360e0ca219 "github.com/microsoftgraph/msgraph-beta-sdk-go/monitoring/alertrules"
    i4a397c92150e0353279b4ef4eeaeb5a4ab922ad0556d24d80908e100407dba3e "github.com/microsoftgraph/msgraph-beta-sdk-go/monitoring/alertrecords/item"
    if486ea48fbeed75043bbdfabe527238f4bb23a0777e0d5085924090762d9628b "github.com/microsoftgraph/msgraph-beta-sdk-go/monitoring/alertrules/item"
)

// MonitoringRequestBuilder provides operations to manage the monitoring singleton.
type MonitoringRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// MonitoringRequestBuilderGetQueryParameters get monitoring
type MonitoringRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// MonitoringRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MonitoringRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MonitoringRequestBuilderGetQueryParameters
}
// MonitoringRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MonitoringRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AlertRecords the alertRecords property
func (m *MonitoringRequestBuilder) AlertRecords()(*idba65f28d93adc7778f8fe81d4f189ab183b31b1afdd469eaefc6bc32666e664.AlertRecordsRequestBuilder) {
    return idba65f28d93adc7778f8fe81d4f189ab183b31b1afdd469eaefc6bc32666e664.NewAlertRecordsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AlertRecordsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.monitoring.alertRecords.item collection
func (m *MonitoringRequestBuilder) AlertRecordsById(id string)(*i4a397c92150e0353279b4ef4eeaeb5a4ab922ad0556d24d80908e100407dba3e.AlertRecordItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["alertRecord%2Did"] = id
    }
    return i4a397c92150e0353279b4ef4eeaeb5a4ab922ad0556d24d80908e100407dba3e.NewAlertRecordItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// AlertRules the alertRules property
func (m *MonitoringRequestBuilder) AlertRules()(*ie61fb3019adb9dd0420256aa6aacdb175005e870f509b997fb24e6360e0ca219.AlertRulesRequestBuilder) {
    return ie61fb3019adb9dd0420256aa6aacdb175005e870f509b997fb24e6360e0ca219.NewAlertRulesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AlertRulesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.monitoring.alertRules.item collection
func (m *MonitoringRequestBuilder) AlertRulesById(id string)(*if486ea48fbeed75043bbdfabe527238f4bb23a0777e0d5085924090762d9628b.AlertRuleItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["alertRule%2Did"] = id
    }
    return if486ea48fbeed75043bbdfabe527238f4bb23a0777e0d5085924090762d9628b.NewAlertRuleItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewMonitoringRequestBuilderInternal instantiates a new MonitoringRequestBuilder and sets the default values.
func NewMonitoringRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MonitoringRequestBuilder) {
    m := &MonitoringRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/monitoring{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewMonitoringRequestBuilder instantiates a new MonitoringRequestBuilder and sets the default values.
func NewMonitoringRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MonitoringRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMonitoringRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation get monitoring
func (m *MonitoringRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration get monitoring
func (m *MonitoringRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *MonitoringRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update monitoring
func (m *MonitoringRequestBuilder) CreatePatchRequestInformation(body i2edb12705e6a63a8a0fb3f8c7a11f4ab12f4be764e61fa1094f401595fb171bf.Monitoringable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update monitoring
func (m *MonitoringRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body i2edb12705e6a63a8a0fb3f8c7a11f4ab12f4be764e61fa1094f401595fb171bf.Monitoringable, requestConfiguration *MonitoringRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Get get monitoring
func (m *MonitoringRequestBuilder) Get(ctx context.Context, requestConfiguration *MonitoringRequestBuilderGetRequestConfiguration)(i2edb12705e6a63a8a0fb3f8c7a11f4ab12f4be764e61fa1094f401595fb171bf.Monitoringable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, i2edb12705e6a63a8a0fb3f8c7a11f4ab12f4be764e61fa1094f401595fb171bf.CreateMonitoringFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i2edb12705e6a63a8a0fb3f8c7a11f4ab12f4be764e61fa1094f401595fb171bf.Monitoringable), nil
}
// Patch update monitoring
func (m *MonitoringRequestBuilder) Patch(ctx context.Context, body i2edb12705e6a63a8a0fb3f8c7a11f4ab12f4be764e61fa1094f401595fb171bf.Monitoringable, requestConfiguration *MonitoringRequestBuilderPatchRequestConfiguration)(error) {
    requestInfo, err := m.CreatePatchRequestInformationWithRequestConfiguration(body, requestConfiguration);
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
