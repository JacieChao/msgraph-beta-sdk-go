package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i1aa4fb44cf4940a0794e33de635a58953960c975237235af78452d02bffe2854 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/resourceaccessprofiles/item/assignments"
    iaac20666e49cb2032811a5fbdaae8facc6a32715040bd58f314cf8d5f006f5bd "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/resourceaccessprofiles/item/assign"
    i4c6649ed921058137d23369f25fb530c052e741c6026b182c9c66a0903b527cd "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/resourceaccessprofiles/item/assignments/item"
)

// DeviceManagementResourceAccessProfileBaseItemRequestBuilder builds and executes requests for operations under \deviceManagement\resourceAccessProfiles\{deviceManagementResourceAccessProfileBase-id}
type DeviceManagementResourceAccessProfileBaseItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// DeviceManagementResourceAccessProfileBaseItemRequestBuilderDeleteOptions options for Delete
type DeviceManagementResourceAccessProfileBaseItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// DeviceManagementResourceAccessProfileBaseItemRequestBuilderGetOptions options for Get
type DeviceManagementResourceAccessProfileBaseItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *DeviceManagementResourceAccessProfileBaseItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// DeviceManagementResourceAccessProfileBaseItemRequestBuilderGetQueryParameters collection of resource access settings associated with account.
type DeviceManagementResourceAccessProfileBaseItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// DeviceManagementResourceAccessProfileBaseItemRequestBuilderPatchOptions options for Patch
type DeviceManagementResourceAccessProfileBaseItemRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceManagementResourceAccessProfileBase;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *DeviceManagementResourceAccessProfileBaseItemRequestBuilder) Assign()(*iaac20666e49cb2032811a5fbdaae8facc6a32715040bd58f314cf8d5f006f5bd.AssignRequestBuilder) {
    return iaac20666e49cb2032811a5fbdaae8facc6a32715040bd58f314cf8d5f006f5bd.NewAssignRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DeviceManagementResourceAccessProfileBaseItemRequestBuilder) Assignments()(*i1aa4fb44cf4940a0794e33de635a58953960c975237235af78452d02bffe2854.AssignmentsRequestBuilder) {
    return i1aa4fb44cf4940a0794e33de635a58953960c975237235af78452d02bffe2854.NewAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.resourceAccessProfiles.item.assignments.item collection
func (m *DeviceManagementResourceAccessProfileBaseItemRequestBuilder) AssignmentsById(id string)(*i4c6649ed921058137d23369f25fb530c052e741c6026b182c9c66a0903b527cd.DeviceManagementResourceAccessProfileAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementResourceAccessProfileAssignment_id"] = id
    }
    return i4c6649ed921058137d23369f25fb530c052e741c6026b182c9c66a0903b527cd.NewDeviceManagementResourceAccessProfileAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewDeviceManagementResourceAccessProfileBaseItemRequestBuilderInternal instantiates a new DeviceManagementResourceAccessProfileBaseItemRequestBuilder and sets the default values.
func NewDeviceManagementResourceAccessProfileBaseItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DeviceManagementResourceAccessProfileBaseItemRequestBuilder) {
    m := &DeviceManagementResourceAccessProfileBaseItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/resourceAccessProfiles/{deviceManagementResourceAccessProfileBase_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceManagementResourceAccessProfileBaseItemRequestBuilder instantiates a new DeviceManagementResourceAccessProfileBaseItemRequestBuilder and sets the default values.
func NewDeviceManagementResourceAccessProfileBaseItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DeviceManagementResourceAccessProfileBaseItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceManagementResourceAccessProfileBaseItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation collection of resource access settings associated with account.
func (m *DeviceManagementResourceAccessProfileBaseItemRequestBuilder) CreateDeleteRequestInformation(options *DeviceManagementResourceAccessProfileBaseItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation collection of resource access settings associated with account.
func (m *DeviceManagementResourceAccessProfileBaseItemRequestBuilder) CreateGetRequestInformation(options *DeviceManagementResourceAccessProfileBaseItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation collection of resource access settings associated with account.
func (m *DeviceManagementResourceAccessProfileBaseItemRequestBuilder) CreatePatchRequestInformation(options *DeviceManagementResourceAccessProfileBaseItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete collection of resource access settings associated with account.
func (m *DeviceManagementResourceAccessProfileBaseItemRequestBuilder) Delete(options *DeviceManagementResourceAccessProfileBaseItemRequestBuilderDeleteOptions)(error) {
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
// Get collection of resource access settings associated with account.
func (m *DeviceManagementResourceAccessProfileBaseItemRequestBuilder) Get(options *DeviceManagementResourceAccessProfileBaseItemRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceManagementResourceAccessProfileBase, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewDeviceManagementResourceAccessProfileBase() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceManagementResourceAccessProfileBase), nil
}
// Patch collection of resource access settings associated with account.
func (m *DeviceManagementResourceAccessProfileBaseItemRequestBuilder) Patch(options *DeviceManagementResourceAccessProfileBaseItemRequestBuilderPatchOptions)(error) {
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