package deviceappmanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// IoslobappprovisioningconfigurationsItemGroupassignmentsCountRequestBuilder provides operations to count the resources in the collection.
type IoslobappprovisioningconfigurationsItemGroupassignmentsCountRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// IoslobappprovisioningconfigurationsItemGroupassignmentsCountRequestBuilderGetQueryParameters get the number of the resource
type IoslobappprovisioningconfigurationsItemGroupassignmentsCountRequestBuilderGetQueryParameters struct {
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
}
// IoslobappprovisioningconfigurationsItemGroupassignmentsCountRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IoslobappprovisioningconfigurationsItemGroupassignmentsCountRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *IoslobappprovisioningconfigurationsItemGroupassignmentsCountRequestBuilderGetQueryParameters
}
// NewIoslobappprovisioningconfigurationsItemGroupassignmentsCountRequestBuilderInternal instantiates a new IoslobappprovisioningconfigurationsItemGroupassignmentsCountRequestBuilder and sets the default values.
func NewIoslobappprovisioningconfigurationsItemGroupassignmentsCountRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IoslobappprovisioningconfigurationsItemGroupassignmentsCountRequestBuilder) {
    m := &IoslobappprovisioningconfigurationsItemGroupassignmentsCountRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceAppManagement/iosLobAppProvisioningConfigurations/{iosLobAppProvisioningConfiguration%2Did}/groupAssignments/$count{?%24filter,%24search}", pathParameters),
    }
    return m
}
// NewIoslobappprovisioningconfigurationsItemGroupassignmentsCountRequestBuilder instantiates a new IoslobappprovisioningconfigurationsItemGroupassignmentsCountRequestBuilder and sets the default values.
func NewIoslobappprovisioningconfigurationsItemGroupassignmentsCountRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IoslobappprovisioningconfigurationsItemGroupassignmentsCountRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewIoslobappprovisioningconfigurationsItemGroupassignmentsCountRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get the number of the resource
// returns a *int32 when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *IoslobappprovisioningconfigurationsItemGroupassignmentsCountRequestBuilder) Get(ctx context.Context, requestConfiguration *IoslobappprovisioningconfigurationsItemGroupassignmentsCountRequestBuilderGetRequestConfiguration)(*int32, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendPrimitive(ctx, requestInfo, "int32", errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(*int32), nil
}
// ToGetRequestInformation get the number of the resource
// returns a *RequestInformation when successful
func (m *IoslobappprovisioningconfigurationsItemGroupassignmentsCountRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *IoslobappprovisioningconfigurationsItemGroupassignmentsCountRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "text/plain;q=0.9")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *IoslobappprovisioningconfigurationsItemGroupassignmentsCountRequestBuilder when successful
func (m *IoslobappprovisioningconfigurationsItemGroupassignmentsCountRequestBuilder) WithUrl(rawUrl string)(*IoslobappprovisioningconfigurationsItemGroupassignmentsCountRequestBuilder) {
    return NewIoslobappprovisioningconfigurationsItemGroupassignmentsCountRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
