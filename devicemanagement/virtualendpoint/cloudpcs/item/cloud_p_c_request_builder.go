package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i13ec1c9a15a271cc8f1a63c29667b2eb16fa66895eadeae5018091e29071ea9d "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/virtualendpoint/cloudpcs/item/changeuseraccounttype"
    i170388c68b05b462f70fd530877e5751fb8c26c036b698799100779eefe3754a "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/virtualendpoint/cloudpcs/item/rename"
    i3ea7449e4e31fbb2dba44ba96548c03bf21e30fa0fbc867460aef4206de9a3ab "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/virtualendpoint/cloudpcs/item/reboot"
    i7174971b3bcb6f4ec85df581ff277b07bbaf56b391c32a3f561718b3a00346fe "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/virtualendpoint/cloudpcs/item/troubleshoot"
    ia073266fa87f992d3a858948753dc53b6ac9007c5febc8b5d74af59df63531cf "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/virtualendpoint/cloudpcs/item/reprovision"
    ic5c3712fbfc3279ed202453cc12bb33fc813457da47b3438a76bf9022c5b2f4a "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/virtualendpoint/cloudpcs/item/endgraceperiod"
)

// CloudPCRequestBuilder builds and executes requests for operations under \deviceManagement\virtualEndpoint\cloudPCs\{cloudPC-id}
type CloudPCRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// CloudPCRequestBuilderDeleteOptions options for Delete
type CloudPCRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// CloudPCRequestBuilderGetOptions options for Get
type CloudPCRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *CloudPCRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// CloudPCRequestBuilderGetQueryParameters cloud managed virtual desktops.
type CloudPCRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// CloudPCRequestBuilderPatchOptions options for Patch
type CloudPCRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CloudPC;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *CloudPCRequestBuilder) ChangeUserAccountType()(*i13ec1c9a15a271cc8f1a63c29667b2eb16fa66895eadeae5018091e29071ea9d.ChangeUserAccountTypeRequestBuilder) {
    return i13ec1c9a15a271cc8f1a63c29667b2eb16fa66895eadeae5018091e29071ea9d.NewChangeUserAccountTypeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewCloudPCRequestBuilderInternal instantiates a new CloudPCRequestBuilder and sets the default values.
func NewCloudPCRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*CloudPCRequestBuilder) {
    m := &CloudPCRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/virtualEndpoint/cloudPCs/{cloudPC_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewCloudPCRequestBuilder instantiates a new CloudPCRequestBuilder and sets the default values.
func NewCloudPCRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*CloudPCRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCloudPCRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation cloud managed virtual desktops.
func (m *CloudPCRequestBuilder) CreateDeleteRequestInformation(options *CloudPCRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.DELETE
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreateGetRequestInformation cloud managed virtual desktops.
func (m *CloudPCRequestBuilder) CreateGetRequestInformation(options *CloudPCRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        requestInfo.AddQueryParameters(*(options.Q))
    }
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation cloud managed virtual desktops.
func (m *CloudPCRequestBuilder) CreatePatchRequestInformation(options *CloudPCRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Delete cloud managed virtual desktops.
func (m *CloudPCRequestBuilder) Delete(options *CloudPCRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *CloudPCRequestBuilder) EndGracePeriod()(*ic5c3712fbfc3279ed202453cc12bb33fc813457da47b3438a76bf9022c5b2f4a.EndGracePeriodRequestBuilder) {
    return ic5c3712fbfc3279ed202453cc12bb33fc813457da47b3438a76bf9022c5b2f4a.NewEndGracePeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get cloud managed virtual desktops.
func (m *CloudPCRequestBuilder) Get(options *CloudPCRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CloudPC, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewCloudPC() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CloudPC), nil
}
// Patch cloud managed virtual desktops.
func (m *CloudPCRequestBuilder) Patch(options *CloudPCRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *CloudPCRequestBuilder) Reboot()(*i3ea7449e4e31fbb2dba44ba96548c03bf21e30fa0fbc867460aef4206de9a3ab.RebootRequestBuilder) {
    return i3ea7449e4e31fbb2dba44ba96548c03bf21e30fa0fbc867460aef4206de9a3ab.NewRebootRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CloudPCRequestBuilder) Rename()(*i170388c68b05b462f70fd530877e5751fb8c26c036b698799100779eefe3754a.RenameRequestBuilder) {
    return i170388c68b05b462f70fd530877e5751fb8c26c036b698799100779eefe3754a.NewRenameRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CloudPCRequestBuilder) Reprovision()(*ia073266fa87f992d3a858948753dc53b6ac9007c5febc8b5d74af59df63531cf.ReprovisionRequestBuilder) {
    return ia073266fa87f992d3a858948753dc53b6ac9007c5febc8b5d74af59df63531cf.NewReprovisionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CloudPCRequestBuilder) Troubleshoot()(*i7174971b3bcb6f4ec85df581ff277b07bbaf56b391c32a3f561718b3a00346fe.TroubleshootRequestBuilder) {
    return i7174971b3bcb6f4ec85df581ff277b07bbaf56b391c32a3f561718b3a00346fe.NewTroubleshootRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
