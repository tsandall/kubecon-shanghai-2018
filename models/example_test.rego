package petclinic_test

import data.petclinic

test_allow_customer_pet_read {
    petclinic.allow with input as {"action": "read", "resource": "pets", "user": "bob"}        
} 

test_allow_customer_pet_create {
    petclinic.allow with input as {"action": "create", "resource": "pets", "user": "bob"}        
}

test_allow_vet_details_read {
    petclinic.allow with input as {"action": "read", "resource": "pet_details", "user": "alice", "sourceIP": "10.1.1.10"}
}

test_allow_vet_details_update {
    petclinic.allow with input as {"action": "update", "resource": "pet_details", "user": "alice", "sourceIP": "10.1.1.10"}
}

test_deny_customer_details_read {
    not petclinic.allow with input as {"action": "read", "resource": "pet_details", "user": "bob"}        
}

test_deny_customer_details_update {
    not petclinic.allow with input as {"action": "update", "resource": "pet_details", "user": "bob"}        
}

test_deny_external_access_to_pet_details {
    not petclinic.allow with input as {"action": "update", "resource": "pet_details", "user": "alice", "sourceIP": "192.168.1.1"}
}