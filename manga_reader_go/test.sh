curl -X POST http://localhost:4000/user/register \
  -H "Content-Type: application/json" \
  -d '{"username": "John Doe", "email": "john@example.com", "password": "thisispassword"}'
