package identity

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// AuthenticationeventsflowsItemConditionsApplicationsIncludeapplicationsAuthenticationConditionApplicationAppItemRequestBuilder provides operations to manage the includeApplications property of the microsoft.graph.authenticationConditionsApplications entity.
type AuthenticationeventsflowsItemConditionsApplicationsIncludeapplicationsAuthenticationConditionApplicationAppItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AuthenticationeventsflowsItemConditionsApplicationsIncludeapplicationsAuthenticationConditionApplicationAppItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AuthenticationeventsflowsItemConditionsApplicationsIncludeapplicationsAuthenticationConditionApplicationAppItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AuthenticationeventsflowsItemConditionsApplicationsIncludeapplicationsAuthenticationConditionApplicationAppItemRequestBuilderGetQueryParameters get includeApplications from identity
type AuthenticationeventsflowsItemConditionsApplicationsIncludeapplicationsAuthenticationConditionApplicationAppItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// AuthenticationeventsflowsItemConditionsApplicationsIncludeapplicationsAuthenticationConditionApplicationAppItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AuthenticationeventsflowsItemConditionsApplicationsIncludeapplicationsAuthenticationConditionApplicationAppItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AuthenticationeventsflowsItemConditionsApplicationsIncludeapplicationsAuthenticationConditionApplicationAppItemRequestBuilderGetQueryParameters
}
// AuthenticationeventsflowsItemConditionsApplicationsIncludeapplicationsAuthenticationConditionApplicationAppItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AuthenticationeventsflowsItemConditionsApplicationsIncludeapplicationsAuthenticationConditionApplicationAppItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewAuthenticationeventsflowsItemConditionsApplicationsIncludeapplicationsAuthenticationConditionApplicationAppItemRequestBuilderInternal instantiates a new AuthenticationeventsflowsItemConditionsApplicationsIncludeapplicationsAuthenticationConditionApplicationAppItemRequestBuilder and sets the default values.
func NewAuthenticationeventsflowsItemConditionsApplicationsIncludeapplicationsAuthenticationConditionApplicationAppItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AuthenticationeventsflowsItemConditionsApplicationsIncludeapplicationsAuthenticationConditionApplicationAppItemRequestBuilder) {
    m := &AuthenticationeventsflowsItemConditionsApplicationsIncludeapplicationsAuthenticationConditionApplicationAppItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identity/authenticationEventsFlows/{authenticationEventsFlow%2Did}/conditions/applications/includeApplications/{authenticationConditionApplication%2DappId}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewAuthenticationeventsflowsItemConditionsApplicationsIncludeapplicationsAuthenticationConditionApplicationAppItemRequestBuilder instantiates a new AuthenticationeventsflowsItemConditionsApplicationsIncludeapplicationsAuthenticationConditionApplicationAppItemRequestBuilder and sets the default values.
func NewAuthenticationeventsflowsItemConditionsApplicationsIncludeapplicationsAuthenticationConditionApplicationAppItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AuthenticationeventsflowsItemConditionsApplicationsIncludeapplicationsAuthenticationConditionApplicationAppItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAuthenticationeventsflowsItemConditionsApplicationsIncludeapplicationsAuthenticationConditionApplicationAppItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete remove or unlink an application from an external identities self-service sign-up user flow that's represented by an externalUsersSelfServiceSignupEventsFlow object. This disables the authentication experience that's defined by the user flow for the application.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/authenticationconditionapplication-delete?view=graph-rest-beta
func (m *AuthenticationeventsflowsItemConditionsApplicationsIncludeapplicationsAuthenticationConditionApplicationAppItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *AuthenticationeventsflowsItemConditionsApplicationsIncludeapplicationsAuthenticationConditionApplicationAppItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get get includeApplications from identity
// returns a AuthenticationConditionApplicationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AuthenticationeventsflowsItemConditionsApplicationsIncludeapplicationsAuthenticationConditionApplicationAppItemRequestBuilder) Get(ctx context.Context, requestConfiguration *AuthenticationeventsflowsItemConditionsApplicationsIncludeapplicationsAuthenticationConditionApplicationAppItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AuthenticationConditionApplicationable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAuthenticationConditionApplicationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AuthenticationConditionApplicationable), nil
}
// Patch update the navigation property includeApplications in identity
// returns a AuthenticationConditionApplicationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AuthenticationeventsflowsItemConditionsApplicationsIncludeapplicationsAuthenticationConditionApplicationAppItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AuthenticationConditionApplicationable, requestConfiguration *AuthenticationeventsflowsItemConditionsApplicationsIncludeapplicationsAuthenticationConditionApplicationAppItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AuthenticationConditionApplicationable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAuthenticationConditionApplicationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AuthenticationConditionApplicationable), nil
}
// ToDeleteRequestInformation remove or unlink an application from an external identities self-service sign-up user flow that's represented by an externalUsersSelfServiceSignupEventsFlow object. This disables the authentication experience that's defined by the user flow for the application.
// returns a *RequestInformation when successful
func (m *AuthenticationeventsflowsItemConditionsApplicationsIncludeapplicationsAuthenticationConditionApplicationAppItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *AuthenticationeventsflowsItemConditionsApplicationsIncludeapplicationsAuthenticationConditionApplicationAppItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation get includeApplications from identity
// returns a *RequestInformation when successful
func (m *AuthenticationeventsflowsItemConditionsApplicationsIncludeapplicationsAuthenticationConditionApplicationAppItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AuthenticationeventsflowsItemConditionsApplicationsIncludeapplicationsAuthenticationConditionApplicationAppItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property includeApplications in identity
// returns a *RequestInformation when successful
func (m *AuthenticationeventsflowsItemConditionsApplicationsIncludeapplicationsAuthenticationConditionApplicationAppItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AuthenticationConditionApplicationable, requestConfiguration *AuthenticationeventsflowsItemConditionsApplicationsIncludeapplicationsAuthenticationConditionApplicationAppItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *AuthenticationeventsflowsItemConditionsApplicationsIncludeapplicationsAuthenticationConditionApplicationAppItemRequestBuilder when successful
func (m *AuthenticationeventsflowsItemConditionsApplicationsIncludeapplicationsAuthenticationConditionApplicationAppItemRequestBuilder) WithUrl(rawUrl string)(*AuthenticationeventsflowsItemConditionsApplicationsIncludeapplicationsAuthenticationConditionApplicationAppItemRequestBuilder) {
    return NewAuthenticationeventsflowsItemConditionsApplicationsIncludeapplicationsAuthenticationConditionApplicationAppItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
