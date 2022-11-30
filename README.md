# CRM Backend

The project represents the backend of a customer relationship management (CRM) web application, realised as final project for the [Go course at Udacity](https://www.udacity.com/course/golang--cd11970).Users can interact with the application by simply make API requests using [Postman](https://www.postman.com/) or [cURL](https://curl.se/).

The API handles the following 5 operations:

- (GET) Getting a single customer through a **/customers/{id}** path
- (GET) Getting all customers through a the **/customers** path
- (POST) Creating a customer through a **/customers** path
- (PATCH) Updating a customer through a **/customers/{id}** path
- (DELETE) Deleting a customer through a **/customers/{id}** path

## Project Set-up

The application use as router [Gin Web Framework](https://gin-gonic.com/).

The project requires only a simple:

- 'go get' command to install the package.
- 'go run' command to launch the application.
