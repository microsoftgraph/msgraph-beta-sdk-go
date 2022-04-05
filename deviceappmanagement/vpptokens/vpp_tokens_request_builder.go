package vpptokens

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i24db185bc1635f2b2be054ec817ced066302c11adf27d59692cd7b262f32698e "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/vpptokens/getlicensesforappwithbundleid"
    ia528e1063400aadddf50613ea016bf5d44eacef1c362b6d342b7013823ef491f "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/vpptokens/count"
    ib956f45626ba71f8b797fc4b16070658a39a611e2d6aa111aea4eca180955477 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/vpptokens/synclicensecounts"
)

// VppTokensRequestBuilder provides operations to manage the vppTokens property of the microsoft.graph.deviceAppManagement entity.
type VppTokensRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// VppTokensRequestBuilderGetOptions options for Get
type VppTokensRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Request query parameters
    QueryParameters *VppTokensRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// VppTokensRequestBuilderGetQueryParameters list of Vpp tokens for this organization.
type VppTokensRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool;
    // Expand related entities
    Expand []string;
    // Filter items by property values
    Filter *string;
    // Order items by property values
    Orderby []string;
    // Search items by search phrases
    Search *string;
    // Select properties to be returned
    Select []string;
    // Skip the first n items
    Skip *int32;
    // Show only the first n items
    Top *int32;
}
// VppTokensRequestBuilderPostOptions options for Post
type VppTokensRequestBuilderPostOptions struct {
    // 
    Body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.VppTokenable;
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// NewVppTokensRequestBuilderInternal instantiates a new VppTokensRequestBuilder and sets the default values.
func NewVppTokensRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VppTokensRequestBuilder) {
    m := &VppTokensRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceAppManagement/vppTokens{?top,skip,search,filter,count,orderby,select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewVppTokensRequestBuilder instantiates a new VppTokensRequestBuilder and sets the default values.
func NewVppTokensRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VppTokensRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewVppTokensRequestBuilderInternal(urlParams, requestAdapter)
}
// Count the count property
func (m *VppTokensRequestBuilder) Count()(*ia528e1063400aadddf50613ea016bf5d44eacef1c362b6d342b7013823ef491f.CountRequestBuilder) {
    return ia528e1063400aadddf50613ea016bf5d44eacef1c362b6d342b7013823ef491f.NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation list of Vpp tokens for this organization.
func (m *VppTokensRequestBuilder) CreateGetRequestInformation(options *VppTokensRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePostRequestInformation create new navigation property to vppTokens for deviceAppManagement
func (m *VppTokensRequestBuilder) CreatePostRequestInformation(options *VppTokensRequestBuilderPostOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
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
// Get list of Vpp tokens for this organization.
func (m *VppTokensRequestBuilder) Get(options *VppTokensRequestBuilderGetOptions)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.VppTokenCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateVppTokenCollectionResponseFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.VppTokenCollectionResponseable), nil
}
// GetLicensesForAppWithBundleId provides operations to call the getLicensesForApp method.
func (m *VppTokensRequestBuilder) GetLicensesForAppWithBundleId(bundleId *string)(*i24db185bc1635f2b2be054ec817ced066302c11adf27d59692cd7b262f32698e.GetLicensesForAppWithBundleIdRequestBuilder) {
    return i24db185bc1635f2b2be054ec817ced066302c11adf27d59692cd7b262f32698e.NewGetLicensesForAppWithBundleIdRequestBuilderInternal(m.pathParameters, m.requestAdapter, bundleId);
}
// Post create new navigation property to vppTokens for deviceAppManagement
func (m *VppTokensRequestBuilder) Post(options *VppTokensRequestBuilderPostOptions)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.VppTokenable, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateVppTokenFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.VppTokenable), nil
}
// SyncLicenseCounts the syncLicenseCounts property
func (m *VppTokensRequestBuilder) SyncLicenseCounts()(*ib956f45626ba71f8b797fc4b16070658a39a611e2d6aa111aea4eca180955477.SyncLicenseCountsRequestBuilder) {
    return ib956f45626ba71f8b797fc4b16070658a39a611e2d6aa111aea4eca180955477.NewSyncLicenseCountsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
