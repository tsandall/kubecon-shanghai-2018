package petclinic.acls

default allow = false

allow {
    input.action = "read"
    input.resource = "pets"
    input.user = "bob"
}

allow {
    input.action = "create"
    input.resource = "pets"
    input.user = "bob"
}

allow {
    input.action = "read"
    input.resource = "pet_details"
    input.user = "alice"
}

allow {
    input.action = "update"
    input.resource = "pet_details"
    input.user = "alice"
}
