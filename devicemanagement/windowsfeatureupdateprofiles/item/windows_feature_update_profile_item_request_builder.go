package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i17ae4c396051acc3d8969fe0d510a15f08f74b0af97608557d2125c16de1bb7d "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/windowsfeatureupdateprofiles/item/assign"
    ib9f19a8a1ed79d6335dbd9524123a1a111e94345212d4f98f256d4d77fb64865 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/windowsfeatureupdateprofiles/item/assignments"
    i9fc8cb618d1d2cb2bd52c185cd35fcf810ea303a5692d3f159788228e6d3514f "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/windowsfeatureupdateprofiles/item/assignments/item"
)

// WindowsFeatureUpdateProfileItemRequestBuilder builds and executes requests for operations under \deviceManagement\windowsFeatureUpdateProfiles\{windowsFeatureUpdateProfile-id}
type WindowsFeatureUpdateProfileItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// WindowsFeatureUpdateProfileItemRequestBuilderDeleteOptions options for Delete
type WindowsFeatureUpdateProfileItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// WindowsFeatureUpdateProfileItemRequestBuilderGetOptions options for Get
type WindowsFeatureUpdateProfileItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *WindowsFeatureUpdateProfileItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// WindowsFeatureUpdateProfileItemRequestBuilderGetQueryParameters a collection of windows feature update profiles
type WindowsFeatureUpdateProfileItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// WindowsFeatureUpdateProfileItemRequestBuilderPatchOptions options for Patch
type WindowsFeatureUpdateProfileItemRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.WindowsFeatureUpdateProfile;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *WindowsFeatureUpdateProfileItemRequestBuilder) Assign()(*i17ae4c396051acc3d8969fe0d510a15f08f74b0af97608557d2125c16de1bb7d.AssignRequestBuilder) {
    return i17ae4c396051acc3d8969fe0d510a15f08f74b0af97608557d2125c16de1bb7d.NewAssignRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WindowsFeatureUpdateProfileItemRequestBuilder) Assignments()(*ib9f19a8a1ed79d6335dbd9524123a1a111e94345212d4f98f256d4d77fb64865.AssignmentsRequestBuilder) {
    return ib9f19a8a1ed79d6335dbd9524123a1a111e94345212d4f98f256d4d77fb64865.NewAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.windowsFeatureUpdateProfiles.item.assignments.item collection
func (m *WindowsFeatureUpdateProfileItemRequestBuilder) AssignmentsById(id string)(*i9fc8cb618d1d2cb2bd52c185cd35fcf810ea303a5692d3f159788228e6d3514f.WindowsFeatureUpdateProfileAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["windowsFeatureUpdateProfileAssignment_id"] = id
    }
    return i9fc8cb618d1d2cb2bd52c185cd35fcf810ea303a5692d3f159788228e6d3514f.NewWindowsFeatureUpdateProfileAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewWindowsFeatureUpdateProfileItemRequestBuilderInternal instantiates a new WindowsFeatureUpdateProfileItemRequestBuilder and sets the default values.
func NewWindowsFeatureUpdateProfileItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*WindowsFeatureUpdateProfileItemRequestBuilder) {
    m := &WindowsFeatureUpdateProfileItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/windowsFeatureUpdateProfiles/{windowsFeatureUpdateProfile_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewWindowsFeatureUpdateProfileItemRequestBuilder instantiates a new WindowsFeatureUpdateProfileItemRequestBuilder and sets the default values.
func NewWindowsFeatureUpdateProfileItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*WindowsFeatureUpdateProfileItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWindowsFeatureUpdateProfileItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation a collection of windows feature update profiles
func (m *WindowsFeatureUpdateProfileItemRequestBuilder) CreateDeleteRequestInformation(options *WindowsFeatureUpdateProfileItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation a collection of windows feature update profiles
func (m *WindowsFeatureUpdateProfileItemRequestBuilder) CreateGetRequestInformation(options *WindowsFeatureUpdateProfileItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation a collection of windows feature update profiles
func (m *WindowsFeatureUpdateProfileItemRequestBuilder) CreatePatchRequestInformation(options *WindowsFeatureUpdateProfileItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete a collection of windows feature update profiles
func (m *WindowsFeatureUpdateProfileItemRequestBuilder) Delete(options *WindowsFeatureUpdateProfileItemRequestBuilderDeleteOptions)(error) {
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
// Get a collection of windows feature update profiles
func (m *WindowsFeatureUpdateProfileItemRequestBuilder) Get(options *WindowsFeatureUpdateProfileItemRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.WindowsFeatureUpdateProfile, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewWindowsFeatureUpdateProfile() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.WindowsFeatureUpdateProfile), nil
}
// Patch a collection of windows feature update profiles
func (m *WindowsFeatureUpdateProfileItemRequestBuilder) Patch(options *WindowsFeatureUpdateProfileItemRequestBuilderPatchOptions)(error) {
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