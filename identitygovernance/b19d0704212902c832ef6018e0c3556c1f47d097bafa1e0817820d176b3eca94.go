package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// EntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageResourceRoleScopesItemAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceScopesItemAccessPackageResourceAccessPackageResourceRolesRequestBuilder provides operations to manage the accessPackageResourceRoles property of the microsoft.graph.accessPackageResource entity.
type EntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageResourceRoleScopesItemAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceScopesItemAccessPackageResourceAccessPackageResourceRolesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageResourceRoleScopesItemAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceScopesItemAccessPackageResourceAccessPackageResourceRolesRequestBuilderGetQueryParameters read-only. Nullable. Supports $expand.
type EntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageResourceRoleScopesItemAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceScopesItemAccessPackageResourceAccessPackageResourceRolesRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// EntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageResourceRoleScopesItemAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceScopesItemAccessPackageResourceAccessPackageResourceRolesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageResourceRoleScopesItemAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceScopesItemAccessPackageResourceAccessPackageResourceRolesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageResourceRoleScopesItemAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceScopesItemAccessPackageResourceAccessPackageResourceRolesRequestBuilderGetQueryParameters
}
// EntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageResourceRoleScopesItemAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceScopesItemAccessPackageResourceAccessPackageResourceRolesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageResourceRoleScopesItemAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceScopesItemAccessPackageResourceAccessPackageResourceRolesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByAccessPackageResourceRoleId provides operations to manage the accessPackageResourceRoles property of the microsoft.graph.accessPackageResource entity.
func (m *EntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageResourceRoleScopesItemAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceScopesItemAccessPackageResourceAccessPackageResourceRolesRequestBuilder) ByAccessPackageResourceRoleId(accessPackageResourceRoleId string)(*EntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageResourceRoleScopesItemAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceScopesItemAccessPackageResourceAccessPackageResourceRolesAccessPackageResourceRoleItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if accessPackageResourceRoleId != "" {
        urlTplParams["accessPackageResourceRole%2Did"] = accessPackageResourceRoleId
    }
    return NewEntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageResourceRoleScopesItemAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceScopesItemAccessPackageResourceAccessPackageResourceRolesAccessPackageResourceRoleItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewEntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageResourceRoleScopesItemAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceScopesItemAccessPackageResourceAccessPackageResourceRolesRequestBuilderInternal instantiates a new AccessPackageResourceRolesRequestBuilder and sets the default values.
func NewEntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageResourceRoleScopesItemAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceScopesItemAccessPackageResourceAccessPackageResourceRolesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageResourceRoleScopesItemAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceScopesItemAccessPackageResourceAccessPackageResourceRolesRequestBuilder) {
    m := &EntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageResourceRoleScopesItemAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceScopesItemAccessPackageResourceAccessPackageResourceRolesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/entitlementManagement/accessPackageAssignments/{accessPackageAssignment%2Did}/accessPackage/accessPackageResourceRoleScopes/{accessPackageResourceRoleScope%2Did}/accessPackageResourceRole/accessPackageResource/accessPackageResourceScopes/{accessPackageResourceScope%2Did}/accessPackageResource/accessPackageResourceRoles{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}", pathParameters),
    }
    return m
}
// NewEntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageResourceRoleScopesItemAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceScopesItemAccessPackageResourceAccessPackageResourceRolesRequestBuilder instantiates a new AccessPackageResourceRolesRequestBuilder and sets the default values.
func NewEntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageResourceRoleScopesItemAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceScopesItemAccessPackageResourceAccessPackageResourceRolesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageResourceRoleScopesItemAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceScopesItemAccessPackageResourceAccessPackageResourceRolesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageResourceRoleScopesItemAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceScopesItemAccessPackageResourceAccessPackageResourceRolesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
func (m *EntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageResourceRoleScopesItemAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceScopesItemAccessPackageResourceAccessPackageResourceRolesRequestBuilder) Count()(*EntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageResourceRoleScopesItemAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceScopesItemAccessPackageResourceAccessPackageResourceRolesCountRequestBuilder) {
    return NewEntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageResourceRoleScopesItemAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceScopesItemAccessPackageResourceAccessPackageResourceRolesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get read-only. Nullable. Supports $expand.
func (m *EntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageResourceRoleScopesItemAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceScopesItemAccessPackageResourceAccessPackageResourceRolesRequestBuilder) Get(ctx context.Context, requestConfiguration *EntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageResourceRoleScopesItemAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceScopesItemAccessPackageResourceAccessPackageResourceRolesRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageResourceRoleCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAccessPackageResourceRoleCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageResourceRoleCollectionResponseable), nil
}
// Post create new navigation property to accessPackageResourceRoles for identityGovernance
func (m *EntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageResourceRoleScopesItemAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceScopesItemAccessPackageResourceAccessPackageResourceRolesRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageResourceRoleable, requestConfiguration *EntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageResourceRoleScopesItemAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceScopesItemAccessPackageResourceAccessPackageResourceRolesRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageResourceRoleable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAccessPackageResourceRoleFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageResourceRoleable), nil
}
// ToGetRequestInformation read-only. Nullable. Supports $expand.
func (m *EntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageResourceRoleScopesItemAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceScopesItemAccessPackageResourceAccessPackageResourceRolesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageResourceRoleScopesItemAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceScopesItemAccessPackageResourceAccessPackageResourceRolesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.TryAdd("Accept", "application/json;q=1")
    return requestInfo, nil
}
// ToPostRequestInformation create new navigation property to accessPackageResourceRoles for identityGovernance
func (m *EntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageResourceRoleScopesItemAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceScopesItemAccessPackageResourceAccessPackageResourceRolesRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageResourceRoleable, requestConfiguration *EntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageResourceRoleScopesItemAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceScopesItemAccessPackageResourceAccessPackageResourceRolesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.TryAdd("Accept", "application/json;q=1")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *EntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageResourceRoleScopesItemAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceScopesItemAccessPackageResourceAccessPackageResourceRolesRequestBuilder) WithUrl(rawUrl string)(*EntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageResourceRoleScopesItemAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceScopesItemAccessPackageResourceAccessPackageResourceRolesRequestBuilder) {
    return NewEntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageResourceRoleScopesItemAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceScopesItemAccessPackageResourceAccessPackageResourceRolesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
