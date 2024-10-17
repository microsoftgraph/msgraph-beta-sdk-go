package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// CloudCertificationAuthorityRequestBuilder provides operations to manage the cloudCertificationAuthority property of the microsoft.graph.deviceManagement entity.
type CloudCertificationAuthorityRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CloudCertificationAuthorityRequestBuilderGetQueryParameters collection of CloudCertificationAuthority records associated with account.
type CloudCertificationAuthorityRequestBuilderGetQueryParameters struct {
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
// CloudCertificationAuthorityRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CloudCertificationAuthorityRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CloudCertificationAuthorityRequestBuilderGetQueryParameters
}
// CloudCertificationAuthorityRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CloudCertificationAuthorityRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByCloudCertificationAuthorityId provides operations to manage the cloudCertificationAuthority property of the microsoft.graph.deviceManagement entity.
// returns a *CloudCertificationAuthorityCloudCertificationAuthorityItemRequestBuilder when successful
func (m *CloudCertificationAuthorityRequestBuilder) ByCloudCertificationAuthorityId(cloudCertificationAuthorityId string)(*CloudCertificationAuthorityCloudCertificationAuthorityItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if cloudCertificationAuthorityId != "" {
        urlTplParams["cloudCertificationAuthority%2Did"] = cloudCertificationAuthorityId
    }
    return NewCloudCertificationAuthorityCloudCertificationAuthorityItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewCloudCertificationAuthorityRequestBuilderInternal instantiates a new CloudCertificationAuthorityRequestBuilder and sets the default values.
func NewCloudCertificationAuthorityRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CloudCertificationAuthorityRequestBuilder) {
    m := &CloudCertificationAuthorityRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/cloudCertificationAuthority{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewCloudCertificationAuthorityRequestBuilder instantiates a new CloudCertificationAuthorityRequestBuilder and sets the default values.
func NewCloudCertificationAuthorityRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CloudCertificationAuthorityRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCloudCertificationAuthorityRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *CloudCertificationAuthorityCountRequestBuilder when successful
func (m *CloudCertificationAuthorityRequestBuilder) Count()(*CloudCertificationAuthorityCountRequestBuilder) {
    return NewCloudCertificationAuthorityCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get collection of CloudCertificationAuthority records associated with account.
// returns a CloudCertificationAuthorityCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CloudCertificationAuthorityRequestBuilder) Get(ctx context.Context, requestConfiguration *CloudCertificationAuthorityRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudCertificationAuthorityCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateCloudCertificationAuthorityCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudCertificationAuthorityCollectionResponseable), nil
}
// Post create new navigation property to cloudCertificationAuthority for deviceManagement
// returns a CloudCertificationAuthorityable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CloudCertificationAuthorityRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudCertificationAuthorityable, requestConfiguration *CloudCertificationAuthorityRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudCertificationAuthorityable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateCloudCertificationAuthorityFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudCertificationAuthorityable), nil
}
// ToGetRequestInformation collection of CloudCertificationAuthority records associated with account.
// returns a *RequestInformation when successful
func (m *CloudCertificationAuthorityRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CloudCertificationAuthorityRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to cloudCertificationAuthority for deviceManagement
// returns a *RequestInformation when successful
func (m *CloudCertificationAuthorityRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudCertificationAuthorityable, requestConfiguration *CloudCertificationAuthorityRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *CloudCertificationAuthorityRequestBuilder when successful
func (m *CloudCertificationAuthorityRequestBuilder) WithUrl(rawUrl string)(*CloudCertificationAuthorityRequestBuilder) {
    return NewCloudCertificationAuthorityRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
