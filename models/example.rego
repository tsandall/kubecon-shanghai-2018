# Pet clinic authorization policy (in English):
#
#   - Pet owners can read and create pets.
#   - Pet veterinarians can read and update pet details.
#
# Pet clinic has two users:
#
#   - Bob is a pet owner.
#   - Alice is a veterinarian.
#
# Pet clinic service provides input to policy:
#
#   - resource : type of resource (e.g., pets, pet_details)
#   - action   : name of operation (e.g., read, update, create)
#   - user     : identity of caller (e.g., bob, alice)
#
# Example Request:
#
#   GET /pets HTTP/1.1
#   Authorization: bob
#
# Example Input:
#
#   action: read
#   resource: pets
#   user: bob
package petclinic

import data.petclinic.acls
import data.petclinic.rbac

default allow = false

allow {
    acls.allow
}

# deny {
#     not net.cidr_overlap("10.1.1.0/24", input.sourceIP)
#     input.resource = "pet_details"
# }