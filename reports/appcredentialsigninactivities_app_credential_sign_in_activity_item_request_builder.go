package reports

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// AppcredentialsigninactivitiesAppCredentialSignInActivityItemRequestBuilder provides operations to manage the appCredentialSignInActivities property of the microsoft.graph.reportRoot entity.
type AppcredentialsigninactivitiesAppCredentialSignInActivityItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AppcredentialsigninactivitiesAppCredentialSignInActivityItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AppcredentialsigninactivitiesAppCredentialSignInActivityItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AppcredentialsigninactivitiesAppCredentialSignInActivityItemRequestBuilderGetQueryParameters get an appCredentialSignInActivity object that contains recent activity of an application credential.
type AppcredentialsigninactivitiesAppCredentialSignInActivityItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// AppcredentialsigninactivitiesAppCredentialSignInActivityItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AppcredentialsigninactivitiesAppCredentialSignInActivityItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AppcredentialsigninactivitiesAppCredentialSignInActivityItemRequestBuilderGetQueryParameters
}
// AppcredentialsigninactivitiesAppCredentialSignInActivityItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AppcredentialsigninactivitiesAppCredentialSignInActivityItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewAppcredentialsigninactivitiesAppCredentialSignInActivityItemRequestBuilderInternal instantiates a new AppcredentialsigninactivitiesAppCredentialSignInActivityItemRequestBuilder and sets the default values.
func NewAppcredentialsigninactivitiesAppCredentialSignInActivityItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AppcredentialsigninactivitiesAppCredentialSignInActivityItemRequestBuilder) {
    m := &AppcredentialsigninactivitiesAppCredentialSignInActivityItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/reports/appCredentialSignInActivities/{appCredentialSignInActivity%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewAppcredentialsigninactivitiesAppCredentialSignInActivityItemRequestBuilder instantiates a new AppcredentialsigninactivitiesAppCredentialSignInActivityItemRequestBuilder and sets the default values.
func NewAppcredentialsigninactivitiesAppCredentialSignInActivityItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AppcredentialsigninactivitiesAppCredentialSignInActivityItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAppcredentialsigninactivitiesAppCredentialSignInActivityItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property appCredentialSignInActivities for reports
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AppcredentialsigninactivitiesAppCredentialSignInActivityItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *AppcredentialsigninactivitiesAppCredentialSignInActivityItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get get an appCredentialSignInActivity object that contains recent activity of an application credential.
// returns a AppCredentialSignInActivityable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/appcredentialsigninactivity-get?view=graph-rest-beta
func (m *AppcredentialsigninactivitiesAppCredentialSignInActivityItemRequestBuilder) Get(ctx context.Context, requestConfiguration *AppcredentialsigninactivitiesAppCredentialSignInActivityItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AppCredentialSignInActivityable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAppCredentialSignInActivityFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AppCredentialSignInActivityable), nil
}
// Patch update the navigation property appCredentialSignInActivities in reports
// returns a AppCredentialSignInActivityable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AppcredentialsigninactivitiesAppCredentialSignInActivityItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AppCredentialSignInActivityable, requestConfiguration *AppcredentialsigninactivitiesAppCredentialSignInActivityItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AppCredentialSignInActivityable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAppCredentialSignInActivityFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AppCredentialSignInActivityable), nil
}
// ToDeleteRequestInformation delete navigation property appCredentialSignInActivities for reports
// returns a *RequestInformation when successful
func (m *AppcredentialsigninactivitiesAppCredentialSignInActivityItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *AppcredentialsigninactivitiesAppCredentialSignInActivityItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation get an appCredentialSignInActivity object that contains recent activity of an application credential.
// returns a *RequestInformation when successful
func (m *AppcredentialsigninactivitiesAppCredentialSignInActivityItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AppcredentialsigninactivitiesAppCredentialSignInActivityItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property appCredentialSignInActivities in reports
// returns a *RequestInformation when successful
func (m *AppcredentialsigninactivitiesAppCredentialSignInActivityItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AppCredentialSignInActivityable, requestConfiguration *AppcredentialsigninactivitiesAppCredentialSignInActivityItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *AppcredentialsigninactivitiesAppCredentialSignInActivityItemRequestBuilder when successful
func (m *AppcredentialsigninactivitiesAppCredentialSignInActivityItemRequestBuilder) WithUrl(rawUrl string)(*AppcredentialsigninactivitiesAppCredentialSignInActivityItemRequestBuilder) {
    return NewAppcredentialsigninactivitiesAppCredentialSignInActivityItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
