package security
type AuditLogUserType int

const (
    REGULAR_AUDITLOGUSERTYPE AuditLogUserType = iota
    RESERVED_AUDITLOGUSERTYPE
    ADMIN_AUDITLOGUSERTYPE
    DCADMIN_AUDITLOGUSERTYPE
    SYSTEM_AUDITLOGUSERTYPE
    APPLICATION_AUDITLOGUSERTYPE
    SERVICEPRINCIPAL_AUDITLOGUSERTYPE
    CUSTOMPOLICY_AUDITLOGUSERTYPE
    SYSTEMPOLICY_AUDITLOGUSERTYPE
    PARTNERTECHNICIAN_AUDITLOGUSERTYPE
    GUEST_AUDITLOGUSERTYPE
    UNKNOWNFUTUREVALUE_AUDITLOGUSERTYPE
)

func (i AuditLogUserType) String() string {
    return []string{"Regular", "Reserved", "Admin", "DcAdmin", "System", "Application", "ServicePrincipal", "CustomPolicy", "SystemPolicy", "PartnerTechnician", "Guest", "unknownFutureValue"}[i]
}
func ParseAuditLogUserType(v string) (any, error) {
    result := REGULAR_AUDITLOGUSERTYPE
    switch v {
        case "Regular":
            result = REGULAR_AUDITLOGUSERTYPE
        case "Reserved":
            result = RESERVED_AUDITLOGUSERTYPE
        case "Admin":
            result = ADMIN_AUDITLOGUSERTYPE
        case "DcAdmin":
            result = DCADMIN_AUDITLOGUSERTYPE
        case "System":
            result = SYSTEM_AUDITLOGUSERTYPE
        case "Application":
            result = APPLICATION_AUDITLOGUSERTYPE
        case "ServicePrincipal":
            result = SERVICEPRINCIPAL_AUDITLOGUSERTYPE
        case "CustomPolicy":
            result = CUSTOMPOLICY_AUDITLOGUSERTYPE
        case "SystemPolicy":
            result = SYSTEMPOLICY_AUDITLOGUSERTYPE
        case "PartnerTechnician":
            result = PARTNERTECHNICIAN_AUDITLOGUSERTYPE
        case "Guest":
            result = GUEST_AUDITLOGUSERTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_AUDITLOGUSERTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeAuditLogUserType(values []AuditLogUserType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AuditLogUserType) isMultiValue() bool {
    return false
}
