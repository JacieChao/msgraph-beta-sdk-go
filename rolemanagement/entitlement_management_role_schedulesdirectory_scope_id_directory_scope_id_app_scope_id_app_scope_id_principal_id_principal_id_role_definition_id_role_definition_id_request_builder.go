package rolemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// EntitlementManagementRoleSchedulesdirectoryScopeIdDirectoryScopeIdAppScopeIdAppScopeIdPrincipalIdPrincipalIdRoleDefinitionIdRoleDefinitionIdRequestBuilder provides operations to call the roleSchedules method.
type EntitlementManagementRoleSchedulesdirectoryScopeIdDirectoryScopeIdAppScopeIdAppScopeIdPrincipalIdPrincipalIdRoleDefinitionIdRoleDefinitionIdRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EntitlementManagementRoleSchedulesdirectoryScopeIdDirectoryScopeIdAppScopeIdAppScopeIdPrincipalIdPrincipalIdRoleDefinitionIdRoleDefinitionIdRequestBuilderGetQueryParameters invoke function roleSchedules
type EntitlementManagementRoleSchedulesdirectoryScopeIdDirectoryScopeIdAppScopeIdAppScopeIdPrincipalIdPrincipalIdRoleDefinitionIdRoleDefinitionIdRequestBuilderGetQueryParameters struct {
    // Usage: appScopeId='@appScopeId'
    AppScopeId *string `uriparametername:"appScopeId"`
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Usage: directoryScopeId='@directoryScopeId'
    DirectoryScopeId *string `uriparametername:"directoryScopeId"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Usage: principalId='@principalId'
    PrincipalId *string `uriparametername:"principalId"`
    // Usage: roleDefinitionId='@roleDefinitionId'
    RoleDefinitionId *string `uriparametername:"roleDefinitionId"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// EntitlementManagementRoleSchedulesdirectoryScopeIdDirectoryScopeIdAppScopeIdAppScopeIdPrincipalIdPrincipalIdRoleDefinitionIdRoleDefinitionIdRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementManagementRoleSchedulesdirectoryScopeIdDirectoryScopeIdAppScopeIdAppScopeIdPrincipalIdPrincipalIdRoleDefinitionIdRoleDefinitionIdRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EntitlementManagementRoleSchedulesdirectoryScopeIdDirectoryScopeIdAppScopeIdAppScopeIdPrincipalIdPrincipalIdRoleDefinitionIdRoleDefinitionIdRequestBuilderGetQueryParameters
}
// NewEntitlementManagementRoleSchedulesdirectoryScopeIdDirectoryScopeIdAppScopeIdAppScopeIdPrincipalIdPrincipalIdRoleDefinitionIdRoleDefinitionIdRequestBuilderInternal instantiates a new RoleSchedulesdirectoryScopeIdDirectoryScopeIdAppScopeIdAppScopeIdPrincipalIdPrincipalIdRoleDefinitionIdRoleDefinitionIdRequestBuilder and sets the default values.
func NewEntitlementManagementRoleSchedulesdirectoryScopeIdDirectoryScopeIdAppScopeIdAppScopeIdPrincipalIdPrincipalIdRoleDefinitionIdRoleDefinitionIdRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementManagementRoleSchedulesdirectoryScopeIdDirectoryScopeIdAppScopeIdAppScopeIdPrincipalIdPrincipalIdRoleDefinitionIdRoleDefinitionIdRequestBuilder) {
    m := &EntitlementManagementRoleSchedulesdirectoryScopeIdDirectoryScopeIdAppScopeIdAppScopeIdPrincipalIdPrincipalIdRoleDefinitionIdRoleDefinitionIdRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/roleManagement/entitlementManagement/roleSchedules(directoryScopeId='@directoryScopeId',appScopeId='@appScopeId',principalId='@principalId',roleDefinitionId='@roleDefinitionId'){?directoryScopeId*,appScopeId*,principalId*,roleDefinitionId*,%24top,%24skip,%24search,%24filter,%24count,%24select,%24orderby}", pathParameters),
    }
    return m
}
// NewEntitlementManagementRoleSchedulesdirectoryScopeIdDirectoryScopeIdAppScopeIdAppScopeIdPrincipalIdPrincipalIdRoleDefinitionIdRoleDefinitionIdRequestBuilder instantiates a new RoleSchedulesdirectoryScopeIdDirectoryScopeIdAppScopeIdAppScopeIdPrincipalIdPrincipalIdRoleDefinitionIdRoleDefinitionIdRequestBuilder and sets the default values.
func NewEntitlementManagementRoleSchedulesdirectoryScopeIdDirectoryScopeIdAppScopeIdAppScopeIdPrincipalIdPrincipalIdRoleDefinitionIdRoleDefinitionIdRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementManagementRoleSchedulesdirectoryScopeIdDirectoryScopeIdAppScopeIdAppScopeIdPrincipalIdPrincipalIdRoleDefinitionIdRoleDefinitionIdRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEntitlementManagementRoleSchedulesdirectoryScopeIdDirectoryScopeIdAppScopeIdAppScopeIdPrincipalIdPrincipalIdRoleDefinitionIdRoleDefinitionIdRequestBuilderInternal(urlParams, requestAdapter)
}
// Get invoke function roleSchedules
// Deprecated: This method is obsolete. Use GetAsRoleSchedulesdirectoryScopeIdDirectoryScopeIdAppScopeIdAppScopeIdPrincipalIdPrincipalIdRoleDefinitionIdRoleDefinitionIdGetResponse instead.
func (m *EntitlementManagementRoleSchedulesdirectoryScopeIdDirectoryScopeIdAppScopeIdAppScopeIdPrincipalIdPrincipalIdRoleDefinitionIdRoleDefinitionIdRequestBuilder) Get(ctx context.Context, requestConfiguration *EntitlementManagementRoleSchedulesdirectoryScopeIdDirectoryScopeIdAppScopeIdAppScopeIdPrincipalIdPrincipalIdRoleDefinitionIdRoleDefinitionIdRequestBuilderGetRequestConfiguration)(EntitlementManagementRoleSchedulesdirectoryScopeIdDirectoryScopeIdAppScopeIdAppScopeIdPrincipalIdPrincipalIdRoleDefinitionIdRoleDefinitionIdResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateEntitlementManagementRoleSchedulesdirectoryScopeIdDirectoryScopeIdAppScopeIdAppScopeIdPrincipalIdPrincipalIdRoleDefinitionIdRoleDefinitionIdResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(EntitlementManagementRoleSchedulesdirectoryScopeIdDirectoryScopeIdAppScopeIdAppScopeIdPrincipalIdPrincipalIdRoleDefinitionIdRoleDefinitionIdResponseable), nil
}
// GetAsRoleSchedulesdirectoryScopeIdDirectoryScopeIdAppScopeIdAppScopeIdPrincipalIdPrincipalIdRoleDefinitionIdRoleDefinitionIdGetResponse invoke function roleSchedules
func (m *EntitlementManagementRoleSchedulesdirectoryScopeIdDirectoryScopeIdAppScopeIdAppScopeIdPrincipalIdPrincipalIdRoleDefinitionIdRoleDefinitionIdRequestBuilder) GetAsRoleSchedulesdirectoryScopeIdDirectoryScopeIdAppScopeIdAppScopeIdPrincipalIdPrincipalIdRoleDefinitionIdRoleDefinitionIdGetResponse(ctx context.Context, requestConfiguration *EntitlementManagementRoleSchedulesdirectoryScopeIdDirectoryScopeIdAppScopeIdAppScopeIdPrincipalIdPrincipalIdRoleDefinitionIdRoleDefinitionIdRequestBuilderGetRequestConfiguration)(EntitlementManagementRoleSchedulesdirectoryScopeIdDirectoryScopeIdAppScopeIdAppScopeIdPrincipalIdPrincipalIdRoleDefinitionIdRoleDefinitionIdGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateEntitlementManagementRoleSchedulesdirectoryScopeIdDirectoryScopeIdAppScopeIdAppScopeIdPrincipalIdPrincipalIdRoleDefinitionIdRoleDefinitionIdGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(EntitlementManagementRoleSchedulesdirectoryScopeIdDirectoryScopeIdAppScopeIdAppScopeIdPrincipalIdPrincipalIdRoleDefinitionIdRoleDefinitionIdGetResponseable), nil
}
// ToGetRequestInformation invoke function roleSchedules
func (m *EntitlementManagementRoleSchedulesdirectoryScopeIdDirectoryScopeIdAppScopeIdAppScopeIdPrincipalIdPrincipalIdRoleDefinitionIdRoleDefinitionIdRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EntitlementManagementRoleSchedulesdirectoryScopeIdDirectoryScopeIdAppScopeIdAppScopeIdPrincipalIdPrincipalIdRoleDefinitionIdRoleDefinitionIdRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *EntitlementManagementRoleSchedulesdirectoryScopeIdDirectoryScopeIdAppScopeIdAppScopeIdPrincipalIdPrincipalIdRoleDefinitionIdRoleDefinitionIdRequestBuilder) WithUrl(rawUrl string)(*EntitlementManagementRoleSchedulesdirectoryScopeIdDirectoryScopeIdAppScopeIdAppScopeIdPrincipalIdPrincipalIdRoleDefinitionIdRoleDefinitionIdRequestBuilder) {
    return NewEntitlementManagementRoleSchedulesdirectoryScopeIdDirectoryScopeIdAppScopeIdAppScopeIdPrincipalIdPrincipalIdRoleDefinitionIdRoleDefinitionIdRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
