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

## API Endpoints

1. **User Registration**
   - `POST /api/v1/signup`
   - Request Body: `{ "username": "example", "password": "example", "email": "example@example.com" }`
   - Expected Response: `{"message": "User registered successfully"}`

2. **User Login**
   - `POST /api/v1/login`
   - Request Body: `{ "username": "example", "password": "example" }`
   - Expected Response: `{"token": "jwt_token"}`

3. **Create Profile**
   - `POST /api/v1/profiles`
   - Authorization: Bearer Token (obtained from login response)
   - Request Body: `{ "name": "John Doe", "age": 30, "bio": "Love to travel and meet new people.", "photo_url": "http://example.com/photo.jpg" }`
   - Expected Response: `{"message": "Profile created successfully", "profile": profile}`

4. **View Profiles**
   - `GET /api/v1/stack-profiles`
   - Authorization: Bearer Token (obtained from login response)
   - Expected Response: List of 10 profiles that the user hasn't swiped on today.

5. **Swipe Profile**
   - `POST /api/v1/swipe`
   - Authorization: Bearer Token (obtained from login response)
   - Request Body: `{ "profile_id": 1, "swipe_type": "like" }`
   - Expected Response: `{"message": "Swipe recorded successfully"}` or appropriate error messages for duplicate swipes or exceeding daily limit.

6. **Purchase Premium**
   - `POST /api/v1/premium`
   - Authorization: Bearer Token (obtained from login response)
   - Request Body: `{ "package_type": "no_swipe_quota" }`
   - Expected Response: `{"message": "Premium package purchased successfully"}`

7. **Update Profile**
   - `PUT /api/v1/profiles`
   - Authorization: Bearer Token (obtained from login response)
   - Request Body: `{ "name": "John Doe Updated", "age": 31, "bio": "Updated bio.", "photo_url": "http://example.com/new_photo.jpg" }`
   - Expected Response: `{"message": "Profile updated successfully", "profile": profile}`

8. **Liked Profiles**
   - `GET /api/v1/profiles`
   - Authorization: Bearer Token (obtained from login response)
   - Expected Response: List of 10 profiles that the user hasn't swiped on today.