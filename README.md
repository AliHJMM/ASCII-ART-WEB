# ASCII ART WEB

## DESCRIPTION
ASCII Art Web is a web-application that allows users to generate ASCII ART from text using different banners. Users can input text, select a banner style, and get the ASCII ART representation of the text.

### AUTHORS
- Ali Hasan Jasim (https://github.com/AliHJMM)
- Habib Mansoor (https://github.com/7abib04)
- Husain Ali (https://github.com/hujaafar)

## Usage

### How to Run
1. Clone the Repo
2. Navigate to the project directory
3. Run `go run main.go`
4. Open `http://localhost:8080` in your web browser.

### Implementation Details

#### Algorithm

The ASCII ART generation is done by reading banner templates from text files. Each banner template file contains ASCII representations for characteristics. The server reads these files, maps characters to their ASCII ART representation, and then constructs the ASCII ART for the given input text.

#### HTTP Endpoints

- `GET /` : Serves the main HTML page where users can input text and select a banner.
- `POST /ascii-art` : Processes the input text and banner selection, generates the ASCII ART, and displays the result.


#### HTTP Status Codes

- `200 OK`: The Request was Successful
- `400 Bad Request`: The Request was invalid (e.g., missing text or banner).
- `404 Not Found`: The Requested resources was not found.
- `500 Internal Server Error` : An Error occured on the server.
