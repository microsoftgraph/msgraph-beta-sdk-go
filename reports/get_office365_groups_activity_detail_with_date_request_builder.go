package reports

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// GetOffice365GroupsActivityDetailWithDateRequestBuilder provides operations to call the getOffice365GroupsActivityDetail method.
type GetOffice365GroupsActivityDetailWithDateRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// GetOffice365GroupsActivityDetailWithDateRequestBuilderGetQueryParameters invoke function getOffice365GroupsActivityDetail
type GetOffice365GroupsActivityDetailWithDateRequestBuilderGetQueryParameters struct {
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
// GetOffice365GroupsActivityDetailWithDateRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GetOffice365GroupsActivityDetailWithDateRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *GetOffice365GroupsActivityDetailWithDateRequestBuilderGetQueryParameters
}
// NewGetOffice365GroupsActivityDetailWithDateRequestBuilderInternal instantiates a new GetOffice365GroupsActivityDetailWithDateRequestBuilder and sets the default values.
func NewGetOffice365GroupsActivityDetailWithDateRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)(*GetOffice365GroupsActivityDetailWithDateRequestBuilder) {
    m := &GetOffice365GroupsActivityDetailWithDateRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/reports/getOffice365GroupsActivityDetail(date={date}){?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    if date != nil {
        m.BaseRequestBuilder.PathParameters["date"] = (*date).String()
    }
    return m
}
// NewGetOffice365GroupsActivityDetailWithDateRequestBuilder instantiates a new GetOffice365GroupsActivityDetailWithDateRequestBuilder and sets the default values.
func NewGetOffice365GroupsActivityDetailWithDateRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetOffice365GroupsActivityDetailWithDateRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetOffice365GroupsActivityDetailWithDateRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get invoke function getOffice365GroupsActivityDetail
// Deprecated: This method is obsolete. Use {TypeName} instead.
// returns a GetOffice365GroupsActivityDetailWithDateResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *GetOffice365GroupsActivityDetailWithDateRequestBuilder) Get(ctx context.Context, requestConfiguration *GetOffice365GroupsActivityDetailWithDateRequestBuilderGetRequestConfiguration)(GetOffice365GroupsActivityDetailWithDateResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateGetOffice365GroupsActivityDetailWithDateResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(GetOffice365GroupsActivityDetailWithDateResponseable), nil
}
// GetAsGetOffice365GroupsActivityDetailWithDateGetResponse invoke function getOffice365GroupsActivityDetail
// returns a GetOffice365GroupsActivityDetailWithDateGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *GetOffice365GroupsActivityDetailWithDateRequestBuilder) GetAsGetOffice365GroupsActivityDetailWithDateGetResponse(ctx context.Context, requestConfiguration *GetOffice365GroupsActivityDetailWithDateRequestBuilderGetRequestConfiguration)(GetOffice365GroupsActivityDetailWithDateGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateGetOffice365GroupsActivityDetailWithDateGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(GetOffice365GroupsActivityDetailWithDateGetResponseable), nil
}
// ToGetRequestInformation invoke function getOffice365GroupsActivityDetail
// returns a *RequestInformation when successful
func (m *GetOffice365GroupsActivityDetailWithDateRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *GetOffice365GroupsActivityDetailWithDateRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *GetOffice365GroupsActivityDetailWithDateRequestBuilder when successful
func (m *GetOffice365GroupsActivityDetailWithDateRequestBuilder) WithUrl(rawUrl string)(*GetOffice365GroupsActivityDetailWithDateRequestBuilder) {
    return NewGetOffice365GroupsActivityDetailWithDateRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
