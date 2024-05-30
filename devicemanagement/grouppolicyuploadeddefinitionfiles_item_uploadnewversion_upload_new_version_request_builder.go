package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// GrouppolicyuploadeddefinitionfilesItemUploadnewversionUploadNewVersionRequestBuilder provides operations to call the uploadNewVersion method.
type GrouppolicyuploadeddefinitionfilesItemUploadnewversionUploadNewVersionRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// GrouppolicyuploadeddefinitionfilesItemUploadnewversionUploadNewVersionRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GrouppolicyuploadeddefinitionfilesItemUploadnewversionUploadNewVersionRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewGrouppolicyuploadeddefinitionfilesItemUploadnewversionUploadNewVersionRequestBuilderInternal instantiates a new GrouppolicyuploadeddefinitionfilesItemUploadnewversionUploadNewVersionRequestBuilder and sets the default values.
func NewGrouppolicyuploadeddefinitionfilesItemUploadnewversionUploadNewVersionRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GrouppolicyuploadeddefinitionfilesItemUploadnewversionUploadNewVersionRequestBuilder) {
    m := &GrouppolicyuploadeddefinitionfilesItemUploadnewversionUploadNewVersionRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/groupPolicyUploadedDefinitionFiles/{groupPolicyUploadedDefinitionFile%2Did}/uploadNewVersion", pathParameters),
    }
    return m
}
// NewGrouppolicyuploadeddefinitionfilesItemUploadnewversionUploadNewVersionRequestBuilder instantiates a new GrouppolicyuploadeddefinitionfilesItemUploadnewversionUploadNewVersionRequestBuilder and sets the default values.
func NewGrouppolicyuploadeddefinitionfilesItemUploadnewversionUploadNewVersionRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GrouppolicyuploadeddefinitionfilesItemUploadnewversionUploadNewVersionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGrouppolicyuploadeddefinitionfilesItemUploadnewversionUploadNewVersionRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action uploadNewVersion
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *GrouppolicyuploadeddefinitionfilesItemUploadnewversionUploadNewVersionRequestBuilder) Post(ctx context.Context, body GrouppolicyuploadeddefinitionfilesItemUploadnewversionUploadNewVersionPostRequestBodyable, requestConfiguration *GrouppolicyuploadeddefinitionfilesItemUploadnewversionUploadNewVersionRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
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
// ToPostRequestInformation invoke action uploadNewVersion
// returns a *RequestInformation when successful
func (m *GrouppolicyuploadeddefinitionfilesItemUploadnewversionUploadNewVersionRequestBuilder) ToPostRequestInformation(ctx context.Context, body GrouppolicyuploadeddefinitionfilesItemUploadnewversionUploadNewVersionPostRequestBodyable, requestConfiguration *GrouppolicyuploadeddefinitionfilesItemUploadnewversionUploadNewVersionRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *GrouppolicyuploadeddefinitionfilesItemUploadnewversionUploadNewVersionRequestBuilder when successful
func (m *GrouppolicyuploadeddefinitionfilesItemUploadnewversionUploadNewVersionRequestBuilder) WithUrl(rawUrl string)(*GrouppolicyuploadeddefinitionfilesItemUploadnewversionUploadNewVersionRequestBuilder) {
    return NewGrouppolicyuploadeddefinitionfilesItemUploadnewversionUploadNewVersionRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
