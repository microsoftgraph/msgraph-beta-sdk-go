package hascustomrolescopetag

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// HasCustomRoleScopeTagRequestBuilder provides operations to call the hasCustomRoleScopeTag method.
type HasCustomRoleScopeTagRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// HasCustomRoleScopeTagRequestBuilderGetOptions options for Get
type HasCustomRoleScopeTagRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// NewHasCustomRoleScopeTagRequestBuilderInternal instantiates a new HasCustomRoleScopeTagRequestBuilder and sets the default values.
func NewHasCustomRoleScopeTagRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*HasCustomRoleScopeTagRequestBuilder) {
    m := &HasCustomRoleScopeTagRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/roleScopeTags/microsoft.graph.hasCustomRoleScopeTag()";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewHasCustomRoleScopeTagRequestBuilder instantiates a new HasCustomRoleScopeTagRequestBuilder and sets the default values.
func NewHasCustomRoleScopeTagRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*HasCustomRoleScopeTagRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewHasCustomRoleScopeTagRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation invoke function hasCustomRoleScopeTag
func (m *HasCustomRoleScopeTagRequestBuilder) CreateGetRequestInformation(options *HasCustomRoleScopeTagRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Get invoke function hasCustomRoleScopeTag
func (m *HasCustomRoleScopeTagRequestBuilder) Get(options *HasCustomRoleScopeTagRequestBuilderGetOptions)(HasCustomRoleScopeTagResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, CreateHasCustomRoleScopeTagResponseFromDiscriminatorValue, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(HasCustomRoleScopeTagResponseable), nil
}
