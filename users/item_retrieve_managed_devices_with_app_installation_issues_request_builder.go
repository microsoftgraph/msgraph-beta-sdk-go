package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemRetrieveManagedDevicesWithAppInstallationIssuesRequestBuilder provides operations to call the retrieveManagedDevicesWithAppInstallationIssues method.
type ItemRetrieveManagedDevicesWithAppInstallationIssuesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemRetrieveManagedDevicesWithAppInstallationIssuesRequestBuilderGetQueryParameters retrieves the list of devices with failed or pending apps
type ItemRetrieveManagedDevicesWithAppInstallationIssuesRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// ItemRetrieveManagedDevicesWithAppInstallationIssuesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemRetrieveManagedDevicesWithAppInstallationIssuesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemRetrieveManagedDevicesWithAppInstallationIssuesRequestBuilderGetQueryParameters
}
// NewItemRetrieveManagedDevicesWithAppInstallationIssuesRequestBuilderInternal instantiates a new ItemRetrieveManagedDevicesWithAppInstallationIssuesRequestBuilder and sets the default values.
func NewItemRetrieveManagedDevicesWithAppInstallationIssuesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemRetrieveManagedDevicesWithAppInstallationIssuesRequestBuilder) {
    m := &ItemRetrieveManagedDevicesWithAppInstallationIssuesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/retrieveManagedDevicesWithAppInstallationIssues(){?%24count,%24filter,%24search,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemRetrieveManagedDevicesWithAppInstallationIssuesRequestBuilder instantiates a new ItemRetrieveManagedDevicesWithAppInstallationIssuesRequestBuilder and sets the default values.
func NewItemRetrieveManagedDevicesWithAppInstallationIssuesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemRetrieveManagedDevicesWithAppInstallationIssuesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemRetrieveManagedDevicesWithAppInstallationIssuesRequestBuilderInternal(urlParams, requestAdapter)
}
// Get retrieves the list of devices with failed or pending apps
// Deprecated: This method is obsolete. Use GetAsRetrieveManagedDevicesWithAppInstallationIssuesGetResponse instead.
// returns a ItemRetrieveManagedDevicesWithAppInstallationIssuesResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemRetrieveManagedDevicesWithAppInstallationIssuesRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemRetrieveManagedDevicesWithAppInstallationIssuesRequestBuilderGetRequestConfiguration)(ItemRetrieveManagedDevicesWithAppInstallationIssuesResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemRetrieveManagedDevicesWithAppInstallationIssuesResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemRetrieveManagedDevicesWithAppInstallationIssuesResponseable), nil
}
// GetAsRetrieveManagedDevicesWithAppInstallationIssuesGetResponse retrieves the list of devices with failed or pending apps
// returns a ItemRetrieveManagedDevicesWithAppInstallationIssuesGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemRetrieveManagedDevicesWithAppInstallationIssuesRequestBuilder) GetAsRetrieveManagedDevicesWithAppInstallationIssuesGetResponse(ctx context.Context, requestConfiguration *ItemRetrieveManagedDevicesWithAppInstallationIssuesRequestBuilderGetRequestConfiguration)(ItemRetrieveManagedDevicesWithAppInstallationIssuesGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemRetrieveManagedDevicesWithAppInstallationIssuesGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemRetrieveManagedDevicesWithAppInstallationIssuesGetResponseable), nil
}
// ToGetRequestInformation retrieves the list of devices with failed or pending apps
// returns a *RequestInformation when successful
func (m *ItemRetrieveManagedDevicesWithAppInstallationIssuesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemRetrieveManagedDevicesWithAppInstallationIssuesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemRetrieveManagedDevicesWithAppInstallationIssuesRequestBuilder when successful
func (m *ItemRetrieveManagedDevicesWithAppInstallationIssuesRequestBuilder) WithUrl(rawUrl string)(*ItemRetrieveManagedDevicesWithAppInstallationIssuesRequestBuilder) {
    return NewItemRetrieveManagedDevicesWithAppInstallationIssuesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
