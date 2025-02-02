# Default API URL if not provided
API_URL=${1:-"http://localhost:8080"}

# Colors for output
GREEN='\033[0;32m'
RED='\033[0;31m'
NC='\033[0m'

# Function to test API endpoint
test_endpoint() {
  local number=$1
  local expected_status=$2
  local description=$3

  echo -e "\nTesting: $description"
  response=$(curl -s -w "\n%{http_code}" "$API_URL/api/classify-number?number=$number")
  status=$(echo "$response" | tail -n1)
  body=$(echo "$response" | sed '$d')

  if [ "$status" -eq "$expected_status" ]; then
    echo -e "${GREEN}✓ Status code matches expected: $status${NC}"
    echo "Response:"
    echo "$body" | jq '.'
  else
    echo -e "${RED}✗ Status code mismatch. Expected: $expected_status, Got: $status${NC}"
    echo "Response:"
    echo "$body"
  fi
}

# Test health endpoint
echo "Testing health endpoint..."
health_response=$(curl -s -w "\n%{http_code}" "$API_URL/health")
health_status=$(echo "$health_response" | tail -n1)
if [ "$health_status" -eq 200 ]; then
  echo -e "${GREEN}✓ Health check passed${NC}"
else
  echo -e "${RED}✗ Health check failed${NC}"
fi

# Test cases
test_endpoint "371" 200 "Armstrong number (371)"
test_endpoint "6" 200 "Perfect number (6)"
test_endpoint "17" 200 "Prime number (17)"
test_endpoint "42" 200 "Regular number (42)"
test_endpoint "-5" 200 "Negative number (-5)"
test_endpoint "0" 200 "Zero (0)"
test_endpoint "abc" 400 "Invalid input"
