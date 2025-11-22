App to track stats for life goals
Purpose: 
- Collect data about myself to discover patterns and trends.
- Map out progress in a more visual way.
- Create a powerful journal. 
- Build more personalized tools to assist with achieving goals.

Planned Features:
- Flexible/dynamic journal
    - Templates
    - Changeable fields
- Morning "newsletter"
    - Goal/Habit focus reminders

test Post curl: curl -X POST http://localhost:8080/tests -H "Content-Type: application/json" -d "{\"name\":\"Meditate\"}"
journal Post curl: curl -X POST http://localhost:8080/journal -H "Content-Type: application/json" -d "{ \"user_id\": \"your mom\", \"category\": \"ooga booga\", \"data\": {\"sleep\": \"1234\", \"notes\": \"sup\"} }"
