# ChoreHound

Simple app to remind people to do their chores.

## Architecture

![Architecture](docs/chorehound.png)

## API Draft

Draft REST API

### CRUD operations for family

TBD: Assume one family per "licensed" user

- GET /users # Get list of all users for the family (for the authenticated user) - includes the calling user
- POST /users # Add a user to the family
- GET /users/[user-id] # Get details of the specified user
- POST /users/[user-id] # Update details for the specified user
- DELETE /users/[user-id] # Remove the user from the family (TBD: how to handle if this is the active/primary user??)

TBD: JSON struct for user (includes contact info)

### CRUD operations for chores
- GET /chores # Get list of all chores (for the authenticated user)
- POST /chores # Create a new chore
- GET /chores/[chore-id]  # Get details of the specified chore
- POST /chores/[chore-id]  # Update details for the specified chore
- DELETE /chores/[chore-id]  # Delete the specified chore

TBD: JSON struct for chore (includes schedule, responsible users)

### Schedule-driven operations
- POST /reminder # Send scheduled reminders; TBD: is this for everything? Should it be chore/user specific??


