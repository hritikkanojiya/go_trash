## Environment Keys for Backend Development

#### General

1. `PORT`: Specifies the port number on which your application will listen for incoming connections.
2. `ENVIRONMENT`: Indicates the current environment in which your application is running, such as "development", "production", or "staging".
3. `LOG_LEVEL`: Sets the logging level for your application. You can use values like "debug", "info", "warn", "error", or "fatal" to control the verbosity of log messages.
4. `LIMIT`: Specifies the maximum number of records to fetch per page. This determines the page size or the number of records returned in each paginated request.
5. `OFFSET`: Specifies the number of records to skip or offset before fetching the current page of records. It determines the starting point of the records to be returned.
6. `SORT_BY`: Specifies the field or column name by which the records should be sorted. It defines the sorting criteria for the fetched records.
7. `SORT_ON`: Specifies the sort order for the records. It can be used to indicate whether the sorting should be in ascending or descending order.
8. `SECRET_KEY`: A secret key used for encryption, session management, or signing tokens.

#### MongoDB

1. `MONGODB_URI`: The connection URI for MongoDB.
2. `MONGODB_PORT`: The port number for MongoDB service.
3. `MONGODB_DATABASE`: The name of the MongoDB database.
4. `MONGODB_USERNAME`: The username for authenticating with MongoDB.
5. `MONGODB_PASSWORD`: The password for authenticating with MongoDB.

#### AWS

1. `AWS_REGION`: Sets the region to interact with AWS services.
2. `AWS_ACCESS_KEY_ID`: The AWS access key ID for accessing DynamoDB.
3. `AWS_SECRET_ACCESS_KEY`: The AWS secret access key for accessing DynamoDB.
4. `DYNAMODB_REGION`: The AWS region where DynamoDB is located.
5. `DYNAMODB_TABLE`: The name of the DynamoDB table.

#### PostgreSQL

1. `POSTGRES_HOST`: The hostname or IP address of the PostgreSQL server.
2. `POSTGRES_PORT`: The port number on which the PostgreSQL server is running.
3. `POSTGRES_DATABASE`: The name of the PostgreSQL database.
4. `POSTGRES_USER`: The username for authenticating with PostgreSQL.
5. `POSTGRES_PASSWORD`: The password for authenticating with PostgreSQL.

#### Brevo

1. `BREVO_SMTP_HOST`: The host of the Brevo SMTP server.
2. `BREVO_SMTP_PORT`: The port number of the Brevo SMTP server.
3. `BREVO_SMTP_USERNAME`: The username for authenticating with Brevo SMTP.
4. `BREVO_SMTP_PASSWORD`: The password for authenticating with Brevo SMTP.
5. `BREVO_SMTP_API_KEY`: The API Key accessing Brevo services.

#### GeneticMinds SMTP

1. `SMTP_HOST`: The hostname or IP address of your SMTP mail server.
2. `SMTP_PORT`: The port number for the SMTP server. Common options are 25, 587, or 465 for SSL/TLS.
3. `SMTP_USERNAME`: The username or email address used for authenticating with the SMTP server.
4. `SMTP_PASSWORD`: The password or authentication token for the SMTP server.
5. `SMTP_USE_TLS`: Set this variable to enable TLS (Transport Layer Security) for secure communication with the SMTP server. Set it to `true` or `false`.
6. `SMTP_FROM_ADDRESS`: The email address that will be used as the "From" address for outgoing emails.
7. `SMTP_FROM_NAME`: The name associated with the "From" email address.

#### MSG91

1. `MSG91_API_KEY`: The authentication key for interacting with the MSG91 service.

#### Google

1. `GOOGLE_APPLICATION_CREDENTIALS`: The path to the JSON file containing the Google service account credentials, if using any Google Cloud services.

#### GOOGLE_OAUTH

1. `GOOGLE_CLIENT_ID`: The client ID of your Google OAuth application.
2. `GOOGLE_CLIENT_SECRET`: The client secret of your Google OAuth application.
3. `GOOGLE_REDIRECT_URI`: The redirect URI configured in your Google OAuth application settings.
4. `GOOGLE_AUTH_SCOPES`: The scopes required for accessing Google APIs. These can be space-separated values or a comma-separated list.

#### Razorpay

1. `RAZORPAY_KEY_ID`: The key ID for authenticating requests with the Razorpay API.
2. `RAZORPAY_KEY_SECRET`: The secret key for authenticating requests with the Razorpay API.

#### Paytm

1. `PAYTM_MERCHANT_ID`: The unique identifier for your Paytm merchant account.
2. `PAYTM_MERCHANT_KEY`: The secret key provided by Paytm for secure communication with the Paytm API.
3. `PAYTM_CALLBACK_URL`: The callback URL to receive payment response from Paytm.

#### Cashfree

1. `CASHFREE_APP_ID`: The unique identifier for your Cashfree merchant account.
2. `CASHFREE_SECRET_KEY`: The secret key provided by Cashfree for secure communication with the Cashfree API.
3. `CASHFREE_CALLBACK_URL`: The callback URL to receive payment response from Cashfree.
