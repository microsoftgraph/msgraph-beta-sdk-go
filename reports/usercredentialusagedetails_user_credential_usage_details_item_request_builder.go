package reports

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// UsercredentialusagedetailsUserCredentialUsageDetailsItemRequestBuilder provides operations to manage the userCredentialUsageDetails property of the microsoft.graph.reportRoot entity.
type UsercredentialusagedetailsUserCredentialUsageDetailsItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// UsercredentialusagedetailsUserCredentialUsageDetailsItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UsercredentialusagedetailsUserCredentialUsageDetailsItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// UsercredentialusagedetailsUserCredentialUsageDetailsItemRequestBuilderGetQueryParameters represents the self-service password reset (SSPR) usage for a given tenant.
type UsercredentialusagedetailsUserCredentialUsageDetailsItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// UsercredentialusagedetailsUserCredentialUsageDetailsItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UsercredentialusagedetailsUserCredentialUsageDetailsItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *UsercredentialusagedetailsUserCredentialUsageDetailsItemRequestBuilderGetQueryParameters
}
// UsercredentialusagedetailsUserCredentialUsageDetailsItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UsercredentialusagedetailsUserCredentialUsageDetailsItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewUsercredentialusagedetailsUserCredentialUsageDetailsItemRequestBuilderInternal instantiates a new UsercredentialusagedetailsUserCredentialUsageDetailsItemRequestBuilder and sets the default values.
func NewUsercredentialusagedetailsUserCredentialUsageDetailsItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UsercredentialusagedetailsUserCredentialUsageDetailsItemRequestBuilder) {
    m := &UsercredentialusagedetailsUserCredentialUsageDetailsItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/reports/userCredentialUsageDetails/{userCredentialUsageDetails%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewUsercredentialusagedetailsUserCredentialUsageDetailsItemRequestBuilder instantiates a new UsercredentialusagedetailsUserCredentialUsageDetailsItemRequestBuilder and sets the default values.
func NewUsercredentialusagedetailsUserCredentialUsageDetailsItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UsercredentialusagedetailsUserCredentialUsageDetailsItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUsercredentialusagedetailsUserCredentialUsageDetailsItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property userCredentialUsageDetails for reports
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UsercredentialusagedetailsUserCredentialUsageDetailsItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *UsercredentialusagedetailsUserCredentialUsageDetailsItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get represents the self-service password reset (SSPR) usage for a given tenant.
// returns a UserCredentialUsageDetailsable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UsercredentialusagedetailsUserCredentialUsageDetailsItemRequestBuilder) Get(ctx context.Context, requestConfiguration *UsercredentialusagedetailsUserCredentialUsageDetailsItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserCredentialUsageDetailsable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
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
// Patch update the navigation property userCredentialUsageDetails in reports
// returns a UserCredentialUsageDetailsable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UsercredentialusagedetailsUserCredentialUsageDetailsItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserCredentialUsageDetailsable, requestConfiguration *UsercredentialusagedetailsUserCredentialUsageDetailsItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserCredentialUsageDetailsable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
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
// ToDeleteRequestInformation delete navigation property userCredentialUsageDetails for reports
// returns a *RequestInformation when successful
func (m *UsercredentialusagedetailsUserCredentialUsageDetailsItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *UsercredentialusagedetailsUserCredentialUsageDetailsItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation represents the self-service password reset (SSPR) usage for a given tenant.
// returns a *RequestInformation when successful
func (m *UsercredentialusagedetailsUserCredentialUsageDetailsItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *UsercredentialusagedetailsUserCredentialUsageDetailsItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property userCredentialUsageDetails in reports
// returns a *RequestInformation when successful
func (m *UsercredentialusagedetailsUserCredentialUsageDetailsItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserCredentialUsageDetailsable, requestConfiguration *UsercredentialusagedetailsUserCredentialUsageDetailsItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// returns a *UsercredentialusagedetailsUserCredentialUsageDetailsItemRequestBuilder when successful
func (m *UsercredentialusagedetailsUserCredentialUsageDetailsItemRequestBuilder) WithUrl(rawUrl string)(*UsercredentialusagedetailsUserCredentialUsageDetailsItemRequestBuilder) {
    return NewUsercredentialusagedetailsUserCredentialUsageDetailsItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
