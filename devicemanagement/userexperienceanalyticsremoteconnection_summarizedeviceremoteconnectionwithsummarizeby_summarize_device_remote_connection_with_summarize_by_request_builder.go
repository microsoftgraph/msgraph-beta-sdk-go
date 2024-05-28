package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// UserexperienceanalyticsremoteconnectionSummarizedeviceremoteconnectionwithsummarizebySummarizeDeviceRemoteConnectionWithSummarizeByRequestBuilder provides operations to call the summarizeDeviceRemoteConnection method.
type UserexperienceanalyticsremoteconnectionSummarizedeviceremoteconnectionwithsummarizebySummarizeDeviceRemoteConnectionWithSummarizeByRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// UserexperienceanalyticsremoteconnectionSummarizedeviceremoteconnectionwithsummarizebySummarizeDeviceRemoteConnectionWithSummarizeByRequestBuilderGetQueryParameters invoke function summarizeDeviceRemoteConnection
type UserexperienceanalyticsremoteconnectionSummarizedeviceremoteconnectionwithsummarizebySummarizeDeviceRemoteConnectionWithSummarizeByRequestBuilderGetQueryParameters struct {
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
// UserexperienceanalyticsremoteconnectionSummarizedeviceremoteconnectionwithsummarizebySummarizeDeviceRemoteConnectionWithSummarizeByRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserexperienceanalyticsremoteconnectionSummarizedeviceremoteconnectionwithsummarizebySummarizeDeviceRemoteConnectionWithSummarizeByRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *UserexperienceanalyticsremoteconnectionSummarizedeviceremoteconnectionwithsummarizebySummarizeDeviceRemoteConnectionWithSummarizeByRequestBuilderGetQueryParameters
}
// NewUserexperienceanalyticsremoteconnectionSummarizedeviceremoteconnectionwithsummarizebySummarizeDeviceRemoteConnectionWithSummarizeByRequestBuilderInternal instantiates a new UserexperienceanalyticsremoteconnectionSummarizedeviceremoteconnectionwithsummarizebySummarizeDeviceRemoteConnectionWithSummarizeByRequestBuilder and sets the default values.
func NewUserexperienceanalyticsremoteconnectionSummarizedeviceremoteconnectionwithsummarizebySummarizeDeviceRemoteConnectionWithSummarizeByRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, summarizeBy *string)(*UserexperienceanalyticsremoteconnectionSummarizedeviceremoteconnectionwithsummarizebySummarizeDeviceRemoteConnectionWithSummarizeByRequestBuilder) {
    m := &UserexperienceanalyticsremoteconnectionSummarizedeviceremoteconnectionwithsummarizebySummarizeDeviceRemoteConnectionWithSummarizeByRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/userExperienceAnalyticsRemoteConnection/summarizeDeviceRemoteConnection(summarizeBy='{summarizeBy}'){?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    if summarizeBy != nil {
        m.BaseRequestBuilder.PathParameters["summarizeBy"] = *summarizeBy
    }
    return m
}
// NewUserexperienceanalyticsremoteconnectionSummarizedeviceremoteconnectionwithsummarizebySummarizeDeviceRemoteConnectionWithSummarizeByRequestBuilder instantiates a new UserexperienceanalyticsremoteconnectionSummarizedeviceremoteconnectionwithsummarizebySummarizeDeviceRemoteConnectionWithSummarizeByRequestBuilder and sets the default values.
func NewUserexperienceanalyticsremoteconnectionSummarizedeviceremoteconnectionwithsummarizebySummarizeDeviceRemoteConnectionWithSummarizeByRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserexperienceanalyticsremoteconnectionSummarizedeviceremoteconnectionwithsummarizebySummarizeDeviceRemoteConnectionWithSummarizeByRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUserexperienceanalyticsremoteconnectionSummarizedeviceremoteconnectionwithsummarizebySummarizeDeviceRemoteConnectionWithSummarizeByRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get invoke function summarizeDeviceRemoteConnection
// Deprecated: This method is obsolete. Use GetAsSummarizeDeviceRemoteConnectionWithSummarizeByGetResponse instead.
// returns a UserexperienceanalyticsremoteconnectionSummarizedeviceremoteconnectionwithsummarizebySummarizeDeviceRemoteConnectionWithSummarizeByResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UserexperienceanalyticsremoteconnectionSummarizedeviceremoteconnectionwithsummarizebySummarizeDeviceRemoteConnectionWithSummarizeByRequestBuilder) Get(ctx context.Context, requestConfiguration *UserexperienceanalyticsremoteconnectionSummarizedeviceremoteconnectionwithsummarizebySummarizeDeviceRemoteConnectionWithSummarizeByRequestBuilderGetRequestConfiguration)(UserexperienceanalyticsremoteconnectionSummarizedeviceremoteconnectionwithsummarizebySummarizeDeviceRemoteConnectionWithSummarizeByResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateUserexperienceanalyticsremoteconnectionSummarizedeviceremoteconnectionwithsummarizebySummarizeDeviceRemoteConnectionWithSummarizeByResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(UserexperienceanalyticsremoteconnectionSummarizedeviceremoteconnectionwithsummarizebySummarizeDeviceRemoteConnectionWithSummarizeByResponseable), nil
}
// GetAsSummarizeDeviceRemoteConnectionWithSummarizeByGetResponse invoke function summarizeDeviceRemoteConnection
// returns a UserexperienceanalyticsremoteconnectionSummarizedeviceremoteconnectionwithsummarizebySummarizeDeviceRemoteConnectionWithSummarizeByGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UserexperienceanalyticsremoteconnectionSummarizedeviceremoteconnectionwithsummarizebySummarizeDeviceRemoteConnectionWithSummarizeByRequestBuilder) GetAsSummarizeDeviceRemoteConnectionWithSummarizeByGetResponse(ctx context.Context, requestConfiguration *UserexperienceanalyticsremoteconnectionSummarizedeviceremoteconnectionwithsummarizebySummarizeDeviceRemoteConnectionWithSummarizeByRequestBuilderGetRequestConfiguration)(UserexperienceanalyticsremoteconnectionSummarizedeviceremoteconnectionwithsummarizebySummarizeDeviceRemoteConnectionWithSummarizeByGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateUserexperienceanalyticsremoteconnectionSummarizedeviceremoteconnectionwithsummarizebySummarizeDeviceRemoteConnectionWithSummarizeByGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(UserexperienceanalyticsremoteconnectionSummarizedeviceremoteconnectionwithsummarizebySummarizeDeviceRemoteConnectionWithSummarizeByGetResponseable), nil
}
// ToGetRequestInformation invoke function summarizeDeviceRemoteConnection
// returns a *RequestInformation when successful
func (m *UserexperienceanalyticsremoteconnectionSummarizedeviceremoteconnectionwithsummarizebySummarizeDeviceRemoteConnectionWithSummarizeByRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *UserexperienceanalyticsremoteconnectionSummarizedeviceremoteconnectionwithsummarizebySummarizeDeviceRemoteConnectionWithSummarizeByRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *UserexperienceanalyticsremoteconnectionSummarizedeviceremoteconnectionwithsummarizebySummarizeDeviceRemoteConnectionWithSummarizeByRequestBuilder when successful
func (m *UserexperienceanalyticsremoteconnectionSummarizedeviceremoteconnectionwithsummarizebySummarizeDeviceRemoteConnectionWithSummarizeByRequestBuilder) WithUrl(rawUrl string)(*UserexperienceanalyticsremoteconnectionSummarizedeviceremoteconnectionwithsummarizebySummarizeDeviceRemoteConnectionWithSummarizeByRequestBuilder) {
    return NewUserexperienceanalyticsremoteconnectionSummarizedeviceremoteconnectionwithsummarizebySummarizeDeviceRemoteConnectionWithSummarizeByRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
