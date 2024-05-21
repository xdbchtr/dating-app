# Dating App Backend

This project is a backend system for a Dating Mobile App, similar to Tinder/Bumble, designed using Golang. The project includes functionalities for user registration, login, daily swipe quota, and premium feature purchases.

## Functional Requirements

1. **User Registration & Login**
    - Users can sign up and log in to the app.
    - Passwords must be hashed and stored securely.

2. **Daily Swipe Quota**
    - Users can view, swipe left (pass), or swipe right (like) up to 10 profiles per day.
    - The same profiles should not appear twice in the same day.

3. **Premium Packages**
    - Users can purchase premium features such as:
        - No swipe quota for the user.
        - Verified label for the user.

## Non-Functional Requirements

1. **Scalability**
    - The system should handle a growing number of users without degradation in performance.

2. **Performance**
    - Ensure fast response times for API requests.

3. **Security**
    - Implement proper authentication and data protection measures using JWT.
    - Securely store sensitive user data such as passwords.

4. **Maintainability**
    - Write clean, modular, and maintainable code.
    - Properly document the codebase and provide clear instructions for setup and usage.

## Tech Stack

- **Programming Language**: Go (Golang)
- **Web Framework**: Gin
- **Database**: PostgreSQL
- **ORM**: GORM
- **Authentication**: JWT (JSON Web Tokens)
- **Caching**: Redis (optional, for enhancing performance)
- **Testing**: Go's testing package, Postman for API tests
- **Deployment**: Docker, Kubernetes (optional)
- **Linting**: Golangci-lint

## Project Structure
