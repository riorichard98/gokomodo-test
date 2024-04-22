# Running the App

## With Docker

### Prerequisites
- Docker installed on your system

### Instructions
1. Clone the repository:
    ```bash
    git clone https://github.com/riorichard98/gokomodo-test
    cd gokomodo-test
    ```

2. Run the following command to build and start the app with Docker:
    ```bash
    make run-all
    ```

---

## Without Docker

### Prerequisites
- Go installed on your system
- PostgreSQL installed and running locally

### Instructions
1. Clone the repository:
    ```bash
    git clone https://github.com/riorichard98/gokomodo-test
    cd gokomodo-test
    ```

2. **(Optional)** Run PostgreSQL in a Docker container:
    ```bash
    make run_postgres
    ```

3. Run the following command to ensure dependencies are in sync:
    ```bash
    go mod tidy
    ```

4. Create a PostgreSQL database with the following details:
    - **Database Name:** shop
    - **Username:** gokomodo
    - **Password:** 12345
    - **Host:** localhost
    - **Port:** 5432

5. Initialize the database with the required schema using the provided `init.sql` file.

6. **Login Credentials:**
    - **Seller:**
        - Email: rio@gmail.com
        - Name: rio
        - Password: 123
        - Address: 123 Main Street, Anytown, USA

    - **Buyer:**
        - Email: rich@gmail.com
        - Name: richard
        - Password: 456
        - Address: 456 Elm Street, Anytown, USA

7. Finally, run the application:
    ```bash
    go run main.go
    ```

---
