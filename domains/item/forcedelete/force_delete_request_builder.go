package forcedelete

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ForceDeleteRequestBuilder provides operations to call the forceDelete method.
type ForceDeleteRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ForceDeleteRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ForceDeleteRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewForceDeleteRequestBuilderInternal instantiates a new ForceDeleteRequestBuilder and sets the default values.
func NewForceDeleteRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ForceDeleteRequestBuilder) {
    m := &ForceDeleteRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/domains/{domain%2Did}/microsoft.graph.forceDelete";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewForceDeleteRequestBuilder instantiates a new ForceDeleteRequestBuilder and sets the default values.
func NewForceDeleteRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ForceDeleteRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewForceDeleteRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation deletes a domain using an asynchronous operation. Prior to calling forceDelete, you must update or remove any references to **Exchange** as the provisioning service. The following actions are performed as part of this operation: After the domain deletion completes, API operations for the deleted domain will return a 404 HTTP response code. To verify deletion of a domain, you can perform a get domain. If the domain was successfully deleted, a 404 HTTP response code will be returned in the response.
func (m *ForceDeleteRequestBuilder) CreatePostRequestInformation(body ForceDeletePostRequestBodyable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePostRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePostRequestInformationWithRequestConfiguration deletes a domain using an asynchronous operation. Prior to calling forceDelete, you must update or remove any references to **Exchange** as the provisioning service. The following actions are performed as part of this operation: After the domain deletion completes, API operations for the deleted domain will return a 404 HTTP response code. To verify deletion of a domain, you can perform a get domain. If the domain was successfully deleted, a 404 HTTP response code will be returned in the response.
func (m *ForceDeleteRequestBuilder) CreatePostRequestInformationWithRequestConfiguration(body ForceDeletePostRequestBodyable, requestConfiguration *ForceDeleteRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Post deletes a domain using an asynchronous operation. Prior to calling forceDelete, you must update or remove any references to **Exchange** as the provisioning service. The following actions are performed as part of this operation: After the domain deletion completes, API operations for the deleted domain will return a 404 HTTP response code. To verify deletion of a domain, you can perform a get domain. If the domain was successfully deleted, a 404 HTTP response code will be returned in the response.
func (m *ForceDeleteRequestBuilder) Post(ctx context.Context, body ForceDeletePostRequestBodyable, requestConfiguration *ForceDeleteRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.CreatePostRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
