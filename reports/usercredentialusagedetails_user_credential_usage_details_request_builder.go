package reports

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// UsercredentialusagedetailsUserCredentialUsageDetailsRequestBuilder provides operations to manage the userCredentialUsageDetails property of the microsoft.graph.reportRoot entity.
type UsercredentialusagedetailsUserCredentialUsageDetailsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// UsercredentialusagedetailsUserCredentialUsageDetailsRequestBuilderGetQueryParameters get a list of userCredentialUsageDetails objects for a given tenant. Details include user information, status of the reset, and the reason for failure.
type UsercredentialusagedetailsUserCredentialUsageDetailsRequestBuilderGetQueryParameters struct {
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
// UsercredentialusagedetailsUserCredentialUsageDetailsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UsercredentialusagedetailsUserCredentialUsageDetailsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *UsercredentialusagedetailsUserCredentialUsageDetailsRequestBuilderGetQueryParameters
}
// UsercredentialusagedetailsUserCredentialUsageDetailsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UsercredentialusagedetailsUserCredentialUsageDetailsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByUserCredentialUsageDetailsId provides operations to manage the userCredentialUsageDetails property of the microsoft.graph.reportRoot entity.
// returns a *UsercredentialusagedetailsUserCredentialUsageDetailsItemRequestBuilder when successful
func (m *UsercredentialusagedetailsUserCredentialUsageDetailsRequestBuilder) ByUserCredentialUsageDetailsId(userCredentialUsageDetailsId string)(*UsercredentialusagedetailsUserCredentialUsageDetailsItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if userCredentialUsageDetailsId != "" {
        urlTplParams["userCredentialUsageDetails%2Did"] = userCredentialUsageDetailsId
    }
    return NewUsercredentialusagedetailsUserCredentialUsageDetailsItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewUsercredentialusagedetailsUserCredentialUsageDetailsRequestBuilderInternal instantiates a new UsercredentialusagedetailsUserCredentialUsageDetailsRequestBuilder and sets the default values.
func NewUsercredentialusagedetailsUserCredentialUsageDetailsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UsercredentialusagedetailsUserCredentialUsageDetailsRequestBuilder) {
    m := &UsercredentialusagedetailsUserCredentialUsageDetailsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/reports/userCredentialUsageDetails{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewUsercredentialusagedetailsUserCredentialUsageDetailsRequestBuilder instantiates a new UsercredentialusagedetailsUserCredentialUsageDetailsRequestBuilder and sets the default values.
func NewUsercredentialusagedetailsUserCredentialUsageDetailsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UsercredentialusagedetailsUserCredentialUsageDetailsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUsercredentialusagedetailsUserCredentialUsageDetailsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *UsercredentialusagedetailsCountRequestBuilder when successful
func (m *UsercredentialusagedetailsUserCredentialUsageDetailsRequestBuilder) Count()(*UsercredentialusagedetailsCountRequestBuilder) {
    return NewUsercredentialusagedetailsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get a list of userCredentialUsageDetails objects for a given tenant. Details include user information, status of the reset, and the reason for failure.
// returns a UserCredentialUsageDetailsCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/reportroot-list-usercredentialusagedetails?view=graph-rest-beta
func (m *UsercredentialusagedetailsUserCredentialUsageDetailsRequestBuilder) Get(ctx context.Context, requestConfiguration *UsercredentialusagedetailsUserCredentialUsageDetailsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserCredentialUsageDetailsCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUserCredentialUsageDetailsCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserCredentialUsageDetailsCollectionResponseable), nil
}
// Post create new navigation property to userCredentialUsageDetails for reports
// returns a UserCredentialUsageDetailsable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UsercredentialusagedetailsUserCredentialUsageDetailsRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserCredentialUsageDetailsable, requestConfiguration *UsercredentialusagedetailsUserCredentialUsageDetailsRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserCredentialUsageDetailsable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUserCredentialUsageDetailsFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserCredentialUsageDetailsable), nil
}
// ToGetRequestInformation get a list of userCredentialUsageDetails objects for a given tenant. Details include user information, status of the reset, and the reason for failure.
// returns a *RequestInformation when successful
func (m *UsercredentialusagedetailsUserCredentialUsageDetailsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *UsercredentialusagedetailsUserCredentialUsageDetailsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to userCredentialUsageDetails for reports
// returns a *RequestInformation when successful
func (m *UsercredentialusagedetailsUserCredentialUsageDetailsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserCredentialUsageDetailsable, requestConfiguration *UsercredentialusagedetailsUserCredentialUsageDetailsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *UsercredentialusagedetailsUserCredentialUsageDetailsRequestBuilder when successful
func (m *UsercredentialusagedetailsUserCredentialUsageDetailsRequestBuilder) WithUrl(rawUrl string)(*UsercredentialusagedetailsUserCredentialUsageDetailsRequestBuilder) {
    return NewUsercredentialusagedetailsUserCredentialUsageDetailsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
