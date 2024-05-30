package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ComanageddevicesItemRecoverpasscodeRecoverPasscodeRequestBuilder provides operations to call the recoverPasscode method.
type ComanageddevicesItemRecoverpasscodeRecoverPasscodeRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ComanageddevicesItemRecoverpasscodeRecoverPasscodeRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ComanageddevicesItemRecoverpasscodeRecoverPasscodeRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewComanageddevicesItemRecoverpasscodeRecoverPasscodeRequestBuilderInternal instantiates a new ComanageddevicesItemRecoverpasscodeRecoverPasscodeRequestBuilder and sets the default values.
func NewComanageddevicesItemRecoverpasscodeRecoverPasscodeRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ComanageddevicesItemRecoverpasscodeRecoverPasscodeRequestBuilder) {
    m := &ComanageddevicesItemRecoverpasscodeRecoverPasscodeRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/comanagedDevices/{managedDevice%2Did}/recoverPasscode", pathParameters),
    }
    return m
}
// NewComanageddevicesItemRecoverpasscodeRecoverPasscodeRequestBuilder instantiates a new ComanageddevicesItemRecoverpasscodeRecoverPasscodeRequestBuilder and sets the default values.
func NewComanageddevicesItemRecoverpasscodeRecoverPasscodeRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ComanageddevicesItemRecoverpasscodeRecoverPasscodeRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewComanageddevicesItemRecoverpasscodeRecoverPasscodeRequestBuilderInternal(urlParams, requestAdapter)
}
// Post recover passcode
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ComanageddevicesItemRecoverpasscodeRecoverPasscodeRequestBuilder) Post(ctx context.Context, requestConfiguration *ComanageddevicesItemRecoverpasscodeRecoverPasscodeRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
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
// ToPostRequestInformation recover passcode
// returns a *RequestInformation when successful
func (m *ComanageddevicesItemRecoverpasscodeRecoverPasscodeRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ComanageddevicesItemRecoverpasscodeRecoverPasscodeRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ComanageddevicesItemRecoverpasscodeRecoverPasscodeRequestBuilder when successful
func (m *ComanageddevicesItemRecoverpasscodeRecoverPasscodeRequestBuilder) WithUrl(rawUrl string)(*ComanageddevicesItemRecoverpasscodeRecoverPasscodeRequestBuilder) {
    return NewComanageddevicesItemRecoverpasscodeRecoverPasscodeRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
