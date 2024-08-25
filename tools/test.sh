#!/bin/bash

PORT="8080"
IP="127.0.0.1"
BASE_URL="http://$IP:$PORT"

# Set the data to be posted
LOST_DATA='{"lost_time": "2021-01-01T12:00:00Z", "kind": "book", "location": "Library", "phone_number": "1234567890"}'

# ping pong
echo "Testing ping pong"
curl -X GET "$BASE_URL/ping"


# Test case 1: POST request to create a lost item
echo "Testing POST /lost"
curl -X POST -H "Content-Type: application/json" -d "$LOST_DATA" "$BASE_URL/lost-items"




# # Set the base URL of the API
# BASE_URL="https://api.example.com"

# # Test case 1: GET request to retrieve a resource
# function test_get_resource() {
#     local endpoint="/resource/123"
#     local expected_response="Resource 123"

#     local response=$(curl -s -o /dev/null -w "%{http_code}" "${BASE_URL}${endpoint}")
#     if [[ $response -eq 200 ]]; then
#         echo "PASS: GET $endpoint returned 200 OK"
#     else
#         echo "FAIL: GET $endpoint returned $response"
#     fi

#     local actual_response=$(curl -s "${BASE_URL}${endpoint}")
#     if [[ $actual_response == "$expected_response" ]]; then
#         echo "PASS: GET $endpoint returned expected response"
#     else
#         echo "FAIL: GET $endpoint returned unexpected response: $actual_response"
#     fi
# }

# # Test case 2: POST request to create a resource
# function test_create_resource() {
#     local endpoint="/resource"
#     local payload='{"name": "New Resource"}'
#     local expected_response="Resource created successfully"

#     local response=$(curl -s -o /dev/null -w "%{http_code}" -X POST -d "$payload" "${BASE_URL}${endpoint}")
#     if [[ $response -eq 201 ]]; then
#         echo "PASS: POST $endpoint returned 201 Created"
#     else
#         echo "FAIL: POST $endpoint returned $response"
#     fi

#     local actual_response=$(curl -s -X POST -d "$payload" "${BASE_URL}${endpoint}")
#     if [[ $actual_response == "$expected_response" ]]; then
#         echo "PASS: POST $endpoint returned expected response"
#     else
#         echo "FAIL: POST $endpoint returned unexpected response: $actual_response"
#     fi
# }

# # Run the test cases
# test_get_resource
# test_create_resource