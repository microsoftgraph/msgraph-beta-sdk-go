package reports

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// AppCredentialSignInActivitiesAppCredentialSignInActivityItemRequestBuilder provides operations to manage the appCredentialSignInActivities property of the microsoft.graph.reportRoot entity.
type AppCredentialSignInActivitiesAppCredentialSignInActivityItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AppCredentialSignInActivitiesAppCredentialSignInActivityItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AppCredentialSignInActivitiesAppCredentialSignInActivityItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AppCredentialSignInActivitiesAppCredentialSignInActivityItemRequestBuilderGetQueryParameters get an appCredentialSignInActivity object that contains recent activity of an application credential.
type AppCredentialSignInActivitiesAppCredentialSignInActivityItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// AppCredentialSignInActivitiesAppCredentialSignInActivityItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AppCredentialSignInActivitiesAppCredentialSignInActivityItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AppCredentialSignInActivitiesAppCredentialSignInActivityItemRequestBuilderGetQueryParameters
}
// AppCredentialSignInActivitiesAppCredentialSignInActivityItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AppCredentialSignInActivitiesAppCredentialSignInActivityItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewAppCredentialSignInActivitiesAppCredentialSignInActivityItemRequestBuilderInternal instantiates a new AppCredentialSignInActivitiesAppCredentialSignInActivityItemRequestBuilder and sets the default values.
func NewAppCredentialSignInActivitiesAppCredentialSignInActivityItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AppCredentialSignInActivitiesAppCredentialSignInActivityItemRequestBuilder) {
    m := &AppCredentialSignInActivitiesAppCredentialSignInActivityItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/reports/appCredentialSignInActivities/{appCredentialSignInActivity%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewAppCredentialSignInActivitiesAppCredentialSignInActivityItemRequestBuilder instantiates a new AppCredentialSignInActivitiesAppCredentialSignInActivityItemRequestBuilder and sets the default values.
func NewAppCredentialSignInActivitiesAppCredentialSignInActivityItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AppCredentialSignInActivitiesAppCredentialSignInActivityItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAppCredentialSignInActivitiesAppCredentialSignInActivityItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property appCredentialSignInActivities for reports
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AppCredentialSignInActivitiesAppCredentialSignInActivityItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *AppCredentialSignInActivitiesAppCredentialSignInActivityItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// [Find more info here]: https://learn.microsoft.com/graph/api/appcredentialsigninactivity-get?view=graph-rest-1.0
func (m *AppCredentialSignInActivitiesAppCredentialSignInActivityItemRequestBuilder) Get(ctx context.Context, requestConfiguration *AppCredentialSignInActivitiesAppCredentialSignInActivityItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AppCredentialSignInActivityable, error) {
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
func (m *AppCredentialSignInActivitiesAppCredentialSignInActivityItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AppCredentialSignInActivityable, requestConfiguration *AppCredentialSignInActivitiesAppCredentialSignInActivityItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AppCredentialSignInActivityable, error) {
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
func (m *AppCredentialSignInActivitiesAppCredentialSignInActivityItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *AppCredentialSignInActivitiesAppCredentialSignInActivityItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, "{+baseurl}/reports/appCredentialSignInActivities/{appCredentialSignInActivity%2Did}", m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation get an appCredentialSignInActivity object that contains recent activity of an application credential.
// returns a *RequestInformation when successful
func (m *AppCredentialSignInActivitiesAppCredentialSignInActivityItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AppCredentialSignInActivitiesAppCredentialSignInActivityItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *AppCredentialSignInActivitiesAppCredentialSignInActivityItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AppCredentialSignInActivityable, requestConfiguration *AppCredentialSignInActivitiesAppCredentialSignInActivityItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, "{+baseurl}/reports/appCredentialSignInActivities/{appCredentialSignInActivity%2Did}", m.BaseRequestBuilder.PathParameters)
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
// returns a *AppCredentialSignInActivitiesAppCredentialSignInActivityItemRequestBuilder when successful
func (m *AppCredentialSignInActivitiesAppCredentialSignInActivityItemRequestBuilder) WithUrl(rawUrl string)(*AppCredentialSignInActivitiesAppCredentialSignInActivityItemRequestBuilder) {
    return NewAppCredentialSignInActivitiesAppCredentialSignInActivityItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
