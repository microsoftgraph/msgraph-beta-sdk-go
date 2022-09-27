package monitoring

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i2edb12705e6a63a8a0fb3f8c7a11f4ab12f4be764e61fa1094f401595fb171bf "github.com/microsoftgraph/msgraph-beta-sdk-go/models/devicemanagement"
    ic3d08d978344fb190fcda69c1207e348d81a12d1378f59743c4eb0a373d59241 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/monitoring/alertrecords"
    ifc74b0db7564744a90ac3bf936445eee7e085edd8a741c1b0e64d66141e0002e "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/monitoring/alertrules"
    i02f8f62eebb75f1ad1675a31e20ad38f03048f2bfb47f3cc693479b68046a5bd "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/monitoring/alertrecords/item"
    i7321368ae7ef7cee965f460ecaba8d359d9ed7610180962fe193dcf9275a0564 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/monitoring/alertrules/item"
)

// MonitoringRequestBuilder provides operations to manage the monitoring property of the microsoft.graph.deviceManagement entity.
type MonitoringRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// MonitoringRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MonitoringRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// MonitoringRequestBuilderGetQueryParameters get monitoring from deviceManagement
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
func (m *MonitoringRequestBuilder) AlertRecords()(*ic3d08d978344fb190fcda69c1207e348d81a12d1378f59743c4eb0a373d59241.AlertRecordsRequestBuilder) {
    return ic3d08d978344fb190fcda69c1207e348d81a12d1378f59743c4eb0a373d59241.NewAlertRecordsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AlertRecordsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.monitoring.alertRecords.item collection
func (m *MonitoringRequestBuilder) AlertRecordsById(id string)(*i02f8f62eebb75f1ad1675a31e20ad38f03048f2bfb47f3cc693479b68046a5bd.AlertRecordItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["alertRecord%2Did"] = id
    }
    return i02f8f62eebb75f1ad1675a31e20ad38f03048f2bfb47f3cc693479b68046a5bd.NewAlertRecordItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// AlertRules the alertRules property
func (m *MonitoringRequestBuilder) AlertRules()(*ifc74b0db7564744a90ac3bf936445eee7e085edd8a741c1b0e64d66141e0002e.AlertRulesRequestBuilder) {
    return ifc74b0db7564744a90ac3bf936445eee7e085edd8a741c1b0e64d66141e0002e.NewAlertRulesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AlertRulesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.monitoring.alertRules.item collection
func (m *MonitoringRequestBuilder) AlertRulesById(id string)(*i7321368ae7ef7cee965f460ecaba8d359d9ed7610180962fe193dcf9275a0564.AlertRuleItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["alertRule%2Did"] = id
    }
    return i7321368ae7ef7cee965f460ecaba8d359d9ed7610180962fe193dcf9275a0564.NewAlertRuleItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewMonitoringRequestBuilderInternal instantiates a new MonitoringRequestBuilder and sets the default values.
func NewMonitoringRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MonitoringRequestBuilder) {
    m := &MonitoringRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/monitoring{?%24select,%24expand}";
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
// CreateDeleteRequestInformation delete navigation property monitoring for deviceManagement
func (m *MonitoringRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property monitoring for deviceManagement
func (m *MonitoringRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration *MonitoringRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation get monitoring from deviceManagement
func (m *MonitoringRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration get monitoring from deviceManagement
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
// CreatePatchRequestInformation update the navigation property monitoring in deviceManagement
func (m *MonitoringRequestBuilder) CreatePatchRequestInformation(body i2edb12705e6a63a8a0fb3f8c7a11f4ab12f4be764e61fa1094f401595fb171bf.Monitoringable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property monitoring in deviceManagement
func (m *MonitoringRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body i2edb12705e6a63a8a0fb3f8c7a11f4ab12f4be764e61fa1094f401595fb171bf.Monitoringable, requestConfiguration *MonitoringRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete navigation property monitoring for deviceManagement
func (m *MonitoringRequestBuilder) Delete(ctx context.Context, requestConfiguration *MonitoringRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration);
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
// Get get monitoring from deviceManagement
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
// Patch update the navigation property monitoring in deviceManagement
func (m *MonitoringRequestBuilder) Patch(ctx context.Context, body i2edb12705e6a63a8a0fb3f8c7a11f4ab12f4be764e61fa1094f401595fb171bf.Monitoringable, requestConfiguration *MonitoringRequestBuilderPatchRequestConfiguration)(i2edb12705e6a63a8a0fb3f8c7a11f4ab12f4be764e61fa1094f401595fb171bf.Monitoringable, error) {
    requestInfo, err := m.CreatePatchRequestInformationWithRequestConfiguration(body, requestConfiguration);
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
