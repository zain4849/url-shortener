# URL Shortener

## Overview
URL Shortener is a simple web application that allows users to shorten long URLs into shorter, more manageable links. It consists of a backend written in Go (echo framework, gorm) with sqlite as the database and a frontend written in React (Typescript) with Vite.

## Use Cases
### Social Media
- Instead of: Check out this amazing blog post on the latest tech trends! https://www.example.com/blog/2024/06/01/the-latest-tech-trends-and-how-they-are-changing-the-world
- Get: Check out this amazing blog post on the latest tech trends! https://localhost:8080/abcdefg (will be a randomly generated key)

### Email Marketing
- Instead of: Don't miss out on our latest product! Click here to learn more: https://www.example.com/products/newest-gadget-2024?utm_source=newsletter&utm_medium=email&utm_campaign=June2024
- Get: Don't miss out on our latest product! Click here to learn more: https://bit.ly/new-gadget (will be a randomly generated key)

## Features
- Shorten long URLs into concise, easy-to-share links.
- Redirect users from short links to the original long URLs.

## Technologies Used
- **Backend:** Go, Echo framework, GORM for ORM
- **Frontend:** React, Vite for fast development, Axios for HTTP requests
- **Database:** SQLite for storing shortened URLs
- **Containerization:** Docker, Docker Compose for managing containers
- **Dev Tools:** Air for live reloading during development

## Installation and Setup
1. Clone the repository to your local machine.
2. Navigate to the root directory and run `docker-compose up` to start the application.
4. Access the application at [http://localhost:5173](http://localhost:5173) in your web browser. 

## Contributing
Contributions are welcome! If you'd like to contribute to the project, please follow these steps:
1. Fork the repository.
2. Create a new branch for your feature or bug fix.
3. Make your changes and commit them to your branch.
4. Push to your fork and submit a pull request.

## License
This project is licensed under the MIT License.
