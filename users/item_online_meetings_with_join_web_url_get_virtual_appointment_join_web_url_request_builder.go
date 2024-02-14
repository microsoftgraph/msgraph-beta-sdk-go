package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemOnlineMeetingsWithJoinWebUrlGetVirtualAppointmentJoinWebUrlRequestBuilder provides operations to call the getVirtualAppointmentJoinWebUrl method.
type ItemOnlineMeetingsWithJoinWebUrlGetVirtualAppointmentJoinWebUrlRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemOnlineMeetingsWithJoinWebUrlGetVirtualAppointmentJoinWebUrlRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemOnlineMeetingsWithJoinWebUrlGetVirtualAppointmentJoinWebUrlRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemOnlineMeetingsWithJoinWebUrlGetVirtualAppointmentJoinWebUrlRequestBuilderInternal instantiates a new ItemOnlineMeetingsWithJoinWebUrlGetVirtualAppointmentJoinWebUrlRequestBuilder and sets the default values.
func NewItemOnlineMeetingsWithJoinWebUrlGetVirtualAppointmentJoinWebUrlRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemOnlineMeetingsWithJoinWebUrlGetVirtualAppointmentJoinWebUrlRequestBuilder) {
    m := &ItemOnlineMeetingsWithJoinWebUrlGetVirtualAppointmentJoinWebUrlRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/onlineMeetings(joinWebUrl='{joinWebUrl}')/getVirtualAppointmentJoinWebUrl()", pathParameters),
    }
    return m
}
// NewItemOnlineMeetingsWithJoinWebUrlGetVirtualAppointmentJoinWebUrlRequestBuilder instantiates a new ItemOnlineMeetingsWithJoinWebUrlGetVirtualAppointmentJoinWebUrlRequestBuilder and sets the default values.
func NewItemOnlineMeetingsWithJoinWebUrlGetVirtualAppointmentJoinWebUrlRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemOnlineMeetingsWithJoinWebUrlGetVirtualAppointmentJoinWebUrlRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemOnlineMeetingsWithJoinWebUrlGetVirtualAppointmentJoinWebUrlRequestBuilderInternal(urlParams, requestAdapter)
}
// Get invoke function getVirtualAppointmentJoinWebUrl
// Deprecated: This method is obsolete. Use {TypeName} instead.
// returns a ItemOnlineMeetingsWithJoinWebUrlGetVirtualAppointmentJoinWebUrlResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemOnlineMeetingsWithJoinWebUrlGetVirtualAppointmentJoinWebUrlRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemOnlineMeetingsWithJoinWebUrlGetVirtualAppointmentJoinWebUrlRequestBuilderGetRequestConfiguration)(ItemOnlineMeetingsWithJoinWebUrlGetVirtualAppointmentJoinWebUrlResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemOnlineMeetingsWithJoinWebUrlGetVirtualAppointmentJoinWebUrlResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemOnlineMeetingsWithJoinWebUrlGetVirtualAppointmentJoinWebUrlResponseable), nil
}
// GetAsGetVirtualAppointmentJoinWebUrlGetResponse invoke function getVirtualAppointmentJoinWebUrl
// returns a ItemOnlineMeetingsWithJoinWebUrlGetVirtualAppointmentJoinWebUrlGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemOnlineMeetingsWithJoinWebUrlGetVirtualAppointmentJoinWebUrlRequestBuilder) GetAsGetVirtualAppointmentJoinWebUrlGetResponse(ctx context.Context, requestConfiguration *ItemOnlineMeetingsWithJoinWebUrlGetVirtualAppointmentJoinWebUrlRequestBuilderGetRequestConfiguration)(ItemOnlineMeetingsWithJoinWebUrlGetVirtualAppointmentJoinWebUrlGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemOnlineMeetingsWithJoinWebUrlGetVirtualAppointmentJoinWebUrlGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemOnlineMeetingsWithJoinWebUrlGetVirtualAppointmentJoinWebUrlGetResponseable), nil
}
// ToGetRequestInformation invoke function getVirtualAppointmentJoinWebUrl
// returns a *RequestInformation when successful
func (m *ItemOnlineMeetingsWithJoinWebUrlGetVirtualAppointmentJoinWebUrlRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemOnlineMeetingsWithJoinWebUrlGetVirtualAppointmentJoinWebUrlRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemOnlineMeetingsWithJoinWebUrlGetVirtualAppointmentJoinWebUrlRequestBuilder when successful
func (m *ItemOnlineMeetingsWithJoinWebUrlGetVirtualAppointmentJoinWebUrlRequestBuilder) WithUrl(rawUrl string)(*ItemOnlineMeetingsWithJoinWebUrlGetVirtualAppointmentJoinWebUrlRequestBuilder) {
    return NewItemOnlineMeetingsWithJoinWebUrlGetVirtualAppointmentJoinWebUrlRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
