package reports

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// AuthenticationmethodsUserregistrationdetailsUserRegistrationDetailsRequestBuilder provides operations to manage the userRegistrationDetails property of the microsoft.graph.authenticationMethodsRoot entity.
type AuthenticationmethodsUserregistrationdetailsUserRegistrationDetailsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AuthenticationmethodsUserregistrationdetailsUserRegistrationDetailsRequestBuilderGetQueryParameters get a list of the authentication methods registered for a user as defined in the userRegistrationDetails object.
type AuthenticationmethodsUserregistrationdetailsUserRegistrationDetailsRequestBuilderGetQueryParameters struct {
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
// AuthenticationmethodsUserregistrationdetailsUserRegistrationDetailsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AuthenticationmethodsUserregistrationdetailsUserRegistrationDetailsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AuthenticationmethodsUserregistrationdetailsUserRegistrationDetailsRequestBuilderGetQueryParameters
}
// AuthenticationmethodsUserregistrationdetailsUserRegistrationDetailsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AuthenticationmethodsUserregistrationdetailsUserRegistrationDetailsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByUserRegistrationDetailsId provides operations to manage the userRegistrationDetails property of the microsoft.graph.authenticationMethodsRoot entity.
// returns a *AuthenticationmethodsUserregistrationdetailsUserRegistrationDetailsItemRequestBuilder when successful
func (m *AuthenticationmethodsUserregistrationdetailsUserRegistrationDetailsRequestBuilder) ByUserRegistrationDetailsId(userRegistrationDetailsId string)(*AuthenticationmethodsUserregistrationdetailsUserRegistrationDetailsItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if userRegistrationDetailsId != "" {
        urlTplParams["userRegistrationDetails%2Did"] = userRegistrationDetailsId
    }
    return NewAuthenticationmethodsUserregistrationdetailsUserRegistrationDetailsItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewAuthenticationmethodsUserregistrationdetailsUserRegistrationDetailsRequestBuilderInternal instantiates a new AuthenticationmethodsUserregistrationdetailsUserRegistrationDetailsRequestBuilder and sets the default values.
func NewAuthenticationmethodsUserregistrationdetailsUserRegistrationDetailsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AuthenticationmethodsUserregistrationdetailsUserRegistrationDetailsRequestBuilder) {
    m := &AuthenticationmethodsUserregistrationdetailsUserRegistrationDetailsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/reports/authenticationMethods/userRegistrationDetails{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewAuthenticationmethodsUserregistrationdetailsUserRegistrationDetailsRequestBuilder instantiates a new AuthenticationmethodsUserregistrationdetailsUserRegistrationDetailsRequestBuilder and sets the default values.
func NewAuthenticationmethodsUserregistrationdetailsUserRegistrationDetailsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AuthenticationmethodsUserregistrationdetailsUserRegistrationDetailsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAuthenticationmethodsUserregistrationdetailsUserRegistrationDetailsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *AuthenticationmethodsUserregistrationdetailsCountRequestBuilder when successful
func (m *AuthenticationmethodsUserregistrationdetailsUserRegistrationDetailsRequestBuilder) Count()(*AuthenticationmethodsUserregistrationdetailsCountRequestBuilder) {
    return NewAuthenticationmethodsUserregistrationdetailsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get a list of the authentication methods registered for a user as defined in the userRegistrationDetails object.
// returns a UserRegistrationDetailsCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/authenticationmethodsroot-list-userregistrationdetails?view=graph-rest-beta
func (m *AuthenticationmethodsUserregistrationdetailsUserRegistrationDetailsRequestBuilder) Get(ctx context.Context, requestConfiguration *AuthenticationmethodsUserregistrationdetailsUserRegistrationDetailsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserRegistrationDetailsCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUserRegistrationDetailsCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserRegistrationDetailsCollectionResponseable), nil
}
// Post create new navigation property to userRegistrationDetails for reports
// returns a UserRegistrationDetailsable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AuthenticationmethodsUserregistrationdetailsUserRegistrationDetailsRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserRegistrationDetailsable, requestConfiguration *AuthenticationmethodsUserregistrationdetailsUserRegistrationDetailsRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserRegistrationDetailsable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUserRegistrationDetailsFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserRegistrationDetailsable), nil
}
// ToGetRequestInformation get a list of the authentication methods registered for a user as defined in the userRegistrationDetails object.
// returns a *RequestInformation when successful
func (m *AuthenticationmethodsUserregistrationdetailsUserRegistrationDetailsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AuthenticationmethodsUserregistrationdetailsUserRegistrationDetailsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to userRegistrationDetails for reports
// returns a *RequestInformation when successful
func (m *AuthenticationmethodsUserregistrationdetailsUserRegistrationDetailsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserRegistrationDetailsable, requestConfiguration *AuthenticationmethodsUserregistrationdetailsUserRegistrationDetailsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *AuthenticationmethodsUserregistrationdetailsUserRegistrationDetailsRequestBuilder when successful
func (m *AuthenticationmethodsUserregistrationdetailsUserRegistrationDetailsRequestBuilder) WithUrl(rawUrl string)(*AuthenticationmethodsUserregistrationdetailsUserRegistrationDetailsRequestBuilder) {
    return NewAuthenticationmethodsUserregistrationdetailsUserRegistrationDetailsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
