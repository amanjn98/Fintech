# Fintech user app 

App enables a customer to create and use his account using emailid and password as per requirements.

Following are the command to run this app:

1)docker compose up 

Following are the features available with this app


## SignUp

POST https://localhost:8080/signup

Please use the above url to create account with the correct emailid and password

Sample example :

'curl --location --request POST 'localhost:8080/signup' \
--header 'Content-Type: application/json' \
--data-raw '{
"username": "john.doe@gmail.com",
"password": "newYork@123"
}' 

## LogIn

POST https://localhost:8080/login

Please use the above url to login into the account with emailid and password 

Sample example :

curl --location --request POST 'localhost:8080/login' \
--header 'Content-Type: application/json' \
--data-raw '{
"username": "john.doe@gmail.com",
"password": "newYork@123"
}'
## Reset Password

POST http://localhost:8080/reset

Please use the above url to reset the account with emailid and password

Sample example :

curl --location --request POST 'localhost:8080/login' \
--header 'Content-Type: application/json' \
--data-raw '{
"username": "john.doe@gmail.com",
"password": "newYork@567"
}'
