package security

import (
    "context"
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DataDiscoveryCloudAppDiscoveryUploadedStreamsItemMicrosoftGraphSecurityAggregatedAppsDetailsWithPeriodRequestBuilder provides operations to call the aggregatedAppsDetails method.
type DataDiscoveryCloudAppDiscoveryUploadedStreamsItemMicrosoftGraphSecurityAggregatedAppsDetailsWithPeriodRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DataDiscoveryCloudAppDiscoveryUploadedStreamsItemMicrosoftGraphSecurityAggregatedAppsDetailsWithPeriodRequestBuilderGetQueryParameters invoke function aggregatedAppsDetails
type DataDiscoveryCloudAppDiscoveryUploadedStreamsItemMicrosoftGraphSecurityAggregatedAppsDetailsWithPeriodRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// DataDiscoveryCloudAppDiscoveryUploadedStreamsItemMicrosoftGraphSecurityAggregatedAppsDetailsWithPeriodRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DataDiscoveryCloudAppDiscoveryUploadedStreamsItemMicrosoftGraphSecurityAggregatedAppsDetailsWithPeriodRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DataDiscoveryCloudAppDiscoveryUploadedStreamsItemMicrosoftGraphSecurityAggregatedAppsDetailsWithPeriodRequestBuilderGetQueryParameters
}
// NewDataDiscoveryCloudAppDiscoveryUploadedStreamsItemMicrosoftGraphSecurityAggregatedAppsDetailsWithPeriodRequestBuilderInternal instantiates a new DataDiscoveryCloudAppDiscoveryUploadedStreamsItemMicrosoftGraphSecurityAggregatedAppsDetailsWithPeriodRequestBuilder and sets the default values.
func NewDataDiscoveryCloudAppDiscoveryUploadedStreamsItemMicrosoftGraphSecurityAggregatedAppsDetailsWithPeriodRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, period *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)(*DataDiscoveryCloudAppDiscoveryUploadedStreamsItemMicrosoftGraphSecurityAggregatedAppsDetailsWithPeriodRequestBuilder) {
    m := &DataDiscoveryCloudAppDiscoveryUploadedStreamsItemMicrosoftGraphSecurityAggregatedAppsDetailsWithPeriodRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/dataDiscovery/cloudAppDiscovery/uploadedStreams/{cloudAppDiscoveryReport%2Did}/microsoft.graph.security.aggregatedAppsDetails(period={period}){?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    if period != nil {
        m.BaseRequestBuilder.PathParameters["period"] = (*period).String()
    }
    return m
}
// NewDataDiscoveryCloudAppDiscoveryUploadedStreamsItemMicrosoftGraphSecurityAggregatedAppsDetailsWithPeriodRequestBuilder instantiates a new DataDiscoveryCloudAppDiscoveryUploadedStreamsItemMicrosoftGraphSecurityAggregatedAppsDetailsWithPeriodRequestBuilder and sets the default values.
func NewDataDiscoveryCloudAppDiscoveryUploadedStreamsItemMicrosoftGraphSecurityAggregatedAppsDetailsWithPeriodRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DataDiscoveryCloudAppDiscoveryUploadedStreamsItemMicrosoftGraphSecurityAggregatedAppsDetailsWithPeriodRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDataDiscoveryCloudAppDiscoveryUploadedStreamsItemMicrosoftGraphSecurityAggregatedAppsDetailsWithPeriodRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get invoke function aggregatedAppsDetails
// Deprecated: This method is obsolete. Use GetAsAggregatedAppsDetailsWithPeriodGetResponse instead.
// returns a DataDiscoveryCloudAppDiscoveryUploadedStreamsItemMicrosoftGraphSecurityAggregatedAppsDetailsWithPeriodAggregatedAppsDetailsWithPeriodResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DataDiscoveryCloudAppDiscoveryUploadedStreamsItemMicrosoftGraphSecurityAggregatedAppsDetailsWithPeriodRequestBuilder) Get(ctx context.Context, requestConfiguration *DataDiscoveryCloudAppDiscoveryUploadedStreamsItemMicrosoftGraphSecurityAggregatedAppsDetailsWithPeriodRequestBuilderGetRequestConfiguration)(DataDiscoveryCloudAppDiscoveryUploadedStreamsItemMicrosoftGraphSecurityAggregatedAppsDetailsWithPeriodAggregatedAppsDetailsWithPeriodResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateDataDiscoveryCloudAppDiscoveryUploadedStreamsItemMicrosoftGraphSecurityAggregatedAppsDetailsWithPeriodAggregatedAppsDetailsWithPeriodResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(DataDiscoveryCloudAppDiscoveryUploadedStreamsItemMicrosoftGraphSecurityAggregatedAppsDetailsWithPeriodAggregatedAppsDetailsWithPeriodResponseable), nil
}
// GetAsAggregatedAppsDetailsWithPeriodGetResponse invoke function aggregatedAppsDetails
// returns a DataDiscoveryCloudAppDiscoveryUploadedStreamsItemMicrosoftGraphSecurityAggregatedAppsDetailsWithPeriodAggregatedAppsDetailsWithPeriodGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DataDiscoveryCloudAppDiscoveryUploadedStreamsItemMicrosoftGraphSecurityAggregatedAppsDetailsWithPeriodRequestBuilder) GetAsAggregatedAppsDetailsWithPeriodGetResponse(ctx context.Context, requestConfiguration *DataDiscoveryCloudAppDiscoveryUploadedStreamsItemMicrosoftGraphSecurityAggregatedAppsDetailsWithPeriodRequestBuilderGetRequestConfiguration)(DataDiscoveryCloudAppDiscoveryUploadedStreamsItemMicrosoftGraphSecurityAggregatedAppsDetailsWithPeriodAggregatedAppsDetailsWithPeriodGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateDataDiscoveryCloudAppDiscoveryUploadedStreamsItemMicrosoftGraphSecurityAggregatedAppsDetailsWithPeriodAggregatedAppsDetailsWithPeriodGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(DataDiscoveryCloudAppDiscoveryUploadedStreamsItemMicrosoftGraphSecurityAggregatedAppsDetailsWithPeriodAggregatedAppsDetailsWithPeriodGetResponseable), nil
}
// ToGetRequestInformation invoke function aggregatedAppsDetails
// returns a *RequestInformation when successful
func (m *DataDiscoveryCloudAppDiscoveryUploadedStreamsItemMicrosoftGraphSecurityAggregatedAppsDetailsWithPeriodRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DataDiscoveryCloudAppDiscoveryUploadedStreamsItemMicrosoftGraphSecurityAggregatedAppsDetailsWithPeriodRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *DataDiscoveryCloudAppDiscoveryUploadedStreamsItemMicrosoftGraphSecurityAggregatedAppsDetailsWithPeriodRequestBuilder when successful
func (m *DataDiscoveryCloudAppDiscoveryUploadedStreamsItemMicrosoftGraphSecurityAggregatedAppsDetailsWithPeriodRequestBuilder) WithUrl(rawUrl string)(*DataDiscoveryCloudAppDiscoveryUploadedStreamsItemMicrosoftGraphSecurityAggregatedAppsDetailsWithPeriodRequestBuilder) {
    return NewDataDiscoveryCloudAppDiscoveryUploadedStreamsItemMicrosoftGraphSecurityAggregatedAppsDetailsWithPeriodRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
