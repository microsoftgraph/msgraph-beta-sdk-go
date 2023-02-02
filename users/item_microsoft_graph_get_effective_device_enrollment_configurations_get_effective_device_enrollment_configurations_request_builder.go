package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemMicrosoftGraphGetEffectiveDeviceEnrollmentConfigurationsGetEffectiveDeviceEnrollmentConfigurationsRequestBuilder provides operations to call the getEffectiveDeviceEnrollmentConfigurations method.
type ItemMicrosoftGraphGetEffectiveDeviceEnrollmentConfigurationsGetEffectiveDeviceEnrollmentConfigurationsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemMicrosoftGraphGetEffectiveDeviceEnrollmentConfigurationsGetEffectiveDeviceEnrollmentConfigurationsRequestBuilderGetQueryParameters invoke function getEffectiveDeviceEnrollmentConfigurations
type ItemMicrosoftGraphGetEffectiveDeviceEnrollmentConfigurationsGetEffectiveDeviceEnrollmentConfigurationsRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
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
// ItemMicrosoftGraphGetEffectiveDeviceEnrollmentConfigurationsGetEffectiveDeviceEnrollmentConfigurationsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemMicrosoftGraphGetEffectiveDeviceEnrollmentConfigurationsGetEffectiveDeviceEnrollmentConfigurationsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemMicrosoftGraphGetEffectiveDeviceEnrollmentConfigurationsGetEffectiveDeviceEnrollmentConfigurationsRequestBuilderGetQueryParameters
}
// NewItemMicrosoftGraphGetEffectiveDeviceEnrollmentConfigurationsGetEffectiveDeviceEnrollmentConfigurationsRequestBuilderInternal instantiates a new GetEffectiveDeviceEnrollmentConfigurationsRequestBuilder and sets the default values.
func NewItemMicrosoftGraphGetEffectiveDeviceEnrollmentConfigurationsGetEffectiveDeviceEnrollmentConfigurationsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMicrosoftGraphGetEffectiveDeviceEnrollmentConfigurationsGetEffectiveDeviceEnrollmentConfigurationsRequestBuilder) {
    m := &ItemMicrosoftGraphGetEffectiveDeviceEnrollmentConfigurationsGetEffectiveDeviceEnrollmentConfigurationsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/microsoft.graph.getEffectiveDeviceEnrollmentConfigurations(){?%24top,%24skip,%24search,%24filter,%24count,%24select,%24orderby}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemMicrosoftGraphGetEffectiveDeviceEnrollmentConfigurationsGetEffectiveDeviceEnrollmentConfigurationsRequestBuilder instantiates a new GetEffectiveDeviceEnrollmentConfigurationsRequestBuilder and sets the default values.
func NewItemMicrosoftGraphGetEffectiveDeviceEnrollmentConfigurationsGetEffectiveDeviceEnrollmentConfigurationsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMicrosoftGraphGetEffectiveDeviceEnrollmentConfigurationsGetEffectiveDeviceEnrollmentConfigurationsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemMicrosoftGraphGetEffectiveDeviceEnrollmentConfigurationsGetEffectiveDeviceEnrollmentConfigurationsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get invoke function getEffectiveDeviceEnrollmentConfigurations
func (m *ItemMicrosoftGraphGetEffectiveDeviceEnrollmentConfigurationsGetEffectiveDeviceEnrollmentConfigurationsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemMicrosoftGraphGetEffectiveDeviceEnrollmentConfigurationsGetEffectiveDeviceEnrollmentConfigurationsRequestBuilderGetRequestConfiguration)(ItemMicrosoftGraphGetEffectiveDeviceEnrollmentConfigurationsGetEffectiveDeviceEnrollmentConfigurationsResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, CreateItemMicrosoftGraphGetEffectiveDeviceEnrollmentConfigurationsGetEffectiveDeviceEnrollmentConfigurationsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemMicrosoftGraphGetEffectiveDeviceEnrollmentConfigurationsGetEffectiveDeviceEnrollmentConfigurationsResponseable), nil
}
// ToGetRequestInformation invoke function getEffectiveDeviceEnrollmentConfigurations
func (m *ItemMicrosoftGraphGetEffectiveDeviceEnrollmentConfigurationsGetEffectiveDeviceEnrollmentConfigurationsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemMicrosoftGraphGetEffectiveDeviceEnrollmentConfigurationsGetEffectiveDeviceEnrollmentConfigurationsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
