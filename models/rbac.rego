package petclinic.rbac

default allow = false

allow {
    roles_for_user[role_name]
    roles_for_operation[role_name]
}

roles_for_user[role_name] {
    binding = data.bindings[_]
    binding.user = input.user
    role_name = binding.roles[_]
}

roles_for_operation[role_name] {
    role = data.roles[_]
    role_name = role.name
    permission = role.permissions[_]
    permission.action = input.action
    permission.resource = input.resource
}