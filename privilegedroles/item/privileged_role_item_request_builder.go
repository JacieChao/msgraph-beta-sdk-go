package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i04754c5841c9fcefc803d8096b9d648d4a4f8e797cd7d5703b0b2fa021349d4f "github.com/microsoftgraph/msgraph-beta-sdk-go/privilegedroles/item/settings"
    i30870af67a7c7e750abaa30f5e84eff1a098db84ecc2ae53c8193d533bc3e771 "github.com/microsoftgraph/msgraph-beta-sdk-go/privilegedroles/item/assignments"
    ib362bf055682cb60eebaecadb64541d74da0307dbd607faad741d862fab711bf "github.com/microsoftgraph/msgraph-beta-sdk-go/privilegedroles/item/summary"
    ib4b9c3bdee734abec86c1274b61654a23bf76937e3c1b4deb6b3a5baf179dcbe "github.com/microsoftgraph/msgraph-beta-sdk-go/privilegedroles/item/selfdeactivate"
    ifc2ecb6356f4c50497a05265b4b0a311af52bfc4aa6302f476a92d1df6adfafb "github.com/microsoftgraph/msgraph-beta-sdk-go/privilegedroles/item/selfactivate"
    i0aae3989cf245a50ac6c762e9583dab23911e23dcb918057ff75fd93264ed4ca "github.com/microsoftgraph/msgraph-beta-sdk-go/privilegedroles/item/assignments/item"
)

// PrivilegedRoleItemRequestBuilder provides operations to manage the collection of privilegedRole entities.
type PrivilegedRoleItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// PrivilegedRoleItemRequestBuilderDeleteOptions options for Delete
type PrivilegedRoleItemRequestBuilderDeleteOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// PrivilegedRoleItemRequestBuilderGetOptions options for Get
type PrivilegedRoleItemRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Request query parameters
    QueryParameters *PrivilegedRoleItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// PrivilegedRoleItemRequestBuilderGetQueryParameters get entity from privilegedRoles by key
type PrivilegedRoleItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// PrivilegedRoleItemRequestBuilderPatchOptions options for Patch
type PrivilegedRoleItemRequestBuilderPatchOptions struct {
    // 
    Body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PrivilegedRoleable;
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// Assignments the assignments property
func (m *PrivilegedRoleItemRequestBuilder) Assignments()(*i30870af67a7c7e750abaa30f5e84eff1a098db84ecc2ae53c8193d533bc3e771.AssignmentsRequestBuilder) {
    return i30870af67a7c7e750abaa30f5e84eff1a098db84ecc2ae53c8193d533bc3e771.NewAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.privilegedRoles.item.assignments.item collection
func (m *PrivilegedRoleItemRequestBuilder) AssignmentsById(id string)(*i0aae3989cf245a50ac6c762e9583dab23911e23dcb918057ff75fd93264ed4ca.PrivilegedRoleAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["privilegedRoleAssignment_id"] = id
    }
    return i0aae3989cf245a50ac6c762e9583dab23911e23dcb918057ff75fd93264ed4ca.NewPrivilegedRoleAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewPrivilegedRoleItemRequestBuilderInternal instantiates a new PrivilegedRoleItemRequestBuilder and sets the default values.
func NewPrivilegedRoleItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrivilegedRoleItemRequestBuilder) {
    m := &PrivilegedRoleItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/privilegedRoles/{privilegedRole_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewPrivilegedRoleItemRequestBuilder instantiates a new PrivilegedRoleItemRequestBuilder and sets the default values.
func NewPrivilegedRoleItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrivilegedRoleItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPrivilegedRoleItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete entity from privilegedRoles
func (m *PrivilegedRoleItemRequestBuilder) CreateDeleteRequestInformation(options *PrivilegedRoleItemRequestBuilderDeleteOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation get entity from privilegedRoles by key
func (m *PrivilegedRoleItemRequestBuilder) CreateGetRequestInformation(options *PrivilegedRoleItemRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update entity in privilegedRoles
func (m *PrivilegedRoleItemRequestBuilder) CreatePatchRequestInformation(options *PrivilegedRoleItemRequestBuilderPatchOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete entity from privilegedRoles
func (m *PrivilegedRoleItemRequestBuilder) Delete(options *PrivilegedRoleItemRequestBuilderDeleteOptions)(error) {
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
// Get get entity from privilegedRoles by key
func (m *PrivilegedRoleItemRequestBuilder) Get(options *PrivilegedRoleItemRequestBuilderGetOptions)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PrivilegedRoleable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePrivilegedRoleFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PrivilegedRoleable), nil
}
// Patch update entity in privilegedRoles
func (m *PrivilegedRoleItemRequestBuilder) Patch(options *PrivilegedRoleItemRequestBuilderPatchOptions)(error) {
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
// SelfActivate the selfActivate property
func (m *PrivilegedRoleItemRequestBuilder) SelfActivate()(*ifc2ecb6356f4c50497a05265b4b0a311af52bfc4aa6302f476a92d1df6adfafb.SelfActivateRequestBuilder) {
    return ifc2ecb6356f4c50497a05265b4b0a311af52bfc4aa6302f476a92d1df6adfafb.NewSelfActivateRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SelfDeactivate the selfDeactivate property
func (m *PrivilegedRoleItemRequestBuilder) SelfDeactivate()(*ib4b9c3bdee734abec86c1274b61654a23bf76937e3c1b4deb6b3a5baf179dcbe.SelfDeactivateRequestBuilder) {
    return ib4b9c3bdee734abec86c1274b61654a23bf76937e3c1b4deb6b3a5baf179dcbe.NewSelfDeactivateRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Settings the settings property
func (m *PrivilegedRoleItemRequestBuilder) Settings()(*i04754c5841c9fcefc803d8096b9d648d4a4f8e797cd7d5703b0b2fa021349d4f.SettingsRequestBuilder) {
    return i04754c5841c9fcefc803d8096b9d648d4a4f8e797cd7d5703b0b2fa021349d4f.NewSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Summary the summary property
func (m *PrivilegedRoleItemRequestBuilder) Summary()(*ib362bf055682cb60eebaecadb64541d74da0307dbd607faad741d862fab711bf.SummaryRequestBuilder) {
    return ib362bf055682cb60eebaecadb64541d74da0307dbd607faad741d862fab711bf.NewSummaryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
