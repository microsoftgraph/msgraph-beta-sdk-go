package reports

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// Getoffice365groupsactivitydetailwithdateGetOffice365GroupsActivityDetailWithDateRequestBuilder provides operations to call the getOffice365GroupsActivityDetail method.
type Getoffice365groupsactivitydetailwithdateGetOffice365GroupsActivityDetailWithDateRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// Getoffice365groupsactivitydetailwithdateGetOffice365GroupsActivityDetailWithDateRequestBuilderGetQueryParameters invoke function getOffice365GroupsActivityDetail
type Getoffice365groupsactivitydetailwithdateGetOffice365GroupsActivityDetailWithDateRequestBuilderGetQueryParameters struct {
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
// Getoffice365groupsactivitydetailwithdateGetOffice365GroupsActivityDetailWithDateRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type Getoffice365groupsactivitydetailwithdateGetOffice365GroupsActivityDetailWithDateRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *Getoffice365groupsactivitydetailwithdateGetOffice365GroupsActivityDetailWithDateRequestBuilderGetQueryParameters
}
// NewGetoffice365groupsactivitydetailwithdateGetOffice365GroupsActivityDetailWithDateRequestBuilderInternal instantiates a new Getoffice365groupsactivitydetailwithdateGetOffice365GroupsActivityDetailWithDateRequestBuilder and sets the default values.
func NewGetoffice365groupsactivitydetailwithdateGetOffice365GroupsActivityDetailWithDateRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)(*Getoffice365groupsactivitydetailwithdateGetOffice365GroupsActivityDetailWithDateRequestBuilder) {
    m := &Getoffice365groupsactivitydetailwithdateGetOffice365GroupsActivityDetailWithDateRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/reports/getOffice365GroupsActivityDetail(date={date}){?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    if date != nil {
        m.BaseRequestBuilder.PathParameters["date"] = (*date).String()
    }
    return m
}
// NewGetoffice365groupsactivitydetailwithdateGetOffice365GroupsActivityDetailWithDateRequestBuilder instantiates a new Getoffice365groupsactivitydetailwithdateGetOffice365GroupsActivityDetailWithDateRequestBuilder and sets the default values.
func NewGetoffice365groupsactivitydetailwithdateGetOffice365GroupsActivityDetailWithDateRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*Getoffice365groupsactivitydetailwithdateGetOffice365GroupsActivityDetailWithDateRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetoffice365groupsactivitydetailwithdateGetOffice365GroupsActivityDetailWithDateRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get invoke function getOffice365GroupsActivityDetail
// Deprecated: This method is obsolete. Use GetAsGetOffice365GroupsActivityDetailWithDateGetResponse instead.
// returns a Getoffice365groupsactivitydetailwithdateGetOffice365GroupsActivityDetailWithDateResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *Getoffice365groupsactivitydetailwithdateGetOffice365GroupsActivityDetailWithDateRequestBuilder) Get(ctx context.Context, requestConfiguration *Getoffice365groupsactivitydetailwithdateGetOffice365GroupsActivityDetailWithDateRequestBuilderGetRequestConfiguration)(Getoffice365groupsactivitydetailwithdateGetOffice365GroupsActivityDetailWithDateResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateGetoffice365groupsactivitydetailwithdateGetOffice365GroupsActivityDetailWithDateResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(Getoffice365groupsactivitydetailwithdateGetOffice365GroupsActivityDetailWithDateResponseable), nil
}
// GetAsGetOffice365GroupsActivityDetailWithDateGetResponse invoke function getOffice365GroupsActivityDetail
// returns a Getoffice365groupsactivitydetailwithdateGetOffice365GroupsActivityDetailWithDateGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *Getoffice365groupsactivitydetailwithdateGetOffice365GroupsActivityDetailWithDateRequestBuilder) GetAsGetOffice365GroupsActivityDetailWithDateGetResponse(ctx context.Context, requestConfiguration *Getoffice365groupsactivitydetailwithdateGetOffice365GroupsActivityDetailWithDateRequestBuilderGetRequestConfiguration)(Getoffice365groupsactivitydetailwithdateGetOffice365GroupsActivityDetailWithDateGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateGetoffice365groupsactivitydetailwithdateGetOffice365GroupsActivityDetailWithDateGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(Getoffice365groupsactivitydetailwithdateGetOffice365GroupsActivityDetailWithDateGetResponseable), nil
}
// ToGetRequestInformation invoke function getOffice365GroupsActivityDetail
// returns a *RequestInformation when successful
func (m *Getoffice365groupsactivitydetailwithdateGetOffice365GroupsActivityDetailWithDateRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *Getoffice365groupsactivitydetailwithdateGetOffice365GroupsActivityDetailWithDateRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *Getoffice365groupsactivitydetailwithdateGetOffice365GroupsActivityDetailWithDateRequestBuilder when successful
func (m *Getoffice365groupsactivitydetailwithdateGetOffice365GroupsActivityDetailWithDateRequestBuilder) WithUrl(rawUrl string)(*Getoffice365groupsactivitydetailwithdateGetOffice365GroupsActivityDetailWithDateRequestBuilder) {
    return NewGetoffice365groupsactivitydetailwithdateGetOffice365GroupsActivityDetailWithDateRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
