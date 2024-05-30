package reports

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// GetskypeforbusinessdeviceusageuserdetailwithdateGetSkypeForBusinessDeviceUsageUserDetailWithDateRequestBuilder provides operations to call the getSkypeForBusinessDeviceUsageUserDetail method.
type GetskypeforbusinessdeviceusageuserdetailwithdateGetSkypeForBusinessDeviceUsageUserDetailWithDateRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// GetskypeforbusinessdeviceusageuserdetailwithdateGetSkypeForBusinessDeviceUsageUserDetailWithDateRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GetskypeforbusinessdeviceusageuserdetailwithdateGetSkypeForBusinessDeviceUsageUserDetailWithDateRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewGetskypeforbusinessdeviceusageuserdetailwithdateGetSkypeForBusinessDeviceUsageUserDetailWithDateRequestBuilderInternal instantiates a new GetskypeforbusinessdeviceusageuserdetailwithdateGetSkypeForBusinessDeviceUsageUserDetailWithDateRequestBuilder and sets the default values.
func NewGetskypeforbusinessdeviceusageuserdetailwithdateGetSkypeForBusinessDeviceUsageUserDetailWithDateRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)(*GetskypeforbusinessdeviceusageuserdetailwithdateGetSkypeForBusinessDeviceUsageUserDetailWithDateRequestBuilder) {
    m := &GetskypeforbusinessdeviceusageuserdetailwithdateGetSkypeForBusinessDeviceUsageUserDetailWithDateRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/reports/getSkypeForBusinessDeviceUsageUserDetail(date={date})", pathParameters),
    }
    if date != nil {
        m.BaseRequestBuilder.PathParameters["date"] = (*date).String()
    }
    return m
}
// NewGetskypeforbusinessdeviceusageuserdetailwithdateGetSkypeForBusinessDeviceUsageUserDetailWithDateRequestBuilder instantiates a new GetskypeforbusinessdeviceusageuserdetailwithdateGetSkypeForBusinessDeviceUsageUserDetailWithDateRequestBuilder and sets the default values.
func NewGetskypeforbusinessdeviceusageuserdetailwithdateGetSkypeForBusinessDeviceUsageUserDetailWithDateRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetskypeforbusinessdeviceusageuserdetailwithdateGetSkypeForBusinessDeviceUsageUserDetailWithDateRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetskypeforbusinessdeviceusageuserdetailwithdateGetSkypeForBusinessDeviceUsageUserDetailWithDateRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get invoke function getSkypeForBusinessDeviceUsageUserDetail
// returns a []byte when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *GetskypeforbusinessdeviceusageuserdetailwithdateGetSkypeForBusinessDeviceUsageUserDetailWithDateRequestBuilder) Get(ctx context.Context, requestConfiguration *GetskypeforbusinessdeviceusageuserdetailwithdateGetSkypeForBusinessDeviceUsageUserDetailWithDateRequestBuilderGetRequestConfiguration)([]byte, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendPrimitive(ctx, requestInfo, "[]byte", errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.([]byte), nil
}
// ToGetRequestInformation invoke function getSkypeForBusinessDeviceUsageUserDetail
// returns a *RequestInformation when successful
func (m *GetskypeforbusinessdeviceusageuserdetailwithdateGetSkypeForBusinessDeviceUsageUserDetailWithDateRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *GetskypeforbusinessdeviceusageuserdetailwithdateGetSkypeForBusinessDeviceUsageUserDetailWithDateRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/octet-stream, application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *GetskypeforbusinessdeviceusageuserdetailwithdateGetSkypeForBusinessDeviceUsageUserDetailWithDateRequestBuilder when successful
func (m *GetskypeforbusinessdeviceusageuserdetailwithdateGetSkypeForBusinessDeviceUsageUserDetailWithDateRequestBuilder) WithUrl(rawUrl string)(*GetskypeforbusinessdeviceusageuserdetailwithdateGetSkypeForBusinessDeviceUsageUserDetailWithDateRequestBuilder) {
    return NewGetskypeforbusinessdeviceusageuserdetailwithdateGetSkypeForBusinessDeviceUsageUserDetailWithDateRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
