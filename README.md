# TSIS_1
# Go Gorilla Mux App

This is a simple web application written in Go using Gorilla Mux that provides information about Harry Potter characters.

## Features

- Provides a list of Harry Potter characters.
- Allows retrieving details of a particular character by name.
- Includes a health check endpoint.

## Installation

1. Ensure you have Go installed. If not, you can download and install it from [here](https://golang.org/dl/).
2. Clone this repository to your local machine:

    ```bash
    git clone https://github.com/PP221B030727/TSIS_1.git
    ```

3. Navigate to the project directory:

    ```bash
    cd TSIS_1
    ```

4. Run the following command to start the server:

    ```bash
    go run main.go
    ```
   or
    ```bash
    make server
    ```

## Usage

- To get the list of all characters, send a GET request to `/characters`.
- To get details of a specific character, send a GET request to `/characters/{name}`, where `{name}` is the name of the character in lowercase.
- To perform a health check, send a GET request to `/health`.

## Project Structure

The project structure follows the Standard Go Project Layout.
- `api`: Contains the main entry point of the application.
- `main.go`: The entry point to our program.
- `handlers/`: Contains HTTP request handlers.
- `data/`: Contains JSON data of Harry Potter characters.
- `README.md`: The document you are currently reading.

## Dependencies

- Gorilla Mux: A powerful URL router and dispatcher for Go.

## Contributing

Contributions are welcome! Feel free to open issues and pull requests.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
