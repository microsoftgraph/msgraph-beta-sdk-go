package education

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// UsersItemClassesEducationClassItemRequestBuilder provides operations to manage the classes property of the microsoft.graph.educationUser entity.
type UsersItemClassesEducationClassItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// UsersItemClassesEducationClassItemRequestBuilderGetQueryParameters classes to which the user belongs. Nullable.
type UsersItemClassesEducationClassItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// UsersItemClassesEducationClassItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UsersItemClassesEducationClassItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *UsersItemClassesEducationClassItemRequestBuilderGetQueryParameters
}
// NewUsersItemClassesEducationClassItemRequestBuilderInternal instantiates a new UsersItemClassesEducationClassItemRequestBuilder and sets the default values.
func NewUsersItemClassesEducationClassItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UsersItemClassesEducationClassItemRequestBuilder) {
    m := &UsersItemClassesEducationClassItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/education/users/{educationUser%2Did}/classes/{educationClass%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewUsersItemClassesEducationClassItemRequestBuilder instantiates a new UsersItemClassesEducationClassItemRequestBuilder and sets the default values.
func NewUsersItemClassesEducationClassItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UsersItemClassesEducationClassItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUsersItemClassesEducationClassItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get classes to which the user belongs. Nullable.
// returns a EducationClassable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UsersItemClassesEducationClassItemRequestBuilder) Get(ctx context.Context, requestConfiguration *UsersItemClassesEducationClassItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EducationClassable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateEducationClassFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EducationClassable), nil
}
// ToGetRequestInformation classes to which the user belongs. Nullable.
// returns a *RequestInformation when successful
func (m *UsersItemClassesEducationClassItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *UsersItemClassesEducationClassItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *UsersItemClassesEducationClassItemRequestBuilder when successful
func (m *UsersItemClassesEducationClassItemRequestBuilder) WithUrl(rawUrl string)(*UsersItemClassesEducationClassItemRequestBuilder) {
    return NewUsersItemClassesEducationClassItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
