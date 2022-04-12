package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i07d40583df1b4aa1b0a3e7e32084dc8b4ab30be70327b272f2ad24d3ccd4095e "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/cloudpcs/item/troubleshoot"
    i97a56d161f68b1fd3ad12ad108342214d2bfeedae4b18c9748dc91bc1e04df9f "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/cloudpcs/item/changeuseraccounttype"
    i97b00529155a730b3ecc9a3f5590a17060250dc44e1d660aed1789b4345f43c0 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/cloudpcs/item/reprovision"
    iabd9a09efde7cff52110da1822c7742e59aeb1e5915ccfe038d929a6b7ad268b "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/cloudpcs/item/endgraceperiod"
    iecadcd2b05f98b70514e4bc46dddb17b2c7871f504fd101f5951b02df655c022 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/cloudpcs/item/rename"
    ief59a775175af2e50c2ba2215af92a5bdc70b7dceef0de8704dccb1033d7ed7e "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/cloudpcs/item/reboot"
    if4dc1d7e15dcd2713c85cab00a5a53d94f288c3f116521764c65b1ad462b0eb0 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/cloudpcs/item/getcloudpclaunchinfo"
)

// CloudPCItemRequestBuilder provides operations to manage the cloudPCs property of the microsoft.graph.user entity.
type CloudPCItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// CloudPCItemRequestBuilderDeleteOptions options for Delete
type CloudPCItemRequestBuilderDeleteOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// CloudPCItemRequestBuilderGetOptions options for Get
type CloudPCItemRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CloudPCItemRequestBuilderGetQueryParameters
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// CloudPCItemRequestBuilderGetQueryParameters get cloudPCs from users
type CloudPCItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// CloudPCItemRequestBuilderPatchOptions options for Patch
type CloudPCItemRequestBuilderPatchOptions struct {
    // 
    Body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPCable
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// ChangeUserAccountType the changeUserAccountType property
func (m *CloudPCItemRequestBuilder) ChangeUserAccountType()(*i97a56d161f68b1fd3ad12ad108342214d2bfeedae4b18c9748dc91bc1e04df9f.ChangeUserAccountTypeRequestBuilder) {
    return i97a56d161f68b1fd3ad12ad108342214d2bfeedae4b18c9748dc91bc1e04df9f.NewChangeUserAccountTypeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewCloudPCItemRequestBuilderInternal instantiates a new CloudPCItemRequestBuilder and sets the default values.
func NewCloudPCItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CloudPCItemRequestBuilder) {
    m := &CloudPCItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/cloudPCs/{cloudPC%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewCloudPCItemRequestBuilder instantiates a new CloudPCItemRequestBuilder and sets the default values.
func NewCloudPCItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CloudPCItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCloudPCItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property cloudPCs for users
func (m *CloudPCItemRequestBuilder) CreateDeleteRequestInformation(options *CloudPCItemRequestBuilderDeleteOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
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
// CreateGetRequestInformation get cloudPCs from users
func (m *CloudPCItemRequestBuilder) CreateGetRequestInformation(options *CloudPCItemRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if options != nil && options.QueryParameters != nil {
        requestInfo.AddQueryParameters(*(options.QueryParameters))
    }
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
// CreatePatchRequestInformation update the navigation property cloudPCs in users
func (m *CloudPCItemRequestBuilder) CreatePatchRequestInformation(options *CloudPCItemRequestBuilderPatchOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
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
// Delete delete navigation property cloudPCs for users
func (m *CloudPCItemRequestBuilder) Delete(options *CloudPCItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// EndGracePeriod the endGracePeriod property
func (m *CloudPCItemRequestBuilder) EndGracePeriod()(*iabd9a09efde7cff52110da1822c7742e59aeb1e5915ccfe038d929a6b7ad268b.EndGracePeriodRequestBuilder) {
    return iabd9a09efde7cff52110da1822c7742e59aeb1e5915ccfe038d929a6b7ad268b.NewEndGracePeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get cloudPCs from users
func (m *CloudPCItemRequestBuilder) Get(options *CloudPCItemRequestBuilderGetOptions)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPCable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateCloudPCFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPCable), nil
}
// GetCloudPcLaunchInfo provides operations to call the getCloudPcLaunchInfo method.
func (m *CloudPCItemRequestBuilder) GetCloudPcLaunchInfo()(*if4dc1d7e15dcd2713c85cab00a5a53d94f288c3f116521764c65b1ad462b0eb0.GetCloudPcLaunchInfoRequestBuilder) {
    return if4dc1d7e15dcd2713c85cab00a5a53d94f288c3f116521764c65b1ad462b0eb0.NewGetCloudPcLaunchInfoRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property cloudPCs in users
func (m *CloudPCItemRequestBuilder) Patch(options *CloudPCItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Reboot the reboot property
func (m *CloudPCItemRequestBuilder) Reboot()(*ief59a775175af2e50c2ba2215af92a5bdc70b7dceef0de8704dccb1033d7ed7e.RebootRequestBuilder) {
    return ief59a775175af2e50c2ba2215af92a5bdc70b7dceef0de8704dccb1033d7ed7e.NewRebootRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Rename the rename property
func (m *CloudPCItemRequestBuilder) Rename()(*iecadcd2b05f98b70514e4bc46dddb17b2c7871f504fd101f5951b02df655c022.RenameRequestBuilder) {
    return iecadcd2b05f98b70514e4bc46dddb17b2c7871f504fd101f5951b02df655c022.NewRenameRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Reprovision the reprovision property
func (m *CloudPCItemRequestBuilder) Reprovision()(*i97b00529155a730b3ecc9a3f5590a17060250dc44e1d660aed1789b4345f43c0.ReprovisionRequestBuilder) {
    return i97b00529155a730b3ecc9a3f5590a17060250dc44e1d660aed1789b4345f43c0.NewReprovisionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Troubleshoot the troubleshoot property
func (m *CloudPCItemRequestBuilder) Troubleshoot()(*i07d40583df1b4aa1b0a3e7e32084dc8b4ab30be70327b272f2ad24d3ccd4095e.TroubleshootRequestBuilder) {
    return i07d40583df1b4aa1b0a3e7e32084dc8b4ab30be70327b272f2ad24d3ccd4095e.NewTroubleshootRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
